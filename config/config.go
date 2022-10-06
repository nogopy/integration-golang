package config

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

var appConfig config

type config struct {
	appPort                   int
	dbHost                    string
	dbPort                    int
	dbUsername                string
	dbPassword                string
	dbName                    string
	dbMaxPoolSize             int
	dbMaxIdleConn             int
	dbMaxOpenConn             int
	dbMaxConnLifetimeDuration time.Duration
	logLevel                  string
}

func Load() error {
	viper.AutomaticEnv()

	appConfig = config{
		appPort:                   viper.GetInt("PORT"),
		dbHost:                    viper.GetString("DB_HOST"),
		dbPort:                    viper.GetInt("DB_PORT"),
		dbUsername:                viper.GetString("DB_USER"),
		dbPassword:                viper.GetString("DB_PASS"),
		dbName:                    viper.GetString("DB_NAME"),
		dbMaxIdleConn:             viper.GetInt("DB_MAX_IDLE_CONN"),
		dbMaxOpenConn:             viper.GetInt("DB_MAX_OPEN_CONN"),
		dbMaxConnLifetimeDuration: viper.GetDuration("DB_CONN_MAX_LIFETIME"),
		logLevel:                  viper.GetString("LOG_LEVEL"),
	}

	return nil
}

func DBConnectionURL() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", appConfig.dbUsername, appConfig.dbPassword, appConfig.dbHost, appConfig.dbPort, appConfig.dbName)
}

func Port() int {
	return appConfig.appPort
}
