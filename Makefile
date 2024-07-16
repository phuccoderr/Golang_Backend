postgre:
	docker run --name postgre16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123456 -d postgres:16.3-alpine
createdb:
	docker exec -it postgre16 createdb --username=root --owner=root simple_bank
dropdb:
	docker exec -it postgre16 dropdb simple_bank

migrate_up:
	migrate -path db/migration -database "postgresql://root:123456@localhost:5432/simple_bank?sslmode=disable" -verbose up
migrate_down:
	migrate -path db/migration -database "postgresql://root:123456@localhost:5432/simple_bank?sslmode=disable" -verbose down

.PHONY: postgre createdb dropdb migrate_up migrate_down