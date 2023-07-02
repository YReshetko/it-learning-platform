package storage

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

func (cs *CourseStorage) CreateTag(tag Tag) (Tag, error) {
	rs := cs.db.Create(&tag)
	if rs.Error != nil {
		return tag, fmt.Errorf("unable to create tag record: %w", rs.Error)
	}
	return tag, nil
}

func (cs *CourseStorage) GetTags() ([]Tag, error) {
	var tags []Tag
	rs := cs.db.Order("name").Find(&tags)
	if rs.Error != nil {
		return nil, fmt.Errorf("unable to get tag records")
	}
	return tags, nil
}

func (cs *CourseStorage) SearchTags(search string) ([]Tag, error) {
	var tags []Tag
	rs := cs.db.Where("name LIKE ?", "%"+search+"%").Order("name").Find(&tags)
	if rs.Error != nil {
		return nil, fmt.Errorf("unable to find tag records")
	}
	return tags, nil
}

func (cs *CourseStorage) RemoveTag(tag Tag) (Tag, error) {
	rs := cs.db.Delete(&tag)
	if rs.Error != nil {
		return tag, fmt.Errorf("unable to delete tag record: %w", rs.Error)
	}
	return tag, nil
}

func (cs *CourseStorage) CreateTopicTag(topicID uuid.UUID, tagName string) (Topic, error) {
	now := time.Now()
	topicTag := TopicTag{
		TopicID:   topicID,
		TagName:   tagName,
		CreatedAt: now,
		UpdatedAt: now,
	}

	rs := cs.db.Create(&topicTag)
	if rs.Error != nil {
		return Topic{}, fmt.Errorf("unable to create topic's tag record: %w", rs.Error)
	}

	return cs.GetTopic(topicID)
}

func (cs *CourseStorage) RemoveTopicTag(topicID uuid.UUID, tagName string) (Topic, error) {
	topicTag := TopicTag{
		TopicID: topicID,
		TagName: tagName,
	}
	rs := cs.db.Delete(&topicTag)
	if rs.Error != nil {
		return Topic{}, fmt.Errorf("unable to delete topic's tag record: %w", rs.Error)
	}
	return cs.GetTopic(topicID)
}

func (cs *CourseStorage) CreateTaskTag(taskID uuid.UUID, tagName string) (Task, error) {
	now := time.Now()
	taskTag := TaskTag{
		TaskID:    taskID,
		TagName:   tagName,
		CreatedAt: now,
		UpdatedAt: now,
	}

	rs := cs.db.Create(&taskTag)
	if rs.Error != nil {
		return Task{}, fmt.Errorf("unable to create task's tag record: %w", rs.Error)
	}

	return cs.GetTask(taskID)
}

func (cs *CourseStorage) RemoveTaskTag(taskID uuid.UUID, tagName string) (Task, error) {
	taskTag := TaskTag{
		TaskID:  taskID,
		TagName: tagName,
	}
	rs := cs.db.Delete(&taskTag)
	if rs.Error != nil {
		return Task{}, fmt.Errorf("unable to delete task's tag record: %w", rs.Error)
	}
	return cs.GetTask(taskID)
}
