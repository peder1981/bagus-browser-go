#!/bin/bash
#
# 📦 Script para Criar .deb com Dependências Embarcadas
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

# Configurações
APP_NAME="bagus-browser"
VERSION="${1:-v5.0.0}"
BUILD_DIR="build"
DIST_DIR="dist"
SOURCE_DIR="cmd/bagus-browser"

echo -e "${CYAN}"
echo "╔════════════════════════════════════════╗"
echo "║   📦 Build .deb com Dependências      ║"
echo "║        Versão: ${VERSION}              ║"
echo "╚════════════════════════════════════════╝"
echo -e "${NC}"

# Verificar se binário existe
if [ ! -f "${BUILD_DIR}/${APP_NAME}" ]; then
    echo -e "${RED}❌ Binário não encontrado: ${BUILD_DIR}/${APP_NAME}${NC}"
    echo "Execute primeiro: ./scripts/bagus build"
    exit 1
fi

# Verificar se bundle existe
if [ ! -d "${BUILD_DIR}/bundle" ]; then
    echo -e "${RED}❌ Bundle não encontrado: ${BUILD_DIR}/bundle${NC}"
    echo "Execute primeiro: ./scripts/bagus bundle"
    exit 1
fi

echo -e "${BLUE}📦 Criando pacote .deb com dependências embarcadas...${NC}"

# Criar estrutura do .deb
DEB_DIR="${BUILD_DIR}/deb-bundled"
rm -rf ${DEB_DIR}
mkdir -p ${DEB_DIR}/DEBIAN
mkdir -p ${DEB_DIR}/usr/bin
mkdir -p ${DEB_DIR}/usr/lib/${APP_NAME}
mkdir -p ${DEB_DIR}/usr/share/applications
mkdir -p ${DEB_DIR}/usr/share/icons/hicolor/48x48/apps
mkdir -p ${DEB_DIR}/usr/share/icons/hicolor/128x128/apps

# Copiar binário
echo -e "${YELLOW}📝 Copiando binário...${NC}"
cp ${BUILD_DIR}/${APP_NAME} ${DEB_DIR}/usr/bin/
chmod +x ${DEB_DIR}/usr/bin/${APP_NAME}

# Copiar dependências embarcadas
echo -e "${YELLOW}📚 Copiando dependências embarcadas...${NC}"
cp -r ${BUILD_DIR}/bundle ${DEB_DIR}/usr/lib/${APP_NAME}/
chmod -R 755 ${DEB_DIR}/usr/lib/${APP_NAME}/bundle

# Copiar .desktop
echo -e "${YELLOW}📝 Copiando .desktop...${NC}"
cp ${BUILD_DIR}/${APP_NAME}.desktop ${DEB_DIR}/usr/share/applications/

# Copiar ícones
echo -e "${YELLOW}🎨 Copiando ícones...${NC}"
if [ -f "assets/icons/bagus-browser-48.png" ]; then
    cp assets/icons/bagus-browser-48.png ${DEB_DIR}/usr/share/icons/hicolor/48x48/apps/${APP_NAME}.png
fi
if [ -f "assets/icons/bagus-browser-128.png" ]; then
    cp assets/icons/bagus-browser-128.png ${DEB_DIR}/usr/share/icons/hicolor/128x128/apps/${APP_NAME}.png
fi

