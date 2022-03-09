package requests

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/vuuvv/errors"
	"github.com/vuuvv/orca/request"
	"github.com/vuuvv/user-api/entity"
)

type settingService struct {
}

var Setting = &settingService{}

func (s *settingService) Get(key string, setting interface{}) (err error) {
	resp := &entity.Setting{}
	_, err = request.Get(
		ServiceUrl("/setting/detail"),
		map[string]string{"key": key},
		resp,
	)

	if err != nil {
		return errors.WithStack(err)
	}
	err = jsoniter.UnmarshalFromString(resp.Value, setting)
	return errors.WithStack(err)
}