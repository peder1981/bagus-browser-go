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
VERSION=$(git describe --tags --abbrev=0 2>/dev/null || echo "v4.2.0")
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

# Gerar ícones se necessário
echo -e "${YELLOW}🎨 Verificando ícones...${NC}"
if [ ! -f "assets/icons/bagus-browser-48.png" ]; then
    echo -e "${YELLOW}   Gerando ícones...${NC}"
    cd assets/icons
    ./generate-icons.sh
    cd ../..
fi
echo -e "${GREEN}✅ Ícones OK${NC}"
echo ""

# Criar .desktop file
echo -e "${YELLOW}📝 Criando arquivo .desktop...${NC}"
cat > ${BUILD_DIR}/${APP_NAME}.desktop <<EOF
[Desktop Entry]
Version=1.0
Type=Application
Name=Bagus Browser
GenericName=Web Browser
Comment=Browser minimalista, seguro e privado
Exec=${APP_NAME} %u
Icon=${APP_NAME}
Terminal=false
Categories=Network;WebBrowser;
Keywords=browser;web;internet;navegador;
MimeType=text/html;text/xml;application/xhtml+xml;x-scheme-handler/http;x-scheme-handler/https;
StartupNotify=true
StartupWMClass=bagus-browser
Actions=NewWindow;NewPrivateWindow;

[Desktop Action NewWindow]
Name=Nova Janela
Exec=${APP_NAME}

[Desktop Action NewPrivateWindow]
Name=Nova Janela Privada
Exec=${APP_NAME} --private
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
mkdir -p ${DEB_DIR}/usr/share/icons/hicolor/{16x16,22x22,24x24,32x32,48x48,64x64,128x128,256x256,512x512}/apps
mkdir -p ${DEB_DIR}/usr/share/pixmaps

# Copiar binário
cp ${BUILD_DIR}/${APP_NAME} ${DEB_DIR}/usr/bin/
chmod +x ${DEB_DIR}/usr/bin/${APP_NAME}

# Copiar .desktop
cp ${BUILD_DIR}/${APP_NAME}.desktop ${DEB_DIR}/usr/share/applications/

# Copiar ícones
echo -e "${YELLOW}🎨 Copiando ícones...${NC}"
for size in 16 22 24 32 48 64 128 256 512; do
    if [ -f "assets/icons/${APP_NAME}-${size}.png" ]; then
        cp "assets/icons/${APP_NAME}-${size}.png" \
           "${DEB_DIR}/usr/share/icons/hicolor/${size}x${size}/apps/${APP_NAME}.png"
    fi
done
# Copiar para pixmaps (fallback)
cp "assets/icons/${APP_NAME}-48.png" "${DEB_DIR}/usr/share/pixmaps/${APP_NAME}.png" 2>/dev/null || true
echo -e "${GREEN}✅ Ícones copiados${NC}"

# Copiar documentação
cp README.md ${DEB_DIR}/usr/share/doc/${APP_NAME}/
cp LICENSE ${DEB_DIR}/usr/share/doc/${APP_NAME}/
cp CHANGELOG.md ${DEB_DIR}/usr/share/doc/${APP_NAME}/ 2>/dev/null || true

# Criar script postinst (pós-instalação)
cat > ${DEB_DIR}/DEBIAN/postinst <<'POSTINST'
#!/bin/bash
set -e

echo "🔧 Configurando Bagus Browser..."

# Atualizar cache de ícones (múltiplas tentativas)
echo "  📦 Atualizando cache de ícones..."
if command -v gtk-update-icon-cache &> /dev/null; then
    gtk-update-icon-cache -f -t /usr/share/icons/hicolor 2>/dev/null || true
    gtk-update-icon-cache -f /usr/share/icons/hicolor 2>/dev/null || true
fi

if command -v update-icon-caches &> /dev/null; then
    update-icon-caches /usr/share/icons/hicolor 2>/dev/null || true
fi

# Atualizar banco de dados de aplicações
echo "  📦 Atualizando banco de dados de aplicações..."
if command -v update-desktop-database &> /dev/null; then
    update-desktop-database /usr/share/applications 2>/dev/null || true
    update-desktop-database 2>/dev/null || true
fi

