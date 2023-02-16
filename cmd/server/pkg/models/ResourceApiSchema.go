package models

import (
	"github.com/harshrastogiexe/sim-companies/pkg/simcompdb/models"
)

type ResourceApiSchema struct {
	ID               int
	Name             string
	Image            string
	Transportation   float64
	Retailable       bool
	Research         bool
	ExchangeTradable bool
	RealmAvailable   bool

	SoldAt            *string
	SoldAtRestaurant  *string
	ProducedAt        string
	NeededFor         []int
	ImprovesQualityOf []int `gorm:"many2many:improves_quality_of"`

	TransportNeeded       float64 `gorm:"not null;"`
	ProducedAnHour        float64 `gorm:"not null;"`
	BaseSalary            float64 `gorm:"not null;"`
	AverageRetailPrice    *float64
	MarketSaturation      *float64
	MarketSaturationLabel *string
	RetailModeling        *string
	StoreBaseSalary       *float64
}

func ConvertResourceMain(resource *models.ResourceMain) ResourceApiSchema {
	schema := ResourceApiSchema{
		ID:              int(resource.ResourceBaseID),
		Name:            resource.ResourceBase.Name,
		Image:           resource.ResourceBase.Image,
		Transportation:  resource.TransportNeeded,
		Retailable:      resource.ResourceBase.Retailable,
		Research:        resource.ResourceBase.Research,
		RealmAvailable:  resource.ResourceBase.RealmAvailable,
		TransportNeeded: resource.TransportNeeded,
		ProducedAnHour:  resource.ProducedAnHour,
		BaseSalary:      resource.BaseSalary,
		ProducedAt:      resource.ProducedAtBuildingID.String,
	}

	if resource.SoldAtBuildingID.Valid {
		schema.SoldAt = &resource.SoldAtBuildingID.String
	}

	if resource.SoldAtRestaurantID.Valid {
		schema.SoldAtRestaurant = &resource.SoldAtRestaurantID.String
	}

	if resource.AverageRetailPrice.Valid {
		schema.AverageRetailPrice = &resource.AverageRetailPrice.Float64
	}

	if resource.MarketSaturation.Valid {
		schema.MarketSaturation = &resource.MarketSaturation.Float64
	}

	if resource.MarketSaturationLabel.Valid {
		schema.MarketSaturationLabel = &resource.MarketSaturationLabel.String
	}

	if resource.RetailModeling.Valid {
		schema.RetailModeling = &resource.RetailModeling.String
	}

	if resource.StoreBaseSalary.Valid {
		schema.StoreBaseSalary = &resource.StoreBaseSalary.Float64
	}

	if resource.ImprovesQualityOf == nil {
		schema.ImprovesQualityOf = []int{}
	} else {
		var resources = []int{}
		for _, rb := range resource.ImprovesQualityOf {
			resources = append(resources, int(rb.ID))
		}
		schema.ImprovesQualityOf = resources
	}

	if resource.NeededFor == nil {
		schema.NeededFor = []int{}
	} else {
		var resources = []int{}
		for _, rb := range resource.NeededFor {
			resources = append(resources, int(rb.ID))
		}
		schema.NeededFor = resources
	}

	return schema
}
