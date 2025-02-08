# AltTube-Go

[![Lint](https://github.com/HackingGate/AltTube-Go/actions/workflows/lint.yml/badge.svg)](https://github.com/HackingGate/AltTube-Go/actions/workflows/lint.yml)
[![golangci-lint](https://github.com/HackingGate/AltTube-Go/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/HackingGate/AltTube-Go/actions/workflows/golangci-lint.yml)
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

## Database Diagram

![DB_Diagram](https://github.com/HackingGate/AltTube-Go/assets/8541644/d5eee81d-75be-489c-8db9-91b0a054b642)
