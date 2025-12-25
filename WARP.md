# WARP.md

This file provides guidance to WARP (warp.dev) when working with code in this repository.

## Project Overview

Vago is a Golang learning portal (Russian language) that provides:
- A book for learning theoretical content about Go
- Quiz system for testing knowledge
- General chat for discussions
- Personal task tracker

The application runs on Docker and includes a full stack: REST API, gRPC/gRPC-Web for chat streaming, PostgreSQL database, web UI using Gin templates, and optional Kafka integration.

Demo: http://vadmark.duckdns.org/

## Tech Stack

- **Backend Framework**: Gin (HTTP), gRPC + gRPC-Web (chat streaming)
- **Database**: PostgreSQL with manual SQL migrations (Goose was removed)
- **ORM**: GORM
- **Auth**: JWT (access + refresh tokens)
- **Logging**: Zap
- **Frontend**: Server-side rendered Gin templates with vanilla JavaScript
- **Message Queue**: Kafka (optional, controlled by `KAFKA_ENABLE` env var)
- **Protobuf**: Protocol Buffers for gRPC communication
- **Frontend Build**: esbuild for bundling JavaScript gRPC clients

## Common Commands

### Docker Operations
```bash
make up          # Start all containers (postgres + app, +kafka if enabled)
make down        # Stop all containers
make down-v      # Stop all containers and remove volumes
make ps          # Show running containers
make logs        # Show last 20 lines of vago container logs
make logs-f      # Follow logs from vago container
make psql        # Open PostgreSQL shell (container: vago-postgres, db: vagodb)
```

### Building and Deployment
```bash
make build       # Build Docker image ghcr.io/vadmark223/vago:latest
make push        # Push image to GitHub Container Registry
make pull        # Pull image from GHCR
```

### gRPC and Protocol Buffers
```bash
make proto-go       # Generate Go gRPC files from .proto definitions in api/proto/
make proto-js-clean # Remove all generated JavaScript files from web/static/js/pb/
make proto-js       # Generate gRPC-Web JavaScript client files
make bundle         # Bundle JavaScript client into web/static/js/bundle.js using esbuild
make proto-js-all   # Full pipeline: clean → generate → bundle (use this for frontend changes)
```

### Kafka (Optional)
```bash
make kafka-up    # Start kafka and kafka UI containers
make kafka-down  # Stop kafka containers
```

### Cleanup
```bash
make clean-all   # ⚠️ Remove ALL Docker resources (containers, images, volumes, networks)
```

## Project Structure

### High-Level Architecture

The codebase follows a layered architecture pattern:

**Entry Points**:
- `cmd/vago/main.go` - Main application entry point with graceful shutdown
- `cmd/tasks/` - Separate executable for task-related utilities/testing

**Core Layers**:
1. **Transport Layer** (`internal/trasport/`) - HTTP/gRPC/WebSocket handlers
   - `http/` - Gin router, handlers, middleware
   - `grpc/` - gRPC server implementation
   - `ws/` - WebSocket hub for real-time chat (Gorilla)

2. **Application Layer** (`internal/application/`) - Business logic services
   - `user/` - User management service
   - `chat/` - Message and chat service
   - `quiz/` - Quiz and question management
   - `task/` - Task tracker service
   - `topic/` - Learning topics service

3. **Domain Layer** (`internal/domain/`) - Domain entities and interfaces
   - Pure domain models (User, Task, Message, Question, etc.)
   - Repository interfaces

4. **Infrastructure Layer** (`internal/infra/`) - External dependencies
   - `persistence/gorm/` - GORM repository implementations (entities + repos)
   - `db/` - Database connection
   - `token/` - JWT provider
   - `logger/` - Zap logger setup
   - `crypto/` - Password hashing utilities
   - `kafka/` - Kafka consumer implementation

5. **Configuration** (`internal/config/`)
   - `config/` - Application config loading from env vars
   - `code/` - String constants used throughout app
   - `route/` - Route path constants
   - `kafka/topic/` - Kafka topic definitions

6. **Context** (`internal/app/`)
   - `context.go` - Application context containing logger, DB, config
   - `localCache.go` - In-memory cache for user data
   - `debug.go` - Debug utilities

### Database Migrations

Manual SQL migrations are located in `db/`:
- `01_init.sql` - Creates users, tasks, and messages tables
- `02_quiz.sql` - Quiz-related tables
- `03_admin.sql` - Admin user seeding

These scripts run automatically when the postgres container starts (mounted to `/docker-entrypoint-initdb.d`).

### Frontend Structure

- `web/templates/` - Go HTML templates (`.html`, `.gohtml`)
- `web/static/` - Static assets (CSS, JS, images)
- `web/static/js/pb/` - Generated gRPC-Web client code (gitignored, regenerated)
- `web/static/js/bundle.js` - Bundled JavaScript output from esbuild

### Protocol Buffers

Proto definitions in `api/proto/`:
- `auth.proto` - Authentication messages
- `chat.proto` - Chat streaming protocol
- `hello.proto` - Hello world example
- `ping.proto` - Health check

Generated Go code goes to `api/pb/<service>/`

## Development Workflow

### Running Locally

