package cqrs

import (
	"testing"

	simcompanies "github.com/harshrastogiexe/pkg/sim-companies-proxy"
)

func Test_convertApiResourceToBuildResource(t *testing.T) {
	r, err := simcompanies.Encyclopedia.GetResource("2")

	if err != nil {
		t.Fatal(err)
	}
	rm := convertApiResourceToBuildResource(r)

	if rm.ID != r.ID {
		t.Error("mismatch id")
	}

	if rm.Name != r.Name {
		t.Error("mismatch id")
	}

	if len(rm.ProducedFrom) != len(r.ProducedFrom) {
		t.Fail()
	}

	for i, producedFrom := range r.ProducedFrom {
		if rm.ProducedFrom[i].ProducedFromResource.ID != producedFrom.Resource.ID || rm.ProducedFrom[i].Amount != producedFrom.Amount {
			t.Error("mismatch resource id or amount during conversion")
		}
	}
}
