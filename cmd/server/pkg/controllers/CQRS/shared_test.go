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

	if rm.ResourceBaseID != r.ID {
		t.Error("mismatch id")
	}
}
