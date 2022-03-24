package requests

import (
	"github.com/unisoftcn/user-api/dto"
	"github.com/unisoftcn/user-api/entity"
	"github.com/unisoftcn/user-api/forms"
	"github.com/vuuvv/errors"
	"github.com/vuuvv/orca/request"
	"github.com/vuuvv/orca/serialize"
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

func (s *orgRequest) GetById(id int64) (org *entity.Org, err error) {
	org = &entity.Org{}
	_, err = request.Get(
		ServiceUrl("/org/detail"),
		map[string]string{"id": strconv.FormatInt(id, 10)},
		org,
	)
	return org, err
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
