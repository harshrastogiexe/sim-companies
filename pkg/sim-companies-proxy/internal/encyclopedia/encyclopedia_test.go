package encyclopedia_test

import (
	"errors"
	"fmt"
	"strconv"
	"testing"

	"github.com/harshrastogiexe/pkg/sim-companies-proxy/internal/encyclopedia"
)

func Test_GetResource(t *testing.T) {
	e := encyclopedia.New()

	t.Run("when function returns error", func(t *testing.T) {
		resourceId := "random_id"
		resource, err := e.GetResource(resourceId)

		if err == nil {
			t.Error("error should not be nil")
		}
		if resource != nil {
			t.Error("resource should be nil")
		}
	})
	t.Run("returns valid resource item with same id", func(t *testing.T) {
		resourceId := "1"
		resource, err := e.GetResource(resourceId)
		if err != nil {
			if !errors.Is(err, encyclopedia.ErrUnknown) {
				t.Fatal(err)
			}
		}

		if id := strconv.Itoa(int(resource.ID)); id != resourceId {
			t.Errorf("expected resource id to be \"%s\", got resource id as \"%s\"", resourceId, id)
		}
	})
}

func Test_GetBuilding(t *testing.T) {
	e := encyclopedia.New()

	t.Run("when function returns error", func(t *testing.T) {
		buildingId := "unknown_id"
		building, err := e.GetBuilding(buildingId)

		if err == nil {
			t.Error("error should not be nil")
		}
		if building != nil {
			t.Error("building should be nil")
		}
	})
	t.Run("returns valid building item with same id", func(t *testing.T) {
		buildingId := "1"
		building, err := e.GetBuilding(buildingId)
		fmt.Printf("%+v", building)
		if err != nil {
			if !errors.Is(err, encyclopedia.ErrUnknown) {
				t.Fatal(err)
			}
		}

		if building.ID != buildingId {
			t.Errorf("expected building id to be \"%s\", got building id as \"%s\"", buildingId, building.ID)
		}
	})
}

func Test_GetLevels(t *testing.T) {
	e := encyclopedia.New()
	t.Run("when levels are fetched successfully", func(t *testing.T) {
		levels, err := e.GetLevels()
		if err != nil {
			if !errors.Is(err, encyclopedia.ErrUnknown) {
				t.Fatal(err)
			}
		}
		const expectedLevelCount = 7
		if actualCount := len(levels); actualCount != expectedLevelCount {
			t.Errorf("expected length of levels to be %d, got %d", expectedLevelCount, actualCount)
		}
	})
}
