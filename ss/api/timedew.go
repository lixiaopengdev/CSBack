package api

import (
	ai "CSBackendTmp/AI"
	"CSBackendTmp/cs_redis"
	"CSBackendTmp/ent"
	"CSBackendTmp/ent/collection"
	"CSBackendTmp/ent/friendship"
	"CSBackendTmp/ent/hidden"
	"CSBackendTmp/ent/timedew"
	"CSBackendTmp/ent/user"
	"CSBackendTmp/utils"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/go-resty/resty/v2"
	"github.com/kataras/iris/v12"
	"golang.org/x/exp/slices"
)

type UpdateTimeDewReq struct {
	TimedewID         string `json:"timedew_id"`
	PromptSeq         string `json:"prompt_seq"`
	PromptSeqFullText string `json:"prompt_seq_full_text"`
	Labels            string `json:"labels"`
	GenContent        string `json:"genContent"`
	PicURL            string `json:"pic_url"`
}

type TimeDewReq struct {
	Data Data `json:"data"`
}
type Segments struct {
	TimeStamp  int `json:"timeStamp"`
	RangeStart int `json:"rangeStart"`
	RangeSize  int `json:"rangeSize"`
}
type Speechs struct {
	Content   string     `json:"content"`
	TimeStamp float64    `json:"timeStamp"`
	Segments  []Segments `json:"segments"`
}
type Place struct {
	Locality              string `json:"locality"`
	Country               string `json:"country"`
	PostalCode            string `json:"postalCode"`
	AdministrativeArea    string `json:"administrativeArea"`
	SubAdministrativeArea string `json:"subAdministrativeArea"`
	Thoroughfare          string `json:"thoroughfare"`
	Name                  string `json:"name"`
}
type Data struct {
	Speechs []Speechs `json:"speechs"`
	Place   Place     `json:"place"`
	Time    int64     `json:"time"`
}

type TimeDewReqV1 struct {
	TDData TDData `json:"data"`
}

type TDData struct {
	TDSpeechs  []TDSpeechs `json:"speechs"`
	Locactions []Locaction `json:"locactions"`
	Ambients   []Ambient   `json:"ambients"`
}
type TDSpeechs struct {
	Label          int64     `json:"label"`
	Speech         string    `json:"speech"`
	StartTimeStamp int64     `json:"startTimeStamp"`
	EndTimeStamp   int64     `json:"endTimeStamp"`
	Features       []float64 `json:"features"`
}
type Locaction struct {
	TimeStamp             int64   `json:"timeStamp"`
	Longitude             float64 `json:"longitude"`
	Latitude              float64 `json:"latitude"`
	Locality              string  `json:"locality"`
	Country               string  `json:"country"`
	PostalCode            string  `json:"postalCode"`
	AdministrativeArea    string  `json:"administrativeArea"`
	SubAdministrativeArea string  `json:"subAdministrativeArea"`
	Thoroughfare          string  `json:"thoroughfare"`
	Name                  string  `json:"name"`
}
type Ambient struct {
	AudioType      string `json:"audio_type"`
	StartTimeStamp int64  `json:"startTimeStamp"`
	EndTimeStamp   int64  `json:"endTimeStamp"`
}

type LifeFlowItem struct {
	ID            int64     `json:"id"`
	TimeStamp     int64     `json:"timeStamp"`
	Title         string    `json:"title"`
	OnwerID       string    `json:"onwerId"`
	OnwerThumb    string    `json:"onwerThumb"`
	Content       string    `json:"content"`
	PicUrl        string    `json:"pic_url,omitempty"`
	IsSaved       bool      `json:"is_saved"`
	CSFieldID     int64     `json:"cs_field_id"`
	Type          string    `json:"time_dew_type"`
	Members       []int64   `json:"members"`
	MembersThumbs []string  `json:"members_thumbs"`
	Reactions     Reactions `json:"reactions"`
}

type Reactions struct {
	LOL  []UserNameID `json:"LOL,omitempty"`
	DAMN []UserNameID `json:"DAMN,omitempty"`
	OMG  []UserNameID `json:"OMG,omitempty"`
	Nooo []UserNameID `json:"Nooo,omitempty"`
	Cool []UserNameID `json:"Cool,omitempty"`
}

type UserNameID struct {
	UserName string `json:"name"`
	UserID   uint64 `json:"id"`
}

type TimeDewResp struct {
	LifeFlow          []LifeFlowItem `json:"lifeFlow"`
	LifeFlowTimeStamp int64          `json:"lifeFlowTimeStamp"`
}

func PostTimeDewDataV1(ctx iris.Context) {
	var tdreq TimeDewReqV1
	data, _ := ctx.GetBody()

	err := ctx.ReadJSON(&tdreq)
	if err != nil {
		log.Println("failed when create timedew: ", err)
		ctx.JSON(iris.Map{"err_msg": err.Error(), "status": iris.StatusBadRequest})
		return
	}
	currUser := GetCurrentUserData(ctx)
	userName := currUser.Name

	if currUser == nil {
		ctx.JSON(iris.Map{"err_msg": "need login", "status": iris.StatusBadRequest})
		return
	}
	place := ""

	var conf map[string]interface{}
	speechForGen := ""
	var last_location Locaction
	tdcf, _ := rdb.Get(ctx, "cstdcf").Result()
	json.Unmarshal([]byte(tdcf), &conf)

	if len(tdreq.TDData.TDSpeechs) == 0 {
		ctx.JSON(iris.Map{"data": "empty_speech", "status": iris.StatusBadRequest})
		return
	}
	for _, speech := range tdreq.TDData.TDSpeechs {
		if speech.Label == 0 {
			speechForGen += userName + ":\n" + speech.Speech + "\n"
		} else {
			speechForGen += "Person" + strconv.FormatInt(speech.Label, 10) + ":\n" + speech.Speech + "\n"
		}
	}

	if len(tdreq.TDData.Locactions) > 0 {
		last_location = tdreq.TDData.Locactions[len(tdreq.TDData.Locactions)-1]
	}
	if last_location.Thoroughfare != "" {
		place = last_location.Thoroughfare
	} else if last_location.Name != "" {
		place = last_location.Name
	} else {
		place = strings.Join([]string{last_location.Country, last_location.AdministrativeArea, last_location.SubAdministrativeArea, last_location.Locality}, " ")
	}

	db_result, err := dbClient.TimeDew.Create().SetRawData([]string{string(data)}).
		SetSpeechs(speechForGen).SetPlace(place).SetOwnerID(currUser.ID).SetCsFieldID(currUser.CurrentCsFieldID).Save(ctx)

	cs_redis.PutMessageQueue("timedew_rawdata_v1", "item", iris.Map{"client_data": tdreq, "user_data": currUser, "time_dew_id": db_result.ID})

	ctx.JSON(iris.Map{"data": db_result, "status": iris.StatusOK})
}

