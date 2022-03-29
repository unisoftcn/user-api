package dto

import (
	"vuuvv.cn/unisoftcn/user-api/entity"
	"vuuvv.cn/unisoftcn/user-api/model"
)

type WxInfo struct {
	User *entity.ThirdPartyUser `json:"user"`
	Site *model.Site            `json:"site"`
}
