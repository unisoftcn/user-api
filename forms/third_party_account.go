package forms

import "vuuvv.cn/unisoftcn/orca/orm"

type ThirdPartyAccount struct {
	orm.Id
	Name          string `json:"name"`
	Seq           int    `json:"seq"`
	Default       bool   `json:"default"`
	OrgPath       string `json:"orgPath"`
	Type          string `json:"type"`
	AppId         string `json:"appId"`
	AppSecret     string `json:"appSecret"`
	OauthAddress  string `json:"oauthAddress"`
	ClientAddress string `json:"clientAddress"`
	Data          string `json:"data"`
}
