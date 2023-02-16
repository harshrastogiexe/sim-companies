package models

import "github.com/harshrastogiexe/sim-companies/pkg/simcompdb/models"

type BuildingApiSchema struct {
	ID             string   `json:"id,omitempty"`
	Name           string   `json:"name,omitempty"`
	Category       string   `json:"category,omitempty"`
	Cost           int      `json:"cost,omitempty"`
	RobotsNeeded   float64  `json:"robots_needed,omitempty"`
	RealmAvailable bool     `json:"realm_available,omitempty"`
	Images         []string `json:"images,omitempty"`
	DoesProduce    []int    `json:"doesProduce,omitempty"`
	DoesSell       []int    `json:"doesSell,omitempty"`
	CostUnits      float64  `json:"cost_units,omitempty"`
	Hours          int64    `json:"hours,omitempty"`
	Wages          float64  `json:"wages,omitempty"`
}

func ConvertBuildingMain(building *models.BuildingMain) BuildingApiSchema {
	var b BuildingApiSchema

	b.ID = building.BuildingBaseID
	b.Name = building.BuildingBase.Name
	b.Category = building.BuildingBase.Category
	b.RobotsNeeded = building.BuildingBase.RobotsNeeded
	b.RealmAvailable = building.BuildingBase.RealmAvailable
	b.CostUnits = building.CostUnits
	b.Hours = building.Hours
	b.Wages = building.Wages

	b.Images = []string{}
	for _, image := range building.BuildingBase.Images {
		b.Images = append(b.Images, image.Image)
	}

	b.DoesProduce = []int{}
	for _, resource := range building.DoesProduce {
		b.DoesProduce = append(b.DoesProduce, int(resource.ID))
	}

	b.DoesSell = []int{}
	for _, resource := range building.DoesSell {
		b.DoesSell = append(b.DoesSell, int(resource.ID))
	}
	return b

}
