package dto

import (
	"github.com/vuuvv/orca/orm"
	"github.com/unisoftcn/user-api/entity"
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
