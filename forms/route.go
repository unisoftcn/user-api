package forms

import "github.com/vuuvv/orca/orm"

type Route struct {
	orm.Id
	Group        string `json:"group" valid:"required~请填写组名"`               // 组名
	Name         string `json:"name" valid:"required~请填写名称"`                // 名称
	Method       string `json:"method" valid:"required~请填写method"`          // http method
	Path         string `json:"path" valid:"required~请填写原路径"`               // 被代理的路径
	ProxyService string `json:"proxyService" valid:"required~请填写代理service"` // 代理指向的service，从注册服务中心获取对应地址
	ProxyPath    string `json:"proxyPath" valid:"required~请填写代理路径"`         // 代理指向的路径，与ProxyService组合成最后请求的地址，如果为空则去原Path
	Guard        string `json:"guard" valid:"required~请选择访问所需权限"`           // 访问需要的权限，
	Seq          int    `json:"seq"`                                        // 排序号，
}