# Criar control file
echo -e "${YELLOW}📝 Criando control file...${NC}"
cat > ${DEB_DIR}/DEBIAN/control <<EOF
Package: bagus-browser
Version: ${VERSION#v}
Section: web
Priority: optional
Architecture: amd64
Depends: libc6 (>= 2.35)
Maintainer: Bagus Browser Team <team@bagus-browser.dev>
Homepage: https://github.com/peder1981/bagus-browser-go
Description: Browser minimalista, seguro e privado
 Bagus Browser é um navegador web focado em privacidade,
 segurança e simplicidade. Zero telemetria, zero rastreamento.
 .
 Inclui todas as dependências necessárias (GTK3, WebKit2GTK, GStreamer).
 Não requer instalação de pacotes adicionais do sistema.
EOF

# Criar postinst script
echo -e "${YELLOW}📝 Criando postinst script...${NC}"
cat > ${DEB_DIR}/DEBIAN/postinst <<'EOF'
#!/bin/bash
set -e

echo "🔧 Configurando Bagus Browser..."

# Atualizar cache de ícones
update-icon-caches /usr/share/icons/hicolor/ 2>/dev/null || true

# Atualizar banco de dados de aplicações
update-desktop-database 2>/dev/null || true

# Criar symlink para wrapper (opcional)
if [ ! -L /usr/bin/bagus-browser-bundled ]; then
    ln -s /usr/lib/bagus-browser/bundle/setup-env.sh /usr/bin/bagus-browser-bundled 2>/dev/null || true
fi

echo "✅ Bagus Browser instalado com sucesso!"
echo ""
echo "Para executar:"
echo "  • Terminal: bagus-browser"
echo "  • Menu: Aplicações → Internet → Bagus Browser"
echo ""
echo "Dependências embarcadas em: /usr/lib/bagus-browser/bundle"
EOF
chmod +x ${DEB_DIR}/DEBIAN/postinst

# Criar postrm script para limpeza
echo -e "${YELLOW}📝 Criando postrm script...${NC}"
cat > ${DEB_DIR}/DEBIAN/postrm <<'EOF'
#!/bin/bash
set -e

if [ "$1" = "remove" ] || [ "$1" = "purge" ]; then
    # Remover symlink se existir
    rm -f /usr/bin/bagus-browser-bundled 2>/dev/null || true
    
    # Atualizar cache de ícones
    update-icon-caches /usr/share/icons/hicolor/ 2>/dev/null || true
    update-desktop-database 2>/dev/null || true
fi
EOF
chmod +x ${DEB_DIR}/DEBIAN/postrm

# Criar preinst para verificações
echo -e "${YELLOW}📝 Criando preinst script...${NC}"
cat > ${DEB_DIR}/DEBIAN/preinst <<'EOF'
#!/bin/bash
set -e

# Verificar compatibilidade de glibc
GLIBC_VERSION=$(ldd --version | head -n1 | awk '{print $NF}')
REQUIRED_GLIBC="2.35"

if ! dpkg --compare-versions "$GLIBC_VERSION" "ge" "$REQUIRED_GLIBC"; then
    echo "❌ Erro: Bagus Browser requer glibc >= $REQUIRED_GLIBC"
    echo "   Seu sistema tem: $GLIBC_VERSION"
    exit 1
fi

echo "✅ Verificações pré-instalação OK"
EOF
chmod +x ${DEB_DIR}/DEBIAN/preinst

# Calcular tamanho instalado
INSTALLED_SIZE=$(du -s ${DEB_DIR} | cut -f1)

# Atualizar control file com tamanho
sed -i "/^Depends:/a Installed-Size: ${INSTALLED_SIZE}" ${DEB_DIR}/DEBIAN/control

# Construir .deb
echo -e "${YELLOW}📦 Construindo pacote .deb...${NC}"
dpkg-deb --build ${DEB_DIR} ${DIST_DIR}/${APP_NAME}_${VERSION}_amd64_bundled.deb

if [ $? -eq 0 ]; then
    SIZE=$(du -h ${DIST_DIR}/${APP_NAME}_${VERSION}_amd64_bundled.deb | cut -f1)
    echo -e "${GREEN}✅ Pacote .deb criado: ${SIZE}${NC}"
else
    echo -e "${RED}❌ Erro ao criar pacote .deb${NC}"
    exit 1
fi

# Gerar checksum
echo -e "${YELLOW}🔐 Gerando checksum...${NC}"
cd ${DIST_DIR}
sha256sum ${APP_NAME}_${VERSION}_amd64_bundled.deb >> SHA256SUMS
cd ..

echo ""
echo -e "${GREEN}╔════════════════════════════════════════╗${NC}"
echo -e "${GREEN}║   ✅ Pacote .deb Criado com Sucesso!  ║${NC}"
echo -e "${GREEN}╚════════════════════════════════════════╝${NC}"
echo ""
echo -e "${CYAN}📦 Pacote:${NC}"
ls -lh ${DIST_DIR}/${APP_NAME}_${VERSION}_amd64_bundled.deb
echo ""
echo -e "${CYAN}🚀 Para instalar:${NC}"
echo "  sudo dpkg -i ${DIST_DIR}/${APP_NAME}_${VERSION}_amd64_bundled.deb"
echo ""
echo -e "${CYAN}📊 Informações:${NC}"
echo "  • Tamanho: $(du -h ${DIST_DIR}/${APP_NAME}_${VERSION}_amd64_bundled.deb | cut -f1)"
echo "  • Dependências embarcadas: /usr/lib/bagus-browser/bundle"
echo "  • Requer apenas: libc6 >= 2.35"
echo ""
