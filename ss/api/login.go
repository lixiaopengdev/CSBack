package api

import (
	"CSBackendTmp/ent"
	"CSBackendTmp/ent/csfield"
	"CSBackendTmp/ent/friendship"
	"CSBackendTmp/ent/invite_code"
	"CSBackendTmp/ent/user"
	"CSBackendTmp/ent/user_auth"
	"CSBackendTmp/utils"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/kataras/iris/v12"
)

func Login(ctx iris.Context) {

	name := strings.TrimSpace(ctx.PostValue("name"))
	currUser, err := dbClient.User.Query().Where(user.Name(name)).First(ctx)
	if err != nil {
		log.Println(err.Error())
	}

	if currUser == nil {
		currUser, err = dbClient.User.Create().SetName(name).SetSystemName(name).Save(ctx)
		err = initUser(currUser)
		currUser, err = dbClient.User.Query().Where(user.ID(currUser.ID)).Only(ctx)
		log.Println(err)
	}
	sess := Sess.Start(ctx)
	user_json, _ := json.Marshal(currUser)

	sess.SetImmutable("user_id", user.ID)
	sessID := sess.ID()
	ctx.SetCookieKV("_session_id", sessID, iris.CookieSameSite(http.SameSiteNoneMode))
	rdb.Set(ctx, sessID, user_json, 0).Result()
	ctx.JSON(iris.Map{"data": currUser, "status": iris.StatusOK})
}

func LoginByID(ctx iris.Context) {

	id, err := ctx.Params().GetUint64("id")

	currUser, err := dbClient.User.Query().Where(user.ID(id)).Only(ctx)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(iris.Map{"data": currUser, "status": iris.StatusBadRequest})
	}

	sess := Sess.Start(ctx)
	user_json, _ := json.Marshal(currUser)

	sess.SetImmutable("user_id", user.ID)
	sessID := sess.ID()
	ctx.SetCookieKV("_session_id", sessID, iris.CookieSameSite(http.SameSiteNoneMode))
	rdb.Set(ctx, sessID, user_json, 0).Result()
	ctx.JSON(iris.Map{"data": currUser, "status": iris.StatusOK})
}

func GenInviteCode(ctx iris.Context) {
	curr_user := GetCurrentUserData(ctx)

	code := utils.RandStr(6, "char")
	curr_invite, err := dbClient.Invite_Code.Create().SetCode(code).SetOwnerID(curr_user.ID).Save(ctx)
	if err != nil {
		log.Println("failed to gen mobile verify code:", err)
		ctx.JSON(iris.Map{"err_msg": err.Error(), "status": iris.StatusBadRequest})
		return
	}
	ctx.JSON(iris.Map{"data": curr_invite, "status": iris.StatusOK})
}

func VerifyRegisterInviteCode(ctx iris.Context) {
	curr_user := GetCurrentUserData(ctx)
	code := strings.TrimSpace(ctx.PostValue("invite_code"))
	codeRecord, err := dbClient.Invite_Code.Query().Where(invite_code.CodeEQ(code), invite_code.TypeEQ(invite_code.TypeRegister)).Only(ctx)
	if err != nil {
		log.Println("VerifyRegisterInviteCode failed", err)
		ctx.JSON(iris.Map{"err_msg": err.Error(), "status": iris.StatusBadRequest})
		return
	}
	if codeRecord.Status == invite_code.StatusUsed {
		log.Println("VerifyRegisterInviteCode failed, already used", code)
		ctx.JSON(iris.Map{"err_msg": "invite code already used", "status": iris.StatusBadRequest})
		return
	}

	codeRecord, err = codeRecord.Update().SetStatus(invite_code.StatusUsed).SetConsumerID(curr_user.ID).Save(ctx)
	if err != nil {
		log.Println("VerifyRegisterInviteCode failed", err)
		ctx.JSON(iris.Map{"err_msg": err.Error(), "status": iris.StatusBadRequest})
		return
	}

	curr_user, err = dbClient.User.Query().Where(user.ID(curr_user.ID)).Only(ctx)
	curr_user, err = curr_user.Update().SetIsInvited(true).Save(ctx)
	err = initUser(curr_user)
	curr_user, err = dbClient.User.Query().Where(user.ID(curr_user.ID)).Only(ctx)
	if err != nil {
		log.Println("VerifyRegisterInviteCode failed", err)
		ctx.JSON(iris.Map{"err_msg": err.Error(), "status": iris.StatusBadRequest})
		return
	}

	ctx.SetUser(curr_user)
	sessID := ctx.GetCookie("_session_id")
	user_json, _ := json.Marshal(curr_user)

	rdb.Set(ctx, sessID, user_json, 0).Result()

	ctx.JSON(iris.Map{"data": curr_user, "status": iris.StatusOK})
}

