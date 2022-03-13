package dto

import (
	"github.com/unisoftcn/user-api/entity"
	"time"
)

type Org struct {
	entity.Org
	TypeName    string    `json:"typeName"`
	ParentName  string    `json:"parentName"`
	HasChildren bool      `json:"hasChildren"`
	Roles       string    `json:"roles"`
	JoinAt      time.Time `json:"joinAt"`
}
