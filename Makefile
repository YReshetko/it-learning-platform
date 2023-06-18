GO_BUILD_SERVICES = svc-users svc-auth api-app ui-app

.PHONY: go-build
go-build:
	$(foreach service,$(GO_BUILD_SERVICES),go build -o bin/$(service) -v ./$(service)/cmd/main.go &) echo Completed

.PHONY: d-up
d-up:
	docker compose up -d

.PHONY: d-stop
d-stop:
	docker compose stop

.PHONY: d-go-clean
d-go-clean:
	docker rm $(foreach service,$(GO_BUILD_SERVICES),$(service) ) 2> /dev/null || true
	docker rmi $(foreach service,$(GO_BUILD_SERVICES),$(service):latest ) 2> /dev/null || true

.PHONY: d-go-restart
d-go-restart: d-stop d-go-clean d-up

.PHONY: rebuild-%
rebuild-%:
	docker stop $* || true
	docker rm $* || true
	docker rmi $*:latest || true
	docker compose up -d