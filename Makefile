.PHONY: "up"
up:
	docker-compose up

.PHONY: "migrate-up"
migrate-up:
	goose -dir="./migrations" postgres "host=localhost user=postgres password=postgres dbname=users_pg sslmode=disable" up

.PHONY: "migrate-reset"
migrate-reset:
	goose -dir="./migrations" postgres "host=localhost user=postgres password=postgres dbname=users_pg sslmode=disable" reset

serve:
	go run main.go