package grpc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/YReshetko/it-learning-platform/lib-app/pkg/model"
	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
)

const (
	userID    = "user_id"
	userRoles = "user_roles"
)

func WithUserIDContext(ctx context.Context, userUUID uuid.UUID) context.Context {
	md := metadataFromOutgoingContext(ctx)
	md.Append(userID, userUUID.String())
	return metadata.NewOutgoingContext(ctx, md)
}

func WithUserRolesContext(ctx context.Context, roles []model.Role) context.Context {
	rolesString, _ := json.Marshal(roles)
	md := metadataFromOutgoingContext(ctx)
	md.Append(userRoles, string(rolesString))
	return metadata.NewOutgoingContext(ctx, md)
}

func metadataFromOutgoingContext(ctx context.Context) metadata.MD {
	md, ok := metadata.FromOutgoingContext(ctx)
	if ok {
		return md
	}
	return metadata.MD{}
}

func GetUserIDFromContext(ctx context.Context) (uuid.UUID, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return uuid.New(), errors.New("metadata not found in GRPC context")
	}
	data := md.Get(userID)
	if len(data) != 1 {
		return uuid.New(), errors.New(fmt.Sprintf("unable to find user id in [%+v]", data))
	}

	return uuid.Parse(data[0])
}

func GetUserRolesFromContext(ctx context.Context) ([]model.Role, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("metadata not found in GRPC context")
	}
	data := md.Get(userRoles)
	if len(data) != 1 {
		return nil, errors.New(fmt.Sprintf("unable to find user roles in [%+v]", data))
	}

	var roles []model.Role
	err := json.Unmarshal([]byte(data[0]), &roles)

	return roles, err
}
