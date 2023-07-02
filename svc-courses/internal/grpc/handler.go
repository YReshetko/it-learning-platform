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
	"strings"
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
	taskMapper                                TaskMapper
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
	logger := h.logger.WithField("method", "GetCategories").WithField("request", rq)

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
	logger := h.logger.WithField("method", "GetTopics").WithField("request", rq)

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

func (h *Handler) GetTopic(_ context.Context, rq *courses.GetTopicRequest) (*courses.TopicResponse, error) {
	logger := h.logger.WithField("method", "GetTopic").WithField("request", rq)
	topicID, err := uuid.Parse(rq.GetTopicId())
	if err != nil {
		logger.WithError(err).Error("unable to parse topic ID")
		return &courses.TopicResponse{}, status.Error(codes.InvalidArgument, "unable to parse topic ID")
	}
	topic, err := h.storage.GetTopic(topicID)
	if err != nil {
		logger.WithError(err).Error("unable to get topic")
		return &courses.TopicResponse{}, status.Error(codes.Internal, "unable to get topic")
	}
	return &courses.TopicResponse{Topic: h.topicMapper.toProto(topic)}, nil
}

func (h *Handler) CreateTag(_ context.Context, rq *courses.CreateTagRequest) (*courses.CreateTagResponse, error) {
	logger := h.logger.WithField("method", "CreateTag").WithField("request", rq)
	now := time.Now()
	tag, err := h.storage.CreateTag(storage.Tag{
		Name:      rq.GetTag().GetName(),
		CreatedAt: now,
		UpdatedAt: now,
	})
	if err != nil {
		logger.WithError(err).Error("unable to create tag")
		return &courses.CreateTagResponse{}, status.Error(codes.Internal, "unable to create tag")
	}
	return &courses.CreateTagResponse{Tag: &courses.Tag{Name: tag.Name}}, nil
}

func (h *Handler) SearchTag(_ context.Context, rq *courses.SearchTagsRequest) (*courses.SearchTagsResponse, error) {
	logger := h.logger.WithField("method", "SearchTag").WithField("request", rq)
	var tags []storage.Tag
	var err error
	if strings.TrimSpace(rq.GetSearch()) == "" {
		tags, err = h.storage.GetTags()
	} else {
		tags, err = h.storage.SearchTags(rq.GetSearch())
	}
	if err != nil {
		logger.WithError(err).Error("unable to find tags")
		return &courses.SearchTagsResponse{}, status.Error(codes.Internal, "unable to find tags")
	}
	protoTags := make([]*courses.Tag, len(tags))
	for i, tag := range tags {
		protoTags[i] = &courses.Tag{Name: tag.Name}
	}

	return &courses.SearchTagsResponse{Tags: protoTags}, nil
}

func (h *Handler) RemoveTag(_ context.Context, rq *courses.RemoveTagRequest) (*emptypb.Empty, error) {
	logger := h.logger.WithField("method", "RemoveTag").WithField("request", rq)
	_, err := h.storage.RemoveTag(storage.Tag{Name: rq.GetTag().GetName()})
	if err != nil {
		logger.WithError(err).Error("unable to remove tag")
		return &emptypb.Empty{}, status.Error(codes.Internal, "unable to remove tag")
	}
	return &emptypb.Empty{}, nil
}

func (h *Handler) AddTopicTag(_ context.Context, rq *courses.AddTopicTagRequest) (*courses.TopicResponse, error) {
	logger := h.logger.WithField("method", "AddTopicTag").WithField("request", rq)
	topicID, err := uuid.Parse(rq.GetTopicId())
	if err != nil {
		logger.WithError(err).Error("unable to parse topic ID")
		return &courses.TopicResponse{}, status.Error(codes.Internal, "unable to parse topic ID")
	}
	topic, err := h.storage.CreateTopicTag(topicID, rq.GetTag().GetName())
	if err != nil {
		logger.WithError(err).Error("unable to add tag to topic")
		return &courses.TopicResponse{}, status.Error(codes.Internal, "unable to add tag to topic")
	}

	return &courses.TopicResponse{Topic: h.topicMapper.toProto(topic)}, nil
}

func (h *Handler) RemoveTopicTag(_ context.Context, rq *courses.RemoveTopicTagRequest) (*courses.TopicResponse, error) {
	logger := h.logger.WithField("method", "RemoveTopicTag").WithField("request", rq)
	topicID, err := uuid.Parse(rq.GetTopicId())
	if err != nil {
		logger.WithError(err).Error("unable to parse topic ID")
		return &courses.TopicResponse{}, status.Error(codes.Internal, "unable to parse topic ID")
	}

	topic, err := h.storage.RemoveTopicTag(topicID, rq.GetTag().GetName())
	if err != nil {
		logger.WithError(err).Error("unable to remove tag from topic")
		return &courses.TopicResponse{}, status.Error(codes.Internal, "unable to remove tag from topic")
	}
	return &courses.TopicResponse{Topic: h.topicMapper.toProto(topic)}, nil
}

