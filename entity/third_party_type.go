package entity

import "vuuvv.cn/unisoftcn/orca/orm"

// ThirdPartyType 第三方平台
type ThirdPartyType struct {
	orm.Id
	Value       string `json:"value" gorm:"not null;index:t_third_party_type_idx_value,unique;comment:实际名称"`
	Title       string `json:"title" gorm:"not null;comment:显示名称"`
	Description string `json:"description" gorm:"comment:描述"`
	orm.Entity
}

func (*ThirdPartyType) TableName() string {
	return "t_third_party_type"
}

func (*ThirdPartyType) TableTitle() string {
	return "第三方用户平台"
}
