package core

type ResourceBase struct {
	Name             string  `json:"name"`
	Image            string  `json:"image"`
	ID               int64   `json:"db_letter"`
	Transportation   float64 `json:"transportation"`
	Retailable       bool    `json:"retailable"`
	Research         bool    `json:"research"`
	ExchangeTradable bool    `json:"exchangeTradable"`
	RealmAvailable   bool    `json:"realmAvailable"`
}

type Resource struct {
	ResourceBase
	ProducedFrom          []ProducedFrom `json:"producedFrom"`
	SoldAt                *BuildingBase  `json:"soldAt,omitempty"`
	SoldAtRestaurant      *BuildingBase  `json:"soldAtRestaurant,omitempty"`
	ProducedAt            BuildingBase   `json:"producedAt"`
	NeededFor             []ResourceBase `json:"neededFor"`
	TransportNeeded       int            `json:"transportNeeded"`
	ProducedAnHour        float64        `json:"producedAnHour"`
	BaseSalary            float64        `json:"baseSalary"`
	AverageRetailPrice    *float64       `json:"averageRetailPrice,omitempty"`
	MarketSaturation      *float64       `json:"marketSaturation,omitempty"`
	MarketSaturationLabel *Label         `json:"marketSaturationLabel,omitempty"`
	RetailModeling        *string        `json:"retailModeling,omitempty"`
	StoreBaseSalary       *float64       `json:"storeBaseSalary,omitempty"`
	RetailData            []RetailData   `json:"retailData"`
	ImprovesQualityOf     []interface{}  `json:"improvesQualityOf"`
}

type ProducedFrom struct {
	Resource ResourceBase `json:"resource"`
	Amount   float64      `json:"amount"`
}

type RetailData struct {
	AveragePrice         int64   `json:"averagePrice"`
	AmountSoldRestaurant int64   `json:"amountSoldRestaurant"`
	Demand               float64 `json:"demand"`
	Date                 string  `json:"date"`
	Label                Label   `json:"label"`
}

type Label string

const (
	Extreme Label = "extreme"
)
