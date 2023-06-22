package models

import "github.com/google/uuid"

type Technology struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

type Technologies struct {
	Technologies []Technology `json:"technologies"`
}
