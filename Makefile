.PHONY: build run docker-build docker-run docker-compose-up docker-compose-down

build:
	go build -o bin/main ./cmd/userdata

run:
	go run ./cmd/userdata

docker-build:
	docker build -t userdata-app .

docker-run:
	docker run -p 8080:8080 userdata-app

docker-compose-up:
	docker-compose up --build

docker-compose-down:
	docker-compose down