package entity

import "vuuvv.cn/unisoftcn/orca/orm"

type OrgUser struct {
	orm.Id
	UserId int64  `json:"userId" gorm:"not null;comment:用户关联Id"`
	OrgId  int64  `json:"orgId" gorm:"not null;comment:机构关联Id"`
	Status string `json:"status" gorm:"default:ok;comment:状态"`
	orm.Entity
}

func (*OrgUser) TableName() string {
	return "t_org_user"
}

func (*OrgUser) TableTitle() string {
	return "机构用户"
}