func GenMobileVerifyCode(ctx iris.Context) {
	area_code := strings.TrimSpace(ctx.URLParam("area_code"))
	mobile_no := strings.TrimSpace(ctx.URLParam("mobile_no"))
	code := utils.RandStr(6, "number")
	_, err := rdb.Set(ctx, fmt.Sprintf("%s-%s-mbvc", area_code, mobile_no), code, 5*time.Minute).Result()
	if err != nil {
		log.Println("failed to gen mobile verify code:", err)
		ctx.JSON(iris.Map{"err_msg": err.Error(), "status": iris.StatusBadRequest})
		return
	}
	ctx.JSON(iris.Map{"data": code, "status": iris.StatusOK})
}

func VerifyMobile(ctx iris.Context) {
	area_code := strings.TrimSpace(ctx.PostValue("area_code"))
	mobile_no := strings.TrimSpace(ctx.PostValue("mobile_no"))
	verify_code := strings.TrimSpace(ctx.PostValue("verify_code"))
	code, err := rdb.Get(ctx, fmt.Sprintf("%s-%s-mbvc", area_code, mobile_no)).Result()
	if err != nil {
		log.Println("failed to verify mobile code:", err)
		ctx.JSON(iris.Map{"err_msg": err.Error(), "status": iris.StatusBadRequest})
		return
	}
	if code != verify_code {
		log.Println("failed to verify mobile code:", area_code, mobile_no, verify_code, code)
		ctx.JSON(iris.Map{"err_msg": "wrong verify_code", "status": iris.StatusBadRequest})
		return
	}

	mobile_no_full := area_code + "-" + mobile_no

	userAuth, err := dbClient.User_auth.Query().Where(user_auth.MobileNoEQ(mobile_no_full)).WithOwner().Only(ctx)
	if userAuth != nil && err == nil {
		log.Println("mobile_no already registed", mobile_no, userAuth.Edges.Owner.ID, userAuth.Edges.Owner.SystemName)
		sess := Sess.Start(ctx)
		user_json, _ := json.Marshal(userAuth.Edges.Owner)

		sess.SetImmutable("user_id", userAuth.Edges.Owner.ID)
		sessID := sess.ID()

		_, err := rdb.Set(ctx, sessID, user_json, 0).Result()
		if err != nil {
			ctx.JSON(iris.Map{"err_msg": err.Error(), "status": iris.StatusBadRequest})
			return
		}
		ctx.SetCookieKV("_session_id", sessID)
		if userAuth.Edges.Owner.IsInvited {
			GenSystemTimeDew(userAuth.Edges.Owner.ID, "online")
			ctx.JSON(iris.Map{"data": iris.Map{"authed": true, "status": iris.StatusOK, "user": userAuth.Edges.Owner}, "status": iris.StatusOK})
		} else {
			ctx.JSON(iris.Map{"data": "need invite code", "status": iris.StatusOK})
		}
		return
	} else if strings.Contains(err.Error(), "not found") {
		{
			tx, err := dbClient.Tx(ctx)
			if err != nil {
				log.Println("starting a transaction: %w", err)
			}
			user, err := tx.User.Create().SetRegionCode(area_code).SetMobileNo(mobile_no).Save(ctx)
			if err != nil {
				log.Println("starting a transaction: %w", err)
			}
			_, err = tx.User_auth.Create().SetType(user_auth.TypeMobile).SetMobileNo(mobile_no_full).SetOwner(user).Save(ctx)
			if err != nil {
				log.Println("starting a transaction: %w", err)
			}
			if err != nil {
				tx.Rollback()
			}
			tx.Commit()
			log.Println("new mobile_no register", mobile_no_full)
			sess := Sess.Start(ctx)
			user_json, _ := json.Marshal(user)

			sess.SetImmutable("user_id", user.ID)
			sessID := sess.ID()

			_, err = rdb.Set(ctx, sessID, user_json, 0).Result()
			if err != nil {
				ctx.JSON(iris.Map{"err_msg": err.Error(), "status": iris.StatusBadRequest})
				return
			}
			ctx.SetCookieKV("_session_id", sessID)
			ctx.JSON(iris.Map{"data": "need invite code", "status": iris.StatusOK})
			return
		}
	}

	// ctx.JSON(iris.Map{"data": "ok", "status": iris.StatusOK})
}

