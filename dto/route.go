package dto

import "vuuvv.cn/unisoftcn/user-api/entity"

type Route struct {
	*entity.Route
	Status string `json:"status"`
}
