package models

type Building struct {
	ID             string  `gorm:"type:NVARCHAR(10)"`
	Name           string  `gorm:"type:NVARCHAR(100); not null;"`
	Image          string  `gorm:"type:NVARCHAR(200); not null;"`
	Category       string  `gorm:"type:NVARCHAR(100);not null;"`
	Cost           int     `gorm:"type:integer;not null;"`
	RobotsNeeded   float64 `gorm:"not null;"`
	RealmAvailable bool    `gorm:"not null;"`
	// DoesProduce    []Resource            `gorm:"many2many:does_produce;"`
	// DoesSell       []Resource            `gorm:"many2many:does_sell;"`
	Images    []BuildingLevelImages
	CostUnits float64 `gorm:"not null;"`
	Hours     int64   `gorm:"not null;type:integer;"`
	Wages     float64 `gorm:"not null;"`
}

type BuildingLevelImages struct {
	BuildingID string `gorm:"type:NVARCHAR(10);foreignKey:BuildingID"`
	Level      uint   `gorm:"type:integer; not null;"`
	Image      string `gorm:"type:NVARCHAR(200); not null;"`
}
