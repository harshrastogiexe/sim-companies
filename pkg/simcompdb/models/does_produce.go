package models

type DoesProduce struct {
	BuildingID string `gorm:"type:NVARCHAR(10)"`
	ResourceID string `gorm:"type:integer"`
}
