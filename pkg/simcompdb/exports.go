package simcompdb

import (
	"github.com/harshrastogiexe/sim-companies/pkg/simcompdb/models"
	"gorm.io/gorm"
)

type repository struct {
	ResourceRepository
	BuildingRepository
}

type ResourceRepository interface {
	GetResourceBase(id string) (*models.ResourceBase, error)
	GetResource(id string, preload ...string) (*models.ResourceMain, error)
}

type BuildingRepository interface {
	GetBuilding(id string, preload ...string) (*models.BuildingMain, error)
	ListBuildings() ([]models.BuildingBase, error)
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		&resourceRepo{db},
		&buildingRepo{db},
	}
}
