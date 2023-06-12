package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/go-redis/redis/v9"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/websocket"

	ai "CSBackendTmp/AI"
	"CSBackendTmp/api"
	"CSBackendTmp/cs_redis"
	cswebsocket "CSBackendTmp/cs_websocket"
	"CSBackendTmp/ent"
	"CSBackendTmp/utils"

	"github.com/milvus-io/milvus-sdk-go/v2/client"
)

const enableJWT = true

func monitorMiddleware(ctx iris.Context) {
	if ctx.Path() != "/health_check" {
		var user *ent.User
		uid := "0"
		m := make(map[string]interface{})
		m["path"] = ctx.Path()
		m["ip"] = ctx.RemoteAddr()
		m["status_code"] = strconv.Itoa(ctx.GetStatusCode())
		ctxUser := ctx.User()
		if ctxUser != nil {
			rawUser, err := ctxUser.GetRaw()
			if err != nil {
				log.Println("Get ctx user empty")
			} else if rawUser != nil {
				user = rawUser.(*ent.User)
				uid = strconv.FormatUint(user.ID, 10)
			}
		}

		utils.ReportDataToMixPanel("request", uid, m)
	}

	ctx.Next()
}

func checkAuthMiddleware(ctx iris.Context) {
	user := api.GetCurrentUserData(ctx)
	if user == nil {
		ctx.StopWithJSON(iris.StatusForbidden, iris.Map{"err_msg": "need login", "status": iris.StatusForbidden})
	} else {
		ctx.SetUser(user)
		ctx.Next()
	}
}

func publicAuthMiddleware(ctx iris.Context) {
	tk := ctx.URLParam("tk")
	if tk != "neoworld503" {
		ctx.StopWithJSON(iris.StatusForbidden, iris.Map{"err_msg": "need login", "status": iris.StatusForbidden})
	} else {
		ctx.Next()
	}
}

func adminAuthMiddleware(ctx iris.Context) {
	tk := ctx.URLParam("tk")
	if tk != "nimda" {
		ctx.StopWithJSON(iris.StatusForbidden, iris.Map{"err_msg": "need admin auth", "status": iris.StatusForbidden})
	} else {
		ctx.Next()
	}
}

func resourceLimit(ctx iris.Context) {
	tk := ctx.URLParam("tk")
	if tk != "neoworld503" {
		ctx.StopWithJSON(iris.StatusForbidden, iris.Map{"err_msg": "need login", "status": iris.StatusForbidden})
	} else {
		ctx.Next()
	}
}

