include .env
export

migrate-up:
	docker-compose run --rm migrate up

migrate-down:
	docker-compose run --rm migrate down 1

migrate-create:
	docker run -v $(shell pwd)/internal/infra/database/migrations:/migrations migrate/migrate create -ext sql -dir /migrations/ -seq $(name)

migrate-force:
	docker-compose run --rm migrate force $(version)
