package entity

import "github.com/vuuvv/orca/orm"

const (
	OrgTypeHeadquarter = "headquarter"
	OrgTypeAgent = "agent"
	OrgTypeSystem = "system"
	OrgTypePark = "park"
)

type OrgType struct {
	orm.Id
	Value       string `json:"value" gorm:"not null;index:t_org_type_idx_value,unique;comment:实际名称"`
	Title       string `json:"title" gorm:"not null;comment:显示名称"`
	Description string `json:"description" gorm:"comment:描述"`
	orm.Entity
}

func (*OrgType) TableName() string {
	return "t_org_type"
}

func (*OrgType) TableTitle() string {
	return "机构类型"
}
