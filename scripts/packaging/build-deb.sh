#!/bin/bash
# Script para criar pacote .deb do Bagus Browser

set -e

VERSION="1.0.0"
ARCH="amd64"
PKG_NAME="bagus-browser"
BUILD_DIR="build/deb"

echo "🔨 Criando pacote .deb do Bagus Browser..."

# Limpa build anterior
rm -rf "$BUILD_DIR"

# Cria estrutura do pacote
mkdir -p "$BUILD_DIR/DEBIAN"
mkdir -p "$BUILD_DIR/usr/bin"
mkdir -p "$BUILD_DIR/usr/share/applications"
mkdir -p "$BUILD_DIR/usr/share/icons/hicolor/48x48/apps"
mkdir -p "$BUILD_DIR/usr/share/doc/$PKG_NAME"

# Compila o navegador se não existir
if [ ! -f "build/bagus" ]; then
    echo "► Compilando Bagus Browser..."
    go build -ldflags "-s -w" -o build/bagus .
fi

# Copia binário
cp build/bagus "$BUILD_DIR/usr/bin/bagus-browser"
chmod 755 "$BUILD_DIR/usr/bin/bagus-browser"

# Cria arquivo .desktop
cat > "$BUILD_DIR/usr/share/applications/bagus-browser.desktop" << 'EOF'
[Desktop Entry]
Version=1.0
Type=Application
Name=Bagus Browser
Comment=Navegador leve e privado
Exec=bagus-browser %U
Icon=bagus-browser
Terminal=false
Categories=Network;WebBrowser;
MimeType=text/html;text/xml;application/xhtml+xml;x-scheme-handler/http;x-scheme-handler/https;
Keywords=browser;web;internet;
EOF

# Cria ícone (placeholder - você pode substituir por um ícone real)
cat > "$BUILD_DIR/usr/share/icons/hicolor/48x48/apps/bagus-browser.png" << 'EOF'
# Placeholder - adicione um ícone PNG real aqui
EOF

# Cria arquivo de controle
cat > "$BUILD_DIR/DEBIAN/control" << EOF
Package: $PKG_NAME
Version: $VERSION
Section: web
Priority: optional
Architecture: $ARCH
Depends: libwebkit2gtk-4.0-37, libgtk-3-0
Maintainer: Bagus Browser Team <bagus@example.com>
Description: Navegador leve e focado em privacidade
 Bagus Browser é um navegador web leve, rápido e focado em privacidade.
 Características:
  - Zero telemetria
  - Bloqueador de ads integrado
  - Interface simples e intuitiva
  - Apenas 4MB
Homepage: https://github.com/peder1981/bagus-browser-go
EOF

# Cria copyright
cat > "$BUILD_DIR/usr/share/doc/$PKG_NAME/copyright" << 'EOF'
Format: https://www.debian.org/doc/packaging-manuals/copyright-format/1.0/
Upstream-Name: Bagus Browser
Source: https://github.com/peder1981/bagus-browser-go

Files: *
Copyright: 2025 Bagus Browser Team
License: MIT
 Permission is hereby granted, free of charge, to any person obtaining a copy
 of this software and associated documentation files (the "Software"), to deal
 in the Software without restriction, including without limitation the rights
 to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 copies of the Software, and to permit persons to whom the Software is
 furnished to do so, subject to the following conditions:
 .
 The above copyright notice and this permission notice shall be included in all
 copies or substantial portions of the Software.
 .
 THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 SOFTWARE.
EOF

# Cria changelog
cat > "$BUILD_DIR/usr/share/doc/$PKG_NAME/changelog.gz" << 'EOF'
bagus-browser (1.0.0) stable; urgency=low

  * Initial release
  * WebKit2GTK engine
  * Zero telemetry
  * Ad blocker integrated
  * Lightweight (4MB)

 -- Bagus Browser Team <bagus@example.com>  Mon, 20 Oct 2025 22:00:00 -0300
EOF
gzip -9 "$BUILD_DIR/usr/share/doc/$PKG_NAME/changelog.gz"

# Cria postinst script
cat > "$BUILD_DIR/DEBIAN/postinst" << 'EOF'
#!/bin/bash
set -e

# Atualiza cache de ícones
if [ -x /usr/bin/update-icon-caches ]; then
    update-icon-caches /usr/share/icons/hicolor
fi

# Atualiza cache de desktop
if [ -x /usr/bin/update-desktop-database ]; then
    update-desktop-database -q
fi

exit 0
EOF
chmod 755 "$BUILD_DIR/DEBIAN/postinst"

# Cria postrm script
cat > "$BUILD_DIR/DEBIAN/postrm" << 'EOF'
#!/bin/bash
set -e

if [ "$1" = "remove" ]; then
    # Atualiza cache de ícones
    if [ -x /usr/bin/update-icon-caches ]; then
        update-icon-caches /usr/share/icons/hicolor
    fi
    
    # Atualiza cache de desktop
    if [ -x /usr/bin/update-desktop-database ]; then
        update-desktop-database -q
    fi
fi

exit 0
EOF
chmod 755 "$BUILD_DIR/DEBIAN/postrm"

# Calcula tamanho instalado
INSTALLED_SIZE=$(du -sk "$BUILD_DIR" | cut -f1)
echo "Installed-Size: $INSTALLED_SIZE" >> "$BUILD_DIR/DEBIAN/control"

# Constrói o pacote
DEB_FILE="${PKG_NAME}_${VERSION}_${ARCH}.deb"
dpkg-deb --build "$BUILD_DIR" "build/$DEB_FILE"

echo ""
echo "✅ Pacote criado com sucesso!"
echo "📦 Arquivo: build/$DEB_FILE"
echo ""
echo "Para instalar:"
echo "  sudo dpkg -i build/$DEB_FILE"
echo "  sudo apt-get install -f  # Se houver dependências faltando"
echo ""
echo "Para testar:"
echo "  dpkg -c build/$DEB_FILE  # Lista conteúdo"
echo "  dpkg -I build/$DEB_FILE  # Mostra informações"
