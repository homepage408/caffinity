include .env
export

# ========================
# RUN APP
# ========================
run:
	go run cmd/main.go

build:
	go build -o bin/$(APP_NAME) cmd/main.go

start: build
	./bin/$(APP_NAME)

# ========================
# DATABASE
# ========================
createdb:
	psql -U postgres -c "CREATE DATABASE caffinity;"

dropdb:
	psql -U postgres -c "DROP DATABASE IF EXISTS caffinity;"

resetdb: dropdb createdb

# ========================
# MIGRATION (golang-migrate)
# ========================
migrate-up:
	migrate -path db/migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path db/migrations -database "$(DB_URL)" down

migrate-force:
	migrate -path db/migrations -database "$(DB_URL)" force 1

# ========================
# SQLC
# ========================
sqlc:
	sqlc generate

# ========================
# DEV
# ========================
dev:
	air

# ========================
# CLEAN
# ========================
clean:
	rm -rf bin/