func UpdateTimeDew(ctx iris.Context) {
	var tdreq UpdateTimeDewReq
	ctx.ReadJSON(&tdreq)
	timedewID := tdreq.TimedewID
	prompt_seq := tdreq.PromptSeq
	prompt_seq_full_text := tdreq.PromptSeqFullText
	labels := tdreq.Labels
	timedewIDInt, _ := strconv.ParseUint(timedewID, 10, 64)
	genContent := tdreq.GenContent
	pic_url := tdreq.PicURL

	_, err := dbClient.TimeDew.Update().Where(timedew.ID(timedewIDInt)).
		SetPromptSeq(prompt_seq).
		SetPromptSeqFullText(prompt_seq_full_text).
		SetJoinedLabel(labels).
		SetGeneratedContent(genContent).SetPicURL(pic_url).Save(ctx)
	if err != nil {
		ctx.JSON(iris.Map{"err_msg": err.Error(), "status": iris.StatusBadRequest})
	} else {
		data, _ := dbClient.TimeDew.Query().Where(timedew.ID(timedewIDInt)).WithOwner().Only(ctx)
		cs_redis.PutMessageQueue("timedew", "item", data)
		ctx.JSON(iris.Map{"data": data, "status": iris.StatusOK})
	}
}

func GenSystemTimeDew(user_id uint64, kind string) error {
	ctx := context.Background()
	cur_user, err := dbClient.User.Query().Where(user.ID(user_id)).Only(ctx)
	tmpl := "%s has entered Come Social."
	if kind == "online" {
		tmpl = "User %s now is online"
	}

	_, err = dbClient.TimeDew.Create().SetGeneratedContent(fmt.Sprintf(tmpl, cur_user.Name)).
		SetOwnerID(user_id).SetCsFieldID(cur_user.CurrentCsFieldID).SetType(timedew.TypeSystem).Save(ctx)

	return err
}