# Atualizar cache MIME
echo "  📦 Atualizando cache MIME..."
if command -v update-mime-database &> /dev/null; then
    update-mime-database /usr/share/mime 2>/dev/null || true
fi

# Atualizar cache do desktop (GNOME/KDE)
echo "  📦 Atualizando cache do desktop..."
if command -v xdg-desktop-menu &> /dev/null; then
    xdg-desktop-menu forceupdate 2>/dev/null || true
fi

# Registrar aplicação no sistema
if command -v xdg-mime &> /dev/null; then
    xdg-mime default bagus-browser.desktop text/html 2>/dev/null || true
    xdg-mime default bagus-browser.desktop x-scheme-handler/http 2>/dev/null || true
    xdg-mime default bagus-browser.desktop x-scheme-handler/https 2>/dev/null || true
fi

# Forçar atualização do menu (para diferentes DEs)
if [ -d /usr/share/applications ]; then
    touch /usr/share/applications 2>/dev/null || true
fi

# Limpar cache do usuário (se executado como root)
if [ "$SUDO_USER" ]; then
    USER_HOME=$(getent passwd "$SUDO_USER" | cut -d: -f6)
    if [ -d "$USER_HOME/.cache/icon-theme.cache" ]; then
        rm -f "$USER_HOME/.cache/icon-theme.cache" 2>/dev/null || true
    fi
fi

echo ""
echo "✅ Bagus Browser instalado com sucesso!"
echo ""
echo "📍 Como executar:"
echo "   • Terminal: bagus-browser"
echo "   • Menu: Procure por 'Bagus Browser' em Aplicações → Internet"
echo "   • Desktop: Clique com botão direito → Criar atalho"
echo ""
echo "💡 Se o ícone não aparecer imediatamente:"
echo "   • Faça logout/login"
echo "   • Ou execute: killall nautilus (GNOME) / killall plasmashell (KDE)"
echo ""

exit 0
POSTINST

chmod +x ${DEB_DIR}/DEBIAN/postinst

# Criar script postrm (pós-remoção)
cat > ${DEB_DIR}/DEBIAN/postrm <<'POSTRM'
#!/bin/bash
set -e

if [ "$1" = "remove" ] || [ "$1" = "purge" ]; then
    # Atualizar cache de ícones
    if command -v gtk-update-icon-cache &> /dev/null; then
        gtk-update-icon-cache -f -t /usr/share/icons/hicolor 2>/dev/null || true
    fi
    
    # Atualizar banco de dados de aplicações
    if command -v update-desktop-database &> /dev/null; then
        update-desktop-database /usr/share/applications 2>/dev/null || true
    fi
fi

exit 0
POSTRM

chmod +x ${DEB_DIR}/DEBIAN/postrm

# Criar arquivo triggers (para atualização automática de caches)
cat > ${DEB_DIR}/DEBIAN/triggers <<'TRIGGERS'
# Triggers para atualização automática de caches do sistema
interest-noawait /usr/share/icons/hicolor
interest-noawait /usr/share/applications
interest-noawait /usr/share/mime/packages
activate-noawait update-icon-caches
activate-noawait update-desktop-database
activate-noawait update-mime-database
TRIGGERS

# Criar arquivo control
cat > ${DEB_DIR}/DEBIAN/control <<EOF
Package: ${APP_NAME}
Version: ${VERSION#v}
Section: web
Priority: optional
Architecture: amd64
Depends: libgtk-3-0, libwebkit2gtk-4.0-37
Recommends: desktop-file-utils, shared-mime-info, hicolor-icon-theme
Suggests: gnome-icon-theme | oxygen-icon-theme
Maintainer: Peder <peder@example.com>
Description: Browser minimalista, seguro e privado
 Bagus Browser é um navegador web construído em Go com foco em:
 - Segurança (AES-256, validação de URLs)
 - Privacidade (zero telemetria, anti-fingerprinting)
 - Leveza (6.4MB)
 - Robustez (15 atalhos, múltiplas abas)
 .
 Features:
 - WebView completo com WebKit2GTK
 - Múltiplas abas independentes
 - Buscar na página (Ctrl+F)
 - Favoritos com criptografia AES-256 (Ctrl+D)
 - Gerenciador de downloads
 - Zoom (Ctrl++, Ctrl+-, Ctrl+0)
 - 15 atalhos de teclado
 - Ícone profissional integrado
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
