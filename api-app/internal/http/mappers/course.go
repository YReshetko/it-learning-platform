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
	// ToProto @Mapping(target="Id", func="uuidToString(in.ID)")
	ToProto(in models.Technology) *courses.Technology
	// ToModels @SliceMapping(target="Technologies", source="in.Technologies", this="ToModel")
	ToModels(in *courses.GetTechnologiesResponse) models.Technologies
	// ToModel @Mapping(target="ID", func="stringToUUID(in.Id)")
	ToModel(in *courses.Technology) models.Technology
}

type ProtoCategories struct {
	Values []*courses.Category
}

// CategoryMapper mapper from proto to API models and vise versa
// @Mapper
type CategoryMapper interface {
	// ToProtos @SliceMapping(target="Values", source="in.Categories", this="ToProto")
	ToProtos(in models.Categories) ProtoCategories
	// ToProto @Mapping(target="Id", func="uuidToString(in.ID)")
	// ToProto @Mapping(target="TechnologyId", func="uuidToString(in.TechnologyID)")
	ToProto(in models.Category) *courses.Category
	// ToModels @SliceMapping(target="Categories", source="in.Categories", this="ToModel")
	ToModels(in *courses.GetCategoriesResponse) models.Categories
	// ToModel @Mapping(target="ID", func="stringToUUID(in.Id)")
	// ToModel @Mapping(target="TechnologyID", func="stringToUUID(in.TechnologyId)")
	ToModel(in *courses.Category) models.Category
}

type ProtoTopics struct {
	Values []*courses.Topic
}

// TopicMapper mapper from proto to API models and vise versa
// @Mapper
type TopicMapper interface {
	// ToProtos @SliceMapping(target="Values", source="in.Topics", this="ToProto")
	ToProtos(in models.Topics) ProtoTopics
	// ToProto @Mapping(target="Id", func="uuidToString(in.ID)")
	// @Mapping(target="CategoryId", func="uuidToString(in.CategoryID)")
	// @SliceMapping(target="Tags", source="in.Tags", this="toTagProto")
	ToProto(in models.Topic) *courses.Topic
	// ToModels @SliceMapping(target="Topics", source="in.Topics", this="ToModel")
	ToModels(in *courses.GetTopicsResponse) models.Topics
	// ToModel @Mapping(target="ID", func="stringToUUID(in.Id)")
	// @Mapping(target="CategoryID", func="stringToUUID(in.CategoryId)")
	// @SliceMapping(target="Tags", source="in.Tags", this="toTagModel")
	ToModel(in *courses.Topic) models.Topic
	toTagProto(in models.Tag) *courses.Tag
	toTagModel(in *courses.Tag) models.Tag
}

type ProtoTasks struct {
	Values []*courses.Task
}

// TaskMapper mapper from proto to API models and vise versa
// @Mapper
type TaskMapper interface {
	// ToProtos @SliceMapping(target="Values", source="in.Tasks", this="ToProto")
	ToProtos(in models.Tasks) ProtoTasks
	// ToProto @Mapping(target="Id", func="uuidToString(in.ID)")
	// @SliceMapping(target="Tags", source="in.Tags", this="toTagProto")
	ToProto(in models.Task) *courses.Task
	// ToModels @SliceMapping(target="Tasks", source="in.Tasks", this="ToModel")
	ToModels(in *courses.TasksResponse) models.Tasks
	// ToModel @Mapping(target="ID", func="stringToUUID(in.Id)")
	// @SliceMapping(target="Tags", source="in.Tags", this="toTagModel")
	ToModel(in *courses.Task) models.Task
	toTagProto(in models.Tag) *courses.Tag
	toTagModel(in *courses.Tag) models.Tag
}

func uuidToString(id uuid.UUID) string {
	if id == emptyUUID {
		return ""
	}
	return id.String()
}

func stringToUUID(id string) uuid.UUID {
	if id == "" {
		return uuid.New()
	}
	out, _ := uuid.Parse(id)
	return out
}
