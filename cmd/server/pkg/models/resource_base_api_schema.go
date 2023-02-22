package models

type ResourceBaseApiSchema struct {
	ID               int64   `json:"id,omitempty"`
	Name             string  `json:"name,omitempty"`
	Image            string  `json:"image,omitempty"`
	Transportation   float64 `json:"transportation,omitempty"`
	Retailable       bool    `json:"retailable,omitempty"`
	Research         bool    `json:"research,omitempty"`
	ExchangeTradable bool    `json:"exchange_tradable,omitempty"`
	RealmAvailable   bool    `json:"realm_available,omitempty"`
}
