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

	errorLog = log.New(os.Stdout, "[ERROR]: ", log.Flags())
)

func (r *DBResourceRepository) SaveResource(resource *types.Resource) error {
	if err := r.DB.Table("Resource").Create(resource).Error; err != nil {
		errorLog.Println(err.Error())
		return ErrFailed
	}
	return nil
}
