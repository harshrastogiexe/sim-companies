package models

import "database/sql"

// ------------- Base Building -------------

type BuildingBase struct {
	ID             string  `gorm:"primaryKey;type:NVARCHAR(10) COLLATE SQL_Latin1_General_CP850_BIN;"`
	Name           string  `gorm:"type:NVARCHAR(100); not null;"`
	Image          string  `gorm:"type:NVARCHAR(200); not null;"`
	Category       string  `gorm:"type:NVARCHAR(100);not null;"`
	Cost           int     `gorm:"type:integer;not null;"`
	RobotsNeeded   float64 `gorm:"not null;"`
	RealmAvailable bool    `gorm:"not null;"`

	Images []BuildingLevelImage
}

func (BuildingBase) TableName() string {
	return "building_base"
}

// ------------- Main Building -------------

type BuildingMain struct {
	BuildingBaseID string       `gorm:"primaryKey;type:NVARCHAR(10) COLLATE SQL_Latin1_General_CP850_BIN;column:building_id"`
	BuildingBase   BuildingBase `gorm:"foreignKey:BuildingBaseID;"`

	DoesProduce []ResourceBase `gorm:"many2many:does_produce"`

	DoesSell []ResourceBase `gorm:"many2many:does_sell"`

	CostUnits float64 `gorm:"not null;"`
	Hours     int64   `gorm:"not null;type:integer;"`
	Wages     float64 `gorm:"not null;"`
}

func (BuildingMain) TableName() string {
	return "building_main"
}

// ------------- Building Level Images -------------

type BuildingLevelImage struct {
	BuildingBaseID string `gorm:"type:COLLATE SQL_Latin1_General_CP850_BIN;foreignKey:BuildingBaseID;column:building_id"`
	Level          int    `gorm:"type:integer;not null;"`
	Image          string `gorm:"type:NVARCHAR(200);not null;"`
}

func (BuildingLevelImage) TableName() string {
	return "building_level_image"
}

// ------------- Resource Base -------------

type ResourceBase struct {
	ID               int64   `gorm:"primaryKey;many2many:does_produce;many2many:does_sell"`
	Name             string  `gorm:"type:NVARCHAR(200);not null;"`
	Image            string  `gorm:"type:NVARCHAR(200);not null;"`
	Transportation   float64 `gorm:"not null;"`
	Retailable       bool    `gorm:"not null;"`
	Research         bool    `gorm:"not null;"`
	ExchangeTradable bool    `gorm:"not null;"`
	RealmAvailable   bool    `gorm:"not null;"`

	ProducedFrom []ProducedFrom
}

func (ResourceBase) TableName() string {
	return "resource_base"
}

// ------------- Resource Main -------------

type ResourceMain struct {
	ResourceBaseID int64        `gorm:"primaryKey;column:resource_id"`
	ResourceBase   ResourceBase `gorm:"foreignKey:ResourceBaseID;"`

	SoldAtBuildingID sql.NullString `gorm:"type:COLLATE SQL_Latin1_General_CP850_BIN"`
	SoldAt           *BuildingBase  `gorm:"foreignKey:SoldAtBuildingID"`

	SoldAtRestaurantID sql.NullString `gorm:"type:COLLATE SQL_Latin1_General_CP850_BIN"`
	SoldAtRestaurant   *BuildingBase  `gorm:"foreignKey:SoldAtRestaurantID"`

	ProducedAtBuildingID sql.NullString
	ProducedAt           *BuildingBase `gorm:"not null;foreignKey:ProducedAtBuildingID"`

	NeededFor []ResourceBase `gorm:"many2many:needed_for"`

	ImprovesQualityOf []ResourceBase `gorm:"many2many:improves_quality_of"`

	TransportNeeded       float64 `gorm:"not null;"`
	ProducedAnHour        float64 `gorm:"not null;"`
	BaseSalary            float64 `gorm:"not null;"`
	AverageRetailPrice    sql.NullFloat64
	MarketSaturation      sql.NullFloat64
	MarketSaturationLabel sql.NullString `gorm:"type:NVARCHAR(200)"`
	RetailModeling        sql.NullString `gorm:"type:NVARCHAR(200)"`
	StoreBaseSalary       sql.NullFloat64
}

func (ResourceMain) TableName() string {
	return "resource_main"
}

// ---------------- Produced From ----------------------------------

type ProducedFrom struct {
	ResourceBaseID int64        `gorm:"column:resource_id"`
	ResourceBase   ResourceBase `gorm:"foreignKey:ResourceBaseID"`
	Amount         float64      `gorm:"not null;"`
}
