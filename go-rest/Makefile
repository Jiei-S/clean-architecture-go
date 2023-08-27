# server
.PHONY: openapi-gen
openapi-gen:
	oapi-codegen -config api/chi-server.config.yaml api/openapi.yaml
	oapi-codegen -config api/model.config.yaml api/openapi.yaml

.PHONY: wire-gen
wire-gen:
	wire cmd/wire.go

.PHONY: run
run:
	docker compose down || true
	docker network create go-rest || true
	docker compose up server db --build

# migration
.PHONY: migrate-create
migrate-create:
	docker compose run migrate create -ext sql -dir /internal/infrastructure/bun/migrations -seq $(name)

.PHONY: migrate-up
migrate-up:
	docker compose run migrate up

.PHONY: migrate-down
migrate-down:
	docker compose run migrate down

# e2e
.PHONY: e2e-up
e2e-up:
	docker compose -f docker-compose.e2e.yml down || true
	docker network create go-rest || true
	docker compose -f docker-compose.e2e.yml up --build -d

.PHONY: e2e-run
e2e-run:
	docker compose -f docker-compose.e2e.yml exec server.e2e go test -v ./test/e2e/...

.PHONY: e2e-down
e2e-down:
	docker compose -f docker-compose.e2e.yml down -v || true