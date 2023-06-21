#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/goose /goose
COPY ./sh/run.sh /run.sh
RUN ["chmod", "+x", "/run.sh"]
ENTRYPOINT ["/run.sh"]
LABEL Name=migration-tool Version=0.0.1