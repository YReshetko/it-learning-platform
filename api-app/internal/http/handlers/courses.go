package handlers

import (
	"errors"
	"github.com/YReshetko/it-learning-platform/api-app/internal/clients"
	"github.com/YReshetko/it-learning-platform/api-app/internal/http"
	"github.com/YReshetko/it-learning-platform/api-app/internal/http/mappers"
	"github.com/YReshetko/it-learning-platform/api-app/internal/http/models"
	"github.com/YReshetko/it-learning-platform/svc-courses/pb/courses"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/emptypb"
	rest "net/http"
)

/*
Courses handler that returns courses info to frontend
@Optional
*/
type Courses struct {
	client           clients.CoursesClient
	technologyMapper mappers.TechnologyMapper
	categoryMapper   mappers.CategoryMapper
	topicMapper      mappers.TopicMapper
	logger           *logrus.Entry
}

func (c *Courses) CreateTechnology(ctx http.Context, technology models.Technology) (models.Technology, http.Status) {
	logger := c.logger.WithField("method", "CreateTechnology").WithField("technology", technology.Name)
	rs, err := c.client.CreateTechnology(ctx.Context(), &courses.CreateTechnologyRequest{
		Technology: c.technologyMapper.ToProto(technology),
	})
	if err != nil {
		logger.WithError(err).Error("unable to create technology")
		return models.Technology{}, http.Status{
			Error:      err,
			StatusCode: rest.StatusInternalServerError,
			Message:    "unable to create technology",
		}
	}

	return c.technologyMapper.ToModel(rs.Technology), http.Status{StatusCode: rest.StatusCreated}
}

func (c *Courses) GetTechnologies(ctx http.Context, _ models.Empty) (models.Technologies, http.Status) {
	logger := c.logger.WithField("method", "CreateTechnology")
	rs, err := c.client.GetTechnologies(ctx.Context(), &emptypb.Empty{})
	if err != nil {
		logger.WithError(err).Error("unable to get technologies")
		return models.Technologies{}, http.Status{
			Error:      err,
			StatusCode: rest.StatusInternalServerError,
			Message:    "unable to get technologies",
		}
	}

	return c.technologyMapper.ToModels(rs), http.Status{StatusCode: rest.StatusOK}
}

func (c *Courses) CreateCategory(ctx http.Context, category models.Category) (models.Category, http.Status) {
	logger := c.logger.WithField("method", "CreateCategory").WithField("category", category.Name)

	rq := &courses.CreateCategoryRequest{
		Category: c.categoryMapper.ToProto(category),
	}

	technologyID := ctx.GinCtx.Param("technology_id")
	if technologyID == "" {
		return models.Category{}, http.Status{
			Error:      errors.New("empty technology_id"),
			StatusCode: rest.StatusBadRequest,
			Message:    "unable to create category",
		}
	}
	rq.Category.TechnologyId = technologyID
	rs, err := c.client.CreateCategory(ctx.Context(), rq)
	if err != nil {
		logger.WithError(err).Error("unable to create category")
		return models.Category{}, http.Status{
			Error:      err,
			StatusCode: rest.StatusInternalServerError,
			Message:    "unable to create category",
		}
	}

	return c.categoryMapper.ToModel(rs.Category), http.Status{StatusCode: rest.StatusCreated}
}

func (c *Courses) GetCategories(ctx http.Context, _ models.Empty) (models.Categories, http.Status) {
	logger := c.logger.WithField("method", "GetCategories")

	technologyID := ctx.GinCtx.Param("technology_id")
	if technologyID == "" {
		return models.Categories{}, http.Status{
			Error:      errors.New("empty technology_id"),
			StatusCode: rest.StatusBadRequest,
			Message:    "unable to get categories",
		}
	}

	rs, err := c.client.GetCategories(ctx.Context(), &courses.GetCategoriesRequest{TechnologyId: technologyID})
	if err != nil {
		logger.WithError(err).Error("unable to get categories")
		return models.Categories{}, http.Status{
			Error:      err,
			StatusCode: rest.StatusInternalServerError,
			Message:    "unable to get categories",
		}
	}

	return c.categoryMapper.ToModels(rs), http.Status{StatusCode: rest.StatusOK}
}

func (c *Courses) CreateTopic(ctx http.Context, topic models.Topic) (models.Topic, http.Status) {
	logger := c.logger.WithField("method", "CreateTopic").WithField("topic", topic.Name)

	rq := &courses.CreateTopicRequest{
		Topic: c.topicMapper.ToProto(topic),
	}

	categoryID := ctx.GinCtx.Param("category_id")
	if categoryID == "" {
		return models.Topic{}, http.Status{
			Error:      errors.New("empty category_id"),
			StatusCode: rest.StatusBadRequest,
			Message:    "unable to create topic",
		}
	}
	rq.Topic.CategoryId = categoryID
	rs, err := c.client.CreateTopic(ctx.Context(), rq)
	if err != nil {
		logger.WithError(err).Error("unable to create topic")
		return models.Topic{}, http.Status{
			Error:      err,
			StatusCode: rest.StatusInternalServerError,
			Message:    "unable to create topic",
		}
	}

	return c.topicMapper.ToModel(rs.Topic), http.Status{StatusCode: rest.StatusCreated}
}

func (c *Courses) GetTopics(ctx http.Context, _ models.Empty) (models.Topics, http.Status) {
	logger := c.logger.WithField("method", "GetTopics")

	categoryID := ctx.GinCtx.Param("category_id")
	if categoryID == "" {
		return models.Topics{}, http.Status{
			Error:      errors.New("empty category_id"),
			StatusCode: rest.StatusBadRequest,
			Message:    "unable to get topics",
		}
	}

	rs, err := c.client.GetTopics(ctx.Context(), &courses.GetTopicsRequest{CategoryId: categoryID})
	if err != nil {
		logger.WithError(err).Error("unable to get topics")
		return models.Topics{}, http.Status{
			Error:      err,
			StatusCode: rest.StatusInternalServerError,
			Message:    "unable to get topics",
		}
	}

	return c.topicMapper.ToModels(rs), http.Status{StatusCode: rest.StatusOK}
}
