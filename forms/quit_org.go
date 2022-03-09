package forms

type QuitOrg struct {
	UserId int64 `json:"userId" valid:"required~请指定用户"`
	OrgId  int64 `json:"orgId" valid:"required~请指定机构"`
}
