package simcompdb

import (
	"gorm.io/gorm"
)

type repository struct {
	ResourceRepository
	BuildingRepository
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		&resourceRepo{db},
		&buildingRepo{db},
	}
}
