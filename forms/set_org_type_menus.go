package forms

import "vuuvv.cn/unisoftcn/user-api/entity"

type SetOrgTypeMenus struct {
	AddList []*entity.OrgTypeMenu
	DelList []*entity.OrgTypeMenu
}
