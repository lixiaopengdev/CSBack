package cswebsocket

import (
	ai "CSBackendTmp/AI"
	"CSBackendTmp/api"
	"CSBackendTmp/cs_redis"
	"CSBackendTmp/db"
	"CSBackendTmp/ent"
	"CSBackendTmp/ent/user"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/websocket"
	"github.com/kataras/neffos"
)

// values should match with the client sides as well.

const namespace = "default"
const websocketAppName = "ws_neoworld"

var WebsocketServer *neffos.Server

// var pongMessage = []byte("pong")

// var events = neffos.Events{
// 	neffos.OnNativeMessage: func(c *neffos.NSConn, msg neffos.Message) error {
// 		log.Printf("Got: %s", string(msg.Body))
// 		log.Printf("Got: %s", msg)

//			if !c.Conn.IsClient() {
//				c.Conn.Socket().WriteText(pongMessage, 0)
//			}
//			c.Conn.Server().Broadcast(c, msg)
//			return nil
//		},
//	}
//
// if namespace is empty then simply websocket.Events{...} can be used instead.
var serverEvents = websocket.Namespaces{
	namespace: websocket.Events{
		websocket.OnNamespaceConnected: func(nsConn *websocket.NSConn, msg websocket.Message) error {
			// with `websocket.GetContext` you can retrieve the Iris' `Context`.
			ctx := websocket.GetContext(nsConn.Conn)
			user_id := ctx.GetHeader("X-UserID")
			log.Printf("[%s] connected to namespace [%s] with IP [%s]",
				user_id, msg.Namespace,
				ctx.RemoteAddr())
			return nil
		},
		websocket.OnNamespaceDisconnect: func(nsConn *websocket.NSConn, msg websocket.Message) error {
			ctx := websocket.GetContext(nsConn.Conn)
			user_id := ctx.GetHeader("X-UserID")
			api.SetUserOffline(user_id)
			log.Printf("[%s] disconnected from namespace [%s]", user_id, msg.Namespace)
			return nil
		},
		"chat": func(nsConn *websocket.NSConn, msg websocket.Message) error {
			ctx := websocket.GetContext(nsConn.Conn)
			user_id := ctx.GetHeader("X-UserID")
			if user_id == "" {
				log.Println("websocket without user_id")
				nsConn.Emit("error", []byte("need user_id from param X-Websocket-Header-X-UserID"))
				return nil
			}
			// room.String() returns -> NSConn.String() returns -> Conn.String() returns -> Conn.ID()

			log.Printf("[%s] sent: %s", user_id, string(msg.Body))
			// Write message back to the client message owner with:
			// nsConn.Emit("chat", msg)
			// Write message to all except this client with:
			nsConn.Conn.Server().Broadcast(nsConn, msg)
			api.ChatByWS(user_id, msg.Room, string(msg.Body))
			return nil
		},
		"json_patch": func(nsConn *websocket.NSConn, msg websocket.Message) error {
			ctx := websocket.GetContext(nsConn.Conn)
			user_id := ctx.GetHeader("X-UserID")
			if user_id == "" {
				log.Println("websocket without user_id")
				nsConn.Emit("error", []byte("need user_id from param X-Websocket-Header-X-UserID"))
				return nil
			}
			log.Printf("[%s] in cs_field [%s]json_patch: %s", user_id, msg.Room, string(msg.Body))

			err := api.JsonPatch(user_id, msg.Room, msg.Body)
			if err != nil {
				nsConn.Emit("error", []byte(err.Error()))
			} else {
				nsConn.Emit("success", []byte("json_patch success"))
				nsConn.Conn.Server().Broadcast(nsConn, msg)
			}
			return nil
		},
		"prop_update": func(nsConn *websocket.NSConn, msg websocket.Message) error {
			ctx := websocket.GetContext(nsConn.Conn)
			user_id := ctx.GetHeader("X-UserID")
			if user_id == "" {
				log.Println("websocket without user_id")
				nsConn.Emit("error", []byte("need user_id from param X-Websocket-Header-X-UserID"))
				return nil
			}
			log.Printf("[%s] in cs_field [%s] prop_update: %s", user_id, msg.Room, string(msg.Body))

			verisonedMsg, err := api.JsonMerge(user_id, msg.Room, msg.Body)

			if err != nil {
				nsConn.Emit("error", []byte(err.Error()))
			} else {
				msg.Body = verisonedMsg
				nsConn.Emit("prop_update", verisonedMsg)
				nsConn.Conn.Server().Broadcast(nsConn, msg)
			}
			return nil
		},
		websocket.OnRoomJoin: func(nsConn *websocket.NSConn, msg websocket.Message) error {
			// with `websocket.GetContext` you can retrieve the Iris' `Context`.
			ctx := websocket.GetContext(nsConn.Conn)
			user_id := ctx.GetHeader("X-UserID")
			if user_id == "" {
				log.Println("websocket without user_id")
				nsConn.Emit("error", []byte("need user_id from param X-Websocket-Header-X-UserID"))
				return nil
			}
			field, curUser := api.JoinInCSFieldByWS(user_id, msg.Room)
			if field == nil {
				nsConn.Emit("error", []byte("field not exist"))
				return nil
			}
			log.Printf("[%s] joining in cs_field [%s] with IP [%s]",
				user_id, msg.Room,
				ctx.RemoteAddr())
			byteUser, _ := json.Marshal(curUser)
			msg.Body = byteUser
			msg.Event = "userJoin"
			nsConn.Conn.Server().Broadcast(nsConn, msg)
			byteField, _ := json.Marshal(field)
			nsConn.Emit("roomInfo", byteField)
			return nil
		},
		websocket.OnRoomJoined: func(nsConn *websocket.NSConn, msg websocket.Message) error {
			// with `websocket.GetContext` you can retrieve the Iris' `Context`.
			ctx := websocket.GetContext(nsConn.Conn)
			user_id := ctx.GetHeader("X-UserID")
			log.Printf("[%s] joined in cs_field [%s] with IP [%s]",
				user_id, msg.Room,
				ctx.RemoteAddr())
			byteUser, _ := json.Marshal((iris.Map{"user_id": user_id}))
			msg.Body = byteUser
			msg.Event = "userJoin"
			nsConn.Conn.Server().Broadcast(nsConn, msg)
			return nil
		},
		websocket.OnRoomLeave: func(nsConn *websocket.NSConn, msg websocket.Message) error {
			// with `websocket.GetContext` you can retrieve the Iris' `Context`.
			ctx := websocket.GetContext(nsConn.Conn)
			user_id := ctx.GetHeader("X-UserID")
			if user_id == "" {
				log.Println("websocket without user_id")
				nsConn.Emit("error", []byte("need user_id from param X-Websocket-Header-X-UserID"))
				return nil
			}
			api.LeaveCSFieldByWS(user_id, msg.Room)
			log.Printf("[%s] leaving  cs_field [%s] with IP [%s]",
				user_id, msg.Room,
				ctx.RemoteAddr())
			byteUser, _ := json.Marshal((iris.Map{"user_id": user_id}))
			msg.Body = byteUser
			msg.Event = "userLeft"
			nsConn.Conn.Server().Broadcast(nsConn, msg)
			return nil
		},
		websocket.OnRoomLeft: func(nsConn *websocket.NSConn, msg websocket.Message) error {
			// with `websocket.GetContext` you can retrieve the Iris' `Context`.
			ctx := websocket.GetContext(nsConn.Conn)
			user_id := ctx.GetHeader("X-UserID")
			log.Printf("[%s] left  cs_field [%s] with IP [%s]",
				user_id, msg.Room,
				ctx.RemoteAddr())
			byteUser, _ := json.Marshal((iris.Map{"user_id": user_id}))
			msg.Body = byteUser
			msg.Event = "userLeft"
			nsConn.Conn.Server().Broadcast(nsConn, msg)
			return nil
		},
		websocket.OnAnyEvent: func(nsConn *websocket.NSConn, msg websocket.Message) error {
			// with `websocket.GetContext` you can retrieve the Iris' `Context`.
			ctx := websocket.GetContext(nsConn.Conn)
			if msg.Event != "_OnNamespaceConnect" {
				log.Printf("[%s] emit event  [%s] with IP [%s]",
					ctx.GetHeader("X-UserID"), msg.Event,
					ctx.RemoteAddr())
				nsConn.Conn.Server().Broadcast(nsConn, msg)
			}
			return nil
		},
	},

	"timedew": websocket.Events{
		websocket.OnNamespaceConnected: func(nsConn *websocket.NSConn, msg websocket.Message) error {
			// with `websocket.GetContext` you can retrieve the Iris' `Context`.
			ctx := websocket.GetContext(nsConn.Conn)
			user_name := strings.TrimSpace(ctx.GetHeader("X-Username"))
			log.Printf("[%s] connected to namespace [%s] with IP [%s]",
				user_name, msg.Namespace,
				ctx.RemoteAddr())
			return nil
		},
		websocket.OnNamespaceDisconnect: func(nsConn *websocket.NSConn, msg websocket.Message) error {
			ctx := websocket.GetContext(nsConn.Conn)
			user_id := ctx.GetHeader("X-UserID")
			log.Printf("[%s] disconnected from namespace [%s]", user_id, msg.Namespace)
			return nil
		},
		"timedew": func(nsConn *websocket.NSConn, msg websocket.Message) error {
			ctx := websocket.GetContext(nsConn.Conn)
			name := ctx.GetHeader("X-Username")

			// room.String() returns -> NSConn.String() returns -> Conn.String() returns -> Conn.ID()
			user := GetuserByName(name)
			nsConn.Emit("timedew", []byte("用户："+user.Name))
			data := string(msg.Body)
			array := strings.Split(data, "###")
			// VerifyTimeDewData(nsConn, user, array[0], array[1])
			VerifyPromptChainData(nsConn, user, array[0], array[1], array[2])
			log.Printf("[%s] sent: %s", name, string(msg.Body))
			// Write message back to the client message owner with:
			// nsConn.Emit("chat", msg)
			// Write message to all except this client with:
			nsConn.Conn.Server().Broadcast(nsConn, msg)

			return nil
		},
	},
}

