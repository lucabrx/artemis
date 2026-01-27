# Artemis

A hybrid project management platform with workspaces, projects, and templates - designed for freelancers and teams.

## Project Structure

```
artemis/
├── apps/
│   ├── api/           # Go backend (Gin, SQLX, PASETO)
│   └── web/           # Svelte frontend (placeholder)
├── docker-compose.yml # Infrastructure services
└── README.md
```

## Tech Stack

### Backend

- **Framework**: Gin
- **Database**: PostgreSQL with pgx/sqlx
- **Cache**: Redis
- **Storage**: MinIO (S3-compatible)
- **Auth**: PASETO tokens (access + refresh)
- **Migrations**: Goose
- **Logging**: zerolog

### Frontend (Coming Soon)

- **Framework**: SvelteKit

## Getting Started

### 1. Start Infrastructure

```bash
docker-compose up -d
```

This starts:

- PostgreSQL on port 5432
- Redis on port 6379
- MinIO on ports 9000 (API) and 9001 (Console)

### 2. Run Migrations

```bash
cd apps/api
go run -tags 'postgres' github.com/pressly/goose/v3/cmd/goose@latest \
  -dir migrations \
  postgres "postgres://artemis:artemis@localhost:5432/artemis?sslmode=disable" up
```

### 3. Start the API

```bash
cd apps/api
cp .env.example .env  # Edit as needed
go run cmd/server/main.go
```

The API will be available at `http://localhost:8080`

## API Endpoints

### Authentication

| Method | Endpoint                | Description                 |
| ------ | ----------------------- | --------------------------- |
| POST   | `/api/v1/auth/register` | Register a new user         |
| POST   | `/api/v1/auth/login`    | Login and get tokens        |
| POST   | `/api/v1/auth/refresh`  | Refresh access token        |
| POST   | `/api/v1/auth/logout`   | Logout (invalidate session) |

### Protected Routes

| Method | Endpoint     | Description      |
| ------ | ------------ | ---------------- |
| GET    | `/api/v1/me` | Get current user |

## Environment Variables

See `apps/api/.env.example` for all configuration options.

## Development

```bash
# Run backend with hot reload (requires air)
cd apps/api
air

# Or run directly
go run cmd/server/main.go
```
