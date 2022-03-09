package forms

type WxArticle struct {
	OrgPath     string `json:"orgPath" valid:"required~请传入orgPath"`
	ToUser      string `json:"toUser" valid:"required~请传入toUser"`
	Title       string `json:"title" valid:"required~请传入title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	H5Url       string `json:"h5Url"`
	PicUrl      string `json:"picUrl"`
}
