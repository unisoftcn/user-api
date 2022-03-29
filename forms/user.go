package forms

import "vuuvv.cn/unisoftcn/orca/orm"

type User struct {
	orm.Id
	Username string `json:"username" valid:"required~请输入用户名"`
	Password string `json:"password" valid:"required~请输入密码"`
	Nickname string `json:"nickname"`
}
