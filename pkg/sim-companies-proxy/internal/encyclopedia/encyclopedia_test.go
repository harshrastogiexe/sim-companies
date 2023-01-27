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