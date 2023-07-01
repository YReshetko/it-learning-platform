package mappers

import (
	"github.com/YReshetko/it-learning-platform/api-app/internal/http/models"
	"github.com/YReshetko/it-learning-platform/svc-courses/pb/courses"
	"github.com/google/uuid"
)

type ProtoTechnologies struct {
	Values []*courses.Technology
}

var emptyUUID uuid.UUID

// TechnologyMapper mapper from proto to API models and vise versa
// @Mapper
type TechnologyMapper interface {
	// ToProtos @SliceMapping(target="Values", source="in.Technologies", this="ToProto")
	ToProtos(in models.Technologies) ProtoTechnologies
	// ToProto @Mapping(target="Id", func="uuidPtrToString(in.ID)")
	ToProto(in models.Technology) *courses.Technology
	// ToModels @SliceMapping(target="Technologies", source="in.Technology", this="ToModel")
	ToModels(in *courses.GetTechnologiesResponse) models.Technologies
	// ToModel @Mapping(target="ID", func="stringToUUIDPtr(in.Id)")
	ToModel(in *courses.Technology) models.Technology
}

func uuidPtrToString(id uuid.UUID) string {
	if id == emptyUUID {
		return ""
	}
	return id.String()
}

func stringToUUIDPtr(id string) uuid.UUID {
	if id == "" {
		return uuid.New()
	}
	out, _ := uuid.Parse(id)
	return out
}
