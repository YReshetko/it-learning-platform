#!/bin/sh

sleep 10
./goose -dir ./sql postgres "postgres://$USER_NAME:$USER_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME" up