package requests

import (
	"github.com/vuuvv/orca/request"
	"github.com/vuuvv/user-api/dto"
	"strconv"
)

type userRequest struct {
}

var User = &userRequest{}

func (s *userRequest) GetById(id int64) (org *dto.SessionUser, err error) {
	org = &dto.SessionUser{}
	_, err = request.Get(
		ServiceUrl("/user/detail"),
		map[string]string{"id": strconv.FormatInt(id, 10)},
		org,
	)
	return org, err
}
