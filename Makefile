.PHONY: build
build:
	swag init -g cmd/server/main.go
	go build -o bin/app ./cmd/server/main.go

.PHONY: start
start:
	swag init -g cmd/server/main.go
	go build -o bin/app ./cmd/server/main.go && bin/app

.PHONY: run dev
run-dev:
	go run ./cmd/server/main.go	

.PHONY: clean
clean:
	go clean -cache

.PHONY: swagger
run-swagger:
	swag init -g cmd/server/main.go

.PHONY: migrate database
# Get database URL from .env
DATABASE_URL := $(shell grep DB_URL .env | cut -d '=' -f2)

# Check migrate if not exist
check-go-migrate:
	@command -v migrate >/dev/null 2>&1 || (echo "Migrate not found, installing..." && \
		go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest)

# Migrate up
migrate: check-go-migrate
	@echo "Migrate is running..."
	@echo "DB_URL: $(DATABASE_URL)"
	@migrate -path ./migrations -database $(DATABASE_URL)?sslmode=disable up


# Migrate down
migrate-down: check-go-migrate
	@echo "Migrate is running down..."
	@migrate -path ./migrations -database $(DATABASE_URL)?sslmode=disable down

# Migrate down all
migrate-down-all: check-go-migrate
	@echo "Migrate is running down all..."
	@migrate -path ./migrations -database $(DATABASE_URL)?sslmode=disable down --all