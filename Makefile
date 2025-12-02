# =========================
# 🎨 HELP SECTION
# =========================
MAKEFLAGS += --no-print-directory
YELLOW:= \033[1;33m
GREEN := \033[1;32m
BLUE  := \033[1;34m
CYAN  := \033[1;36m
ORANGE := \033[38;5;208m
RESET := \033[0m

# =========================
# Read .env.prod
# =========================
ifneq (,$(wildcard .env.prod))
    include .env.prod
    export $(shell sed -n 's/^\([^#[:space:]]\+\)=.*/\1/p' .env.prod)
endif
ifeq ($(KAFKA_ENABLE), true)
	KAFKA_YML = -f docker-compose.kafka.yml
else
    KAFKA_YML =
endif

PROJECT_NAME = vago
COMPOSE = docker compose -p $(PROJECT_NAME)
COMPOSE_FULL = $(COMPOSE) -f docker-compose.yml $(KAFKA_YML)

PROTO_DIR = api/proto
PROTO_FILES := $(wildcard $(PROTO_DIR)/*.proto)
PROTOC = protoc

build:
	docker build -t ghcr.io/vadmark223/vago:latest .

push:
	docker push ghcr.io/vadmark223/vago:latest

pull:
	docker pull ghcr.io/vadmark223/vago:latest

up:
	docker compose -p $(PROJECT_NAME) -f docker-compose.yml $(KAFKA_YML) up -d

down:
	docker compose -p $(PROJECT_NAME) down

down-v:
	docker compose -p $(PROJECT_NAME) down -v

ps:
	$(COMPOSE) ps --format 'table {{.Name}}\t{{.Ports}}'

logs:
	docker compose -p $(PROJECT_NAME) logs --tail=20 vago

logs-f:
	docker compose -p $(PROJECT_NAME) logs -f --tail=20 vago

psql:
	docker exec -it vado-postgres psql -U vadmark -d vadodb

clean-all:
	docker system prune -af --volumes

proto-go:
	@echo "Generating Go gRPC files..."
	@for file in $(PROTO_FILES); do \
		echo "  -> Compilation $$file"; \
		$(PROTOC) -I=$(PROTO_DIR) $$file \
			--go_out=. \
			--go-grpc_out=. ; \
	done
	@echo "✅ Generation complete."

PB_WEB_OUT_DIR = ./web/static/js/pb
GRPC_WEB_PLUGIN = /usr/local/bin/protoc-gen-grpc-web

proto-js-clean:
	@echo "$(ORANGE)⚠️ Clear all *.js$(PB_WEB_OUT_DIR)...$(RESET)"
	@find $(PB_WEB_OUT_DIR) -type f \( -name "*.ts" -o -name "*.js" \) -delete
	@echo "$(GREEN)✅️ Cleaning is complete$(RESET)"

proto-js:
	@echo "🔧 Generating gRPC-Web JS files..."
	@mkdir -p $(PB_WEB_OUT_DIR)
	@for file in $(PROTO_FILES); do \
        echo "  🔵 Compilation $$file"; \
        $(PROTOC) -I=$(PROTO_DIR) $$file \
            --js_out=import_style=commonjs,binary:$(PB_WEB_OUT_DIR) \
            --plugin=protoc-gen-grpc-web=$(GRPC_WEB_PLUGIN) \
            --grpc-web_out=import_style=commonjs,mode=grpcwebtext:$(PB_WEB_OUT_DIR); \
    done
	@echo "$(GREEN)✅ Generation complete. Files in $(PB_WEB_OUT_DIR)$(RESET)"

bundle:
	@echo "$(BLUE)📦 Bundling JavaScript client...$(RESET)"
	npx esbuild web/static/js/index.js \
			--bundle \
			--format=esm \
			--outfile=web/static/js/bundle.js \
			--platform=browser \
			--target=es2020 \
			--define:process.env.GRPC_WEB_PORT="'$(GRPC_WEB_PORT)'"
	@echo "$(GREEN)✅ Bundle created → web/static/js/bundle.js$(RESET)"

proto-js-all: ## 🚀 Full pipeline: clean → generate → bundle
	@echo "$(BLUE)🚀 Starting full gRPC-Web JavaScript build pipeline...$(RESET)"
	@$(MAKE) proto-js-clean || { echo "$(ORANGE)❌ Stage failed: proto-ts-clean$(RESET)"; exit 1; }
	@$(MAKE) proto-js || { echo "$(ORANGE)❌ Stage failed: proto-ts$(RESET)"; exit 1; }
	@$(MAKE) bundle || { echo "$(ORANGE)❌ Stage failed: bundle$(RESET)"; exit 1; }
	@echo "$(GREEN)✅ All stages completed successfully!$(RESET)"

gen-questions:
	@echo "Run convert Json in SQL"
	go run ./cmd/genQuestions
	@echo "==> Выполнение SQL..."
	psql "postgresql://localhost:5432/vadodb" -f db/04_questions.sql

kafka-up:
	$(COMPOSE) $(KAFKA_YML) up -d

kafka-down:
	$(COMPOSE) $(KAFKA_YML) down

help:
	@echo "$(YELLOW)🧩 Available Make targets:$(RESET)"
	@echo ""
	@echo "  $(GREEN)make build$(RESET)     - 🔧 build image ghcr.io/vadmark223/vago:latest from Dockerfile"
	@echo "  $(GREEN)make push$(RESET)      - 📤 push image in GHCR"
	@echo "  $(GREEN)make pull$(RESET)      - 📥 pull image from GHCR"
	@echo "  $(GREEN)make up$(RESET)        - 🚀 start all containers"
	@echo "  $(GREEN)make down$(RESET)      - 🧯 stop all containers"
	@echo "  $(GREEN)make down-v$(RESET)    - 🧯 stop all containers (remove volumes)"
	@echo "  $(GREEN)make ps$(RESET)        - show containers"
	@echo "  $(GREEN)make logs$(RESET)      - 🧾 show logs"
	@echo "  $(GREEN)make logs-f$(RESET)    - 🧾 show logs (Follow)"
	@echo "  $(GREEN)make psql$(RESET)      - 🐘 open psql shell"
	@echo "  $(GREEN)make clean-all$(RESET) - ⚠️ clean all Docker (containers, images, volumes, networks)"
	@echo "  $(GREEN)make proto-go$(RESET)  - 🧠generating gRPC Go files"
	@echo ""
	@echo "$(CYAN)JavaScript proto:$(RESET)"
	@echo "  $(GREEN)make proto-js-clean$(RESET) - 🧹 Clean generated *.js, files from $(PB_WEB_OUT_DIR)"
	@echo "  $(GREEN)make proto-js$(RESET)       - 🔧 Generate gRPC-Web client files (.js,)"
	@echo "  $(GREEN)make bundle$(RESET)         - 📦 Bundle JavaScript client into a single bundle.js"
	@echo "  $(GREEN)make proto-js-all$(RESET)   - 🚀 Run the full pipeline: clean → generate → bundle"
	@echo ""
	@echo "$(CYAN)Others:$(RESET)"
	@echo "  $(GREEN)make kafka-up$(RESET)   - start kafka and kafka UI containers"
	@echo "  $(GREEN)make kafka-down$(RESET) - stop kafka and kafka UI containers"
	@echo ""
	@echo "$(CYAN)Qiuz:$(RESET)"
	@echo "  $(GREEN)make gen-questions$(RESET)   - generate questions from JSON"
.DEFAULT_GOAL := help