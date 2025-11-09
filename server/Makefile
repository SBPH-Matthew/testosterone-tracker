#!make
include .env

DB_URL=${DB_ENV}

new-migration:
	goose -dir ./migrations create $(NAME) sql

migrate-up:
	goose -dir ./migrations postgres "$(DB_URL)" up

migrate-down:
	goose -dir ./migrations postgres "$(DB_URL)" down

run:
	go run server.go

tidy:
	go mod tidy
