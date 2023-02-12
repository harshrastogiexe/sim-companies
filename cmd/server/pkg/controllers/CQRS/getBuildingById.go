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

func GetBuildingById(ctx *gin.Context) {
	app, found := global.Get[*core.Application](core.APPLICATION_TOKEN)
	if !found {
		panic(fmt.Errorf("instance not found: TOKEN = %s", core.APPLICATION_TOKEN))
	}
	repo := simcompdb.NewRepository(app.DB)
	building, err := repo.GetBuilding(ctx.Param("id"), ctx.QueryArray("include")...)

	if err != nil {
		code := http.StatusInternalServerError
		ctx.JSON(code, models.NewApiError(code, err.Error()))
		return
	}

	if building == nil {
		code := http.StatusNotFound
		ctx.JSON(code, models.NewApiError(code, "building not found"))
		return
	}

	ctx.JSON(200, building)
}
