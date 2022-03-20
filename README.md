# stegoer - server

Server is using Go, Postgres, GraphQL and Redis.

---

Server endpoint: https://stegoer-server.herokuapp.com/

---

## Installation

### Install dependencies

```sh
go get ./...
```

### Create the `.env` file

Create a `.env` file and copy the contents of `.env.example` file into the `.env` file

```sh
cp .env.example .env
```

### Initialize database

```sh
make db-init
```

## Development

### Dev server

```sh
make dev
```

### Redis server

```sh
redis-server
```

### Make

```sh
make help
```
