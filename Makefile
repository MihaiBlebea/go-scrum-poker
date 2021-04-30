setup: build-server

local: build up

build-server:
	go build -o=./scrum-poker . && ./scrum-poker start

build-webapp:
	cd ./webapp && npm run build && cp -r ./dist/* ./../server/webapp/

watch-webapp:
	cd ./webapp && npm run dev

build:
	docker-compose build

up:
	docker-compose up -d

refresh:
	docker-compose stop &&\
	docker-compose rm &&\
	rm -r ./volume &&\
	make local