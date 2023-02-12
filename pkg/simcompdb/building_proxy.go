package simcompdb

import (
	"github.com/harshrastogiexe/sim-companies/pkg/simcompdb/models"
	"gorm.io/gorm"
)

type buildingRepo struct {
	*gorm.DB
}

func (repo *buildingRepo) GetBuilding(id string, preload ...string) (*models.BuildingMain, error) {
	var item models.BuildingMain

	var tx = repo.Debug()
	for _, preloadTable := range preload {
		tx = tx.Preload(preloadTable)
	}

	result := tx.Where(&models.BuildingMain{BuildingBaseID: id}).Find(&item)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, nil
	}
	return &item, nil
}

func (repo *buildingRepo) ListBuildings() ([]models.BuildingBase, error) {
	var building []models.BuildingBase
	result := repo.Preload("Images").Find(&building)
	if result.Error != nil {
		return []models.BuildingBase{}, repo.Error
	}
	return building, nil
}
