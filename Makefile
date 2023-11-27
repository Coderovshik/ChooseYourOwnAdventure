BINARY_NAME = cyoa
.DEFAULT_GOAL = run

run: build
	@./bin/$(BINARY_NAME)

build:
	@GOARCH=amd64 GOOS=windows go build -o ./bin/$(BINARY_NAME)