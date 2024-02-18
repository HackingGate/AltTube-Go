# AltTube-Go

## Getting Started

### Generate Docuemntation

At the root of the project, run

```sh
swag init --parseDependency --parseInternal --parseDepth 1
```

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
