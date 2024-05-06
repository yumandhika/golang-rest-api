build:
	@go build -o yumandhika/golang-rest-api cmd/main.go

test:
	@go test -v ./...
	
run: build
	@./yumandhika/golang-rest-api

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down