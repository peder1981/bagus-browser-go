# Bagus Browser - CEF Implementation

## Migração para CEF (Chromium Embedded Framework)

### Por que CEF?

CEF fornece um navegador Chromium completo, resolvendo TODAS as limitações:
- ✅ 100% compatibilidade com sites (Google, Facebook, Twitter, etc.)
- ✅ Múltiplas janelas/abas nativas
- ✅ DevTools integrado
- ✅ Extensões (opcional)
- ✅ Performance completa do Chrome

### Arquitetura

```
┌─────────────────────────────────────────┐
│ Go Application (main.go)                │
│ ├─ História (criptografada)             │
│ ├─ Configuração                         │
│ └─ Lógica de negócio                    │
└─────────────────────────────────────────┘
         ↓ CGO
┌─────────────────────────────────────────┐
│ C++ CEF Wrapper (cef_browser.cpp)       │
│ ├─ CefApp                               │
│ ├─ CefClient                            │
│ └─ CefBrowserHost                       │
└─────────────────────────────────────────┘
         ↓
┌─────────────────────────────────────────┐
│ Chromium Embedded Framework             │
│ (Navegador completo)                    │
└─────────────────────────────────────────┘
```

## Instalação

### 1. Dependências

#### Ubuntu/Debian
```bash
sudo apt-get update
sudo apt-get install -y \
    build-essential \
    cmake \
    git \
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
    libasound2-dev
```

### 2. Download CEF

```bash
cd /tmp
wget https://cef-builds.spotifycdn.com/cef_binary_120.1.10+g3ce3184+chromium-120.0.6099.129_linux64.tar.bz2
tar -xjf cef_binary_*.tar.bz2
sudo mv cef_binary_* /opt/cef
```

### 3. Build CEF

```bash
cd /opt/cef
mkdir build
cd build
cmake -G "Unix Makefiles" -DCMAKE_BUILD_TYPE=Release ..
make -j$(nproc)
```

### 4. Build Bagus Browser

```bash
cd ~/bagus-browser-go
./scripts/build_cef.sh
```

## Estrutura de Arquivos

```
bagus-browser-go/
├── cef/
│   ├── cef_browser.h          # Header C++
│   ├── cef_browser.cpp        # Implementação CEF
│   ├── cef_app.h              # CefApp
│   ├── cef_app.cpp            # Implementação CefApp
│   └── CMakeLists.txt         # Build config
├── internal/
│   └── cef/
│       ├── cef.go             # Bindings Go
│       └── cef_wrapper.go     # Wrapper Go
├── scripts/
│   └── build_cef.sh           # Script de build
└── main_cef.go                # Main com CEF
```

## Uso

```bash
# Build
./scripts/build_cef.sh

# Run
./build/bagus-cef
```

## Funcionalidades

### Completas
- ✅ Navegação em 100% dos sites
- ✅ Múltiplas abas
- ✅ DevTools (F12)
- ✅ História criptografada
- ✅ Favoritos
- ✅ Downloads
- ✅ Cookies
- ✅ JavaScript completo
- ✅ WebGL, WebRTC, etc.

### Planejadas
- 🔄 Extensões
- 🔄 Sincronização
- 🔄 Perfis múltiplos
- 🔄 VPN integrada

## Performance

- **Memória**: ~200MB base + ~100MB por aba
- **CPU**: Equivalente ao Chrome
- **Startup**: ~2-3 segundos

## Licença

CEF é BSD-licensed. Compatível com uso comercial.
