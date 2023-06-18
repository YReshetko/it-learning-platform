package handlers

import (
	"fmt"
	"github.com/YReshetko/it-learning-platform/api-app/internal/clients"
	"github.com/YReshetko/it-learning-platform/svc-auth/pb/auth"
	"github.com/sirupsen/logrus"
	rest "net/http"

	"github.com/YReshetko/it-learning-platform/api-app/internal/http"
	"github.com/YReshetko/it-learning-platform/api-app/internal/http/models"
)

/*
Registration the handlers to interact with svc-auth
@Constructor
*/
type Registration struct {
	client clients.AuthClient
	logger *logrus.Entry
}

func (r *Registration) CreateUser(context http.Context, user models.AuthUser) (models.AuthResponse, http.Status) {
	logger := r.logger.WithField("method", "CreateUser").WithField("user", user)
	err := r.createUser(context, user)
	if err != nil {
		logger.WithError(err).Error("unable to create user")
		return models.AuthResponse{}, http.Status{
			Error:      err,
			StatusCode: rest.StatusInternalServerError,
			Message:    "unable to create user",
		}
	}

	return models.AuthResponse{}, http.Status{StatusCode: rest.StatusOK}
}

func (r *Registration) CreateUsers(context http.Context, users models.AuthUsers) (models.AuthResponse, http.Status) {
	logger := r.logger.WithField("method", "CreateUsers").WithField("user", users)
	for i, user := range users.Users {
		err := r.createUser(context, user)
		if err != nil {
			logger.WithField("user", user).WithError(err).Error("unable to create user")
			return models.AuthResponse{}, http.Status{
				Error:      err,
				StatusCode: rest.StatusInternalServerError,
				Message:    fmt.Sprintf("unable to create %d user", i+1),
			}
		}
	}

	return models.AuthResponse{}, http.Status{StatusCode: rest.StatusOK}
}

func (r *Registration) createUser(context http.Context, user models.AuthUser) error {
	_, err := r.client.CreateUser(context.GinCtx.Request.Context(), &auth.CreateAuthUserRequest{
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
