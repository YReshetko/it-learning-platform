// Code generated by Constructor annotation processor. DO NOT EDIT.
// versions:
//		go: go1.20.4
//		go-annotation: 0.1.0
//		Constructor: 1.0.0

package clients

import (
	config "github.com/YReshetko/it-learning-platform/svc-auth/internal/config"
)

func NewKeycloakClient(cfg config.KeycloakClient) KeycloakClient {
	returnValue := KeycloakClient{
		cfg: cfg,
	}
	returnValue.postConstruct()

	return returnValue
}

func NewUsersClient(cfg config.UsersClient) UsersClient {
	returnValue := UsersClient{
		cfg: cfg,
	}
	returnValue.postConstruct()

	return returnValue
}
