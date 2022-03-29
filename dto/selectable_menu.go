package dto

import "vuuvv.cn/unisoftcn/user-api/entity"

type SelectableMenu struct {
	Menus    []*entity.Menu `json:"menus"`
	Selected []*entity.Menu `json:"selected"`
}
