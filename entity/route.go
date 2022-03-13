package entity

import (
	"fmt"
	"github.com/vuuvv/orca/orm"
)

type Route struct {
	orm.Id
	Group      string `json:"group" gorm:"comment:组名"`
	Name       string `json:"name" gorm:"comment:名称"`
	Type       int    `json:"type" gorm:"default:1;comment:类型,固定为1,反代模式"`
	Schema     string `json:"schema" gorm:"comment:http 或 https"`
	Host       string `json:"host" gorm:"comment:入站主机名称"`
	Method     string `json:"method" gorm:"comment:http method"`
	Path       string `json:"path" gorm:"comment:路径"`
	Upstream   string `json:"upstream" gorm:"comment:上游主机"`
	UrlRewrite string `json:"urlRewrite" gorm:"comment:url重写，不填则path不重写完整传入"`
	Timeout    int    `json:"timeout" gorm:"comment:超时时间"`
	Guard      string `json:"guard" gorm:"comment:名称"`
	Seq        int    `json:"seq" gorm:"comment:显示顺序"`
	orm.Entity
}

func (*Route) TableName() string {
	return "t_route"
}

func (*Route) TableTitle() string {
	return "路由"
}

func (this *Route) Key() string {
	return fmt.Sprintf("%s:%s:%s:%s", this.Schema, this.Host, this.Method, this.Path)
}
