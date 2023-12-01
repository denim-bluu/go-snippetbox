postgres:
	docker run --name postgres12 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:12-alpine  
createdb:
	docker exec -it postgres12 createdb --username=root --owner=root snippet_app
dropdb:
	docker exec -it postgres12 dropdb snippet_app
docker_postgres:
	docker exec -it postgres12 psql -U root
migrateup:
	migrate -path db/migrations -database "postgresql://myappuser:myapppassword@localhost:5432/myappdb?sslmode=disable" up
docker_compose_up:
	docker-compose up -d

.PHONY: createdb, dropdb, postgres