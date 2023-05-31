include .env

MIGRATE_SCHEME?="up"

run:
	go run ./src/cmd/app.go ./src/cmd/main.go

migrate:
	./src/docs/sql/migrate -database ${DATA_SOURCE} -path "./src/docs/sql/migration" ${MIGRATE_SCHEME} 
