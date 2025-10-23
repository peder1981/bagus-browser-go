#
# Bagus Browser - Makefile Simplificado
# Todos os comandos agora usam o script master ./bagus
#

.PHONY: help build install clean test version release publish status

help:
	@./bagus help

build:
	@./bagus build

install:
	@./bagus install

clean:
	@./bagus clean

test:
	@./bagus test

version:
	@./bagus version

status:
	@./bagus status

release:
ifdef VERSION
	@./bagus release $(VERSION)
else
	@echo "❌ Erro: Especifique a versão"
	@echo "Uso: make release VERSION=4.5.1"
	@echo "Ou use: ./bagus release 4.5.1"
	@exit 1
endif

publish:
	@./bagus publish-auto

run:
	@./bagus run

# Atalhos para compatibilidade
all: build
.DEFAULT_GOAL := help
