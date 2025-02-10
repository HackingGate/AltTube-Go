# AltTube-Go

[![Lint](https://github.com/HackingGate/AltTube-Go/actions/workflows/lint.yml/badge.svg)](https://github.com/HackingGate/AltTube-Go/actions/workflows/lint.yml)
[![golangci-lint](https://github.com/HackingGate/AltTube-Go/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/HackingGate/AltTube-Go/actions/workflows/golangci-lint.yml)
[![Atlas](https://github.com/HackingGate/AltTube-Go/actions/workflows/ci-atlas.yaml/badge.svg)](https://github.com/HackingGate/AltTube-Go/actions/workflows/ci-atlas.yaml)
[![Build and Test](https://github.com/HackingGate/AltTube-Go/actions/workflows/build-and-test.yml/badge.svg)](https://github.com/HackingGate/AltTube-Go/actions/workflows/build-and-test.yml)
[![CodeQL](https://github.com/HackingGate/AltTube-Go/actions/workflows/github-code-scanning/codeql/badge.svg)](https://github.com/HackingGate/AltTube-Go/actions/workflows/github-code-scanning/codeql)
[![Dependabot Updates](https://github.com/HackingGate/AltTube-Go/actions/workflows/dependabot/dependabot-updates/badge.svg)](https://github.com/HackingGate/AltTube-Go/actions/workflows/dependabot/dependabot-updates)

## Getting Started

### Common steps

Generate `DB_PASSWORD` and `JWT_KEY` in `.env` and `.env.docker` file

```sh
python3 generate_credentials.py
```

Remove `data/` directory after generating `DB_PASSWORD`

```sh
rm -rf data/
```

### Local environment

Build and run

```sh
docker compose up -d --build alttube-postgres piped piped-postgres piped-proxy
go mod download
go run main.go
```

### Docker environment

Build and run

```sh
docker compose build
docker compose up
```

Build and run in background in one command

```sh
docker compose up --build -d
```

## Database Development

#### Dependencies

- [ent](https://entgo.io)
- [atlas](https://atlasgo.io)

#### Generate ent if ent/schema is updated

```sh
go generate ./ent
```

#### Generate Migration SQL

```sh
atlas migrate diff add_something_awesome \
  --dir "file://ent/migrate/migrations" \
  --to "ent://ent/schema" \
  --dev-url "postgres://AltTube:AltTube@localhost:5432/AltTube?search_path=public&sslmode=disable"
```

#### Apply Migration SQL

```sh
atlas migrate apply \
  --dir "file://ent/migrate/migrations" \
  -u "postgres://AltTube:AltTube@localhost:5432/AltTube?search_path=public&sslmode=disable"
```

#### Visualize Database Schema

```sh
atlas schema inspect -w \
  -u "postgres://AltTube:AltTube@localhost:5432/AltTube?search_path=public&sslmode=disable"
```

## API Documentation

#### Generate

Install [swag](https://github.com/swaggo/swag)

```sh
go install github.com/swaggo/swag/cmd/swag@latest
```

At the root of the project, run

```sh
swag init --parseDependency --parseInternal --parseDepth 1
```

#### Access

http://localhost:8072/swagger/index.html

https://efficiency-node-alttube.hackinggate.com/swagger/index.html
