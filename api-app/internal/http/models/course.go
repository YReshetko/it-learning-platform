package models

import (
	"github.com/google/uuid"
)

type Technology struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

type Technologies struct {
	Technologies []Technology `json:"technologies"`
}

type Category struct {
	ID           uuid.UUID `json:"id"`
	TechnologyID uuid.UUID `json:"technology_id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
}

type Categories struct {
	Categories []Category `json:"categories"`
}

type Topic struct {
	ID          uuid.UUID `json:"id"`
	CategoryID  uuid.UUID `json:"category_id"`
	SeqNo       int       `json:"seq_no"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Active      bool      `json:"active"`
	Tags        []Tag     `json:"tags"`
}

type Topics struct {
	Topics []Topic `json:"topics"`
}

type Task struct {
	ID          uuid.UUID `json:"id"`
	SeqNo       int       `json:"seq_no"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Active      bool      `json:"active"`
	Tags        []Tag     `json:"tags"`
}

type Tasks struct {
	Tasks []Task `json:"tasks"`
}

type Tag struct {
	Name string `json:"name"`
}

type Tags struct {
	Tags []Tag `json:"tags"`
}

type Course struct {
	ID          uuid.UUID     `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	SeqNo       int           `json:"seq_no"`
	Active      bool          `json:"active"`
	Topics      []CourseTopic `json:"topics"`
}

type Courses struct {
	Courses []Course `json:"courses"`
}

type CourseTopic struct {
	ID          uuid.UUID `json:"id"`
	SeqNo       int       `json:"seq_no"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Active      bool      `json:"active"`
	Tags        []Tag     `json:"tags"`
}

type FullTechnologyList struct {
	Technologies []TechnologyWithCategory `json:"technologies"`
}

type TechnologyWithCategory struct {
	Technology Technology           `json:"technology"`
	Categories []CategoryWithTopics `json:"categories"`
}

type CategoryWithTopics struct {
	Category Category `json:"category"`
	Topics   []Topic  `json:"topics"`
}
