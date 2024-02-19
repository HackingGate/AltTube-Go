# AltTube-Go

## Getting Started

### Local environment

Generate `JWT_KEY` and `DB_PASSWORD` in `.env` file

```sh
python3 generate_credentials.py .env
```

Remove `data/` directory after generating `DB_PASSWORD`

```sh
rm -rf data/
```

Source `.env` file

```sh
. ./.env
```

Run

```sh
docker compose up -d --build postgres piped piped-postgres piped-proxy
go mod download
go run main.go
```

### Docker environment

Generate `JWT_KEY` and `DB_PASSWORD` in `.env.docker` file

```sh
python3 generate_credentials.py .env.docker
```

Remove `data/` directory after generating `DB_PASSWORD`

```sh
rm -rf data/
```

Source `.env.docker` file

```sh
. ./.env.docker
```

Build and run

```sh
docker compose build
docker compose up
```

Build and run in background in one command

```sh
docker compose up --build -d
```

### API Documentation

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
