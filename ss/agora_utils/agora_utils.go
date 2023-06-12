package agora_utils

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/AgoraIO-Community/go-tokenbuilder/rtctokenbuilder"
	"github.com/AgoraIO-Community/go-tokenbuilder/rtmtokenbuilder"
	"github.com/joho/godotenv"
)

const (
	JoinChannel  = "join_channel"  // 加入频道
	PublishAudio = "publish_audio" //发送音频流
	PublishVideo = "publish_video" // 发送视频流
)

// Recorder manages cloud recording
type Recorder struct {
	http.Client
	Channel string
	Token   string
	UID     int32
	RID     string
	SID     string
}

type AcquireClientRequest struct {
	ResourceExpiredHour int `json:"resourceExpiredHour,omitempty"`
}

type AcquireRequest struct {
	Cname         string               `json:"cname"`
	UID           string               `json:"uid"`
	ClientRequest AcquireClientRequest `json:"clientRequest"`
}

type TranscodingConfig struct {
	Height           int    `json:"height"`
	Width            int    `json:"width"`
	Bitrate          int    `json:"bitrate"`
	Fps              int    `json:"fps"`
	MixedVideoLayout int    `json:"mixedVideoLayout"`
	MaxResolutionUID string `json:"maxResolutionUid,omitempty"`
	BackgroundColor  string `json:"backgroundColor"`
}

type RecordingConfig struct {
	MaxIdleTime       int               `json:"maxIdleTime"`
	StreamTypes       int               `json:"streamTypes"`
	ChannelType       int               `json:"channelType"`
	DecryptionMode    int               `json:"decryptionMode,omitempty"`
	Secret            string            `json:"secret,omitempty"`
	TranscodingConfig TranscodingConfig `json:"transcodingConfig"`
}

type StorageConfig struct {
	Vendor         int      `json:"vendor"`
	Region         int      `json:"region"`
	Bucket         string   `json:"bucket"`
	AccessKey      string   `json:"accessKey"`
	SecretKey      string   `json:"secretKey"`
	FileNamePrefix []string `json:"fileNamePrefix"`
}

type RecordingFileConfig struct {
	AVFileType []string `json:"avFileType"`
}

type ClientRequest struct {
	Token               string              `json:"token"`
	RecordingConfig     RecordingConfig     `json:"recordingConfig"`
	RecordingFileConfig RecordingFileConfig `json:"recordingFileConfig"`
	StorageConfig       StorageConfig       `json:"storageConfig"`
}

type StartRecordRequest struct {
	Cname         string        `json:"cname"`
	UID           string        `json:"uid"`
	ClientRequest ClientRequest `json:"clientRequest"`
}

type PushRtmpConfig struct {
	Vendor         int      `json:"vendor"`
	Region         int      `json:"region"`
	Bucket         string   `json:"bucket"`
	AccessKey      string   `json:"accessKey"`
	SecretKey      string   `json:"secretKey"`
	FileNamePrefix []string `json:"fileNamePrefix"`
}

var agoraCustomerKey string
var agoraCustomerSecret string
var agoraAppID string
var agoraAppCertificate string
var agoraBase64Credentials string
var agoraProjectsUrl string
var agoraChannelsUrl string
var agoraPushToCDNUrl string
var agoraConvertersUrl string
var agoraKickRuleUrl = "https://api.agora.io/dev/v1/kicking-rule"

