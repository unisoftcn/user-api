package entity

import "github.com/vuuvv/orca/orm"

type OrgTypeRoute struct {
	orm.Id
	OrgTypeId int64 `json:"orgTypeId" gorm:"not null;index:t_rog_type_route_idx_org_type_id;comment:角色关联Id"`
	RouteId   int64 `json:"routeId" gorm:"not null;index:t_rog_type_route_idx_router_id;comment:路由关联Id"`
	orm.Entity
}

func (*OrgTypeRoute) TableName() string {
	return "t_org_type_route"
}

func (*OrgTypeRoute) TableTitle() string {
	return "机构类型路由"
}
