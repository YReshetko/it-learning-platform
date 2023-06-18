package mapper

import (
	"fmt"
	"github.com/YReshetko/it-learning-platform/svc-users/internal/storage"
	"github.com/YReshetko/it-learning-platform/svc-users/pb/users"
	"github.com/google/uuid"
)

func PbToDBUser(user users.User) (storage.User, error) {
	var ID *uuid.UUID
	if user.Id != "" {
		pbID, err := uuid.Parse(user.Id)
		if err != nil {
			return storage.User{}, fmt.Errorf("unable to parse user ID: %w", err)
		} else {
			ID = &pbID
		}
	}

	externalID, err := uuid.Parse(user.ExternalId)
	if err != nil {
		return storage.User{}, fmt.Errorf("unable to parse user ExternalID: %W", err)
	}
	return storage.User{
		ID:         ID,
		ExternalID: externalID,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Email:      user.Email,
		Active:     user.Active,
	}, nil
}

func DBToPbUser(user storage.User) users.User {
	ID := ""
	if user.ID != nil {
		ID = user.ID.String()
	}
	return users.User{
		Id:         ID,
		ExternalId: user.ExternalID.String(),
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Active:     user.Active,
	}
}
