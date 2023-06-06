package storage

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

/*
UserStorage the user storage
@Constructor
*/
type UserStorage struct {
	db *gorm.DB
}

func (us UserStorage) Create(user User) (User, error) {
	rs := us.db.Create(&user)
	if rs.Error != nil {
		return User{}, fmt.Errorf("unable to create user: %w", rs.Error)
	}
	return user, nil
}

func (us UserStorage) FindByID(ID uuid.UUID) (User, error) {
	user := User{ID: &ID}
	rs := us.db.First(&user)
	if rs.Error != nil {
		return User{}, fmt.Errorf("unable to find user %s: %w", ID.String(), rs.Error)
	}
	return user, nil
}

func (us UserStorage) FindByExternalID(ID uuid.UUID) (User, error) {
	user := User{ID: &ID}
	rs := us.db.First(&user, "external_id = ?", ID.String())
	if rs.Error != nil {
		return User{}, fmt.Errorf("unable to find user by external ID %s: %w", ID.String(), rs.Error)
	}
	return user, nil
}

func (us UserStorage) FindByIDs(IDs []uuid.UUID) ([]User, error) {
	var users []User
	rs := us.db.Find(&users, IDs)
	if rs.Error != nil {
		return nil, fmt.Errorf("unable to find user %s: %w", IDs, rs.Error)
	}
	return users, nil
}

func (us UserStorage) Update(user User) error {
	user.UpdatedAt = time.Now()
	rs := us.db.Save(user)
	if rs.Error != nil {
		return fmt.Errorf("unable to update user %s: %w", user.ID.String(), rs.Error)
	}
	return nil
}
