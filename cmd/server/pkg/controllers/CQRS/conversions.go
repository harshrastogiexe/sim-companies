package cqrs

import (
	"database/sql"

	"github.com/harshrastogiexe/pkg/sim-companies-proxy/core"
	"github.com/harshrastogiexe/sim-companies/pkg/simcompdb/models"
)

func convertBuildingApiToBuildingModel(building *core.Building) *models.Building {
	return &models.Building{
		ID:             building.ID,
		Name:           building.Name,
		Image:          building.Image,
		Category:       building.Category,
		Cost:           building.Cost,
		RobotsNeeded:   building.RobotsNeeded,
		RealmAvailable: building.RealmAvailable,
		CostUnits:      building.CostUnits,
		Hours:          building.Hours,
		Wages:          building.Wages,
		Images: func() []models.BuildingLevelImages {
			var images []models.BuildingLevelImages
			for i, image := range building.Images {
				images = append(images, models.BuildingLevelImages{
					BuildingID: building.ID,
					Level:      uint(i + 1),
					Image:      image,
				})
			}
			return images
		}(),
	}
}

func convertApiResourceToBuildResource(resource *core.Resource) *models.ResourceMain {
	m := &models.ResourceMain{
		ResourceBaseID:       resource.ID,
		SoldAtBuildingID:     sql.NullString{String: resource.SoldAt.ID},
		SoldAtRestaurantID:   sql.NullString{String: resource.SoldAtRestaurant.ID},
		ProducedAtBuildingID: sql.NullString{String: resource.ProducedAt.ID},
	}

	return m
}
