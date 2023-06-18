package grpc

import (
	"context"
	"github.com/YReshetko/it-learning-platform/svc-auth/internal/clients"
	"github.com/YReshetko/it-learning-platform/svc-auth/internal/model"
	"github.com/YReshetko/it-learning-platform/svc-auth/pb/auth"
	"github.com/YReshetko/it-learning-platform/svc-users/pb/users"
	"github.com/sirupsen/logrus"
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
	logger                              *logrus.Entry
}

func (h *Handler) CreateUser(ctx context.Context, rq *auth.CreateAuthUserRequest) (*auth.CreateAuthUserResponse, error) {
	logger := h.logger.WithField("method", "CreateUser").WithField("request", rq)
	keycloakUserID, err := h.keycloakClient.CreateUser(ctx, model.User{
		Login:     rq.GetUser().GetLogin(),
		FirstName: rq.GetUser().GetFirstName(),
		LastName:  rq.GetUser().GetLastName(),
		Email:     rq.GetUser().GetEmail(),
		Active:    true,
		Roles:     mapRoles(rq.GetUser().GetRoles()),
	})
	if err != nil {
		logger.WithError(err).Error("unable to create user in keycloak")
		return &auth.CreateAuthUserResponse{}, status.Error(codes.Internal, "unable to register keycloak user")
	}

	logger = logger.WithField("keycloak_user_id", keycloakUserID)
	_, err = h.userClient.CreateUser(ctx, &users.CreateUserRequest{
		User: &users.User{
			ExternalId: keycloakUserID,
			FirstName:  rq.GetUser().GetFirstName(),
			LastName:   rq.GetUser().GetLastName(),
			Email:      rq.GetUser().GetEmail(),
			Active:     true,
		}})
	if err != nil {
		logger.WithError(err).Error("unable to create user in svc-users")
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
	logger := h.logger.WithField("method", "AccessTokenExchange").WithField("request", rq)
	ok, err := h.keycloakClient.ValidateAccessToken(ctx, rq.AccessToken.GetToken())
	if err != nil {
		logger.WithError(err).Error("unable to validate access token")
		return &auth.AccessTokenExchangeResponse{}, status.Error(codes.Internal, "unable to validate access token")
	}
	if !ok {
		logger.Error("invalid token")
		return &auth.AccessTokenExchangeResponse{}, status.Error(codes.Unauthenticated, "invalid access token")
	}

	keycloakUserId, roles, err := h.keycloakClient.GetUserIDAndRoles(ctx, rq.AccessToken.GetToken())
	if err != nil {
		logger.WithError(err).Error("unable to get user ID and role from keycloak")
		return &auth.AccessTokenExchangeResponse{}, status.Error(codes.Unauthenticated, "unable to authorize user")
	}

	user, err := h.userClient.FindUserByExternalID(ctx, &users.FindUserByExternalIDRequest{ExternalId: keycloakUserId.String()})
	if err != nil {
		logger.WithError(err).Error("unable to get user ID by external user ID")
		return &auth.AccessTokenExchangeResponse{}, status.Error(codes.Unauthenticated, "unable to authorize user")
	}

	var userRoles []auth.UserRole
	for _, role := range roles {
		r, ok := auth.UserRole_value[string(role)]
		if !ok {
			logger.WithField("model_role", role).
				WithField("proto_roles", auth.UserRole_value).
				Warn("Model role does not match any proto roles")
			continue
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
