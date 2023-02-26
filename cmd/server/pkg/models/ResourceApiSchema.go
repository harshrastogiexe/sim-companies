package models

import (
	"github.com/harshrastogiexe/sim-companies/pkg/simcompdb/models"
)

type ResourceApiSchema struct {
	ID               int     `json:"id"`
	Name             string  `json:"name"`
	Image            string  `json:"image"`
	Transportation   float64 `json:"transportation"`
	Retailable       bool    `json:"retailable"`
	Research         bool    `json:"research"`
	ExchangeTradable bool    `json:"exchange_tradable"`
	RealmAvailable   bool    `json:"realm_available"`

	SoldAt            *string `json:"sold_at"`
	SoldAtRestaurant  *string `json:"sold_at_restaurant"`
	ProducedAt        string  `json:"produced_at"`
	NeededFor         []int   `json:"needed_for"`
	ImprovesQualityOf []int   `gorm:"many2many:improves_quality_of" json:"improves_quality_of"`
	ProducedFrom      []struct {
		Amount     float64
		ResourceId int
	} `json:"producedFrom"`

	TransportNeeded       float64  `gorm:"not null;" json:"transport_needed"`
	ProducedAnHour        float64  `gorm:"not null;" json:"produced_an_hour"`
	BaseSalary            float64  `gorm:"not null;" json:"base_salary"`
	AverageRetailPrice    *float64 `json:"average_retail_price"`
	MarketSaturation      *float64 `json:"market_saturation"`
	MarketSaturationLabel *string  `json:"market_saturation_label"`
	RetailModeling        *string  `json:"retail_modeling"`
	StoreBaseSalary       *float64 `json:"store_base_salary"`
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

	schema.ProducedFrom = []struct {
		Amount     float64
		ResourceId int
	}{}
	for _, item := range resource.ResourceBase.ProducedFrom {
		schema.ProducedFrom = append(schema.ProducedFrom, struct {
			Amount     float64
			ResourceId int
		}{
			Amount:     item.Amount,
			ResourceId: int(item.FromResourceBaseId),
		})
	}

	return schema
}