func main() {
	app := iris.Default()
	app.Use(iris.Compression)
	api.Init()
	cswebsocket.Init()
	app.Use(api.Sess.Handler(), monitorMiddleware)
	app.UseError(monitorMiddleware)
	tmpl := iris.HTML("./templates", ".html")
	app.RegisterView(tmpl)

	tmpl.Reload(true)

	j := jwt.New(jwt.Config{
		Extractor: jwt.FromParameter("token"),

		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("My Secret"), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	idGen := func(ctx iris.Context) string {
		if userID := ctx.GetHeader("X-UserID"); userID != "" {
			log.Println("Websocket get userID:", userID)
		}
		if username := ctx.GetHeader("X-Username"); username != "" {
			return username
		}

		return websocket.DefaultIDGenerator(ctx)
	}

	websocketRoute := app.Get("/ws", websocket.Handler(cswebsocket.WebsocketServer, idGen))

	if enableJWT {
		websocketRoute.Use(j.Serve)
	}

	app.Get("/", publicAuthMiddleware, func(ctx iris.Context) {
		ctx.ServeFile("./templates/index.html")
	})

	app.Get("/health_check", func(ctx iris.Context) {
		ctx.JSON(utils.OKResp(iris.Map{"health_check": "passed"}))
	})

	app.HandleDir("/assets", iris.Dir("./assets"))
	app.Favicon("./assets")

	app.Post("/upload", api.Upload)

	ApiRouterGroup := app.Party("api/v0/")
	{
		ApiRouterGroup.Post("signin", api.SignIn)
		ApiRouterGroup.Post("signup", api.SignUp)
		ApiRouterGroup.Get("mobile_verify", api.GenMobileVerifyCode)
		ApiRouterGroup.Post("mobile_verify", api.VerifyMobile)
		ApiRouterGroup.Get("login_check_mobile", api.LoginCheckMobile)
		ApiRouterGroup.Get("login_check_pw_format", api.CheckPassword)
		ApiRouterGroup.Get("login_check_birthday", api.CheckBirthday)
		ApiRouterGroup.Get("get_agora_base64", api.GetAgoraBase64)
		ApiRouterGroup.Get("get_agora_channels", api.GetAgoraChannels)
		ApiRouterGroup.Get("get_agora_user", api.GetAgoraUser)
		ApiRouterGroup.Post("push_agora_to_cdn", api.PushAgoraStreamToCDN)
		ApiRouterGroup.Get("converters", api.GetAgoraConverters)
		ApiRouterGroup.Delete("converters", api.DelAgoraConverter)
		ApiRouterGroup.Get("converter", api.GetAgoraConverterStatus)
		ApiRouterGroup.Get("agora_rules", api.GetAgoraRules)
		ApiRouterGroup.Post("agora_rules", api.AppleAgoraRules)
		ApiRouterGroup.Put("agora_rules", api.UpdateAgoraRules)
		ApiRouterGroup.Delete("agora_rules", api.DeleteAgoraRules)
		ApiRouterGroup.Get("rtc/:channelName/:role/:tokenType/:rtcuid", api.GetRtcToken)
		ApiRouterGroup.Get("rtm/:rtmuid", api.GetRtmToken)
		ApiRouterGroup.Get("rte/:channelName/:role/:tokenType/:rtcuid", api.GetRtcRtmToken)
		ApiRouterGroup.Get("rte/:channelName/:role/:tokenType/:rtcuid/:rtmuid", api.GetRtcRtmToken)
		ApiRouterGroup.Get("login/{id}", api.LoginByID)
		ApiRouterGroup.Get("oauth", api.Oauth2Redirect)
		ApiRouterGroup.Get("oauth_callback", api.Oauth2CallBack)
		ApiRouterGroup.Get("terms", api.TermsAndCondition)
		ApiRouterGroup.Get("privacy", api.PrivacyPolicy)
	}

	ApiWithAuthRouterGroup := app.Party("api/v0/")
	ApiWithAuthRouterGroup.Use(checkAuthMiddleware)
	{
		ApiWithAuthRouterGroup.Post("privacy_confirm", api.ConfirmPrivacy)
		ApiWithAuthRouterGroup.Get("privacy_text", api.PrivacyAgreement)
		ApiWithAuthRouterGroup.Get("users", api.GetUsers)
		ApiWithAuthRouterGroup.Get("user/{id}", api.GetUserInfoByID)
		ApiWithAuthRouterGroup.Post("user_thumbnail", api.UpdateUserThumbnail)
		ApiWithAuthRouterGroup.Get("check_dup_name", api.CheckDupUserName)
		ApiWithAuthRouterGroup.Put("user", api.UpdateUserInfo)
		ApiWithAuthRouterGroup.Put("user_info", api.UpdateUserInfoSingle)
		ApiWithAuthRouterGroup.Get("curr_user", api.GetCurrUser)
		ApiWithAuthRouterGroup.Put("user/collection_switch", api.UpdateUserCollectionSwitch)
		ApiWithAuthRouterGroup.Get("invite_code", api.GenInviteCode)
		ApiWithAuthRouterGroup.Post("invite_code", api.VerifyRegisterInviteCode)
		ApiWithAuthRouterGroup.Get("friends", api.GetFriends)
		ApiWithAuthRouterGroup.Post("friend", api.AddFriend)
		ApiWithAuthRouterGroup.Get("action", api.ActionForFriend)
		ApiWithAuthRouterGroup.Get("invite_friends", api.InviteFriends)
		ApiWithAuthRouterGroup.Get("collections", api.GetCollections)
		ApiWithAuthRouterGroup.Post("collection", api.AddCollection)
		ApiWithAuthRouterGroup.Delete("collection", api.DeleteCollection)
		ApiWithAuthRouterGroup.Post("timedew", api.PostTimeDewDataV1)
		ApiWithAuthRouterGroup.Get("timedews", api.GetUserRelatedTimedew)
		ApiWithAuthRouterGroup.Get("timedew_with_reactions", api.GetTimeDewsReactions)
		ApiWithAuthRouterGroup.Post("timedew_reaction", api.TimeDewReaction)
		ApiWithAuthRouterGroup.Delete("timedew_reaction", api.DeleteTimeDewReaction)
		ApiWithAuthRouterGroup.Get("friend/{id}", api.GetUserInfoByID)
		ApiWithAuthRouterGroup.Post("invite_friendship", api.InviteFriendShip)
		ApiWithAuthRouterGroup.Post("deal_friendship_invitation", api.DealFriendShipInvitation)
		ApiWithAuthRouterGroup.Get("connect_request", api.ConectionRequest)
		ApiWithAuthRouterGroup.Get("get_invited_friend", api.GetFriendsAskedFor)
		ApiWithAuthRouterGroup.Get("get_friendship_request", api.GetFriendsInvitations)
		ApiWithAuthRouterGroup.Get("cards", api.GetCards)
		ApiWithAuthRouterGroup.Get("search_user", api.UserSearchBySystemName)
		ApiWithAuthRouterGroup.Get("add_user", api.UserSearchAddOthers)
		ApiWithAuthRouterGroup.Get("search_field", api.FieldSearch)
		ApiWithAuthRouterGroup.Post("card", api.CreateCard)
		ApiWithAuthRouterGroup.Get("settings", api.GetSettings)
		ApiWithAuthRouterGroup.Put("setting", api.UpdateSetting)
		ApiWithAuthRouterGroup.Post("setting", api.CreateSetting)
		ApiWithAuthRouterGroup.Get("messages", api.GetV1Messages)
		ApiWithAuthRouterGroup.Post("message", api.PostV1Messages)
		ApiWithAuthRouterGroup.Get("csfields", api.GetCSFields)
		ApiWithAuthRouterGroup.Get("csfield/{cs_field_id}", api.GetCSFieldByID)
		ApiWithAuthRouterGroup.Get("csfield_live/{cs_field_id}", api.GetCSFieldLiveByID)
		ApiWithAuthRouterGroup.Post("csfield_live/{cs_field_id}", api.UpdateCSFieldLiveByID)
		ApiWithAuthRouterGroup.Post("csfield", api.CreateCSField)
		ApiWithAuthRouterGroup.Put("csfield_private_level", api.CSFieldPrivateLevel)
		ApiWithAuthRouterGroup.Post("csfieldbyuser", api.CreateCSFieldByUserID)
		ApiWithAuthRouterGroup.Post("signout", api.SignOut)
		ApiWithAuthRouterGroup.Get("masks", api.GetMasks)
		ApiWithAuthRouterGroup.Post("mask", api.CreateMask)
		ApiWithAuthRouterGroup.Get("feedback_items", api.GetFeedbackReason)
		ApiWithAuthRouterGroup.Post("feedback", api.CreateFeedback)
		ApiWithAuthRouterGroup.Get("hidden_users", api.GetHiddenUsers)
		ApiWithAuthRouterGroup.Post("hidden_user", api.HiddenUser)
		ApiWithAuthRouterGroup.Post("recover_hidden_user", api.RecoverHiddenUser)
		ApiWithAuthRouterGroup.Get("search_hidden_users", api.SearchHiddenUsers)
		ApiWithAuthRouterGroup.Post("gen_timedew", api.PostTimeDew)
		ApiWithAuthRouterGroup.Post("invite_join_field", api.InviteJoinCSField)
		ApiWithAuthRouterGroup.Get("change_password_check", api.ChangePWCheck)
		ApiWithAuthRouterGroup.Post("change_password", api.ChangePW)
		ApiWithAuthRouterGroup.Get("invite_friends_to_field", api.InviteToFieldList)
		ApiWithAuthRouterGroup.Delete("break_connection", api.BreakConnection)
		ApiWithAuthRouterGroup.Post("upload_contacts", api.UploadContacts)
	}

	CSPreASpRouterGroup := app.Party("api/v1/")
	{
		CSPreASpRouterGroup.Get("rtc/{channelName}/{role}/{tokenType}/{rtcuid}", api.GetRtcToken)
		CSPreASpRouterGroup.Get("rtm/{rtmuid}", api.GetRtmToken)
		CSPreASpRouterGroup.Get("rte/{channelName}/{role}/{tokenType}/{rtcuid}", api.GetRtcRtmToken)
		CSPreASpRouterGroup.Get("rte/{channelName}/{role}/{tokenType}/{rtcuid}/{rtmuid}", api.GetRtcRtmToken)
		CSPreASpRouterGroup.Get("messages", checkAuthMiddleware, api.GetV1Messages)
		CSPreASpRouterGroup.Post("message", checkAuthMiddleware, api.PostV1Messages)
		CSPreASpRouterGroup.Get("cards", publicAuthMiddleware, api.GetV1Cards)
		CSPreASpRouterGroup.Post("card", publicAuthMiddleware, api.PostV1Cards)
		CSPreASpRouterGroup.Get("delete_card", publicAuthMiddleware, api.DelectV1Card)
		CSPreASpRouterGroup.Get("card_config", publicAuthMiddleware, func(ctx iris.Context) {
			ctx.ServeFile("./templates/cardconfig.html")
		})
		CSPreASpRouterGroup.Get("realState", checkAuthMiddleware, api.GetTimeDews)
		CSPreASpRouterGroup.Post("lifeFlows", checkAuthMiddleware, api.PostTimeDewData)
		CSPreASpRouterGroup.Get("timedews", checkAuthMiddleware, api.GetTimeDews)
		CSPreASpRouterGroup.Post("timedew", checkAuthMiddleware, api.PostTimeDewDataV1)
		CSPreASpRouterGroup.Get("get_rec_tds", checkAuthMiddleware, api.GetRecTimeDews)
		CSPreASpRouterGroup.Get("timedews_web", publicAuthMiddleware, api.GetWebTimeDews)
		CSPreASpRouterGroup.Get("timedew_config", publicAuthMiddleware, api.GetTimeDewConf)
		CSPreASpRouterGroup.Post("timedew_config", publicAuthMiddleware, api.SetTimeDewConf)
		CSPreASpRouterGroup.Get("timedew_offlinegen", publicAuthMiddleware, api.OfflineGen)
		CSPreASpRouterGroup.Get("check_timedew", publicAuthMiddleware, api.CheckTimeDews)
		CSPreASpRouterGroup.Post("upload_prompt", publicAuthMiddleware, api.Upload)
		CSPreASpRouterGroup.Get("try_read_stream", publicAuthMiddleware, api.ConsumeTimedewStream)
		CSPreASpRouterGroup.Put("timedew", publicAuthMiddleware, api.UpdateTimeDew)
		CSPreASpRouterGroup.Post("tmp_pic", publicAuthMiddleware, api.TmpUpload)

		CSPreASpRouterGroup.Post("login", api.Login)
	}

	ApiAdminGroup := app.Party("api/admin/ruleless/")
	ApiAdminGroup.Use(adminAuthMiddleware)
	{
		ApiAdminGroup.Get("users", api.GetAdminUsers)
		ApiAdminGroup.Get("timedews", api.GetAdminTimeDews)
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	AppPort, _ := os.LookupEnv("APP_PORT")

	app.Configure(iris.WithOptimizations, iris.WithLogLevel("info"),
		iris.WithRemoteAddrHeader("X-Real-IP"),
		iris.WithPostMaxMemory(100*iris.MB))
	app.Build()
	// go ConsumeStream("timedew")
	// go ConsumeStream("user")
	// go ConsumeStream("timedew_rawdata_v1")
	app.Listen(":" + AppPort)

}

func ConsumeStream(streamName string) {

	log.Println("ConsumeStream start:", streamName)
	ctx := context.Background()
	cs_redis.RdbClient.XGroupCreate(ctx, streamName, "test2", "0").Err()
	for {
		entries, err := cs_redis.RdbClient.XReadGroup(ctx, &redis.XReadGroupArgs{
			Group:    "test2",
			Consumer: "test1",
			Streams:  []string{streamName, ">"},
			Count:    2,
			Block:    0,
			NoAck:    false,
		}).Result()
		if err != nil {
			log.Println(err)
			return
		} else {
			for i := 0; i < len(entries[0].Messages); i++ {
				messageID := entries[0].Messages[i].ID
				values := entries[0].Messages[i].Values
				item := []byte(values["item"].(string))
				// if streamName == "timedew" {
				// 	var tdent *ent.TimeDew

				// 	err := json.Unmarshal(item, &tdent)
				// 	var labels []string
				// 	if err != nil {
				// 		log.Println("consume redis stream Unmarshal err:", streamName, err.Error())

				// 	} else {
				// 		// var lf api.LifeFlowItem
				// 		// lf.ID = int64(tdent.ID)
				// 		// lf.Content = tdent.GeneratedContent
				// 		// lf.OnwerID = strconv.FormatUint(tdent.UserID, 10)
				// 		// lf.Title = tdent.Edges.Owner.Name
				// 		// lf.TimeStamp = tdent.CreateTime.UnixMilli()

				// 		// lf.PicUrl = tdent.PicURL
				// 		// tdbyte, err := json.Marshal(lf)
				// 		// if err != nil {
				// 		// 	log.Println("consume redis stream Marshal err:", streamName, err.Error())
				// 		// }
				// 		// cswebsocket.BroadcastToClient("td_update", tdbyte)
				// 		prompt := ai.GenTimeDewLabelPrompt(tdent.GeneratedContent)
				// 		gptLabels := ai.GenerateContent(prompt)
				// 		for _, item := range strings.Split(gptLabels, "##") {
				// 			labels = append(labels, item)
				// 		}
				// 		r, err := utils.PushTimeDewToRec(strconv.FormatUint(tdent.ID, 10), []string{"time_dew"}, labels)
				// 		log.Println(r, err)
				// 	}

				// }
				if streamName == "user" {
					var user map[string]interface{}

					err := json.Unmarshal(item, &user)

					if err != nil {
						log.Println("consume redis stream Unmarshal err:", streamName, err.Error())
					} else {
						labels := []string{user["name"].(string)}
						prompt := ai.GenUserProfilePrompt(labels[0])
						gptLabels := ai.GenerateContent(prompt)
						for _, item := range strings.Split(gptLabels, "##") {
							labels = append(labels, item)
						}
						uid := strconv.FormatFloat(user["id"].(float64), 'f', 0, 64)
						r, err := utils.PushUserToRec(uid, labels)
						log.Println(r, err)
						r, err = utils.PushFeedBackToRec(uid, "83", "like")
						log.Println(r, err)
					}
				}
				if streamName == "timedew_rawdata_v1" {
					_, err := client.NewGrpcClient(
						context.Background(),   // ctx
						"192.168.50.169:19530", // addr
					)
					if err != nil {
						log.Println("failed to connect to Milvus:", err.Error())
					}
					// log.Println("milvusClient", milvusClient)
					// has, err := milvusClient.HasCollection(ctx, "test")
					// log.Println("Milvus HasCollection test: ", has, err)
				}
				cs_redis.RdbClient.XAck(ctx, streamName, "test", messageID)
			}
		}
	}
}
