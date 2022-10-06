package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/tuanbieber/integration-golang/internal/handler"
	"github.com/tuanbieber/integration-golang/internal/repository"
	"github.com/tuanbieber/integration-golang/internal/service"
	"gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
)

type Server struct {
	Port        int
	DBConn      *gorm.DB
	ServerReady chan bool
}

func (s *Server) Start() {
	appPort := fmt.Sprintf(":%d", s.Port)

	repo := repository.NewPhoneRepository(s.DBConn)
	phoneService := service.NewPhoneService(repo)
	phoneHandler := handler.NewPhoneHandler(phoneService)

	e := echo.New()

	e.POST("/phones", phoneHandler.CreateOnePhone)
	e.GET("/phones", phoneHandler.GetAllPhone)
	e.GET("/phones/:id", phoneHandler.GetOnePhoneById)

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"pong": "ok",
		})
	})

	go func() {
		if err := e.Start(appPort); err != nil {
			logrus.Errorf(err.Error())
			logrus.Infof("shutting down the server")
		}
	}()

	if s.ServerReady != nil {
		s.ServerReady <- true
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		logrus.Fatalf("failed to gracefully shutdown the server: %s", err)
	}
}
