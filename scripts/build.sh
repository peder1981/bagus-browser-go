#!/bin/bash
# Bagus Browser Go - Script de Build (Versão Webview)
# Uso: ./scripts/build.sh [linux|windows|macos|all]

set -e

# Cores
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Variáveis
APP_NAME="bagus"
VERSION="2.0.0-alpha"
BUILD_DIR="build"
CMD_DIR="."

# Banner
echo -e "${BLUE}╔════════════════════════════════════════╗${NC}"
echo -e "${BLUE}║   Bagus Browser - Versão Webview     ║${NC}"
echo -e "${BLUE}║   Leve | Rápido | 70% Compatibilidade║${NC}"
echo -e "${BLUE}╚════════════════════════════════════════╝${NC}"
echo ""
echo -e "${YELLOW}Nota: Para 100% compatibilidade, use: ./scripts/build_cef.sh${NC}"
echo ""

# Função para build
build_platform() {
    local os=$1
    local arch=$2
    local output=$3
    
    echo -e "${YELLOW}► Compilando para ${os}/${arch}...${NC}"
    
    GOOS=$os GOARCH=$arch go build \
        -ldflags "-s -w -X main.appVersion=${VERSION}" \
        -o "${BUILD_DIR}/${output}" \
        ./${CMD_DIR}
    
    if [ $? -eq 0 ]; then
        echo -e "${GREEN}✓ Build concluído: ${BUILD_DIR}/${output}${NC}"
    else
        echo -e "${RED}✗ Erro ao compilar${NC}"
        exit 1
    fi
}

# Criar diretório de build
mkdir -p ${BUILD_DIR}

# Instalar dependências
echo -e "${YELLOW}► Instalando dependências...${NC}"
go mod download
go mod tidy
echo -e "${GREEN}✓ Dependências instaladas${NC}"
echo ""

# Determinar plataforma
PLATFORM=${1:-current}

case $PLATFORM in
    linux)
        echo -e "${BLUE}═══ Build Linux ═══${NC}"
        build_platform "linux" "amd64" "${APP_NAME}-linux-amd64"
        build_platform "linux" "arm64" "${APP_NAME}-linux-arm64"
        ;;
    windows)
        echo -e "${BLUE}═══ Build Windows ═══${NC}"
        build_platform "windows" "amd64" "${APP_NAME}-windows-amd64.exe"
        ;;
    macos)
        echo -e "${BLUE}═══ Build macOS ═══${NC}"
        build_platform "darwin" "amd64" "${APP_NAME}-darwin-amd64"
        build_platform "darwin" "arm64" "${APP_NAME}-darwin-arm64"
        ;;
    all)
        echo -e "${BLUE}═══ Build Todas as Plataformas ═══${NC}"
        build_platform "linux" "amd64" "${APP_NAME}-linux-amd64"
        build_platform "linux" "arm64" "${APP_NAME}-linux-arm64"
        build_platform "windows" "amd64" "${APP_NAME}-windows-amd64.exe"
        build_platform "darwin" "amd64" "${APP_NAME}-darwin-amd64"
        build_platform "darwin" "arm64" "${APP_NAME}-darwin-arm64"
        ;;
    current)
        echo -e "${BLUE}═══ Build Plataforma Atual ═══${NC}"
        go build \
            -ldflags "-s -w -X main.appVersion=${VERSION}" \
            -o "${BUILD_DIR}/${APP_NAME}" \
            ./${CMD_DIR}
        echo -e "${GREEN}✓ Build concluído: ${BUILD_DIR}/${APP_NAME}${NC}"
        ;;
    *)
        echo -e "${RED}Plataforma inválida: $PLATFORM${NC}"
        echo "Uso: $0 [linux|windows|macos|all|current]"
        exit 1
        ;;
esac

echo ""
echo -e "${GREEN}╔════════════════════════════════════════╗${NC}"
echo -e "${GREEN}║      Build Concluído com Sucesso!     ║${NC}"
echo -e "${GREEN}╚════════════════════════════════════════╝${NC}"
echo ""
echo -e "${BLUE}Versão:${NC} Webview (Leve)"
echo -e "${BLUE}Tamanho:${NC} ~5MB"
echo -e "${BLUE}Compatibilidade:${NC} ~70% dos sites"
echo ""
echo -e "${GREEN}✅ Sites compatíveis:${NC}"
echo -e "   - DuckDuckGo, Wikipedia, Stack Overflow"
echo -e "   - YouTube, Reddit, Medium"
echo -e "   - Maioria dos blogs e sites"
echo ""
echo -e "${YELLOW}⚠️  Sites incompatíveis:${NC}"
echo -e "   - Google, Facebook, Twitter"
echo -e "   - GitHub (algumas páginas)"
echo -e "   - Sites com X-Frame-Options"
echo ""
echo -e "${BLUE}Para 100% compatibilidade:${NC}"
echo -e "   ./scripts/install_cef.sh  # Instalar CEF (uma vez)"
echo -e "   ./scripts/build_cef.sh    # Compilar versão CEF"
echo ""
