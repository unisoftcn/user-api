package forms

import "github.com/vuuvv/orca/orm"

type Route struct {
	orm.Id
	Group      string `json:"group" valid:"required~请填写组名"`
	Name       string `json:"name" valid:"required~请填写名称"`
	Type       int    `json:"type"`
	Schema     string `json:"schema"`
	Host       string `json:"host"`
	Method     string `json:"method" valid:"required~请填写method"`
	Path       string `json:"path" valid:"required~请填写Path"`
	Upstream   string `json:"upstream" valid:"required~请填写Upstream"`
	UrlRewrite string `json:"urlRewrite"`
	Guard      string `json:"guard" valid:"required~请选择访问所需权限"`
	Seq        int    `json:"seq"`
}
