package dto

import "github.com/vuuvv/user-api/entity"

type ThirdPartyAccount struct {
	entity.ThirdPartyAccount
	TypeName string `json:"typeName"`
	OrgName  string `json:"orgName"`
}
