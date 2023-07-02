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

type Tag struct {
	Name string `json:"name"`
}

type Tags struct {
	Tags []Tag `json:"tags"`
}
