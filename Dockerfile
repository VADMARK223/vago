# Этап 1: Сборка (builder) слой с зависимостями
FROM golang:1.25 AS builder

# Рабочая директория
WORKDIR /app

# Копируем go.mod и go.sum - чтобы кэшировался слой зависимостей
COPY go.mod go.sum ./
# Качаем зависимости
RUN go mod download

# Копируем только исходники
COPY api ./api
COPY "cmd" "./cmd"
COPY internal ./internal
COPY pkg ./pkg
COPY web ./web
COPY migrations ./migrations

# Ставим goose как бинарник
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# Сборка бинарника (статическая, без CGO)
# CGO_ENABLED=0 компилято выключает исользование С, и Go собирает чистый статический бинарник. Если чистое CLI, для GUI может все сломать
# -trimpath убирает пути из бинаря (безопасность + меньше размер)
# -o указывает название выходного бинарника
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -trimpath -o vago ./cmd/vago

# Этап 2: рантайм (минимальный финальный образ)
FROM debian:bookworm-slim AS runtime

WORKDIR /app

# Копируем бинарник из builder-этапа
COPY --from=builder /app/vago .
# Копируем миграции
COPY --from=builder /app/migrations ./migrations
# Копируем шаблоны и статику
COPY --from=builder /app/web ./web

# Копируем goose бинарник
COPY --from=builder /go/bin/goose /usr/local/bin/goose

# Порт для gRPC и HTTP
EXPOSE 50051 5555 8090

# Устанавливаем базовые env (не мешают docker-compose)
ENV APP_ENV=production \
    GIN_MODE=release

# Запускаем сервер
CMD ["./vago"]