run: build
	@./bin/api

build:
	@go build -o ./bin/api

test:
	@go test -v ./...
	
templ:
	@templ generate -watch -proxy=http://localhost:8080

tailwind:
	@tailwindcss -i views/css/styles.css -o public/styles.css --watch

migration: # add migration name at the end (ex: make migration create-jobs-table)
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down
