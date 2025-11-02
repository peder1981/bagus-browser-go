#!/bin/bash
#
# ⚡ Quick Bundle Test - Teste rápido do workflow completo
# Versão: 1.0
#

set -e

# Cores
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
CYAN='\033[0;36m'
NC='\033[0m'

echo -e "${CYAN}"
echo "╔════════════════════════════════════════╗"
echo "║   ⚡ Quick Bundle Test                ║"
echo "║        Versão: 1.0                    ║"
echo "╚════════════════════════════════════════╝"
echo -e "${NC}"

# Configurações
BUILD_DIR="build"
DIST_DIR="dist"
VERSION="v5.0.0"

# Passo 1: Build
echo -e "${BLUE}📦 Passo 1: Compilando...${NC}"
if ./scripts/bagus build > /tmp/build.log 2>&1; then
    echo -e "${GREEN}✅ Build OK${NC}"
else
    echo -e "${RED}❌ Build falhou${NC}"
    tail -20 /tmp/build.log
    exit 1
fi

# Passo 2: Bundle
echo -e "${BLUE}📦 Passo 2: Embarcando dependências...${NC}"
if ./scripts/bagus bundle > /tmp/bundle.log 2>&1; then
    echo -e "${GREEN}✅ Bundle OK${NC}"
else
    echo -e "${RED}❌ Bundle falhou${NC}"
    tail -20 /tmp/bundle.log
    exit 1
fi

# Passo 3: Teste do Bundle
echo -e "${BLUE}📦 Passo 3: Testando bundle...${NC}"
if ./scripts/test-bundled.sh > /tmp/test.log 2>&1; then
    echo -e "${GREEN}✅ Testes OK${NC}"
else
    echo -e "${RED}❌ Testes falharam${NC}"
    tail -20 /tmp/test.log
    exit 1
fi

# Passo 4: Criar .deb bundled
echo -e "${BLUE}📦 Passo 4: Criando .deb bundled...${NC}"
if ./scripts/build-deb-bundled.sh ${VERSION} > /tmp/deb.log 2>&1; then
    echo -e "${GREEN}✅ .deb OK${NC}"
else
    echo -e "${RED}❌ .deb falhou${NC}"
    tail -20 /tmp/deb.log
    exit 1
fi

# Resumo
echo ""
echo -e "${GREEN}╔════════════════════════════════════════╗${NC}"
echo -e "${GREEN}║   ✅ Teste Completo com Sucesso!     ║${NC}"
echo -e "${GREEN}╚════════════════════════════════════════╝${NC}"
echo ""

# Estatísticas
echo -e "${CYAN}📊 Estatísticas:${NC}"
echo ""

# Binário
if [ -f "${BUILD_DIR}/bagus-browser" ]; then
    BINARY_SIZE=$(du -h ${BUILD_DIR}/bagus-browser | cut -f1)
    echo "  Binário: ${BINARY_SIZE}"
fi

# Bundle
if [ -d "${BUILD_DIR}/bundle" ]; then
    BUNDLE_SIZE=$(du -sh ${BUILD_DIR}/bundle | cut -f1)
    LIB_COUNT=$(find ${BUILD_DIR}/bundle/lib -type f | wc -l)
    echo "  Bundle: ${BUNDLE_SIZE} (${LIB_COUNT} arquivos)"
fi

# .deb files
echo ""
echo -e "${CYAN}📦 Pacotes criados:${NC}"
ls -lh ${DIST_DIR}/*.deb 2>/dev/null | awk '{print "  " $9 " (" $5 ")"}'

# Checksums
echo ""
echo -e "${CYAN}🔐 Checksums:${NC}"
if [ -f "${DIST_DIR}/SHA256SUMS" ]; then
    tail -3 ${DIST_DIR}/SHA256SUMS | awk '{print "  " $2 ": " substr($1, 1, 16) "..."}'
fi

echo ""
echo -e "${CYAN}🚀 Próximos passos:${NC}"
echo "  1. Testar instalação:"
echo "     sudo dpkg -i ${DIST_DIR}/bagus-browser_${VERSION}_amd64_bundled.deb"
echo ""
echo "  2. Executar:"
echo "     bagus-browser"
echo ""
echo "  3. Publicar (opcional):"
echo "     ./scripts/bagus publish"
echo ""
