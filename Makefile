include .env

MIGRATE_SCHEME?="up"

run:
	go run ./src/cmd/app.go ./src/cmd/main.go

migrate:
	./src/docs/sql/migrate -database ${DATA_SOURCE} -path "./src/docs/sql/migration" ${MIGRATE_SCHEME} 

migrate-generate:
	./src/docs/sql/migrate create -ext sql -dir "./src/docs/sql/migration" -seq ${MIGRATE_NAME}

seed:
	go run ./src/docs/sql/seeder/main seed ${SEED_NAME}
