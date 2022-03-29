package entity

import "vuuvv.cn/unisoftcn/orca/orm"

type UserLogin struct {
	orm.Id
	UserId int64  `json:"userId"`
	Type   string `json:"type" gorm:"comment:登录类型"`
	Value  string `json:"value" gorm:"comment:值"`
	Data   string `json:"data" gorm:"comment:附加数据"`
	orm.Entity
}

func (*UserLogin) TableName() string {
	return "t_user_login"
}

func (*UserLogin) TableTitle() string {
	return "登录认证信息"
}