func InitAgora() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	AgoraCustomerKey, AgoraCustomerKeyExists := os.LookupEnv("AGORA_CUSTOMER_KEY")
	AgoraCustomerrSecret, AgoraCustomerrSecretExists := os.LookupEnv("AGORA_CUSTOMER_SECRET")
	AgoraAppID, AgoraAppIDExists := os.LookupEnv("AGORA_APP_ID")
	AgoraAppCertificate, AgoraAppCertificateExists := os.LookupEnv("AGORA_APP_CERTIFICATE")
	if !AgoraCustomerKeyExists || !AgoraCustomerrSecretExists || !AgoraAppIDExists || !AgoraAppCertificateExists || len(AgoraCustomerKey) == 0 || len(AgoraCustomerrSecret) == 0 || len(AgoraAppID) == 0 || len(AgoraAppCertificate) == 0 {
		log.Fatal("FATAL ERROR: ENV not properly configured, check .env file or AGORA PARAM GROUP")
	}
	agoraCustomerKey = AgoraCustomerKey
	agoraCustomerSecret = AgoraCustomerrSecret
	agoraAppID = AgoraAppID
	agoraAppCertificate = AgoraAppCertificate
	agoraBase64Credentials = base64.StdEncoding.EncodeToString([]byte(agoraCustomerKey + ":" + agoraCustomerSecret))
	agoraProjectsUrl = "https://api.agora.io/dev/v1/projects"
	agoraChannelsUrl = fmt.Sprintf("https://api.agora.io/dev/v1/channel/%s", AgoraAppID)
	agoraPushToCDNUrl = fmt.Sprintf("https://api.agora.io/na/v1/projects/%s/rtmp-converters", AgoraAppID)
	agoraConvertersUrl = fmt.Sprintf("https://api.agora.io/v1/projects/%s/rtmp-converters", AgoraAppID)

}

func GenAgoraBase64Authorization() string {

	method := "GET"

	payload := strings.NewReader(``)

	client := &http.Client{}
	req, err := http.NewRequest(method, agoraProjectsUrl, payload)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	req.Header.Add("Authorization", "Basic "+agoraBase64Credentials)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Println(string(body))
	return string(body)

}

func GetAgoraChannels() string {

	method := "GET"

	payload := strings.NewReader(``)

	client := &http.Client{}
	req, err := http.NewRequest(method, agoraChannelsUrl, payload)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	// 增加 Authorization header
	req.Header.Add("Authorization", "Basic "+agoraBase64Credentials)
	req.Header.Add("Content-Type", "application/json")

	// 发送 HTTP 请求
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Println(string(body))
	return string(body)

}

func GetAgoraUserinChannel(channel string) string {

	url := fmt.Sprintf("https://api.agora.io/dev/v1/channel/user/%s/%s", agoraAppID, channel)
	method := "GET"

	payload := strings.NewReader(``)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	req.Header.Add("Authorization", "Basic "+agoraBase64Credentials)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Println(string(body))
	return string(body)

}

func PushAgoraStreamToCDN(agoraUserID string, agoraChannel string) string {

	method := "POST"

	rq := make(map[string]interface{})
	rq1 := make(map[string]interface{})
	rq2 := make(map[string]interface{})
	agoraUserIDInt64, _ := strconv.ParseInt(agoraUserID, 10, 64)
	rq2["rtcChannel"] = agoraChannel
	rq2["rtcStreamUid"] = agoraUserIDInt64

	rq1["name"] = "myconverter"
	rq1["rawOptions"] = rq2
	rq1["rtmpUrl"] = "rtmp://ec2-54-147-47-143.compute-1.amazonaws.com:1935/live/test110"
	rq1["idleTimeout"] = 60
	rq["converter"] = rq1
	// rqJson := `{
	// 	"converter":{
	// 		"name": "mysconverter",
	// 		"rawOptions": {
	// 			"rtcChannel": "demo",
	// 			"rtcStreamUid": 1673247400,
	// 		},
	// 		"rtmpUrl":"rtmp://ec2-54-147-47-143.compute-1.amazonaws.com:1935/live/test110",
	// 		"idleTimeout": 60
	// 	}
	// }`
	rq_data, _ := json.Marshal(rq)
	payload := bytes.NewBuffer(rq_data)

	client := &http.Client{}
	req, err := http.NewRequest(method, agoraPushToCDNUrl, payload)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	// 增加 Authorization header
	req.Header.Add("Authorization", "Basic "+agoraBase64Credentials)
	req.Header.Add("Content-Type", "application/json")

	// 发送 HTTP 请求
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Println(string(body))
	return string(body)

}

