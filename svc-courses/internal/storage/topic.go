package storage

import (
	"fmt"
	"github.com/google/uuid"
)

func (cs *CourseStorage) CreateTopic(topic Topic) (Topic, error) {
	rs := cs.db.Create(&topic)
	if rs.Error != nil {
		return topic, fmt.Errorf("unable to create topic record: %w", rs.Error)
	}
	return topic, nil
}

func (cs *CourseStorage) GetTopics(categoryID uuid.UUID) ([]Topic, error) {
	var topics []Topic
	rs := cs.db.Where("category_id = ?", categoryID.String()).Find(&topics)
	if rs.Error != nil {
		return nil, fmt.Errorf("unable to get topic records")
	}
	return topics, nil
}
