HOME = /home/$(USER)

PORT ?= 3000

run-api: kill-api
	clear
	@echo "Run development environment..."
	go run cmd/api/main.go cmd/api/config.go

kill-api:
	clear
	@echo "Kill api..."
	kill -9 $$(lsof -t -i:$(PORT)) || true

create-migration:
	clear
	@if [ -z "$(name)" ]; then \
		echo "Usage: make create-migration name=<migration_name>"; \
		exit 1; \
	fi
	migrate create -ext sql -dir db/migrations $(name)

run-migration:
	clear
	echo "Running migration" 
	migrate -database "$$DSN_MIGRATE" -path db/migrations up

down-migration:
	clear
	echo "Running migration" 
	migrate -database "$$DSN_MIGRATE" -path db/migrations down 1

force-version:
	clear
	@if [ -z "$(version)" ]; then \
		echo "Usage: make force-version version=<migration_version>"; \
		exit 1; \
	fi
	migrate -database "$$DSN_MIGRATE" -path db/migrations force $(version)
