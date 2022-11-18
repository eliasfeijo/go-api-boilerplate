generate-swagger:
	swag i -g server/web.go

build:generate-swagger
	go build -o api

up:
	docker-compose up -d --build

up-db:
	docker-compose up -d --build db

down:
	docker-compose down

logs:
	docker-compose logs -f