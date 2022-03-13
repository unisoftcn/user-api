package forms

import "github.com/unisoftcn/user-api/entity"

type SetRoleMenus struct {
	AddList []*entity.RoleMenu
	DelList []*entity.RoleMenu
}