func initUser(cur_user *ent.User) error {
	err := genDefaultCSField(cur_user)
	err = GenSystemTimeDew(cur_user.ID, "register")
	return err
}

func genDefaultCSField(cur_user *ent.User) error {
	ctx := context.Background()
	cur_field, err := dbClient.CSField.Create().SetStatus(csfield.StatusOpening).SetType(csfield.TypeEmpty).SetMode(csfield.ModeSingle).
		SetName(fmt.Sprintf("%s's Field", cur_user.Name)).SetUserID(cur_user.ID).AddJoinedUserIDs(cur_user.ID).Save(ctx)
	if err != nil {
		log.Println("failed when create field: ", err)
	}
	prop := map[string]int{"version": 0}
	propInit, _ := json.Marshal(prop)
	_, err = cur_user.Update().SetIsOnline(true).SetCurrentCsFieldID(cur_field.ID).SetCurrentCsFieldName(cur_field.Name).SetPrivateCsFieldID(cur_field.ID).SetPrivateCsFieldName(cur_field.Name).Save(ctx)
	if err != nil {
		log.Println("failed when create field: ", err)
	}
	_, err = rdb.Set(ctx, fmt.Sprintf("csfl_%d", cur_field.ID), propInit, 0).Result()

	if err != nil {
		log.Println("failed when create field: ", err)
	}
	return err
}

func LoginCheckMobile(ctx iris.Context) {
	region_code := strings.TrimSpace(ctx.URLParam("region_code"))
	mobile_no := strings.TrimSpace(ctx.URLParam("mobile_no"))
	if !utils.CheckMobileFormat(mobile_no) {
		ctx.JSON(iris.Map{"err_msg": "Incorrect format", "status": iris.StatusBadRequest})
		return
	}
	mobile_no_auth := region_code + "-" + mobile_no
	exist, err := dbClient.User_auth.Query().Where(user_auth.MobileNoEQ(mobile_no_auth), user_auth.IsFinished(true)).Exist(ctx)
	if err != nil {
		ctx.JSON(iris.Map{"err_msg": err.Error(), "status": iris.StatusBadRequest})
		return
	}
	if exist {
		ctx.JSON(iris.Map{"data": iris.Map{"registed": true}, "status": iris.StatusOK})
		return
	} else {
		ctx.JSON(iris.Map{"data": iris.Map{"registed": false}, "status": iris.StatusOK})
		return
	}
}

func CheckPassword(ctx iris.Context) {
	pw := strings.TrimSpace(ctx.URLParam("pw"))
	log.Println(pw)
	ok, reasons := utils.CheckPasswordFormat(pw)

	if !ok {
		im := make(map[string]interface{})
		for _, reason := range reasons {

			if reason == "6 characters minimum" {
				im["count"] = reason

			}
			if reason == "20 characters maximum" {
				im["count"] = reason
			}
			if reason == "must contain one number" {
				im["digit"] = reason

			}
			if reason == "must contain one uppercase" {
				im["uppercase"] = reason

			}
		}
		ctx.JSON(iris.Map{"data": iris.Map{"detail_error": im, "pw_valid": false}, "status": iris.StatusOK})
		return
	}

	ctx.JSON(iris.Map{"data": iris.Map{"detail_error": iris.Map{}, "pw_valid": true}, "status": iris.StatusOK})
}

func CheckBirthday(ctx iris.Context) {
	birthday := strings.TrimSpace(ctx.URLParam("birthday"))
	log.Println(birthday)
	if !utils.CheckAge(birthday) {
		ctx.JSON(iris.Map{"data": iris.Map{"detail_error": iris.Map{"reason": "You are underage"}, "pw_valid": false}, "status": iris.StatusOK})
		return
	}
	ctx.JSON(iris.Map{"data": iris.Map{"detail_error": iris.Map{}, "pw_valid": true}, "status": iris.StatusOK})
	return
}

