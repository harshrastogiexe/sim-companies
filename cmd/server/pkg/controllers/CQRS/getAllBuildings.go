package cqrs

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harshrastogiexe/cmd/server/pkg/core"
	"github.com/harshrastogiexe/cmd/server/pkg/global"
	"github.com/harshrastogiexe/cmd/server/pkg/models"
	"github.com/harshrastogiexe/sim-companies/pkg/simcompdb"
)

func GetAllBuildings(ctx *gin.Context) {
	app, found := global.Get[*core.Application](core.APPLICATION_TOKEN)
	if !found {
		panic(fmt.Errorf("instance not found: TOKEN = %s", core.APPLICATION_TOKEN))
	}

	simDB := simcompdb.SimcompaniesDB{Database: app.DB}

	buildings, err := simDB.GetBuildings()
	if err != nil {
		code := http.StatusInternalServerError
		ctx.JSON(code, models.NewApiError(code, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, buildings)
}
