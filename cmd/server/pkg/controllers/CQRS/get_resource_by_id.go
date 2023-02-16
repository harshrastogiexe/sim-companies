package cqrs

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/harshrastogiexe/sim-companies/pkg/logger"
	"github.com/harshrastogiexe/sim-companies/pkg/simcompdb"

	"github.com/harshrastogiexe/cmd/server/pkg/core"
	"github.com/harshrastogiexe/cmd/server/pkg/global"
	"github.com/harshrastogiexe/cmd/server/pkg/models"
)

func GetResourceById(ctx *gin.Context) {
	app, found := global.Get[*core.Application](core.APPLICATION_TOKEN)
	if !found {
		panic(fmt.Errorf("instance not found: TOKEN = %s", core.APPLICATION_TOKEN))
	}

	repo := simcompdb.NewRepository(app.DB)
	include := []string{"ResourceBase", "SoldAt", "SoldAtRestaurant", "ProducedAt", "NeededFor", "ImprovesQualityOf"}

	r, err := repo.GetResource(ctx.Param("id"), include...)

	if err != nil {
		logger.Log(logger.Fail, err.Error())
		ctx.JSON(http.StatusInternalServerError, models.NewApiError(http.StatusInternalServerError, err.Error()))
		return
	}
	if r == nil {
		ctx.JSON(http.StatusNotFound, models.NewApiError(http.StatusNotFound, "resource not found"))
		return
	}
	ctx.JSON(http.StatusOK, models.ConvertResourceMain(r))
}
