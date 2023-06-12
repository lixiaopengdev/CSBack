package utils

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/bwmarrin/snowflake"
	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
	"github.com/mixpanel/mixpanel-go"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
)

var AWS_S3_REGION = ""
var AWS_S3_BUCKET = ""
var AWS_S3_KeyID = ""
var AWS_S3_SecretKey = ""
var node *snowflake.Node
var DISCORD_ClientID = ""
var DISCORD_ClientSecret = ""
var MIXPANEL_PROJECT_TOKEN = ""
var OauthConf *oauth2.Config
var charset = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var numberset = []byte("0123456789")
var lenN = len(numberset)
var lenC = len(charset)

var MPClient *mixpanel.ApiClient

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	AWS_S3_REGION, _ = os.LookupEnv("AWS_S3_REGION")
	AWS_S3_BUCKET, _ = os.LookupEnv("AWS_S3_BUCKET")
	AWS_S3_KeyID, _ = os.LookupEnv("AWS_S3_KeyID")
	AWS_S3_SecretKey, _ = os.LookupEnv("AWS_S3_SecretKey")
	DISCORD_ClientID, _ = os.LookupEnv("DISCORD_ClientID")
	DISCORD_ClientSecret, _ = os.LookupEnv("DISCORD_ClientSecret")
	MIXPANEL_PROJECT_TOKEN, _ = os.LookupEnv("MIXPANEL_PROJECT_TOKEN")

	// if !AWS_S3_REGIONExists || !AWS_S3_BUCKETExists || !AWS_S3_KeyIDExists || !AWS_S3_SecretKeyExists {
	// 	log.Fatal("FATAL ERROR: ENV not properly configured, check .env file or aws s3 config")
	// }
	node, _ = snowflake.NewNode(1)
	OauthConf = &oauth2.Config{
		ClientID:     DISCORD_ClientID,
		ClientSecret: DISCORD_ClientSecret,
		Scopes:       []string{"webhook.incoming", "identify"},
		Endpoint: oauth2.Endpoint{
			TokenURL: "https://discord.com/api/oauth2/token",
			AuthURL:  "https://discord.com/oauth2/authorize",
		},
	}
	if err != nil {
		log.Println(err)

	}
	if MIXPANEL_PROJECT_TOKEN != "" {
		MPClient = mixpanel.NewApiClient(MIXPANEL_PROJECT_TOKEN)
	} else {
		MPClient = nil
	}

}

func GenID() int64 {
	// Create a new Node with a Node number of 1

	id := node.Generate()

	return id.Int64()
}

func OKResp(obj map[string]interface{}) iris.Map {
	obj["status"] = iris.StatusOK
	return iris.Map(obj)
}

func BadReqResp(obj map[string]interface{}) iris.Map {
	obj["status"] = iris.StatusBadRequest
	return iris.Map(obj)
}

func ServerErrorResp(obj map[string]interface{}) iris.Map {
	obj["status"] = iris.StatusInternalServerError
	return iris.Map(obj)
}

func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

// func SendMonitorData(ctx iris.Context) {
// 	p := influxdb2.NewPointWithMeasurement("webstatus").AddTag("status_code", strconv.Itoa(ctx.GetStatusCode())).AddTag("ip", ctx.RemoteAddr()).AddTag("method", ctx.Method()).AddTag("path", ctx.Path()).AddField("code", ctx.GetStatusCode()).SetTime(time.Now())
// 	db.MonAPI.WritePoint(p)
// }

func StreamToByte(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.Bytes()
}

func UploadFileToS3(ImageBytes []byte, key string) error {
	session, err := session.NewSession(&aws.Config{Region: aws.String(AWS_S3_REGION), Credentials: credentials.NewStaticCredentials(AWS_S3_KeyID, AWS_S3_SecretKey, "")})

	_, err = s3.New(session).PutObject(&s3.PutObjectInput{
		Bucket: aws.String(AWS_S3_BUCKET),
		Key:    aws.String(fmt.Sprintf("images/%s", key)),
		ACL:    aws.String("public-read"),
		Body:   bytes.NewReader(ImageBytes),
	})
	return err
}

type PyTDReq struct {
	UserName              string `json:"user_name"`
	TranscriptionWithName string `json:"transcription_with_name"`
	TimedewID             string `json:"timedew_id"`
	LastPlace             string `json:"last_place"`
}

type UserForRecReq struct {
	Comment   string   `json:"Comment"`
	Labels    []string `json:"Labels"`
	Subscribe []string `json:"Subscribe"`
	UserID    string   `json:"UserId"`
}

type ItemForRecReq struct {
	Categories []string  `json:"Categories"`
	Comment    string    `json:"Comment"`
	IsHidden   bool      `json:"IsHidden"`
	ItemID     string    `json:"ItemId"`
	Labels     []string  `json:"Labels"`
	Timestamp  time.Time `json:"Timestamp"`
}

type FeedBackForRecReq struct {
	Comment      string    `json:"Comment"`
	FeedbackType string    `json:"FeedbackType"`
	ItemID       string    `json:"ItemId"`
	Timestamp    time.Time `json:"Timestamp"`
	UserID       string    `json:"UserId"`
}

func PushUserToRec(userId string, labels []string) (string, error) {
	// Create a Resty Client
	client := resty.New()
	// POST JSON string
	// No need to set content type, if you have client level setting
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody([]UserForRecReq{UserForRecReq{UserID: userId, Labels: labels}}).
		Post("http://192.168.50.193:7259/api/users")

	return resp.String(), err
}

