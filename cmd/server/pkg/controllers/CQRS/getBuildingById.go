package cqrs

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harshrastogiexe/cmd/server/pkg/core"
	"github.com/harshrastogiexe/cmd/server/pkg/global"
	"github.com/harshrastogiexe/cmd/server/pkg/models"

	simcompanies "github.com/harshrastogiexe/pkg/sim-companies-proxy"

	"github.com/harshrastogiexe/sim-companies/pkg/logger"
	"github.com/harshrastogiexe/sim-companies/pkg/simcompdb"
)

func GetBuildingById(ctx *gin.Context) {
	app, found := global.Get[*core.Application](core.APPLICATION_TOKEN)
	if !found {
		panic(fmt.Errorf("instance not found: TOKEN = %s", core.APPLICATION_TOKEN))
	}

	var (
		simDB      = simcompdb.SimcompaniesDB{Database: app.DB}
		buildingID = ctx.Param("id")
	)

	building, err := simDB.GetBuildingById(buildingID)

	if err != nil {
		code := http.StatusInternalServerError
		ctx.JSON(code, models.NewApiError(code, err.Error()))
		return
	}

	if building != nil {
		ctx.JSON(http.StatusOK, building)
		return
	}

	logger.Log(logger.Info, "fetching building from simcompanies.com")
	buildingApiData, err := simcompanies.Encyclopedia.GetBuilding(buildingID)

	if err != nil {
		var code int
		switch {
		case err.Error() == "item not found":
			code = http.StatusNotFound
		default:
			code = http.StatusInternalServerError
		}
		ctx.JSON(code, models.NewApiError(code, err.Error()))
		return
	}

	buildingModelData := convertBuildingApiToBuildingModel(buildingApiData)
	err = simDB.SaveBuilding(buildingModelData)
	if err != nil {
		logger.Log(logger.Fail, "failed to save building data")
	}
	ctx.JSON(200, buildingModelData)
}
