package cqrs

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/harshrastogiexe/cmd/server/pkg/core"
	"github.com/harshrastogiexe/cmd/server/pkg/global"
	simcompanies "github.com/harshrastogiexe/pkg/sim-companies-proxy"
	"github.com/harshrastogiexe/sim-companies/pkg/logger"
	"github.com/harshrastogiexe/sim-companies/pkg/simcompdb"
)

func PopulateDatabase(ctx *gin.Context) {
	app, found := global.Get[*core.Application](core.APPLICATION_TOKEN)
	if !found {
		panic(fmt.Errorf("instance not found: TOKEN = %s", core.APPLICATION_TOKEN))
	}

	var passed, failed uint

	for i := '1'; i <= 'z'; i++ {
		building, err := simcompanies.Encyclopedia.GetBuilding(string(i))

		if err != nil {
			failed++
			logger.Log(logger.Fail, err.Error())
			continue
		}

		simDB := simcompdb.SimcompaniesDB{Database: app.DB}
		err = simDB.SaveBuilding(convertBuildingApiToBuildingModel(building))

		if err != nil {
			failed++
			logger.Log(logger.Fail, err.Error())
			continue
		}
		passed++

		switch i {
		case '9':
			i = 'A'
		case 'Z':
			i = 'a'
		}
	}
	ctx.String(200, "Result For Saving Buildings\nSaved: %d\nFailed: %d", passed, failed)
}
