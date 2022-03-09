package forms

import "github.com/vuuvv/orca/orm"

type OrgUserEx struct {
	orm.Id
	Nickname string `json:"nickname" valid:"required~请填写用户昵称"`
	RoleId int64 `json:"roleId" valid:"required~请选择用户角色"`
	Phone string `json:"phone"`
}