func GetAgoraConverters() string {

	method := "GET"

	payload := strings.NewReader(``)

	client := &http.Client{}
	req, err := http.NewRequest(method, agoraConvertersUrl, payload)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	// 增加 Authorization header
	req.Header.Add("Authorization", "Basic "+agoraBase64Credentials)
	req.Header.Add("Content-Type", "application/json")

	// 发送 HTTP 请求
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Println(string(body))
	return string(body)

}

func DeleteAgoraConverter() string {

	method := "DELETE"

	payload := strings.NewReader(``)

	client := &http.Client{}
	req, err := http.NewRequest(method, agoraConvertersUrl, payload)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	// 增加 Authorization header
	req.Header.Add("Authorization", "Basic "+agoraBase64Credentials)
	req.Header.Add("Content-Type", "application/json")

	// 发送 HTTP 请求
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Println(string(body))
	return string(body)

}

func GetAgoraRuls() string {

	method := "GET"

	payload := strings.NewReader(``)
	url := fmt.Sprintf("%s?appid=%s", agoraKickRuleUrl, agoraAppID)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	// 增加 Authorization header
	req.Header.Add("Authorization", "Basic "+agoraBase64Credentials)
	req.Header.Add("Content-Type", "application/json")

	// 发送 HTTP 请求
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Println(string(body))
	return string(body)

}

func ApplyAgoraRules(agoraChannel string, agoraUserID string, rules []string) string {

	method := "POST"
	agoraUserIDInt64, _ := strconv.ParseInt(agoraUserID, 10, 64)

	rq := make(map[string]interface{})
	rq["appid"] = agoraAppID
	rq["cname"] = agoraChannel
	rq["uid"] = agoraUserIDInt64
	rq["privileges"] = rules
	rq["time"] = 1              // 封禁时间，单位为分钟，取值范围为 [1,1440]。注意事项:当设置的值在 0 到 1 之间时，服务端自动改设为 1。当设置的值大于 1440 时，服务端自动改设为 1440。当设置的值为 0 时，表示不封禁。服务端会对频道内符合设定规则的用户进行下线一次的操作。用户可以重新登录进入频道。time 和 time_in_seconds 两个参数只需设置其中的一个。如果同时设置，则 time_in_seconds 生效；如果都不设置，服务端会自动将封禁时间设为 60 分钟，即 3600 秒。
	rq["time_in_seconds	"] = 60 // 封禁时间，单位为秒，取值范围为 [10,86430]。注意事项：当设置的值在 0 到 10 之间，服务端自动改设为 10。当设置的值大于 86430 时，服务端自动改设为 86430。当设置的值为 0 时，表示不封禁。服务端会对频道内符合设定规则的用户进行下线一次的操作。用户可以重新登录进入频道。time 和 time_in_seconds 两个参数只需设置其中的一个。如果同时设置，则 time_in_seconds 生效，如果都不设置，服务端会自动将封禁时间设为 60 分钟，即 3600 秒。

	rq_data, _ := json.Marshal(rq)
	payload := bytes.NewBuffer(rq_data)

	client := &http.Client{}
	req, err := http.NewRequest(method, agoraKickRuleUrl, payload)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	// 增加 Authorization header
	req.Header.Add("Authorization", "Basic "+agoraBase64Credentials)
	req.Header.Add("Content-Type", "application/json")

	// 发送 HTTP 请求
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Println(string(body))
	return string(body)

}

