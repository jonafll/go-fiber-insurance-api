.PHONY: clean test build

clean:
	rm -rf ./build

test:
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out

build: clean test
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) main.go

run: swag
	go run main.go

docker.build:
	docker build -t insurance-api .

docker.run: docker.build
	docker run --rm -d \
		--name insurance-api \
		-p 8080:8080 \
		insurance-api

docker.stop:
	docker stop insurance-api

swag:
	swag init