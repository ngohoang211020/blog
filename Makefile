## help: print this help message
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]
## run/api: run the cmd/api application
cmd/api:
	@go run ./cmd/api
## db/psql: connect to the database using psql
psql:
	psql -U postgres -d blog -p 5432 -h localhost
## db/migrations/up: apply all up database migrations
migrations/up: confirm
	@migrate -path migrations -database postgresql://postgres:123456@localhost:5432/blog?sslmode=disable up
## db/migrations/down: apply all up database migrations
migrations/down: confirm
	@migrate -path migrations -database postgresql://postgres:123456@localhost:5432/blog?sslmode=disable down
## db/migrations/new name=$1: create a new database migration
migrations/new:
	@echo 'Creating migration files for ${name}...'
	migrate create -seq -ext=.sql -dir=./migrations ${name}
.PHONY: migrations/new psql help migrations/up migrations/down