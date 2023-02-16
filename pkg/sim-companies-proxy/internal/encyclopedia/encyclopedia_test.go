package encyclopedia_test

import (
	"errors"
	"strconv"
	"testing"

	"github.com/harshrastogiexe/pkg/sim-companies-proxy/internal/encyclopedia"
)

func Test_GetResource(t *testing.T) {
	e := encyclopedia.New()

	t.Run("when function returns error", func(t *testing.T) {
		id := "random_id"
		r, err := e.GetResource(id)

		if err == nil {
			t.Error("error should not be nil")
		}
		if r != nil {
			t.Error("resource should be nil")
		}
	})
	t.Run("returns valid resource item with same id", func(t *testing.T) {
		rid := "1"
		r, err := e.GetResource(rid)
		if err != nil {
			if !errors.Is(err, encyclopedia.ErrUnknown) {
				t.Fatal(err)
			}
		}

		if id := strconv.Itoa(int(r.ID)); id != rid {
			t.Errorf("expected resource id to be \"%s\", got resource id as \"%s\"", rid, id)
		}
	})
}

func Test_GetBuilding(t *testing.T) {
	e := encyclopedia.New()

	t.Run("when function returns error", func(t *testing.T) {
		id := "unknown_id"
		b, err := e.GetBuilding(id)

		if err == nil {
			t.Error("error should not be nil")
		}
		if b != nil {
			t.Error("building should be nil")
		}
	})
	t.Run("returns valid building item with same id", func(t *testing.T) {
		id := "1"
		b, err := e.GetBuilding(id)
		if err != nil {
			if !errors.Is(err, encyclopedia.ErrUnknown) {
				t.Fatal(err)
			}
		}

		if b.ID != id {
			t.Errorf("expected building id to be \"%s\", got building id as \"%s\"", id, b.ID)
		}
	})
}

func Test_GetLevels(t *testing.T) {
	e := encyclopedia.New()
	t.Run("when levels are fetched successfully", func(t *testing.T) {
		l, err := e.GetLevels()
		if err != nil {
			if !errors.Is(err, encyclopedia.ErrUnknown) {
				t.Fatal(err)
			}
		}
		const expectedLevelCount = 7
		if actualCount := len(l); actualCount != expectedLevelCount {
			t.Errorf("expected length of levels to be %d, got %d", expectedLevelCount, actualCount)
		}
	})
}

func Test_GetRatings(t *testing.T) {
	e := encyclopedia.New()
	t.Run("when rating are fetched successfully", func(t *testing.T) {
		ratings, err := e.GetRatings()
		if err != nil {
			if !errors.Is(err, encyclopedia.ErrUnknown) {
				t.Fatal(err)
			}
		}
		const expectedLevelCount = 18
		if actualCount := len(ratings); actualCount != expectedLevelCount {
			t.Errorf("expected length of ratings to be %d, got %d", expectedLevelCount, actualCount)
		}
	})
}
