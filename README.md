# Jsask Forum Backend

The backend of my first web application!

## Prerequisites

Ensure you have the following installed:

- [Go](https://go.dev/dl/)
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/)
- [PostgreSQL](https://www.postgresql.org/download/)
- [GNU Make](https://www.gnu.org/software/make/) \*optional

## Setup

Clone the repository and navigate into the project directory:

```sh
git clone https://github.com/tjingsheng/jsask-forum-backend
cd jsask-forum-backend
```

Copy the environment file:

```sh
cp .env.example .env
```

Install Go dependencies:

```sh
go mod tidy
```

## Makefile Commands

The following commands help manage the backend.

### Start Development Server

```sh
make dev
```

### Start Docker Services

```sh
make docker
```

### Run Database Migrations

```sh
make migrate
```

### Seed the Database

```sh
make seed
```

### Reset Database

```sh
make reset
```

### Build the Application

```sh
make build
```

### Run the Compiled Build

```sh
make start
```

## Running Without Makefile

If you don’t want to use `make`, you can run commands manually:

### Start the Server

```sh
go run cmd/server/main.go
```

### Run Migrations

```sh
go run cmd/migrate/main.go
```

### Seed the Database

```sh
go run cmd/seed/main.go
```

### Reset the Database

```sh
docker exec -i jsask-postgres psql -U postgres -d jsask -c "DROP SCHEMA IF EXISTS public CASCADE; CREATE SCHEMA public; GRANT ALL ON SCHEMA public TO postgres; GRANT ALL ON SCHEMA public TO public;"
```

### Build the Application

```sh
go build -o bin/server cmd/server/main.go
```

### Run the Built Binary

```sh
./bin/server
```

---
