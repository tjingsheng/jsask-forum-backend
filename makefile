.PHONY: docker migrate

docker:
	docker compose --profile development --env-file .env up --remove-orphans --detach

migrate:
	go run cmd/migrate/main.go