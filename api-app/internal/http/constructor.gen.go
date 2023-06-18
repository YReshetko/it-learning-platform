// Code generated by Constructor annotation processor. DO NOT EDIT.
// versions:
//		go: go1.20.4
//		go-annotation: 0.1.0
//		Constructor: 1.0.0

package http

import (
	config "github.com/YReshetko/it-learning-platform/api-app/internal/config"
	logrus "github.com/sirupsen/logrus"
)

func NewServer(cfg config.HTTP, logger *logrus.Entry) Server {
	returnValue := Server{
		cfg:    cfg,
		logger: logger,
	}
	returnValue.postConstruct()

	return returnValue
}
