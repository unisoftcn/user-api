package vuser

import (
	"github.com/gin-gonic/gin"
	"github.com/vuuvv/errors"
	"vuuvv.cn/unisoftcn/orca/server"
	"vuuvv.cn/unisoftcn/orca/utils"
)

const (
//SessionKeyTrdUserId     = "/vuser/session/key/third_party_user_id"
//SessionKeyWxRedirectUrl = "/vuser/session/key/redirect_url"
//SessionKeyWxLoginToken  = "/vuser/session/key/wx_login_token"
//SessionKeyCaptcha       = "/vuser/session/key/captcha"
//SessionKeyLogin         = "/vuser/session/key/login"
)

type WechatLoginInfo struct {
	TrdUserId int64  `json:"trdUserId"`
	UserId    int64  `json:"userId"`
	Nickname  string `json:"nickname"`
	OrgId     int64  `json:"orgId"`
	AppId     string `json:"appId"`
	OrgPath   string `json:"orgPath"`
}

func CheckWechatLogin(ctx *gin.Context) (ret *WechatLoginInfo, err error) {
	ret, err = server.GetSessionP[WechatLoginInfo](ctx, CookieTrdUserId)
	//ret = &WechatLoginInfo{}
	//err = sessions.GetSession(ctx, SessionKeyTrdUserId, &ret)
	if utils.RecordNotFound(err) {
		return nil, errors.New("请从微信客户端登录")
	}
	return ret, errors.WithStack(err)
}

//
//func GetLoginUser(ctx *gin.Context) (user *server.AccessToken, err error) {
//	user = &server.AccessToken{}
//	body := ctx.Request.Header.Get(HeadAccessUser)
//	if body == "" {
//		return nil, request.NewErrorUnauthorized()
//	}
//	err = serialize.JsonParse(body, user)
//	if err != nil {
//		return nil, err
//	}
//	if user.OrgId == 0 {
//		return nil, request.NewErrorForbidden()
//	}
//	return user, nil
//}

//
//func CheckLogin(ctx *gin.Context) *dto.SessionUser {
//	user, err := GetLoginUser(ctx)
//	utils.PanicIf(err)
//	return user
//}
//
//func CheckLoginWithOrg(ctx *gin.Context) *dto.SessionUser {
//	user, err := GetLoginUser(ctx)
//	utils.PanicIf(err)
//	return user
//}
