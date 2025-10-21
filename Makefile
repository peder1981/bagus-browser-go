.PHONY: help build release publish clean install test

help:
	@echo "Bagus Browser - Makefile"
	@echo ""
	@echo "Comandos disponíveis:"
	@echo "  make build     - Compilar e empacotar"
	@echo "  make release   - Criar release no GitHub"
	@echo "  make publish   - Build + Release completo"
	@echo "  make clean     - Limpar builds"
	@echo "  make install   - Instalar localmente"
	@echo "  make test      - Testar compilação"
	@echo ""

build:
	@chmod +x scripts/build.sh
	@./scripts/build.sh

release:
	@chmod +x scripts/release.sh
	@./scripts/release.sh

publish:
	@chmod +x scripts/publish.sh
	@./scripts/publish.sh

clean:
	@echo "🗑️  Limpando..."
	@rm -rf build/ dist/
	@rm -f bagus-browser
	@echo "✅ Limpo!"

install: build
	@echo "📦 Instalando..."
	@sudo cp build/bagus-browser /usr/local/bin/
	@sudo cp build/bagus-browser.desktop /usr/share/applications/
	@echo "✅ Instalado em /usr/local/bin/bagus-browser"

test:
	@echo "🧪 Testando compilação..."
	@go build -o /tmp/bagus-browser-test .
	@rm /tmp/bagus-browser-test
	@echo "✅ Teste OK!"
