# Go REST

This is a boilerplate for clean architecture and Go REST projects.

# Architecture

![Architecture](../architecture.drawio.png)

For more information, see [Boilerplate Clean Architecture](../README.md).

## Directories

```
.
├── Makefile
├── api
│   ├── chi-server.config.yaml
│   ├── model.config.yaml
│   └── openapi.yaml
├── build
│   └── dev
│       ├── Dockerfile.e2e
│       ├── Dockerfile.migrate
│       ├── Dockerfile.server
│       └── entrypoint.sh
├── cmd
│   ├── main.go
│   ├── wire.go
│   └── wire_gen.go
├── configs
│   └── mysql
│       └── my.cnf
├── docker-compose.e2e.yml
├── docker-compose.yml
├── go.mod
├── go.sum
├── internal
│   ├── adapter
│   │   ├── controller
│   │   │   ├── model.gen.go
│   │   │   ├── server.gen.go
│   │   │   ├── user_handler.go
│   │   │   └── user_handler_mapper.go
│   │   └── gateway
│   │       ├── db.go
│   │       ├── user_repository_impl.go
│   │       └── user_repository_mapper.go
│   ├── domain
│   │   └── entity
│   │       └── user.go
│   └── usecase
│       ├── user_repository_port.go
│       ├── user_usecase_impl.go
│       ├── user_usecase_mapper.go
│       └── user_usecase_port.go
├── migrations
│   ├── 000001_create_user.down.sql
│   └── 000001_create_user.up.sql
├── pkg
│   └── error
│       └── error.go
└── test
    └── e2e
        └── user_test.go
```

# Get Started

Run the following command. Server will start on http://localhost:8080.

```bash
$ make run
```

# How To Use

## Add user

```bash
$ curl --location 'http://localhost:8080/users' \
  --header 'Content-Type: application/json' \
  --header 'Accept: application/json' \
  --data '{
    "firstName": "test",
    "lastName": "user",
    "age": 20
  }'
```

```json
{
  "age": 20,
  "firstName": "test",
  "id": "<USER_ID>",
  "lastName": "user"
}
```

## Find user

```bash
$ curl --location 'http://localhost:8080/users/<USER_ID>'
```

```json
{
  "age": 20,
  "firstName": "test",
  "id": "<USER_ID>",
  "lastName": "user"
}
```

# How To Development

## API Schema

You need to update `api/openapi.yaml` and run the following command to update OpenAPI schema.

```bash
$ make openapi-gen
```

Generated files are `internal/adapter/controller/model.gen.go` and `internal/adapter/controller/server.gen.go`.

## Dependency Injection

You need to run the following command to update dependency injection.

```bash
$ make wire-gen
```

## Migration

You need to run the following commands to update migration.

```bash
# generate
$ make migrate-create name=<NAME>

# up
$ make migrate-up

# down
$ make migrate-down
```

## Test

### E2E Test

To run E2E test, you need to run the following commands or GitHub Actions Workflow.

```bash
# e2e container up and migration up
$ make e2e-up

# e2e run
$ make e2e-run
```

# References

- [Standard Go Project Layout](https://github.com/golang-standards/project-layout/tree/master)
