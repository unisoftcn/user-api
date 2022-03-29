package entity

import "vuuvv.cn/unisoftcn/orca/orm"

const (
	// RoleSystemManager 系统管理员
	RoleSystemManager = "system_manager"
	// RoleHeadquarterManager 平台管理员
	RoleHeadquarterManager = "headquarter_manager"
	// RoleHeadquarterMember 普通用户
	RoleHeadquarterMember = "headquarter_member"
	// RolePrimaryFranchiseeManager 一级加盟商管理员
	RolePrimaryFranchiseeManager = "primary_franchisee_manager"
	// RoleSecondaryFranchiseeManager 二级加盟商管理员
	RoleSecondaryFranchiseeManager = "secondary_franchisee_manager"
	// RoleEstateManager 园区管理员
	RoleEstateManager = "estate_manager"
)

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
