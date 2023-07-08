package storage

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

func (cs *CourseStorage) CreateCourse(course Course) (Course, error) {
	now := time.Now()
	course.CreatedAt = now
	course.UpdatedAt = now
	rs := cs.db.Create(&course)
	if rs.Error != nil {
		return course, fmt.Errorf("unable to create course record: %w", rs.Error)
	}
	return course, nil
}

func (cs *CourseStorage) GetOwnerCourses(ownerID uuid.UUID) ([]Course, error) {
	var courses []Course
	rs := cs.db.Where("owner_id = ?", ownerID.String()).Order("seq_no").Find(&courses)
	if rs.Error != nil {
		return nil, fmt.Errorf("unable to find owner courses: %w", rs.Error)
	}
	return courses, nil
}

func (cs *CourseStorage) GetCourse(ID uuid.UUID) (Course, error) {
	course := Course{ID: &ID}
	rs := cs.db.First(&course)
	if rs.Error != nil {
		return course, fmt.Errorf("unable to find course: %w", rs.Error)
	}
	return course, nil
}

func (cs *CourseStorage) CreateCourseTopic(courseTopic CourseTopic) (CourseTopic, error) {
	now := time.Now()
	courseTopic.CreatedAt = now
	courseTopic.UpdatedAt = now
	rs := cs.db.Create(&courseTopic)
	if rs.Error != nil {
		return courseTopic, fmt.Errorf("unable to create course topic record: %w", rs.Error)
	}
	return courseTopic, nil
}

func (cs *CourseStorage) GetCourseTopics(courseID uuid.UUID) ([]CourseTopic, error) {
	var courseTopics []CourseTopic
	rs := *cs.db.Preload("Topic").Where("course_id = ?", courseID.String()).Find(&courseTopics)
	if rs.Error != nil {
		return courseTopics, fmt.Errorf("unable to course topics: %w", rs.Error)
	}
	return courseTopics, nil
}

func (cs *CourseStorage) GetCourseTopicsCount(courseID uuid.UUID) (int, error) {
	var count int64
	rs := cs.db.Model(&CourseTopic{ID: &courseID}).Count(&count)
	if rs.Error != nil {
		return 0, fmt.Errorf("unable to calculate course topics count: %w", rs.Error)
	}
	return int(count), nil
}

func (cs *CourseStorage) CreateCourseTopicTask(courseTopicTask CourseTopicTask) (CourseTopicTask, error) {
	now := time.Now()
	courseTopicTask.CreatedAt = now
	courseTopicTask.UpdatedAt = now
	rs := cs.db.Create(&courseTopicTask)
	if rs.Error != nil {
		return courseTopicTask, fmt.Errorf("unable to create course topic task record: %w", rs.Error)
	}
	return courseTopicTask, nil
}

func (cs *CourseStorage) GetCourseTopicTasks(courseTopicID uuid.UUID) ([]CourseTopicTask, error) {
	var courseTopicTasks []CourseTopicTask
	rs := *cs.db.Preload("Task").Where("course_topic_id = ?", courseTopicID.String()).Find(&courseTopicTasks)
	if rs.Error != nil {
		return courseTopicTasks, fmt.Errorf("unable to get course topic tasks: %w", rs.Error)
	}
	return courseTopicTasks, nil
}
