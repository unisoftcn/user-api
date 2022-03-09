package forms

type Base64UploadFile struct {
	FileType string `json:"fileType"`
	FileName string `json:"fileName"`
	Base64   string `json:"base64" valid:"required~请传入base64内容"`
}
