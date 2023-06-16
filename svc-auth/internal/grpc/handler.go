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

func (h *Handler) AccessTokenExchange(ctx context.Context, rq *auth.AccessTokenExchangeRequest) (*auth.AccessTokenExchangeResponse, error) {
	ok, err := h.keycloakClient.ValidateAccessToken(ctx, rq.AccessToken.GetToken())
	if err != nil {
		fmt.Println("unable to validate access token", err.Error())
		return &auth.AccessTokenExchangeResponse{}, status.Error(codes.Internal, "unable to validate access token")
	}
	if !ok {
		fmt.Println("invalid token: ", rq.AccessToken)
		return &auth.AccessTokenExchangeResponse{}, status.Error(codes.Unauthenticated, "invalid access token")
	}

	keycloakUserId, roles, err := h.keycloakClient.GetUserIDAndRoles(ctx, rq.AccessToken.GetToken())
	if err != nil {
		fmt.Println("unable to get user ID and role from keycloak", err.Error())
		return &auth.AccessTokenExchangeResponse{}, status.Error(codes.Unauthenticated, "unable to authorize user")
	}

	user, err := h.userClient.FindUserByExternalID(ctx, &users.FindUserByExternalIDRequest{ExternalId: keycloakUserId.String()})
	if err != nil {
		fmt.Println("unable to get user ID by external user ID", err.Error())
		return &auth.AccessTokenExchangeResponse{}, status.Error(codes.Unauthenticated, "unable to authorize user")
	}

	var userRoles []auth.UserRole
	for _, role := range roles {
		r, ok := auth.UserRole_value[string(role)]
		if !ok {
			fmt.Printf("WARNING! model role %s does not match any pb roles \n", role)
		}
		userRoles = append(userRoles, auth.UserRole(r))
	}

	return &auth.AccessTokenExchangeResponse{
		UserInfo: &auth.UserInfo{
			Id:    user.User.Id,
			Roles: userRoles,
		},
	}, nil
}
