# stegoer - server

- Server is using Go, Postgres and GraphQL.

## Installation

### Install dependencies

```sh
go get ./...
```

### Copy and fill in environment variables

```sh
cp .env.example .env
```

### Initialize database

```sh
createdb stegoer
make migrate
```

## Development

### Dev server

```sh
make dev
```

### Codegen

```sh
make gen
```

### Tests

```sh
make test
```

### Coverage

```sh
make cover
```