func PushTimeDewToRec(id string, categories, labels []string) (string, error) {
	// Create a Resty Client
	client := resty.New()
	// POST JSON string
	// No need to set content type, if you have client level setting
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(ItemForRecReq{ItemID: id, Categories: categories, Labels: labels, Timestamp: time.Now()}).
		Post("http://192.168.50.193:7259/api/item")

	return resp.String(), err
}

func PushFeedBackToRec(userid, ItemId, feedback string) (string, error) {
	// Create a Resty Client
	client := resty.New()
	// POST JSON string
	// No need to set content type, if you have client level setting
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody([]FeedBackForRecReq{FeedBackForRecReq{FeedbackType: feedback, ItemID: ItemId, UserID: userid, Timestamp: time.Now()}}).
		Put("http://192.168.50.193:7259/api/feedback")

	return resp.String(), err
}

func GetTimeDewsIDByRec(userid string) (string, error) {
	// Create a Resty Client
	client := resty.New()
	// POST JSON string
	// No need to set content type, if you have client level setting
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		Get("http://192.168.50.193:7259//api/recommend/" + userid)

	return resp.String(), err
}

func RequestPyTDProcesser(userName, speech, tdID, last_location string) {
	// Create a Resty Client
	client := resty.New()
	// POST JSON string
	// No need to set content type, if you have client level setting
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(PyTDReq{UserName: userName, TranscriptionWithName: speech, TimedewID: tdID, LastPlace: last_location}).
		Post("http://192.168.50.171:5000/v1/summarization")
	log.Println("RequestPyTDProcesser resp :", resp, err)
}

// n is the length of random string we want to generate
func RandStr(n int, source string) string {

	b := make([]byte, n)
	for i := range b {
		// randomly select 1 character from given charset
		if source == "char" {
			b[i] = charset[rand.Intn(lenC)]
		} else if source == "number" {
			b[i] = numberset[rand.Intn(lenN)]
		}

	}
	return string(b)
}

func CheckMobileFormat(phone string) bool {

	regRuler := "^\\d{9,10}$"

	reg := regexp.MustCompile(regRuler)

	return reg.MatchString(phone)

}

func CheckPasswordFormat(pw string) (bool, []string) {
	result := []string{}
	if len(pw) < 6 {
		result = append(result, "6 characters minimum")
	}

	if len(pw) > 20 {
		result = append(result, "20 characters maximum")
	}
	num := `[0-9]{1}`
	A_Z := `[A-Z]{1}`
	if b, err := regexp.MatchString(num, pw); !b || err != nil {
		result = append(result, "must contain one number")
	}

	if b, err := regexp.MatchString(A_Z, pw); !b || err != nil {
		result = append(result, "must contain one uppercase")
	}
	if len(result) > 0 {
		return false, result
	} else {
		return true, result
	}

}

func CheckAge(birthday string) bool {
	t := time.Now()
	year := t.Year()
	month := t.Month()
	day := t.Day()
	var timeLayoutStr = "20060102"
	st, _ := time.ParseInLocation(timeLayoutStr, birthday, time.UTC)
	curyear := st.Year()
	curmonth := st.Month()
	curday := st.Day()
	log.Println(year, month, day, curyear, curmonth, curday)
	if year-curyear > 13 {
		return true
	} else if year-curyear == 13 {
		if month > curmonth {
			return true
		} else if month == curmonth {
			if day >= curday {
				return true
			}
		}
	}

	return false

}

func GetConstellation(birthday string) (star string) {
	month, err := strconv.ParseInt(birthday[4:6], 10, 0)
	day, err := strconv.ParseInt(birthday[6:], 10, 0)
	if err != nil {
		star = "None"
	}
	if month <= 0 || month >= 13 {
		star = "None"
	}
	if day <= 0 || day >= 32 {
		star = "None"
	}
	if (month == 1 && day >= 20) || (month == 2 && day <= 18) {
		star = "Aquarius"
	}
	if (month == 2 && day >= 19) || (month == 3 && day <= 20) {
		star = "Pisces"
	}
	if (month == 3 && day >= 21) || (month == 4 && day <= 19) {
		star = "Aries"
	}
	if (month == 4 && day >= 20) || (month == 5 && day <= 20) {
		star = "Taurus"
	}
	if (month == 5 && day >= 21) || (month == 6 && day <= 21) {
		star = "Gemini"
	}
	if (month == 6 && day >= 22) || (month == 7 && day <= 22) {
		star = "Cancer"
	}
	if (month == 7 && day >= 23) || (month == 8 && day <= 22) {
		star = "Leo"
	}
	if (month == 8 && day >= 23) || (month == 9 && day <= 22) {
		star = "Virgo"
	}
	if (month == 9 && day >= 23) || (month == 10 && day <= 22) {
		star = "Libra"
	}
	if (month == 10 && day >= 23) || (month == 11 && day <= 21) {
		star = "Scorpio"
	}
	if (month == 11 && day >= 22) || (month == 12 && day <= 21) {
		star = "Sagittarius"
	}
	if (month == 12 && day >= 22) || (month == 1 && day <= 19) {
		star = "Capricorn"
	}

	return star
}

func ReportDataToMixPanel(eventName, userID string, data map[string]interface{}) {
	if MPClient != nil {
		ctx := context.Background()

		if err := MPClient.Track(ctx, []*mixpanel.Event{
			MPClient.NewEvent(eventName, userID, data),
		}); err != nil {
			log.Println(err)
		}
	}

}
