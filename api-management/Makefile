SVC_NAME=api-management

.PHONY: annotations
annotations:
	annotations .

.PHONY: build
build:
	go build -o ./bin/$(SVC_NAME) ./cmd/main.go

.PHONY: run
run:
	go run ./cmd/main.go

