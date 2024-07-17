# Docker

- Set up Postgre

```
docker pull postgres:16.3-alpine
docker run --name postgre16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123456 -d postgres:16.3-alpine
docker exec -it postgre16 psql -U root

services.msc (stop postgre)
```

# Golang Migration

- https://github.com/golang-migrate/migrate?tab=readme-ov-file

```
// in Folder Project
migrate create -ext sql -dir db/migration -seq init_schema
docker exec -it postgre16 createdb --username=root --owner=root simple_bank
docker exec -it postgre16 psql -U root simple_bank
```

# Swagger
- Lưu ý: Phải thêm docs !!
~~~
swag init -g ./cmd/server/main.go -o cmd/server/docs
~~~
