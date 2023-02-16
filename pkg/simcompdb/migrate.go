package simcompdb

import (
	"github.com/harshrastogiexe/sim-companies/pkg/simcompdb/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	tables := []any{models.BuildingBase{},
		models.ResourceBase{},
		models.BuildingMain{},
		models.ResourceMain{},
		models.ProducedFrom{},
		models.BuildingLevelImage{}}

	return db.AutoMigrate(tables...)
}
