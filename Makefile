build:
	@go build -o bin/go-api cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/go-api
migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS)) 
migration-up:
	@go run cmd/migrate/main.go up 
migration-down:
	@go run cmd/migrate/main.go down 
migration-redo:
	@go run cmd/migrate/main.go redo 
