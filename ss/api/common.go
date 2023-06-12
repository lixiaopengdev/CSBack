package api

import (
	ai "CSBackendTmp/AI"
	"CSBackendTmp/agora_utils"
	"CSBackendTmp/cs_redis"
	"CSBackendTmp/db"
	"CSBackendTmp/ent"
	"CSBackendTmp/ent/bundle"
	"CSBackendTmp/ent/mask"
	"CSBackendTmp/ent/user"
	"CSBackendTmp/utils"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"mime"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/AgoraIO-Community/go-tokenbuilder/rtctokenbuilder"
	jsonpatch "github.com/evanphx/json-patch"
	"github.com/go-redis/redis/v9"
	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"golang.org/x/oauth2"
)

var dbClient *ent.Client

var Sess *sessions.Sessions

var rdb *redis.Client

var systemUserCSTeamID, systemUserRulyID uint64

var reportReasons = []string{
	"I just don’t like it",
	"It’s spam",
	"Nudity or sexual activity",
	"Hate speech or symbols",
	"Bullying or harassment",
	"Violence or dangerous organizations",
	"False information",
	"Suicide or self-injury",
	"Sale of illegal or regulated goods",
}

func Init() {
	db.Init()
	dbClient = db.DBClient
	ctx := context.Background()
	systemUserCSTeamID, _ = dbClient.User.Query().Where(user.SystemNameEQ("Come Social Team")).OnlyID(ctx)
	systemUserRulyID, _ = dbClient.User.Query().Where(user.SystemNameEQ("Ruly")).OnlyID(ctx)
	log.Println("systemUserCSTeamID, systemUserRulyID", systemUserCSTeamID, systemUserRulyID)
	cs_redis.Init()
	rdb = cs_redis.RdbClient
	Sess = sessions.New(sessions.Config{
		Cookie:  "_session_id",
		Expires: 0, // defaults to 0: unlimited life. Another good value is: 45 * time.Minute,
	})
	agora_utils.InitAgora()
	ai.InitGPT()
	ai.InitGPT()
	utils.Init()
}

func preUpload(ctx iris.Context, file *multipart.FileHeader) bool {
	if strings.HasPrefix(file.Filename, "=?utf-8?Q") {
		log.Println("Unity filename: ", file.Filename)
		dec := new(mime.WordDecoder)
		name, err := dec.Decode(file.Filename)
		if err != nil {
			log.Println("Unity filename process fail, name=", file.Filename)
		} else {
			log.Println("decode name:", name)
			file.Filename = name
		}
	}
	return true
}

func Upload(ctx iris.Context) {
	uploadType := ctx.PostValue("type")
	versionID := ctx.PostValue("version_id")
	maskID := ctx.PostValue("mask_id")
	platform := ctx.PostValue("platform")
	if platform == "" {
		platform = "iPhone"
	}
	user := GetCurrentUserData(ctx)
	if uploadType == "mask" {
		if user == nil {
			ctx.StopWithJSON(iris.StatusForbidden, iris.Map{"err_msg": "need login", "status": iris.StatusForbidden})
		}
		log.Println(uploadType, versionID, maskID)
	}

	files, _, err := ctx.UploadFormFiles("./assets", preUpload)
	if err != nil {
		log.Println(err)
		ctx.StopWithStatus(iris.StatusInternalServerError)
		return
	}

	var file_urls []string

	for _, file := range files {
		content, _ := file.Open()
		buf := new(bytes.Buffer)
		buf.ReadFrom(content)
		utils.UploadFileToS3(buf.Bytes(), file.Filename)
		file_url := fmt.Sprintf("http://sagemaker-us-west-2-887392381071.s3.us-west-2.amazonaws.com/images/%s", file.Filename)
		file_urls = append(file_urls, file_url)
		if uploadType == "mask" {
			vid, _ := strconv.ParseUint(versionID, 0, 64)
			mid, _ := strconv.ParseUint(maskID, 0, 64)
			createMaskBundlebyID(ctx, mid, vid, platform, file_url)
		}
		if uploadType == "mask_thumbnail" {
			mid, _ := strconv.ParseUint(maskID, 0, 64)
			updateMaskThumbnail(ctx, mid, file_url)
		}
		if uploadType == "prompt" {
			prompt := ReadPromptFile(file_url)
			ctx.JSON(iris.Map{"data": iris.Map{"file_urls": file_urls, "prompt": prompt}, "status": iris.StatusOK})
			return
		}

	}

	ctx.JSON(iris.Map{"data": iris.Map{"file_urls": file_urls}, "status": iris.StatusOK})
}

