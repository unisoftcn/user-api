package dto

import (
	"github.com/unisoftcn/user-api/entity"
	"github.com/unisoftcn/user-api/model"
)

type WxInfo struct {
	User *entity.ThirdPartyUser `json:"user"`
	Site *model.Site `json:"site"`
}
