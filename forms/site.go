package forms

import "github.com/vuuvv/orca/orm"

type Site struct {
	orm.Id
	Protocol string `json:"protocol" valid:"required~请选择协议"`
	Domain   string `json:"domain" valid:"required~请填写域名"`
	Title    string `json:"title" valid:"required~请填写标题"`
	Path     string `json:"path"`
	OrgPath  string `json:"orgPath"`
}
