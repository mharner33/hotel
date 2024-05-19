build:
	@go build -o bin/api cmd/web/*.go

run: build
	@go run cmd/web/*.go

test:
	@go test -v ./...

clean:
	@rm bin/api/*