func SignIn(ctx iris.Context) {
	region_code := strings.TrimSpace(ctx.PostValue("region_code"))
	mobile_no := strings.TrimSpace(ctx.PostValue("mobile_no"))
	mobile_no = region_code + "-" + mobile_no
	pw := strings.TrimSpace(ctx.PostValue("password"))
	signType := strings.TrimSpace(ctx.PostValue("type"))
	userAuth, err := dbClient.User_auth.Query().Where(user_auth.And(user_auth.MobileNoEQ(mobile_no), user_auth.TypeEQ(user_auth.Type(signType)))).WithOwner().Only(ctx)
	if err != nil {
		log.Println("failed query user", err.Error())
		ctx.JSON(iris.Map{"err_msg": err.Error(), "status": iris.StatusBadRequest})
		return
	}
	if userAuth != nil {
		if utils.ComparePasswords(userAuth.Password, []byte(pw)) {
			sess := Sess.Start(ctx)
			user_json, _ := json.Marshal(userAuth.Edges.Owner)

			sess.SetImmutable("user_id", userAuth.Edges.Owner.ID)
			sessID := sess.ID()

			_, err := rdb.Set(ctx, sessID, user_json, 0).Result()
			if err != nil {
				ctx.JSON(iris.Map{"err_msg": err.Error(), "status": iris.StatusBadRequest})
				return
			}
			ctx.SetCookieKV("_session_id", sessID)
			ctx.JSON(iris.Map{"data": iris.Map{"authed": true, "status": iris.StatusOK, "user": userAuth.Edges.Owner}, "status": iris.StatusOK})
			return
		} else {
			ctx.JSON(iris.Map{"err_msg": "password wrong", "status": iris.StatusBadRequest})
			return
		}

	} else {
		ctx.JSON(iris.Map{"data": iris.Map{"authed": false}, "status": iris.StatusOK})
	}
}

func SignUp(ctx iris.Context) {
	region_code := strings.TrimSpace(ctx.PostValue("region_code"))
	mobile_no := strings.TrimSpace(ctx.PostValue("mobile_no"))
	pw := strings.TrimSpace(ctx.PostValue("password"))
	mobile_no_auth := region_code + "-" + mobile_no
	birthday := strings.TrimSpace(ctx.PostValue("birthday"))

	star := utils.GetConstellation(birthday)
	ip := ctx.RemoteAddr()

	if !utils.CheckMobileFormat(mobile_no) {
		ctx.JSON(iris.Map{"err_msg": "Incorrect format", "status": iris.StatusBadRequest})
		return
	}
	ok, _ := utils.CheckPasswordFormat(pw)
	if !ok {
		ctx.JSON(iris.Map{"err_msg": "Incorrect format", "status": iris.StatusBadRequest})
		return
	}

	if !utils.CheckAge(birthday) {
		ctx.JSON(iris.Map{"err_msg": "Age under 13 is not allowed", "status": iris.StatusBadRequest})
		return
	}

	name := strings.TrimSpace(ctx.PostValue("name"))
	pwhash := utils.HashAndSalt([]byte(pw))
	exist, err := dbClient.User_auth.Query().Where(user_auth.MobileNoEQ(mobile_no_auth)).Exist(ctx)
	if exist {
		log.Println("mobile_no exist")
		ctx.JSON(iris.Map{"err_msg": "mobile_no exist", "status": iris.StatusBadRequest})
		return
	}
	default_thumbnai := fmt.Sprintf("https://ruleless.s3.eu-west-1.amazonaws.com/images/default_thumbnail_%d.png", rand.Intn(8))
	curr_user, err := dbClient.User.Create().SetRegionCode(region_code).SetMobileNo(mobile_no).SetBirthday(birthday).SetName(name).SetRegisterIP(ip).SetConstellation(star).Save(ctx)
	dbClient.User.Update().Where(user.IDEQ(curr_user.ID)).AddFriendIDs(systemUserCSTeamID, systemUserRulyID).SetThumbnailURL(default_thumbnai).Save(ctx)
	dbClient.Friendship.Update().Where(friendship.UserID(curr_user.ID), friendship.FriendID(systemUserCSTeamID)).SetStatus(friendship.StatusEstablished).SetCurrType(friendship.CurrTypeNormal).Save(ctx)
	dbClient.Friendship.Update().Where(friendship.UserID(systemUserCSTeamID), friendship.FriendID(curr_user.ID)).SetStatus(friendship.StatusEstablished).SetCurrType(friendship.CurrTypeNormal).Save(ctx)
	dbClient.Friendship.Update().Where(friendship.UserID(curr_user.ID), friendship.FriendID(systemUserRulyID)).SetStatus(friendship.StatusEstablished).SetCurrType(friendship.CurrTypeNormal).Save(ctx)
	dbClient.Friendship.Update().Where(friendship.UserID(systemUserRulyID), friendship.FriendID(curr_user.ID)).SetStatus(friendship.StatusEstablished).SetCurrType(friendship.CurrTypeNormal).Save(ctx)
	_, err = dbClient.User_auth.Create().SetMobileNo(mobile_no_auth).SetType(user_auth.TypeMobile).SetPassword(pwhash).SetIsFinished(true).SetOwner(curr_user).OnConflict().UpdateNewValues().ID(ctx)
	err = initUser(curr_user)
	curr_user, err = dbClient.User.Query().Where(user.IDEQ(curr_user.ID)).Only(ctx)
	if err != nil {
		log.Println("failed creating user", err)
		ctx.JSON(iris.Map{"err_msg": err, "status": iris.StatusBadRequest})
		return
	}
	sess := Sess.Start(ctx)
	user_json, _ := json.Marshal(curr_user)

	sess.SetImmutable("user_id", curr_user.ID)
	sessID := sess.ID()

	_, err = rdb.Set(ctx, sessID, user_json, 0).Result()
	if err != nil {
		ctx.JSON(iris.Map{"err_msg": err.Error(), "status": iris.StatusBadRequest})
		return
	}
	ctx.SetCookieKV("_session_id", sessID)
	ctx.JSON(iris.Map{"data": curr_user, "status": iris.StatusOK})
	return

}

