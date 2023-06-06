package mapper

import (
	"fmt"
	"github.com/YReshetko/it-academy-cources/svc-users/internal/storage"
	"github.com/YReshetko/it-academy-cources/svc-users/pb/users"
	"github.com/google/uuid"
)

func PbToDBUser(user users.User) storage.User {
	var ID *uuid.UUID
	if user.Id != "" {
		pbID, err := uuid.Parse(user.Id)
		if err != nil {
			fmt.Println("User ID: unable to parse UUID from string:", user.Id)
			fmt.Println(err)
		} else {
			ID = &pbID
		}
	}

	externalID, err := uuid.Parse(user.ExternalId)
	if err != nil {
		fmt.Println("ExternalId: unable to parse UUID from string:", user.ExternalId)
		fmt.Println(err)
	}
	return storage.User{
		ID:         ID,
		ExternalID: externalID,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Email:      user.Email,
		Active:     user.Active,
	}
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
