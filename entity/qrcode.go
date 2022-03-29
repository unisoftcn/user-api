package entity

import (
	"time"
	"vuuvv.cn/unisoftcn/orca/orm"
)

// Qrcode 二维码
// 生成：根据特征码找到二维码，如果是临时二维码要对比expiredTime看二维码是否失效。如果表里有就取出，没有就新生成一个。
// 识别：流程，扫描二维码后，由/wx/msg识别，识别后用id取出二维码，然后发送id到消息队列channel中
type Qrcode struct {
	orm.Id
	Platform    string     `json:"platform" gorm:"comment:所属平台"`
	Type        string     `json:"type" gorm:"comment:二维码类型"`
	Sn          string     `json:"sn" gorm:"index:t_qrcode_idx_sn;comment:二维码识别特征,由它判定二维码是否存在"`
	Tmp         bool       `json:"tmp" gorm:"是否临时二维码，临时二维码必须有ExpiredTime"`
	ExpiredTime *time.Time `json:"expiredTime" gorm:"comment:过期时间"`
	Channel     string     `json:"channel" gorm:"comment:扫描二维码后，发布到消息队列的channel"`
	Data        string     `json:"data" gorm:"comment:扫描二维码后，请求的微服务的数据"`
	Ticket      string     `json:"ticket" gorm:"comment:ticket"`
	Content     string     `json:"content" gorm:"comment:内容"`
	Url         string     `json:"url" gorm:"comment:地址"`
	orm.Entity
}

func (*Qrcode) TableName() string {
	return "t_qrcode"
}

func (*Qrcode) TableTitle() string {
	return "二维码"
}
