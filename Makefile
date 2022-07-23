include .env

run:
	go run .

migrateup:
	migrate -database ${DSN} -path migration up

migratedown:
	migrate -database ${DSN} -path migration down
