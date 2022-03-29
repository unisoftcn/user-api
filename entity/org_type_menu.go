package entity

import "vuuvv.cn/unisoftcn/orca/orm"

type OrgTypeMenu struct {
	orm.Id
	OrgTypeId int64 `json:"orgTypeId" gorm:"not null;index:t_org_type_menu_idx_org_type_id;comment:机构类型关联Id"`
	MenuId    int64 `json:"menuId" gorm:"not null;index:t_org_type_menu_idx_menu_id;comment:菜单关联Id"`
	orm.Entity
}

func (*OrgTypeMenu) TableName() string {
	return "t_org_type_menu"
}

func (*OrgTypeMenu) TableTitle() string {
	return "机构类型菜单"
}
