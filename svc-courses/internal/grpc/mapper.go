package grpc

import (
	"github.com/YReshetko/it-learning-platform/svc-courses/internal/storage"
	"github.com/YReshetko/it-learning-platform/svc-courses/pb/courses"
	"github.com/google/uuid"
)

type modelTechnologies struct {
	Values []storage.Technology
}
type protoTechnologies struct {
	Values []*courses.Technology
}

// @Mapper
type TechnologyMapper interface {
	// @SliceMapping(target="Values", source="in.Values", this="toProto")
	toProtos(in modelTechnologies) protoTechnologies
	// @Mapping(target="Id", func="uuidPtrToString(in.ID)")
	toProto(in storage.Technology) *courses.Technology
	// @SliceMapping(target="Values", source="in.Values", this="toModel")
	toModels(in protoTechnologies) modelTechnologies
	// @Mapping(target="ID", func="stringToUUIDPtr(in.Id)")
	toModel(in *courses.Technology) storage.Technology
}

func uuidPtrToString(id *uuid.UUID) string {
	if id == nil {
		return ""
	}
	return id.String()
}

func stringToUUIDPtr(id string) *uuid.UUID {
	if id == "" {
		return nil
	}
	out, _ := uuid.Parse(id)
	return &out
}
