package vuser

const (
	//SessionKeyTrdUserId     = "/vuser/session/key/third_party_user_id"
	//SessionKeyWxRedirectUrl = "/vuser/session/key/redirect_url"
	//SessionKeyWxLoginToken  = "/vuser/session/key/wx_login_token"
	//SessionKeyCaptcha       = "/vuser/session/key/captcha"
	//SessionKeyLogin         = "/vuser/session/key/login"
)

type WechatLoginInfo struct {
	TrdUserId int64  `json:"trdUserId"`
	Nickname  string `json:"nickname"`
	AppId     string `json:"appId"`
	OrgPath   string `json:"orgPath"`
}

//func CheckWechatLogin(ctx *gin.Context) (ret *WechatLoginInfo, err error) {
//	ret = &WechatLoginInfo{}
//	err = sessions.GetSession(ctx, SessionKeyTrdUserId, &ret)
//	if utils.RecordNotFound(err) {
//		return nil, errors.New("请从微信客户端登录")
//	}
//	if err != nil {
//		return nil, errors.WithStack(err)
//	}
//	return ret, nil
//}
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
