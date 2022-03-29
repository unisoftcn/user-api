package forms

import "vuuvv.cn/unisoftcn/orca/orm"

type Menu struct {
	orm.Tree
	Title       string `json:"title" valid:"required~请填写标题"`
	To          string `json:"to"`
	Icon        string `json:"icon"`
	Seq         int    `json:"seq"`
	Description string `json:"description"`
}