1. Copy and configure environment:
   ```bash
   cp .env.prod .env.local
   # Edit .env.local with local values
   export APP_ENV=local
   ```

2. Start services:
   ```bash
   make up
   ```

3. Application will be available at:
   - HTTP: `http://localhost:5555`
   - gRPC: `localhost:50051`
   - gRPC-Web: `localhost:8090`

### Making gRPC Changes

1. Edit `.proto` files in `api/proto/`
2. Regenerate Go code: `make proto-go`
3. Regenerate JavaScript client (if needed): `make proto-js-all`
4. Rebuild and restart: `make down && make build && make up`

### Environment Variables

Key environment variables (see `.env.prod` for full list):
- `APP_ENV` - Environment mode (`local`, `prod`)
- `PORT` - HTTP server port (default: 5555)
- `GRPC_PORT` - gRPC server port (default: 50051)
- `GRPC_WEB_PORT` - gRPC-Web proxy port (default: 8090)
- `POSTGRES_DSN` - PostgreSQL connection string
- `JWT_SECRET` - Secret for JWT signing
- `TOKEN_TTL` - Access token TTL in seconds
- `REFRESH_TTL` - Refresh token TTL in seconds
- `KAFKA_ENABLE` - Enable/disable Kafka (true/false)
- `KAFKA_BROKER` - Kafka broker address
- `GIN_MODE` - Gin framework mode (`debug`, `release`)
- `CORS_ALLOWED_ORIGINS` - Comma-separated CORS origins

### Authentication Flow

The app uses JWT-based auth with access and refresh tokens:

1. **Middleware Chain** (in `internal/trasport/http/middleware/`):
   - `SessionMiddleware()` - Gin sessions
   - `CheckJWT()` - Validates JWT, auto-refreshes if needed
   - `LoadUserContext()` - Loads user from DB/cache into context
   - `NoCache` - Prevents browser caching
   - `TemplateContext` - Adds common template variables

2. **Protected Routes**: Wrapped in `middleware.CheckAuthAndRedirect()` which redirects to login if unauthenticated

3. **Token Storage**: JWTs stored in session cookies (`vago_token`, `vago_refresh_token`)

### Graceful Shutdown

The application (`cmd/vago/main.go`) handles graceful shutdown:
- Listens for `SIGTERM`/`SIGINT`
- Cancels context to notify all goroutines
- In DEV mode (`APP_ENV=local`): instant shutdown
- In PROD mode: 
  - Closes Kafka consumer first
  - GracefulStop for gRPC (10s timeout)
  - Waits for all goroutines via `sync.WaitGroup`

### Seeding Data

The application includes seed functions accessible via protected routes:
- `POST /runTopicsSeed` - Seeds learning topics
- `POST /run_questions_seed` - Seeds quiz questions

Admin user is created via `db/03_admin.sql` on first database init.

## Code Patterns and Conventions

### Repository Pattern

All data access goes through repository interfaces defined in `domain` and implemented in `infra/persistence/gorm/`:
- Entities are GORM models with `gorm` tags (e.g., `userEntity.go`)
- Repositories expose domain-model methods (e.g., `userRepo.go`)
- Services in `application/` depend on repository interfaces, not implementations

### Service Initialization

Services are instantiated in `internal/trasport/http/router.go`:
```go
taskSvc := task.NewService(gorm.NewTaskRepo(ctx.DB))
userSvc := user.NewService(userRepo, tokenProvider)
```

Handlers receive services via constructor injection:
```go
authH := handler.NewAuthHandler(userSvc, ctx.Cfg.JwtSecret, ...)
```

### Logging

Use structured logging via Zap (available as `appCtx.Log`):
```go
appCtx.Log.Infow("message", "key1", value1, "key2", value2)
appCtx.Log.Errorw("error occurred", "error", err)
```

### Configuration Constants

Use constants from `internal/config/code/` instead of string literals:
```go
c.SetCookie(code.VagoToken, token, ...)
currentUser := c.MustGet(code.CurrentUser)
```

## Important Notes

- **No test files exist** - there is no existing test suite or testing framework configured
- **Typo in directory name**: `trasport` should be `transport` (but don't fix without coordination)
- **Russian language**: User-facing content, comments, and README are in Russian
- **Database migrations**: Manual SQL files, not a migration tool
- **Environment switching**: Controlled by `APP_ENV` env var, loads `.env.local` only in local mode

## Шпаргалка
### Postgres
#### Локально
```shell
psql -h localhost -p 5432 -U vadmark -d vagodb
\i db/01_init.sql
\i db/02_quiz.sql
```
#### VDS
```postgresql
UPDATE users
SET
    role = 'admin',
    username = 'Vadmark'
WHERE id = 1;
```


### Перенос
```shell
scp -P 2499 -i ~/.ssh/id_vado .env.prod vadmark@159.255.33.142:~/vago/.env.prod
```
 
## Доработки
- При регистрации просить повторить пароль
- При логине и регистрации фокус сразу кидать на поле ввода логина
- Отловить 503 в quiz
- В книге продумать и реализовать адекватную навигацию
- Снять ограничения в длине сообщений в чате
- Длинное описание задач ломают верстку
- Одно название викторина или quiz
- Синхронизировать порядок в меню и на главной странице