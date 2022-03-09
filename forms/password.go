package forms

import "github.com/vuuvv/orca/orm"

type Password struct {
	orm.Id
	Password string `json:"password" valid:"required~请输入密码"`
}
