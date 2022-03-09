package dto

import "github.com/vuuvv/user-api/entity"

type SelectableRoute struct {
	Routes   []*entity.Route `json:"routes"`
	Selected []*entity.Route `json:"selected"`
}