func PostTimeDewData(ctx iris.Context) {
	currUser := GetCurrentUserData(ctx)
	userName := currUser.Name
	needReplaceName := false
	if currUser == nil {
		ctx.JSON(iris.Map{"err_msg": "need login", "status": iris.StatusBadRequest})
		return
	}
	var prompt_seq []string
	var labels []string
	var prompt_seq_full_text []string
	genContent := ""
	needGen := true
	locationChanged := false
	place := ""
	pic_url := ""
	var conf map[string]interface{}
	speech := ""
	speechForGen := ""
	tdcf, _ := rdb.Get(ctx, "cstdcf").Result()
	json.Unmarshal([]byte(tdcf), &conf)
	duration := conf["duration"].(string)
	duration_int, _ := strconv.Atoi(duration)
	word_switch := conf["word_switch"].(string)
	word_switch_int, _ := strconv.Atoi(word_switch)
	gen_pic_prefix := conf["gen_pic_prefix"].(string)
	gen_pic_subfix := conf["gen_pic_subfix"].(string)
	// route_switch := conf["route_switch"].(string)
	name_replace_switch := conf["name_replace_switch"].(string)
	if name_replace_switch == "Y" {
		userName = "张三"
		needReplaceName = true
	}
	now := time.Now()
	startTime := now.Add(-time.Minute * time.Duration(duration_int))

	var tdreq TimeDewReq
	data, _ := ctx.GetBody()
	cs_redis.PutMessageQueue("timedew_rawdata", "item", data)
	err := ctx.ReadJSON(&tdreq)
	if len(tdreq.Data.Speechs) == 0 {
		ctx.JSON(iris.Map{"data": "empty_speech", "status": iris.StatusBadRequest})
		return
	}
	speechs := tdreq.Data.Speechs[0].Content
	result, err := rdb.ZRangeByScoreWithScores(ctx, fmt.Sprintf("td_user_zset_%d", currUser.ID), &redis.ZRangeBy{Min: strconv.FormatInt(1000*startTime.Unix(), 10), Max: strconv.FormatInt(1000*now.Unix(), 10)}).Result()
	for _, item := range result {
		var tdstruct TimeDewReq
		json.Unmarshal([]byte(item.Member.(string)), &tdstruct)
		speech += tdstruct.Data.Speechs[0].Content
		if tdstruct.Data.Speechs[0].Content == speechs {
			log.Println("temp deal with dup report:", speechs)
			ctx.JSON(iris.Map{"data": "dup_speech", "status": iris.StatusBadRequest})
			return
		}
	}
	last_place, err := rdb.Get(ctx, fmt.Sprintf("td_user_last_location_%d", currUser.ID)).Result()

	if tdreq.Data.Place.Thoroughfare != "" {
		place = tdreq.Data.Place.Thoroughfare
	} else if tdreq.Data.Place.Name != "" {
		place = tdreq.Data.Place.Name
	} else {
		place = strings.Join([]string{tdreq.Data.Place.Country, tdreq.Data.Place.AdministrativeArea, tdreq.Data.Place.SubAdministrativeArea, tdreq.Data.Place.Locality}, " ")
	}

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
			statePrompt := ai.GenerateStatePrompt([]string{userName}, place, speechForGen)
			state := ai.GenerateContent(statePrompt)
			prompt := ai.GenerateSummaryPrompt([]string{userName}, place, speechForGen, state)

			genContent = ai.GenerateContent(prompt)
			if needReplaceName {
				genContent = strings.ReplaceAll(genContent, "张三", currUser.Name)
			}
			labels = []string{place, "状态描述"}
			prompt_seq = []string{"状态总结", "总结1"}
			prompt_seq_full_text = []string{statePrompt, state, prompt, genContent}
			log.Println("Gen 总结1", genContent)

		} else {
			rnum := rand.Intn(10)
			if rnum >= 5 {
				prompt2 := ai.GenerateSummaryPrompt2([]string{userName}, place, speechForGen)
				prompt2_gen := ai.GenerateContent(prompt2)
				style_prompt := ai.GenerateCompressionPrompt(prompt2_gen)
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
				prompt_seq = []string{"总结3"}
				labels = []string{place, "文学性", "比喻句"}
				prompt_seq_full_text = []string{prompt3}
				genContent = ai.GenerateContent(prompt3)
				if needReplaceName {
					genContent = strings.ReplaceAll(genContent, "张三", currUser.Name)
				}
				log.Println("Gen 总结3", genContent)

			}
		}
	} else {
		rdb.ZAdd(ctx, fmt.Sprintf("td_user_zset_%d", currUser.ID), redis.Z{Score: float64(tdreq.Data.Time), Member: data}).Result()
	}
	log.Println("Gen Content: ", genContent)
	if len(genContent) > 0 {
		rnum := rand.Intn(100)
		gen_pic_upscale, _ := strconv.ParseFloat(conf["gen_pic_upscale"].(string), 64)
		gen_pic_rate, _ := strconv.Atoi(conf["gen_pic_rate"].(string))
		if rnum < gen_pic_rate {
			sdhostfmt := fmt.Sprintf("http://%s/sdapi/v1/", conf["gen_pic_host"].(string)) + "%s"
			pic_prompt := ai.GenPicPrompt(genContent)
			prompts := strings.Split(pic_prompt, "prompt:")
			if len(prompts) < 2 {
				log.Println("gen sd pic prompt failed", pic_prompt)
			} else {
				pic_prompt = strings.Split(pic_prompt, "prompt:")[1]

				pic_prompt = gen_pic_prefix + pic_prompt + gen_pic_subfix
				sdimgbs64, err := ai.TrySDT2I(pic_prompt, gen_pic_upscale, sdhostfmt)

				if err != nil || len(sdimgbs64) < 50 {
					log.Println("gen sd pic failed", err.Error())
				} else {
					pic_name := sdimgbs64[220:230] + ".png"

					imageBytes, _ := base64.StdEncoding.DecodeString(sdimgbs64)
					// os.WriteFile("./assets/"+pic_name, imageBytes, 0644)
					utils.UploadFileToS3(imageBytes, pic_name)
					pic_url = fmt.Sprintf("http://sagemaker-us-west-2-887392381071.s3.us-west-2.amazonaws.com/images/%s", pic_name)
					log.Println("Gen pic url: ", pic_url)
				}
			}

		}
	}
	db_result, err := dbClient.TimeDew.Create().SetRawData([]string{string(data)}).
		SetPromptSeq(strings.Join(prompt_seq, ",")).
		SetPromptSeqFullText(strings.Join(prompt_seq_full_text, "###")).
		SetJoinedLabel(strings.Join(labels, ",")).
		SetSpeechs(speechs).SetPlace(place).SetOwnerID(currUser.ID).SetGeneratedContent(genContent).SetPicURL(pic_url).Save(ctx)
	if err != nil {
		log.Println("failed when create timedew: ", err)
		ctx.JSON(iris.Map{"err_msg": err, "status": iris.StatusBadRequest})
		return
	}
	if genContent != "" {
		data, _ := dbClient.TimeDew.Query().Where(timedew.ID(db_result.ID)).WithOwner().Only(ctx)
		// timedew_byte, _ := json.Marshal(timedew)
		// result, err := rdb.XAdd(ctx, &redis.XAddArgs{Stream: "timedew", Values: []string{"item", string(timedew_byte)}}).Result()
		cs_redis.PutMessageQueue("timedew", "item", data)
		// log.Println("xadd redis result, err:", result, err)
		// 	cswebsocket.WebsocketServer.Broadcast(nil, neffos.Message{Namespace: "default", Event: "td_update", Body: timedew_byte})
	}
	ctx.JSON(iris.Map{"data": db_result, "status": iris.StatusOK})

}

func GetTimeDews(ctx iris.Context) {
	currUser := GetCurrentUserData(ctx)

	log.Println("get timedews by ", currUser.Name)
	if currUser == nil {
		ctx.JSON(iris.Map{"err_msg": "need login", "status": iris.StatusBadRequest})
		return
	}
	var conf map[string]interface{}
	tdcf, _ := rdb.Get(ctx, "cstdcf").Result()
	json.Unmarshal([]byte(tdcf), &conf)

	batch_no, _ := strconv.Atoi(conf["batch_no"].(string))
	show_pic_rate, _ := strconv.Atoi(conf["show_pic_rate"].(string))

	timedews, err := dbClient.TimeDew.Query().Where(timedew.GeneratedContentNEQ("")).WithOwner().Order(ent.Desc(timedew.FieldID)).Limit(batch_no).All(ctx)
	var tdresp TimeDewResp
	for _, timedew := range timedews {
		var lf LifeFlowItem
		lf.ID = int64(timedew.ID)
		lf.Content = timedew.GeneratedContent
		lf.OnwerID = strconv.FormatUint(timedew.UserID, 10)
		lf.TimeStamp = timedew.CreateTime.UnixMilli()
		rnum := rand.Intn(100)
		if rnum < show_pic_rate {
			lf.PicUrl = timedew.PicURL
		}
		tdresp.LifeFlowTimeStamp = timedew.CreateTime.UnixMilli()
		tdresp.LifeFlow = append(tdresp.LifeFlow, lf)
	}

	if err != nil {
		log.Println("failed when GetTimeDews: ", err)
		ctx.JSON(iris.Map{"err_msg": err, "status": iris.StatusBadRequest})
		return
	}
	ctx.JSON(iris.Map{"data": tdresp, "status": iris.StatusOK})

}

