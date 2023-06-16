package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Context struct {
	GinCtx      *gin.Context
	AccessToken string
	UserID      uuid.UUID
}

func (c *Context) Context() context.Context {
	if c == nil || c.GinCtx == nil || c.GinCtx.Request == nil || c.GinCtx.Request.Context() == nil {
		return context.Background()
	}
	return c.GinCtx.Request.Context()
}

type Status struct {
	Error      error
	StatusCode int
	Message    string
}

type HandlerFunc[Rq any, Rs any] func(Context, Rq) (Rs, Status)
