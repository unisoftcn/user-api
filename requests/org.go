package requests

import (
	"github.com/vuuvv/errors"
	"github.com/vuuvv/orca/request"
	"github.com/vuuvv/orca/serialize"
	"github.com/vuuvv/orca/utils"
	"github.com/vuuvv/user-api/dto"
	"github.com/vuuvv/user-api/entity"
	"github.com/vuuvv/user-api/forms"
	"strconv"
)

type orgRequest struct {
}

var Org = &orgRequest{}

func (s *orgRequest) GetByPath(path string) (org *entity.Org, err error) {
	org = &entity.Org{}
	_, err = request.Get(ServiceUrl("/org/detail"), map[string]string{"path": path}, org)
	return
}

func (s *orgRequest) GetByPathC(path string) *entity.Org {
	org, err := s.GetByPath(path)
	utils.PanicIf(err)
	return org
}

func (s *orgRequest) GetById(id int64) (org *entity.Org, err error) {
	org = &entity.Org{}
	_, err = request.Get(
		ServiceUrl("/org/detail"),
		map[string]string{"id": strconv.FormatInt(id, 10)},
		org,
	)
	return org, err
}

func (s *orgRequest) GetByIdC(id int64) *entity.Org {
	org, err := s.GetById(id)
	if err != nil {
		panic(err)
	}
	return org
}

func (s *orgRequest) FindByOrgPaths(orgPath []string) (org []*entity.Org, err error) {
	paths, err := serialize.JsonStringify(orgPath)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	org = []*entity.Org{}
	_, err = request.Get(
		ServiceUrl("/org/list"),
		map[string]string{"paths": paths},
		&org,
	)

	return org, err
}

func (s *orgRequest) FindByIds(ids []int64) (org []*entity.Org, err error) {
	idList, err := serialize.JsonStringify(ids)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	org = []*entity.Org{}
	_, err = request.Get(
		ServiceUrl("/org/list"),
		map[string]string{"ids": idList},
		&org,
	)

	return org, err
}

func (s *orgRequest) Edit(form *forms.Org) (org *entity.Org, err error) {
	org = &entity.Org{}
	_, err = request.Put(ServiceUrl("/org"), form, &org)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return org, nil
}

type DescendantRequest struct {
	Path   string
	Id     int64
	Type   string
	TypeId int64
}

func (s *orgRequest) Descendant(req *DescendantRequest) (ret []*dto.Org, err error) {
	query := map[string]string{}
	if req.Path != "" {
		query["path"] = req.Path
	}
	if req.Id != 0 {
		query["id"] = strconv.FormatInt(req.Id, 10)
	}
	if req.Type != "" {
		query["type"] = req.Type
	}
	if req.TypeId != 0 {
		query["TypeId"] = strconv.FormatInt(req.TypeId, 10)
	}
	_, err = request.Get(
		ServiceUrl("/org/descendant"),
		query,
		&ret,
	)
	return ret, errors.WithStack(err)
}
