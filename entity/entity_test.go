package entity

import (
	"fmt"
	"github.com/vuuvv/errors"
	"github.com/vuuvv/orca/config"
	"github.com/vuuvv/orca/id"
	"github.com/vuuvv/orca/orm"
	"github.com/vuuvv/orca/serialize"
	"gorm.io/gorm"
	"io/ioutil"
	"testing"
)

func getDb(t *testing.T) *gorm.DB {
	loader := config.MustLoad("../../user/resources/application.yaml")
	ormConfig := &orm.Config{}
	err := loader.Unmarshal(ormConfig, "database")
	_, _ = id.NewGenerator()
	if err != nil {
		t.Error(err)
	}
	db, err := orm.New(ormConfig)
	if err != nil {
		t.Error(err)
	}
	return db
}

func TestMigration(t *testing.T) {
	db := getDb(t)

	err := db.AutoMigrate(
		&orm.Sequence{},
		&Area{},
		&User{},
		&UserLogin{},
		&Route{},
		&Menu{},
		&Org{},
		&OrgType{},
		&OrgUser{},
		&OrgUserRole{},
		&Role{},
		&RoleMenu{},
		&RoleRoute{},
		&OrgTypeMenu{},
		&OrgTypeRoute{},
		&Setting{},
		&Site{},
		&ThirdPartyType{},
		&ThirdPartyAccount{},
		&ThirdPartyUser{},
		&ThirdPartyUserRelation{},
		&Qrcode{},
		&ActionToken{},
	)
	if err != nil {
		t.Error(err)
	}
}

type areaForm struct {
	Name         string `json:"name"`
	Code         string `json:"code"`
	StreetCode   string `json:"streetCode"`
	AreaCode     string `json:"areaCode"`
	CityCode     string `json:"cityCode"`
	ProvinceCode string `json:"provinceCode"`
}

func TestCreateArea(t *testing.T) {
	db := getDb(t)
	t.Log("province...")
	err := createAreas(db, "provinces", "province", func(form *areaForm) string {
		return ""
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("city...")
	err = createAreas(db, "cities", "city", func(form *areaForm) string {
		return form.ProvinceCode
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("county...")
	err = createAreas(db, "areas", "county", func(form *areaForm) string {
		return form.CityCode
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("street...")
	err = createAreas(db, "streets", "street", func(form *areaForm) string {
		return form.AreaCode
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("village...")
	err = createAreas(db, "villages", "village", func(form *areaForm) string {
		return form.StreetCode
	})
	if err != nil {
		t.Error(err)
		return
	}
}

func createAreas(db *gorm.DB, filename string, typ string, parentCode func(form *areaForm) string) error {
	body, err := ioutil.ReadFile(fmt.Sprintf("E:/repo/area/Administrative-divisions-of-China/dist/%s.json", filename))
	if err != nil {
		return errors.WithStack(err)
	}
	var areaList []*areaForm
	err = serialize.JsonParseBytes(body, &areaList)
	if err != nil {
		return errors.WithStack(err)
	}

	var entityList []*Area
	for _, a := range areaList {
		entity := &Area{
			Code: a.Code,
			Name: a.Name,
			Type: typ,
		}
		pCode := parentCode(a)
		if pCode != "" {
			parent := &Area{}
			err = db.First(parent, "code=?", pCode).Error
			if err != nil {
				return errors.WithStack(err)
			}
			entity.ParentId = parent.GetId()
			entity.ParentCode = parent.Code
			entity.FullName = parent.FullName + "-" + entity.Name
		} else {
			entity.FullName = entity.Name
		}
		entityList = append(entityList, entity)
	}
	err = db.CreateInBatches(entityList, 1000).Error
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func TestJoin(t *testing.T) {
	db := getDb(t)

	user := &User{}
	//s := db.Debug().Statement
	//s.AddClause(clause.From{
	//	Tables:
	//})
	err := db.
		Select("u.id, u.nickname").
		Table("t_user as u").
		Joins("join t_user_login as ul on ul.user_id=u.id").First(&user).Error

	fmt.Println(err)
	//db.Model(&user).Joins("t_user").Scan(&user)
	//db.Joins("Account").Find(&user)
}

func TestPage(t *testing.T) {
	paginator := &orm.Paginator{
		Page:     1,
		PageSize: 20,
		Filters: map[string]interface{}{
			"group": "管理",
			"name":  "创建",
		},
	}
	var users []User
	page, err := orm.
		NewPage().
		Count(`select count(*)`).
		Select(`select r.*`).
		Join(`
from t_route as r
where ${group:-1=1} and ${name:-1=1} and ${path:-1=1}
`).
		OrderBy("r", "id", false).
		Criteria("group", "r", "group", orm.OP_Contain).
		Criteria("name", "r", "name", orm.OP_StartsWith).
		Criteria("path", "r", "path", "like").
		Query(getDb(t), paginator, &users)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(page.Total)
}
