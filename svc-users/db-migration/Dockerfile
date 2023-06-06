FROM golang:1.20-alpine
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
COPY ./sql/ /sql/
COPY ./sh/run.sh /sql/run.sh

RUN chmod 777 /sql/run.sh

WORKDIR /sql
ENTRYPOINT ["./run.sh"]
