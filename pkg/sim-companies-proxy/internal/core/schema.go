package core

type Resource struct {
	ID                    uint           `json:"db_letter"`
	Name                  string         `json:"name"`
	Image                 string         `json:"image"`
	Transportation        float64        `json:"transportation"`
	Retailable            bool           `json:"retailable"`
	Research              bool           `json:"research"`
	ExchangeTradable      bool           `json:"exchangeTradable"`
	RealmAvailable        bool           `json:"realmAvailable"`
	ProducedFrom          []ProducedFrom `json:"producedFrom"`
	SoldAt                *Building      `json:"soldAt,omitempty"`
	SoldAtRestaurant      *Building      `json:"soldAtRestaurant,omitempty"`
	ProducedAt            Building       `json:"producedAt"`
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

type Building struct {
	Id             string   `json:"db_letter"`
	Image          string   `json:"image"`
	Images         []string `json:"images"`
	Name           string   `json:"name"`
	Category       string   `json:"category"`
	Cost           int      `json:"cost"`
	RobotsNeeded   float64  `json:"robotsNeeded"`
	RealmAvailable bool     `json:"realmAvailable"`
}

type ProducedFrom struct {
	Resource ResourceBase `json:"resource"`
	Amount   float64      `json:"amount"`
}

type ResourceBase struct {
	Name             string  `json:"name"`
	Image            string  `json:"image"`
	DBLetter         int64   `json:"db_letter"`
	Transportation   float64 `json:"transportation"`
	Retailable       bool    `json:"retailable"`
	Research         bool    `json:"research"`
	ExchangeTradable bool    `json:"exchangeTradable"`
	RealmAvailable   bool    `json:"realmAvailable"`
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
