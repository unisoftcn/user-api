package forms

import "vuuvv.cn/unisoftcn/orca/orm"

type Org struct {
	orm.Tree
	Name     string      `json:"name" valid:"required~请填写机构名称"`
	Sn       string      `json:"sn"`
	TypeId   int64       `json:"typeId"` // TypeId 和 Type 选一个传入
	Type     string      `json:"type"`
	Address  string      `json:"address"`
	AreaCode string      `json:"areaCode"`
	RawData  interface{} `json:"rawData"`
	Data     string      `json:"-"`
	AreaName string      `json:"-"`
}
