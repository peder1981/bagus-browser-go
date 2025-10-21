#!/bin/bash
set -e

# Cores para output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Informações do projeto
APP_NAME="bagus-browser"
VERSION=$(git describe --tags --abbrev=0 2>/dev/null || echo "v4.1.0")
BUILD_DIR="build"
DIST_DIR="dist"

echo -e "${BLUE}🌐 Bagus Browser - Build & Package Script${NC}"
echo -e "${BLUE}=========================================${NC}"
echo ""
echo -e "${GREEN}Versão: ${VERSION}${NC}"
echo ""

# Limpar builds anteriores
echo -e "${YELLOW}🗑️  Limpando builds anteriores...${NC}"
rm -rf ${BUILD_DIR} ${DIST_DIR}
mkdir -p ${BUILD_DIR} ${DIST_DIR}

# Verificar dependências
echo -e "${YELLOW}🔍 Verificando dependências...${NC}"
if ! command -v go &> /dev/null; then
    echo -e "${RED}❌ Go não encontrado. Instale Go 1.23+${NC}"
    exit 1
fi

if ! pkg-config --exists gtk+-3.0; then
    echo -e "${RED}❌ GTK3 não encontrado. Instale: sudo apt-get install libgtk-3-dev${NC}"
    exit 1
fi

if ! pkg-config --exists webkit2gtk-4.0; then
    echo -e "${RED}❌ WebKit2GTK não encontrado. Instale: sudo apt-get install libwebkit2gtk-4.0-dev${NC}"
    exit 1
fi

echo -e "${GREEN}✅ Todas as dependências OK${NC}"
echo ""

# Compilar
echo -e "${YELLOW}🔨 Compilando ${APP_NAME}...${NC}"
go build -ldflags="-s -w" -o ${BUILD_DIR}/${APP_NAME} .

if [ $? -eq 0 ]; then
    echo -e "${GREEN}✅ Compilação bem-sucedida!${NC}"
    SIZE=$(du -h ${BUILD_DIR}/${APP_NAME} | cut -f1)
    echo -e "${GREEN}   Tamanho: ${SIZE}${NC}"
else
    echo -e "${RED}❌ Erro na compilação${NC}"
    exit 1
fi
echo ""

# Criar .desktop file
echo -e "${YELLOW}📝 Criando arquivo .desktop...${NC}"
cat > ${BUILD_DIR}/${APP_NAME}.desktop <<EOF
[Desktop Entry]
Version=1.0
Type=Application
Name=Bagus Browser
Comment=Browser minimalista, seguro e privado
Exec=${APP_NAME}
Icon=${APP_NAME}
Terminal=false
Categories=Network;WebBrowser;
Keywords=browser;web;internet;
StartupNotify=true
EOF

echo -e "${GREEN}✅ Arquivo .desktop criado${NC}"
echo ""

# Criar estrutura Debian
echo -e "${YELLOW}📦 Criando pacote .deb...${NC}"
DEB_DIR="${BUILD_DIR}/${APP_NAME}_${VERSION}_amd64"
mkdir -p ${DEB_DIR}/DEBIAN
mkdir -p ${DEB_DIR}/usr/bin
mkdir -p ${DEB_DIR}/usr/share/applications
mkdir -p ${DEB_DIR}/usr/share/doc/${APP_NAME}

# Copiar binário
cp ${BUILD_DIR}/${APP_NAME} ${DEB_DIR}/usr/bin/
chmod +x ${DEB_DIR}/usr/bin/${APP_NAME}

# Copiar .desktop
cp ${BUILD_DIR}/${APP_NAME}.desktop ${DEB_DIR}/usr/share/applications/

# Copiar documentação
cp README.md ${DEB_DIR}/usr/share/doc/${APP_NAME}/
cp LICENSE ${DEB_DIR}/usr/share/doc/${APP_NAME}/
cp CHANGELOG.md ${DEB_DIR}/usr/share/doc/${APP_NAME}/ 2>/dev/null || true

# Criar arquivo control
cat > ${DEB_DIR}/DEBIAN/control <<EOF
Package: ${APP_NAME}
Version: ${VERSION#v}
Section: web
Priority: optional
Architecture: amd64
Depends: libgtk-3-0, libwebkit2gtk-4.0-37
Maintainer: Peder <peder@example.com>
Description: Browser minimalista, seguro e privado
 Bagus Browser é um navegador web construído em Go com foco em:
 - Segurança (AES-256, validação de URLs)
 - Privacidade (zero telemetria, anti-fingerprinting)
 - Leveza (6.4MB)
 - Robustez (15 atalhos, múltiplas abas)
Homepage: https://github.com/peder1981/bagus-browser-go
EOF

# Construir .deb
dpkg-deb --build ${DEB_DIR}
mv ${DEB_DIR}.deb ${DIST_DIR}/

echo -e "${GREEN}✅ Pacote .deb criado: ${DIST_DIR}/${APP_NAME}_${VERSION}_amd64.deb${NC}"
echo ""

# Criar tarball
echo -e "${YELLOW}📦 Criando tarball...${NC}"
TAR_NAME="${APP_NAME}_${VERSION}_linux_amd64.tar.gz"
tar -czf ${DIST_DIR}/${TAR_NAME} \
    -C ${BUILD_DIR} ${APP_NAME} ${APP_NAME}.desktop \
    -C .. README.md LICENSE CHANGELOG.md 2>/dev/null || \
tar -czf ${DIST_DIR}/${TAR_NAME} \
    -C ${BUILD_DIR} ${APP_NAME} ${APP_NAME}.desktop \
    -C .. README.md LICENSE

echo -e "${GREEN}✅ Tarball criado: ${DIST_DIR}/${TAR_NAME}${NC}"
echo ""

# Criar checksums
echo -e "${YELLOW}🔐 Gerando checksums...${NC}"
cd ${DIST_DIR}
sha256sum *.deb *.tar.gz > SHA256SUMS
cd ..

echo -e "${GREEN}✅ Checksums gerados${NC}"
echo ""

# Resumo
echo -e "${BLUE}=========================================${NC}"
echo -e "${GREEN}✅ Build completo!${NC}"
echo ""
echo -e "${BLUE}📦 Pacotes criados:${NC}"
ls -lh ${DIST_DIR}/
echo ""
echo -e "${BLUE}📝 Checksums:${NC}"
cat ${DIST_DIR}/SHA256SUMS
echo ""
echo -e "${GREEN}🎉 Pronto para release!${NC}"
