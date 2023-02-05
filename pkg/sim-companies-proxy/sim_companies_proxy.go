package simcompanies

import (
	"github.com/harshrastogiexe/pkg/sim-companies-proxy/core"
	"github.com/harshrastogiexe/pkg/sim-companies-proxy/internal/encyclopedia"
)

var (
	Encyclopedia core.Encyclopedia
)

func init() {
	Encyclopedia = encyclopedia.New()
}
