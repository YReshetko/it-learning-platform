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
