package requests

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/vuuvv/errors"
	"vuuvv.cn/unisoftcn/orca/request"
	"vuuvv.cn/unisoftcn/user-api/entity"
	"vuuvv.cn/unisoftcn/user-api/model"
	"vuuvv.cn/unisoftcn/user-api/vuser"
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

func (s *settingService) GetSetting() (site *model.Site, err error) {
	site = &model.Site{}
	err = s.Get(vuser.SettingSite, site)
	return site, err
}
