SVC_NAME=svc-auth

.PHONY: annotations
annotations:
	annotations .

.PHONY: gen-proto
gen-proto:
	protoc --go_out=. \
	--go_opt=Mproto/auth.proto=./pb/auth \
	--go-grpc_out=. \
	--go-grpc_opt=Mproto/auth.proto=./pb/auth \
	proto/auth.proto
