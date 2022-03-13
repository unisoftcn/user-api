package forms

type SyncRoute struct {
	Path     string `json:"path"`
	Upstream string `json:"upstream" valid:"required~请传入上游服务"`
	DataUrl  string `json:"dataUrl"`
	OnlyAdd  bool   `json:"onlyAdd"`
}
