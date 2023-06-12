package api

import (
	"CSBackendTmp/ent"
	"CSBackendTmp/ent/card"
	"CSBackendTmp/ent/collection"
	"CSBackendTmp/ent/csfield"
	"CSBackendTmp/ent/feedback"
	"CSBackendTmp/ent/friendship"
	"CSBackendTmp/ent/hidden"
	"CSBackendTmp/ent/message"
	"CSBackendTmp/ent/predicate"
	"CSBackendTmp/ent/setting"
	"CSBackendTmp/ent/timedew"
	"CSBackendTmp/ent/user"
	"CSBackendTmp/utils"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/kataras/iris/v12"
	"golang.org/x/exp/slices"
)

type UserNameStatus struct {
	UserName      string `json:"name"`
	UserID        uint64 `json:"id"`
	UserStatus    string `json:"user_status"`
	Status        string `json:"status,omitempty"`
	UserThumbnail string `json:"user_thumbnail"`
	CanAction     bool   `json:"can_action,omitempty"`
	ActionUrl     string `json:"action_url,omitempty"`
}

type InviteFriendsResp struct {
	UserList  []UserNameStatus `json:"you_might_know_them"`
	InviteUrl string           `json:"invite_via_link"`
}

type SearchFriendsResp struct {
	UserListSuggested []UserNameStatus `json:"suggested"`
	UserListMore      []UserNameStatus `json:"more"`
}

type FriendsNoticeResp struct {
	FriendsNoticeList []FriendsNotice `json:"notice_list"`
}

type FriendsNotice struct {
	UserList []UserNameStatus `json:"user_list"`
	Content  string           `json:"content"`
}

type InviteToFieldListResp struct {
	UserListSuggested []UserNameStatus `json:"suggested"`
	UserListMore      []UserNameStatus `json:"more"`
}

type ConectionRequestResp struct {
	UserList []UserNameStatus `json:"user_list"`
}

type brReq struct {
	TargetID uint64 `json:"target_id"`
}

func GetUsers(ctx iris.Context) {
	users, err := dbClient.User.Query().All(ctx)
	if err != nil {
		log.Println("failed when get all users: ", err)
	}
	ctx.JSON(users)
}

func GetUserByID(ctx iris.Context) {
	id := strings.TrimSpace(ctx.Params().Get("id"))
	id32, _ := strconv.ParseUint(id, 10, 32)
	user, err := dbClient.User.Query().Where(user.ID(id32)).WithSetting().Only(ctx)
	if err != nil {
		log.Println("failed when get all users: ", err)
	}
	ctx.JSON(user)
}

func CheckDupUserName(ctx iris.Context) {
	systemName := strings.TrimSpace(ctx.URLParam("user_name"))
	isExist, err := dbClient.User.Query().Where(user.SystemNameEQ(systemName)).Exist(ctx)
	if err != nil {
		log.Println("failed creating user", err)
		ctx.JSON(iris.Map{"err_msg": err.Error(), "status": iris.StatusBadRequest})
		return
	}
	if isExist {
		ctx.JSON(iris.Map{"err_msg": "Username exists already, please try another one", "status": iris.StatusBadRequest})
		return
	}
	ctx.JSON(iris.Map{"data": "Username not exists", "status": iris.StatusOK})
}

func UpdateUserThumbnail(ctx iris.Context) {
	curr_user := GetCurrentUserData(ctx)
	thumbnai_url := ""
	files, _, err := ctx.UploadFormFiles("./assets")
	if err != nil {
		log.Println(err)
		ctx.StopWithStatus(iris.StatusInternalServerError)
		return
	}

	for _, file := range files {
		content, _ := file.Open()
		buf := new(bytes.Buffer)
		buf.ReadFrom(content)
		name := fmt.Sprintf("%d-%s", curr_user.ID, file.Filename)
		utils.UploadFileToS3(buf.Bytes(), name)
		thumbnai_url = fmt.Sprintf("https://ruleless.s3.eu-west-1.amazonaws.com/images/%s", name)
	}
	_, err = dbClient.User.Update().Where(user.ID(curr_user.ID)).SetThumbnailURL(thumbnai_url).Save(ctx)
	curr_user, err = dbClient.User.Query().Where(user.ID(curr_user.ID)).Only(ctx)
	ctx.SetUser(curr_user)
	sessID := ctx.GetCookie("_session_id")
	user_json, _ := json.Marshal(curr_user)

	rdb.Set(ctx, sessID, user_json, 0).Result()

	if err != nil {
		log.Println("failed when update user info: ", err)
		ctx.JSON(iris.Map{"err_msg": err.Error(), "status": iris.StatusBadRequest})
		return
	}
	ctx.JSON(iris.Map{"data": thumbnai_url, "status": iris.StatusOK})
}

func UpdateUserInfo(ctx iris.Context) {
	curr_user := GetCurrentUserData(ctx)
	id := curr_user.ID
	name := strings.TrimSpace(ctx.PostValue("name"))
	username := strings.TrimSpace(ctx.PostValue("usernameID"))

	school := strings.TrimSpace(ctx.PostValue("school"))

	bio := strings.TrimSpace(ctx.PostValue("bio"))
	birthday := strings.TrimSpace(ctx.PostValue("birthday"))

	_, err := dbClient.User.Update().Where(user.ID(id)).SetBirthday(birthday).SetName(name).SetSchoolName(school).SetSystemName(username).SetBio(bio).Save(ctx)
	curr_user, err = dbClient.User.Query().Where(user.ID(curr_user.ID)).Only(ctx)
	ctx.SetUser(curr_user)
	sessID := ctx.GetCookie("_session_id")
	user_json, _ := json.Marshal(curr_user)

	rdb.Set(ctx, sessID, user_json, 0).Result()

	if err != nil {
		log.Println("failed when update user info: ", err)
		ctx.JSON(iris.Map{"err_msg": err.Error(), "status": iris.StatusBadRequest})
		return
	}

	ctx.JSON(iris.Map{"data": "updated", "status": iris.StatusOK})
}

func UpdateUserInfoSingle(ctx iris.Context) {

	curr_user := GetCurrentUserData(ctx)
	id := curr_user.ID
	name := strings.TrimSpace(ctx.PostValue("item_name"))
	value := strings.TrimSpace(ctx.PostValue("value"))

	var err error

	switch name {
	case "name":
		_, err = dbClient.User.Update().Where(user.ID(id)).SetName(value).Save(ctx)
	case "usernameID":
		_, err = dbClient.User.Update().Where(user.ID(id)).SetSystemName(value).Save(ctx)
	case "school":
		_, err = dbClient.User.Update().Where(user.ID(id)).SetSchoolName(value).Save(ctx)
	case "bio":
		_, err = dbClient.User.Update().Where(user.ID(id)).SetBio(value).Save(ctx)
	case "birthday":
		star := utils.GetConstellation(value)
		_, err = dbClient.User.Update().Where(user.ID(id)).SetBirthday(value).SetConstellation(star).Save(ctx)
	default:
		log.Println("user_info update with name", name)
	}

	if err != nil {
		ctx.JSON(iris.Map{"err_msg": err.Error(), "status": iris.StatusBadRequest})
		return
	}

	curr_user, err = dbClient.User.Query().Where(user.ID(curr_user.ID)).Only(ctx)
	ctx.SetUser(curr_user)
	sessID := ctx.GetCookie("_session_id")
	user_json, _ := json.Marshal(curr_user)

	rdb.Set(ctx, sessID, user_json, 0).Result()

	if err != nil {
		log.Println("failed when update user info: ", err)
		ctx.JSON(iris.Map{"err_msg": err.Error(), "status": iris.StatusBadRequest})
		return
	}

	ctx.JSON(iris.Map{"data": "updated", "status": iris.StatusOK})

}

