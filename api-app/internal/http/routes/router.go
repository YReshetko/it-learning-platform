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
	// COMMON API
	engine.GET("/api/v1/self", protected(r.self.GetUserInfo, r.services, model.AllRoles()))
	engine.POST("/api/v1/logout", protected(r.self.Logout, r.services, model.AllRoles()))

	// ADMIN API
	engine.POST("/api/v1/registration/users", protected(r.registration.CreateUsers, r.services, []model.Role{model.ADMIN}))
	engine.POST("/api/v1/registration/user", protected(r.registration.CreateUser, r.services, []model.Role{model.ADMIN}))

	engine.POST("/api/v1/admin/technology", protected(r.courses.CreateTechnology, r.services, []model.Role{model.ADMIN}))
	engine.GET("/api/v1/admin/technologies", protected(r.courses.GetTechnologies, r.services, []model.Role{model.ADMIN}))
	engine.POST("/api/v1/admin/technologies/:technology_id/category", protected(r.courses.CreateCategory, r.services, []model.Role{model.ADMIN}))
	engine.GET("/api/v1/admin/technologies/:technology_id/categories", protected(r.courses.GetCategories, r.services, []model.Role{model.ADMIN}))
	engine.POST("/api/v1/admin/categories/:category_id/topic", protected(r.courses.CreateTopic, r.services, []model.Role{model.ADMIN}))
	engine.GET("/api/v1/admin/categories/:category_id/topics", protected(r.courses.GetTopics, r.services, []model.Role{model.ADMIN}))
	engine.GET("/api/v1/admin/topics/:topic_id", protected(r.courses.GetTopic, r.services, []model.Role{model.ADMIN}))

	engine.POST("/api/v1/admin/tasks", protected(r.courses.CreateTask, r.services, []model.Role{model.ADMIN}))
	engine.GET("/api/v1/admin/tasks/:task_id", protected(r.courses.GetTask, r.services, []model.Role{model.ADMIN}))
	engine.GET("/api/v1/admin/tasks", protected(r.courses.GetTasks, r.services, []model.Role{model.ADMIN}))

	engine.POST("/api/v1/admin/tags", protected(r.courses.CreateTag, r.services, []model.Role{model.ADMIN}))
	engine.GET("/api/v1/admin/tags", protected(r.courses.SearchTags, r.services, []model.Role{model.ADMIN}))
	engine.DELETE("/api/v1/admin/tags/:tag_name", protected(r.courses.RemoveTag, r.services, []model.Role{model.ADMIN}))
	engine.POST("/api/v1/admin/topics/:topic_id/tags", protected(r.courses.AddTopicTag, r.services, []model.Role{model.ADMIN}))
	engine.DELETE("/api/v1/admin/topics/:topic_id/tags/:tag_name", protected(r.courses.RemoveTopicTag, r.services, []model.Role{model.ADMIN}))
	engine.POST("/api/v1/admin/tasks/:task_id/tags", protected(r.courses.AddTaskTag, r.services, []model.Role{model.ADMIN}))
	engine.DELETE("/api/v1/admin/tasks/:task_id/tags/:tag_name", protected(r.courses.RemoveTaskTag, r.services, []model.Role{model.ADMIN}))

	// TEACHER API
	engine.POST("/api/v1/teacher/courses", protected(r.courses.CreateCourse, r.services, []model.Role{model.TEACHER}))
	engine.GET("/api/v1/teacher/courses", protected(r.courses.GetOwnerCourses, r.services, []model.Role{model.TEACHER}))
}

func protected[Rq any, Rs any](fn http.HandlerFunc[Rq, Rs], service *RouterServices, roles []model.Role) func(*gin.Context) {
	fn = auth.Authorize(fn, service.authService, roles)
	fn = auth.Authenticate(fn)
	return mgin.Wrap(fn, service.redirectURL, service.logger)
}