func (h *Handler) CreateTask(_ context.Context, request *courses.CreateTaskRequest) (*courses.TaskResponse, error) {
	logger := h.logger.WithField("method", "CreateTask").WithField("request", request)
	var err error
	t := h.taskMapper.toModel(request.GetTask())
	now := time.Now()
	t.CreatedAt = now
	t.UpdatedAt = now

	t, err = h.storage.CreateTask(t)
	if err != nil {
		logger.WithError(err).Error("unable to save task")
		return &courses.TaskResponse{}, status.Error(codes.Internal, "unable to save task")
	}

	return &courses.TaskResponse{Task: h.taskMapper.toProto(t)}, nil
}

func (h *Handler) GetTask(_ context.Context, rq *courses.GetTaskRequest) (*courses.TaskResponse, error) {
	logger := h.logger.WithField("method", "GetTask").WithField("request", rq)
	taskID, err := uuid.Parse(rq.GetTaskId())
	if err != nil {
		logger.WithError(err).Error("unable to parse task ID")
		return &courses.TaskResponse{}, status.Error(codes.InvalidArgument, "unable to parse task ID")
	}
	task, err := h.storage.GetTask(taskID)
	if err != nil {
		logger.WithError(err).Error("unable to get task")
		return &courses.TaskResponse{}, status.Error(codes.Internal, "unable to get task")
	}
	return &courses.TaskResponse{Task: h.taskMapper.toProto(task)}, nil
}

func (h *Handler) FindTasks(_ context.Context, rq *courses.FindTasksRequest) (*courses.TasksResponse, error) {
	logger := h.logger.WithField("method", "FindTasks").WithField("request", rq)

	var tasks []storage.Task
	var err error
	if len(rq.GetTags()) == 0 {
		tasks, err = h.storage.GetTasks()
	} else {
		tags := make([]storage.Tag, len(rq.GetTags()))
		for i, tag := range rq.GetTags() {
			tags[i] = storage.Tag{
				Name: tag.GetName(),
			}
		}
		tasks, err = h.storage.FindTasks(tags)
	}

	if err != nil {
		logger.WithError(err).Error("unable to find tasks")
		return &courses.TasksResponse{}, status.Error(codes.Internal, "unable to find tasks")
	}

	protoTasksValues := h.taskMapper.toProtos(modelTasks{tasks}).Values
	sort.SliceIsSorted(protoTasksValues, func(i, j int) bool {
		return protoTasksValues[i].SeqNo < protoTasksValues[j].SeqNo
	})

	return &courses.TasksResponse{
		Tasks: protoTasksValues,
	}, nil
}

func (h *Handler) AddTaskTag(_ context.Context, rq *courses.AddTaskTagRequest) (*courses.TaskResponse, error) {
	logger := h.logger.WithField("method", "AddTaskTag").WithField("request", rq)
	taskID, err := uuid.Parse(rq.GetTaskId())
	if err != nil {
		logger.WithError(err).Error("unable to parse task ID")
		return &courses.TaskResponse{}, status.Error(codes.Internal, "unable to parse taskID")
	}
	task, err := h.storage.CreateTaskTag(taskID, rq.GetTag().GetName())
	if err != nil {
		logger.WithError(err).Error("unable to add tag to task")
		return &courses.TaskResponse{}, status.Error(codes.Internal, "unable to add tag to task")
	}

	return &courses.TaskResponse{Task: h.taskMapper.toProto(task)}, nil
}

func (h *Handler) RemoveTaskTag(_ context.Context, rq *courses.RemoveTaskTagRequest) (*courses.TaskResponse, error) {
	logger := h.logger.WithField("method", "RemoveTaskTag").WithField("request", rq)
	taskID, err := uuid.Parse(rq.GetTaskId())
	if err != nil {
		logger.WithError(err).Error("unable to parse task ID")
		return &courses.TaskResponse{}, status.Error(codes.Internal, "unable to parse taskID")
	}

	task, err := h.storage.RemoveTaskTag(taskID, rq.GetTag().GetName())
	if err != nil {
		logger.WithError(err).Error("unable to remove tag from task")
		return &courses.TaskResponse{}, status.Error(codes.Internal, "unable to remove tag from task")
	}
	return &courses.TaskResponse{Task: h.taskMapper.toProto(task)}, nil
}
