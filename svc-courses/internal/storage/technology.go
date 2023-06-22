package storage

import "fmt"

func (cs *CourseStorage) CreateTechnology(technology Technology) (Technology, error) {
	rs := cs.db.Create(&technology)
	if rs.Error != nil {
		return technology, fmt.Errorf("unable to create technology record: %w", rs.Error)
	}
	return technology, nil
}

func (cs *CourseStorage) GetTechnologies() ([]Technology, error) {
	var technologies []Technology
	rs := cs.db.Find(&technologies)
	if rs.Error != nil {
		return nil, fmt.Errorf("unable to get tecnology records")
	}
	return technologies, nil
}
