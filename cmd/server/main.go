package main

import (
	"os"
	"os/signal"
	"time"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"

	"github.com/joho/godotenv"

	"github.com/harshrastogiexe/cmd/server/pkg/controllers"
	"github.com/harshrastogiexe/cmd/server/pkg/core"
	"github.com/harshrastogiexe/cmd/server/pkg/global"

	"github.com/harshrastogiexe/sim-companies/pkg/logger"

	"github.com/harshrastogiexe/sim-companies/pkg/simcompdb"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	start := time.Now()

	var app = SetupApplication()

	global.Add(core.APPLICATION_TOKEN, app)

	ConnectSqlServer(app)
	MigrateDatabase(app.DB)

	app.Listen()
	logger.Log(logger.Info, "time lapsed to start server: %d milliseconds", time.Since(start).Milliseconds())

	// Application Shutdown Logic
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	handleShutDown(app)
	logger.Log(logger.Info, "closing application")
	err := app.Close()
	if err != nil {
		logger.Log(logger.Fail, err.Error())
	}
	os.Exit(0)
}

func SetupSqlServer(dsn string) (db *gorm.DB) {
	dialect := sqlserver.Open(dsn)
	db, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

// create new application instance,
//
// - setup sql server instance using "DSN" from env variable
//
// - setup routers
func SetupApplication() (app *core.Application) {
	app = core.NewApplication()

	app.UseGormDatabase(SetupSqlServer(os.Getenv("DSN")))
	app.UseHandler(controllers.Router())
	return
}

// ConnectSqlServer pings the database, panics if error occurs
func ConnectSqlServer(app *core.Application) {
	if err := app.PingDatabase(); err != nil {
		panic(err)
	} else {
		logger.Log(logger.Info, "connected to database")
	}
}

// MigrateDatabase run the migration scripts
func MigrateDatabase(db *gorm.DB) {
	if err := simcompdb.Migrate(db); err != nil {
		panic(err)
	}
}

// handleShutDown handle gracefully shutdown of server
func handleShutDown(app *core.Application) {
	if err := app.ShutDown(); err != nil {
		logger.Log(logger.Fail, "failed to stop server gracefully")
		os.Exit(1)
	}

	logger.Log(logger.Info, "server stopped listening")
}
