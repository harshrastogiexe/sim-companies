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

func convertApiResourceToBuildResource(resource *core.Resource) *models.Resource {
	resourceModel := &models.Resource{
		ID:                   resource.ID,
		Name:                 resource.Name,
		Image:                resource.Image,
		Transportation:       resource.Transportation,
		Retailable:           resource.Retailable,
		Research:             resource.Research,
		ExchangeTradable:     resource.ExchangeTradable,
		RealmAvailable:       resource.RealmAvailable,
		ProducedAtBuildingId: resource.ProducedAt.ID,
		TransportNeeded:      resource.TransportNeeded,
		ProducedAnHour:       resource.ProducedAnHour,
		BaseSalary:           resource.BaseSalary,
	}

	if resource.ProducedFrom != nil {
		resourceModel.ProducedFrom = func() (producedFrom []models.ProducedFrom) {
			for _, item := range resource.ProducedFrom {
				producedFrom = append(producedFrom, models.ProducedFrom{
					ResourceID:           resource.ID,
					ProducedFromResource: models.Resource{ID: item.Resource.ID},
					Amount:               item.Amount,
				})
			}
			return producedFrom
		}()
	}

	if resource.SoldAt != nil {
		resourceModel.SoldAtBuildingId = sql.NullString{String: resource.SoldAt.ID}
	}

	if resource.MarketSaturationLabel != nil {
		resourceModel.MarketSaturationLabel = sql.NullString{String: string(*resource.MarketSaturationLabel)}
	}

	if resource.RetailModeling != nil {
		resourceModel.RetailModeling = sql.NullString{String: *resource.RetailModeling}
	}

	if resource.SoldAtRestaurant != nil {
		resourceModel.SoldAtRestaurantId = sql.NullString{String: resource.SoldAtRestaurant.ID}
	}

	if resource.AverageRetailPrice != nil {
		resourceModel.AverageRetailPrice = sql.NullFloat64{Float64: *resource.AverageRetailPrice}
	}

	if resource.MarketSaturation != nil {
		resourceModel.MarketSaturation = sql.NullFloat64{Float64: *resource.MarketSaturation}
	}

	if resource.StoreBaseSalary != nil {
		resourceModel.StoreBaseSalary = sql.NullFloat64{Float64: *resource.StoreBaseSalary}
	}

	return resourceModel
}
