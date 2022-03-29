package dto

import (
	"strings"
	"vuuvv.cn/unisoftcn/orca/orm"
	"vuuvv.cn/unisoftcn/user-api/entity"
)

type User struct {
	entity.User
	Org   string         `json:"org"`
	Roles []int64        `json:"roles" gorm:"-"`
	Menus []*entity.Menu `json:"menus" gorm:"-"`
}

type SessionUser struct {
	orm.Id
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	Realname  string `json:"realname"`
	Phone     string `json:"phone"`
	Avatar    string `json:"avatar"`
	IsSuper   bool   `json:"isSuper"`   // IsSuper 是否超级管理员
	OrgId     int64  `json:"orgId"`     // OrgId 默认机构
	OrgPath   string `json:"orgPath"`   // OrgPath 机构编码
	OrgName   string `json:"orgName"`   // OrgName 机构名称
	RoleId    string `json:"roleId"`    // RoleId 角色Id
	RoleValue string `json:"roleValue"` // RoleValue 角色标识
	RoleName  string `json:"roleName"`  // RoleName 角色名称
	Status    string `json:"status"`    // Status 用户状态
	OrgUserId int64  `json:"orgUserId"` // OrgUserId 机构用户Id
	Logins    string `json:"logins"`

	Roles []*SimpleRole  `json:"roles" gorm:"-"`
	Menus []*entity.Menu `json:"menus" gorm:"-"`
}

func (u *SessionUser) IsValid() bool {
	return u.Status == entity.UserStatusOK
}

func (u *SessionUser) GetName() string {
	return u.Username
}

func (u *SessionUser) InAncestorOrg(orgPath string) bool {
	return u.OrgId == 0 || !strings.HasPrefix(orgPath, u.OrgPath)
}

func (u *SessionUser) HasRole(role string) bool {
	for _, r := range u.Roles {
		if r.Value == role {
			return true
		}
	}
	return false
}
