package simcompdb

import (
	"github.com/harshrastogiexe/sim-companies/pkg/simcompdb/models"
	"gorm.io/gorm"
)

type resourceRepo struct {
	*gorm.DB
}

// find item in database, if error occur returns error, if value not found return value and error as nil
func (repo *resourceRepo) GetResourceBase(id string) (*models.ResourceBase, error) {
	return findById[models.ResourceBase](id, repo.DB)
}

func (repo *resourceRepo) GetResource(id string, preload ...string) (*models.ResourceMain, error) {
	var item models.ResourceMain

	var tx = repo.Debug()
	for _, preloadTable := range preload {
		tx = tx.Preload(preloadTable)
	}

	result := tx.Find(&item, id)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, nil
	}
	return &item, nil
}
