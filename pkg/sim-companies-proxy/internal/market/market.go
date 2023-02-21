package market

import (
	"fmt"
	"net/http"

	"github.com/harshrastogiexe/pkg/sim-companies-proxy/core"
	"github.com/harshrastogiexe/pkg/sim-companies-proxy/internal/helper"
	"github.com/harshrastogiexe/sim-companies/pkg/logger"
)

const BASE_URI_CURRENT_MARKET_DATA = "https://www.simcompanies.com/api/v3/market/0/%s/"

type market struct{}

func New() *market {
	return &market{}
}

func (market) GetMarketData(id string) ([]core.MarketItem, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(BASE_URI_CURRENT_MARKET_DATA, id), nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := res.Body.Close(); err != nil {
			logger.Log(logger.Warn, "failed to close response body")
		}
	}()

	mis, err := helper.ParseResponseJsonBody[[]core.MarketItem](res)
	if err != nil {
		return nil, err
	}

	return *mis, nil
}
