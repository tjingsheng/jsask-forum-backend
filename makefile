.PHONY: dev docker migrate seed reset build-windows build-unix build

dev: 
	go run cmd/server/main.go

docker:
	docker compose --profile development --env-file .env up --remove-orphans --detach

migrate:
	go run cmd/migrate/main.go

seed:
	go run cmd/seed/main.go

reset: 
	docker exec -i jsask-postgres psql -U postgres -d jsask -c "DROP SCHEMA IF EXISTS public CASCADE; CREATE SCHEMA public; GRANT ALL ON SCHEMA public TO postgres; GRANT ALL ON SCHEMA public TO public;"

build-windows:
	GOOS=windows GOARCH=amd64 go build -o bin/server.exe cmd/server/main.go 
	
build-unix:
	go build -o bin/server cmd/server/main.go

build: build-windows build-unix

start:
	./bin/server
