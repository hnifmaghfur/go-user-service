.PHONY: build
build:
	go build -o bin/app ./cmd/server/main.go

.PHONY: start
start:
	go build -o bin/app ./cmd/server/main.go && bin/app

.PHONY: run dev
dev:
	go run ./cmd/server/main.go	
