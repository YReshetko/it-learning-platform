package http

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Context struct {
	GinCtx      *gin.Context
	AccessToken string
	UserID      uuid.UUID
}

type Status struct {
	Error      error
	StatusCode int
	Message    string
}

type HandlerFunc[Rq any, Rs any] func(Context, Rq) (Rs, Status)
