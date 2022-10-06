package test

import (
	"github.com/tuanbieber/integration-golang/config"
	"log"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

const (
	sourceURL    = "file://../database/migration"
	driverPrefix = "mysql://"
)

type RepositoryTestSuite struct {
	suite.Suite
	gormDB      *gorm.DB
	dbMigration *migrate.Migrate
}

func TestRepositoryTestSuite(t *testing.T) {
	suite.Run(t, &RepositoryTestSuite{})
}

func (p *RepositoryTestSuite) SetupSuite() {
	p.Require().NoError(config.Load())

	connStr := config.DBConnectionURL()
	gormDB, err := gorm.Open(mysql.Open(connStr), &gorm.Config{})

	if err != nil {
		log.Fatal("Can't connect to mysql", err)
	}

	p.gormDB = gormDB

	dbMigration, err := migrate.New(sourceURL, driverPrefix+connStr)

	if err != nil {
		log.Fatal("cannot set up migration for testing: ", err)
	}

	p.dbMigration = dbMigration
}

func (p *RepositoryTestSuite) TearDownSuite() {
	if err := p.dbMigration.Down(); err != nil {
		if err == migrate.ErrNoChange {
			return
		}

		panic(err)
	}
}

func (p *RepositoryTestSuite) SetupTest() {
	if err := p.dbMigration.Up(); err != nil {
		if err == migrate.ErrNoChange {
			return
		}

		panic(err)
	}
}

func (p *RepositoryTestSuite) TearDownTest() {
	if err := p.dbMigration.Down(); err != nil {
		if err == migrate.ErrNoChange {
			return
		}

		panic(err)
	}
}
