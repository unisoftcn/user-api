package dto

import "github.com/unisoftcn/user-api/entity"

type Route struct {
	*entity.Route
	Status       string `json:"status"`
}