func GetTimeDewsReactions(ctx iris.Context) {
	currUser := GetCurrentUserData(ctx)

	var item_ids []uint64
	var hidden_ids []uint64
	var friend_ids []uint64
	err := dbClient.Friendship.Query().Where(friendship.FriendID(currUser.ID), friendship.StatusEQ(friendship.StatusEstablished)).Select(friendship.FieldUserID).Scan(ctx, &friend_ids)
	err = dbClient.Hidden.Query().Where(hidden.UserID(currUser.ID)).Select(hidden.FieldHiddenID).Scan(ctx, &hidden_ids)
	friend_ids = append(friend_ids, currUser.ID)

	log.Println("get timedews by ", currUser.ID, currUser.Name, friend_ids)

	err = dbClient.Collection.Query().Where(collection.UserID(currUser.ID)).Select(collection.FieldItemID).Scan(ctx, &item_ids)

	reacts, err := dbClient.TimeDew.Query().Where(timedew.UserIDIn(friend_ids...), timedew.GeneratedContentNEQ("")).WithOwner().Order(ent.Desc(timedew.FieldID)).Limit(50).WithReactions().All(ctx)
	var tdresp TimeDewResp
	for _, item := range reacts {
		if slices.Contains(hidden_ids, item.UserID) {
			continue
		}
		if item.Type == timedew.TypeInvite {
			if item.UserID == currUser.ID {
				continue
			}
			if item.TargetID != currUser.ID {
				continue
			}
		}
		var lf LifeFlowItem
		lf.ID = int64(item.ID)
		if slices.Contains(item_ids, item.ID) {
			lf.IsSaved = true
		} else {
			lf.IsSaved = false
		}
		lf.Content = item.GeneratedContent
		lf.OnwerID = strconv.FormatUint(item.UserID, 10)
		if item.Edges.Owner != nil {
			lf.Title = item.Edges.Owner.Name
			lf.OnwerThumb = item.Edges.Owner.ThumbnailURL
		}
		lf.TimeStamp = item.CreateTime.UnixMilli()
		lf.PicUrl = item.PicURL
		lf.CSFieldID = int64(item.CsFieldID)
		if item.Type == timedew.TypeCsField {
			lf.Type = "field"
			lf.Members = []int64{int64(item.UserID)}
			lf.MembersThumbs = []string{item.Edges.Owner.ThumbnailURL}
		} else if item.Type == timedew.TypeUser {
			lf.Type = "single"
		} else if item.Type == timedew.TypeSystem {
			lf.Type = "system"
		} else if item.Type == timedew.TypeInvite {
			lf.Type = "invite"
		}
		self_td := false
		if item.UserID == currUser.ID {
			self_td = true
		}
		if item.Edges.Reactions != nil {
			self_react := false
			for _, react := range item.Edges.Reactions {

				if react.UserID == currUser.ID {
					self_react = true
				}

				if react.IsCool {
					if self_td || self_react {
						name, err := dbClient.User.Query().Where(user.IDEQ(react.UserID)).Select(user.FieldName).Strings(ctx)
						if err == nil {
							lf.Reactions.Cool = append(lf.Reactions.Cool, UserNameID{UserName: name[0], UserID: react.UserID})
						}
					}
				}
				if react.IsDAMN {
					if self_td || self_react {
						name, err := dbClient.User.Query().Where(user.IDEQ(react.UserID)).Select(user.FieldName).Strings(ctx)
						if err == nil {
							lf.Reactions.DAMN = append(lf.Reactions.DAMN, UserNameID{UserName: name[0], UserID: react.UserID})
						}
					}
				}
				if react.IsLOL {
					if self_td || self_react {
						name, err := dbClient.User.Query().Where(user.IDEQ(react.UserID)).Select(user.FieldName).Strings(ctx)
						if err == nil {
							lf.Reactions.LOL = append(lf.Reactions.LOL, UserNameID{UserName: name[0], UserID: react.UserID})
						}
					}
				}
				if react.IsNooo {
					if self_td || self_react {
						name, err := dbClient.User.Query().Where(user.IDEQ(react.UserID)).Select(user.FieldName).Strings(ctx)
						if err == nil {
							lf.Reactions.Nooo = append(lf.Reactions.Nooo, UserNameID{UserName: name[0], UserID: react.UserID})
						}
					}
				}
				if react.IsOMG {
					if self_td || self_react {
						name, err := dbClient.User.Query().Where(user.IDEQ(react.UserID)).Select(user.FieldName).Strings(ctx)
						if err == nil {
							lf.Reactions.OMG = append(lf.Reactions.OMG, UserNameID{UserName: name[0], UserID: react.UserID})
						}
					}
				}
			}
		}
		tdresp.LifeFlowTimeStamp = item.CreateTime.UnixMilli()
		tdresp.LifeFlow = append(tdresp.LifeFlow, lf)
	}

	if err != nil {
		log.Println("failed when GetTimeDews: ", err)
		ctx.JSON(iris.Map{"err_msg": err, "status": iris.StatusBadRequest})
		return
	}
	ctx.JSON(iris.Map{"data": tdresp, "status": iris.StatusOK})

}

