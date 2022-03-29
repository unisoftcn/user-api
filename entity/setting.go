package entity

import "vuuvv.cn/unisoftcn/orca/orm"

type Setting struct {
	orm.Id
	Title       string `json:"title" gorm:"not null;comment:键"`
	Value       string `json:"value" gorm:"comment:值"`
	Description string `json:"description" gorm:"comment:描述"`
	orm.Entity
}

func (*Setting) TableName() string {
	return "t_setting"
}

func (*Setting) TableTitle() string {
	return "系统配置"
}
