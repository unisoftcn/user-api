package forms

import "github.com/vuuvv/orca/orm"

type ValueTitleForm struct {
	orm.Id
	Value       string `json:"value" valid:"required~请填写值"`
	Title       string `json:"title" valid:"required~请填写标题"`
	Description string `json:"description"`
}
