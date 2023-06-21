package storage

import "gorm.io/gorm"

/*
CourseStorage the courses storage
@Constructor
*/
type CourseStorage struct {
	db *gorm.DB
}
