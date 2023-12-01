docker_psql:
	docker exec -it snippetbox-postgres-1 psql -U myappuser myappdb
migrateup:
	migrate -path db/migrations -database "postgresql://myappuser:myapppassword@localhost:5433/myappdb?sslmode=disable" up
migratedown:
	migrate -path db/migrations -database "postgresql://myappuser:myapppassword@localhost:5433/myappdb?sslmode=disable" down
docker_compose_up:
	docker-compose up -d
docker_compose_down:
	docker-compose down

.PHONY: createdb, dropdb, postgres