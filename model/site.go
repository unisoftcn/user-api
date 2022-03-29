package model

import "vuuvv.cn/unisoftcn/orca/utils"

// Site 每个机构都可以用自己的网站配置或者使用最近的上级机构的网站配置，如果都没有就使用网站默认的配置
// 在setting表中的site中
type Site struct {
	Title         string `json:"title"`         // 网站标题
	Logo          string `json:"logo"`          // Logo
	Protocol      string `json:"protocol"`      // 协议
	Domain        string `json:"domain"`        // 域名
	Path          string `json:"path"`          // 路径
	Dashboard     string `json:"dashboard"`     // 后台管理地址
	Wechat        string `json:"wechat"`        // 微信端地址
	WechatOauth   string `json:"wechatOauth"`   // 微信授权地址
	UploadUrl     string `json:"uploadUrl"`     // 上传文件的访问地址
	Custom        bool   `json:"custom"`        // 是否自定义的
	HeadquarterId int64  `json:"headquarterId"` // 总部机构Id
}

func (s *Site) Url() (string, error) {
	return utils.UrlJoin(s.Protocol+s.Domain, s.Path)
}
