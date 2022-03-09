package forms

import "github.com/vuuvv/user-api/entity"

type SetOrgTypeMenus struct {
	AddList []*entity.OrgTypeMenu
	DelList []*entity.OrgTypeMenu
}
