#
# Bagus Browser - Makefile Simplificado
# Todos os comandos agora usam o script master ./bagus
#

.PHONY: help build install clean test version release publish status

help:
	@./scripts/bagus help

build:
	@./scripts/bagus build

install:
	@./scripts/bagus install

clean:
	@./scripts/bagus clean

test:
	@./scripts/bagus test

version:
	@./scripts/bagus version

status:
	@./scripts/bagus status

release:
ifdef VERSION
	@./scripts/bagus release $(VERSION)
else
	@echo "❌ Erro: Especifique a versão"
	@echo "Uso: make release VERSION=4.5.1"
	@echo "Ou use: ./scripts/bagus release 4.5.1"
	@exit 1
endif

publish:
	@./scripts/bagus publish-auto

run:
	@./scripts/bagus run

# Atalhos para compatibilidade
all: build
.DEFAULT_GOAL := help
