package simcompdb

import (
	"github.com/harshrastogiexe/sim-companies/pkg/simcompdb/models"
	"gorm.io/gorm"
)

type ResourceRepository interface {
	GetResourceBase(id string) (*models.ResourceBase, error)
	GetResource(id string, preload ...string) (*models.ResourceMain, error)
}

type resourceRepo struct {
	*gorm.DB
}

// find item in database, if error occur returns error, if value not found return value and error as nil
func (repo *resourceRepo) GetResourceBase(id string) (*models.ResourceBase, error) {
	return findById[models.ResourceBase](id, repo.DB)
}

func (repo *resourceRepo) GetResource(id string, preload ...string) (*models.ResourceMain, error) {
	var resource models.ResourceMain

	var tx = repo.Debug()
	for _, preloadTable := range preload {
		tx = tx.Preload(preloadTable)
	}

	r := tx.Find(&resource, id)
	if r.Error != nil {
		return nil, r.Error
	}

	if r.RowsAffected == 0 {
		return nil, nil
	}
	return &resource, nil
}
