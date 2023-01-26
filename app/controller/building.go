package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harshrastogiexe/app/logger"
	"github.com/harshrastogiexe/simcompanies/encyclopedia"
	"github.com/harshrastogiexe/simcompanies/types"
	"gorm.io/gorm"
)

const (
	TABLE_BUILDING       = "Building"
	TABLE_BUILDING_IMAGE = "BuildingImages"
)

type BuildingController struct {
	GormDB *gorm.DB
}

type BuildingWithoutImages struct {
	BuildingID     string  `json:"db_letter" gorm:"primaryKey; column:id"`
	Image          string  `json:"image"`
	Name           string  `json:"name"`
	Category       string  `json:"category"`
	Cost           int64   `json:"cost"`
	RobotsNeeded   float64 `json:"robotsNeeded"`
	RealmAvailable bool    `json:"realmAvailable"`
}

type BuildingImage struct {
	Id    string
	Level int
	Path  string
}

func (c *BuildingController) GetBuildingById(ctx *gin.Context) {
	id := ctx.Param("id")
	var building BuildingWithoutImages

	// check in db

	rc := c.GormDB.Table(TABLE_BUILDING).Find(&building, &id).RowsAffected

	// if not_found
	if rc == 0 {
		logger.Info.Println("fetching from simcompanies.com")
		building, err := encyclopedia.GetBuildingInfo(id)
		if err != nil {
			logger.Error.Println(fmt.Errorf("failed to retrieve data from simcompanies api: %w", err))
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		go func() {
			logger.Info.Println("attempting to save building to database")
			if err := c.saveBuildingToDb(building); err != nil {
				logger.Error.Println("failed to save building info:", err)
				return
			}
			logger.Info.Printf("Building id=%s data saved successfully", building.BuildingID)
		}()
		// save to database
		ctx.JSON(http.StatusOK, building)
		return
	}
	ctx.JSON(http.StatusOK, building)
}

func (c *BuildingController) saveBuildingToDb(b *types.Building) error {
	errTransaction := c.GormDB.Transaction(func(tx *gorm.DB) error {
		err := tx.Table(TABLE_BUILDING).Create(b).Error
		if err != nil {
			return fmt.Errorf("failed to save building:\n%w", err)
		}
		logger.Info.Println("building info save successfully")
		images := convertToBuildingImage(b)
		err = tx.Table(TABLE_BUILDING_IMAGE).Create(&images).Error
		if err != nil {
			return fmt.Errorf("failed to save building images:\n\n%w", err)
		}
		logger.Info.Println("building level images info save successfully")
		return nil
	})
	if errTransaction != nil {
		logger.Error.Println("transaction failed")
	}
	return errTransaction
}

func convertToBuildingImage(b *types.Building) []BuildingImage {
	var images []BuildingImage
	for i, image := range b.Images {
		images = append(images, BuildingImage{Id: b.BuildingID, Level: i + 1, Path: image})
	}
	return images
}