func GetAdminTimeDews(ctx iris.Context) {
	cursor := strings.TrimSpace(ctx.URLParam("cursor"))
	if cursor == "" || cursor == "0" {
		cursor = "100000000000000"
	}
	limit := strings.TrimSpace(ctx.URLParam("limit"))
	LimitInt, _ := strconv.ParseInt(limit, 10, 0)
	if LimitInt > 500 {
		LimitInt = 500
	}
	cursorInt, _ := strconv.ParseUint(cursor, 10, 64)
	timedews, err := dbClient.TimeDew.Query().Where(timedew.GeneratedContentNEQ(""), timedew.IDLTE(cursorInt)).WithOwner().Order(ent.Desc(timedew.FieldID)).Limit(int(LimitInt)).All(ctx)
	var tdresp TimeDewResp
	for _, item := range timedews {
		var lf LifeFlowItem
		lf.ID = int64(item.ID)
		lf.Content = item.GeneratedContent
		lf.OnwerID = strconv.FormatUint(item.UserID, 10)
		if item.Edges.Owner != nil {
			lf.Title = item.Edges.Owner.Name
			lf.OnwerThumb = item.Edges.Owner.ThumbnailURL
		}
		lf.TimeStamp = item.CreateTime.UnixMilli()
		lf.PicUrl = item.PicURL
		lf.CSFieldID = int64(item.CsFieldID)
		if item.Type == timedew.TypeCsField {
			lf.Type = "field"
			lf.Members = []int64{int64(item.UserID)}
		} else if item.Type == timedew.TypeUser {
			lf.Type = "single"
		} else if item.Type == timedew.TypeSystem {
			lf.Type = "system"
		}
		tdresp.LifeFlowTimeStamp = item.CreateTime.UnixMilli()
		tdresp.LifeFlow = append(tdresp.LifeFlow, lf)
	}

	if err != nil {
		log.Println("failed when GetTimeDews: ", err)
		ctx.JSON(iris.Map{"err_msg": err, "status": iris.StatusBadRequest})
		return
	}
	ctx.JSON(iris.Map{"data": tdresp, "status": iris.StatusOK})

}

func GetAdminUsers(ctx iris.Context) {
	cursor := strings.TrimSpace(ctx.URLParam("cursor"))
	log.Println("cursor", cursor, len(cursor))
	if cursor == "" || cursor == "0" {
		cursor = "100000000000000"
	}
	limit := strings.TrimSpace(ctx.URLParam("limit"))
	LimitInt, _ := strconv.ParseInt(limit, 10, 0)
	if LimitInt > 500 {
		LimitInt = 500
	}
	cursorInt, _ := strconv.ParseUint(cursor, 10, 64)
	users, err := dbClient.User.Query().Where(user.IDLT(cursorInt)).Order(ent.Desc(user.FieldID)).Limit(int(LimitInt)).All(ctx)

	if err != nil {
		log.Println("failed when GetTimeDews: ", err)
		ctx.JSON(iris.Map{"err_msg": err, "status": iris.StatusBadRequest})
		return
	}
	ctx.JSON(iris.Map{"data": users, "status": iris.StatusOK})

}

func GetUserRelatedTimedew(ctx iris.Context) {
	currUser := GetCurrentUserData(ctx)
	var item_ids []uint64
	err := dbClient.Friendship.Query().Where(friendship.FriendID(currUser.ID), friendship.StatusEQ(friendship.StatusEstablished)).Select(friendship.FieldUserID).Scan(ctx, &item_ids)
	item_ids = append(item_ids, currUser.ID)
	log.Println("get timedews by ", currUser.ID, currUser.Name, item_ids)

	// var conf map[string]interface{}
	// tdcf, _ := rdb.Get(ctx, "cstdcf").Result()
	// json.Unmarshal([]byte(tdcf), &conf)

	// batch_no, _ := strconv.Atoi(conf["batch_no"].(string))
	// show_pic_rate, _ := strconv.Atoi(conf["show_pic_rate"].(string))

	timedews, err := dbClient.TimeDew.Query().Where(timedew.UserIDIn(item_ids...), timedew.GeneratedContentNEQ("")).WithOwner().Order(ent.Desc(timedew.FieldID)).Limit(50).All(ctx)
	var tdresp TimeDewResp
	for _, item := range timedews {
		var lf LifeFlowItem
		lf.ID = int64(item.ID)
		lf.Content = item.GeneratedContent
		lf.OnwerID = strconv.FormatUint(item.UserID, 10)
		if item.Edges.Owner != nil {
			lf.Title = item.Edges.Owner.Name
			lf.OnwerThumb = item.Edges.Owner.ThumbnailURL
		}
		lf.TimeStamp = item.CreateTime.UnixMilli()
		lf.PicUrl = item.PicURL
		lf.CSFieldID = int64(item.CsFieldID)
		if item.Type == timedew.TypeCsField {
			lf.Type = "field"
			lf.Members = []int64{int64(item.UserID)}
		} else if item.Type == timedew.TypeUser {
			lf.Type = "single"
		} else if item.Type == timedew.TypeSystem {
			lf.Type = "system"
		}

		tdresp.LifeFlowTimeStamp = item.CreateTime.UnixMilli()
		tdresp.LifeFlow = append(tdresp.LifeFlow, lf)
	}

	if err != nil {
		log.Println("failed when GetTimeDews: ", err)
		ctx.JSON(iris.Map{"err_msg": err, "status": iris.StatusBadRequest})
		return
	}
	ctx.JSON(iris.Map{"data": tdresp, "status": iris.StatusOK})

}

