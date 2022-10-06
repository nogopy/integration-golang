package test

import (
	"github.com/tuanbieber/integration-golang/internal/model"
	"github.com/tuanbieber/integration-golang/internal/repository"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/lib/pq"
	_ "gorm.io/driver/mysql"
)

func (p *RepositoryTestSuite) TestCreateOnePhone() {
	phone := &model.Phone{
		Name:  "iphone 6s",
		Brand: "apple",
	}

	r := repository.NewPhoneRepository(p.gormDB)
	p.Assert().NoError(r.CreateOnePhone(phone))

	id := phone.ID // new id is created

	var result model.Phone

	p.Assert().NoError(p.gormDB.First(&result, model.Phone{ID: id}).Error)
	p.Assert().Equal(phone.Name, result.Name)
	p.Assert().Equal(phone.Brand, result.Brand)
}

func (p *RepositoryTestSuite) TestGetOnePhoneById() {
	phone := &model.Phone{
		Name:  "11 pro max",
		Brand: "apple",
	}

	r := repository.NewPhoneRepository(p.gormDB)
	p.Assert().NoError(r.CreateOnePhone(phone))

	result, err := r.GetOnePhoneById(phone.ID)

	p.Assert().NoError(err)
	p.Assert().Equal(phone.Name, result.Name)
	p.Assert().Equal(phone.Brand, result.Brand)
}

func (p *RepositoryTestSuite) TestGetAllPhone() {
	galaxy := &model.Phone{
		Name:  "s7",
		Brand: "galaxy",
	}

	xiaomi := &model.Phone{
		Name:  "mi 8 se",
		Brand: "xiaomi",
	}

	r := repository.NewPhoneRepository(p.gormDB)
	p.Assert().NoError(r.CreateOnePhone(galaxy))
	p.Assert().NoError(r.CreateOnePhone(xiaomi))

	result, err := r.GetAllPhone()
	p.Assert().NoError(err)
	p.Assert().Len(result, 2)
}
