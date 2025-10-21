#!/bin/bash
# Script para publicar release no GitHub

set -e

VERSION="2.0.0"
TAG="v${VERSION}"

# Cores
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m'

echo -e "${BLUE}╔═══════════════════════════════════════╗${NC}"
echo -e "${BLUE}║   BAGUS BROWSER - PUBLISH RELEASE    ║${NC}"
echo -e "${BLUE}╚═══════════════════════════════════════╝${NC}"
echo ""

# Verifica se gh CLI está instalado
if ! command -v gh &> /dev/null; then
    echo -e "${RED}✗ GitHub CLI (gh) não está instalado${NC}"
    echo -e "${YELLOW}Instale com: sudo apt install gh${NC}"
    exit 1
fi

# Verifica se está autenticado
if ! gh auth status &> /dev/null; then
    echo -e "${YELLOW}► Autenticando no GitHub...${NC}"
    gh auth login
fi

# Verifica se os arquivos de distribuição existem
if [ ! -d "dist" ] || [ -z "$(ls -A dist)" ]; then
    echo -e "${RED}✗ Diretório dist/ vazio. Execute ./build-all.sh primeiro${NC}"
    exit 1
fi

echo -e "${YELLOW}► Criando tag ${TAG}...${NC}"
git tag -a "${TAG}" -m "Release ${VERSION}" || echo -e "${YELLOW}  Tag já existe${NC}"

echo -e "${YELLOW}► Fazendo push da tag...${NC}"
git push origin "${TAG}" || echo -e "${YELLOW}  Tag já existe no remote${NC}"

echo -e "${YELLOW}► Criando release no GitHub...${NC}"

# Cria release notes
RELEASE_NOTES="## Bagus Browser ${VERSION}

### 🚀 Novidades
- ✅ Navegação direta com webview nativo (sem iframe)
- ✅ Proteção contra múltiplas instâncias
- ✅ Suporte completo a todos os sites
- ✅ Interface moderna e responsiva
- ✅ Histórico de navegação com busca
- ✅ Bloqueador de anúncios integrado

### 📦 Downloads

**Linux:**
- \`bagus-${VERSION}-linux-amd64.tar.gz\` - Para sistemas Linux 64-bit

**Windows:**
- \`bagus-${VERSION}-windows-amd64.zip\` - Para Windows 64-bit (requer build nativo)

**Debian/Ubuntu:**
- \`bagus-browser_1.0.0_amd64.deb\` - Pacote .deb para instalação fácil

### 🔧 Instalação

**Linux (método rápido):**
\`\`\`bash
tar -xzf bagus-${VERSION}-linux-amd64.tar.gz
sudo mv bagus-linux-amd64 /usr/local/bin/bagus
sudo chmod +x /usr/local/bin/bagus
bagus
\`\`\`

**Debian/Ubuntu:**
\`\`\`bash
sudo dpkg -i bagus-browser_1.0.0_amd64.deb
bagus-browser
\`\`\`

### 🐛 Correções
- Corrigido problema de múltiplas instâncias
- Corrigido problema de sites não carregarem
- Melhorada estabilidade geral

### 📝 Notas
- Primeira versão estável
- Zero telemetria
- 100% open source
"

# Cria release
gh release create "${TAG}" \
    --title "Bagus Browser ${VERSION}" \
    --notes "${RELEASE_NOTES}" \
    dist/* \
    build/bagus-browser_1.0.0_amd64.deb

echo ""
echo -e "${GREEN}╔═══════════════════════════════════════╗${NC}"
echo -e "${GREEN}║   RELEASE PUBLICADA COM SUCESSO!      ║${NC}"
echo -e "${GREEN}╚═══════════════════════════════════════╝${NC}"
echo ""
echo -e "${BLUE}URL da release:${NC}"
gh release view "${TAG}" --web
