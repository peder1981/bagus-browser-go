#!/bin/bash
# Bagus Browser Go - Script de Publicação no GitHub (Linux/macOS)
# Uso: ./scripts/publish.sh [mensagem do commit]

set -e

# Cores
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color

# Banner
echo -e "${CYAN}╔════════════════════════════════════════╗${NC}"
echo -e "${CYAN}║  Bagus Browser Go - Publish Script   ║${NC}"
echo -e "${CYAN}╚════════════════════════════════════════╝${NC}"
echo ""

# Verificar se está em um repositório git
if [ ! -d ".git" ]; then
    echo -e "${RED}✗ Erro: Não é um repositório Git${NC}"
    echo "Execute 'git init' primeiro"
    exit 1
fi

# Verificar se há mudanças
if git diff-index --quiet HEAD -- 2>/dev/null; then
    echo -e "${YELLOW}⚠ Nenhuma mudança para commitar${NC}"
    read -p "Deseja continuar mesmo assim? (s/N): " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Ss]$ ]]; then
        echo -e "${BLUE}Publicação cancelada${NC}"
        exit 0
    fi
fi

# Mensagem de commit
COMMIT_MSG="${1:-Update: $(date +'%Y-%m-%d %H:%M')}"

echo -e "${BLUE}═══ Preparando Publicação ═══${NC}"
echo ""

# 1. Verificar qualidade do código
echo -e "${YELLOW}► Verificando qualidade do código...${NC}"

# Formatar
echo "  • Formatando código..."
go fmt ./... > /dev/null 2>&1
echo -e "    ${GREEN}✓${NC} Código formatado"

# Lint (opcional, não bloqueia)
if command -v golangci-lint &> /dev/null; then
    echo "  • Executando lint..."
    if golangci-lint run ./... > /dev/null 2>&1; then
        echo -e "    ${GREEN}✓${NC} Lint passou"
    else
        echo -e "    ${YELLOW}⚠${NC} Lint encontrou problemas (não bloqueante)"
    fi
fi

# Testes
echo "  • Executando testes..."
if go test ./... > /dev/null 2>&1; then
    echo -e "    ${GREEN}✓${NC} Todos os testes passaram"
else
    echo -e "    ${RED}✗${NC} Alguns testes falharam"
    read -p "Deseja continuar mesmo assim? (s/N): " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Ss]$ ]]; then
        echo -e "${RED}Publicação cancelada${NC}"
        exit 1
    fi
fi

# Build
echo "  • Compilando..."
if go build -o build/bagus ./cmd/bagus > /dev/null 2>&1; then
    echo -e "    ${GREEN}✓${NC} Build bem-sucedido"
else
    echo -e "    ${RED}✗${NC} Build falhou"
    exit 1
fi

echo ""

# 2. Git status
echo -e "${BLUE}═══ Status do Git ═══${NC}"
git status --short
echo ""

# 3. Adicionar arquivos
echo -e "${YELLOW}► Adicionando arquivos...${NC}"
git add .
echo -e "${GREEN}✓ Arquivos adicionados${NC}"
echo ""

# 4. Commit
echo -e "${YELLOW}► Criando commit...${NC}"
echo -e "  Mensagem: ${CYAN}\"$COMMIT_MSG\"${NC}"
git commit -m "$COMMIT_MSG"
echo -e "${GREEN}✓ Commit criado${NC}"
echo ""

# 5. Verificar remote
echo -e "${YELLOW}► Verificando remote...${NC}"
if git remote | grep -q "origin"; then
    REMOTE_URL=$(git remote get-url origin)
    echo -e "  Remote: ${CYAN}$REMOTE_URL${NC}"
else
    echo -e "${YELLOW}⚠ Remote 'origin' não configurado${NC}"
    echo ""
    echo "Configure o remote com:"
    echo -e "${CYAN}git remote add origin git@github.com:peder1981/bagus-browser-go.git${NC}"
    echo "ou"
    echo -e "${CYAN}git remote add origin https://github.com/peder1981/bagus-browser-go.git${NC}"
    exit 1
fi
echo ""

# 6. Push
echo -e "${YELLOW}► Publicando no GitHub...${NC}"
BRANCH=$(git branch --show-current)
echo -e "  Branch: ${CYAN}$BRANCH${NC}"

if git push origin "$BRANCH"; then
    echo -e "${GREEN}✓ Push bem-sucedido${NC}"
else
    echo -e "${RED}✗ Push falhou${NC}"
    echo ""
    echo "Se for o primeiro push, tente:"
    echo -e "${CYAN}git push -u origin $BRANCH${NC}"
    exit 1
fi

echo ""
echo -e "${GREEN}╔════════════════════════════════════════╗${NC}"
echo -e "${GREEN}║     Publicação Concluída! 🚀          ║${NC}"
echo -e "${GREEN}╚════════════════════════════════════════╝${NC}"
echo ""
echo -e "Branch: ${CYAN}$BRANCH${NC}"
echo -e "Commit: ${CYAN}$COMMIT_MSG${NC}"
echo ""
