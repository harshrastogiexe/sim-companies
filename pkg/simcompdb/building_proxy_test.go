package simcompdb_test

import (
	"testing"

	"github.com/harshrastogiexe/sim-companies/pkg/simcompdb"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var DSN = "server=localhost; Database=SIMCOMPANIES; User Id=SA; Password=ComplexP@ssword1234"
	var err error
	db, err = openDB(DSN)

	if err != nil {
		panic(err)
	}
}

func Test_GetBuilding(t *testing.T) {
	repo := simcompdb.NewRepository(db)
	t.Run("returns the correct building", func(t *testing.T) {
		id := "a"
		building, err := repo.GetBuilding(id)
		if err != nil {
			t.Fatal(err)
		}
		if building.BuildingBaseID != id {
			t.Errorf("expected building with id %s, got %s", id, building.BuildingBaseID)
		}
	})

	t.Run("returns the nil when building not found", func(t *testing.T) {
		id := "not_found_id"
		building, err := repo.GetBuilding(id)

		if building != nil {
			t.Errorf("expected building to be nil when not found")
		}

		if err != nil {
			t.Errorf("expected error to be nil if not found")
		}
	})

	t.Run("pre load populate building data", func(t *testing.T) {
		id := "a"
		populate := "BuildingBase"
		building, err := repo.GetBuilding(id, populate)
		if err != nil {
			t.Fatal(err)
		}
		if building.BuildingBase.ID != id {
			t.Errorf("failed to populate '%s', expected id to be '%s', got '%s'", populate, id, building.BuildingBase.ID)
		}
	})
}

func openDB(dsn string) (*gorm.DB, error) {
	dialect := sqlserver.Open(dsn)
	return gorm.Open(dialect, &gorm.Config{})
}
