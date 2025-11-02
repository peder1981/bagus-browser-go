# 🔧 Configuração WebKit com WebRTC - Bagus Browser v5.0.0

## 📋 Visão Geral

Este documento descreve como o Bagus Browser v5.0.0 é compilado e executado com suporte a WebRTC via WebKit compilado.

---

## 🏗️ Estrutura de Compilação

### 1. **Build do WebKit com WebRTC** (`build-webkit-webrtc.sh`)

```bash
sudo bash ./scripts/build-webkit-webrtc.sh
```

**Resultado:**
- Instala em: `/opt/webkitgtk-webrtc/`
- Estrutura:
  ```
  /opt/webkitgtk-webrtc/
  ├── lib/
  │   ├── libwebkit2gtk-4.0.so      (GTK3)
  │   ├── libwebkitgtk-6.0.so       (GTK4)
  │   └── pkgconfig/
  │       ├── webkit2gtk-4.0.pc     (GTK3)
  │       └── webkitgtk-6.0.pc      (GTK4)
  ├── include/
  │   ├── webkitgtk-4.0/            (GTK3)
  │   └── webkitgtk-6.0/            (GTK4)
  └── bin/
      └── MiniBrowser               (Teste)
  ```

**Flags CMake:**
```cmake
-DENABLE_WEB_RTC=ON              # ✅ Habilita WebRTC
-DUSE_GSTREAMER_WEBRTC=ON        # ✅ Usa GStreamer para WebRTC
-DENABLE_MINIBROWSER=ON          # ✅ Habilita MiniBrowser
```

---

## 🔨 Build do Bagus Browser (`scripts/bagus build`)

### Detecção Automática

O script `bagus` detecta automaticamente qual versão do WebKit foi compilada:

```bash
if [ -d "/opt/webkitgtk-webrtc" ]; then
    if [ -f "/opt/webkitgtk-webrtc/lib/pkgconfig/webkitgtk-6.0.pc" ]; then
        # GTK4 + WebRTC
        export PKG_CONFIG_PATH="/opt/webkitgtk-webrtc/lib/pkgconfig:$PKG_CONFIG_PATH"
        export LD_LIBRARY_PATH="/opt/webkitgtk-webrtc/lib:$LD_LIBRARY_PATH"
        export CGO_CFLAGS="-I/opt/webkitgtk-webrtc/include"
        export CGO_LDFLAGS="-L/opt/webkitgtk-webrtc/lib"
    elif [ -f "/opt/webkitgtk-webrtc/lib/pkgconfig/webkit2gtk-4.0.pc" ]; then
        # GTK3 + WebRTC
        export PKG_CONFIG_PATH="/opt/webkitgtk-webrtc/lib/pkgconfig:$PKG_CONFIG_PATH"
        export LD_LIBRARY_PATH="/opt/webkitgtk-webrtc/lib:$LD_LIBRARY_PATH"
        export CGO_CFLAGS="-I/opt/webkitgtk-webrtc/include"
        export CGO_LDFLAGS="-L/opt/webkitgtk-webrtc/lib"
    fi
fi
```

### Variáveis de Ambiente

| Variável | Valor | Propósito |
|----------|-------|----------|
| `PKG_CONFIG_PATH` | `/opt/webkitgtk-webrtc/lib/pkgconfig:...` | Encontrar `.pc` do WebKit compilado |
| `LD_LIBRARY_PATH` | `/opt/webkitgtk-webrtc/lib:...` | Carregar bibliotecas compiladas |
| `CGO_CFLAGS` | `-I/opt/webkitgtk-webrtc/include` | Headers do WebKit compilado |
| `CGO_LDFLAGS` | `-L/opt/webkitgtk-webrtc/lib` | Linkar com WebKit compilado |

---

## 🚀 Execução do Browser

### Scripts de Execução

Todos os scripts de execução configuram as variáveis de ambiente:

