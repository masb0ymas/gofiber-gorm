include .env

BUILD_DIR = $(PWD)/build

swagger:
	swag init -g main.go

dev: swagger
	./bin/air server --port $(APP_PORT)

clean:
	rm -rf ./build

security:
	gosec ./...

build: clean
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) main.go

run: build
	$(BUILD_DIR)/$(APP_NAME)

