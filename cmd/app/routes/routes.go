package routes

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/harshrastogiexe/cmd/repository"
	"github.com/harshrastogiexe/sim-companies/cmd/app/controller"
	"github.com/harshrastogiexe/sim-companies/cmd/app/database"
	"github.com/harshrastogiexe/sim-companies/cmd/app/logger"
)

func SetupHandler() http.Handler {
	router := gin.Default()
	router.Use(cors.Default())

	setupResourceRoutes(router.Group("/resource"))
	setupBuildingRoutes(router.Group("/building"))

	return router
}

func setupResourceRoutes(router *gin.RouterGroup) {
	logger.Info.Println("setting up resource routes ")
	db, err := database.GetConnection()
	if err != nil {
		logger.Error.Panic(err)
		return
	}
	rController := controller.ResourceController{
		ResourceRepo: repository.DBResourceRepository{DB: db},
	}

	// routes
	router.GET("/", rController.GetAllResource)
	router.GET("/:id", rController.GetResourceByID)

	logger.Info.Println("resource routes setup successfully")
}

func setupBuildingRoutes(router *gin.RouterGroup) {
	logger.Info.Println("setting up building routes ")
	db, err := database.GetConnection()
	if err != nil {
		logger.Error.Println("failed to setup building routes\n%w", err)
		logger.Error.Panic(err)
		return
	}

	bController := controller.BuildingController{
		BuildingsRepository: &repository.SqlServerBuildingRepository{GormDB: db},
	}
	router.GET("/:id", bController.GetBuildingById)

	logger.Info.Println("building routes setup successfully")
}
