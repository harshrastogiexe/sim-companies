package controllers

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	cqrs "github.com/harshrastogiexe/cmd/server/pkg/controllers/CQRS"
)

func Router() http.Handler {
	router := gin.Default()
	router.Use(cors.Default())

	encyclopediaRouter := router.Group("/encyclopedia")
	{
		encyclopediaRouter.GET("/resource/:id", cqrs.GetResourceById)
		encyclopediaRouter.GET("/building/:id", cqrs.GetBuildingById)
		encyclopediaRouter.GET("/building", cqrs.GetAllBuildings)
	}
	return router
}
