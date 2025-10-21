#!/bin/bash
# Script para instalar CEF (Chromium Embedded Framework)

set -e

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

echo -e "${BLUE}╔════════════════════════════════════════╗${NC}"
echo -e "${BLUE}║   Instalação do CEF                   ║${NC}"
echo -e "${BLUE}╚════════════════════════════════════════╝${NC}"
echo ""

# Verifica se já está instalado
if [ -d "/opt/cef" ]; then
    echo -e "${YELLOW}CEF já está instalado em /opt/cef${NC}"
    read -p "Deseja reinstalar? (s/N): " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Ss]$ ]]; then
        echo -e "${GREEN}Usando instalação existente${NC}"
        exit 0
    fi
    sudo rm -rf /opt/cef
fi

# Instala dependências
echo -e "${YELLOW}► Instalando dependências...${NC}"
sudo apt-get update
sudo apt-get install -y \
    build-essential \
    cmake \
    git \
    wget \
    libgtk-3-dev \
    libglib2.0-dev \
    libnss3-dev \
    libatk1.0-dev \
    libatk-bridge2.0-dev \
    libcups2-dev \
    libdrm-dev \
    libxkbcommon-dev \
    libxcomposite-dev \
    libxdamage-dev \
    libxrandr-dev \
    libgbm-dev \
    libpango1.0-dev \
    libasound2-dev \
    libx11-dev

echo -e "${GREEN}✓ Dependências instaladas${NC}"
echo ""

# Download CEF
echo -e "${YELLOW}► Baixando CEF...${NC}"
cd /tmp

# CEF versão 120 (Chromium 120)
CEF_VERSION="120.1.10+g3ce3184+chromium-120.0.6099.129"
CEF_FILE="cef_binary_${CEF_VERSION}_linux64.tar.bz2"
CEF_URL="https://cef-builds.spotifycdn.com/${CEF_FILE}"

if [ ! -f "$CEF_FILE" ]; then
    wget "$CEF_URL"
else
    echo -e "${YELLOW}Arquivo já existe, pulando download${NC}"
fi

echo -e "${GREEN}✓ CEF baixado${NC}"
echo ""

# Extrai CEF
echo -e "${YELLOW}► Extraindo CEF...${NC}"
tar -xjf "$CEF_FILE"
CEF_DIR=$(ls -d cef_binary_* | head -n 1)

echo -e "${GREEN}✓ CEF extraído${NC}"
echo ""

# Move para /opt
echo -e "${YELLOW}► Instalando em /opt/cef...${NC}"
sudo mv "$CEF_DIR" /opt/cef
sudo chown -R $USER:$USER /opt/cef

echo -e "${GREEN}✓ CEF instalado${NC}"
echo ""

# Build CEF
echo -e "${YELLOW}► Compilando CEF (isso pode demorar)...${NC}"
cd /opt/cef
mkdir -p build
cd build

cmake -G "Unix Makefiles" \
    -DCMAKE_BUILD_TYPE=Release \
    -DUSE_SANDBOX=OFF \
    ..

make -j$(nproc) libcef_dll_wrapper

echo -e "${GREEN}✓ CEF compilado${NC}"
echo ""

# Verifica instalação
if [ -f "/opt/cef/build/libcef_dll_wrapper/libcef_dll_wrapper.a" ]; then
    echo -e "${GREEN}╔════════════════════════════════════════╗${NC}"
    echo -e "${GREEN}║   CEF Instalado com Sucesso!          ║${NC}"
    echo -e "${GREEN}╚════════════════════════════════════════╝${NC}"
    echo ""
    echo -e "${BLUE}Localização:${NC} /opt/cef"
    echo -e "${BLUE}Versão:${NC} $CEF_VERSION"
    echo ""
    echo -e "${YELLOW}Próximo passo:${NC} ./scripts/build_cef.sh"
else
    echo -e "${RED}✗ Erro na instalação do CEF${NC}"
    exit 1
fi