func Init() {

	WebsocketServer = websocket.New(
		websocket.DefaultGorillaUpgrader, /* DefaultGobwasUpgrader can be used too. */
		serverEvents)
}

func BroadcastToClient(event string, msg []byte) {
	WebsocketServer.Broadcast(nil, neffos.Message{Namespace: "default", Event: event, Body: msg})
}

func BroadcastForTDVerifyClient(msg string) {
	WebsocketServer.Broadcast(nil, neffos.Message{Namespace: "timedew", Event: "timedew", Body: []byte(msg)})
}

func GetuserByName(username string) *ent.User {
	ctx := context.Background()
	user, err := db.DBClient.User.Query().Where(user.NameEQ(username)).First(ctx)
	if err != nil {
		log.Println("failed when get user: ", username, err.Error())
	}
	if user == nil {
		user, _ = db.DBClient.User.Create().SetName(username).SetSystemName(username).Save(ctx)
		cs_redis.PutMessageQueue("user", "item", user)
	}
	return user
}

func VerifyTimeDewData(nsConn *websocket.NSConn, currUser *ent.User, speechs, place string) {
	ctx := context.Background()
	userName := currUser.Name
	needReplaceName := false
	rdb := cs_redis.RdbClient
	var prompt_seq []string
	var labels []string
	var prompt_seq_full_text []string
	genContent := ""
	needGen := true
	locationChanged := false
	pic_url := ""
	var conf map[string]interface{}
	speech := ""
	speechForGen := ""
	tdcf, _ := rdb.Get(ctx, "cstdcf").Result()
	json.Unmarshal([]byte(tdcf), &conf)
	// duration := conf["duration"].(string)
	// duration_int, _ := strconv.Atoi(duration)
	word_switch := conf["word_switch"].(string)
	word_switch_int, _ := strconv.Atoi(word_switch)
	// gen_pic_prefix := conf["gen_pic_prefix"].(string)
	// gen_pic_subfix := conf["gen_pic_subfix"].(string)
	// route_switch := conf["route_switch"].(string)
	name_replace_switch := conf["name_replace_switch"].(string)
	if name_replace_switch == "Y" {
		userName = "张三"
		needReplaceName = true
		nsConn.Emit("timedew", []byte("进行用户名替换,生成过程中间使用 张三 "))
	}

	last_place, err := rdb.Get(ctx, fmt.Sprintf("td_user_last_locatio_%d", currUser.ID)).Result()
	nsConn.Emit("timedew", []byte("上一个位置"+last_place))
	if last_place != place {
		locationChanged = true
	}

	rdb.Set(ctx, fmt.Sprintf("td_user_last_locatio_%d", currUser.ID), place, 0).Result()

	if len(speech+speechs) < word_switch_int {
		log.Println("word_switch_int and  get speech word", word_switch_int, len(speech+speechs))
		needGen = false
	} else {
		speechForGen = speechs + speech
		log.Println("length of total speech: ", len(speech+speechs), speechForGen, word_switch_int)
	}

	if needGen {
		rdb.Del(ctx, fmt.Sprintf("td_user_zset_%d", currUser.ID)).Result()

		if locationChanged {
			nsConn.Emit("timedew", []byte("位置发生变化"))
			statePrompt := ai.GenerateStatePrompt([]string{userName}, place, speechForGen)
			nsConn.Emit("timedew", []byte("状态总结Prompt："+statePrompt))
			state := ai.GenerateContent(statePrompt)
			nsConn.Emit("timedew", []byte("状态总结生成结果："+state))
			prompt := ai.GenerateSummaryPrompt([]string{userName}, place, speechForGen, state)
			nsConn.Emit("timedew", []byte("总结1Prompt："+prompt))
			genContent = ai.GenerateContent(prompt)
			if needReplaceName {
				genContent = strings.ReplaceAll(genContent, "张三", currUser.Name)
			}
			labels = []string{place, "状态描述"}
			prompt_seq = []string{"状态总结", "总结1"}
			prompt_seq_full_text = []string{statePrompt, state, prompt, genContent}
			log.Println("Gen 总结1", genContent)
			nsConn.Emit("timedew", []byte("生成结果："+genContent))
		} else {
			nsConn.Emit("timedew", []byte("位置未发生变化"))
			rnum := rand.Intn(10)
			if rnum >= 5 {
				prompt2 := ai.GenerateSummaryPrompt2([]string{userName}, place, speechForGen)
				nsConn.Emit("timedew", []byte("总结2 Prompt"+prompt2))
				prompt2_gen := ai.GenerateContent(prompt2)
				nsConn.Emit("timedew", []byte("总结2 生成结果"+prompt2_gen))
				style_prompt := ai.GenerateCompressionPrompt(prompt2_gen)
				nsConn.Emit("timedew", []byte("风格化 Prompt"+style_prompt))
				labels = []string{place, "状态描述"}
				prompt_seq = []string{"总结2", "风格化"}
				prompt_seq_full_text = []string{prompt2, prompt2_gen, style_prompt}
				genContent = ai.GenerateContent(style_prompt)
				if needReplaceName {
					genContent = strings.ReplaceAll(genContent, "张三", currUser.Name)
				}
				log.Println("Gen 总结2 + 风格化", prompt2_gen)
			} else if rnum <= 1 {
				prompt4 := ai.GenerateSummaryPrompt4([]string{userName}, place, speechForGen)
				nsConn.Emit("timedew", []byte("总结4 Prompt"+prompt4))
				prompt_seq = []string{"总结4"}
				labels = []string{place, "文学性"}
				prompt_seq_full_text = []string{prompt4}
				genContent = ai.GenerateContent(prompt4)
				if needReplaceName {
					genContent = strings.ReplaceAll(genContent, "张三", currUser.Name)
				}
				log.Println("Gen 总结4", genContent)
			} else {
				prompt3 := ai.GenerateSummaryPrompt3([]string{userName}, place, speechForGen)
				nsConn.Emit("timedew", []byte("总结3 Prompt"+prompt3))
				prompt_seq = []string{"总结3"}
				labels = []string{place, "文学性", "比喻句"}
				prompt_seq_full_text = []string{prompt3}
				genContent = ai.GenerateContent(prompt3)
				if needReplaceName {
					genContent = strings.ReplaceAll(genContent, "张三", currUser.Name)
				}
				log.Println("Gen 总结3", genContent)

			}
			nsConn.Emit("timedew", []byte("生成结果："+genContent))
		}
	}
	log.Println("Gen Content: ", genContent)
	// if len(genContent) > 0 {
	// 	rnum := rand.Intn(100)
	// 	gen_pic_upscale, _ := strconv.ParseFloat(conf["gen_pic_upscale"].(string), 64)
	// 	gen_pic_rate, _ := strconv.Atoi(conf["gen_pic_rate"].(string))
	// 	if rnum < gen_pic_rate {
	// 		sdhostfmt := fmt.Sprintf("http://%s/sdapi/v1/", conf["gen_pic_host"].(string)) + "%s"
	// 		pic_prompt := ai.GenPicPrompt(genContent)
	// 		prompts := strings.Split(pic_prompt, "prompt:")
	// 		if len(prompts) < 2 {
	// 			log.Println("gen sd pic prompt failed", pic_prompt)
	// 		} else {
	// 			pic_prompt = strings.Split(pic_prompt, "prompt:")[1]

	// 			pic_prompt = gen_pic_prefix + pic_prompt + gen_pic_subfix
	// 			sdimgbs64, err := ai.TrySDT2I(pic_prompt, gen_pic_upscale, sdhostfmt)

	// 			if err != nil || len(sdimgbs64) < 50 {
	// 				log.Println("gen sd pic failed", err.Error())
	// 			} else {
	// 				pic_name := sdimgbs64[220:230] + ".png"

	// 				imageBytes, _ := base64.StdEncoding.DecodeString(sdimgbs64)
	// 				// os.WriteFile("./assets/"+pic_name, imageBytes, 0644)
	// 				utils.UploadFileToS3(imageBytes, pic_name)
	// 				pic_url = fmt.Sprintf("http://sagemaker-us-west-2-887392381071.s3.us-west-2.amazonaws.com/images/%s", pic_name)
	// 				log.Println("Gen pic url: ", pic_url)
	// 			}
	// 		}

	// 	}
	// }
	db_result, err := db.DBClient.TimeDew.Create().SetRawData([]string{speech, place}).
		SetPromptSeq(strings.Join(prompt_seq, ",")).
		SetPromptSeqFullText(strings.Join(prompt_seq_full_text, "###")).
		SetJoinedLabel(strings.Join(labels, ",")).
		SetSpeechs(speechs).SetPlace(place).SetOwnerID(currUser.ID).SetGeneratedContent(genContent).SetPicURL(pic_url).Save(ctx)
	if err != nil {
		log.Println("failed when create timedew: ", err)
		return
	}
	byteResult, _ := json.Marshal(db_result)
	nsConn.Emit("timedew", []byte(byteResult))
}

