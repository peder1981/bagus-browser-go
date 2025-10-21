#!/bin/bash
# Bagus Browser Go - Script de Setup Inicial do GitHub
# Uso: ./scripts/setup-github.sh

set -e

# Cores
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
CYAN='\033[0;36m'
NC='\033[0m'

echo -e "${CYAN}╔════════════════════════════════════════╗${NC}"
echo -e "${CYAN}║   Bagus Browser Go - GitHub Setup    ║${NC}"
echo -e "${CYAN}╚════════════════════════════════════════╝${NC}"
echo ""

# Verificar se já é um repositório git
if [ -d ".git" ]; then
    echo -e "${YELLOW}⚠ Já é um repositório Git${NC}"
    read -p "Deseja reconfigurar? (s/N): " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Ss]$ ]]; then
        echo -e "${BLUE}Setup cancelado${NC}"
        exit 0
    fi
else
    echo -e "${YELLOW}► Inicializando repositório Git...${NC}"
    git init
    echo -e "${GREEN}✓ Repositório inicializado${NC}"
    echo ""
fi

# Configurar remote
echo -e "${BLUE}═══ Configuração do Remote ═══${NC}"
echo ""
echo "Escolha o método de autenticação:"
echo "  1) SSH (recomendado)"
echo "  2) HTTPS"
echo ""
read -p "Opção (1/2): " -n 1 -r
echo ""
echo ""

if [[ $REPLY == "1" ]]; then
    REMOTE_URL="git@github.com:peder1981/bagus-browser-go.git"
    echo -e "${CYAN}URL SSH: $REMOTE_URL${NC}"
else
    REMOTE_URL="https://github.com/peder1981/bagus-browser-go.git"
    echo -e "${CYAN}URL HTTPS: $REMOTE_URL${NC}"
fi

# Remover remote existente se houver
if git remote | grep -q "origin"; then
    echo -e "${YELLOW}► Removendo remote existente...${NC}"
    git remote remove origin
fi

echo -e "${YELLOW}► Adicionando remote...${NC}"
git remote add origin "$REMOTE_URL"
echo -e "${GREEN}✓ Remote configurado${NC}"
echo ""

# Configurar branch principal
echo -e "${YELLOW}► Configurando branch principal...${NC}"
git branch -M main
echo -e "${GREEN}✓ Branch 'main' configurada${NC}"
echo ""

# Primeiro commit
echo -e "${BLUE}═══ Primeiro Commit ═══${NC}"
echo ""

if git rev-parse HEAD >/dev/null 2>&1; then
    echo -e "${YELLOW}⚠ Já existem commits${NC}"
else
    echo -e "${YELLOW}► Adicionando arquivos...${NC}"
    git add .
    echo -e "${GREEN}✓ Arquivos adicionados${NC}"
    
    echo -e "${YELLOW}► Criando commit inicial...${NC}"
    git commit -m "Initial commit: Bagus Browser Go v2.0.0-alpha

- Estrutura completa do projeto
- Motor do navegador básico
- Sistema de armazenamento
- Testes unitários (29 testes)
- Documentação completa
- Scripts de build/teste multiplataforma
- Manuais de usuário e desenvolvedor"
    
    echo -e "${GREEN}✓ Commit inicial criado${NC}"
fi

echo ""

# Push inicial
echo -e "${BLUE}═══ Publicação Inicial ═══${NC}"
echo ""
echo -e "${YELLOW}► Fazendo push para o GitHub...${NC}"

if git push -u origin main; then
    echo -e "${GREEN}✓ Push bem-sucedido${NC}"
else
    echo -e "${RED}✗ Push falhou${NC}"
    echo ""
    echo -e "${YELLOW}Possíveis causas:${NC}"
    echo "  1. Repositório não existe no GitHub"
    echo "  2. Sem permissão de acesso"
    echo "  3. Chave SSH não configurada (se usando SSH)"
    echo ""
    echo -e "${CYAN}Passos para resolver:${NC}"
    echo ""
    echo "1. Crie o repositório no GitHub:"
    echo "   https://github.com/new"
    echo ""
    echo "2. Se usando SSH, configure sua chave:"
    echo "   ssh-keygen -t ed25519 -C \"seu-email@example.com\""
    echo "   cat ~/.ssh/id_ed25519.pub"
    echo "   # Adicione a chave em: https://github.com/settings/keys"
    echo ""
    echo "3. Tente novamente:"
    echo "   git push -u origin main"
    exit 1
fi

echo ""
echo -e "${GREEN}╔════════════════════════════════════════╗${NC}"
echo -e "${GREEN}║      Setup Concluído! 🎉              ║${NC}"
echo -e "${GREEN}╚════════════════════════════════════════╝${NC}"
echo ""
echo -e "Repositório: ${CYAN}https://github.com/peder1981/bagus-browser-go${NC}"
echo -e "Branch: ${CYAN}main${NC}"
echo ""
echo -e "${YELLOW}Próximos passos:${NC}"
echo "  • Configure proteção de branch no GitHub"
echo "  • Adicione colaboradores se necessário"
echo "  • Configure GitHub Actions (CI/CD)"
echo "  • Adicione topics/tags ao repositório"
echo ""
echo -e "${CYAN}Para publicar mudanças futuras:${NC}"
echo "  ./scripts/publish.sh \"mensagem do commit\""
echo ""
