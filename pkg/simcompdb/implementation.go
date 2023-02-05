package simcompdb

import (
	"github.com/harshrastogiexe/sim-companies/pkg/simcompdb/models"
	"gorm.io/gorm"
)

type SimcompaniesDB struct {
	Database *gorm.DB
}

func (simDB *SimcompaniesDB) GetResourceById(id string) (*models.Resource, error) {
	var resource models.Resource
	preload := []string{"SoldAt", "ProducedFrom"}
	for _, item := range preload {
		simDB.Database.Preload(item)
	}
	result := simDB.Database.Find(&resource, id)
	if err := result.Error; err != nil {
		return nil, err
	}

	if result.RowsAffected == 0 {
		return nil, nil
	}

	return &resource, nil
}

func (simDB *SimcompaniesDB) SaveResource(resource *models.Resource) error {
	return simDB.Database.Debug().Create(resource).Error
}

func (simDB *SimcompaniesDB) SaveBuilding(building *models.Building) error {
	return simDB.Database.Create(building).Error
}

func (simDB *SimcompaniesDB) GetBuildingById(id string) (*models.Building, error) {
	var building models.Building

	buildingResult := simDB.Database.Preload("Images").Find(&building, map[string]any{"ID": id})

	if buildingResult.Error != nil {
		return nil, buildingResult.Error
	}
	if buildingResult.RowsAffected == 0 {
		return nil, nil
	}

	return &building, buildingResult.Error
}

func (simDB *SimcompaniesDB) GetBuildings() ([]models.Building, error) {
	var buildings []models.Building
	tx := simDB.Database.Model(&models.Building{})

	tx.Preload("Images").Find(&buildings).Order("ID")
	return buildings, nil
}
