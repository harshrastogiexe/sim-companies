package repository

import (
	"github.com/harshrastogiexe/simcompanies/types"
	"gorm.io/gorm"
)

const (
	TABLE_BUILDING       = "Building"
	TABLE_BUILDING_IMAGE = "BuildingImages"
)

type BuildingRepository interface {
	Save(b *types.Building) error

	// queries the database for id, if building is not found returns "nil" as `types.Building`
	FindById(id string) (*types.Building, error)
}

type SqlServerBuildingRepository struct {
	GormDB *gorm.DB
}

func (r *SqlServerBuildingRepository) Save(b *types.Building) error {
	builder := NewBuildingBuilder(b)
	buildingInfo, images := builder.BuildingInfo(), builder.BuildingLevelImages()
	if err := r.GormDB.Table(TABLE_BUILDING).Create(&buildingInfo).Error; err != nil {
		return err
	}
	if err := r.GormDB.Table(TABLE_BUILDING_IMAGE).Create(&images).Error; err != nil {
		return err
	}
	return nil
}

// queries the database for id, if building is not found returns "nil" as `types.Building`
func (r *SqlServerBuildingRepository) FindById(id string) (*types.Building, error) {
	buildingInfo, err := r.getBuildingInfo(id)
	if err != nil || buildingInfo == nil {
		return nil, err
	}

	levelImages, err := r.getBuildingLevelImages(id)
	building := newBuilding(buildingInfo, levelImages)
	if err != nil {
		return building, err
	}
	return building, nil
}

// queries the database for id, if building is not found returns "nil" as `types.Building`
func (r *SqlServerBuildingRepository) getBuildingInfo(id string) (*buildingInfo, error) {
	var buildingInfo buildingInfo
	result := r.GormDB.Table(TABLE_BUILDING).Find(&buildingInfo, &id)
	if result.Error != nil {
		return nil, result.Error
	}
	// building not found
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return &buildingInfo, nil
}

// queries the building images for id, if building id not found empty slice
func (r *SqlServerBuildingRepository) getBuildingLevelImages(id string) ([]buildingImage, error) {
	var images []buildingImage
	result := r.GormDB.Table(TABLE_BUILDING_IMAGE).Find(&images, &id)
	if result.Error != nil {
		return images, result.Error
	}
	return images, nil
}

func newBuilding(b *buildingInfo, buildingImageLevels []buildingImage) *types.Building {
	return &types.Building{
		BuildingID:     b.BuildingID,
		Image:          b.Image,
		Name:           b.Name,
		Category:       b.Category,
		Cost:           b.Cost,
		RobotsNeeded:   b.RobotsNeeded,
		RealmAvailable: b.RealmAvailable,
		Images: func() []string {
			images := []string{}
			for _, levelImage := range buildingImageLevels {
				images = append(images, levelImage.Path)
			}
			return images
		}(),
	}
}
