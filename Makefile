# Bagus Browser Go - Makefile
# Builds multiplataforma e automação de tarefas
# Suporta duas versões: Webview (leve) e CEF (100% compatível)

# Variáveis
APP_NAME := bagus
VERSION := 2.0.0-alpha
BUILD_DIR := build
DIST_DIR := dist
CMD_DIR := .

# Detectar sistema operacional
UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Linux)
	OS := linux
endif
ifeq ($(UNAME_S),Darwin)
	OS := darwin
endif
ifeq ($(OS),Windows_NT)
	OS := windows
endif

# Flags de compilação
LDFLAGS := -ldflags "-s -w -X main.appVersion=$(VERSION)"
GO_BUILD := go build $(LDFLAGS)

# Cores para output
RED := \033[0;31m
GREEN := \033[0;32m
YELLOW := \033[1;33m
NC := \033[0m # No Color

.PHONY: all build clean test lint fmt deps help menu build-webview build-cef install-cef verify-privacy run-cef

# Target padrão
all: clean deps fmt lint test build

# Help
help:
	@echo "$(GREEN)╔════════════════════════════════════════╗$(NC)"
	@echo "$(GREEN)║   Bagus Browser Go - Makefile        ║$(NC)"
	@echo "$(GREEN)╚════════════════════════════════════════╝$(NC)"
	@echo ""
	@echo "$(YELLOW)📦 Build:$(NC)"
	@echo "  make menu           - Menu interativo de build"
	@echo "  make build          - Compila versão Webview (padrão)"
	@echo "  make build-webview  - Compila versão Webview (leve)"
	@echo "  make build-cef      - Compila versão CEF (100% compatível)"
	@echo "  make install-cef    - Instala CEF (necessário para build-cef)"
	@echo ""
	@echo "$(YELLOW)🔒 Privacidade:$(NC)"
	@echo "  make verify-privacy - Verifica zero telemetria"
	@echo ""
	@echo "$(YELLOW)🚀 Execução:$(NC)"
	@echo "  make run            - Executa versão Webview"
	@echo "  make run-cef        - Executa versão CEF"
	@echo ""
	@echo "$(YELLOW)🧪 Testes:$(NC)"
	@echo "  make test           - Executa testes (pula UI/CGO)"
	@echo "  make test-coverage  - Executa testes com coverage (pula UI/CGO)"
	@echo "  make test-all       - Executa TODOS os testes (requer GTK/WebKit)"
	@echo ""
	@echo "$(YELLOW)🛠️  Utilitários:$(NC)"
	@echo "  make lint           - Executa linter"
	@echo "  make fmt            - Formata código"
	@echo "  make deps           - Instala dependências"
	@echo "  make clean          - Remove arquivos de build"
	@echo ""
	@echo "$(YELLOW)📦 Distribuição:$(NC)"
	@echo "  make build-all      - Compila para todas as plataformas"
	@echo "  make dist           - Cria pacotes de distribuição"
	@echo ""

# Menu interativo
menu:
	@./scripts/build_menu.sh

# Instalar dependências
deps:
	@echo "$(GREEN)Instalando dependências...$(NC)"
	go mod download
	go mod tidy
	@echo "$(GREEN)✓ Dependências instaladas$(NC)"

# Build para plataforma atual (Webview - padrão)
build: build-webview

# Build versão Webview (leve, 70% compatibilidade)
build-webview: deps
	@echo "$(GREEN)╔════════════════════════════════════════╗$(NC)"
	@echo "$(GREEN)║   Compilando Versão Webview           ║$(NC)"
	@echo "$(GREEN)╚════════════════════════════════════════╝$(NC)"
	@mkdir -p $(BUILD_DIR)
	$(GO_BUILD) -o $(BUILD_DIR)/$(APP_NAME) ./$(CMD_DIR)
	@echo "$(GREEN)✓ Build concluído: $(BUILD_DIR)/$(APP_NAME)$(NC)"

# Build versão CEF (100% compatibilidade)
build-cef:
	@echo "$(GREEN)╔════════════════════════════════════════╗$(NC)"
	@echo "$(GREEN)║   Compilando Versão CEF               ║$(NC)"
	@echo "$(GREEN)╚════════════════════════════════════════╝$(NC)"
	@./scripts/build/build_cef.sh

# Instalar CEF
install-cef:
	@echo "$(GREEN)╔════════════════════════════════════════╗$(NC)"
	@echo "$(GREEN)║   Instalando CEF                      ║$(NC)"
	@echo "$(GREEN)╚════════════════════════════════════════╝$(NC)"
	@./scripts/setup/install_cef.sh

# Verificar privacidade (zero telemetria)
verify-privacy:
	@./scripts/testing/verify_privacy.sh

# Build para Linux
build-linux:
	@echo "$(GREEN)Compilando para Linux...$(NC)"
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 $(GO_BUILD) -o $(BUILD_DIR)/$(APP_NAME)-linux-amd64 ./$(CMD_DIR)
	GOOS=linux GOARCH=arm64 $(GO_BUILD) -o $(BUILD_DIR)/$(APP_NAME)-linux-arm64 ./$(CMD_DIR)
	@echo "$(GREEN)✓ Build Linux concluído$(NC)"

# Build para Windows
build-windows:
	@echo "$(GREEN)Compilando para Windows...$(NC)"
	@mkdir -p $(BUILD_DIR)
	GOOS=windows GOARCH=amd64 $(GO_BUILD) -o $(BUILD_DIR)/$(APP_NAME)-windows-amd64.exe ./$(CMD_DIR)
	@echo "$(GREEN)✓ Build Windows concluído$(NC)"

# Build para macOS
build-macos:
	@echo "$(GREEN)Compilando para macOS...$(NC)"
	@mkdir -p $(BUILD_DIR)
	GOOS=darwin GOARCH=amd64 $(GO_BUILD) -o $(BUILD_DIR)/$(APP_NAME)-darwin-amd64 ./$(CMD_DIR)
	GOOS=darwin GOARCH=arm64 $(GO_BUILD) -o $(BUILD_DIR)/$(APP_NAME)-darwin-arm64 ./$(CMD_DIR)
	@echo "$(GREEN)✓ Build macOS concluído$(NC)"

# Build para todas as plataformas
build-all: build-linux build-windows build-macos
	@echo "$(GREEN)✓ Build completo para todas as plataformas$(NC)"

# Executar testes (pula pacotes que requerem CGO)
test:
	@echo "$(GREEN)Executando testes...$(NC)"
	@echo "$(YELLOW)Nota: Pulando pacotes que requerem GTK/WebKit (main, cmd/bagus, internal/ui)$(NC)"
	go test -v ./internal/browser/... ./internal/config/... ./internal/lockfile/... ./internal/security/... ./internal/storage/... ./pkg/...
	@echo "$(GREEN)✓ Testes concluídos$(NC)"

# Executar testes com coverage (pula pacotes que requerem CGO)
test-coverage:
	@echo "$(GREEN)Executando testes com coverage...$(NC)"
	@echo "$(YELLOW)Nota: Pulando pacotes que requerem GTK/WebKit (main, cmd/bagus, internal/ui)$(NC)"
	go test -v -coverprofile=coverage.out ./internal/browser/... ./internal/config/... ./internal/lockfile/... ./internal/security/... ./internal/storage/... ./pkg/...
	@if [ -f coverage.out ]; then \
		go tool cover -html=coverage.out -o coverage.html; \
		echo "$(GREEN)✓ Coverage gerado: coverage.html$(NC)"; \
		go tool cover -func=coverage.out | tail -1; \
	fi

# Executar TODOS os testes (requer GTK/WebKit instalados)
test-all:
	@echo "$(GREEN)Executando TODOS os testes (incluindo UI)...$(NC)"
	@echo "$(YELLOW)Requer: libwebkit2gtk-4.1-dev libgtk-3-dev$(NC)"
	go test -v ./...
	@echo "$(GREEN)✓ Todos os testes concluídos$(NC)"

# Lint
lint:
	@echo "$(GREEN)Executando linter...$(NC)"
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run ./...; \
	else \
		echo "$(YELLOW)golangci-lint não encontrado. Instalando...$(NC)"; \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
		golangci-lint run ./...; \
	fi
	@echo "$(GREEN)✓ Lint concluído$(NC)"

# Formatar código
fmt:
	@echo "$(GREEN)Formatando código...$(NC)"
	go fmt ./...
	@echo "$(GREEN)✓ Código formatado$(NC)"

# Executar aplicação (Webview)
run: build
	@echo "$(GREEN)Executando $(APP_NAME) (Webview)...$(NC)"
	./$(BUILD_DIR)/$(APP_NAME)

# Executar versão CEF
run-cef:
	@echo "$(GREEN)Executando $(APP_NAME) (CEF)...$(NC)"
	cd $(BUILD_DIR) && ./run_bagus_cef.sh

# Executar com debug
run-debug: build
	@echo "$(GREEN)Executando $(APP_NAME) em modo debug...$(NC)"
	./$(BUILD_DIR)/$(APP_NAME) --debug

# Limpar arquivos de build
clean:
	@echo "$(YELLOW)Limpando arquivos de build...$(NC)"
	rm -rf $(BUILD_DIR)
	rm -rf $(DIST_DIR)
	rm -f coverage.out coverage.html
	@echo "$(GREEN)✓ Limpeza concluída$(NC)"

# Criar pacotes de distribuição
dist: build-all
	@echo "$(GREEN)Criando pacotes de distribuição...$(NC)"
	@mkdir -p $(DIST_DIR)
	cd $(BUILD_DIR) && tar -czf ../$(DIST_DIR)/$(APP_NAME)-$(VERSION)-linux-amd64.tar.gz $(APP_NAME)-linux-amd64
	cd $(BUILD_DIR) && tar -czf ../$(DIST_DIR)/$(APP_NAME)-$(VERSION)-linux-arm64.tar.gz $(APP_NAME)-linux-arm64
	cd $(BUILD_DIR) && tar -czf ../$(DIST_DIR)/$(APP_NAME)-$(VERSION)-darwin-amd64.tar.gz $(APP_NAME)-darwin-amd64
	cd $(BUILD_DIR) && tar -czf ../$(DIST_DIR)/$(APP_NAME)-$(VERSION)-darwin-arm64.tar.gz $(APP_NAME)-darwin-arm64
	cd $(BUILD_DIR) && zip ../$(DIST_DIR)/$(APP_NAME)-$(VERSION)-windows-amd64.zip $(APP_NAME)-windows-amd64.exe
	@echo "$(GREEN)✓ Pacotes criados em $(DIST_DIR)/$(NC)"

# Instalar localmente
install: build
	@echo "$(GREEN)Instalando $(APP_NAME)...$(NC)"
	cp $(BUILD_DIR)/$(APP_NAME) /usr/local/bin/
	@echo "$(GREEN)✓ Instalado em /usr/local/bin/$(APP_NAME)$(NC)"

# Desinstalar
uninstall:
	@echo "$(YELLOW)Desinstalando $(APP_NAME)...$(NC)"
	rm -f /usr/local/bin/$(APP_NAME)
	@echo "$(GREEN)✓ Desinstalado$(NC)"
