package controllers

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	cqrs "github.com/harshrastogiexe/cmd/server/pkg/controllers/CQRS"
)

func Router() http.Handler {
	r := gin.Default()
	r.Use(cors.Default())

	encyclopediaRouter := r.Group("/encyclopedia")
	{
		encyclopediaRouter.GET("/resource/:id", cqrs.GetResourceById)
		encyclopediaRouter.GET("/building/:id", cqrs.GetBuildingById)
		encyclopediaRouter.GET("/building", cqrs.GetAllBuildings)
		encyclopediaRouter.GET("/resource", cqrs.GetAllResource)
	}

	marketRouter := r.Group("/market")
	{
		marketRouter.GET("/:id", cqrs.GetResourceMarketInfo)
	}
	return r
}
