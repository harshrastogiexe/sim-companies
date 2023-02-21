package core

type Market interface {
	// gets the current active market items based on resource id
	GetMarketData(r string) ([]MarketItem, error)
}
