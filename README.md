# Quotes Service

A minimal Gin HTTP service that serves random weather, nature, and science facts from a PostgreSQL database.

## Features

- **Gin engine**: with default logging/recovery middleware
- **PostgreSQL**: via GORM ORM with migration support (up/down)
- **Provider interface**: database layer can be swapped by implementing the `Provider` interface
- **Auto-seed**: seeds 149 quotes on first startup (idempotent)
- **Swagger docs**: interactive API docs served at `/swagger/index.html`
- **CORS middleware**: handles preflight requests and sets appropriate headers
- **Graceful shutdown**: handles SIGINT/SIGTERM with 5s timeout
- **Standardized responses**: consistent `APIResponse` format
- **Docker support**: multi-stage Dockerfile and docker-compose with PostgreSQL

## API Endpoints

| Method | Path                     | Auth | Description            |
|--------|--------------------------|------|------------------------|
| GET    | `/`                      | No   | Returns a random quote |
| GET    | `/swagger/index.html`    | No   | Swagger UI             |

## Getting Started

### Prerequisites

- Go 1.24+
- PostgreSQL 16+

### Environment Variables

Copy `.env.example` to `.env` and configure:

```bash
cp .env.example .env
```

| Variable  | Description       | Default     |
|-----------|-------------------|-------------|
| `DB_HOST` | Database host     | `localhost` |
| `DB_PORT` | Database port     | `5432`      |
| `DB_NAME` | Database name     | `postgres`  |
| `DB_USER` | Database user     | `postgres`  |
| `DB_PASS` | Database password | `postgres`  |
| `DB_SSL`  | SSL mode          | `disable`   |
| `PORT`    | Server port       | `3000`      |

### Run with Docker

```bash
docker-compose up --build
```

### Run Locally

```bash
# (Optional) run migrations explicitly — the app also auto-migrates on startup
go run scripts/migrate.go -action up

# Start the server (auto-migrates and seeds on first run)
go run cmd/main.go
```

Then open the Swagger UI at http://localhost:3000/swagger/index.html

### Regenerate Swagger Docs

The `docs/` package is generated from annotations using [swag](https://github.com/swaggo/swag):

```bash
go install github.com/swaggo/swag/cmd/swag@latest
swag init -g cmd/main.go -o docs --parseDependency --parseInternal
```

## CI/CD

`.github/workflows/build-and-publish.yml` runs on every push/PR to `main`/`master`, on `v*` tags, and on manual dispatch:

1. **Build & Test** — `go build`, `go vet`, `go test`.
2. **Publish** — builds the Docker image and pushes it to **GHCR** (`ghcr.io/<owner>/quotesservice`) using the built-in `GITHUB_TOKEN`. Tags are derived from the branch, semver tags, commit SHA, and `latest` on the default branch. Skipped on pull requests.

### Make the image public (one-time)

GHCR images are **private by default**. After the first successful publish:

1. Go to your GitHub profile/org → **Packages** → select the `quotesservice` package.
2. **Package settings** → **Danger Zone** → **Change visibility** → **Public**.

(Optionally connect the package to this repository in the same settings page so it inherits repo access.) Once set, future pushes keep it public.

## Project Structure

```
├── cmd/
│   └── main.go                  # Application entrypoint
├── docs/                        # Generated Swagger docs (swag init)
├── scripts/
│   └── migrate.go               # Database migration CLI
├── internal/
│   ├── app/
│   │   └── app.go               # App bootstrap, migrate, seed, graceful shutdown
│   ├── config/
│   │   └── config.go            # Environment-based configuration
│   ├── models/
│   │   └── quote.go             # GORM data models
│   ├── provider/
│   │   ├── provider.go          # Provider interface
│   │   └── postgres.go          # PostgreSQL implementation
│   ├── seed/
│   │   └── seed.go              # Quote seed data and seeding logic
│   └── api/
│       ├── types/
│       │   └── types.go         # Request/response DTOs
│       ├── controllers/
│       │   └── controller.go    # Route handlers
│       ├── middlewares/
│       │   └── middleware.go    # CORS
│       └── routes/
│           └── routes.go        # Route definitions
├── Dockerfile
├── docker-compose.yml
├── .env.example
└── .gitignore
```
