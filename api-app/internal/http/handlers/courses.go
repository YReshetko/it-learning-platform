package handlers

import (
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
@Constructor
*/
type Courses struct {
	client           clients.CoursesClient
	technologyMapper mappers.TechnologyMapper
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
		logger.WithError(err).Error("unable to create technology")
		return models.Technologies{}, http.Status{
			Error:      err,
			StatusCode: rest.StatusInternalServerError,
			Message:    "unable to get technologies",
		}
	}

	return c.technologyMapper.ToModels(rs), http.Status{StatusCode: rest.StatusOK}
}
