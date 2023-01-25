package repository

import (
	"errors"
	"log"
	"os"

	"github.com/harshrastogiexe/simcompanies/types"

	"gorm.io/gorm"
)

type IResourceRepository interface {
	SaveResource(resource *types.Resource) error
}

type DBResourceRepository struct {
	DB *gorm.DB
}

var (
	ErrFailed = errors.New("query failed")
	errorLog  = log.New(os.Stdout, "[ERROR]: ", log.Flags())
)

const (
	RESOURCE_TABLE_NAME = "Resource"
)

func (r *DBResourceRepository) SaveResource(resource *types.Resource) error {
	if err := r.DB.Table(RESOURCE_TABLE_NAME).Create(resource).Error; err != nil {
		errorLog.Println(err.Error())
		return ErrFailed
	}
	return nil
}

func (r *DBResourceRepository) ListResource() []types.Resource {
	var resources []types.Resource
	r.DB.Table(RESOURCE_TABLE_NAME).Find(&resources)
	return resources
}

func (r *DBResourceRepository) GetResourceById(id uint) *types.Resource {
	var resource = new(types.Resource)
	rows := r.DB.Table(RESOURCE_TABLE_NAME).Find(&resource, id).RowsAffected
	if rows == 0 {
		return nil
	}
	return resource
}
