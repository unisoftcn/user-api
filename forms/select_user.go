package forms

type SelectUser struct {
	TokenId int64 `json:"tokenId" valid:"required~请指定登录token"`
	LoginId int64 `json:"loginId" valid:"required~请指定登录用户"`
}