func Json(ctx iris.Context) {
	ctx.JSON(iris.Map{"message": "hello", "status": iris.StatusOK})
}

func JsonSecure(ctx iris.Context) {
	response := []string{"val1", "val2", "val3"}
	options := iris.JSON{Indent: "", Secure: true}
	ctx.JSON(response, options)

	// Will output: while(1);["val1","val2","val3"]
}

// Use ASCII field to generate ASCII-only JSON
// with escaped non-ASCII characters.
func JsonAscii(ctx iris.Context) {
	response := iris.Map{"lang": "GO-虹膜", "tag": "<br>"}
	options := iris.JSON{Indent: "    ", ASCII: true}
	ctx.JSON(response, options)

	/* Will output:
	   {
	       "lang": "GO-\u8679\u819c",
	       "tag": "\u003cbr\u003e"
	   }
	*/
}

// Normally, JSON replaces special HTML characters with their unicode entities.
// If you want to encode such characters literally,
// you SHOULD set the UnescapeHTML field to true.
func JsonRaw(ctx iris.Context) {
	options := iris.JSON{UnescapeHTML: true}
	ctx.JSON(iris.Map{
		"html": "<b>Hello, world!</b>",
	}, options)

	// Will output: {"html":"<b>Hello, world!</b>"}
}

func JsonStruct(ctx iris.Context) {
	// You also can use a struct.
	var msg struct {
		Name    string `json:"user"`
		Message string
		Number  int
	}
	msg.Name = "Mariah"
	msg.Message = "hello"
	msg.Number = 42
	// Note that msg.Name becomes "user" in the JSON.
	// Will output: {"user": "Mariah", "Message": "hello", "Number": 42}
	ctx.JSON(msg)
}

func Jsonp(ctx iris.Context) {
	ctx.JSONP(iris.Map{"hello": "jsonp"}, iris.JSONP{Callback: "callbackName"})
}

func Markdown(ctx iris.Context) {
	ctx.Markdown([]byte("# Hello Dynamic Markdown -- iris"))
}

func Yaml(ctx iris.Context) {
	ctx.YAML(iris.Map{"message": "hello", "status": iris.StatusOK})
}

func Msgpack(ctx iris.Context) {
	var msg struct {
		Name    string
		Message string
		Number  int
	}
	msg.Name = "Mariah"
	msg.Message = "hello"
	msg.Number = 42

	ctx.MsgPack(msg)
}

func parseRtcParams(ctx iris.Context) (channelName, tokenType, uidStr string, rtmuid string, role rtctokenbuilder.Role, expireTimestamp uint32, err error) {
	// get param values
	channelName = ctx.Params().Get("channelName")
	roleStr := ctx.Params().Get("role")
	tokenType = ctx.Params().Get("tokenType")
	uidStr = ctx.Params().Get("rtcuid")
	rtmuid = ctx.Params().Get("rtmuid")

	if uidStr == "" {
		// If the uid is missing, just set to 0,
		// meaning it allows for any user ID
		uidStr = "0"
	}
	if rtmuid == "" {
		if uidStr == "0" {
			err = fmt.Errorf("Failed to parse rtm user ID. Cannot be empty or \"0\"")
		}
		rtmuid = uidStr
	}

	if roleStr == "publisher" {
		role = rtctokenbuilder.RolePublisher
	} else {
		// Making an assumption that !publisher == subscriber
		role = rtctokenbuilder.RoleSubscriber
	}
	expireTime := ctx.Params().GetInt64Default("expiry", 3600)
	expireTimeInSeconds := uint32(expireTime)
	currentTimestamp := uint32(time.Now().UTC().Unix())
	expireTimestamp = currentTimestamp + expireTimeInSeconds

	return channelName, tokenType, uidStr, rtmuid, role, expireTimestamp, err
}

func parseRtmParams(ctx iris.Context) (uidStr string, expireTimestamp uint32, err error) {

	uidStr = ctx.Params().Get("rtmuid")
	expireTime := ctx.Params().GetInt64Default("expiry", 3600)

	expireTimeInSeconds := uint32(expireTime)
	currentTimestamp := uint32(time.Now().UTC().Unix())
	expireTimestamp = currentTimestamp + expireTimeInSeconds

	return uidStr, expireTimestamp, err
}

