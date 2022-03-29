package entity

import "vuuvv.cn/unisoftcn/orca/orm"

type ThirdPartyUser struct {
	orm.Id
	Subscribe int    `json:"subscribe" gorm:"default:0;comment:是否关注"`
	AppId     string `json:"appId" gorm:"comment:AppId"`
	UserId    int64  `json:"userId" gorm:"comment:关联用户Id"`
	OpenId    string `json:"openId" gorm:"index:t_third_party_user_idx_open_id;comment:第三方平台用户Id"`
	UnionId   string `json:"unionId" gorm:"index:t_third_party_user_idx_union_id;comment:微信的unionId"`
	Nickname  string `json:"nickname" gorm:"comment:第三方平台昵称"`
	Avatar    string `json:"avatar" gorm:"comment:头像"`
	Data      string `json:"data" gorm:"comment:其他数据"`
	orm.Entity

	OrgPath string `json:"orgPath" gorm:"-"`
	OrgId   int64  `json:"orgId" gorm:"-"`
}

func (*ThirdPartyUser) TableName() string {
	return "t_third_party_user"
}

func (*ThirdPartyUser) TableTitle() string {
	return "第三方用户"
}
