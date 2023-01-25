package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harshrastogiexe/app/controller"
	"github.com/harshrastogiexe/app/database"
	"github.com/harshrastogiexe/app/logger"
	"github.com/harshrastogiexe/cmd/repository"
)

func SetupHandler() http.Handler {
	router := gin.Default()
	setupResourceRoutes(router.Group("/resource"))
	return router
}

func setupResourceRoutes(router *gin.RouterGroup) {
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
}
