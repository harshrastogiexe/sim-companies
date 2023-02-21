package simcompanies

import (
	"github.com/harshrastogiexe/pkg/sim-companies-proxy/core"
	"github.com/harshrastogiexe/pkg/sim-companies-proxy/internal/encyclopedia"
	"github.com/harshrastogiexe/pkg/sim-companies-proxy/internal/market"
)

var (
	Encyclopedia core.Encyclopedia
	Market       core.Market
)

func init() {
	Encyclopedia = encyclopedia.New()
	Market = market.New()
}
