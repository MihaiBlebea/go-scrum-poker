setup: build-server

local: build up

build-server:
	go build -o=./scrum-poker . && ./scrum-poker start

build-webapp:
	cd ./webapp && npm run build

build:
	docker-compose build

up:
	docker-compose up -d