func UpdateUserCollectionSwitch(ctx iris.Context) {
	curr_user := GetCurrentUserData(ctx)
	id := curr_user.ID

	switch_bool := false

	switch_value := strings.TrimSpace(ctx.PostValue("switch_value"))
	if switch_value == "1" {
		switch_bool = true
	}

	_, err := dbClient.User.Update().Where(user.ID(id)).SetIsShowCollections(switch_bool).Save(ctx)
	if err != nil {
		log.Println("failed when get all users: ", err)
		ctx.JSON(iris.Map{"err_msg": err.Error(), "status": iris.StatusBadRequest})
		return

	}
	curr_user, err = dbClient.User.Query().Where(user.ID(curr_user.ID)).Only(ctx)
	ctx.SetUser(curr_user)
	sessID := ctx.GetCookie("_session_id")
	user_json, _ := json.Marshal(curr_user)

	rdb.Set(ctx, sessID, user_json, 0).Result()

	if err != nil {
		log.Println("failed when update user info: ", err)
		ctx.JSON(iris.Map{"err_msg": err.Error(), "status": iris.StatusBadRequest})
		return
	}

	ctx.JSON(iris.Map{"data": "updated", "status": iris.StatusOK})
}

func UserSearchBySystemName(ctx iris.Context) {
	curr_user := GetCurrentUserData(ctx)
	system_name := strings.TrimSpace(ctx.URLParam("name"))
	var users []*ent.User
	var err error
	var ifresp SearchFriendsResp
	var friend_item_ids []uint64
	err = dbClient.Friendship.Query().Where(friendship.UserID(curr_user.ID), friendship.StatusEQ(friendship.StatusEstablished)).Select(friendship.FieldFriendID).Scan(ctx, &friend_item_ids)
	if system_name == "" {
		users, err = dbClient.User.Query().Where(user.IDIn(friend_item_ids...)).All(ctx)
	} else {

		users, err = dbClient.User.Query().Where(user.IDNEQ(curr_user.ID), user.SystemNameContains(system_name)).All(ctx)
		if len(users) == 0 {
			users, err = dbClient.User.Query().Where(user.IDNEQ(curr_user.ID), user.MobileNoContains(system_name)).All(ctx)
		}
	}

	if err != nil {
		log.Println("failed when UserSearchBySystemName: ", err)
	}
	for _, item := range users {
		if slices.Contains(friend_item_ids, item.ID) {

			var userStatus UserNameStatus
			userStatus.UserID = item.ID
			userStatus.UserName = item.Name
			userStatus.UserStatus = "in " + item.CurrentCsFieldName

			userStatus.UserThumbnail = item.ThumbnailURL
			ifresp.UserListSuggested = append(ifresp.UserListSuggested, userStatus)
			// ifresp.UserListMore = append(ifresp.UserListMore, userStatus)
		}
	}
	ctx.JSON(iris.Map{"data": ifresp, "status": iris.StatusOK})
}

func UserSearchAddOthers(ctx iris.Context) {
	curr_user := GetCurrentUserData(ctx)
	system_name := strings.TrimSpace(ctx.URLParam("name"))
	var users []*ent.User
	var err error
	var ifresp SearchFriendsResp
	if system_name == "" {
		users, err = dbClient.User.Query().Where(user.IDNEQ(curr_user.ID)).Order(ent.Desc(user.FieldID)).Limit(10).All(ctx)
	} else {
		users, err = dbClient.User.Query().Where(user.IDNEQ(curr_user.ID), user.SystemNameContains(system_name)).All(ctx)
		if len(users) == 0 {
			users, err = dbClient.User.Query().Where(user.IDNEQ(curr_user.ID), user.MobileNoContainsFold(system_name)).All(ctx)
		}
		if len(users) == 0 {
			users, err = dbClient.User.Query().Where(user.IDNEQ(curr_user.ID), user.NameContains(system_name)).All(ctx)
		}
	}

	var item_ids []uint64
	err = dbClient.Friendship.Query().Where(friendship.UserID(curr_user.ID), friendship.StatusEQ(friendship.StatusEstablished)).Select(friendship.FieldFriendID).Scan(ctx, &item_ids)
	var invited_item_ids []uint64
	err = dbClient.Friendship.Query().Where(friendship.UserID(curr_user.ID), friendship.StatusEQ(friendship.StatusInvite)).Select(friendship.FieldFriendID).Scan(ctx, &invited_item_ids)
	if err != nil {
		log.Println("failed when UserSearchAddOthers: ", err)
	}
	for _, item := range users {
		if !slices.Contains(item_ids, item.ID) {

			var userStatus UserNameStatus
			userStatus.UserID = item.ID
			userStatus.UserName = item.Name
			userStatus.UserStatus = "in " + item.CurrentCsFieldName
			if slices.Contains(invited_item_ids, item.ID) {
				userStatus.Status = "Wating..."
				userStatus.CanAction = false
			} else {
				userStatus.Status = "Add"
				userStatus.ActionUrl = fmt.Sprintf("action?type=add&target_id=%d", item.ID)
				userStatus.CanAction = true
			}

			userStatus.UserThumbnail = item.ThumbnailURL
			ifresp.UserListSuggested = append(ifresp.UserListSuggested, userStatus)
			// ifresp.UserListMore = append(ifresp.UserListMore, userStatus)
		}
	}

	ctx.JSON(iris.Map{"data": ifresp, "status": iris.StatusOK})
}

