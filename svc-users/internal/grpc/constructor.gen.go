// Code generated by Constructor annotation processor. DO NOT EDIT.
// versions:
//		go: go1.20.4
//		go-annotation: 0.1.0
//		Constructor: 1.0.0

package grpc

import (
	config "github.com/YReshetko/it-academy-cources/svc-users/internal/config"
	storage "github.com/YReshetko/it-academy-cources/svc-users/internal/storage"
)

func NewHandler(storage storage.UserStorage) Handler {
	returnValue := Handler{
		storage: storage,
	}

	return returnValue
}

func NewServer(cfg config.GRPC, handler Handler) Server {
	returnValue := Server{
		cfg:     cfg,
		handler: handler,
	}

	return returnValue
}
