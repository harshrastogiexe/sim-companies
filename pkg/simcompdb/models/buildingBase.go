package models

type BuildingBase struct {
	ID             string                `gorm:"type:NVARCHAR(10)"`
	Name           string                `gorm:"type:NVARCHAR(100); not null;"`
	Image          string                `gorm:"type:NVARCHAR(200); not null;"`
	Images         []BuildingLevelImages `gorm:"foreignKey:BuildingID"`
	Category       string                `gorm:"type:NVARCHAR(100);not null;"`
	Cost           int                   `gorm:"type:integer;not null;"`
	RobotsNeeded   float64               `gorm:"not null;"`
	RealmAvailable bool                  `gorm:"not null;"`
}

type BuildingLevelImages struct {
	BuildingID string `gorm:"type:NVARCHAR(10);not null;"`
	Level      uint   `gorm:"type:integer; not null;"`
	Image      string `gorm:"type:NVARCHAR(200); not null;"`
}