#### `scripts/run-v5.sh`
```bash
if [ -d "/opt/webkitgtk-webrtc" ]; then
    export LD_LIBRARY_PATH="/opt/webkitgtk-webrtc/lib:$LD_LIBRARY_PATH"
    export PKG_CONFIG_PATH="/opt/webkitgtk-webrtc/lib/pkgconfig:$PKG_CONFIG_PATH"
fi
./build/bagus-browser
```

#### `scripts/test-webrtc-local.sh`
```bash
if [ -d "/opt/webkitgtk-webrtc" ]; then
    export LD_LIBRARY_PATH="/opt/webkitgtk-webrtc/lib:$LD_LIBRARY_PATH"
    export PKG_CONFIG_PATH="/opt/webkitgtk-webrtc/lib/pkgconfig:$PKG_CONFIG_PATH"
fi
./build/bagus-browser
```

#### `scripts/test-meet.sh`
```bash
if [ -d "/opt/webkitgtk-webrtc" ]; then
    export LD_LIBRARY_PATH="/opt/webkitgtk-webrtc/lib:$LD_LIBRARY_PATH"
    export PKG_CONFIG_PATH="/opt/webkitgtk-webrtc/lib/pkgconfig:$PKG_CONFIG_PATH"
fi
./build/bagus-browser
```

---

## 📊 Fluxo Completo

```
1. Compilar WebKit com WebRTC
   └─ sudo bash ./scripts/build-webkit-webrtc.sh
      └─ Instala em /opt/webkitgtk-webrtc/

2. Build do Bagus Browser
   └─ ./scripts/bagus build
      └─ Detecta WebKit compilado
      └─ Configura PKG_CONFIG_PATH, LD_LIBRARY_PATH, CGO_*
      └─ Compila com WebKit + WebRTC

3. Executar Browser
   └─ ./scripts/run-v5.sh
      └─ Configura LD_LIBRARY_PATH
      └─ Executa ./build/bagus-browser
      └─ Carrega WebKit compilado com WebRTC

4. Testar WebRTC
   └─ file:///tmp/test-webrtc.html
      └─ RTCPeerConnection disponível ✅
```

---

## ✅ Verificação

### Verificar WebKit Compilado

```bash
# Verificar instalação
ls -la /opt/webkitgtk-webrtc/lib/

# Verificar .pc files
ls -la /opt/webkitgtk-webrtc/lib/pkgconfig/

# Testar MiniBrowser
/opt/webkitgtk-webrtc/bin/MiniBrowser
```

### Verificar Compilação do Bagus

```bash
# Verificar se foi compilado com WebKit correto
ldd ./build/bagus-browser | grep webkit

# Deve mostrar:
# libwebkit2gtk-4.0.so => /opt/webkitgtk-webrtc/lib/libwebkit2gtk-4.0.so
```

### Verificar WebRTC em Tempo de Execução

```bash
# Abrir browser
./scripts/run-v5.sh

# Navegar para: file:///tmp/test-webrtc.html
# Verificar se RTCPeerConnection está disponível
```

---

## 🔍 Troubleshooting

### WebKit compilado não é detectado

```bash
# Verificar se diretório existe
ls -la /opt/webkitgtk-webrtc/

# Verificar se .pc files existem
ls -la /opt/webkitgtk-webrtc/lib/pkgconfig/
```

### Erro ao compilar Bagus Browser

```bash
# Limpar e recompilar
./scripts/bagus clean
./scripts/bagus build

# Verificar variáveis de ambiente
echo $PKG_CONFIG_PATH
echo $LD_LIBRARY_PATH
```

### WebRTC não funciona após compilar

```bash
# Verificar se bibliotecas estão carregadas
ldd ./build/bagus-browser | grep webrtc

# Executar com debug
LD_DEBUG=libs ./build/bagus-browser 2>&1 | grep webkit
```

---

## 📝 Notas

- **Tempo de compilação:** 8-12 horas
- **Espaço necessário:** ~30GB
- **RAM recomendada:** 16GB
- **Versão WebKit:** 2.48.7
- **Versão GTK:** 4.6.9 (compilado com GTK4)

---

**Última atualização:** 31/10/2025  
**Versão:** v5.0.0
