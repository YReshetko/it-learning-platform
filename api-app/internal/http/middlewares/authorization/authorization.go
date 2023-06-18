package authorization

import (
	"fmt"
	"github.com/YReshetko/it-learning-platform/api-app/internal/clients"
	rest "github.com/YReshetko/it-learning-platform/api-app/internal/http"
	"github.com/YReshetko/it-learning-platform/svc-auth/pb/auth"
	"github.com/google/uuid"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"strings"
)

// Service @Constructor
type Service struct {
	client clients.AuthClient
}

func (s *Service) getUserRoles(ctx context.Context, token string) (uuid.UUID, []Role, error) {
	fmt.Printf("USER INFO TOKEN: %+v\n", token)
	userInfo, err := s.client.AccessTokenExchange(ctx, &auth.AccessTokenExchangeRequest{AccessToken: &auth.AccessToken{Token: token}})
	if err != nil {
		return uuid.UUID{}, nil, err
	}
	fmt.Printf("USER INFO FROM TOKEN: %+v\n", userInfo)
	userId, err := uuid.Parse(userInfo.GetUserInfo().GetId())
	if err != nil {
		return uuid.UUID{}, nil, err
	}

	userRoles := userInfo.GetUserInfo().GetRoles()
	roles := make([]Role, len(userRoles))
	for i, role := range userRoles {
		roles[i] = Role(auth.UserRole_name[int32(role)])
	}

	return userId, roles, nil
}

func Authenticate[Rq any, Rs any](fn rest.HandlerFunc[Rq, Rs]) rest.HandlerFunc[Rq, Rs] {
	return func(ctx rest.Context, request Rq) (Rs, rest.Status) {
		ctx.GinCtx.GetHeader("Authorization")
		authHeader := ctx.GinCtx.GetHeader("Authorization")
		var rs Rs
		if len(authHeader) < 1 {
			return rs, rest.Status{
				Error:      nil,
				StatusCode: http.StatusUnauthorized,
				Message:    "unauthorized",
			}
		}
		ctx.AccessToken = strings.Split(authHeader, " ")[1]
		return fn(ctx, request)
	}
}

func Authorize[Rq any, Rs any](fn rest.HandlerFunc[Rq, Rs], service *Service, roles []Role) rest.HandlerFunc[Rq, Rs] {
	return func(context rest.Context, request Rq) (Rs, rest.Status) {
		var rs Rs
		userId, userRoles, err := service.getUserRoles(context.Context(), context.AccessToken)
		if err != nil {
			s, ok := status.FromError(err)
			if !ok {
				return rs, rest.Status{
					Error:      err,
					StatusCode: http.StatusInternalServerError,
					Message:    "unable to verify user role",
				}
			}
			switch s.Code() {
			case codes.Unauthenticated:
				return rs, rest.Status{
					Error:      err,
					StatusCode: http.StatusUnauthorized,
					Message:    "invalid token",
				}
			case codes.Internal:
				return rs, rest.Status{
					Error:      err,
					StatusCode: http.StatusInternalServerError,
					Message:    "unable to authorize user",
				}
			default:
				return rs, rest.Status{
					Error:      err,
					StatusCode: http.StatusInternalServerError,
					Message:    "unknown error",
				}
			}
		}

		if !matchRoles(userRoles, roles) {
			return rs, rest.Status{
				Error:      nil,
				StatusCode: http.StatusForbidden,
				Message:    "access denied",
			}
		}

		context.UserID = userId
		return fn(context, request)
	}
}

func matchRoles(r1, r2 []Role) bool {
	for _, role1 := range r1 {
		for _, role2 := range r2 {
			if role1 == role2 {
				return true
			}
		}
	}
	return false
}
