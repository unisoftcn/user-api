package forms

import "github.com/vuuvv/orca/orm"

type Role struct {
	orm.Id
	Value       string `json:"value" valid:"required~请填写值"`
	Title       string `json:"title" valid:"required~请填写标题"`
	OrgTypeId   int64  `json:"orgTypeId"`
	OrgId       int64  `json:"orgId"`
	Description string `json:"description"`
}
