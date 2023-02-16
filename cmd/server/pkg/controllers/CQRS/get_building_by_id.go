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
	include := []string{"BuildingBase", "BuildingBase.Images", "DoesProduce", "DoesSell"}

	b, err := repo.GetBuilding(ctx.Param("id"), include...)

	if err != nil {
		code := http.StatusInternalServerError
		ctx.JSON(code, models.NewApiError(code, err.Error()))
		return
	}

	if b == nil {
		code := http.StatusNotFound
		ctx.JSON(code, models.NewApiError(code, "building not found"))
		return
	}

	ctx.JSON(200, models.ConvertBuildingMain(b))
}
