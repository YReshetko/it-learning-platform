SVC_NAME=svc-migration

.PHONY: annotations
annotations:
	annotations .

.PHONY: build-migration
build-migration:
	echo $(SVC_NAME)

.PHONY: new-sql
new-sql:
	goose -dir ./db-migration/sql/ create $(sql) sql

.PHONY: new-test-sql
new-test-sql:
	goose -dir ./db-migration/sql/test/ create $(sql) sql
