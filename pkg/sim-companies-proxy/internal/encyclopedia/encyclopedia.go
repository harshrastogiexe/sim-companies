package encyclopedia

import (
	"errors"
	"net/http"

	"github.com/harshrastogiexe/pkg/sim-companies-proxy/internal/core"
	"github.com/harshrastogiexe/pkg/sim-companies-proxy/internal/helper"
	"github.com/harshrastogiexe/sim-companies/pkg/logger"
)

const (
	baseUriResourceEncyclopedia string = "https://www.simcompanies.com/api/v4/en/0/encyclopedia/resources/1/"
	baseUriBuildingEncyclopedia string = "https://www.simcompanies.com/api/v3/0/encyclopedia/buildings/"
	baseUriLevelsEncyclopedia   string = "https://www.simcompanies.com/api/v2/encyclopedia/levels/"
)

var (
	ErrNotFound = errors.New("item not found")
	ErrUnknown  = errors.New("something went wrong")
)

type encyclopedia struct{}

// create new encyclopedia instance
func New() encyclopedia {
	return encyclopedia{}
}

// fetches resource from https://www.simcompanies.com api service
func (encyclopedia) GetResource(id string) (*core.Resource, error) {
	url := baseUriResourceEncyclopedia + id
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := response.Body.Close(); err != nil {
			logger.Log(logger.Warn, "failed to close response body")
		}
	}()

	if response.StatusCode == http.StatusNotFound {
		return nil, ErrNotFound
	} else if startWith := response.StatusCode / 100; startWith == 4 || startWith == 5 {
		return nil, ErrUnknown
	}

	data, err := helper.ParseResponseJsonBody[core.Resource](response)
	if err != nil {
		return nil, err
	}

	return data, err
}

// fetches building from https://www.simcompanies.com api service
func (encyclopedia) GetBuilding(id string) (*core.Building, error) {
	url := baseUriBuildingEncyclopedia + id
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := response.Body.Close(); err != nil {
			logger.Log(logger.Warn, "failed to close response body")
		}
	}()

	switch {
	case response.StatusCode == http.StatusNotFound:
		return nil, ErrNotFound
	case !helper.IsResponseStatusOk(response):
		return nil, ErrUnknown
	}

	data, err := helper.ParseResponseJsonBody[core.Building](response)
	if err != nil {
		return nil, err
	}

	return data, err
}

// fetches building from https://www.simcompanies.com api service
func (encyclopedia) GetLevels() ([]core.Level, error) {
	request, err := http.NewRequest(http.MethodGet, baseUriLevelsEncyclopedia, nil)
	if err != nil {
		return nil, err
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := response.Body.Close(); err != nil {
			logger.Log(logger.Warn, "failed to close response body")
		}
	}()

	switch {
	case response.StatusCode == http.StatusNotFound:
		return nil, ErrNotFound
	case !helper.IsResponseStatusOk(response):
		return nil, ErrUnknown
	}

	data, err := helper.ParseResponseJsonBody[[]core.Level](response)
	if err != nil {
		return nil, err
	}

	return *data, err
}
