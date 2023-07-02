// Code generated by Mapper annotation processor. DO NOT EDIT.
// versions:
//		go: go1.20.4
//		go-annotation: 0.1.0
//		Mapper: 0.0.1-alpha

package mappers

import (
	_imp_2 "github.com/YReshetko/it-learning-platform/api-app/internal/http/models"
	_imp_1 "github.com/YReshetko/it-learning-platform/svc-courses/pb/courses"
)

var _ TechnologyMapper = (*TechnologyMapperImpl)(nil)

type TechnologyMapperImpl struct{}

func (_this_ TechnologyMapperImpl) ToProtos(in _imp_2.Technologies) ProtoTechnologies {
	out0 := ProtoTechnologies{}

	_var_0 := in.Technologies
	_var_1 := make([]*_imp_1.Technology, len(_var_0), len(_var_0))
	for _var_2, _var_3 := range _var_0 {
		_var_1[_var_2] = _this_.ToProto(_var_3)
	}
	out0.Values = _var_1

	return out0
}

func (_this_ TechnologyMapperImpl) ToProto(in _imp_2.Technology) *_imp_1.Technology {
	out0 := &_imp_1.Technology{}
	out0.Id = uuidToString(in.ID)
	out0.Name = in.Name
	out0.Description = in.Description

	return out0
}

func (_this_ TechnologyMapperImpl) ToModels(in *_imp_1.GetTechnologiesResponse) _imp_2.Technologies {
	out0 := _imp_2.Technologies{}
	if in != nil {

		_var_0 := in.Technologies
		_var_1 := make([]_imp_2.Technology, len(_var_0), len(_var_0))
		for _var_2, _var_3 := range _var_0 {
			_var_1[_var_2] = _this_.ToModel(_var_3)
		}
		out0.Technologies = _var_1

	}

	return out0
}

func (_this_ TechnologyMapperImpl) ToModel(in *_imp_1.Technology) _imp_2.Technology {
	out0 := _imp_2.Technology{}
	if in != nil {
		out0.ID = stringToUUID(in.Id)
		out0.Name = in.Name
		out0.Description = in.Description
	}

	return out0
}

var _ CategoryMapper = (*CategoryMapperImpl)(nil)

type CategoryMapperImpl struct{}

func (_this_ CategoryMapperImpl) ToProtos(in _imp_2.Categories) ProtoCategories {
	out0 := ProtoCategories{}

	_var_0 := in.Categories
	_var_1 := make([]*_imp_1.Category, len(_var_0), len(_var_0))
	for _var_2, _var_3 := range _var_0 {
		_var_1[_var_2] = _this_.ToProto(_var_3)
	}
	out0.Values = _var_1

	return out0
}

func (_this_ CategoryMapperImpl) ToProto(in _imp_2.Category) *_imp_1.Category {
	out0 := &_imp_1.Category{}
	out0.Id = uuidToString(in.ID)
	out0.TechnologyId = uuidToString(in.TechnologyID)
	out0.Name = in.Name
	out0.Description = in.Description

	return out0
}

func (_this_ CategoryMapperImpl) ToModels(in *_imp_1.GetCategoriesResponse) _imp_2.Categories {
	out0 := _imp_2.Categories{}
	if in != nil {

		_var_0 := in.Categories
		_var_1 := make([]_imp_2.Category, len(_var_0), len(_var_0))
		for _var_2, _var_3 := range _var_0 {
			_var_1[_var_2] = _this_.ToModel(_var_3)
		}
		out0.Categories = _var_1

	}

	return out0
}

func (_this_ CategoryMapperImpl) ToModel(in *_imp_1.Category) _imp_2.Category {
	out0 := _imp_2.Category{}
	if in != nil {
		out0.ID = stringToUUID(in.Id)
		out0.TechnologyID = stringToUUID(in.TechnologyId)
		out0.Name = in.Name
		out0.Description = in.Description
	}

	return out0
}

var _ TopicMapper = (*TopicMapperImpl)(nil)

type TopicMapperImpl struct{}

func (_this_ TopicMapperImpl) ToProtos(in _imp_2.Topics) ProtoTopics {
	out0 := ProtoTopics{}

	_var_0 := in.Topics
	_var_1 := make([]*_imp_1.Topic, len(_var_0), len(_var_0))
	for _var_2, _var_3 := range _var_0 {
		_var_1[_var_2] = _this_.ToProto(_var_3)
	}
	out0.Values = _var_1

	return out0
}

func (_this_ TopicMapperImpl) ToProto(in _imp_2.Topic) *_imp_1.Topic {
	out0 := &_imp_1.Topic{}
	out0.Id = uuidToString(in.ID)
	out0.CategoryId = uuidToString(in.CategoryID)
	out0.SeqNo = func(v int) int32 {
		res := int32(v)
		return res
	}(in.SeqNo)
	out0.Name = in.Name
	out0.Description = in.Description
	out0.Active = in.Active

	_var_0 := in.Tags
	_var_1 := make([]*_imp_1.Tag, len(_var_0), len(_var_0))
	for _var_2, _var_3 := range _var_0 {
		_var_1[_var_2] = _this_.toTagProto(_var_3)
	}
	out0.Tags = _var_1

	return out0
}

func (_this_ TopicMapperImpl) ToModels(in *_imp_1.GetTopicsResponse) _imp_2.Topics {
	out0 := _imp_2.Topics{}
	if in != nil {

		_var_0 := in.Topics
		_var_1 := make([]_imp_2.Topic, len(_var_0), len(_var_0))
		for _var_2, _var_3 := range _var_0 {
			_var_1[_var_2] = _this_.ToModel(_var_3)
		}
		out0.Topics = _var_1

	}

	return out0
}

func (_this_ TopicMapperImpl) ToModel(in *_imp_1.Topic) _imp_2.Topic {
	out0 := _imp_2.Topic{}
	if in != nil {
		out0.ID = stringToUUID(in.Id)
		out0.CategoryID = stringToUUID(in.CategoryId)
		out0.SeqNo = func(v int32) int {
			res := int(v)
			return res
		}(in.SeqNo)
		out0.Name = in.Name
		out0.Description = in.Description
		out0.Active = in.Active

		_var_0 := in.Tags
		_var_1 := make([]_imp_2.Tag, len(_var_0), len(_var_0))
		for _var_2, _var_3 := range _var_0 {
			_var_1[_var_2] = _this_.toTagModel(_var_3)
		}
		out0.Tags = _var_1

	}

	return out0
}

func (_this_ TopicMapperImpl) toTagProto(in _imp_2.Tag) *_imp_1.Tag {
	out0 := &_imp_1.Tag{}
	out0.Name = in.Name

	return out0
}

func (_this_ TopicMapperImpl) toTagModel(in *_imp_1.Tag) _imp_2.Tag {
	out0 := _imp_2.Tag{}
	if in != nil {
		out0.Name = in.Name
	}

	return out0
}
