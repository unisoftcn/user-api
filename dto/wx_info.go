package dto

import (
	"github.com/vuuvv/user-api/entity"
	"github.com/vuuvv/user-api/model"
)

type WxInfo struct {
	User *entity.ThirdPartyUser `json:"user"`
	Site *model.Site `json:"site"`
}
