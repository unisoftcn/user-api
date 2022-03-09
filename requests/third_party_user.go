package requests

import (
	"github.com/vuuvv/errors"
	"github.com/vuuvv/orca/request"
	"github.com/vuuvv/user-api/entity"
	"strconv"
)

type thirdPartyUserService struct {
}

var ThirdPartyUser = &thirdPartyUserService{}

func (s *thirdPartyUserService) GetById(id int64) (ret *entity.ThirdPartyUser, err error) {
	ret = &entity.ThirdPartyUser{}
	_, err = request.Get(
		ServiceUrl("/third_party_user/detail"),
		map[string]string{"id": strconv.FormatInt(id, 10)},
		ret,
	)
	return ret, errors.WithStack(err)
}
