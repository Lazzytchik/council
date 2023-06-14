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

serve:
	APP_DIR="${PWD}" bin/serve.sh