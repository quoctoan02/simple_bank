DB_URL=postgresql://root:root@localhost:5432/simple_bank?sslmode=disable

postgres:
	docker run --name postgres16 -network bank-network -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -p 5432:5432 -d postgres:16-alpine

createdb:
	docker exec -it postgres16 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres16 dropdb simple_bank

# migrateup:
# 	migrate -path db/migration -database "postgresql://root:9ncIPDeNFzd2Vst7jRDd@simple-bank.c1kyqseqe2qi.ap-southeast-1.rds.amazonaws.com:5432/simple_bank" -verbose up

# migrateup1:
# 	migrate -path db/migration -database "postgresql://root:9ncIPDeNFzd2Vst7jRDd@simple-bank.c1kyqseqe2qi.ap-southeast-1.rds.amazonaws.com:5432/simple_bank" -verbose up 1

# migratedown:
# 	migrate -path db/migration -database "postgresql://root:9ncIPDeNFzd2Vst7jRDd@simple-bank.c1kyqseqe2qi.ap-southeast-1.rds.amazonaws.com:5432/simple_bank" -verbose down

# migratedown1:
# 	migrate -path db/migration -database "postgresql://root:9ncIPDeNFzd2Vst7jRDd@simple-bank.c1kyqseqe2qi.ap-southeast-1.rds.amazonaws.com:5432/simple_bank" -verbose down 1

migrateup:
	migrate -path db/migration -database "${DB}" -verbose up

migrateup1:
	migrate -path db/migration -database "${DB}" -verbose up 1

migratedown:
	migrate -path db/migration -database "${DB}" -verbose down

migratedown1:
	migrate -path db/migration -database "${DB}" -verbose down 1

db_docs:
	dbdocs build doc/db.dbml

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go simple_bank/db/sqlc Store

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

.PHONY: postgres createdb dropdb migratedown migrateup sqlc test server mock migratedown1 migrateup1 new_migration db_docs db_schema 

