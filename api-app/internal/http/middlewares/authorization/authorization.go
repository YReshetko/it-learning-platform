package authorization

import (
	rest "github.com/YReshetko/it-academy-cources/api-app/internal/http"
	"github.com/google/uuid"
	"net/http"
	"strings"
)

type Service struct {
}

func (s Service) verifyToken(token string) (bool, error) {
	return true, nil
}

func (s Service) getUserRoles(token string) (uuid.UUID, []Role, error) {
	return uuid.UUID{}, []Role{ADMIN}, nil
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

func Authorize[Rq any, Rs any](fn rest.HandlerFunc[Rq, Rs], service Service, roles []Role) rest.HandlerFunc[Rq, Rs] {
	return func(context rest.Context, request Rq) (Rs, rest.Status) {
		var rs Rs
		ok, err := service.verifyToken(context.AccessToken)
		if err != nil {
			return rs, rest.Status{
				Error:      err,
				StatusCode: http.StatusInternalServerError,
				Message:    "unable to authorize user",
			}
		}
		if !ok {
			return rs, rest.Status{
				Error:      nil,
				StatusCode: http.StatusUnauthorized,
				Message:    "invalid token",
			}
		}

		userID, userRoles, err := service.getUserRoles(context.AccessToken)
		if err != nil {
			return rs, rest.Status{
				Error:      err,
				StatusCode: http.StatusInternalServerError,
				Message:    "unable to verify user roles",
			}
		}

		if !matchRoles(userRoles, roles) {
			return rs, rest.Status{
				Error:      nil,
				StatusCode: http.StatusForbidden,
				Message:    "access denied",
			}
		}

		context.UserID = userID
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
