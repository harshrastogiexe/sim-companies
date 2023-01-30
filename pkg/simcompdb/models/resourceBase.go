package models

type ResourceBase struct {
	ID               int64   `gorm:"not null;type:int;"`
	Name             string  `gorm:"type:NVARCHAR(200);not null;"`
	Image            string  `gorm:"type:NVARCHAR(200);not null;"`
	Transportation   float64 `gorm:"not null;"`
	Retailable       bool    `gorm:"not null;"`
	Research         bool    `gorm:"not null;"`
	ExchangeTradable bool    `gorm:"not null;"`
	RealmAvailable   bool    `gorm:"not null;"`
}
