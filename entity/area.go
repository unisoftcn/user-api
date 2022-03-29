package entity

import (
	"vuuvv.cn/unisoftcn/orca/orm"
)

type Area struct {
	orm.Id
	ParentId     int64  `json:"parentId" gorm:"index:t_area_idx_parent_id;comment:父Id"`
	Code         string `json:"code" gorm:"index:t_area_idx_code;comment:代码"`
	ParentCode   string `json:"parentCode" gorm:"index:t_area_idx_parent_code;comment:父区域代码"`
	Name         string `json:"name" gorm:"comment:名称"`
	FullName     string `json:"fullName" gorm:"comment:全称"`
	Type         string `json:"type" gorm:"类型: province, city, county, street, village"`
	Abbreviation string `json:"abbreviation" gorm:"comment:简称"`
	PhoneCode    string `json:"phoneCode" gorm:"comment:电话区号"`
}

func (*Area) TableName() string {
	return "t_area"
}

func (*Area) TableTitle() string {
	return "行政区划"
}
