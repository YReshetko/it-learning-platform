// Code generated by Constructor annotation processor. DO NOT EDIT.
// versions:
//		go: go1.20.4
//		go-annotation: 0.1.0
//		Constructor: 1.0.0

package grpc

import (
	storage "github.com/YReshetko/it-learning-platform/svc-courses/internal/storage"
	logrus "github.com/sirupsen/logrus"
)

type HandlerOption func(*Handler)

func NewHandler(opts ...HandlerOption) Handler {
	rt := &Handler{}
	for _, o := range opts {
		o(rt)
	}

	return *rt
}

func WithCategoryMapper(v CategoryMapper) HandlerOption {
	return func(rt *Handler) {
		rt.categoryMapper = v
	}
}

func WithCourseMapper(v CourseMapper) HandlerOption {
	return func(rt *Handler) {
		rt.courseMapper = v
	}
}

func WithLogger(v *logrus.Entry) HandlerOption {
	return func(rt *Handler) {
		rt.logger = v
	}
}

func WithStorage(v storage.CourseStorage) HandlerOption {
	return func(rt *Handler) {
		rt.storage = v
	}
}

func WithTaskMapper(v TaskMapper) HandlerOption {
	return func(rt *Handler) {
		rt.taskMapper = v
	}
}

func WithTechnologyMapper(v TechnologyMapper) HandlerOption {
	return func(rt *Handler) {
		rt.technologyMapper = v
	}
}

func WithTopicMapper(v TopicMapper) HandlerOption {
	return func(rt *Handler) {
		rt.topicMapper = v
	}
}
