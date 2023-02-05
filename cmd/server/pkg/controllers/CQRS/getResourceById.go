package cqrs

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	simcompanies "github.com/harshrastogiexe/pkg/sim-companies-proxy"
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
	var (
		simDB      = simcompdb.SimcompaniesDB{Database: app.DB}
		resourceId = ctx.Param("id")
	)

	resource, err := simDB.GetResourceById(resourceId)
	if err != nil {
		logger.Log(logger.Fail, err.Error())
	}
	if resource != nil {
		ctx.JSON(http.StatusOK, resource)
		return
	}
	logger.Log(logger.Info, "fetching resource information from simcompanies.com [ID=%s]", resourceId)
	resourceApiData, err := simcompanies.Encyclopedia.GetResource(resourceId)

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

	resourceModel := convertApiResourceToBuildResource(resourceApiData)
	fmt.Printf("%+v", resourceModel)
	err = simDB.SaveResource(resourceModel)
	if err != nil {
		logger.Log(logger.Fail, "failed to save resource to database")
		ctx.JSON(200, resourceApiData)
		return
	}
	logger.Log(logger.Info, "resource saved to database")
	ctx.JSON(200, resourceApiData)
}
