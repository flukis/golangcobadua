build: 
	@go build -o bin/jokes

run: build
	./bin/jokes

test: 
	go test -v ./..