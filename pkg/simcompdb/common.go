package simcompdb

import "gorm.io/gorm"

// find item in database, if error occur returns error, if value not found return value and error as nil
func findById[T any](id string, db *gorm.DB) (*T, error) {
	var item T
	result := db.Find(&item, id)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, nil
	}
	return &item, nil
}
