package requests

import (
	"github.com/vuuvv/errors"
	"github.com/vuuvv/orca/request"
	"github.com/vuuvv/user-api/forms"
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