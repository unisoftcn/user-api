package dto

import "github.com/unisoftcn/user-api/entity"

type SelectableMenu struct {
	Menus    []*entity.Menu `json:"menus"`
	Selected []*entity.Menu `json:"selected"`
}
