package simcompdb

import (
	"github.com/harshrastogiexe/sim-companies/pkg/simcompdb/models"
	"gorm.io/gorm"
)

type BuildingRepository interface {
	GetBuilding(id string, preload ...string) (*models.BuildingMain, error)
	ListBuildings() ([]models.BuildingBase, error)
}

type buildingRepo struct {
	*gorm.DB
}

func (repo *buildingRepo) GetBuilding(id string, preload ...string) (*models.BuildingMain, error) {
	var buildings models.BuildingMain

	var tx = repo.Debug()
	for _, preloadTable := range preload {
		tx = tx.Preload(preloadTable)
	}

	r := tx.Where(&models.BuildingMain{BuildingBaseID: id}).Find(&buildings)
	if r.Error != nil {
		return nil, r.Error
	}

	if r.RowsAffected == 0 {
		return nil, nil
	}
	return &buildings, nil
}

func (repo *buildingRepo) ListBuildings() ([]models.BuildingBase, error) {
	var buildings []models.BuildingBase
	result := repo.Preload("Images").Find(&buildings)
	if result.Error != nil {
		return []models.BuildingBase{}, repo.Error
	}
	return buildings, nil
}
