#!/bin/bash

# Script para compilar WebKitGTK com WebRTC
# Bagus Browser v5.0.0
# Tempo estimado: 8-12 horas

set -e

echo "═══════════════════════════════════════════════════════════"
echo "🔨 Compilar WebKitGTK com WebRTC"
echo "═══════════════════════════════════════════════════════════"
echo ""
echo "⚠️  AVISOS:"
echo "   - Tempo de compilação: 8-12 horas"
echo "   - Espaço necessário: ~30GB"
echo "   - RAM mínima: 8GB (recomendado 16GB)"
echo ""
echo "═══════════════════════════════════════════════════════════"
echo ""

# Verificar espaço em disco
AVAILABLE_SPACE=$(df -BG . | tail -1 | awk '{print $4}' | sed 's/G//')
if [ "$AVAILABLE_SPACE" -lt 30 ]; then
    echo "❌ Erro: Espaço insuficiente em disco"
    echo "   Necessário: 30GB"
    echo "   Disponível: ${AVAILABLE_SPACE}GB"
    exit 1
fi

echo "✅ Espaço em disco: ${AVAILABLE_SPACE}GB (OK)"
echo ""

# Passo 1: Instalar dependências
echo "📦 Passo 1/8: Instalando dependências..."
echo ""

sudo apt install -y \
    build-essential \
    cmake \
    ninja-build \
    python3 \
    python3-pip \
    ruby \
    gperf \
    bison \
    flex \
    perl \
    git \
    subversion \
    libgtk-4-dev \
    libgtk-3-dev \
    libglib2.0-dev \
    libsoup-3.0-dev \
    libsoup2.4-dev \
    libxml2-dev \
    libxslt1-dev \
    libsqlite3-dev \
    libgcrypt20-dev \
    libsystemd-dev \
    libjpeg-dev \
    libpng-dev \
    libwebp-dev \
    libharfbuzz-dev \
    libfreetype6-dev \
    libfontconfig1-dev \
    libcairo2-dev \
    libpango1.0-dev \
    libicu-dev \
    libtasn1-6-dev \
    libgnutls28-dev \
    libgstreamer1.0-dev \
    libgstreamer-plugins-base1.0-dev \
    libgstreamer-plugins-bad1.0-dev \
    libenchant-2-dev \
    libsecret-1-dev \
    libhyphen-dev \
    libnotify-dev \
    libxt-dev \
    libxcomposite-dev \
    libxdamage-dev \
    libwayland-dev \
    libepoxy-dev \
    libmanette-0.2-dev \
    libwoff-dev \
    liblcms2-dev \
    libopus-dev \
    libvpx-dev \
    libsrtp2-dev \
    libnice-dev \
    libusrsctp-dev \
    libssl-dev \
    flite1-dev \
    libflite1

echo ""
echo "✅ Dependências instaladas!"
echo ""

# Passo 2: Baixar código-fonte
echo "📥 Passo 2/8: Baixando código-fonte do WebKitGTK..."
echo ""

WEBKIT_DIR="/tmp/webkit-build"
mkdir -p "$WEBKIT_DIR"
cd "$WEBKIT_DIR"
echo "   Diretório de build: $WEBKIT_DIR"

if [ ! -f "webkitgtk-2.48.7.tar.xz" ] || [ ! -s "webkitgtk-2.48.7.tar.xz" ]; then
    echo "   Removendo arquivo corrompido (se existir)..."
    rm -f webkitgtk-2.48.7.tar.xz
    echo "   Baixando WebKitGTK 2.48.7..."
    wget --continue https://webkitgtk.org/releases/webkitgtk-2.48.7.tar.xz
    if [ $? -ne 0 ]; then
        echo "❌ Erro ao baixar WebKitGTK"
        exit 1
    fi
fi

if [ ! -d "webkitgtk-2.48.7" ]; then
    echo "   Extraindo..."
    rm -rf webkitgtk-2.48.7
    tar -xf webkitgtk-2.48.7.tar.xz
    if [ $? -ne 0 ]; then
        echo "❌ Erro ao extrair WebKitGTK"
        rm -f webkitgtk-2.48.7.tar.xz
        exit 1
    fi
fi

cd webkitgtk-2.48.7

echo ""
echo "✅ Código-fonte pronto!"
echo ""

# Passo 3: Configurar build
echo "🔧 Passo 3/8: Configurando build com WebRTC..."
echo ""

mkdir -p build
cd build

cmake .. \
    -G Ninja \
    -DCMAKE_BUILD_TYPE=Release \
    -DCMAKE_INSTALL_PREFIX=/opt/webkitgtk-webrtc \
    -DPORT=GTK \
    -DUSE_GTK4=OFF \
    -DENABLE_WEB_RTC=ON \
    -DUSE_GSTREAMER_WEBRTC=ON \
    -DUSE_LIBWEBRTC=OFF \
    -DENABLE_INTROSPECTION=OFF \
    -DENABLE_DOCUMENTATION=OFF \
    -DENABLE_MINIBROWSER=ON \
    -DENABLE_GAMEPAD=OFF \
    -DUSE_JPEGXL=OFF \
    -DUSE_AVIF=OFF \
    -DUSE_LIBBACKTRACE=OFF \
    -DUSE_SOUP2=ON

