package main

import (
	"github.com/tuanbieber/integration-golang/config"
	"github.com/tuanbieber/integration-golang/server"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/mattes/migrate/source/file"
)

func main() {
	config.Load()

	dbStr := config.DBConnectionURL()

	dbConn, err := gorm.Open(mysql.Open(dbStr), &gorm.Config{})

	if err != nil {
		log.Fatal("Can't connect to mysql database. ", err)
	}

	serverReady := make(chan bool)

	httpServer := server.Server{
		DBConn:      dbConn,
		Port:        config.Port(),
		ServerReady: serverReady,
	}

	httpServer.Start()
}
