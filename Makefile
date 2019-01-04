MIGRATE_DSN='user=postgres password=postgres dbname=golangvideos sslmode=disable'
build:
	go build -o bin/golangvideos cmd/golangvideos/main.go
run: build
	./bin/golangvideos
deps:
	go mod tidy
migrate:
	go run cmd/migrate/main.go postgres $(MIGRATE_DSN) up	
recreate-db:
	go run cmd/migrate/main.go postgres $(MIGRATE_DSN) down | true
	go run cmd/migrate/main.go postgres $(MIGRATE_DSN) up
