package main

import (
	"net/http"
	"sync"

	"github.com/harshrastogiexe/sim-companies/cmd/app/database"
	"github.com/harshrastogiexe/sim-companies/cmd/app/logger"
	"github.com/harshrastogiexe/sim-companies/cmd/app/routes"
)

func main() {
	SetupDatabaseConnection()
	handler := routes.SetupHandler()

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := http.ListenAndServe(":8080", handler); err != nil {
			logger.Error.Panic(err)
		}
	}()

	logger.Info.Println("listening on port :8080")
	wg.Wait()
}

func SetupDatabaseConnection() {
	var (
		wg = sync.WaitGroup{}
	)

	DB, err := database.GetConnection()
	if err != nil {
		logger.Error.Panic(err)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := database.PingDB(DB); err != nil {
			logger.Error.Panic(err)
		}
		logger.Info.Printf("successfully connected to database name=%s", DB.Name())
	}()
	wg.Wait()
}
