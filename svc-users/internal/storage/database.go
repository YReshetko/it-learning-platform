package storage

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/YReshetko/it-academy-cources/svc-users/internal/config"
)

func DatabaseConnection(cfg config.Database) (*gorm.DB, error) {
	dialector := postgres.New(postgres.Config{
		DSN: cfg.CreatePostgresDSN(),
	})
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("unable to connect to DB: %w", err)
	}
	return db, nil
}
