package storage

import (
	"fmt"
	"github.com/google/uuid"
)

func (cs *CourseStorage) CreateCategory(category Category) (Category, error) {
	rs := cs.db.Create(&category)
	if rs.Error != nil {
		return category, fmt.Errorf("unable to create category record: %w", rs.Error)
	}
	return category, nil
}

func (cs *CourseStorage) GetCategories(technologyID uuid.UUID) ([]Category, error) {
	var categories []Category
	rs := cs.db.Where("technology_id = ?", technologyID.String()).Find(&categories)
	if rs.Error != nil {
		return nil, fmt.Errorf("unable to get category records")
	}
	return categories, nil
}
