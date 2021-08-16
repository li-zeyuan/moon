package config

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// InitDatabase gorm支持连接池
func InitDatabase(conf *Config) *gorm.DB {
	if db != nil {
		return db
	}

	var err error
	db, err = gorm.Open(mysql.Open(conf.DB.DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDb, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDb.SetMaxIdleConns(conf.DB.MaxConn)
	sqlDb.SetMaxOpenConns(conf.DB.MaxOpen)
	sqlDb.SetConnMaxIdleTime(time.Duration(conf.DB.Timeout))
	return db
}
