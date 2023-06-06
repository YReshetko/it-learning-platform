package grpc

import (
	"context"
	"fmt"
	"github.com/YReshetko/it-academy-cources/svc-auth/internal/clients"
	"github.com/YReshetko/it-academy-cources/svc-auth/internal/model"
	"github.com/YReshetko/it-academy-cources/svc-auth/pb/auth"
	"github.com/YReshetko/it-academy-cources/svc-users/pb/users"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

/*
Handler the GRPC requests handler
@Constructor
*/
type Handler struct {
	auth.UnimplementedAuthServiceServer // @Exclude
	userClient                          *clients.UsersClient
	keycloakClient                      *clients.KeycloakClient
}

func (h *Handler) CreateUser(ctx context.Context, rq *auth.CreateAuthUserRequest) (*auth.CreateAuthUserResponse, error) {
	keycloakUserID, err := h.keycloakClient.CreateUser(ctx, model.User{
		Login:     rq.GetUser().GetLogin(),
		FirstName: rq.GetUser().GetFirstName(),
		LastName:  rq.GetUser().GetLastName(),
		Email:     rq.GetUser().GetEmail(),
		Active:    true,
		Roles:     mapRoles(rq.GetUser().GetRoles()),
	})
	if err != nil {
		fmt.Println("unable to create user in keycloak:", err)
		return &auth.CreateAuthUserResponse{}, status.Error(codes.Internal, "unable to register keycloak user")
	}

	fmt.Println("Keycloak user ID:", keycloakUserID)
	_, err = h.userClient.CreateUser(ctx, &users.CreateUserRequest{
		User: &users.User{
			ExternalId: keycloakUserID,
			FirstName:  rq.GetUser().GetFirstName(),
			LastName:   rq.GetUser().GetLastName(),
			Email:      rq.GetUser().GetEmail(),
			Active:     true,
		}})
	if err != nil {
		fmt.Println("unable to create user in svc-users", err)
		return &auth.CreateAuthUserResponse{}, status.Error(codes.Internal, "unable to create user in svc-users")
	}

	return &auth.CreateAuthUserResponse{}, nil
}

func mapRoles(roles []auth.UserRole) model.Roles {
	out := make(model.Roles, len(roles))
	for i, role := range roles {
		out[i] = model.Role(auth.UserRole_name[int32(role)])
	}
	return out
}
