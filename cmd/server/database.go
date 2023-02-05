package main

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func SetupSqlServer(dsn string) (db *gorm.DB) {
	dialect := sqlserver.Open(dsn)
	db, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
