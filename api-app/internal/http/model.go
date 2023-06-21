package http

import (
	"context"
	"github.com/YReshetko/it-learning-platform/lib-app/pkg/grpc"
	"github.com/YReshetko/it-learning-platform/lib-app/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Context struct {
	GinCtx      *gin.Context
	AccessToken string
	UserID      uuid.UUID
	UserRoles   []model.Role
}

func (c *Context) Context() context.Context {
	if c == nil || c.GinCtx == nil || c.GinCtx.Request == nil || c.GinCtx.Request.Context() == nil {
		return c.contextWithUserMetadata(context.Background())
	}
	return c.contextWithUserMetadata(c.GinCtx.Request.Context())
}

func (c *Context) contextWithUserMetadata(ctx context.Context) context.Context {
	nc := grpc.WithUserIDContext(ctx, c.UserID)
	return grpc.WithUserRolesContext(nc, c.UserRoles)
}

type Status struct {
	Error      error
	StatusCode int
	Message    string
}

type HandlerFunc[Rq any, Rs any] func(Context, Rq) (Rs, Status)
