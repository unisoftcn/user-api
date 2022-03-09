package entity

import "github.com/vuuvv/orca/orm"

type User struct {
	orm.Id
	Username  string `json:"username" gorm:"comment:用户登录名"`
	Password  string `json:"-" gorm:"comment:密码"`
	Salt      string `json:"-" gorm:"comment:密码盐"`
	Nickname  string `json:"nickname" gorm:"comment:电子邮箱"`
	Realname  string `json:"realname" gorm:"comment:真实姓名"`
	Phone     string `json:"phone" gorm:"comment:手机号码"`
	Email     string `json:"email" gorm:"comment:电子邮箱"`
	Avatar    string `json:"avatar" gorm:"comment:用户头像"`
	IsSuper   bool   `json:"isSuper"`                                       // IsSuper 是否超级管理员
	OrgId     int64  `json:"orgId"`                                         // OrgId 默认机构
	OrgPath   string `json:"orgPath"`                                       // OrgPath 默认机构Path
	OrgUserId int64  `json:"orgUserId"`                                     // OrgUserId 机构用户Id, 冗余字段
	Status    string `json:"status" gorm:"default:ok;comment:用户状态"` // Status 用户状态
	Logins    string `json:"logins" gorm:"comment:可登录方式"`                   // Logins 可登录方式
	orm.Entity
}

const (
	UserStatusOK string = "ok"
)

func (u *User) TableName() string {
	return "t_user"
}

func (*User) TableTitle() string {
	return "用户"
}

func (u *User) IsValid() bool {
	return u.Status == UserStatusOK
}

func (u *User) GetName() string {
	return u.Username
}
