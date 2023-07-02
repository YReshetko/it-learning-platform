package storage

import (
	"fmt"
	"github.com/google/uuid"
)

func (cs *CourseStorage) CreateTask(task Task) (Task, error) {
	rs := cs.db.Create(&task)
	if rs.Error != nil {
		return task, fmt.Errorf("unable to create task record: %w", rs.Error)
	}
	return task, nil
}

func (cs *CourseStorage) GetTasks() ([]Task, error) {
	var tasks []Task
	rs := cs.db.Find(&tasks)
	if rs.Error != nil {
		return nil, fmt.Errorf("unable to get tasks records: %w", rs.Error)
	}
	return tasks, nil
}

func (cs *CourseStorage) GetTask(taskID uuid.UUID) (Task, error) {
	task := Task{
		ID: &taskID,
	}
	rs := cs.db.Preload("Tags").First(&task)
	if rs.Error != nil {
		return Task{}, fmt.Errorf("unable to get task records: %w", rs.Error)
	}
	return task, nil
}

func (cs *CourseStorage) FindTasks(tags []Tag) ([]Task, error) {
	names := make([]string, len(tags))
	for i, tag := range tags {
		names[i] = tag.Name
	}

	var tasks []Task
	rs := cs.db.Preload("Tags").Where("Tags.Name IN (?)", names).Find(&tasks)
	if rs.Error != nil {
		return nil, fmt.Errorf("unable to find tasks records: %w", rs.Error)
	}
	return tasks, nil
}
