package storage

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID         *uuid.UUID `gorm:"primaryKey;default:uuid_generate_v4()"`
	ExternalID uuid.UUID
	FirstName  string
	LastName   string
	Email      string
	Active     bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (u User) TableName() string {
	return "users"
}
