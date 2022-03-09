package entity

import (
	"github.com/vuuvv/orca/orm"
)

type Org struct {
	orm.Tree
	Name        string `json:"name" gorm:"not null;comment:名称"`
	Sn          string `json:"sn" gorm:"index:t_org_idx_sn;comment:序号"`
	TypeId      int64  `json:"typeId" gorm:"not null;index:t_org_idx_type_id;comment:机构类型"`
	Description string `json:"description" gorm:"text;comment:机构描述"`
	Address     string `json:"address" gorm:"text;comment:机构地址"`
	Logo        string `json:"logo" gorm:"comment:机构商标"`
	Data        string `json:"data" gorm:"comment:机构数据,json格式"`
	WechatId    int64  `json:"wechatId" gorm:"default:0;index:t_org_idx_wechat_id;comment:机构微信关联Id"`
	Status      string `json:"status" gorm:"comment:机构状态"`
	Seq         int    `json:"seq" gorm:"comment:排序号"`
	AreaCode    string `json:"areaCode" gorm:"index:t_org_idx_area_code;comment:行政区划代码"`
	AreaName    string `json:"areaName" gorm:"comment:行政区划名称"`
	orm.Entity

	Type *OrgType `json:"type" gorm:"-"`
}

func (*Org) TableName() string {
	return "t_org"
}

func (*Org) TableTitle() string {
	return "机构"
}

func IsOrgType(org *Org, typ string) bool {
	if org == nil {
		return false
	}
	return org.Type.Value == typ
}
