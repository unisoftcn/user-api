package dto

import "github.com/vuuvv/user-api/entity"

type SelectableMenu struct {
	Menus    []*entity.Menu `json:"menus"`
	Selected []*entity.Menu `json:"selected"`
}
