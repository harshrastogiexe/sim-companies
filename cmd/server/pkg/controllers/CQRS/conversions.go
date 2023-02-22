package cqrs

import (
	"database/sql"

	"github.com/harshrastogiexe/pkg/sim-companies-proxy/core"
	"github.com/harshrastogiexe/sim-companies/pkg/simcompdb/models"
)

func convertBuildingApiToBuildingModel(b *core.Building) *models.Building {
	return &models.Building{
		ID:             b.ID,
		Name:           b.Name,
		Image:          b.Image,
		Category:       b.Category,
		Cost:           b.Cost,
		RobotsNeeded:   b.RobotsNeeded,
		RealmAvailable: b.RealmAvailable,
		CostUnits:      b.CostUnits,
		Hours:          b.Hours,
		Wages:          b.Wages,
		Images: func() []models.BuildingLevelImages {
			var images []models.BuildingLevelImages
			for i, img := range b.Images {
				images = append(images, models.BuildingLevelImages{
					BuildingID: b.ID,
					Level:      uint(i + 1), // using image
					Image:      img,
				})
			}
			return images
		}(),
	}
}

func convertApiResourceToBuildResource(r *core.Resource) *models.ResourceMain {
	rm := &models.ResourceMain{
		ResourceBaseID:       r.ID,
		SoldAtBuildingID:     sql.NullString{String: r.SoldAt.ID},
		SoldAtRestaurantID:   sql.NullString{String: r.SoldAtRestaurant.ID},
		ProducedAtBuildingID: sql.NullString{String: r.ProducedAt.ID},
	}

	return rm
}

func convertResourceBaseToResourceApi(r *core.ResourceBase) {

}
