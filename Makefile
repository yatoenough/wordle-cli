BINARY_NAME=wordle
DIST_DIR=dist

build:
	@mkdir -p $(DIST_DIR)
	@go build -o $(DIST_DIR)/$(BINARY_NAME) ./cmd/wordle

run:
	@go build -o $(DIST_DIR)/$(BINARY_NAME) ./cmd/wordle
	@./$(DIST_DIR)/$(BINARY_NAME)

clean:
	@rm -rf $(DIST_DIR)

test:
	@go test ./...

.PHONY: build run clean test
