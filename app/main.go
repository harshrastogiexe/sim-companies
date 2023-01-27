package main

import (
	"net/http"
	"sync"

	"github.com/harshrastogiexe/app/database"
	legacyLogger "github.com/harshrastogiexe/app/logger"
	"github.com/harshrastogiexe/app/routes"
)

func main() {
	SetupDatabaseConnection()
	handler := routes.SetupHandler()

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := http.ListenAndServe(":8080", handler); err != nil {
			legacyLogger.Error.Panic(err)
		}
	}()

	legacyLogger.Info.Println("listening on port :8080")
	wg.Wait()
}

func SetupDatabaseConnection() {
	var (
		wg = sync.WaitGroup{}
	)

	DB, err := database.GetConnection()
	if err != nil {
		legacyLogger.Error.Panic(err)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := database.PingDB(DB); err != nil {
			legacyLogger.Error.Panic(err)
		}
		legacyLogger.Info.Printf("successfully connected to database name=%s", DB.Name())
	}()
	wg.Wait()
}
