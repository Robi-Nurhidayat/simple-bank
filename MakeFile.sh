# run postgres
docker run -p 5432:5432 --name postgres -e POSTGRES_PASSWORD=secret -e POSTGRES_USER=root  -d postgres:16-alpine

#Createdb:
docker exec -it postgres createdb --username=root --owner=root simple_bank2

#dropdb
docker exec -it postgres dropdb simple_bank2

# masuk ke database
docker exec -it postgres psql -U root simple_bank2

# migrate up
migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/simple_bank2?sslmode=disable" -verbose up

# migrate down
migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/simple_bank2?sslmode=disable" -verbose down


# auth token simple_bank
sqlc_01HJMGGE5CBQRSPVPD3Y6HPYAK