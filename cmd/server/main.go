package main

import (
	"os"
	"os/signal"
	"time"

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

	app := core.NewApplication()

	app.UseGormDatabase(SetupSqlServer(os.Getenv("DSN")))
	app.UseHandler(controllers.Router())

	global.Add(core.APPLICATION_TOKEN, app)

	if err := app.PingDatabase(); err != nil {
		panic(err)
	} else {
		logger.Log(logger.Info, "connected to database")
	}

	if err := simcompdb.Migrate(app.DB); err != nil {
		panic(err)
	}

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

func handleShutDown(app *core.Application) {
	if err := app.ShutDown(); err != nil {
		logger.Log(logger.Fail, "failed to stop server gracefully")
		os.Exit(1)
	}

	logger.Log(logger.Info, "server stopped listening")
}
