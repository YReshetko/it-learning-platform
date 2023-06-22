package mappers

import (
	"github.com/YReshetko/it-learning-platform/api-app/internal/http/models"
	"github.com/YReshetko/it-learning-platform/svc-courses/pb/courses"
	"github.com/google/uuid"
)

type ProtoTechnologies struct {
	Values []*courses.Technology
}

// @Mapper
type TechnologyMapper interface {
	// @SliceMapping(target="Values", source="in.Technologies", this="ToProto")
	ToProtos(in models.Technologies) ProtoTechnologies
	// @Mapping(target="Id", func="uuidPtrToString(in.ID)")
	ToProto(in models.Technology) *courses.Technology
	// @SliceMapping(target="Technologies", source="in.Technology", this="ToModel")
	ToModels(in *courses.GetTechnologiesResponse) models.Technologies
	// @Mapping(target="ID", func="stringToUUIDPtr(in.Id)")
	ToModel(in *courses.Technology) models.Technology
}

func uuidPtrToString(id uuid.UUID) string {
	return id.String()
}

func stringToUUIDPtr(id string) uuid.UUID {
	if id == "" {
		return uuid.New()
	}
	out, _ := uuid.Parse(id)
	return out
}
