package entity

import (
	"vuuvv.cn/unisoftcn/orca/orm"

	"time"
)

const (
	ActionTokenTypeWxLogin   = "wx_login"
	ActionTokenStatusUnused  = "unused"  // 未使用
	ActionTokenStatusScan    = "scan"    // 用户扫描完成，等待下一步动作，例如选择账户
	ActionTokenStatusActive  = "active"  // 用户动作已完成，等待系统验证
	ActionTokenStatusDone    = "done"    // 已经结束，不可再使用
	ActionTokenStatusExpired = "expired" // 已超时
)

type ActionToken struct {
	orm.Id
	Type       string     `json:"type" gorm:"not null;comment:行为类型"`
	LoginId    int64      `json:"userId" gorm:"not null;comment:用户Id"`
	ScanId     int64      `json:"scanId" gorm:"comment:扫描者Id"`
	Status     string     `json:"status" gorm:"not null;comment:状态"`
	Data       string     `json:"data" gorm:"comment:附加数据"`
	Ip         string     `json:"ip" gorm:"comment:执行者Ip"`
	ExpireTime *time.Time `json:"expireTime" gorm:"comment:超时时间"`
	ActionTime *time.Time `json:"actionTime" gorm:"comment:执行时间"`
	DoneTime   *time.Time `json:"doneTime" gorm:"comment:完成时间"`
	orm.Entity
}

func (*ActionToken) TableName() string {
	return "t_action_token"
}

func (*ActionToken) TableTitle() string {
	return "令牌"
}
