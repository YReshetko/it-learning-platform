package routes

import (
	"github.com/YReshetko/it-learning-platform/api-app/internal/http"
	auth "github.com/YReshetko/it-learning-platform/api-app/internal/http/middlewares/authorization"

	"github.com/YReshetko/it-learning-platform/api-app/internal/http/handlers"
	mgin "github.com/YReshetko/it-learning-platform/api-app/internal/http/middlewares/gin"
	"github.com/gin-gonic/gin"
)

/*
Router REST API router
@Optional
*/
type Router struct {
	authHandler handlers.Auth

	authService auth.Service
}

func (r Router) Init(engine *gin.Engine) {
	engine.POST("/api/v1/auth/user", protected(r.authHandler.CreateUser, r.authService, []auth.Role{auth.ADMIN}))
	engine.POST("/api/v1/auth/users", protected(r.authHandler.CreateUsers, r.authService, []auth.Role{auth.ADMIN}))

}

func protected[Rq any, Rs any](fn http.HandlerFunc[Rq, Rs], service auth.Service, roles []auth.Role) func(*gin.Context) {
	fn = auth.Authorize(fn, service, roles)
	fn = auth.Authenticate(fn)
	return mgin.Wrap(fn)
}
