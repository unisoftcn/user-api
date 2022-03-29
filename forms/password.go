package forms

import "vuuvv.cn/unisoftcn/orca/orm"

type Password struct {
	orm.Id
	Password string `json:"password" valid:"required~请输入密码"`
}
