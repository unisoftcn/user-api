package entity

import "vuuvv.cn/unisoftcn/orca/orm"

type ThirdPartyUserRelation struct {
	orm.Id
	UserId    int64 `json:"userId" gorm:"comment:用户Id"`
	TrdUserId int64 `json:"trdUserId" gorm:"comment:第三方用户Id"`
	orm.Entity
}

func (*ThirdPartyUserRelation) TableName() string {
	return "t_third_party_user_relation"
}

func (*ThirdPartyUserRelation) TableTitle() string {
	return "第三方用户关联"
}
