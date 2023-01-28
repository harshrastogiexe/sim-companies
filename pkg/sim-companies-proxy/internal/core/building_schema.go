package core



type BuildingBase struct {
	ID             string   `json:"db_letter"`
	Image          string   `json:"image"`
	Images         []string `json:"images"`
	Name           string   `json:"name"`
	Category       string   `json:"category"`
	Cost           int      `json:"cost"`
	RobotsNeeded   float64  `json:"robotsNeeded"`
	RealmAvailable bool     `json:"realmAvailable"`
}

type Building struct {
	BuildingBase
	DoesProduce []ResourceBase `json:"doesProduce"`
	DoesSell    []ResourceBase `json:"doesSell"`
	CostUnits   float64        `json:"costUnits"`
	Hours       int64          `json:"hours"`
	Wages       float64        `json:"wages"`
}
