package requests

import (
	"github.com/vuuvv/errors"
	"vuuvv.cn/unisoftcn/orca/request"
	"vuuvv.cn/unisoftcn/user-api/forms"
)

type siteService struct {
}

var Site = &siteService{}

func (s *siteService) Upload(form *forms.Base64UploadFile) (url string, err error) {
	_, err = request.Post(
		ServiceUrl("/site/upload"),
		form,
		&url,
	)

	return url, errors.WithStack(err)
}
