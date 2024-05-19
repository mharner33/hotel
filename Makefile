build:
	@go build -o bin/api cmd/*.go

run: build
	@go run cmd/*.go

test:
	@go test -v ./...

clean:
	@rm bin/api/*