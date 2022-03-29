package requests

import (
	"github.com/vuuvv/errors"
	"strconv"
	"vuuvv.cn/unisoftcn/orca/orm"
	"vuuvv.cn/unisoftcn/orca/request"
	"vuuvv.cn/unisoftcn/user-api/entity"
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

func (s *thirdPartyUserService) BindUser(id int64) (ret *entity.ThirdPartyUser, err error) {
	ret = &entity.ThirdPartyUser{}
	_, err = request.Post(
		ServiceUrl("/third_party_user/bind"),
		&orm.Id{Id: id},
		ret,
	)
	return ret, errors.WithStack(err)
}
