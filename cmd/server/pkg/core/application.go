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

// Listen start server to listen and serve, return a channel for error which will emit values when
// server stops
func (app *Application) Listen() <-chan error {
	var errCh = make(chan error)

	go func(server http.Server) {
		errCh <- server.ListenAndServe()
	}(app.server)

	logger.Log(logger.Info, "listening on http://localhost%s/", app.server.Addr)
	return errCh
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
	db, err := app.DB.DB()
	if err != nil {
		return err
	}
	return db.Ping()
}

// Close closed the application instance
func (app *Application) Close() (err error) {
	db, _ := app.DB.DB()
	if err := db.Close(); err != nil {
		return err
	}
	return
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
