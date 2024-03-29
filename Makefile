.PHONY:
run:
	go run cmd/app/main.go

migrateup:
	migrate -path pkg/database/postgres/migration -database "postgresql://postgres:marta2010@localhost:5432/databaselab3" -verbose up

migratefix:
	migrate -path pkg/database/postgres/migration/ -database "postgresql://postgres:marta2010@localhost:5432/databaselab3" force 1