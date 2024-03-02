run: build
	@./bin/api

build:
	@go build -o ./bin/api

test:
	@go test -v ./...
	
templ:
	@templ generate -watch -proxy=http://localhost:8080
