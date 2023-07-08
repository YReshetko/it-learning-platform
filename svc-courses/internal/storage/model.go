package storage

import (
	"github.com/google/uuid"
	"time"
)

type Technology struct {
	ID          *uuid.UUID `gorm:"primaryKey;default:uuid_generate_v4()"`
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (u Technology) TableName() string {
	return "technologies"
}

type Category struct {
	ID           *uuid.UUID `gorm:"primaryKey;default:uuid_generate_v4()"`
	TechnologyID uuid.UUID
	Name         string
	Description  string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (u Category) TableName() string {
	return "categories"
}

type Course struct {
	ID          *uuid.UUID `gorm:"primaryKey;default:uuid_generate_v4()"`
	SeqNo       int
	Name        string
	Description string
	Active      bool
	OwnerID     uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (u Course) TableName() string {
	return "courses"
}

type CourseTopic struct {
	ID        *uuid.UUID `gorm:"primaryKey;default:uuid_generate_v4()"`
	CourseID  uuid.UUID
	Topic     Topic `gorm:"foreignKey:topic_id"`
	SeqNo     int
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u CourseTopic) TableName() string {
	return "courses_topics"
}

type CourseTopicTask struct {
	ID            *uuid.UUID `gorm:"primaryKey;default:uuid_generate_v4()"`
	CourseTopicID uuid.UUID
	Task          Task `gorm:"foreignKey:task_id"`
	Weight        int
	Active        bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (u CourseTopicTask) TableName() string {
	return "tasks_courses_topics"
}

type Tag struct {
	Name      string `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u Tag) TableName() string {
	return "tags"
}

type Task struct {
	ID          *uuid.UUID `gorm:"primaryKey;default:uuid_generate_v4()"`
	SeqNo       int
	Name        string
	Description string
	Active      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Tags        []Tag `gorm:"many2many:tasks_tags;"`
}

func (u Task) TableName() string {
	return "tasks"
}

type Topic struct {
	ID          *uuid.UUID `gorm:"primaryKey;default:uuid_generate_v4()"`
	CategoryID  uuid.UUID
	SeqNo       int
	Name        string
	Description string
	Active      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Tags        []Tag `gorm:"many2many:topics_tags;"`
}

func (u Topic) TableName() string {
	return "topics"
}

type TopicTag struct {
	TopicID   uuid.UUID `gorm:"primaryKey"`
	TagName   string    `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u TopicTag) TableName() string {
	return "topics_tags"
}

type TaskTag struct {
	TaskID    uuid.UUID `gorm:"primaryKey"`
	TagName   string    `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u TaskTag) TableName() string {
	return "tasks_tags"
}
