package model

import (
	"github.com/YReshetko/it-learning-platform/lib-app/pkg/model"
	"github.com/google/uuid"
)

type User struct {
	ID         *uuid.UUID
	KeycloakID *uuid.UUID
	Login      string
	FirstName  string
	LastName   string
	Email      string
	Active     bool
	Roles      []model.Role
}
