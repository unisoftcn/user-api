package entity

import "vuuvv.cn/unisoftcn/orca/orm"

type RoleMenu struct {
	orm.Id
	RoleId int64 `json:"roleId" gorm:"not null;index:t_role_menu_idx_role_id;comment:角色关联Id"`
	MenuId int64 `json:"menuId" gorm:"not null;index:t_role_menu_idx_menu_id;comment:菜单关联Id"`
	orm.Entity
}

func (*RoleMenu) TableName() string {
	return "t_role_menu"
}

func (*RoleMenu) TableTitle() string {
	return "角色菜单"
}