func GetRtcToken(ctx iris.Context) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	// get param values
	channelName, _, uidStr, _, role, expireTimestamp, err := parseRtcParams(ctx)

	rtcToken, tokenErr := agora_utils.GenerateRtcToken(channelName, uidStr, "uid", role, expireTimestamp)
	if tokenErr != nil {
		log.Println(tokenErr) // token failed to generate

		errMsg := "Error Generating RTC token - " + tokenErr.Error()
		ctx.JSON(iris.Map{
			"status": iris.StatusBadRequest,
			"error":  errMsg,
		})
	} else {
		ctx.JSON(iris.Map{"data": iris.Map{"rtcToken": rtcToken}, "status": iris.StatusOK})
	}
}

func GetRtmToken(ctx iris.Context) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	rtmuid, expireTimestamp, err := parseRtmParams(ctx)
	if err != nil {
		log.Println(err) // token failed to generate

		errMsg := "Error Generating RTM token - " + err.Error()
		ctx.JSON(iris.Map{
			"status": iris.StatusBadRequest,
			"error":  errMsg,
		})
	}
	rtmToken, tokenErr := agora_utils.GenerateRtmToken(rtmuid, expireTimestamp)
	if tokenErr != nil {
		log.Println(err) // token failed to generate

		errMsg := "Error Generating RTM token - " + err.Error()
		ctx.JSON(iris.Map{
			"status": iris.StatusBadRequest,
			"error":  errMsg,
		})
	} else {
		ctx.JSON(iris.Map{"data": iris.Map{"rtmToken": rtmToken}, "status": iris.StatusOK})
	}
}

func GetRtcRtmToken(ctx iris.Context) {

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	// get param values
	channelName, _, uidStr, rtmuid, role, expireTimestamp, err := parseRtcParams(ctx)

	rtcToken, tokenErr := agora_utils.GenerateRtcToken(channelName, uidStr, "uid", role, expireTimestamp)
	rtmToken, _ := agora_utils.GenerateRtmToken(rtmuid, expireTimestamp)
	if tokenErr != nil {
		log.Println(tokenErr) // token failed to generate

		errMsg := "Error Generating RTC token - " + tokenErr.Error()
		ctx.JSON(iris.Map{
			"status": iris.StatusBadRequest,
			"error":  errMsg,
		})
	} else {
		ctx.JSON(iris.Map{"data": iris.Map{
			"rtcToken": rtcToken,
			"rtmToken": rtmToken,
		}, "status": iris.StatusOK})
	}

}

func GetCurrentUserData(ctx iris.Context) *ent.User {

	var curr_user *ent.User

	ctxUser := ctx.User()
	if ctxUser != nil {
		rawUser, err := ctxUser.GetRaw()
		if err != nil {
			log.Println("Get ctx user empty")
		} else if rawUser != nil {
			curr_user = rawUser.(*ent.User)
			curr_user.IsOnline = true
			return curr_user
		}
	}
	// log.Println("session_id", ctx.GetCookie("_session_id"))
	user_json, _ := rdb.Get(ctx, ctx.GetCookie("_session_id")).Result()
	if user_json == "" {
		return nil
	}

	err := json.Unmarshal([]byte(user_json), &curr_user)
	if err == nil {
		log.Println("user: ", curr_user.ID, curr_user.Name)
		curr_user.IsOnline = true
		dbClient.User.Update().Where(user.ID(curr_user.ID)).SetIsOnline(true).Exec(ctx)
		return curr_user
	} else {
		return nil
	}
}

func JsonPatch(user_id string, cs_field_id string, patchMsg []byte) error {
	ctx := context.Background()
	csFieldID := strings.TrimSpace(cs_field_id)
	data, err := rdb.Get(ctx, fmt.Sprintf("csfl_%s", csFieldID)).Result()
	if err != nil {
		log.Println("JsonPatch failed: ", user_id, cs_field_id, string(patchMsg))
		return err
	} else {
		dataByte := []byte(data)
		patch, err := jsonpatch.DecodePatch(patchMsg)
		if err != nil {
			log.Println("JsonPatch DecodePatch failed: ", err)
			return err
		}

		modified, err := patch.Apply(dataByte)
		if err != nil {
			log.Println("JsonPatch Apply failed: ", err)
			return err
		}

		log.Println("Original document: ", string(dataByte))
		log.Println("Modified document: ", string(modified))
		_, err = rdb.Set(ctx, fmt.Sprintf("csfl_%s", cs_field_id), modified, 0).Result()
		if err != nil {
			log.Println("JsonPatch Save failed: ", err)
			return err
		}
	}
	return nil
}

