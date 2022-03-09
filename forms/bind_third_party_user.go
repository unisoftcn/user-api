package forms

type BindThirdPartyUser struct {
	UserId int64  `json:"userId" valid:"required~请指定用户"`
	OpenId string `json:"openId" valid:"required~请指定openId"`
}
