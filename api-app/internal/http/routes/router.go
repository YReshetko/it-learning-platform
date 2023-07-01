package routes

import (
	"github.com/YReshetko/it-learning-platform/api-app/internal/http"
	auth "github.com/YReshetko/it-learning-platform/api-app/internal/http/middlewares/authorization"
	"github.com/YReshetko/it-learning-platform/lib-app/pkg/model"
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
	courses      handlers.Courses

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
	engine.POST("/api/v1/registration/user", protected(r.registration.CreateUser, r.services, []model.Role{model.ADMIN}))
	engine.POST("/api/v1/registration/users", protected(r.registration.CreateUsers, r.services, []model.Role{model.ADMIN}))
	engine.GET("/api/v1/self", protected(r.self.GetUserInfo, r.services, model.AllRoles()))
	engine.POST("/api/v1/logout", protected(r.self.Logout, r.services, model.AllRoles()))

	engine.POST("/api/v1/admin/technology", protected(r.courses.CreateTechnology, r.services, []model.Role{model.ADMIN}))
	engine.GET("/api/v1/admin/technologies", protected(r.courses.GetTechnologies, r.services, []model.Role{model.ADMIN}))
	engine.POST("/api/v1/admin/technologies/:technology_id/category", protected(r.courses.CreateCategory, r.services, []model.Role{model.ADMIN}))
	engine.GET("/api/v1/admin/technologies/:technology_id/categories", protected(r.courses.GetCategories, r.services, []model.Role{model.ADMIN}))
	engine.POST("/api/v1/admin/categories/:category_id/topic", protected(r.courses.CreateTopic, r.services, []model.Role{model.ADMIN}))
	engine.GET("/api/v1/admin/categories/:category_id/topics", protected(r.courses.GetTopics, r.services, []model.Role{model.ADMIN}))
}

func protected[Rq any, Rs any](fn http.HandlerFunc[Rq, Rs], service *RouterServices, roles []model.Role) func(*gin.Context) {
	fn = auth.Authorize(fn, service.authService, roles)
	fn = auth.Authenticate(fn)
	return mgin.Wrap(fn, service.redirectURL, service.logger)
}
