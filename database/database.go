package database

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strings"
)

type ConnectOptions struct {
	Driver   string
	Host     string
	Port     int16
	DBName   string
	Username string
	Password string
	Charset  string
	Loc      string
}

func NewDataBase(options *ConnectOptions, l *logrus.Logger) *gorm.DB {
	db, err := gorm.Open(getDialector(options), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
		Logger: NewMyLogger(logrus.DebugLevel, l),
	})

	if err != nil {
		panic("failed to connect database,err: " + err.Error())
	}

	return db

	//return nil
}

func getDialector(options *ConnectOptions) gorm.Dialector {

	switch strings.ToLower(options.Driver) {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			options.Username, options.Password, options.Host, options.Port, options.DBName)
		return mysql.Open(dsn)
	case "postgres":
		dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			options.Host, options.Port, options.Username, options.Password, options.DBName)
		return postgres.Open(dsn)
	default:
		panic(fmt.Sprintf("unsupported driver: %s", options.Driver))
	}
}
