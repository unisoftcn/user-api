package entity

import "vuuvv.cn/unisoftcn/orca/orm"

type Site struct {
	orm.Id
	OrgPath  string `json:"orgPath" gorm:"default:;comment:机构代码,0为全局默认"`
	Protocol string `json:"protocol" gorm:"not null;comment:协议,http或https"`
	Domain   string `json:"domain" gorm:"not null;index:t_site_idx_domain,unique;comment:域名"`
	Title    string `json:"title" gorm:"not null;comment:网站标题"`
	orm.Entity
}

func (*Site) TableName() string {
	return "t_site"
}

func (*Site) TableTitle() string {
	return "网站配置"
}
