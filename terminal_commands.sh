docker pull postgres
docker ps
docker run --name=todo-db -e POSTGRES_PASSWORD='qwerty' -p 5432:5432 -d --rm postgres
docker ps
docker exec -it 45e13b7d7e31 /bin/bash
migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' down
docker exec -it 45e13b7d7e31 /bin/bash  