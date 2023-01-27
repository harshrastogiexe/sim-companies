package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harshrastogiexe/app/logger"
	"github.com/harshrastogiexe/cmd/repository"
	"github.com/harshrastogiexe/simcompanies/encyclopedia"
	"github.com/harshrastogiexe/simcompanies/types"
)

const (
	TABLE_BUILDING       = "Building"
	TABLE_BUILDING_IMAGE = "BuildingImages"
)

type BuildingController struct {
	BuildingsRepository repository.BuildingRepository
}

func (c *BuildingController) GetBuildingById(ctx *gin.Context) {
	id := ctx.Param("id")
	building, err := c.BuildingsRepository.FindById(id)

	if err != nil {
		logger.Error.Println(err)
	}

	// if not_found
	if building == nil {
		logger.Info.Println("building not found in database, fetching from simcompanies.com")
		building, err := encyclopedia.GetBuildingInfo(id)
		if err != nil {
			logger.Error.Println(fmt.Errorf("failed to retrieve data from simcompanies api: %w", err))
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		// save to database
		go func() {
			c.saveBuilding(building)
		}()
		ctx.JSON(http.StatusOK, building)
		return
	}
	ctx.JSON(http.StatusOK, building)
}

func (c *BuildingController) saveBuilding(building *types.Building) {
	logger.Info.Println("attempting to save building to database")
	if err := c.BuildingsRepository.Save(building); err != nil {
		logger.Error.Println("failed to save building info:", err)
		return
	}
	logger.Info.Printf("Building id=%s data saved successfully", building.BuildingID)
}
