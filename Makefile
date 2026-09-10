.PHONY: test generate build clean

test:
	@set -e; \
	trap 'docker compose down -v --remove-orphans' EXIT; \
	sqlc generate; \
	go build ./...; \
	docker compose down -v --remove-orphans; \
	docker compose up -d; \
	echo "Esperando a PostgreSQL..."; \
	until docker compose exec -T postgres pg_isready -U tp_user -d tp_db > /dev/null 2>&1; do \
		sleep 1; \
	done; \
	echo "PostgreSQL está listo."; \
	go test ./...

generate:
	sqlc generate

build:
	go build ./...

clean:
	docker compose down -v --remove-orphans