.PHONY: help build release publish clean install test version

help:
	@echo "Bagus Browser - Makefile"
	@echo ""
	@echo "Comandos disponíveis:"
	@echo "  make build     - Compilar e empacotar"
	@echo "  make release   - Criar release completa (recomendado)"
	@echo "  make version   - Ver versão atual"
	@echo "  make clean     - Limpar builds"
	@echo "  make install   - Build + Instalar localmente"
	@echo "  make test      - Testar compilação"
	@echo ""
	@echo "Novo workflow (recomendado):"
	@echo "  make release VERSION=4.5.0  - Cria release completa"
	@echo ""

build:
	@chmod +x scripts/build.sh
	@./scripts/build.sh

release:
	@chmod +x scripts/version.sh
ifdef VERSION
	@./scripts/version.sh release $(VERSION)
else
	@echo "❌ Erro: Especifique a versão"
	@echo "Uso: make release VERSION=4.5.0"
	@exit 1
endif

version:
	@chmod +x scripts/version.sh
	@./scripts/version.sh current

publish:
	@chmod +x scripts/publish.sh
	@./scripts/publish.sh

clean:
	@echo "🗑️  Limpando..."
	@rm -rf build/ dist/
	@rm -f bagus-browser *.log
	@echo "✅ Limpo!"

install: build
	@echo "📦 Instalando..."
	@sudo dpkg -i dist/bagus-browser_*.deb || sudo apt-get install -f
	@echo "✅ Instalado!"
	@echo ""
	@echo "🧹 Limpando temporários..."
	@rm -rf build/
	@echo "✅ Limpeza concluída!"

test:
	@echo "🧪 Testando compilação..."
	@go build -o /tmp/bagus-browser-test .
	@rm /tmp/bagus-browser-test
	@echo "✅ Teste OK!"
