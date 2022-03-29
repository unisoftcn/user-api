package requests

import (
	"github.com/vuuvv/errors"
	"vuuvv.cn/unisoftcn/orca/orm"
	"vuuvv.cn/unisoftcn/orca/request"
	"vuuvv.cn/unisoftcn/user-api/forms"
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
