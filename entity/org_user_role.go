package entity

import "vuuvv.cn/unisoftcn/orca/orm"

type OrgUserRole struct {
	orm.Id
	OrgUserId int64 `json:"orgUserId" gorm:"index:t_org_user_role_idx_org_user_id;机构用户关联Id"`
	RoleId    int64 `json:"roleId" gorm:"index:t_org_user_role_idx_role_id;角色关联Id"`
	orm.Entity
}

func (*OrgUserRole) TableName() string {
	return "t_org_user_role"
}

func (*OrgUserRole) TableTitle() string {
	return "机构用户角色"
}
