package core

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/harshrastogiexe/sim-companies/pkg/logger"
	"gorm.io/gorm"
)

// Application Instance
type Application struct {
	// Application configuration
	Config *configuration

	// Application server instance
	server http.Server

	// Application gorm db instance
	DB *gorm.DB
}

// it build up the application and start listening to port
func (app *Application) Listen() <-chan error {
	var errorChan = make(chan error)

	go func(server http.Server) {
		err := server.ListenAndServe()
		errorChan <- err
	}(app.server)

	logger.Log(logger.Info, "listening on http://localhost%s/", app.server.Addr)
	return errorChan
}

// gracefully stop the server, once stopped can't be started again
func (app *Application) ShutDown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	logger.Log(logger.Info, "attempting to stop server")
	return app.server.Shutdown(ctx)
}

// configure a http handler for application
func (app *Application) UseGormDatabase(db *gorm.DB) {
	app.DB = db
}

// configure a http handler for application
func (app *Application) UseHandler(handler http.Handler) {
	app.server.Handler = handler
}

// Pings the database if, return error nil if something does wrong
func (app *Application) PingDatabase() error {
	sqlDB, err := app.DB.DB()
	if err != nil {
		return err
	}

	return sqlDB.Ping()
}

// Close closed the application instance
func (app *Application) Close() error {
	sqlDB, _ := app.DB.DB()
	var closeErr error = nil

	if err := sqlDB.Close(); err != nil {
		closeErr = err
	}
	return closeErr
}

// create a new application instance with default config
func NewApplication() *Application {
	var app Application
	app.Config = DefaultConfig
	app.server = http.Server{
		Addr: os.Getenv("PORT"),
	}

	return &app
}
