SVC_NAME=svc-courses

.PHONY: annotations
annotations:
	annotations .

.PHONY: new-sql
new-sql:
	goose -dir ./db-migration/sql/ create $(sql) sql

.PHONY: new-test-sql
new-test-sql:
	goose -dir ./db-migration/sql/test/ create $(sql) sql

.PHONY: gen-proto
gen-proto:
	protoc --go_out=. \
	--go_opt=Mproto/courses.proto=./pb/courses \
	--go-grpc_out=. \
	--go-grpc_opt=Mproto/courses.proto=./pb/courses \
	proto/courses.proto

