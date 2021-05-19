package model

import (
	"fmt"
	"github.com/dfzhou6/goblog/pkg/config"
	"github.com/dfzhou6/goblog/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	var err error
	var (
		host     = config.GetString("database.mysql.host")
		port     = config.GetString("database.mysql.port")
		database = config.GetString("database.mysql.database")
		username = config.GetString("database.mysql.username")
		password = config.GetString("database.mysql.password")
		charset  = config.GetString("database.mysql.charset")
	)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s",
		username, password, host, port, database, charset, true, "Local")
	gormConfig := mysql.New(mysql.Config{
		DSN: dsn,
	})

	var level gormLogger.LogLevel
	if config.GetBool("app.debug") {
		level = gormLogger.Warn
	} else {
		level = gormLogger.Error
	}

	DB, err = gorm.Open(gormConfig, &gorm.Config{
		Logger: gormLogger.Default.LogMode(level),
	})

	logger.LogError(err)

	return DB
}
