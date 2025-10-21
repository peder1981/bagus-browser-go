#!/bin/bash
# Bagus Browser Go - Script de Testes (Linux/macOS)
# Uso: ./scripts/test.sh [unit|coverage|all]

set -e

# Cores
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Banner
echo -e "${BLUE}╔════════════════════════════════════════╗${NC}"
echo -e "${BLUE}║   Bagus Browser Go - Test Script     ║${NC}"
echo -e "${BLUE}╚════════════════════════════════════════╝${NC}"
echo ""

# Função para testes unitários
run_unit_tests() {
    echo -e "${YELLOW}► Executando testes unitários...${NC}"
    go test -v ./...
    
    if [ $? -eq 0 ]; then
        echo -e "${GREEN}✓ Todos os testes passaram${NC}"
    else
        echo -e "${RED}✗ Alguns testes falharam${NC}"
        exit 1
    fi
}

# Função para coverage
run_coverage() {
    echo -e "${YELLOW}► Executando testes com coverage...${NC}"
    go test -v -coverprofile=coverage.out ./...
    
    if [ $? -eq 0 ]; then
        echo -e "${GREEN}✓ Testes concluídos${NC}"
        echo -e "${YELLOW}► Gerando relatório HTML...${NC}"
        go tool cover -html=coverage.out -o coverage.html
        echo -e "${GREEN}✓ Relatório gerado: coverage.html${NC}"
        
        # Mostrar resumo
        echo ""
        echo -e "${BLUE}═══ Coverage Summary ═══${NC}"
        go tool cover -func=coverage.out | tail -n 1
    else
        echo -e "${RED}✗ Alguns testes falharam${NC}"
        exit 1
    fi
}

# Função para lint
run_lint() {
    echo -e "${YELLOW}► Executando linter...${NC}"
    
    if ! command -v golangci-lint &> /dev/null; then
        echo -e "${YELLOW}golangci-lint não encontrado. Instalando...${NC}"
        go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
    fi
    
    golangci-lint run ./...
    
    if [ $? -eq 0 ]; then
        echo -e "${GREEN}✓ Lint passou${NC}"
    else
        echo -e "${RED}✗ Problemas encontrados no lint${NC}"
        exit 1
    fi
}

# Função para formatar
run_fmt() {
    echo -e "${YELLOW}► Formatando código...${NC}"
    go fmt ./...
    echo -e "${GREEN}✓ Código formatado${NC}"
}

# Determinar tipo de teste
TEST_TYPE=${1:-all}

case $TEST_TYPE in
    unit)
        run_unit_tests
        ;;
    coverage)
        run_coverage
        ;;
    lint)
        run_lint
        ;;
    fmt)
        run_fmt
        ;;
    all)
        echo -e "${BLUE}═══ Executando Todos os Testes ═══${NC}"
        run_fmt
        run_lint
        run_unit_tests
        run_coverage
        ;;
    *)
        echo -e "${RED}Tipo de teste inválido: $TEST_TYPE${NC}"
        echo "Uso: $0 [unit|coverage|lint|fmt|all]"
        exit 1
        ;;
esac

echo ""
echo -e "${GREEN}╔════════════════════════════════════════╗${NC}"
echo -e "${GREEN}║      Testes Concluídos com Sucesso!   ║${NC}"
echo -e "${GREEN}╚════════════════════════════════════════╝${NC}"
