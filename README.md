# AltTube-Go

## Getting Started

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

### Local environment

Change `JWT_SECRET` in `.env` file

```sh
NEW_SECRET=$(openssl rand -base64 32)
sed -i '' "s/JWT_SECRET=.*/JWT_SECRET=$NEW_SECRET/" .env
```

Run

```sh
go mod download
go run main.go
```

### Docker environment

Change `JWT_SECRET` in `.env` file

```sh
NEW_SECRET=$(openssl rand -base64 32)
sed -i '' "s/JWT_SECRET=.*/JWT_SECRET=$NEW_SECRET/" .env.docker
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
