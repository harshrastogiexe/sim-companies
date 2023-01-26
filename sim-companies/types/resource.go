package types

import "time"

type Resource struct {
	Name             string  `json:"name" gorm:"column:name"`
	Image            string  `json:"image" gorm:"column:image"`
	ResourceID       uint    `json:"db_letter" gorm:"primaryKey; column:id"`
	Transportation   float64 `json:"transportation" gorm:"column:transporation"`
	Retailable       bool    `json:"retailable" gorm:"column:retailable"`
	Research         bool    `json:"research" gorm:"column:research"`
	ExchangeTradable bool    `json:"exchangeTradable" gorm:"column:ExchangeTradable"`
	RealmAvailable   bool    `json:"realmAvailable" gorm:"column:RealmAvailable"`
}

type Building struct {
	BuildingID     string   `json:"db_letter" gorm:"primaryKey; column:id"`
	Image          string   `json:"image"`
	Images         []string `json:"images"`
	Name           string   `json:"name"`
	Category       string   `json:"category"`
	Cost           int64    `json:"cost"`
	RobotsNeeded   float64  `json:"robotsNeeded"`
	RealmAvailable bool     `json:"realmAvailable"`
}

type EncyclopediaResource struct {
	Name                  string   `json:"name"`
	Image                 string   `json:"image"`
	ResourceID            uint     `json:"db_letter"`
	Transportation        float64  `json:"transportation"`
	Retailable            bool     `json:"retailable"`
	Research              bool     `json:"research"`
	ExchangeTradable      bool     `json:"exchangeTradable"`
	RealmAvailable        bool     `json:"realmAvailable"`
	TransportNeeded       int      `json:"transportNeeded"`
	ProducedAnHour        float64  `json:"producedAnHour"`
	BaseSalary            float64  `json:"baseSalary"`
	MarketSaturation      *float64 `json:"marketSaturation"`
	MarketSaturationLevel *string  `json:"marketSaturationLevel"`

	NeededFor         []Resource
	ImprovesQualityOf []Resource `json:"improvesQualityOf"`
	ProducedAt        Building   `json:"producedAt"`
	SoldAt            *Building  `json:"soldAt"`
	SoldAtRestaurant  *Building  `json:"soldAtRestaurant"`
	ProducesFrom      []struct {
		Resource `json:"resource"`
	} `json:"producedFrom"`
}

type Seller struct {
	ID           int          `json:"id"`
	Company      string       `json:"company"`
	RealmID      int          `json:"realmId"`
	Logo         string       `json:"logo"`
	Certificates int          `json:"certificates"`
	ContestWINS  int          `json:"contest_wins"`
	Npc          bool         `json:"npc"`
	CourseID     *interface{} `json:"courseId"`
	IP           string       `json:"ip"`
}

type MarketItem struct {
	ID       int       `json:"id"`
	Kind     uint      `json:"kind"`
	Quantity uint      `json:"quantity"`
	Quality  uint      `json:"quality"`
	Posted   time.Time `json:"Posted"`
	Fees     float64   `json:"fees"`

	Seller `json:"seller"`
}
