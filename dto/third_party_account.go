package dto

import "vuuvv.cn/unisoftcn/user-api/entity"

type ThirdPartyAccount struct {
	entity.ThirdPartyAccount
	TypeName string `json:"typeName"`
	OrgName  string `json:"orgName"`
}