func JsonMerge(user_id string, cs_field_id string, mergeMsg []byte) ([]byte, error) {
	ctx := context.Background()
	csFieldID := strings.TrimSpace(cs_field_id)
	data, err := rdb.Get(ctx, fmt.Sprintf("csfl_%s", csFieldID)).Result()
	var statusStruct map[string]interface{}
	var mergeMsgStruct map[string]interface{}
	if err != nil {
		log.Println("JsonMerge failed: ", user_id, cs_field_id, string(mergeMsg))
		return nil, err
	} else {

		dataByte := []byte(data)
		err = json.Unmarshal(dataByte, &statusStruct)
		if err != nil {
			log.Println("JsonMerge failed: ", user_id, cs_field_id, data)
			return nil, err
		}
		err = json.Unmarshal(mergeMsg, &mergeMsgStruct)
		if err != nil {
			log.Println("JsonMerge failed: ", user_id, cs_field_id, string(mergeMsg))
			return nil, err
		} else {
			var reqVersion_id = float64(0)
			var versionKeyString = ""
			versionKey := mergeMsgStruct["version_key"]
			if versionKey != nil {
				versionKeyString = versionKey.(string)
			}
			var isUpgradeBool = false
			prop := make(map[string]interface{})

			version_id := statusStruct["version"].(float64)
			propRaw := mergeMsgStruct["prop"]
			if propRaw != nil {
				prop = mergeMsgStruct["prop"].(map[string]interface{})
			} else {
				prop["version"] = version_id + 1
				clearProp, _ := json.Marshal(prop)
				_, err = rdb.Set(ctx, fmt.Sprintf("csfl_%s", cs_field_id), clearProp, 0).Result()
				return mergeMsg, err
			}
			isUpgrade := mergeMsgStruct["upgrade"]
			if isUpgrade != nil {
				isUpgradeBool = isUpgrade.(bool)
			}
			if prop[versionKeyString] != nil {
				reqVersion_id = prop[versionKeyString].(float64)
			}

			log.Println(version_id, reqVersion_id)
			if (isUpgradeBool && reqVersion_id == version_id+1) || (!isUpgradeBool && reqVersion_id == version_id) || versionKey == nil {
				prop["version"] = version_id + 1

				propMsg, err := json.Marshal(prop)
				mergeMsgStruct["prop"] = prop
				mergeMsg, err = json.Marshal(mergeMsgStruct)

				if err != nil {
					log.Println("JsonMerge Marshal failed: ", err.Error())
					return nil, err
				}
				modifideData, err := jsonpatch.MergePatch(dataByte, propMsg)

				if err != nil {
					log.Println("JsonMerge MergePatch failed: ", err.Error())
					return nil, err
				}
				log.Println("Original document: ", string(dataByte))
				log.Println("Modified document: ", string(modifideData))
				_, err = rdb.Set(ctx, fmt.Sprintf("csfl_%s", cs_field_id), modifideData, 0).Result()
				if err != nil {
					log.Println("JsonMerge Save failed: ", err.Error())
					return nil, err
				}
			} else {
				return nil, errors.New("version check failed")
			}
		}
	}
	return mergeMsg, nil
}

func Oauth2Redirect(ctx iris.Context) {
	url := utils.OauthConf.AuthCodeURL("state", oauth2.AccessTypeOnline)
	log.Printf("Visit the URL for the auth dialog: %v", url)

	ctx.Redirect(url)
}

func Oauth2CallBack(ctx iris.Context) {

	code := ctx.URLParam("code")
	// Use the custom HTTP client when requesting a token.
	httpClient := &http.Client{Timeout: 2 * time.Second}
	req_ctx := context.Background()
	req_ctx = context.WithValue(req_ctx, oauth2.HTTPClient, httpClient)

	tok, err := utils.OauthConf.Exchange(ctx, code)
	if err != nil {
		log.Println(err)
		ctx.JSON(iris.Map{"err_msg": err.Error(), "status": iris.StatusInternalServerError})
	}

	client := utils.OauthConf.Client(ctx, tok)
	resp, err := client.Get("https://discord.com/api/v10/users/@me")
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Printf("Visit the URL for the auth dialog: %v", err)
	}
	var respMap map[string]interface{}
	json.Unmarshal(body, &respMap)
	user, err := LoginByOauth(tok.AccessToken, respMap["username"].(string), respMap["id"].(string), "discord")
	if err == nil {
		sess := Sess.Start(ctx)
		user_json, _ := json.Marshal(user)

		sess.SetImmutable("user_id", user.ID)
		sessID := sess.ID()

		_, err := rdb.Set(ctx, sessID, user_json, 0).Result()
		if err != nil {
			ctx.JSON(iris.Map{"err_msg": err.Error(), "status": iris.StatusBadRequest})
			return
		}
		ctx.SetCookieKV("_session_id", sessID)
		ctx.Redirect("http://zingy.land/api/v1/timedews_web?tk=neoworld503")
		// ctx.JSON(iris.Map{"data": iris.Map{"authed": true, "status": iris.StatusOK, "user": user}, "debug": iris.Map{"data": respMap, "tok": tok, "status": iris.StatusOK}, "status": iris.StatusOK})
		return
	} else {
		ctx.JSON(iris.Map{"err_msg": "oauth login failed", "status": iris.StatusBadRequest})
		return
	}
}

