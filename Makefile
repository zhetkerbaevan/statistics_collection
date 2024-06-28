build:
	@go build -o bin/e-commerce cmd/e-commerce/main.go

test:
	@go test -v ./...

run: build
	@./bin/e-commerce