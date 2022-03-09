package forms

import (
	"github.com/vuuvv/orca/orm"
)

type WxAccount struct {
	orm.Id
	Name           string           `json:"name" valid:"required~请填写公众号名称"`
	AppId          string           `json:"appId" valid:"required~请填写AppId"`
	AppSecret      string           `json:"appSecret" valid:"required~请填写AppSecret"`
	Token          string           `json:"token"`
	EncodingAESKey string           `json:"encodingAesKey"`
	OrgPath        string           `json:"orgPath"`
	Data           string           `json:"data"`
	Template       map[string]string `json:"template"`
}
