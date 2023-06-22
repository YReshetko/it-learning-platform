// Code generated by Mapper annotation processor. DO NOT EDIT.
// versions:
//		go: go1.20.4
//		go-annotation: 0.1.0
//		Mapper: 0.0.1-alpha

package grpc

import (
	_imp_3 "github.com/YReshetko/it-learning-platform/svc-courses/internal/storage"
	_imp_1 "github.com/YReshetko/it-learning-platform/svc-courses/pb/courses"
)

var _ TechnologyMapper = (*TechnologyMapperImpl)(nil)

type TechnologyMapperImpl struct{}

func (_this_ TechnologyMapperImpl) toProtos(in modelTechnologies) protoTechnologies {
	out0 := protoTechnologies{}

	_var_0 := in.Values
	_var_1 := make([]*_imp_1.Technology, len(_var_0), len(_var_0))
	for _var_2, _var_3 := range _var_0 {
		_var_1[_var_2] = _this_.toProto(_var_3)
	}
	out0.Values = _var_1

	return out0
}

func (_this_ TechnologyMapperImpl) toProto(in _imp_3.Technology) *_imp_1.Technology {
	out0 := &_imp_1.Technology{}
	out0.Id = uuidPtrToString(in.ID)
	out0.Name = in.Name
	out0.Description = in.Description

	return out0
}

func (_this_ TechnologyMapperImpl) toModels(in protoTechnologies) modelTechnologies {
	out0 := modelTechnologies{}

	_var_0 := in.Values
	_var_1 := make([]_imp_3.Technology, len(_var_0), len(_var_0))
	for _var_2, _var_3 := range _var_0 {
		_var_1[_var_2] = _this_.toModel(_var_3)
	}
	out0.Values = _var_1

	return out0
}

func (_this_ TechnologyMapperImpl) toModel(in *_imp_1.Technology) _imp_3.Technology {
	out0 := _imp_3.Technology{}
	if in != nil {
		out0.ID = stringToUUIDPtr(in.Id)
		out0.Name = in.Name
		out0.Description = in.Description
	}

	return out0
}
