package forms

import "vuuvv.cn/unisoftcn/user-api/entity"

type SetRoleMenus struct {
	AddList []*entity.RoleMenu
	DelList []*entity.RoleMenu
}
