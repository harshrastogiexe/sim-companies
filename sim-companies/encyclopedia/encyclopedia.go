package encyclopedia

import (
	"log"
	"net/http"

	"github.com/harshrastogiexe/simcompanies/types"
)

const bASE_URI_ENCYCLOPEDIA string = "https://www.simcompanies.com/api/v4/en/0/encyclopedia/resources/1/"
const bASE_URI_MARKET string = "https://www.simcompanies.com/api/v3/market/0/"
const bASE_URI_BUILDING string = "https://www.simcompanies.com/api/v3/0/encyclopedia/buildings/"

func GetInfo(id string) (*types.EncyclopediaResource, error) {
	log.Println("fetching info for resource id", id)
	req, err := http.NewRequest(http.MethodGet, bASE_URI_ENCYCLOPEDIA+id, nil)
	if err != nil {
		return nil, err
	}
	resource, err := sendRequest[types.EncyclopediaResource](req)
	if err != nil {
		return nil, err
	}
	return resource, nil
}

func GetResourceInfoFromMarket(id string) ([]types.MarketItem, error) {
	log.Println("fetching info for resource id", id)
	req, err := http.NewRequest(http.MethodGet, bASE_URI_MARKET+id, nil)
	if err != nil {
		return nil, err
	}
	resource, err := sendRequest[[]types.MarketItem](req)
	if err != nil {
		return nil, err
	}
	return *resource, nil
}

func GetBuildingInfo(id string) (*types.Building, error) {
	log.Println("fetching info for building id", id)
	req, err := http.NewRequest(http.MethodGet, bASE_URI_BUILDING+id, nil)
	if err != nil {
		return nil, err
	}
	building, err := sendRequest[types.Building](req)
	if err != nil {
		return nil, err
	}
	return building, nil
}
