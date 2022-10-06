package test

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"syscall"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/tuanbieber/integration-golang/config"
	"github.com/tuanbieber/integration-golang/internal/model"
	"github.com/tuanbieber/integration-golang/server"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/lib/pq"
)

type integrationTestSuite struct {
	suite.Suite
	dbConnectionStr string
	port            int
	dbConn          *gorm.DB
	dbMigration     *migrate.Migrate
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, &integrationTestSuite{})
}

func (s *integrationTestSuite) SetupSuite() {
	s.Require().NoError(config.Load())

	s.port = config.Port()
	s.dbConnectionStr = config.DBConnectionURL()
	dbConn, err := gorm.Open(mysql.Open(s.dbConnectionStr), &gorm.Config{})

	s.Require().NoError(err)
	s.dbConn = dbConn

	m, err := migrate.New(sourceURL, driverPrefix+s.dbConnectionStr)
	if err != nil {
		log.Fatal("cannot set up migration for testing: ", err)
	}
	s.dbMigration = m

	serverReady := make(chan bool)

	htppServer := server.Server{
		Port:        config.Port(),
		DBConn:      s.dbConn,
		ServerReady: serverReady,
	}

	go htppServer.Start()
	<-serverReady
}

func (s *integrationTestSuite) TearDownSuite() {
	p, _ := os.FindProcess(syscall.Getpid())
	p.Signal(syscall.SIGINT)
}

func (s *integrationTestSuite) SetupTest() {
	if err := s.dbMigration.Up(); err != nil {
		if err == migrate.ErrNoChange {
			return
		}

		panic(err)
	}
}

func (s *integrationTestSuite) TearDownTest() {
	if err := s.dbMigration.Down(); err != nil {
		if err == migrate.ErrNoChange {
			return
		}

		panic(err)
	}
}

func (s *integrationTestSuite) Test_EndToEnd_CreateOnePhone() {
	reqStr := `{"phone":"iphone 6s", "brand": "apple"}`
	req, err := http.NewRequest(echo.POST, fmt.Sprintf("http://localhost:%d/phones", s.port), strings.NewReader(reqStr))
	s.NoError(err)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusCreated, response.StatusCode)

	byteBody, err := ioutil.ReadAll(response.Body)
	s.NoError(err)

	s.Equal(`{"status":200,"message":"Success","data":{"id":1}}`, strings.Trim(string(byteBody), "\n"))
	response.Body.Close()
}

func (s *integrationTestSuite) Test_EndToEnd_GetOnePhoneById() {
	phone := model.Phone{
		Name:  "galaxy 8s",
		Brand: "samsung",
	}

	s.NoError(s.dbConn.Create(&phone).Error)

	req, err := http.NewRequest(echo.GET, fmt.Sprintf("http://localhost:%d/phones/1", s.port), nil)
	s.NoError(err)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusOK, response.StatusCode)

	byteBody, err := ioutil.ReadAll(response.Body)
	s.NoError(err)

	s.Equal(`{"status":200,"message":"Success","data":{"id":1,"name":"galaxy 8s","brand":"samsung"}}`, strings.Trim(string(byteBody), "\n"))
	response.Body.Close()
}

func (s *integrationTestSuite) Test_EndToEnd_GetAllPhone() {
	xiaomi := model.Phone{
		Name:  "xiaomi mi 8 se",
		Brand: "xiaomi",
	}

	galaxy := model.Phone{
		Name:  "galaxy a18",
		Brand: "samsung",
	}

	s.NoError(s.dbConn.Create(&xiaomi).Error)
	s.NoError(s.dbConn.Create(&galaxy).Error)

	req, err := http.NewRequest(echo.GET, fmt.Sprintf("http://localhost:%d/phones", s.port), nil)
	s.NoError(err)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusOK, response.StatusCode)

	byteBody, err := ioutil.ReadAll(response.Body)
	s.NoError(err)

	s.Equal(`{"status":200,"message":"Success","data":[{"id":1,"name":"xiaomi mi 8 se","brand":"xiaomi"},{"id":2,"name":"galaxy a18","brand":"samsung"}]}`, strings.Trim(string(byteBody), "\n"))
	response.Body.Close()
}