func GetRecTimeDews(ctx iris.Context) {
	currUser := GetCurrentUserData(ctx)
	var conf map[string]interface{}
	tdcf, _ := rdb.Get(ctx, "cstdcf").Result()
	json.Unmarshal([]byte(tdcf), &conf)

	batch_no, _ := strconv.Atoi(conf["batch_no"].(string))
	show_pic_rate, _ := strconv.Atoi(conf["show_pic_rate"].(string))
	var recResp []string
	var recIDs []uint64

	resp, err := utils.GetTimeDewsIDByRec(strconv.FormatUint(currUser.ID, 10))
	log.Println("get rec timedews by ", currUser.Name, resp)
	if err != nil {
		ctx.JSON(iris.Map{"err_msg": "get rec timedews failed, " + err.Error(), "status": iris.StatusBadRequest})
		return
	}

	err = json.Unmarshal([]byte(resp), &recResp)

	if err != nil {
		ctx.JSON(iris.Map{"err_msg": "get rec timedews failed, " + err.Error(), "status": iris.StatusBadRequest})
		return
	}
	for _, id := range recResp {
		uid, _ := strconv.ParseUint(id, 10, 64)
		recIDs = append(recIDs, uid)
	}

	timedews, err := dbClient.TimeDew.Query().Where(timedew.GeneratedContentNEQ(""), timedew.IDIn(recIDs...)).WithOwner().Order(ent.Desc(timedew.FieldID)).Limit(batch_no).All(ctx)
	var tdresp TimeDewResp
	for _, timedew := range timedews {
		var lf LifeFlowItem
		lf.ID = int64(timedew.ID)
		lf.Content = timedew.GeneratedContent
		lf.OnwerID = strconv.FormatUint(timedew.UserID, 10)
		lf.Title = timedew.Edges.Owner.Name
		lf.TimeStamp = timedew.CreateTime.UnixMilli()
		rnum := rand.Intn(100)
		if rnum < show_pic_rate {
			lf.PicUrl = timedew.PicURL
		}
		tdresp.LifeFlowTimeStamp = timedew.CreateTime.UnixMilli()
		tdresp.LifeFlow = append(tdresp.LifeFlow, lf)
	}

	if err != nil {
		log.Println("failed when GetTimeDews: ", err)
		ctx.JSON(iris.Map{"err_msg": err, "status": iris.StatusBadRequest})
		return
	}
	ctx.JSON(iris.Map{"data": tdresp, "status": iris.StatusOK})

}

func GetWebTimeDews(ctx iris.Context) {

	var conf map[string]interface{}
	tdcf, _ := rdb.Get(ctx, "cstdcf").Result()
	json.Unmarshal([]byte(tdcf), &conf)

	batch_no, _ := strconv.Atoi(conf["batch_no"].(string))
	show_pic_rate, _ := strconv.Atoi(conf["show_pic_rate"].(string))

	timedews, err := dbClient.TimeDew.Query().Where(timedew.GeneratedContentNEQ("")).WithOwner().Order(ent.Desc(timedew.FieldID)).Limit(batch_no).All(ctx)
	var tdresp TimeDewResp
	for _, timedew := range timedews {
		var lf LifeFlowItem
		lf.ID = int64(timedew.ID)
		lf.Content = timedew.GeneratedContent
		lf.OnwerID = strconv.FormatUint(timedew.UserID, 10)
		if timedew.Edges.Owner != nil {
			lf.Title = timedew.Edges.Owner.Name
		}
		lf.TimeStamp = timedew.CreateTime.UnixMilli()
		rnum := rand.Intn(100)
		if rnum < show_pic_rate {
			lf.PicUrl = timedew.PicURL
		}
		tdresp.LifeFlowTimeStamp = timedew.CreateTime.UnixMilli()
		tdresp.LifeFlow = append(tdresp.LifeFlow, lf)
	}

	if err != nil {
		log.Println("failed when GetTimeDews: ", err)
		ctx.JSON(iris.Map{"err_msg": err, "status": iris.StatusBadRequest})
		return
	}

	ctx.ViewData("tds", tdresp.LifeFlow)

	ctx.View("timedew_web.html")
	// ctx.JSON(iris.Map{"td": tdresp, "status": iris.StatusBadRequest})

}

func CheckTimeDews(ctx iris.Context) {

	var conf map[string]interface{}
	tdcf, _ := rdb.Get(ctx, "cstdcf").Result()
	json.Unmarshal([]byte(tdcf), &conf)

	batch_no, _ := strconv.Atoi(conf["batch_no"].(string))

	timedews, err := dbClient.TimeDew.Query().Where().WithOwner().Order(ent.Desc(timedew.FieldID)).Limit(batch_no).All(ctx)

	if err != nil {

		ctx.JSON(iris.Map{"err_msg": err, "status": iris.StatusBadRequest})
		return
	}
	var jsonOptions = iris.JSON{
		Indent: "    ",
	}
	ctx.JSON(iris.Map{"data": timedews, "status": iris.StatusOK}, jsonOptions)

}

func GetTimeDewConf(ctx iris.Context) {
	var conf map[string]interface{}
	resp := ctx.URLParam("resp")
	tdcf, _ := rdb.Get(ctx, "cstdcf").Result()
	json.Unmarshal([]byte(tdcf), &conf)
	log.Println("cstdcf", conf)
	ctx.ViewData("report_time_switch", conf["report_time_switch"])
	ctx.ViewData("gen_time_switch", conf["gen_time_switch"])
	ctx.ViewData("word_switch", conf["word_switch"])
	ctx.ViewData("seq", conf["seq"])
	ctx.ViewData("pro1", conf["pro1"])
	ctx.ViewData("pro2", conf["pro2"])
	ctx.ViewData("name", conf["name"])
	ctx.ViewData("speech", conf["speech"])
	ctx.ViewData("place", conf["place"])
	ctx.ViewData("duration", conf["duration"])
	ctx.ViewData("batch_no", conf["batch_no"])
	ctx.ViewData("route_switch", conf["route_switch"])
	ctx.ViewData("gen_pic_prefix", conf["gen_pic_prefix"])
	ctx.ViewData("gen_pic_subfix", conf["gen_pic_subfix"])
	ctx.ViewData("gen_pic_rate", conf["gen_pic_rate"])
	ctx.ViewData("show_pic_rate", conf["show_pic_rate"])
	ctx.ViewData("gen_pic_upscale", conf["gen_pic_upscale"])
	ctx.ViewData("gen_pic_host", conf["gen_pic_host"])
	ctx.ViewData("name_replace_switch", conf["name_replace_switch"])
	ctx.ViewData("rec_switch", conf["rec_switch"])

	// Render template file: ./views/hi.html
	if resp == "json" {
		ctx.JSON(iris.Map{"data": conf, "status": iris.StatusOK})
	} else {
		ctx.View("timedew_config.html")
	}

	// ctx.ServeFile("./templates/timedew_config.html")
}

