package forms

import "github.com/vuuvv/orca/orm"

type Qrcode struct {
	orm.Id
	OrgPath string `json:"orgPath"`
	AppId   string `json:"appId"`
	Sn      string `json:"sn" valid:"required~请传入Sn"`
	Type    string `json:"type" valid:"required~请传入Type"`
	Channel string `json:"channel"`
	Data    string `json:"data"`
	Expires int64  `json:"expires" comment:"超时毫秒数"`
}
