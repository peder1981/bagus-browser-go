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
./build.sh
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

# 3. Release
echo -e "${BLUE}[3/3] 🚀 Creating GitHub release...${NC}"
./release.sh
if [ $? -ne 0 ]; then
    echo -e "${RED}❌ Release failed${NC}"
    exit 1
fi

echo ""
echo -e "${GREEN}╔════════════════════════════════════════╗${NC}"
echo -e "${GREEN}║  ✅ PUBLISH COMPLETE!                  ║${NC}"
echo -e "${GREEN}╚════════════════════════════════════════╝${NC}"
echo ""
echo -e "${BLUE}Next steps:${NC}"
echo -e "  1. Check release: https://github.com/peder1981/bagus-browser-go/releases"
echo -e "  2. Test installation"
echo -e "  3. Share with users!"
echo ""
