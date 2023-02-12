package models

import "database/sql"

type Resource struct {
	ID                    int64   `gorm:"not null;autoIncrement;many2many:does_produce;many2many:does_sell;"`
	Name                  string  `gorm:"type:NVARCHAR(200);not null;"`
	Image                 string  `gorm:"type:NVARCHAR(200);not null;"`
	Transportation        float64 `gorm:"not null;"`
	Retailable            bool    `gorm:"not null;"`
	Research              bool    `gorm:"not null;"`
	ExchangeTradable      bool    `gorm:"not null;"`
	RealmAvailable        bool    `gorm:"not null;"`
	ProducedFrom          []ProducedFromLegacy
	SoldAtBuildingId      sql.NullString
	SoldAt                *Building `gorm:"foreignKey:SoldAtBuildingId"`
	SoldAtRestaurantId    sql.NullString
	SoldAtRestaurant      *Building `gorm:"foreignKey:SoldAtRestaurantId"`
	ProducedAtBuildingId  string
	ProducedAt            *Building  `gorm:"foreignKey:ProducedAtBuildingId"`
	NeededFor             []Resource `gorm:"many2many:needed_for"`
	TransportNeeded       float64
	ProducedAnHour        float64
	BaseSalary            float64
	AverageRetailPrice    sql.NullFloat64
	MarketSaturation      sql.NullFloat64
	MarketSaturationLabel sql.NullString
	RetailModeling        sql.NullString
	StoreBaseSalary       sql.NullFloat64
	ImprovesQualityOf     []Resource `gorm:"many2many:improves_quality"`
}

type ProducedFromLegacy struct {
	ResourceID             int64
	Resource               Resource `gorm:"not null;foreignKey:ResourceID"`
	ProducedFromResourceID int64
	ProducedFromResource   Resource `gorm:"not null;foreignKey:ProducedFromResourceID"`
	Amount                 float64  `gorm:"not null;"`
}

func (ProducedFromLegacy) TableName() string {
	return "produced_from"
}
