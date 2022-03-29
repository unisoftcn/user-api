package entity

import "vuuvv.cn/unisoftcn/orca/orm"

type Menu struct {
	orm.Tree
	Title       string `json:"title" gorm:"not null;comment:标题"`
	To          string `json:"to" gorm:"comment:链接"`
	Icon        string `json:"icon" gorm:"comment:图标"`
	Seq         int    `json:"seq" gorm:"default:0;not null;comment:排序号"`
	Description string `json:"description" gorm:"comment:描述"`
	orm.Entity
}

func (*Menu) TableName() string {
	return "t_menu"
}

func (*Menu) TableTitle() string {
	return "菜单"
}
