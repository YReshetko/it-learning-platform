package routes

import (
	"github.com/YReshetko/it-learning-platform/api-app/internal/http"
	auth "github.com/YReshetko/it-learning-platform/api-app/internal/http/middlewares/authorization"
	"github.com/sirupsen/logrus"

	"github.com/YReshetko/it-learning-platform/api-app/internal/http/handlers"
	mgin "github.com/YReshetko/it-learning-platform/api-app/internal/http/middlewares/gin"
	"github.com/gin-gonic/gin"
)

/*
Router REST API router
@Optional
*/
type Router struct {
	registration handlers.Registration
	self         handlers.Self

	services *RouterServices
}

/*
RouterServices contains all services requires for middlewares
@Optional
*/
type RouterServices struct {
	authService *auth.Service
	logger      *logrus.Entry
	redirectURL string
}

func (r Router) Init(engine *gin.Engine) {
	engine.POST("/api/v1/registration/user", protected(r.registration.CreateUser, r.services, []auth.Role{auth.ADMIN}))
	engine.POST("/api/v1/registration/users", protected(r.registration.CreateUsers, r.services, []auth.Role{auth.ADMIN}))
	engine.GET("/api/v1/self", protected(r.self.GetUserInfo, r.services, []auth.Role{auth.ADMIN, auth.STUDENT, auth.MANAGER, auth.TEACHER}))

}

func protected[Rq any, Rs any](fn http.HandlerFunc[Rq, Rs], service *RouterServices, roles []auth.Role) func(*gin.Context) {
	fn = auth.Authorize(fn, service.authService, roles)
	fn = auth.Authenticate(fn)
	return mgin.Wrap(fn, service.redirectURL, service.logger)
}
