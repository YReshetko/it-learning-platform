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
		logger.Error("no category id")
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

func (c *Courses) GetTopic(ctx http.Context, _ models.Empty) (models.Topic, http.Status) {
	logger := c.logger.WithField("method", "GetTopic")
	topicID := ctx.GinCtx.Param("topic_id")
	if topicID == "" {
		logger.Error("no topic id")
		return models.Topic{}, http.Status{
			Error:      errors.New("empty topic_id"),
			StatusCode: rest.StatusBadRequest,
			Message:    "unable to get topics",
		}
	}
	topic, err := c.client.GetTopic(ctx.Context(), &courses.GetTopicRequest{TopicId: topicID})
	if err != nil {
		logger.WithError(err).Error("unable to get topic")
		return models.Topic{}, http.Status{
			Error:      err,
			StatusCode: rest.StatusInternalServerError,
			Message:    "unable to get topic",
		}
	}
	return c.topicMapper.ToModel(topic.GetTopic()), http.Status{StatusCode: rest.StatusOK}
}

func (c *Courses) CreateTag(ctx http.Context, tag models.Tag) (models.Tag, http.Status) {
	logger := c.logger.WithField("method", "CreateTag")
	rs, err := c.client.CreateTag(ctx.Context(), &courses.CreateTagRequest{Tag: &courses.Tag{Name: tag.Name}})
	if err != nil {
		logger.WithError(err).Error("unable to create tag")
		return models.Tag{}, http.Status{
			Error:      err,
			StatusCode: rest.StatusInternalServerError,
			Message:    "unable to create tag",
		}
	}
	return models.Tag{Name: rs.GetTag().GetName()}, http.Status{StatusCode: rest.StatusCreated}
}

func (c *Courses) SearchTags(ctx http.Context, _ models.Empty) (models.Tags, http.Status) {
	logger := c.logger.WithField("method", "SearchTags")
	rs, err := c.client.SearchTag(ctx.Context(), &courses.SearchTagsRequest{Search: ctx.GinCtx.Query("search")})
	if err != nil {
		logger.WithError(err).Error("unable to find tags")
		return models.Tags{}, http.Status{
			Error:      err,
			StatusCode: rest.StatusInternalServerError,
			Message:    "unable to find tags",
		}
	}

	tags := make([]models.Tag, len(rs.GetTags()))
	for i, tag := range rs.GetTags() {
		tags[i] = models.Tag{Name: tag.GetName()}
	}
	return models.Tags{Tags: tags}, http.Status{StatusCode: rest.StatusOK}
}

func (c *Courses) RemoveTag(ctx http.Context, _ models.Empty) (models.Empty, http.Status) {
	logger := c.logger.WithField("method", "RemoveTag")
	tagName := ctx.GinCtx.Param("tag_name")
	_, err := c.client.RemoveTag(ctx.Context(), &courses.RemoveTagRequest{Tag: &courses.Tag{Name: tagName}})
	if err != nil {
		logger.WithError(err).Error("unable to remove tag")
		return models.Empty{}, http.Status{
			Error:      err,
			StatusCode: rest.StatusInternalServerError,
			Message:    "unable to remove tag",
		}
	}
	return models.Empty{}, http.Status{StatusCode: rest.StatusNoContent}
}

func (c *Courses) AddTopicTag(ctx http.Context, tag models.Tag) (models.Topic, http.Status) {
	logger := c.logger.WithField("method", "AddTopicTag")
	topicID := ctx.GinCtx.Param("topic_id")

	topic, err := c.client.AddTopicTag(ctx.Context(), &courses.AddTopicTagRequest{
		TopicId: topicID,
		Tag:     &courses.Tag{Name: tag.Name},
	})
	if err != nil {
		logger.WithError(err).Error("unable to add topic tag")
		return models.Topic{}, http.Status{
			Error:      err,
			StatusCode: rest.StatusInternalServerError,
			Message:    "unable to add topic tag",
		}
	}

	return c.topicMapper.ToModel(topic.GetTopic()), http.Status{StatusCode: rest.StatusOK}
}

func (c *Courses) RemoveTopicTag(ctx http.Context, _ models.Empty) (models.Topic, http.Status) {
	logger := c.logger.WithField("method", "RemoveTopicTag")
	topicID := ctx.GinCtx.Param("topic_id")
	tagName := ctx.GinCtx.Param("tag_name")

	topic, err := c.client.RemoveTopicTag(ctx.Context(), &courses.RemoveTopicTagRequest{
		TopicId: topicID,
		Tag:     &courses.Tag{Name: tagName},
	})
	if err != nil {
		logger.WithError(err).Error("unable to remove topic tag")
		return models.Topic{}, http.Status{
			Error:      err,
			StatusCode: rest.StatusInternalServerError,
			Message:    "unable to remove topic tag",
		}
	}

	return c.topicMapper.ToModel(topic.GetTopic()), http.Status{StatusCode: rest.StatusOK}
}