func VerifyPromptChainData(nsConn *websocket.NSConn, currUser *ent.User, speechs, place, flow string) {
	ctx := context.Background()
	var promptsT []string
	flows := strings.Split(flow, ">")
	nsConn.Emit("timedew", []byte("当前使用Prompt: "))
	client := resty.New()
	for _, item := range flows {
		resp, _ := client.R().
			Get("http://192.168.50.193/assets/" + item)
		result := string(resp.Body())
		nsConn.Emit("timedew", []byte(item+":"+result))
		promptsT = append(promptsT, result)

	}
	userName := currUser.Name
	needReplaceName := false

	rdb := cs_redis.RdbClient
	var prompt_seq []string
	var labels []string
	var prompt_seq_full_text []string
	genContent := ""
	needGen := true
	locationChanged := false
	pic_url := ""
	var conf map[string]interface{}
	speech := ""
	speechForGen := ""
	tdcf, _ := rdb.Get(ctx, "cstdcf").Result()
	json.Unmarshal([]byte(tdcf), &conf)
	// duration := conf["duration"].(string)
	// duration_int, _ := strconv.Atoi(duration)
	word_switch := conf["word_switch"].(string)
	word_switch_int, _ := strconv.Atoi(word_switch)
	// gen_pic_prefix := conf["gen_pic_prefix"].(string)
	// gen_pic_subfix := conf["gen_pic_subfix"].(string)
	// route_switch := conf["route_switch"].(string)
	name_replace_switch := conf["name_replace_switch"].(string)
	if name_replace_switch == "Y" {
		userName = "张三"
		needReplaceName = true
		nsConn.Emit("timedew", []byte("进行用户名替换,生成过程中间使用 张三 "))
	}

	last_place, _ := rdb.Get(ctx, fmt.Sprintf("td_user_last_locatio_%d", currUser.ID)).Result()
	nsConn.Emit("timedew", []byte("上一个位置"+last_place))
	if last_place != place {
		locationChanged = true
	}

	rdb.Set(ctx, fmt.Sprintf("td_user_last_locatio_%d", currUser.ID), place, 0).Result()

	if len(speech+speechs) < word_switch_int {
		log.Println("word_switch_int and  get speech word", word_switch_int, len(speech+speechs))
		needGen = false
	} else {
		speechForGen = speechs + speech
		log.Println("length of total speech: ", len(speech+speechs), speechForGen, word_switch_int)
	}

	if needGen {
		rdb.Del(ctx, fmt.Sprintf("td_user_zset_%d", currUser.ID)).Result()

		if locationChanged {
			genContent = FilePropmtChainWithWS(nsConn, promptsT, userName, speechs, place, "")
		}
		if needReplaceName {
			genContent = strings.ReplaceAll(genContent, "张三", currUser.Name)
			nsConn.Emit("timedew", []byte("GPT最终过程生成内容:\n"+genContent))
		}
	}
	db_result, err := db.DBClient.TimeDew.Create().SetRawData([]string{speech, place}).
		SetPromptSeq(strings.Join(prompt_seq, ",")).
		SetPromptSeqFullText(strings.Join(prompt_seq_full_text, "###")).
		SetJoinedLabel(strings.Join(labels, ",")).
		SetSpeechs(speechs).SetPlace(place).SetOwnerID(currUser.ID).SetGeneratedContent(genContent).SetPicURL(pic_url).Save(ctx)
	if err != nil {
		log.Println("failed when create timedew: ", err)
		return
	}

	byteResult, _ := json.Marshal(db_result)
	nsConn.Emit("timedew", []byte(byteResult))
}

func FilePropmtChainWithWS(nsConn *websocket.NSConn, prompts []string, user, speech, place, output string) string {
	for _, item := range prompts {
		prompt := ai.GenPromptByTemplate(item, user, speech, place, output)
		nsConn.Emit("timedew", []byte("生成所用Prompt:\n"+prompt))
		output = ai.GenerateContent(prompt)
		nsConn.Emit("timedew", []byte("GPT过程生成内容:\n"+output))

	}
	return output
}
