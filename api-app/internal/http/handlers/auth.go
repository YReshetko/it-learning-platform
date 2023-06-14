package handlers

import (
	"fmt"
	"github.com/YReshetko/it-academy-cources/api-app/internal/clients"
	"github.com/YReshetko/it-academy-cources/svc-auth/pb/auth"
	rest "net/http"

	"github.com/YReshetko/it-academy-cources/api-app/internal/http"
	"github.com/YReshetko/it-academy-cources/api-app/internal/http/models"
)

/*
Auth the handlers to interact with svc-auth
@Constructor
*/
type Auth struct {
	client clients.AuthClient
}

func (a *Auth) CreateUser(context http.Context, user models.AuthUser) (models.AuthResponse, http.Status) {
	err := a.createUser(context, user)
	if err != nil {
		return models.AuthResponse{}, http.Status{
			Error:      err,
			StatusCode: rest.StatusInternalServerError,
			Message:    "unable to create user",
		}
	}

	return models.AuthResponse{}, http.Status{StatusCode: rest.StatusOK}
}

func (a *Auth) CreateUsers(context http.Context, users models.AuthUsers) (models.AuthResponse, http.Status) {
	for i, user := range users.Users {
		err := a.createUser(context, user)
		if err != nil {
			return models.AuthResponse{}, http.Status{
				Error:      err,
				StatusCode: rest.StatusInternalServerError,
				Message:    fmt.Sprintf("unable to create %d user", i+1),
			}
		}
	}

	return models.AuthResponse{}, http.Status{StatusCode: rest.StatusOK}
}

func (a *Auth) createUser(context http.Context, user models.AuthUser) error {
	_, err := a.client.CreateUser(context.GinCtx.Request.Context(), &auth.CreateAuthUserRequest{
		User: &auth.AuthUser{
			Login:     user.Login,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Roles:     user.Roles.ToAuthProtoRoles(),
		},
	})

	if err != nil {
		return fmt.Errorf("unable to create user")
	}
	return nil
}
