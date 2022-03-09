package entity

import "github.com/vuuvv/orca/orm"

type ThirdPartyAccount struct {
	orm.Id
	Name          string `json:"name" gorm:"comment:账户名称"`
	Seq           int    `json:"seq" gorm:"default:100;comment:排序号"`
	Default       bool   `json:"default" gorm:"default:0;comment:是否系统默认账号"`
	OrgPath       string `json:"orgPath" gorm:"index:t_third_party_account_idx_org_code;comment:所属机构代码"`
	Type          string `json:"type" gorm:"index:t_third_party_account_idx_org_type;comment:第三方平台类型与third_party表的value关联"`
	AppId         string `json:"appId" gorm:"index:t_third_party_account_idx_app_id,unique;comment:第三方平台Id"`
	AppSecret     string `json:"appSecret" gorm:"comment:密钥"`
	OauthAddress  string `json:"oauthAddress" gorm:"comment:授权地址"`
	ClientAddress string `json:"clientAddress" gorm:"comment:客户端地址"`
	Data          string `json:"data" gorm:"comment:第三方平台其他数据"`
	orm.Entity
}

func (*ThirdPartyAccount) TableName() string {
	return "t_third_party_account"
}

func (*ThirdPartyAccount) TableTitle() string {
	return "第三方用户账号"
}