func GetMasks(ctx iris.Context) {

	curr_user := GetCurrentUserData(ctx)
	masks, err := dbClient.Mask.Query().Where(mask.HasOwnerWith(user.IDEQ(curr_user.ID))).WithBundle().All(ctx)
	if err != nil {
		ctx.JSON(iris.Map{"err_msg": err.Error(), "status": iris.StatusBadRequest})
		return
	}
	ctx.JSON(iris.Map{"data": masks, "status": iris.StatusOK})
}

func CreateMask(ctx iris.Context) {

	curr_user := GetCurrentUserData(ctx)
	name := strings.TrimSpace(ctx.PostValue("name"))
	mask, err := dbClient.Mask.Create().SetOwnerID(curr_user.ID).SetName(name).SetGUID(strconv.FormatInt(utils.GenID(), 10)).Save(ctx)
	if err != nil {
		ctx.JSON(iris.Map{"err_msg": err.Error(), "status": iris.StatusBadRequest})
		return
	}
	ctx.JSON(iris.Map{"data": mask, "status": iris.StatusOK})
}

func CreateMaskBundle(ctx iris.Context) {
	// curr_user := GetCurrentUserData(ctx)
	versionID := ctx.PostValue("version_id")
	vid, _ := strconv.ParseUint(versionID, 0, 64)
	bundle_url := ctx.PostValue("bundle_url")
	maskID := ctx.PostValue("mask_id")
	mid, _ := strconv.ParseUint(maskID, 0, 64)
	platform := ctx.PostValue("platform")

	bundle := createMaskBundlebyID(ctx, mid, vid, platform, bundle_url)

	ctx.JSON(iris.Map{"data": bundle, "status": iris.StatusOK})
}

func createMaskBundlebyID(ctx iris.Context, maskID, version_id uint64, platform, bundle_url string) (b *ent.Bundle) {
	var err error
	var plat bundle.Platform
	if platform == "iPhone" {
		plat = bundle.PlatformIPhone
	} else {
		plat = bundle.PlatformAndroid
	}
	b, _ = dbClient.Bundle.Query().Where(bundle.MaskID(maskID), bundle.VerionID(version_id), bundle.PlatformEQ(plat)).Only(ctx)
	if b != nil {
		b, err = b.Update().SetBundleURL(bundle_url).Save(ctx)
	} else {
		b, err = dbClient.Bundle.Create().SetMaskID(maskID).SetBundleURL(bundle_url).SetVerionID(version_id).SetPlatform(plat).Save(ctx)
	}
	if err != nil {
		return nil
	} else {
		return b
	}
}

func updateMaskThumbnail(ctx iris.Context, maskID uint64, thumbnai_url string) {
	_, err := dbClient.Mask.Update().Where(mask.ID(maskID)).SetThumbnailURL(thumbnai_url).Save(ctx)
	if err != nil {
		log.Println("updateMaskThumbnail fail, err:", err.Error())
	}
}

func GetAgoraBase64(ctx iris.Context) {
	result := agora_utils.GenAgoraBase64Authorization()
	ctx.JSON(iris.Map{"data": result, "status": iris.StatusOK})
}

func GetAgoraChannels(ctx iris.Context) {
	result := agora_utils.GetAgoraChannels()
	ctx.JSON(iris.Map{"data": result, "status": iris.StatusOK})
}

func GetAgoraUser(ctx iris.Context) {
	channel := ctx.URLParam("channel")
	result := agora_utils.GetAgoraUserinChannel(channel)
	ctx.JSON(iris.Map{"data": result, "status": iris.StatusOK})
}

