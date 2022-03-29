package dto

import (
	"vuuvv.cn/unisoftcn/orca/orm"
	"vuuvv.cn/unisoftcn/user-api/entity"
)

type Role struct {
	entity.Role
	OrgTypeName string `json:"orgTypeName"`
	OrgName     string `json:"orgName"`
	Selected    bool   `json:"selected"`
}

type RoleRoutes struct {
	RoleId int64 `json:"roleId"`
	Routes []*string
}

type SimpleRole struct {
	orm.Id
	Value string `json:"value"`
	Title string `json:"title"`
}