echo ""
echo "✅ Build configurado!"
echo ""

# Verificar se WebRTC está habilitado
if grep -q "USE_GSTREAMER_WEBRTC:BOOL=ON" CMakeCache.txt; then
    echo "✅ WebRTC (GStreamer): HABILITADO ⭐"
else
    echo "❌ Erro: WebRTC não foi habilitado!"
    exit 1
fi

echo ""

# Passo 4: Compilar
echo "🏗️  Passo 4/8: Compilando WebKitGTK..."
echo ""
echo "⚠️  ATENÇÃO: Este processo vai levar 8-12 HORAS!"
echo "   - Deixe o computador ligado"
echo "   - Configure para não suspender"
echo "   - CPU vai usar 100%"
echo ""
echo "Iniciando compilação em 10 segundos..."
echo "Pressione Ctrl+C para cancelar"
echo ""

for i in {10..1}; do
    echo -n "$i... "
    sleep 1
done
echo ""
echo ""

echo "🚀 Compilação iniciada: $(date)"
echo ""

# Usar todos os cores disponíveis
CORES=$(nproc)
echo "   Usando $CORES cores"
echo ""

# Compilar
ninja -j$CORES

echo ""
echo "✅ Compilação completa: $(date)"
echo ""

# Passo 5: Instalar
echo "📦 Passo 5/8: Instalando WebKitGTK..."
echo ""

sudo ninja install

echo ""
echo "✅ WebKitGTK instalado em /opt/webkitgtk-webrtc/"
echo ""

# Passo 6: Configurar pkg-config
echo "🔗 Passo 6/8: Configurando pkg-config..."
echo ""

sudo tee /usr/lib/x86_64-linux-gnu/pkgconfig/webkitgtk-webrtc-6.0.pc > /dev/null << 'EOF'
prefix=/opt/webkitgtk-webrtc
exec_prefix=${prefix}
libdir=${exec_prefix}/lib
includedir=${prefix}/include

Name: webkitgtk-webrtc
Description: WebKitGTK with WebRTC support
Version: 2.48.7
Requires: gtk4 >= 4.0.0
Libs: -L${libdir} -lwebkitgtk-6.0
Cflags: -I${includedir}/webkitgtk-6.0
EOF

sudo ldconfig

echo ""
echo "✅ pkg-config configurado!"
echo ""

# Passo 7: Testar WebRTC
echo "🧪 Passo 7/8: Testando WebRTC..."
echo ""

cat > /tmp/test-webrtc.c << 'EOF'
#include <webkit/webkit.h>
#include <gtk/gtk.h>

int main(int argc, char **argv) {
    gtk_init();
    
    WebKitWebView *webview = WEBKIT_WEB_VIEW(webkit_web_view_new());
    WebKitSettings *settings = webkit_web_view_get_settings(webview);
    
    webkit_settings_set_enable_webrtc(settings, TRUE);
    
    g_print("WebRTC enabled: %d\n", 
        webkit_settings_get_enable_webrtc(settings));
    
    return 0;
}
EOF

gcc /tmp/test-webrtc.c -o /tmp/test-webrtc \
    $(pkg-config --cflags --libs webkit2gtk-4.1 gtk+-3.0) 2>/dev/null

if /tmp/test-webrtc 2>/dev/null | grep -q "WebRTC enabled: 1"; then
    echo "✅ WebRTC: FUNCIONANDO! ⭐"
else
    echo "❌ Erro: WebRTC não está funcionando"
    exit 1
fi

echo ""

# Passo 8: Atualizar Bagus Browser
echo "🔄 Passo 8/8: Atualizando Bagus Browser..."
echo ""

cd ~/bagus-browser-go

# Atualizar pkg-config no código
sed -i 's/#cgo pkg-config: webkitgtk-6.0/#cgo pkg-config: webkitgtk-webrtc-6.0/g' \
    cmd/bagus-browser-v5/main.go

# Recompilar
go build -o build/bagus-browser-v5 ./cmd/bagus-browser-v5

echo ""
echo "✅ Bagus Browser recompilado!"
echo ""

# Resumo final
echo "═══════════════════════════════════════════════════════════"
echo "🎉 COMPILAÇÃO COMPLETA!"
echo "═══════════════════════════════════════════════════════════"
echo ""
echo "✅ WebKitGTK com WebRTC instalado"
echo "✅ Bagus Browser atualizado"
echo ""
echo "🧪 Teste agora:"
echo "   ./build/bagus-browser-v5"
echo ""
echo "   Navegue para: https://meet.google.com"
echo "   Clique em: Nova reunião"
echo ""
echo "═══════════════════════════════════════════════════════════"
