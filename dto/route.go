package dto

import "github.com/vuuvv/user-api/entity"

type Route struct {
	*entity.Route
	Status       string `json:"status"`
}
