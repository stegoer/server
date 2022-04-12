# server

[![stegoer-server deployment](https://img.shields.io/github/deployments/stegoer/server/stegoer-server?label=heroku&logo=heroku&logoColor=heroku)](https://stegoer-server.herokuapp.com/)
[![Continuous Integration](https://github.com/stegoer/server/actions/workflows/ci.yml/badge.svg)](https://github.com/stegoer/server/actions/workflows/ci.yml)
[![pre-commit.ci status](https://results.pre-commit.ci/badge/github/stegoer/server/main.svg)](https://results.pre-commit.ci/latest/github/stegoer/server/main)
[![Go Reference](https://pkg.go.dev/badge/github.com/stegoer/server.svg)](https://pkg.go.dev/github.com/stegoer/server)
[![Go Report Card](https://goreportcard.com/badge/github.com/stegoer/server)](https://goreportcard.com/report/github.com/stegoer/server)

Server is using Go, Postgres, GraphQL and Redis.

---

Server endpoint: https://stegoer-server.herokuapp.com/

Development documentation: https://github.com/stegoer/server/blob/main/README.md

Reference documentation: https://stegoer.github.io/client

Source code: https://github.com/stegoer/server

---

## Installation

### Install instructions

1. Install Go https://go.dev/doc/install
2. Install PostgreSQL https://www.postgresql.org/download/
3. Install Redis https://redis.io/docs/getting-started/
4. Clone this repository

```sh
git clone git@github.com:stegoer/server.git
```

5. Install dependencies

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

## Contributing

```sh
pre-commit install
```

## License

Developed under the [MIT](https://github.com/stegoer/server/blob/master/LICENSE) license.
