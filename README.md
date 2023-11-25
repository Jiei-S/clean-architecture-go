# Clean Architecture for Go

This is a boilerplate for Clean Architecture.

# Architecture

![Architecture](architecture.drawio.png)

For more information, please refer to the following article.
[Qiita](https://qiita.com/Jiei-S/items/dbf06194f4858519bc61).


## Technology Stack

| Category              | Technology Stack                    |
| --------------------- | ----------------------------------- |
| Frameworks/Libraries  | OpenAPI, chi, wire, bun, go-migrate |
| Environment setup     | Docker                              |


# Development

## Get Started

Run the following command. Server will start on http://localhost:8080.

```bash
$ make dev
```

## How To Use

### Add user

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

### Find user

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

## How To Development

### API Schema

You need to update `api/openapi.yaml` and run the following command to update OpenAPI schema.

```bash
$ make openapi-gen
```

Generated files are `internal/infrastructure/openapi/model.gen.go` and `internal/infrastructure/openapi/server.gen.go`.

### Dependency Injection

You need to run the following command to update dependency injection.

```bash
$ make wire-gen
```

### Migration

You need to run the following commands to update migration.

```bash
# generate
$ make migrate-create name=<NAME>

# up
$ make migrate-up

# down
$ make migrate-down
```

### Test

#### E2E Test

To run E2E test, you need to run the following commands or GitHub Actions Workflow.

```bash
# e2e container up and migration up
$ make e2e-up

# e2e run
$ make e2e-run
```

# References

- [Standard Go Project Layout](https://github.com/golang-standards/project-layout/tree/master)
- [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Clean Architecture 達人に学ぶソフトウェアの構造と設計](https://www.amazon.co.jp/Clean-Architecture-%E9%81%94%E4%BA%BA%E3%81%AB%E5%AD%A6%E3%81%B6%E3%82%BD%E3%83%95%E3%83%88%E3%82%A6%E3%82%A7%E3%82%A2%E3%81%AE%E6%A7%8B%E9%80%A0%E3%81%A8%E8%A8%AD%E8%A8%88-Robert-C-Martin/dp/4048930656)
