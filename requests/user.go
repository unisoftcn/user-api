package requests

import (
	"strconv"
	"vuuvv.cn/unisoftcn/orca/request"
	"vuuvv.cn/unisoftcn/user-api/dto"
	"vuuvv.cn/unisoftcn/user-api/forms"
)

type userRequest struct {
}

var User = &userRequest{}

func (s *userRequest) GetById(id int64) (user *dto.SessionUser, err error) {
	user = &dto.SessionUser{}
	_, err = request.Get(
		ServiceUrl("/user/detail"),
		map[string]string{"id": strconv.FormatInt(id, 10)},
		user,
	)
	return user, err
}

func (s *userRequest) Create(form *forms.OrgUser) (user *dto.SessionUser, err error) {
	user = &dto.SessionUser{}
	_, err = request.Post(
		ServiceUrl("/user/detail"),
		form,
		user,
	)
	return user, err
}
