.PHONY: build
build:
	go build -o ./bin/annotations ./cmd/annotations/main.go

.PHONY: install
install:
	cp ./bin/annotations $(GOBIN)

.PHONY: build-docker
build-docker:
	docker build -t migration-tool:latest -f docker/db/Dockerfile ./docker/db
