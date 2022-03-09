package forms

type WxTemplate struct {
	ToUser      int64                    `json:"toUser" valid:"required~请传入发送对象"` // 必须, 接受者OpenID
	Template    string                    `json:"template" valid:"required~请传入模板名称"`
	URL         string                    `json:"url"`   // 可选, 用户点击后跳转的URL, 该URL必须处于开发者在公众平台网站中设置的域中
	Color       string                    `json:"color"` // 可选, 整个消息的颜色, 可以不设置
	Params      map[string]*TemplateParam `json:"params"`
	MiniProgram *MiniProgram              `json:"miniprogram"` //可选,跳转至小程序地址
}

type TemplateParam struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

type MiniProgram struct {
	AppID    string `json:"appid"`    //所需跳转到的小程序appid（该小程序appid必须与发模板消息的公众号是绑定关联关系）
	PagePath string `json:"pagepath"` //所需跳转到小程序的具体页面路径，支持带参数,（示例index?foo=bar）
}