func SetTimeDewConf(ctx iris.Context) {
	user := strings.TrimSpace(ctx.PostValue("name"))
	speech := strings.TrimSpace(ctx.PostValue("speech"))
	place := strings.TrimSpace(ctx.PostValue("place"))
	seq := strings.TrimSpace(ctx.PostValue("seq"))
	pro1 := strings.TrimSpace(ctx.PostValue("pro1"))
	duration := strings.TrimSpace(ctx.PostValue("duration"))
	batch_no := strings.TrimSpace(ctx.PostValue("batch_no"))
	route_switch := strings.TrimSpace(ctx.PostValue("route_switch"))
	// pro2 := strings.TrimSpace(ctx.PostValue("pro2"))
	report_time_switch := strings.TrimSpace(ctx.PostValue("report_time_switch"))
	gen_time_switch := strings.TrimSpace(ctx.PostValue("gen_time_switch"))
	word_switch := strings.TrimSpace(ctx.PostValue("word_switch"))

	gen_pic_subfix := strings.TrimSpace(ctx.PostValue("gen_pic_subfix"))
	gen_pic_prefix := strings.TrimSpace(ctx.PostValue("gen_pic_prefix"))
	gen_pic_rate := strings.TrimSpace(ctx.PostValue("gen_pic_rate"))
	show_pic_rate := strings.TrimSpace(ctx.PostValue("show_pic_rate"))
	gen_pic_upscale := strings.TrimSpace(ctx.PostValue("gen_pic_upscale"))
	gen_pic_host := strings.TrimSpace(ctx.PostValue("gen_pic_host"))
	name_replace_switch := strings.TrimSpace(ctx.PostValue("name_replace_switch"))
	rec_switch := strings.TrimSpace(ctx.PostValue("rec_switch"))

	conf := make(map[string]interface{})
	// ctx.ReadJSON(&conf)
	conf["batch_no"] = batch_no
	conf["route_switch"] = route_switch
	conf["report_time_switch"] = report_time_switch
	conf["gen_time_switch"] = gen_time_switch
	conf["word_switch"] = word_switch
	conf["duration"] = duration
	conf["seq"] = seq
	conf["pro1"] = pro1
	conf["gen_pic_prefix"] = gen_pic_prefix
	conf["gen_pic_subfix"] = gen_pic_subfix
	conf["gen_pic_rate"] = gen_pic_rate
	conf["show_pic_rate"] = show_pic_rate
	conf["gen_pic_upscale"] = gen_pic_upscale
	conf["gen_pic_host"] = gen_pic_host
	conf["name_replace_switch"] = name_replace_switch
	conf["rec_switch"] = rec_switch

	// conf["pro2"] = pro2
	conf["name"] = user
	conf["speech"] = speech
	conf["place"] = place
	confBuff, _ := json.Marshal(conf)
	rdb.Set(ctx, "cstdcf", confBuff, 0).Result()
	genContent := ai.ProcessContentWithCustomPrompt([]string{user}, place, speech, pro1)
	ctx.JSON(iris.Map{"data": genContent, "status": iris.StatusOK})
}

func OfflineGen(ctx iris.Context) {
	var conf map[string]interface{}
	// ctx := context.Background()

	tdcf, _ := rdb.Get(ctx, "cstdcf").Result()
	json.Unmarshal([]byte(tdcf), &conf)
	gen_time_switch := conf["gen_time_switch"].(string)
	gen_time_switch_int, _ := strconv.Atoi(gen_time_switch)
	now := time.Now()
	startTime := now.Add(-time.Minute * time.Duration(gen_time_switch_int))

	timedews, _ := dbClient.TimeDew.Query().Where(timedew.CreateTimeGTE(startTime), timedew.GeneratedContent("")).WithOwner().Order(ent.Asc(timedew.FieldCreateTime)).All(ctx)

	for i, item := range timedews {
		log.Println(i, item)

	}
}

func ReadPromptFile(url string) string {
	client := resty.New()
	var conf map[string]interface{}
	ctx := context.Background()
	tdcf, _ := rdb.Get(ctx, "cstdcf").Result()
	json.Unmarshal([]byte(tdcf), &conf)
	user := conf["name"].(string)
	place := conf["place"].(string)
	state := "状态"
	speech := conf["speech"].(string)
	var params []interface{}
	resp, _ := client.R().
		Get(url)
	log.Println(string(resp.Body()))

	regexp, _ := regexp.Compile(`{.*}`)
	ma := regexp.FindAllString(string(resp.Body()), -1)
	matchParamLen := 0
	for _, st := range ma {

		if strings.Contains(st, "用户名") {
			params = append(params, user)
		}

		if strings.Contains(st, "地点") {
			params = append(params, place)
		}

		if strings.Contains(st, "语音转文本") {
			params = append(params, speech)
		}
		if strings.Contains(st, "场景概括") {
			params = append(params, state)
		}
		matchParamLen += 1
	}

	rp := regexp.ReplaceAllString(string(resp.Body()), "%s")
	return fmt.Sprintf(rp, params...)

	// ctx.JSON(iris.Map{"data": ma, "rp": rp, "result": fmt.Sprintf(rp, params...), "status": iris.StatusOK})

}

func FilePropmtChain(prompts []string, user, speech, place, output string) string {
	for _, item := range prompts {
		prompt := ai.GenPromptByTemplate(item, user, speech, place, output)
		output = ai.GenerateContent(prompt)

	}
	return output
}

func ConsumeTimedewStream(ctx iris.Context) {
	rdb.XGroupCreate(ctx, "timedew", "test2", "0").Err()
	for {
		entries, err := rdb.XReadGroup(ctx, &redis.XReadGroupArgs{
			Group:    "test2",
			Consumer: "test1",
			Streams:  []string{"timedew", ">"},
			Count:    2,
			Block:    0,
			NoAck:    false,
		}).Result()
		if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < len(entries[0].Messages); i++ {
			messageID := entries[0].Messages[i].ID
			values := entries[0].Messages[i].Values
			timedew := fmt.Sprintf("%v", values["td_item"])
			log.Println("consume redis stream:", timedew)
			// timedew := fmt.Sprintf("%v", values["ticketID"])
			rdb.XAck(ctx, "timedew", "test", messageID)
		}
	}
}

