package entity

import "github.com/vuuvv/orca/orm"

type Role struct {
	orm.Id
	Value       string `json:"value" gorm:"not null;index:t_role_idx_value,unique;comment:实际名称"`
	Title       string `json:"title" gorm:"not null;comment:显示名称"`
	Description string `json:"description" gorm:"comment:描述"`
	OrgTypeId   int64  `json:"orgTypeId" gorm:"default:0;index:t_role_idx_org_type_id;comment:机构类型关联Id"`
	OrgId       int64  `json:"orgId" gorm:"default:0;index:t_role_idx_org_id;comment:机构关联Id"`
	orm.Entity
}

func (*Role) TableName() string {
	return "t_role"
}

func (*Role) TableTitle() string {
	return "角色"
}
