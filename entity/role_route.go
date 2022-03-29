package entity

import "vuuvv.cn/unisoftcn/orca/orm"

type RoleRoute struct {
	orm.Id
	RoleId  int64 `json:"roleId" gorm:"not null;index:t_role_route_idx_role_id;comment:角色关联Id"`
	RouteId int64 `json:"routeId" gorm:"not null;index:t_role_route_idx_route_id;comment:路由关联Id"`
	orm.Entity
}

func (*RoleRoute) TableName() string {
	return "t_role_route"
}

func (*RoleRoute) TableTitle() string {
	return "角色路由"
}
