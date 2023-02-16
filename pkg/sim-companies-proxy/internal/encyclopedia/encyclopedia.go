package encyclopedia

import (
	"errors"
	"net/http"

	"github.com/harshrastogiexe/pkg/sim-companies-proxy/core"
	"github.com/harshrastogiexe/pkg/sim-companies-proxy/internal/helper"
	"github.com/harshrastogiexe/sim-companies/pkg/logger"
)

// sim companies relative url
const (
	baseUriResourceEncyclopedia string = "https://www.simcompanies.com/api/v4/en/0/encyclopedia/resources/1/"
	baseUriBuildingEncyclopedia string = "https://www.simcompanies.com/api/v3/0/encyclopedia/buildings/"
	baseUriLevelsEncyclopedia   string = "https://www.simcompanies.com/api/v2/encyclopedia/levels/"
	baseUriRatingEncyclopedia   string = "https://www.simcompanies.com/api/v2/encyclopedia/ratings/"
)

// errors
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
	req, err := http.NewRequest(http.MethodGet, url, nil)
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

	if res.StatusCode == http.StatusNotFound {
		return nil, ErrNotFound
	} else if startWith := res.StatusCode / 100; startWith == 4 || startWith == 5 {
		return nil, ErrUnknown
	}

	r, err := helper.ParseResponseJsonBody[core.Resource](res)
	if err != nil {
		return nil, err
	}

	return r, err
}

// fetches building from https://www.simcompanies.com api service
func (encyclopedia) GetBuilding(id string) (*core.Building, error) {
	url := baseUriBuildingEncyclopedia + id
	req, err := http.NewRequest(http.MethodGet, url, nil)
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

	switch {
	case res.StatusCode == http.StatusNotFound:
		return nil, ErrNotFound
	case !helper.IsResponseStatusOk(res):
		return nil, ErrUnknown
	}

	b, err := helper.ParseResponseJsonBody[core.Building](res)
	if err != nil {
		return nil, err
	}

	return b, err
}

// fetches level information from https://www.simcompanies.com api service "Contractor","Family Business" etc.
func (encyclopedia) GetLevels() ([]core.Level, error) {
	req, err := http.NewRequest(http.MethodGet, baseUriLevelsEncyclopedia, nil)
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

	switch {
	case res.StatusCode == http.StatusNotFound:
		return nil, ErrNotFound
	case !helper.IsResponseStatusOk(res):
		return nil, ErrUnknown
	}

	levels, err := helper.ParseResponseJsonBody[[]core.Level](res)
	if err != nil {
		return nil, err
	}

	return *levels, err
}

// fetches rating from https://www.simcompanies.com api service
func (encyclopedia) GetRatings() ([]core.Rating, error) {
	req, err := http.NewRequest(http.MethodGet, baseUriRatingEncyclopedia, nil)
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

	switch {
	case res.StatusCode == http.StatusNotFound:
		return nil, ErrNotFound
	case !helper.IsResponseStatusOk(res):
		return nil, ErrUnknown
	}

	rs, err := helper.ParseResponseJsonBody[[]core.Rating](res)
	if err != nil {
		return nil, err
	}

	return *rs, err
}
