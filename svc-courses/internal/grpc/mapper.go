package grpc

import (
	"github.com/YReshetko/it-learning-platform/svc-courses/internal/storage"
	"github.com/YReshetko/it-learning-platform/svc-courses/pb/courses"
	"github.com/google/uuid"
)

var emptyUUID uuid.UUID

type modelTechnologies struct {
	Values []storage.Technology
}
type protoTechnologies struct {
	Values []*courses.Technology
}

// TechnologyMapper the mapper model storage.Technology from\to proto courses.Technology
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

type modelCategories struct {
	Values []storage.Category
}
type protoCategories struct {
	Values []*courses.Category
}

// CategoryMapper the mapper model storage.Category from\to proto courses.Category
// @Mapper
type CategoryMapper interface {
	// @SliceMapping(target="Values", source="in.Values", this="toProto")
	toProtos(in modelCategories) protoCategories
	// @Mapping(target="Id", func="uuidPtrToString(in.ID)")
	// @Mapping(target="TechnologyId", func="uuidToString(in.TechnologyID)")
	toProto(in storage.Category) *courses.Category
	// @SliceMapping(target="Values", source="in.Values", this="toModel")
	toModels(in protoCategories) modelCategories
	// @Mapping(target="ID", func="stringToUUIDPtr(in.Id)")
	// @Mapping(target="TechnologyID", func="stringToUUID(in.TechnologyId)")
	toModel(in *courses.Category) storage.Category
}

type modelTopics struct {
	Values []storage.Topic
}
type protoTopics struct {
	Values []*courses.Topic
}

// TopicMapper the mapper model storage.Topic from\to proto courses.Topic
// @Mapper
type TopicMapper interface {
	// @SliceMapping(target="Values", source="in.Values", this="toProto")
	toProtos(in modelTopics) protoTopics
	// @Mapping(target="Id", func="uuidPtrToString(in.ID)")
	// @Mapping(target="CategoryId", func="uuidToString(in.CategoryID)")
	// @SliceMapping(target="Tags", source="in.Tags", this="toTagProto")
	toProto(in storage.Topic) *courses.Topic
	// @SliceMapping(target="Values", source="in.Values", this="toModel")
	toModels(in protoTopics) modelTopics
	// @Mapping(target="ID", func="stringToUUIDPtr(in.Id)")
	// @Mapping(target="CategoryID", func="stringToUUID(in.CategoryId)")
	// @SliceMapping(target="Tags", source="in.Tags", this="toTagModel")
	toModel(in *courses.Topic) storage.Topic
	toTagProto(in storage.Tag) *courses.Tag
	toTagModel(in *courses.Tag) storage.Tag
}

type modelTasks struct {
	Values []storage.Task
}
type protoTasks struct {
	Values []*courses.Task
}

// TaskMapper the mapper model storage.Task from\to proto courses.Task
// @Mapper
type TaskMapper interface {
	// @SliceMapping(target="Values", source="in.Values", this="toProto")
	toProtos(in modelTasks) protoTasks
	// @Mapping(target="Id", func="uuidPtrToString(in.ID)")
	// @SliceMapping(target="Tags", source="in.Tags", this="toTagProto")
	toProto(in storage.Task) *courses.Task
	// @SliceMapping(target="Values", source="in.Values", this="toModel")
	toModels(in protoTasks) modelTasks
	// @Mapping(target="ID", func="stringToUUIDPtr(in.Id)")
	// @SliceMapping(target="Tags", source="in.Tags", this="toTagModel")
	toModel(in *courses.Task) storage.Task
	toTagProto(in storage.Tag) *courses.Tag
	toTagModel(in *courses.Tag) storage.Tag
}
type modelCourses struct {
	Values []storage.Course
}

// CourseMapper the mapper model storage.Course from\to proto courses.Course
// @Mapper
type CourseMapper interface {
	// @Mapping(target="ID", func="stringToUUIDPtr(in.Id)")
	courseToModel(in *courses.Course) storage.Course
	// @SliceMapping(target="Course", source="in.Values", this="courseToProto")
	coursesToProto(in modelCourses) *courses.CoursesResponse
	// @Mapping(target="Id", func="uuidPtrToString(in.ID)")
	courseToProto(in storage.Course) *courses.Course
}

func uuidPtrToString(id *uuid.UUID) string {
	if id == nil {
		return ""
	}
	return id.String()
}

func uuidToString(id uuid.UUID) string {
	if id == emptyUUID {
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

func stringToUUID(id string) uuid.UUID {
	if id == "" {
		return emptyUUID
	}
	out, _ := uuid.Parse(id)
	return out
}
