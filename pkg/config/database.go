package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var _db *gorm.DB

func InitialDB() {
	configDb := getConfigDb()
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)
	db, err := gorm.Open(sqlserver.Open(configDb+"&charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{Logger: newLogger})
	if err != nil {
		log.Println(err)
		panic(err)
	}
	_db = db
	//Auto Migrate
	// db.Table("User").AutoMigrate(&User{})
}

func getConfigDb() string {
	conf := Database()
	if conf.Port > 0 {
		return fmt.Sprintf(
			"sqlserver://%s:%s@%s:%d?database=%s",
			conf.User,
			conf.Pass,
			conf.Server,
			conf.Port,
			conf.DatabaseName,
		)
	} else {
		return fmt.Sprintf(
			"sqlserver://%s:%s@%s?database=%s",
			conf.User,
			conf.Pass,
			conf.Server,
			conf.DatabaseName,
		)
	}
}

func Db() *gorm.DB {
	return _db
}
