#!/bin/bash
set -e

# Cores
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

echo -e "${BLUE}╔════════════════════════════════════════╗${NC}"
echo -e "${BLUE}║  Bagus Browser - Publish Complete     ║${NC}"
echo -e "${BLUE}╚════════════════════════════════════════╝${NC}"
echo ""

# 1. Build
echo -e "${BLUE}[1/3] 🔨 Building...${NC}"
bash scripts/build.sh
if [ $? -ne 0 ]; then
    echo -e "${RED}❌ Build failed${NC}"
    exit 1
fi
echo ""

# 2. Commit e Push
echo -e "${BLUE}[2/3] 📤 Committing and pushing...${NC}"
read -p "Commit message: " COMMIT_MSG
if [ -n "$COMMIT_MSG" ]; then
    git add -A
    git commit -m "$COMMIT_MSG" || echo "Nothing to commit"
    git push origin main
    echo -e "${GREEN}✅ Pushed to GitHub${NC}"
fi
echo ""

# 3. Preparar Release
echo -e "${BLUE}[3/3] 🚀 Preparando release...${NC}"
bash scripts/release.sh

echo ""
echo -e "${GREEN}╔════════════════════════════════════════╗${NC}"
echo -e "${GREEN}║  ✅ BUILD & COMMIT COMPLETE!           ║${NC}"
echo -e "${GREEN}╚════════════════════════════════════════╝${NC}"
echo ""
echo -e "${YELLOW}📋 Próximos passos:${NC}"
echo -e "  1. ${BLUE}Publicar release manualmente no GitHub${NC}"
echo -e "     ${GREEN}https://github.com/peder1981/bagus-browser-go/releases/new${NC}"
echo -e "  2. ${BLUE}Fazer upload dos arquivos de dist/${NC}"
echo -e "  3. ${BLUE}Testar instalação${NC}"
echo -e "  4. ${BLUE}Compartilhar com usuários!${NC}"
echo ""
