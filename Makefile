BIN_DIR = bin
BIN_NAME = transporter
CMD_PATH = ./main.go

.PHONY: build run test clean

build:
	@go build -o $(BIN_DIR)/$(BIN_NAME) $(CMD_PATH)

run: build
	@./$(BIN_DIR)/$(BIN_NAME)

test:
	@go test -v ./...

clean:
	@rm -rf $(BIN_DIR)/$(BIN_NAME)