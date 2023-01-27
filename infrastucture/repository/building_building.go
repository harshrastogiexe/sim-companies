package repository

import "github.com/harshrastogiexe/simcompanies/types"

type buildingInfo struct {
	BuildingID     string  `json:"db_letter" gorm:"primaryKey; column:id"`
	Image          string  `json:"image"`
	Name           string  `json:"name"`
	Category       string  `json:"category"`
	Cost           int64   `json:"cost"`
	RobotsNeeded   float64 `json:"robotsNeeded"`
	RealmAvailable bool    `json:"realmAvailable"`
}

type buildingImage struct {
	Id    string
	Level int
	Path  string
}

func NewBuildingBuilder(building *types.Building) *buildingBuilder {
	return &buildingBuilder{building: building}
}

type buildingBuilder struct {
	building *types.Building
}

func (b buildingBuilder) BuildingInfo() buildingInfo {
	return buildingInfo{
		BuildingID:     b.building.BuildingID,
		Image:          b.building.Image,
		Name:           b.building.Name,
		Category:       b.building.Category,
		Cost:           b.building.Cost,
		RobotsNeeded:   b.building.RobotsNeeded,
		RealmAvailable: b.building.RealmAvailable,
	}
}

func (b buildingBuilder) BuildingLevelImages() (images []buildingImage) {
	for i, image := range b.building.Images {
		images = append(images, buildingImage{Id: b.building.BuildingID, Level: i + 1, Path: image})
	}
	return images
}
