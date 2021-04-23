setup: build-server

local: build up

build-server:
	go build -o=./scrum-poker . && ./scrum-poker start

build-webapp:
	cd ./webapp && npm run build

watch-webapp:
	cd ./webapp && npm run dev

build:
	docker-compose build

up:
	docker-compose up -d