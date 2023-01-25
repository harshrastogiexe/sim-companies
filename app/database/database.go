package database

import (
	"fmt"
	"sync"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

const (
	dSN string = "server=localhost; Database=SIMCOMPANIES; User Id=SA; Password=ComplexP@ssword1234"
)

var (
	globalDB *gorm.DB
	m        = sync.Mutex{}
)

func getSqlServerDialect(dsn string) gorm.Dialector {
	return sqlserver.Open(dsn)
}

// setup the connection data base takes input as dialector
func setupConnection(dialect gorm.Dialector) error {
	db, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to initialize db session:\n%w", err)
	}
	globalDB = db
	return nil
}

// Ping the database
func PingDB(gormDB *gorm.DB) error {
	sqlDB, err := gormDB.DB()
	if err != nil {
		return fmt.Errorf("failed to ping database %s:\n%w", gormDB.Name(), err)
	}
	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database %s:\n%w", gormDB.Name(), err)
	}
	return nil
}

func GetConnection() (*gorm.DB, error) {
	m.Lock()
	if globalDB != nil {
		return globalDB, nil
	}

	dialect := getSqlServerDialect(dSN)
	if err := setupConnection(dialect); err != nil {
		return nil, err
	}
	m.Unlock()
	return globalDB, nil
}
