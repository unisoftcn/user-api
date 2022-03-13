package requests

import (
	"github.com/vuuvv/errors"
	"github.com/vuuvv/orca/orm"
	"github.com/vuuvv/orca/request"
	"github.com/unisoftcn/user-api/forms"
)

type wxService struct {
}

var Wx = &wxService{}

func (s *wxService) SendTemplate(form *forms.WxTemplate) (err error) {
	ret := &orm.Id{}
	_, err = request.Post(
		ServiceUrl("/wx/template"),
		form,
		ret,
	)
	return errors.WithStack(err)
}