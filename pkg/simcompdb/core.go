package simcompdb

import (
	"github.com/harshrastogiexe/sim-companies/pkg/simcompdb/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		// &models.Building{},
		&models.BuildingBase{},
		&models.BuildingLevelImages{},
		// &models.Resource{},
		&models.ResourceBase{},
	)
}
