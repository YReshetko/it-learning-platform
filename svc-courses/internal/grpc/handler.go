package grpc

import (
	"context"
	"github.com/YReshetko/it-learning-platform/svc-courses/internal/storage"
	"github.com/YReshetko/it-learning-platform/svc-courses/pb/courses"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

/*
Handler the GRPC handler
@Constructor
*/
type Handler struct {
	courses.UnimplementedCoursesServiceServer // @Exclude
	storage                                   storage.CourseStorage
	technologyMapper                          TechnologyMapper
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
		Technology: h.technologyMapper.toProtos(modelTechnologies{technologies}).Values,
	}, nil
}
