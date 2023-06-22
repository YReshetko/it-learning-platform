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
	out0.Id = uuidPtrToString(in.ID)
	out0.Name = in.Name
	out0.Description = in.Description

	return out0
}

func (_this_ TechnologyMapperImpl) ToModels(in *_imp_1.GetTechnologiesResponse) _imp_2.Technologies {
	out0 := _imp_2.Technologies{}
	if in != nil {

		_var_0 := in.Technology
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
		out0.ID = stringToUUIDPtr(in.Id)
		out0.Name = in.Name
		out0.Description = in.Description
	}

	return out0
}
