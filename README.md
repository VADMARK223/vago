# Vago

**Vago** - учебный портал для изучения Golang, включающий:
- теоретические материалы
- систему квизов
- realtime-чат
- персональный трекер задач

Проект реализован как **production-style backend-приложение** с REST и gRPC (streaming), PostgreSQL и деплоем в Docker на VDS.

🌐 Demo: http://vadmark.duckdns.org/

---

## Ключевые возможности

- ✅ REST API (Gin)
- ✅ gRPC + gRPC-Web (realtime chat streaming)
- ✅ JWT аутентификация (access + refresh, auto-refresh)
- ✅ PostgreSQL + SQL-миграции
- ✅ Docker / Docker Compose
- ✅ Graceful shutdown (context, WaitGroup)
- ✅ Server-side rendered web UI
- ✅ Kafka (опционально, через feature flag)

---

## Технологии

- **Go**
- **Gin** (HTTP)
- **gRPC / gRPC-Web**
- **PostgreSQL**
- **GORM**
- **JWT**
- **Docker / Docker Compose**
- **Zap logger**
- **Kafka (optional)**

---

## Архитектура

Проект построен по слоистой архитектуре с разделением ответственности:

- `cmd/` - точки входа приложения
- `internal/domain` - доменные модели и интерфейсы
- `internal/application` - бизнес-логика (services)
- `internal/infra` - инфраструктура (DB, JWT, Kafka, logger)
- `internal/transport`
    - `http` - REST API (Gin)
    - `grpc` - gRPC сервер
    - `ws` - WebSocket hub
- `db/` - SQL-миграции
- `web/` - HTML templates и static assets

---

## Быстрый старт (Docker)

```bash
git clone https://github.com/VADMARK223/vago.git
cd vago

cp .env.prod .env.local
make up
