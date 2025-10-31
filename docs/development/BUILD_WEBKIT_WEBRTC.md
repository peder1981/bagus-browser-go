# 🔨 Compilar WebKitGTK com WebRTC - Guia Completo

**Objetivo**: Compilar WebKitGTK do zero com suporte completo a WebRTC  
**Tempo estimado**: 8-12 horas de compilação  
**Espaço necessário**: ~30GB

## ⚠️ Avisos Importantes

1. **Tempo de compilação**: 8-12 horas em máquina moderna
2. **Espaço em disco**: ~30GB necessários
3. **RAM**: Mínimo 8GB, recomendado 16GB
4. **CPU**: Quanto mais cores, melhor (usa todos os cores)

## 📋 Passo 1: Instalar Dependências

### Dependências de Build

```bash
sudo apt update
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
    subversion
```

### Dependências do WebKit

```bash
sudo apt install -y \
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
    libavif-dev \
    libjxl-dev
```

### Dependências WebRTC (ESSENCIAL)

```bash
sudo apt install -y \
    libopus-dev \
    libvpx-dev \
    libsrtp2-dev \
    libnice-dev \
    libusrsctp-dev \
    libopenssl-dev
```

## 📥 Passo 2: Baixar Código-Fonte

### Opção A: Tarball (Recomendado - Mais Rápido)

```bash
cd ~
mkdir webkit-build
cd webkit-build

# Baixar WebKitGTK 2.48.7 (mesma versão do sistema)
wget https://webkitgtk.org/releases/webkitgtk-2.48.7.tar.xz

# Extrair
tar -xf webkitgtk-2.48.7.tar.xz
cd webkitgtk-2.48.7
```

### Opção B: Git (Mais Lento, Mais Recente)

```bash
cd ~
mkdir webkit-build
cd webkit-build

# Clonar repositório (GRANDE - ~3GB)
git clone https://github.com/WebKit/WebKit.git
cd WebKit

# Checkout versão estável
git checkout webkitgtk-2.48.7
```

## 🔧 Passo 3: Configurar Build

### Criar diretório de build

```bash
mkdir build
cd build
```

### Configurar com CMake

```bash
cmake .. \
    -GNinja \
    -DCMAKE_BUILD_TYPE=Release \
    -DCMAKE_INSTALL_PREFIX=/opt/webkitgtk-webrtc \
    -DPORT=GTK \
    -DENABLE_WEBRTC=ON \
    -DENABLE_MEDIA_STREAM=ON \
    -DENABLE_WEB_AUDIO=ON \
    -DENABLE_VIDEO=ON \
    -DENABLE_WEB_CRYPTO=ON \
    -DENABLE_ENCRYPTED_MEDIA=ON \
    -DENABLE_GAMEPAD=ON \
    -DENABLE_MINIBROWSER=ON \
    -DUSE_LIBWEBRTC=ON \
    -DUSE_GSTREAMER_WEBRTC=ON \
    -DUSE_SOUP2=OFF \
    -DUSE_GTK4=ON
```

### Verificar Configuração

```bash
# Verificar se WebRTC está habilitado
grep -i "ENABLE_WEBRTC" CMakeCache.txt
# Deve mostrar: ENABLE_WEBRTC:BOOL=ON
```

## 🏗️ Passo 4: Compilar (8-12 horas)

### Iniciar Compilação

```bash
# Usar todos os cores disponíveis
ninja -j$(nproc)
```

**IMPORTANTE**: Este processo vai levar MUITAS horas!

### Monitorar Progresso

```bash
# Em outro terminal
watch -n 60 'ps aux | grep ninja'
```

### Se der erro de memória

```bash
# Reduzir jobs paralelos
ninja -j4  # Usar apenas 4 cores
```

## 📦 Passo 5: Instalar

```bash
sudo ninja install
```

Isso vai instalar em `/opt/webkitgtk-webrtc/`

## 🔗 Passo 6: Configurar pkg-config

### Criar arquivo .pc customizado

```bash
sudo nano /usr/lib/x86_64-linux-gnu/pkgconfig/webkitgtk-webrtc-6.0.pc
```

Conteúdo:
```
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
```

### Atualizar cache do pkg-config

```bash
sudo ldconfig
pkg-config --modversion webkitgtk-webrtc-6.0
```

## 🧪 Passo 7: Testar WebRTC

### Criar teste simples

```bash
cat > /tmp/test-webrtc.c << 'EOF'
#include <webkit/webkit.h>
#include <gtk/gtk.h>

int main(int argc, char **argv) {
    gtk_init(&argc, &argv);
    
    WebKitWebView *webview = WEBKIT_WEB_VIEW(webkit_web_view_new());
    WebKitSettings *settings = webkit_web_view_get_settings(webview);
    
    webkit_settings_set_enable_webrtc(settings, TRUE);
    
    g_print("WebRTC enabled: %d\n", 
        webkit_settings_get_enable_webrtc(settings));
    
    return 0;
}
EOF

# Compilar
gcc /tmp/test-webrtc.c -o /tmp/test-webrtc \
    $(pkg-config --cflags --libs webkitgtk-webrtc-6.0 gtk4)

# Executar
/tmp/test-webrtc
```

Deve mostrar: `WebRTC enabled: 1`

## 🔄 Passo 8: Atualizar Bagus Browser

### Atualizar go.mod (não necessário)

### Atualizar pkg-config no código

Editar `cmd/bagus-browser-v5/main.go`:

```go
/*
#cgo pkg-config: webkitgtk-webrtc-6.0 gtk4
#include <webkit/webkit.h>
#include <gtk/gtk.h>
*/
```

### Recompilar

```bash
cd ~/bagus-browser-go
go build -o build/bagus-browser-v5 ./cmd/bagus-browser-v5
```

### Executar e Testar

```bash
./build/bagus-browser-v5
```

Navegue para `https://meet.google.com` e teste!

## 📊 Checklist de Progresso

- [ ] Dependências instaladas
- [ ] Código-fonte baixado
- [ ] Build configurado
- [ ] Compilação iniciada
- [ ] Compilação completa (8-12h)
- [ ] WebKitGTK instalado
- [ ] pkg-config configurado
- [ ] Teste WebRTC passou
- [ ] Bagus Browser recompilado
- [ ] Google Meet funcionando ⭐

## ⚠️ Problemas Comuns

### Erro: Out of memory

**Solução**: Reduzir jobs paralelos
```bash
ninja -j2  # Usar apenas 2 cores
```

### Erro: Missing dependency

**Solução**: Instalar dependência faltante
```bash
sudo apt install <pacote-faltante>-dev
```

### Erro: WebRTC não habilitado

**Solução**: Verificar dependências WebRTC
```bash
sudo apt install libopus-dev libvpx-dev libsrtp2-dev libnice-dev
```

## 📝 Notas Importantes

1. **Backup**: Faça backup antes de começar
2. **Espaço**: Verifique espaço em disco (30GB+)
3. **Tempo**: Deixe compilando overnight
4. **Energia**: Configure para não suspender
5. **Ventilação**: CPU vai esquentar muito

## 🎯 Próximos Passos Após Compilação

1. Testar Google Meet
2. Testar Microsoft Teams
3. Testar Zoom
4. Criar script de build automatizado
5. Documentar processo
6. Criar pacote .deb customizado

---

**Status**: 🚧 Pronto para iniciar  
**Tempo estimado**: 8-12 horas  
**Dificuldade**: Alta  
**Recompensa**: WebRTC funcionando! 🎉
