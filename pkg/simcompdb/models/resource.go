package models

import "database/sql"

type Resource struct {
	ResourceBase          ResourceBase `gorm:"embedded"`
	ProducedFrom          []ProducedFrom
	SoldAtBuildingId      string
	SoldAt                *BuildingBase `gorm:"foreignKey:SoldAtBuildingId"`
	SoldAtRestaurantId    string
	SoldAtRestaurant      *BuildingBase `gorm:"foreignKey:SoldAtRestaurantId"`
	ProductAtBuildingId   string
	ProducedAt            BuildingBase `gorm:"foreignKey:ProductAtBuildingId"`
	NeededFor             []ResourceBase
	TransportNeeded       float64
	ProducedAnHour        float64
	BaseSalary            float64
	AverageRetailPrice    sql.NullFloat64
	MarketSaturation      sql.NullFloat64
	MarketSaturationLabel sql.NullString
	RetailModeling        sql.NullString
	StoreBaseSalary       sql.NullFloat64
	ImprovesQualityOf     []ResourceBase
}
