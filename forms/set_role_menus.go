package forms

import "github.com/vuuvv/user-api/entity"

type SetRoleMenus struct {
	AddList []*entity.RoleMenu
	DelList []*entity.RoleMenu
}