func TmpUpload(ctx iris.Context) {

	bs64 := strings.TrimSpace(ctx.PostValue("image"))
	pic_name := strconv.FormatInt(utils.GenID(), 10) + ".png"

	imageBytes, _ := base64.StdEncoding.DecodeString(bs64)
	// os.WriteFile("./assets/"+pic_name, imageBytes, 0644)
	utils.UploadFileToS3(imageBytes, pic_name)
	pic_url := fmt.Sprintf("http://sagemaker-us-west-2-887392381071.s3.us-west-2.amazonaws.com/images/%s", pic_name)
	log.Println("Gen pic url: ", pic_url)
	ctx.JSON(iris.Map{"data": pic_url, "status": iris.StatusOK})

}

func TimeDewReaction(ctx iris.Context) {
	currUser := GetCurrentUserData(ctx)
	tid, _ := strconv.ParseUint(ctx.PostValue("timedew_id"), 10, 64)
	reaction := ctx.PostValue("reaction")
	var err error
	switch reaction {
	case "LOL":
		_, err = dbClient.Reaction.Create().SetIsLOL(true).SetTimeDewID(tid).SetUserID(currUser.ID).OnConflict().SetIsLOL(true).ID(ctx)
	case "DAMN":
		_, err = dbClient.Reaction.Create().SetIsDAMN(true).SetTimeDewID(tid).SetUserID(currUser.ID).OnConflict().SetIsDAMN(true).ID(ctx)
	case "Nooo":
		_, err = dbClient.Reaction.Create().SetIsNooo(true).SetTimeDewID(tid).SetUserID(currUser.ID).OnConflict().SetIsNooo(true).ID(ctx)
	case "OMG":
		_, err = dbClient.Reaction.Create().SetIsOMG(true).SetTimeDewID(tid).SetUserID(currUser.ID).OnConflict().SetIsOMG(true).ID(ctx)
	case "Cool":
		_, err = dbClient.Reaction.Create().SetIsCool(true).SetTimeDewID(tid).SetUserID(currUser.ID).OnConflict().SetIsCool(true).ID(ctx)
	default:
		log.Println("reaction", currUser.ID, tid, reaction)
	}
	if err != nil {
		ctx.JSON(iris.Map{"err_msg": err.Error(), "status": iris.StatusBadRequest})
		return
	}
	ctx.JSON(iris.Map{"data": "ok", "status": iris.StatusOK})
}

func DeleteTimeDewReaction(ctx iris.Context) {
	currUser := GetCurrentUserData(ctx)
	tid, _ := strconv.ParseUint(ctx.URLParam("timedew_id"), 10, 64)
	reaction := ctx.URLParam("reaction")
	var err error
	switch reaction {
	case "LOL":
		_, err = dbClient.Reaction.Create().SetIsLOL(false).SetTimeDewID(tid).SetUserID(currUser.ID).OnConflict().SetIsLOL(false).ID(ctx)
	case "DAMN":
		_, err = dbClient.Reaction.Create().SetIsDAMN(false).SetTimeDewID(tid).SetUserID(currUser.ID).OnConflict().SetIsDAMN(false).ID(ctx)
	case "Nooo":
		_, err = dbClient.Reaction.Create().SetIsNooo(false).SetTimeDewID(tid).SetUserID(currUser.ID).OnConflict().SetIsNooo(false).ID(ctx)
	case "OMG":
		_, err = dbClient.Reaction.Create().SetIsOMG(false).SetTimeDewID(tid).SetUserID(currUser.ID).OnConflict().SetIsOMG(false).ID(ctx)
	case "Cool":
		_, err = dbClient.Reaction.Create().SetIsCool(false).SetTimeDewID(tid).SetUserID(currUser.ID).OnConflict().SetIsCool(false).ID(ctx)
	default:
		log.Println("reaction", currUser.ID, tid, reaction)
	}
	if err != nil {
		ctx.JSON(iris.Map{"err_msg": err.Error(), "status": iris.StatusBadRequest})
		return
	}
	ctx.JSON(iris.Map{"data": "ok", "status": iris.StatusOK})
}

func PostTimeDew(ctx iris.Context) {
	curr_user := GetCurrentUserData(ctx)
	genContent := strings.TrimSpace(ctx.PostValue("content"))
	pic_url := strings.TrimSpace(ctx.PostValue("pic_url"))
	time_dew_type := strings.TrimSpace(ctx.PostValue("time_dew_type"))
	td_type := timedew.TypeUser
	if time_dew_type == "field" {
		td_type = timedew.TypeCsField
	}
	td, err := dbClient.TimeDew.Create().SetGeneratedContent(genContent).
		SetPicURL(pic_url).SetOwnerID(curr_user.ID).SetCsFieldID(curr_user.CurrentCsFieldID).SetType(td_type).Save(ctx)
	if err != nil {
		ctx.JSON(iris.Map{"err_msg": err.Error(), "status": iris.StatusBadRequest})
	} else {
		ctx.JSON(iris.Map{"data": td, "status": iris.StatusOK})
	}
}

func genInviteTimedew(curUser *ent.User, target_id uint64) error {
	ctx := context.Background()
	genContent := fmt.Sprintf("%s invites you to join %s", curUser.Name, curUser.CurrentCsFieldName)
	td, err := dbClient.TimeDew.Create().SetGeneratedContent(genContent).
		SetOwnerID(curUser.ID).SetCsFieldID(curUser.CurrentCsFieldID).SetType(timedew.TypeInvite).SetTargetID(target_id).Save(ctx)
	rdb.HMSet(ctx, fmt.Sprintf("%d_%d_invite_to_join", target_id, curUser.CurrentCsFieldID), curUser.ID, td.ID)
	return err
}
