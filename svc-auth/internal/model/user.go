package model

import "github.com/google/uuid"

type User struct {
	ID         *uuid.UUID
	KeycloakID *uuid.UUID
	Login      string
	FirstName  string
	LastName   string
	Email      string
	Active     bool
	Roles      Roles
}

// TODO duplication in api-app
type Role string
type Roles []Role

const (
	ADMIN   Role = "ADMIN"
	MANAGER Role = "MANAGER"
	TEACHER Role = "TEACHER"
	STUDENT Role = "STUDENT"
)

var AllRoles = []Role{ADMIN, MANAGER, TEACHER, STUDENT}

func (r Roles) ToStringPtr() []string {
	l := len(r)
	if l == 0 {
		return nil
	}
	out := make([]string, l)
	for i, role := range r {
		out[i] = string(role)
	}
	return out
}
