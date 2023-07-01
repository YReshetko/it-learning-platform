package grpc

import (
	"context"
	"github.com/YReshetko/it-learning-platform/svc-courses/internal/storage"
	"github.com/YReshetko/it-learning-platform/svc-courses/pb/courses"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"sort"
	"time"
)

/*
Handler the GRPC handler
@Optional
*/
type Handler struct {
	courses.UnimplementedCoursesServiceServer // @Exclude
	storage                                   storage.CourseStorage
	technologyMapper                          TechnologyMapper
	categoryMapper                            CategoryMapper
	topicMapper                               TopicMapper
	logger                                    *logrus.Entry
}

func (h *Handler) CreateTechnology(_ context.Context, request *courses.CreateTechnologyRequest) (*courses.CreateTechnologyResponse, error) {
	logger := h.logger.WithField("method", "CreateTechnology").WithField("request", request)
	var err error
	t := h.technologyMapper.toModel(request.GetTechnology())
	now := time.Now()
	t.CreatedAt = now
	t.UpdatedAt = now

	t, err = h.storage.CreateTechnology(t)
	if err != nil {
		logger.WithError(err).Error("unable to save technology")
		return &courses.CreateTechnologyResponse{}, status.Error(codes.Internal, "unable to save technology")
	}

	return &courses.CreateTechnologyResponse{Technology: h.technologyMapper.toProto(t)}, nil
}
func (h *Handler) GetTechnologies(context.Context, *emptypb.Empty) (*courses.GetTechnologiesResponse, error) {
	logger := h.logger.WithField("method", "GetTechnologies")
	technologies, err := h.storage.GetTechnologies()
	if err != nil {
		logger.WithError(err).Error("unable to get technologies")
		return &courses.GetTechnologiesResponse{}, status.Error(codes.Internal, "unable to get technologies")
	}

	return &courses.GetTechnologiesResponse{
		Technologies: h.technologyMapper.toProtos(modelTechnologies{technologies}).Values,
	}, nil
}

func (h *Handler) CreateCategory(_ context.Context, request *courses.CreateCategoryRequest) (*courses.CreateCategoryResponse, error) {
	logger := h.logger.WithField("method", "CreateCategory").WithField("request", request)
	var err error
	c := h.categoryMapper.toModel(request.GetCategory())
	now := time.Now()
	c.CreatedAt = now
	c.UpdatedAt = now

	c, err = h.storage.CreateCategory(c)
	if err != nil {
		logger.WithError(err).Error("unable to save category")
		return &courses.CreateCategoryResponse{}, status.Error(codes.Internal, "unable to save category")
	}

	return &courses.CreateCategoryResponse{Category: h.categoryMapper.toProto(c)}, nil
}

func (h *Handler) GetCategories(_ context.Context, rq *courses.GetCategoriesRequest) (*courses.GetCategoriesResponse, error) {
	logger := h.logger.WithField("method", "GetCategories")

	technologyID, err := uuid.Parse(rq.TechnologyId)
	if err != nil {
		logger.WithError(err).Error("unable to parse technology ID")
		return &courses.GetCategoriesResponse{}, status.Error(codes.InvalidArgument, "unable to parse technology ID")
	}

	categories, err := h.storage.GetCategories(technologyID)
	if err != nil {
		logger.WithError(err).Error("unable to get categories")
		return &courses.GetCategoriesResponse{}, status.Error(codes.Internal, "unable to get categories")
	}

	return &courses.GetCategoriesResponse{
		Categories: h.categoryMapper.toProtos(modelCategories{categories}).Values,
	}, nil
}

func (h *Handler) CreateTopic(_ context.Context, request *courses.CreateTopicRequest) (*courses.CreateTopicResponse, error) {
	logger := h.logger.WithField("method", "CreateTopic").WithField("request", request)
	var err error
	t := h.topicMapper.toModel(request.GetTopic())
	now := time.Now()
	t.CreatedAt = now
	t.UpdatedAt = now

	t, err = h.storage.CreateTopic(t)
	if err != nil {
		logger.WithError(err).Error("unable to save topic")
		return &courses.CreateTopicResponse{}, status.Error(codes.Internal, "unable to save topic")
	}

	return &courses.CreateTopicResponse{Topic: h.topicMapper.toProto(t)}, nil
}

func (h *Handler) GetTopics(_ context.Context, rq *courses.GetTopicsRequest) (*courses.GetTopicsResponse, error) {
	logger := h.logger.WithField("method", "GetTopics")

	categoryID, err := uuid.Parse(rq.CategoryId)
	if err != nil {
		logger.WithError(err).Error("unable to parse category ID")
		return &courses.GetTopicsResponse{}, status.Error(codes.InvalidArgument, "unable to parse category ID")
	}

	topics, err := h.storage.GetTopics(categoryID)
	if err != nil {
		logger.WithError(err).Error("unable to get topics")
		return &courses.GetTopicsResponse{}, status.Error(codes.Internal, "unable to get topics")
	}

	protoTopicValues := h.topicMapper.toProtos(modelTopics{topics}).Values
	sort.SliceIsSorted(protoTopicValues, func(i, j int) bool {
		return protoTopicValues[i].SeqNo < protoTopicValues[j].SeqNo
	})

	return &courses.GetTopicsResponse{
		Topics: protoTopicValues,
	}, nil
}