func SignOut(ctx iris.Context) {
	Sess.Start(ctx).Destroy()
	ctx.RemoveCookie("_session_id")
	ctx.JSON(iris.Map{"data": "ok", "status": iris.StatusOK})
}

func LoginByOauth(accessToken, name, sourceID, source string) (*ent.User, error) {
	ctx := context.Background()
	auth, err := dbClient.User_auth.Query().Where(user_auth.OauthID(sourceID), user_auth.OauthSource(source)).WithOwner().Only(ctx)
	if err == nil {
		auth.Update().SetAccessToken(accessToken).Save(ctx)
		return auth.Edges.Owner, err
	} else {
		tx, err := dbClient.Tx(ctx)
		if err != nil {
			log.Println("starting a transaction: %w", err)
		}
		user, err := tx.User.Create().SetName(name).SetSystemName(sourceID + source).Save(ctx)
		_, err = tx.User_auth.Create().SetType(user_auth.TypeOauth2).SetAccessToken(accessToken).SetOauthID(sourceID).SetOauthSource(source).SetOwner(user).Save(ctx)
		if err != nil {
			tx.Rollback()
		}
		tx.Commit()
		return user, nil
	}
}

func ChangePWCheck(ctx iris.Context) {
	currUser := GetCurrentUserData(ctx)
	auth, err := dbClient.User_auth.Query().Where(user_auth.UserID(currUser.ID)).Only(ctx)
	if err != nil {
		ctx.JSON(iris.Map{"err_msg": err.Error(), "status": iris.StatusBadRequest})
		return
	}
	pw := strings.TrimSpace(ctx.URLParam("cur_pw"))
	if utils.ComparePasswords(auth.Password, []byte(pw)) {
		ctx.JSON(iris.Map{"data": "ok", "status": iris.StatusOK})
	} else {
		ctx.JSON(iris.Map{"err_msg": "password wrong", "status": iris.StatusBadRequest})
	}
}

func ChangePW(ctx iris.Context) {
	currUser := GetCurrentUserData(ctx)
	auth, err := dbClient.User_auth.Query().Where(user_auth.UserID(currUser.ID)).Only(ctx)
	if err != nil {
		ctx.JSON(iris.Map{"err_msg": err.Error(), "status": iris.StatusBadRequest})
		return
	}
	pw := strings.TrimSpace(strings.TrimSpace(ctx.PostValue("cur_pw")))
	newPW := strings.TrimSpace(strings.TrimSpace(ctx.PostValue("new_pw")))

	ok, reasons := utils.CheckPasswordFormat(newPW)

	if !ok {
		im := make(map[string]interface{})
		for _, reason := range reasons {

			if reason == "6 characters minimum" {
				im["count"] = reason

			}
			if reason == "20 characters maximum" {
				im["count"] = reason
			}
			if reason == "must contain one number" {
				im["digit"] = reason

			}
			if reason == "must contain one uppercase" {
				im["uppercase"] = reason

			}
		}
		ctx.JSON(iris.Map{"data": iris.Map{"detail_error": im, "pw_valid": false}, "status": iris.StatusOK})
		return
	}

	if utils.ComparePasswords(auth.Password, []byte(pw)) {
		pwhash := utils.HashAndSalt([]byte(newPW))
		err = auth.Update().SetPassword(pwhash).Exec(ctx)
		if err == nil {
			ctx.JSON(iris.Map{"data": "ok", "status": iris.StatusOK})
			return
		} else {
			ctx.JSON(iris.Map{"err_msg": err.Error(), "status": iris.StatusBadRequest})
		}

	} else {
		ctx.JSON(iris.Map{"err_msg": "password wrong", "status": iris.StatusBadRequest})
	}
}
