package forms

type Login struct {
	Username string            `json:"username" valid:"required~请输入用户名"`
	Password string            `json:"password" valid:"required~请输入密码"`
	Captcha  string            `json:"captcha"`
	Type     string            `json:"type"`
	Data     map[string]string `json:"data"`
}
