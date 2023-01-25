package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harshrastogiexe/app/database"
	"github.com/harshrastogiexe/app/logger"
	"github.com/harshrastogiexe/simcompanies/types"
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
	// repo := repository.DBResourceRepository{DB: db}

	router.GET("/", func(ctx *gin.Context) {
		var resources []types.Resource
		db.Table("Resource").Find(&resources)
		ctx.JSON(http.StatusOK, resources)
	})

}
