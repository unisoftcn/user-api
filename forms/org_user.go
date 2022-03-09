package forms

import (
	"github.com/vuuvv/errors"
	"github.com/vuuvv/orca/orm"

)

type OrgUser struct {
	orm.Id
	Nickname   string   `json:"nickname"`
	Phone      string   `json:"phone"`
	Username   string   `json:"username"`   // 传入用户名时一定要传入密码
	Password   string   `json:"password"`
	RoleIdList []int64  `json:"roleIdList"` // 角色Id，数组
	RoleList   []string `json:"roleList"`   // 角色名称，数组，和RoleId必须传入一个
	OrgId      int64    `json:"orgId"`      // 当前机构不用传入Id，否则需传入机构Id
}

func (form *OrgUser) Check() (err error) {
	if form.Username == "" && form.Nickname == "" && form.Phone == "" {
		return errors.New("请填写用户名或昵称或手机号")
	}
	if form.Nickname == "" {
		if form.Username != "" {
			form.Nickname = form.Username
		} else if form.Phone != "" {
			form.Nickname = form.Phone
		}
	}

	if form.OrgId == 0 {
		return errors.New("请传入用户机构")
	}
	return nil
}
