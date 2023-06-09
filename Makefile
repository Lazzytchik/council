-include .env
export $(shell sed 's/=.*//' .env)

.PHONY: "up"
up:
	docker-compose up

.PHONY: "migrate-up"
migrate-up:
	goose -dir="./migrations" postgres "host=${POSTGRES_HOST} user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} dbname=${POSTGRES_DB} sslmode=disable" up

.PHONY: "migrate-reset"
migrate-reset:
	goose -dir="./migrations" postgres "host=${POSTGRES_HOST} user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} dbname=${POSTGRES_DB} sslmode=disable" reset

serve: migrate-up
	APP_DIR="${PWD}" bin/serve.sh

buf-gen:
	protoc -I ./proto \
       --go_out ./grpc --go_opt paths=source_relative \
       --go-grpc_out ./grpc --go-grpc_opt paths=source_relative \
       --grpc-gateway_out ./grpc --grpc-gateway_opt paths=source_relative \
       ./proto/auth/user.proto