all: run

run:  build migrate compose

migrate:
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up

build:
	docker build -t todo-app .

compose:
	docker-compose up todo-app