func GetCollections(ctx iris.Context) {
	curr_user := GetCurrentUserData(ctx)
	var item_ids []uint64

	err := dbClient.Collection.Query().Where(collection.UserID(curr_user.ID)).Select(collection.FieldItemID).Scan(ctx, &item_ids)
	collections, err := dbClient.TimeDew.Query().Where(timedew.IDIn(item_ids...)).WithOwner().All(ctx)
	var tdresp TimeDewResp
	for _, item := range collections {
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
		} else if item.Type == timedew.TypeUser {
			lf.Type = "single"
		} else if item.Type == timedew.TypeSystem {
			lf.Type = "system"
		} else if item.Type == timedew.TypeInvite {
			lf.Type = "invite"
		}
		self_td := false

		if item.UserID == curr_user.ID {
			self_td = true
		}
		if item.Edges.Reactions != nil {
			self_react := false
			for _, react := range item.Edges.Reactions {

				if react.UserID == curr_user.ID {
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
		log.Println("failed when get collection info: ", err)
		ctx.JSON(iris.Map{"err_msg": "failed when get friend info", "status": iris.StatusInternalServerError})
		return
	}
	ctx.JSON(iris.Map{"data": tdresp, "status": iris.StatusOK})

}

func AddCollection(ctx iris.Context) {
	curr_user := GetCurrentUserData(ctx)
	item_id := strings.TrimSpace(ctx.PostValue("item_id"))
	item_id_int, _ := strconv.ParseUint(item_id, 10, 64)
	item_type := strings.TrimSpace(ctx.PostValue("item_type"))
	td, err := dbClient.TimeDew.Query().Where(timedew.ID(item_id_int)).Only(ctx)
	if td.UserID != curr_user.ID {
		ctx.JSON(iris.Map{"err_msg": "only can save self timedew", "status": iris.StatusBadRequest})
		return
	}

	collection, err := dbClient.Collection.Create().SetOwnerID(curr_user.ID).SetItemID(item_id_int).SetType(collection.TypeTimedew).Save(ctx)
	if err != nil {
		log.Println("failed when AddCollection: ", err, item_type)
		ctx.JSON(iris.Map{"err_msg": "fail to add collection", "status": iris.StatusInternalServerError})
		return
	}
	ctx.JSON(iris.Map{"data": collection, "status": iris.StatusOK})

}

func DeleteCollection(ctx iris.Context) {
	curr_user := GetCurrentUserData(ctx)
	id := strings.TrimSpace(ctx.URLParam("id"))
	id_int, _ := strconv.ParseUint(id, 10, 64)

	_, err := dbClient.Collection.Delete().Where(collection.UserID(curr_user.ID), collection.ID(id_int)).Exec(ctx)
	if err != nil {
		log.Println("failed when delete collection: ", err)
		ctx.JSON(iris.Map{"err_msg": "failed when delete collection", "status": iris.StatusInternalServerError})
		return
	}
	ctx.JSON(iris.Map{"data": "deleted", "status": iris.StatusOK})

}

func GetUserInfoByID(ctx iris.Context) {
	currUser := GetCurrentUserData(ctx)
	userID := strings.TrimSpace(ctx.Params().Get("id"))
	userIDInt, _ := strconv.ParseUint(userID, 10, 64)
	friendship_status := friendship.CurrTypeNone
	isWaiting := false
	isHidden := false
	var item_ids []uint64
	userInfo, err := dbClient.User.Query().Where(user.IDEQ(userIDInt)).Only(ctx)
	if err != nil {
		log.Println("failed when update user info: ", err)
		ctx.JSON(iris.Map{"err_msg": err.Error(), "status": iris.StatusBadRequest})
		return
	}
	friend, err := dbClient.Friendship.Query().Where(friendship.UserID(currUser.ID), friendship.FriendID(userIDInt)).WithUser().Only(ctx)
	if err == nil {
		friendship_status = friend.CurrType
		if friend.Status == friendship.StatusInvite {
			isWaiting = true
		}
	} else {
		log.Println(err)
	}

	isHidden, _ = dbClient.Hidden.Query().Where(hidden.UserID(currUser.ID), hidden.HiddenID(userIDInt)).Exist(ctx)

	var tdresp TimeDewResp
	if userInfo.IsShowCollections {
		err = dbClient.Collection.Query().Where(collection.UserID(userIDInt)).Select(collection.FieldItemID).Scan(ctx, &item_ids)
		collections, _ := dbClient.TimeDew.Query().Where(timedew.IDIn(item_ids...)).WithOwner().All(ctx)

		for _, item := range collections {
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
			if item.Edges.Reactions != nil {
				for _, react := range item.Edges.Reactions {
					if react.IsCool {
						name, err := dbClient.User.Query().Where(user.IDEQ(react.UserID)).Select(user.FieldName).Strings(ctx)
						if err == nil {
							lf.Reactions.Cool = append(lf.Reactions.Cool, UserNameID{UserName: name[0], UserID: react.UserID})
						}

					}
					if react.IsDAMN {
						name, err := dbClient.User.Query().Where(user.IDEQ(react.UserID)).Select(user.FieldName).Strings(ctx)
						if err == nil {
							lf.Reactions.DAMN = append(lf.Reactions.DAMN, UserNameID{UserName: name[0], UserID: react.UserID})
						}

					}
					if react.IsLOL {
						name, err := dbClient.User.Query().Where(user.IDEQ(react.UserID)).Select(user.FieldName).Strings(ctx)
						if err == nil {
							lf.Reactions.LOL = append(lf.Reactions.LOL, UserNameID{UserName: name[0], UserID: react.UserID})
						}

					}
					if react.IsNooo {
						name, err := dbClient.User.Query().Where(user.IDEQ(react.UserID)).Select(user.FieldName).Strings(ctx)
						if err == nil {
							lf.Reactions.Nooo = append(lf.Reactions.Nooo, UserNameID{UserName: name[0], UserID: react.UserID})
						}

					}
					if react.IsOMG {
						name, err := dbClient.User.Query().Where(user.IDEQ(react.UserID)).Select(user.FieldName).Strings(ctx)
						if err == nil {
							lf.Reactions.OMG = append(lf.Reactions.OMG, UserNameID{UserName: name[0], UserID: react.UserID})
						}

					}
				}
			}
			tdresp.LifeFlowTimeStamp = item.CreateTime.UnixMilli()
			tdresp.LifeFlow = append(tdresp.LifeFlow, lf)
		}
	}

	ctx.JSON(iris.Map{"data": iris.Map{"user_info": userInfo, "collections": tdresp, "friendship_status": friendship_status, "is_waiting": isWaiting, "is_hidden": isHidden}, "status": iris.StatusOK})
}

func GetCurrUser(ctx iris.Context) {
	curr_user := GetCurrentUserData(ctx)
	curr_user, err := dbClient.User.Query().Where(user.ID(curr_user.ID)).Only(ctx)
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

func GetFriends(ctx iris.Context) {
	curr_user := GetCurrentUserData(ctx)
	closeFriends, err := dbClient.Friendship.Query().Where(friendship.FriendID(curr_user.ID), friendship.CurrTypeEQ(friendship.CurrTypeClose)).WithUser().All(ctx)
	friends, err := dbClient.Friendship.Query().Where(friendship.FriendID(curr_user.ID), friendship.CurrTypeEQ(friendship.CurrTypeGood)).WithUser().All(ctx)
	all, err := dbClient.Friendship.Query().Where(friendship.FriendID(curr_user.ID), friendship.StatusEQ(friendship.StatusEstablished)).WithUser().All(ctx)
	friendsRequests, err := dbClient.Friendship.Query().Where(friendship.FriendID(curr_user.ID), friendship.StatusEQ(friendship.StatusInvite)).All(ctx)
	var normal, good, closed FriendsNotice
	for _, item := range friendsRequests {
		if item.RequestType == friendship.RequestTypeNormal {

			var userStatus UserNameStatus
			userStatus.UserID = item.UserID
			u, err := dbClient.User.Query().Where(user.IDEQ(item.UserID)).Only(ctx)
			if err == nil {
				userStatus.UserName = u.Name
				userStatus.UserStatus = "in " + u.CurrentCsFieldName

				userStatus.UserThumbnail = u.ThumbnailURL
			}

			normal.UserList = append(normal.UserList, userStatus)
			normal.Content = fmt.Sprintf("%s wants to be connection with you", userStatus.UserName)
		}
		if item.RequestType == friendship.RequestTypeGood {

			var userStatus UserNameStatus
			userStatus.UserID = item.UserID

			u, err := dbClient.User.Query().Where(user.IDEQ(item.UserID)).Only(ctx)
			if err == nil {
				userStatus.UserName = u.Name
				userStatus.UserStatus = "in " + u.CurrentCsFieldName
				userStatus.UserThumbnail = u.ThumbnailURL
			}

			good.UserList = append(good.UserList, userStatus)
			good.Content = fmt.Sprintf("%s wants to be connection with you", userStatus.UserName)
		}
		if item.RequestType == friendship.RequestTypeClose {

			var userStatus UserNameStatus
			userStatus.UserID = item.UserID
			u, err := dbClient.User.Query().Where(user.IDEQ(item.UserID)).Only(ctx)
			if err != nil {
				userStatus.UserName = u.Name
				userStatus.UserStatus = "in " + u.CurrentCsFieldName
				userStatus.UserThumbnail = u.ThumbnailURL
			}

			closed.UserList = append(normal.UserList, userStatus)
			closed.Content = fmt.Sprintf("%s wants to be your close friend", userStatus.UserName)
		}
	}
	var fn FriendsNoticeResp
	if len(good.UserList) > 0 {
		fn.FriendsNoticeList = append(fn.FriendsNoticeList, good)
	}
	if len(normal.UserList) > 0 {
		fn.FriendsNoticeList = append(fn.FriendsNoticeList, normal)
	}
	if len(closed.UserList) > 0 {
		fn.FriendsNoticeList = append(fn.FriendsNoticeList, closed)
	}

	if err != nil {
		log.Println("failed when get all users: ", err)
	}
	ctx.JSON(iris.Map{"data": iris.Map{"friends_notice": fn, "close_friends": closeFriends, "close_friends_nums": fmt.Sprintf("%d/%d", len(closeFriends), len(closeFriends)), "friends": friends, "friends_nums": fmt.Sprintf("%d/%d", len(friends), len(friends)), "all": all, "all_nums": fmt.Sprintf("%d/%d", len(friends), len(friends))}, "status": iris.StatusOK})
}

func GetFriendsInvitations(ctx iris.Context) {
	curr_user := GetCurrentUserData(ctx)
	friends, err := dbClient.User.Query().Where(user.IDEQ(curr_user.ID), predicate.User(friendship.StatusNEQ(friendship.StatusEstablished))).QueryFriends().WithFriendships().All(ctx)
	if err != nil {
		log.Println("failed when get all users: ", err)
	}
	ctx.JSON(iris.Map{"data": friends, "status": iris.StatusOK})
}

func GetFriendsAskedFor(ctx iris.Context) {
	curr_user := GetCurrentUserData(ctx)
	friends, err := dbClient.User.Query().Where(user.IDEQ(curr_user.ID), predicate.User(friendship.StatusNEQ(friendship.StatusEstablished))).QueryFriends().WithFriendships().All(ctx)
	if err != nil {
		log.Println("failed when get all users: ", err)
	}
	ctx.JSON(iris.Map{"data": friends, "status": iris.StatusOK})
}

func AddFriend(ctx iris.Context) {

	target_id, _ := strconv.ParseUint(ctx.URLParam("target_id"), 10, 64)

	curr_user := GetCurrentUserData(ctx)
	ok, err := dbClient.User.Update().Where(user.IDEQ(curr_user.ID)).AddFriendIDs(target_id).Save(ctx)
	dbClient.Friendship.Update().Where(friendship.UserID(curr_user.ID), friendship.FriendID(target_id)).SetStatus(friendship.StatusEstablished).Save(ctx)

	if err != nil {
		log.Println("failed when get all users: ", err)
		ctx.JSON(iris.Map{"err_msg": err.Error(), "status": iris.StatusBadRequest})
		return
	}
	ctx.JSON(iris.Map{"data": ok, "status": iris.StatusOK})
}

func ActionForFriend(ctx iris.Context) {
	currUser := GetCurrentUserData(ctx)
	target_id := ctx.URLParam("target_id")
	target_idInt, _ := strconv.ParseUint(target_id, 10, 64)
	action_type := strings.TrimSpace(ctx.URLParam("type"))
	if action_type == "add" {
		dbClient.User.Update().Where(user.IDEQ(currUser.ID)).AddFriendIDs(target_idInt).Save(ctx)
		dbClient.Friendship.Update().Where(friendship.UserID(currUser.ID), friendship.FriendID(target_idInt)).SetStatus(friendship.StatusInvite).SetRequestType(friendship.RequestTypeNormal).Save(ctx)
		dbClient.Friendship.Update().Where(friendship.FriendID(currUser.ID), friendship.UserID(target_idInt)).SetStatus(friendship.StatusInvited).SetRequestType(friendship.RequestTypeNormal).Save(ctx)
	} else if action_type == "invite" {
		err := genInviteTimedew(currUser, target_idInt)
		// v := make(map[uint64]bool)
		// v[target_idInt] = true
		// vbyte, _ := json.Marshal(v)
		_, Rerr := rdb.HMSet(ctx, fmt.Sprintf("%d_%d_invite_map", currUser.ID, currUser.CurrentCsFieldID), target_idInt, true).Result()
		if Rerr != nil {
			log.Println("invite_map failed, err ", Rerr.Error())
		}
		if err != nil {
			ctx.JSON(iris.Map{"err_msg": err.Error(), "status": iris.StatusBadRequest})
			return
		}
	}
	ctx.JSON(iris.Map{"data": "ok", "status": iris.StatusOK})

}

func InviteFriendShip(ctx iris.Context) {

	target_id, _ := strconv.ParseUint(strings.TrimSpace(ctx.PostValue("target_id")), 10, 64)
	friendship_type := strings.TrimSpace(ctx.PostValue("friendship_type"))
	var friendship_type_enum friendship.RequestType
	target, err := dbClient.User.Query().Where(user.ID(target_id)).Only(ctx)
	if friendship_type == "" || friendship_type == "normal" {
		friendship_type_enum = friendship.RequestTypeNormal
	} else if friendship_type == "close" {
		friendship_type_enum = friendship.RequestTypeClose
	} else if friendship_type == "good" {
		friendship_type_enum = friendship.RequestTypeGood
	}
	curr_user := GetCurrentUserData(ctx)
	log.Println("invite friend from ", curr_user.ID, target_id)
	fp, err := dbClient.Friendship.Query().Where(friendship.UserID(target_id), friendship.FriendID(curr_user.ID)).Only(ctx)
	if err == nil && fp != nil {
		if fp.Status == friendship.StatusInvite {
			dbClient.Friendship.Update().Where(friendship.UserID(curr_user.ID), friendship.FriendID(target.ID)).SetStatus(friendship.StatusEstablished).SetRequestType(friendship_type_enum).Save(ctx)
			dbClient.Friendship.Update().Where(friendship.UserID(target.ID), friendship.FriendID(curr_user.ID)).SetStatus(friendship.StatusEstablished).SetRequestType(friendship_type_enum).Save(ctx)
		}
	} else {
		_, err = dbClient.User.Update().Where(user.IDEQ(curr_user.ID)).AddFriendIDs(target.ID).Save(ctx)
		dbClient.Friendship.Update().Where(friendship.UserID(curr_user.ID), friendship.FriendID(target.ID)).SetStatus(friendship.StatusInvite).SetRequestType(friendship_type_enum).Save(ctx)
		dbClient.Friendship.Update().Where(friendship.FriendID(curr_user.ID), friendship.UserID(target.ID)).SetStatus(friendship.StatusInvited).SetRequestType(friendship_type_enum).Save(ctx)
	}

	ctx.JSON(iris.Map{"data": "invited", "status": iris.StatusOK})
}

func DealFriendShipInvitation(ctx iris.Context) {
	op := strings.TrimSpace(ctx.PostValue("op"))
	target_id, _ := strconv.ParseUint(strings.TrimSpace(ctx.PostValue("target_id")), 10, 64)
	friendship_type := strings.TrimSpace(ctx.PostValue("friendship_type"))
	var friendship_type_enum friendship.CurrType
	if friendship_type == "" || friendship_type == "normal" {
		friendship_type_enum = friendship.CurrTypeNormal
	} else if friendship_type == "close" {
		friendship_type_enum = friendship.CurrTypeClose
	} else if friendship_type == "good" {
		friendship_type_enum = friendship.CurrTypeGood
	} else if friendship_type == "none" {
		friendship_type_enum = friendship.CurrTypeNone
	}
	curr_user := GetCurrentUserData(ctx)
	if op == "accept" {
		dbClient.Friendship.Update().Where(friendship.UserID(curr_user.ID), friendship.FriendID(target_id)).SetStatus(friendship.StatusEstablished).SetCurrType(friendship_type_enum).Save(ctx)
		dbClient.Friendship.Update().Where(friendship.FriendID(curr_user.ID), friendship.UserID(target_id)).SetStatus(friendship.StatusEstablished).SetCurrType(friendship_type_enum).Save(ctx)
	} else {
		dbClient.Friendship.Update().Where(friendship.UserID(curr_user.ID), friendship.FriendID(target_id)).SetStatus(friendship.StatusRejected).SetCurrType(friendship_type_enum).Save(ctx)
		dbClient.Friendship.Update().Where(friendship.FriendID(curr_user.ID), friendship.UserID(target_id)).SetStatus(friendship.StatusRejected).SetCurrType(friendship_type_enum).Save(ctx)
	}
	total, _ := dbClient.Friendship.Query().Where(friendship.FriendID(curr_user.ID), friendship.StatusEQ(friendship.StatusEstablished)).Count(ctx)
	dbClient.User.UpdateOneID(curr_user.ID).SetTotalConnections(total).Exec(ctx)

	ctx.JSON(iris.Map{"data": op, "status": iris.StatusOK})
}

func InviteFriends(ctx iris.Context) {
	currUser := GetCurrentUserData(ctx)
	var item_ids []uint64
	err := dbClient.Friendship.Query().Where(friendship.FriendID(currUser.ID)).Select(friendship.FieldUserID).Scan(ctx, &item_ids)
	item_ids = append(item_ids, currUser.ID)
	users, err := dbClient.User.Query().Where(user.IDNotIn(item_ids...)).Order(ent.Desc(user.FieldID)).Limit(10).All(ctx)
	var ifresp InviteFriendsResp

	if err != nil {
		log.Println("failed when InviteFriends get users: ", err)
		ctx.JSON(iris.Map{"err_msg": err, "status": iris.StatusBadRequest})
		return
	}
	for _, item := range users {
		var userStatus UserNameStatus
		userStatus.UserID = item.ID
		userStatus.UserName = item.Name
		userStatus.UserStatus = ""
		userStatus.Status = "Add"
		if userStatus.Status == "Add" || userStatus.Status == "Invite" {
			userStatus.CanAction = true
			userStatus.ActionUrl = fmt.Sprintf("action?type=add&target_id=%d", item.ID)
		} else {
			userStatus.CanAction = false
		}
		userStatus.UserThumbnail = item.ThumbnailURL
		ifresp.InviteUrl = "http://zingy.social"
		ifresp.UserList = append(ifresp.UserList, userStatus)
	}
	ctx.JSON(iris.Map{"data": ifresp, "status": iris.StatusOK})
}

func ConectionRequest(ctx iris.Context) {
	currUser := GetCurrentUserData(ctx)
	requsets, err := dbClient.Friendship.Query().Where(friendship.FriendID(currUser.ID), friendship.StatusEQ(friendship.StatusInvite)).All(ctx)
	if err != nil {
		log.Println("failed when ConectionRequest get users: ", err)
		ctx.JSON(iris.Map{"err_msg": err, "status": iris.StatusBadRequest})
		return
	}
	var res ConectionRequestResp
	for _, item := range requsets {
		var userStatus UserNameStatus
		userStatus.UserID = item.UserID
		u, err := dbClient.User.Query().Where(user.IDEQ(item.UserID)).Only(ctx)
		if err == nil {
			userStatus.UserName = u.Name
			userStatus.UserStatus = "wants to be your friend"
			userStatus.UserThumbnail = u.ThumbnailURL
		}
		res.UserList = append(res.UserList, userStatus)
	}
	ctx.JSON(iris.Map{"data": res, "status": iris.StatusOK})
}

func BreakConnection(ctx iris.Context) {
	var br brReq
	ctx.ReadJSON(&br)
	target_id := br.TargetID
	url_target_id, err := strconv.ParseUint(strings.TrimSpace(ctx.URLParam("target_id")), 10, 64)
	if err == nil {
		target_id = url_target_id
	}
	currUser := GetCurrentUserData(ctx)
	_, err = dbClient.Friendship.Delete().Where(friendship.UserID(currUser.ID), friendship.FriendID(target_id)).Exec(ctx)
	_, err = dbClient.Friendship.Delete().Where(friendship.FriendID(currUser.ID), friendship.UserID(target_id)).Exec(ctx)
	if err != nil {
		log.Println("failed when BreakConnection get users: ", err)
		ctx.JSON(iris.Map{"err_msg": err, "status": iris.StatusBadRequest})
		return
	}
	total, _ := dbClient.Friendship.Query().Where(friendship.FriendID(currUser.ID), friendship.StatusEQ(friendship.StatusEstablished)).Count(ctx)
	dbClient.User.UpdateOneID(currUser.ID).SetTotalConnections(total).Exec(ctx)

	ctx.JSON(iris.Map{"data": "ok", "status": iris.StatusOK})
}

func GetCards(ctx iris.Context) {
	curr_user := GetCurrentUserData(ctx)
	friends, err := dbClient.Card.Query().Where(card.HasOwnerWith(user.IDEQ(curr_user.ID))).All(ctx)
	if err != nil {
		log.Println("failed when get all users: ", err)
	}
	ctx.JSON(iris.Map{"data": friends, "status": iris.StatusOK})
}

func CreateCard(ctx iris.Context) {

	name := strings.TrimSpace(ctx.PostValue("name"))
	curr_user := GetCurrentUserData(ctx)

	card, err := dbClient.Card.Create().SetName(name).SetOwnerID(curr_user.ID).SetStatus(card.StatusStatus1).SetType(card.TypeType1).Save(ctx)

	if err != nil {
		log.Println("failed when create card: ", err)
		ctx.JSON(iris.Map{"err_msg": err, "status": iris.StatusBadRequest})
		return
	}
	ctx.JSON(iris.Map{"data": card, "status": iris.StatusOK})
}

func FieldSearch(ctx iris.Context) {
	key := strings.TrimSpace(ctx.URLParam("key"))
	fields, err := dbClient.CSField.Query().Where(csfield.NameContains(key)).All(ctx)
	if err != nil {
		log.Println("failed when get all users: ", err)
	}
	ctx.JSON(iris.Map{"data": fields, "status": iris.StatusOK})
}

func GetCSFields(ctx iris.Context) {

	currUser := GetCurrentUserData(ctx)
	if currUser == nil {
		ctx.JSON(iris.Map{"err_msg": "need login", "status": iris.StatusBadRequest})
		return
	}
	csfields, err := dbClient.CSField.Query().Where(csfield.HasJoinedUserWith(user.IDEQ(currUser.ID))).WithJoinedUser().All(ctx)
	if err != nil {
		log.Println("failed when get all users: ", err)
		ctx.JSON(iris.Map{"err_msg": err.Error(), "status": iris.StatusBadRequest})
		return
	}
	ctx.JSON(iris.Map{"data": csfields, "status": iris.StatusOK})
}

func GetCSFieldByID(ctx iris.Context) {
	csFieldID, _ := strconv.ParseUint(strings.TrimSpace(ctx.Params().Get("cs_field_id")), 10, 64)
	currUser := GetCurrentUserData(ctx)
	if currUser == nil {
		ctx.JSON(iris.Map{"err_msg": "need login", "status": iris.StatusBadRequest})
		return
	}
	csfield, err := dbClient.CSField.Query().Where(csfield.ID(csFieldID)).WithJoinedUser().Only(ctx)
	if err != nil {
		log.Println("failed when get all users: ", err)
		ctx.JSON(iris.Map{"err_msg": err.Error(), "status": iris.StatusBadRequest})
		return
	}
	liveStatus, err := rdb.Get(ctx, fmt.Sprintf("csfl_%d", csFieldID)).Result()
	if err != nil {
		ctx.JSON(iris.Map{"err_msg": err.Error(), "status": iris.StatusBadRequest})
		return
	}
	var liveStruct interface{}
	err = json.Unmarshal([]byte(liveStatus), &liveStruct)

	ctx.JSON(iris.Map{"data": iris.Map{"cs_field": csfield, "live_status": liveStruct}, "status": iris.StatusOK})
}

func GetCSFieldLiveByID(ctx iris.Context) {
	csFieldID := strings.TrimSpace(ctx.Params().Get("cs_field_id"))
	data, err := rdb.Get(ctx, fmt.Sprintf("csfl_%s", csFieldID)).Result()
	if err != nil {
		ctx.JSON(iris.Map{"err_msg": err.Error(), "status": iris.StatusBadRequest})
		return
	} else {
		var liveStruct interface{}
		json.Unmarshal([]byte(data), &liveStruct)
		ctx.JSON(iris.Map{"data": liveStruct, "status": iris.StatusOK})
	}
}

func UpdateCSFieldLiveByID(ctx iris.Context) {
	csFieldID := strings.TrimSpace(ctx.Params().Get("cs_field_id"))
	liveStatus := ctx.PostValue("live_status")
	log.Println("cs field update liveStatus: ", csFieldID, liveStatus)
	_, err := rdb.Set(ctx, fmt.Sprintf("csfl_%s", csFieldID), liveStatus, 0).Result()
	if err != nil {
		ctx.JSON(iris.Map{"err_msg": err.Error(), "status": iris.StatusBadRequest})
	} else {
		ctx.JSON(iris.Map{"data": liveStatus, "status": iris.StatusOK})
	}
}

func CreateCSField(ctx iris.Context) {

	curr_user := GetCurrentUserData(ctx)
	name := strings.TrimSpace(ctx.PostValue("name"))
	csfield, err := dbClient.CSField.Create().SetStatus(csfield.StatusOpening).SetType(csfield.TypeEmpty).SetMode(csfield.ModeSingle).SetName(name).SetUserID(curr_user.ID).AddJoinedUserIDs(curr_user.ID).Save(ctx)
	if err != nil {
		log.Println("failed when create field: ", err)
	}
	ctx.JSON(iris.Map{"data": csfield, "status": iris.StatusOK})
}

func CSFieldPrivateLevel(ctx iris.Context) {

	// curr_user := GetCurrentUserData(ctx)
	private_level := strings.TrimSpace(ctx.PostValue("private_level"))
	var pl csfield.PrivateLevel
	if private_level == "public" {
		pl = csfield.PrivateLevelPublic
	}
	if private_level == "private" {
		pl = csfield.PrivateLevelPrivate
	}
	if private_level == "ghost" {
		pl = csfield.PrivateLevelGhost
	}
	cs_field_id, _ := strconv.ParseUint(strings.TrimSpace(ctx.PostValue("cs_field_id")), 10, 64)
	csfield, err := dbClient.CSField.UpdateOneID(cs_field_id).SetPrivateLevel(pl).Save(ctx)
	if err != nil {
		log.Println("failed when create field: ", err)
	}
	ctx.JSON(iris.Map{"data": csfield, "status": iris.StatusOK})
}

func CreateCSFieldByUserID(ctx iris.Context) {

	name := strings.TrimSpace(ctx.PostValue("name"))
	user_id, _ := strconv.ParseUint(ctx.PostValue("user_id"), 10, 64)
	csfield, err := dbClient.CSField.Create().SetStatus(csfield.StatusOpening).SetType(csfield.TypeEmpty).SetName(name).AddJoinedUserIDs(user_id).Save(ctx)
	if err != nil {
		ctx.JSON(iris.Map{"err_msg": err.Error(), "status": iris.StatusBadRequest})
		return
	}
	ctx.JSON(csfield)
}

func InviteJoinCSField(ctx iris.Context) {
	currUser := GetCurrentUserData(ctx)
	target_id, err := strconv.ParseUint(ctx.PostValue("target_id"), 10, 64)
	err = genInviteTimedew(currUser, target_id)
	if err != nil {
		ctx.JSON(iris.Map{"err_msg": err.Error(), "status": iris.StatusBadRequest})
		return
	}
	ctx.JSON(iris.Map{"data": "ok", "status": iris.StatusOK})
}

func GetSettings(ctx iris.Context) {

	curr_user := GetCurrentUserData(ctx)
	curSetting, err := dbClient.Setting.Query().Where(setting.HasOwnerWith(user.IDEQ(curr_user.ID))).Only(ctx)
	if err != nil {
		if err.Error() == "ent: setting not found" {
			curSetting, err = dbClient.Setting.Create().SetOwnerID(curr_user.ID).Save(ctx)
			if err != nil {
				ctx.JSON(iris.Map{"err_msg": err.Error(), "status": iris.StatusBadRequest})
				return
			} else {
				ctx.JSON(iris.Map{"data": curSetting, "status": iris.StatusOK})
				return
			}
		}
		ctx.JSON(iris.Map{"err_msg": err.Error(), "status": iris.StatusBadRequest})
		return
	}

	ctx.JSON(iris.Map{"data": curSetting, "status": iris.StatusOK})
}

func CreateSetting(ctx iris.Context) {
	curr_user := GetCurrentUserData(ctx)
	setting, err := dbClient.Setting.Create().SetOwnerID(curr_user.ID).Save(ctx)
	if err != nil {
		ctx.JSON(iris.Map{"err_msg": err.Error(), "status": iris.StatusBadRequest})
		return
	}
	ctx.JSON(iris.Map{"data": setting, "status": iris.StatusOK})
}

func UpdateSetting(ctx iris.Context) {
	curr_user := GetCurrentUserData(ctx)
	switch_bool := false
	name := strings.TrimSpace(ctx.PostValue("name"))
	switch_value := strings.TrimSpace(ctx.PostValue("switch_value"))
	if switch_value == "1" {
		switch_bool = true
	}
	var err error

	switch name {
	case "friends_online":
		_, err = dbClient.Setting.Update().Where(setting.UserIDEQ(curr_user.ID)).SetFriendsOnline(switch_bool).Save(ctx)
	case "time_dew_from_friends":
		_, err = dbClient.Setting.Update().Where(setting.UserIDEQ(curr_user.ID)).SetTimeDewFromFriends(switch_bool).Save(ctx)
	case "detailed_notification":
		_, err = dbClient.Setting.Update().Where(setting.UserIDEQ(curr_user.ID)).SetDetailedNotification(switch_bool).Save(ctx)
	case "receive_field_invitation":
		_, err = dbClient.Setting.Update().Where(setting.UserIDEQ(curr_user.ID)).SetReceiveFieldInvitation(switch_bool).Save(ctx)
	case "see_my_location":
		_, err = dbClient.Setting.Update().Where(setting.UserIDEQ(curr_user.ID)).SetSeeMyLocation(switch_bool).Save(ctx)
	case "camera":
		_, err = dbClient.Setting.Update().Where(setting.UserIDEQ(curr_user.ID)).SetCamera(switch_bool).Save(ctx)
	case "microphone":
		_, err = dbClient.Setting.Update().Where(setting.UserIDEQ(curr_user.ID)).SetMicrophone(switch_bool).Save(ctx)
	case "health_data":
		_, err = dbClient.Setting.Update().Where(setting.UserIDEQ(curr_user.ID)).SetHealthData(switch_bool).Save(ctx)
	case "time_dew_location":
		_, err = dbClient.Setting.Update().Where(setting.UserIDEQ(curr_user.ID)).SetTimeDewLocation(switch_bool).Save(ctx)
	case "time_dew_microphone":
		_, err = dbClient.Setting.Update().Where(setting.UserIDEQ(curr_user.ID)).SetTimeDewMicrophone(switch_bool).Save(ctx)
	case "time_dew_Lora":
		_, err = dbClient.Setting.Update().Where(setting.UserIDEQ(curr_user.ID)).SetTimeDewLora(switch_bool).Save(ctx)
	default:
		log.Println("settings update with name", name)
	}

	if err != nil {
		ctx.JSON(iris.Map{"err_msg": err.Error(), "status": iris.StatusBadRequest})
		return
	}
	ctx.JSON(iris.Map{"data": "updated", "status": iris.StatusOK})
}

func GetV1Messages(ctx iris.Context) {
	messages, err := dbClient.Message.Query().Order(ent.Asc(message.FieldCreateTime)).WithOwner().All(ctx)
	if err != nil {
		ctx.JSON(iris.Map{"err_msg": err.Error(), "status": iris.StatusBadRequest})
		return
	} else {
		ctx.JSON(iris.Map{"data": messages, "status": iris.StatusOK})
	}
}

func PostV1Messages(ctx iris.Context) {
	currUser := GetCurrentUserData(ctx)
	content := strings.TrimSpace(ctx.PostValue("content"))
	message, _ := dbClient.Message.Create().SetOwnerID(currUser.ID).SetContent(content).Save(ctx)
	ctx.JSON(iris.Map{"data": message, "status": iris.StatusOK})
}

func GetV1Cards(ctx iris.Context) {
	cards, err := dbClient.Card.Query().All(ctx)
	if err != nil {
		ctx.JSON(iris.Map{"err_msg": err.Error(), "status": iris.StatusBadRequest})
		return
	} else {
		ctx.JSON(iris.Map{"data": cards, "status": iris.StatusOK})
	}
}

func DelectV1Card(ctx iris.Context) {
	card_id := ctx.URLParam("card_id")
	card_idInt, _ := strconv.ParseUint(card_id, 10, 64)
	err := dbClient.Card.DeleteOneID(card_idInt).Exec(ctx)
	if err != nil {
		ctx.JSON(iris.Map{"err_msg": err.Error(), "status": iris.StatusBadRequest})
		return
	} else {
		ctx.JSON(iris.Map{"data": fmt.Sprintf("card ID= %s deleted", card_id), "status": iris.StatusOK})
	}
}

func PostV1Cards(ctx iris.Context) {
	name := strings.TrimSpace(ctx.PostValue("name"))
	description := strings.TrimSpace(ctx.PostValue("description"))
	pic_url := strings.TrimSpace(ctx.PostValue("pic_url"))
	script_url := ""
	files, _, err := ctx.UploadFormFiles("./assets", preUpload)
	if err != nil {
		log.Println(err)
		ctx.StopWithStatus(iris.StatusInternalServerError)
		return
	}

	var file_urls []string

	for _, file := range files {
		file_url := fmt.Sprintf("%s%s/assets/%s", "http://", ctx.Host(), file.Filename)

		if strings.Contains(file.Filename, ".js") || strings.Contains(file.Filename, ".zip") {
			script_url = file_url
		} else if strings.Contains(file.Filename, ".png") || strings.Contains(file.Filename, ".jpg") || strings.Contains(file.Filename, ".jpeg") {
			pic_url = file_url
		}
		file_urls = append(file_urls, file_url)
	}

	cur_card, err := dbClient.Card.Create().SetName(name).SetDescription(description).SetPicURL(pic_url).SetThumbnailURL(pic_url).SetScriptURL(script_url).Save(ctx)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(iris.Map{"err_msg": err.Error(), "status": iris.StatusBadRequest})
		return
	}
	ctx.JSON(iris.Map{"data": cur_card, "status": iris.StatusOK})

}

func JoinInCSFieldByWS(userID string, cs_field_id string) (*ent.CSField, *ent.User) {

	ctx := context.Background()
	curUserID, _ := strconv.ParseUint(userID, 10, 64)
	fieldID, _ := strconv.ParseUint(cs_field_id, 10, 64)
	if fieldID == 0 || curUserID == 0 {
		return nil, nil
	}
	currUser, err := dbClient.User.Query().Where(user.ID(curUserID)).Only(ctx)
	field, err := dbClient.CSField.Query().Where(csfield.ID(fieldID)).WithJoinedUser().Only(ctx)
	if err != nil {
		return nil, nil
	}
	_, err = field.Update().AddJoinedUser(currUser).Save(ctx)
	if err != nil {
		log.Println("JoinInCSFieldByWS Error:", err.Error())
	}
	isExist, err := rdb.Exists(ctx, fmt.Sprintf("csfl_%s", cs_field_id)).Result()
	if isExist < 1 || err != nil {
		prop := map[string]int{"version": 0}
		propInit, _ := json.Marshal(prop)
		_, err = rdb.Set(ctx, fmt.Sprintf("csfl_%s", cs_field_id), propInit, 0).Result()
	}

	inviteMap, err := rdb.HGetAll(ctx, fmt.Sprintf("%d_%d_invite_to_join", curUserID, fieldID)).Result()
	if err == nil {
		for from_id, td_id := range inviteMap {
			td_idInt, err := strconv.ParseUint(td_id, 10, 64)
			if err == nil {
				dbClient.TimeDew.DeleteOneID(td_idInt).Exec(ctx)
			}
			from_idInt, err := strconv.ParseUint(from_id, 10, 64)
			if err == nil {
				_, err := rdb.HMSet(ctx, fmt.Sprintf("%d_%d_invite_map", from_idInt, field.ID), currUser.ID, false).Result()
				if err != nil {
					log.Println("clear invite map fail", err.Error())
				}
			}
		}
	}

	return field, currUser

}

func LeaveCSFieldByWS(userID string, cs_field_id string) {

	ctx := context.Background()
	curUserID, _ := strconv.ParseUint(userID, 10, 64)
	fieldID, _ := strconv.ParseUint(cs_field_id, 10, 64)
	if fieldID == 0 || curUserID == 0 {

	} else {

		currUser, err := dbClient.User.Query().Where(user.ID(curUserID)).Only(ctx)
		field, err := dbClient.CSField.Query().Where(csfield.ID(fieldID)).WithJoinedUser().Only(ctx)

		if currUser.PrivateCsFieldID == field.ID {
			genDefaultCSField(currUser)
		}
		if err != nil {
			_, err = field.Update().RemoveJoinedUser(currUser).Save(ctx)
			if err != nil {
				log.Println("LeaveCSFieldByWS Error:", err.Error())
			}
		}
	}

}

func ChatByWS(userID, cs_field_id, msgBody string) {

	ctx := context.Background()
	curUserID, _ := strconv.ParseUint(userID, 10, 64)
	fieldID, _ := strconv.ParseUint(cs_field_id, 10, 64)
	dbClient.CSField.Query().Where(csfield.ID(fieldID)).Only(ctx)
	_, err := dbClient.Message.Create().SetContent(strings.TrimSpace(string(msgBody))).SetOwnerID(curUserID).Save(ctx)
	if err != nil {
		log.Println("ChatByWS Error:", err.Error())
	}

}

func CreateFeedback(ctx iris.Context) {
	curr_user := GetCurrentUserData(ctx)
	feedbackType := strings.TrimSpace(ctx.PostValue("feedback_type"))
	var fb feedback.Type
	item_id := strings.TrimSpace(ctx.PostValue("item_id"))
	item_id_int, _ := strconv.ParseUint(item_id, 10, 64)
	reason := strings.TrimSpace(ctx.PostValue("reason"))
	if feedbackType == "field" {
		fb = feedback.TypeCsField
	} else if feedbackType == "time_dew" {
		fb = feedback.TypeTimeDew
	} else if feedbackType == "user" {
		fb = feedback.TypeUser
	}
	result, err := dbClient.Feedback.Create().SetItemID(item_id_int).SetUserID(curr_user.ID).SetReason(reason).SetType(fb).Save(ctx)
	if err != nil {
		ctx.JSON(iris.Map{"err_msg": err.Error(), "status": iris.StatusBadRequest})
		return
	}
	ctx.JSON(iris.Map{"data": result, "status": iris.StatusOK})

}

func GetFeedbackReason(ctx iris.Context) {
	feedbackType := strings.TrimSpace(ctx.URLParam("feedback_type"))
	log.Println("get feedback reason", feedbackType)
	ctx.JSON(iris.Map{"data": reportReasons, "status": iris.StatusOK})
}

func GetHiddenUsers(ctx iris.Context) {
	currUser := GetCurrentUserData(ctx)
	var item_ids []uint64

	err := dbClient.Hidden.Query().Where(hidden.UserID(currUser.ID)).Select(hidden.FieldHiddenID).Scan(ctx, &item_ids)

	users, err := dbClient.User.Query().Where(user.IDIn(item_ids...)).All(ctx)
	if err != nil {
		ctx.JSON(iris.Map{"err_msg": err.Error(), "status": iris.StatusBadRequest})
		return
	}
	ctx.JSON(iris.Map{"data": users, "status": iris.StatusOK})
}

func SearchHiddenUsers(ctx iris.Context) {
	key := strings.TrimSpace(ctx.URLParam("key"))
	currUser := GetCurrentUserData(ctx)
	var item_ids []uint64

	err := dbClient.Hidden.Query().Where(hidden.UserID(currUser.ID)).Select(hidden.FieldHiddenID).Scan(ctx, &item_ids)

	users, err := dbClient.User.Query().Where(user.IDIn(item_ids...), user.NameContains(key)).All(ctx)
	if err != nil {
		ctx.JSON(iris.Map{"err_msg": err.Error(), "status": iris.StatusBadRequest})
		return
	}
	ctx.JSON(iris.Map{"data": users, "status": iris.StatusOK})
}

func HiddenUser(ctx iris.Context) {
	currUser := GetCurrentUserData(ctx)

	targetID, _ := strconv.ParseUint(strings.TrimSpace(ctx.PostValue("target_id")), 10, 64)

	_, err := dbClient.Hidden.Create().SetUserID(currUser.ID).SetHiddenID(uint64(targetID)).Save(ctx)
	if err != nil {
		ctx.JSON(iris.Map{"err_msg": "hidden failed", "status": iris.StatusBadRequest})
		return
	}
	ctx.JSON(iris.Map{"data": "OK", "status": iris.StatusOK})
}

func RecoverHiddenUser(ctx iris.Context) {
	currUser := GetCurrentUserData(ctx)

	targetID, _ := strconv.ParseUint(strings.TrimSpace(ctx.PostValue("target_id")), 10, 64)

	_, err := dbClient.Hidden.Delete().Where(hidden.UserID(currUser.ID), hidden.HiddenID(targetID)).Exec(ctx)
	if err != nil {
		ctx.JSON(iris.Map{"err_msg": "hidden failed", "status": iris.StatusBadRequest})
		return
	}
	ctx.JSON(iris.Map{"data": "OK", "status": iris.StatusOK})
}

func InviteToFieldList(ctx iris.Context) {
	curr_user := GetCurrentUserData(ctx)
	system_name := strings.TrimSpace(ctx.URLParam("name"))
	var users []*ent.User
	var err error
	var friend_item_ids []uint64
	err = dbClient.Friendship.Query().Where(friendship.UserID(curr_user.ID), friendship.StatusEQ(friendship.StatusEstablished)).Select(friendship.FieldFriendID).Scan(ctx, &friend_item_ids)
	if system_name == "" {
		users, err = dbClient.User.Query().Where(user.IDIn(friend_item_ids...)).All(ctx)
	} else {
		users, err = dbClient.User.Query().Where(user.IDNEQ(curr_user.ID), user.SystemNameContains(system_name)).All(ctx)
		if len(users) == 0 {
			users, err = dbClient.User.Query().Where(user.IDNEQ(curr_user.ID), user.MobileNoContainsFold(system_name)).All(ctx)
		}
		if len(users) == 0 {
			users, err = dbClient.User.Query().Where(user.IDNEQ(curr_user.ID), user.NameContains(system_name)).All(ctx)
		}
	}
	var item_ids []uint64
	err = dbClient.Friendship.Query().Where(friendship.FriendID(curr_user.ID), friendship.StatusEQ(friendship.StatusEstablished)).Select(friendship.FieldUserID).Scan(ctx, &item_ids)
	if err != nil {
		ctx.JSON(iris.Map{"err_msg": err.Error(), "status": iris.StatusBadRequest})
		return
	}
	invite_map, _ := rdb.HGetAll(ctx, fmt.Sprintf("%d_%d_invite_map", curr_user.ID, curr_user.CurrentCsFieldID)).Result()
	var ifresp InviteToFieldListResp
	for _, item := range users {
		if slices.Contains(friend_item_ids, item.ID) {
			var userStatus UserNameStatus
			userStatus.UserID = item.ID
			userStatus.UserName = item.Name
			userStatus.UserStatus = "in " + item.CurrentCsFieldName
			userStatus.Status = "Invite"
			userStatus.CanAction = true
			userStatus.ActionUrl = fmt.Sprintf("action?type=invite&target_id=%d", item.ID)
			if invite_map != nil {
				if invite_map[strconv.FormatUint(item.ID, 10)] == "1" {
					userStatus.Status = "Inviting..."
					userStatus.CanAction = false
					userStatus.ActionUrl = ""
				}
			}
			userStatus.UserThumbnail = item.ThumbnailURL
			ifresp.UserListSuggested = append(ifresp.UserListSuggested, userStatus)
		}

	}
	ctx.JSON(iris.Map{"data": ifresp, "status": iris.StatusOK})
}