func UpdateAgoraRules(ruleID string, time_in_seconds int64) string {

	method := "PUT"

	rq := make(map[string]interface{})
	rq["appid"] = agoraAppID
	rq["id"] = ruleID
	rq["time_in_seconds	"] = time_in_seconds // 封禁时间，单位为秒，取值范围为 [10,86430]。注意事项：当设置的值在 0 到 10 之间，服务端自动改设为 10。当设置的值大于 86430 时，服务端自动改设为 86430。当设置的值为 0 时，表示不封禁。服务端会对频道内符合设定规则的用户进行下线一次的操作。用户可以重新登录进入频道。time 和 time_in_seconds 两个参数只需设置其中的一个。如果同时设置，则 time_in_seconds 生效，如果都不设置，服务端会自动将封禁时间设为 60 分钟，即 3600 秒。

	rq_data, _ := json.Marshal(rq)
	payload := bytes.NewBuffer(rq_data)

	client := &http.Client{}
	req, err := http.NewRequest(method, agoraKickRuleUrl, payload)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	// 增加 Authorization header
	req.Header.Add("Authorization", "Basic "+agoraBase64Credentials)
	req.Header.Add("Content-Type", "application/json")

	// 发送 HTTP 请求
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Println(string(body))
	return string(body)

}

func DeleteAgoraRules(ruleID string) string {

	method := "DELETE"

	rq := make(map[string]interface{})
	rq["appid"] = agoraAppID
	rq["id"] = ruleID

	rq_data, _ := json.Marshal(rq)
	payload := bytes.NewBuffer(rq_data)

	client := &http.Client{}
	req, err := http.NewRequest(method, agoraKickRuleUrl, payload)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	// 增加 Authorization header
	req.Header.Add("Authorization", "Basic "+agoraBase64Credentials)
	req.Header.Add("Content-Type", "application/json")

	// 发送 HTTP 请求
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Println(string(body))
	return string(body)

}

func GenAgoraToken() string {
	var (
		// 输入你生成的 Token
		tokenValue = "input the token value here"
		// 输入你的 AppID
		appID     = "input your app ID here"
		urlstr    = fmt.Sprintf("https://api.agora.io/dev/v2/project/%s/rtm/vendor/user_events", appID)
		authValue = fmt.Sprintf("agora token=%s", tokenValue)
	)

	// 构造 HTTP 请求
	req, err := http.NewRequest(http.MethodGet, urlstr, nil)
	if err != nil {
		log.Println(fmt.Errorf("failed to new http request, %w", err))
	}
	// 设置 Authorization header
	req.Header.Set("Authorization", authValue)
	req.Header.Set("Content-Type", "application/json")

	// 发送 HTTP 请求
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(fmt.Errorf("failed to send request, %w", err))
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(fmt.Errorf("failed to read response body, %w", err))
	}

	// 判定 StatusCode
	if resp.StatusCode/100 != 2 {
		log.Println(fmt.Errorf("StatusCode(%d) != 2xx, %s", resp.StatusCode, string(body)))
	}

	// 打印出正常返回的响应体
	fmt.Println(string(body))
	return string(body)
}

func GenerateRtcToken(channelName, uidStr, tokenType string, role rtctokenbuilder.Role, expireTimestamp uint32) (rtcToken string, err error) {

	if tokenType == "userAccount" {
		log.Printf("Building Token with userAccount: %s\n", uidStr)
		rtcToken, err = rtctokenbuilder.BuildTokenWithUserAccount(agoraAppID, agoraAppCertificate, channelName, uidStr, role, expireTimestamp)
		return rtcToken, err

	} else if tokenType == "uid" {
		uid64, parseErr := strconv.ParseUint(uidStr, 10, 64)
		// check if conversion fails
		if parseErr != nil {
			err = fmt.Errorf("failed to parse uidStr: %s, to uint causing error: %s", uidStr, parseErr)
			return "", err
		}

		uid := uint32(uid64) // convert uid from uint64 to uint 32
		rtcToken, err = rtctokenbuilder.BuildTokenWithUID(agoraAppID, agoraAppCertificate, channelName, uid, role, expireTimestamp)
		return rtcToken, err
	} else {
		err = fmt.Errorf("failed to generate RTC token for Unknown Tokentype: %s", tokenType)
		log.Println(err)
		return "", err
	}
}

func GenerateRtmToken(rtmuid string, expireTimestamp uint32) (rtcToken string, err error) {
	return rtmtokenbuilder.BuildToken(agoraAppID, agoraAppCertificate, rtmuid, rtmtokenbuilder.RoleRtmUser, expireTimestamp)

}