func PushAgoraStreamToCDN(ctx iris.Context) {
	channel := ctx.PostValue("channel")
	agoraUserID := ctx.PostValue("agoraid")

	result := agora_utils.PushAgoraStreamToCDN(agoraUserID, channel)
	ctx.JSON(iris.Map{"data": result, "status": iris.StatusOK})
}

func GetAgoraConverters(ctx iris.Context) {
	result := agora_utils.GetAgoraConverters()
	ctx.JSON(iris.Map{"data": result, "status": iris.StatusOK})
}

func DelAgoraConverter(ctx iris.Context) {
	result := agora_utils.DeleteAgoraConverter()
	ctx.JSON(iris.Map{"data": result, "status": iris.StatusOK})
}

func GetAgoraConverterStatus(ctx iris.Context) {
	result := agora_utils.GetAgoraConverters()
	ctx.JSON(iris.Map{"data": result, "status": iris.StatusOK})
}

func GetAgoraRules(ctx iris.Context) {
	result := agora_utils.GetAgoraRuls()
	ctx.JSON(iris.Map{"data": result, "status": iris.StatusOK})
}

func AppleAgoraRules(ctx iris.Context) {
	channel := ctx.PostValue("channel")
	agoraUserID := ctx.PostValue("agoraid")
	result := agora_utils.ApplyAgoraRules(channel, agoraUserID, []string{agora_utils.JoinChannel, agora_utils.PublishAudio, agora_utils.PublishVideo})
	ctx.JSON(iris.Map{"data": result, "status": iris.StatusOK})
}

func UpdateAgoraRules(ctx iris.Context) {
	seconds := ctx.PostValue("seconds")
	secondsInt, err := strconv.ParseInt(seconds, 10, 64)
	if err != nil {
		log.Println("err seconds")
		ctx.JSON(iris.Map{"err_msg": err, "status": iris.StatusBadRequest})
		return
	}
	ruleID := ctx.PostValue("rule_id")
	result := agora_utils.UpdateAgoraRules(ruleID, secondsInt)
	ctx.JSON(iris.Map{"data": result, "status": iris.StatusOK})
}

func DeleteAgoraRules(ctx iris.Context) {
	ruleID := ctx.PostValue("rule_id")
	result := agora_utils.DeleteAgoraRules(ruleID)
	ctx.JSON(iris.Map{"data": result, "status": iris.StatusOK})
}

func ConfirmPrivacy(ctx iris.Context) {
	curr_user := GetCurrentUserData(ctx)
	curr_user, err := dbClient.User.UpdateOneID(curr_user.ID).SetNeedPrivacyConfirm(false).Save(ctx)
	if err != nil {
		ctx.JSON(iris.Map{"err_msg": err, "status": iris.StatusBadRequest})
		return
	} else {
		ctx.SetUser(curr_user)
		sessID := ctx.GetCookie("_session_id")
		user_json, _ := json.Marshal(curr_user)

		rdb.Set(ctx, sessID, user_json, 0).Result()

		if err != nil {
			log.Println("failed when update user info: ", err)
			ctx.JSON(iris.Map{"err_msg": err.Error(), "status": iris.StatusBadRequest})
			return
		}

		ctx.JSON(iris.Map{"data": curr_user, "status": iris.StatusOK})
	}
}

func PrivacyAgreement(ctx iris.Context) {

	ctx.JSON(iris.Map{"data": iris.Map{"title": "title", "body": "body"}, "status": iris.StatusOK})
}

func TermsAndCondition(ctx iris.Context) {

	ctx.View("terms.html")
}

func PrivacyPolicy(ctx iris.Context) {
	ctx.View("privacy.html")
}

func SetUserOffline(user_id string) {
	ctx := context.Background()
	userID, _ := strconv.ParseUint(user_id, 10, 64)
	dbClient.User.Update().Where(user.ID(userID)).SetIsOnline(false).Exec(ctx)
}

func UploadContacts(ctx iris.Context) {
	var tdreq TimeDewReqV1
	// data, _ := ctx.GetBody()

	err := ctx.ReadJSON(&tdreq)
	if err != nil {
		ctx.JSON(iris.Map{"data": tdreq, "status": iris.StatusBadRequest})
	}
	// bulk := make([]*ent.Contact, 5)
	// for i, name := range TimeDewReqV1.TDData.Locactions {
	// 	bulk[i] = dbClient.Contact.Create().SetName(name)
	// }
	// dbClient.Contact.CreateBulk()
	ctx.JSON(iris.Map{"data": "ok", "status": iris.StatusOK})
}
