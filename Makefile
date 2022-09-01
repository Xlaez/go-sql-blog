postgres: 
	docker run --name postgres13 -p 5000:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres\:13.8-alpine
createdb:
	docker exec -it postgres13 createdb --username=root --owner=root test
dropdb:
	docker exec -it postgres13 dropdb test
migrateup: 
	migrate -path ./db/migration/ -database "postgresql://root:password@localhost:5000/test?sslmode=disable" -verbose up
migratedown:
	migrate -path ./db/migration/ -database "postgresql://root:password@localhost:5000/test?sslmode=disable" -verbose down
initmigrate:
	migrate create -ext sql -dir db/migration -seq add_user
test:
	go test -v -cover ./...
sqlc:
	sqlc generate
serve:
	go run main.go
opendb:
	docker exec -it postgres13-old psql -U root -d bank
.PHONY: 
	postgres createdb dropdb migrateup migratedown sqlc test serve opendb initmigrate