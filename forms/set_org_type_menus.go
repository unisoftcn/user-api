package forms

import "github.com/unisoftcn/user-api/entity"

type SetOrgTypeMenus struct {
	AddList []*entity.OrgTypeMenu
	DelList []*entity.OrgTypeMenu
}
