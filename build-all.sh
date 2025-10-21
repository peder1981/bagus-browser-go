#!/bin/bash
# Script de build unificado para Bagus Browser
# Compila para Linux e Windows

set -e

VERSION="2.0.0"
APP_NAME="bagus"

# Cores
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m'

echo -e "${BLUE}╔═══════════════════════════════════════╗${NC}"
echo -e "${BLUE}║   BAGUS BROWSER - BUILD ALL          ║${NC}"
echo -e "${BLUE}╚═══════════════════════════════════════╝${NC}"
echo ""

# Limpa builds anteriores
echo -e "${YELLOW}► Limpando builds anteriores...${NC}"
rm -rf build dist
mkdir -p build dist

# Instala dependências
echo -e "${YELLOW}► Instalando dependências...${NC}"
go mod download
go mod tidy

# Build para Linux AMD64
echo -e "${GREEN}► Compilando para Linux AMD64...${NC}"
GOOS=linux GOARCH=amd64 go build -ldflags "-s -w -X main.version=${VERSION}" -o "build/${APP_NAME}-linux-amd64" .
chmod +x "build/${APP_NAME}-linux-amd64"

# Build para Windows AMD64 (CGO_ENABLED=0 para cross-compile)
echo -e "${GREEN}► Compilando para Windows AMD64...${NC}"
if GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-s -w -X main.version=${VERSION}" -o "build/${APP_NAME}-windows-amd64.exe" . 2>/dev/null; then
    echo -e "${GREEN}  ✓ Windows AMD64 compilado${NC}"
else
    echo -e "${YELLOW}  ⚠ Windows AMD64 falhou (webview requer CGO, use build nativo no Windows)${NC}"
fi

# Cria pacotes de distribuição
echo -e "${YELLOW}► Criando pacotes de distribuição...${NC}"

# Linux AMD64
cd build
tar -czf "../dist/${APP_NAME}-${VERSION}-linux-amd64.tar.gz" "${APP_NAME}-linux-amd64"
echo -e "${GREEN}  ✓ dist/${APP_NAME}-${VERSION}-linux-amd64.tar.gz${NC}"

# Linux ARM64 (se existir)
if [ -f "${APP_NAME}-linux-arm64" ]; then
    tar -czf "../dist/${APP_NAME}-${VERSION}-linux-arm64.tar.gz" "${APP_NAME}-linux-arm64"
    echo -e "${GREEN}  ✓ dist/${APP_NAME}-${VERSION}-linux-arm64.tar.gz${NC}"
fi

# Windows AMD64 (se existir)
if [ -f "${APP_NAME}-windows-amd64.exe" ]; then
    zip -q "../dist/${APP_NAME}-${VERSION}-windows-amd64.zip" "${APP_NAME}-windows-amd64.exe"
    echo -e "${GREEN}  ✓ dist/${APP_NAME}-${VERSION}-windows-amd64.zip${NC}"
fi

cd ..

# Cria pacote .deb para Debian/Ubuntu
echo -e "${YELLOW}► Criando pacote .deb...${NC}"
./build-deb.sh

echo ""
echo -e "${GREEN}╔═══════════════════════════════════════╗${NC}"
echo -e "${GREEN}║   BUILD CONCLUÍDO COM SUCESSO!        ║${NC}"
echo -e "${GREEN}╚═══════════════════════════════════════╝${NC}"
echo ""
echo -e "${BLUE}Arquivos gerados:${NC}"
echo -e "  📦 build/${APP_NAME}-linux-amd64"
echo -e "  📦 build/${APP_NAME}-linux-arm64"
echo -e "  📦 build/${APP_NAME}-windows-amd64.exe"
echo -e "  📦 build/bagus-browser_1.0.0_amd64.deb"
echo ""
echo -e "  🗜️  dist/${APP_NAME}-${VERSION}-linux-amd64.tar.gz"
echo -e "  🗜️  dist/${APP_NAME}-${VERSION}-linux-arm64.tar.gz"
echo -e "  🗜️  dist/${APP_NAME}-${VERSION}-windows-amd64.zip"
echo ""
echo -e "${YELLOW}Para testar:${NC}"
echo -e "  ./build/${APP_NAME}-linux-amd64"
echo ""
echo -e "${YELLOW}Para instalar (Linux):${NC}"
echo -e "  sudo cp build/${APP_NAME}-linux-amd64 /usr/local/bin/bagus"
echo -e "  sudo chmod +x /usr/local/bin/bagus"
echo ""
