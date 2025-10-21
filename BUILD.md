# 🔨 Guia de Compilação - Bagus Browser

## 🎯 Opções de Build

O Bagus Browser oferece **três formas de compilar**:

### 1️⃣ Menu Interativo (Recomendado)
```bash
make menu
# ou
./scripts/build_menu.sh
```

**Vantagens:**
- Interface amigável
- Guia passo a passo
- Detecta se CEF está instalado
- Oferece instalar CEF automaticamente

### 2️⃣ Makefile (Rápido)
```bash
# Ver opções
make help

# Compilar versão Webview (padrão)
make build

# Compilar versão CEF
make build-cef

# Instalar CEF
make install-cef

# Verificar privacidade
make verify-privacy
```

### 3️⃣ Scripts Diretos (Avançado)
```bash
# Versão Webview
./scripts/build.sh

# Versão CEF
./scripts/install_cef.sh  # Uma vez
./scripts/build_cef.sh

# Verificar privacidade
./scripts/verify_privacy.sh
```

---

## 📦 Versão Webview (Padrão)

### Características
- ✅ **Leve**: ~5MB
- ✅ **Rápido**: Compilação em ~1 minuto
- ⚠️ **Compatibilidade**: ~70% dos sites
- ❌ Google, Facebook, Twitter não funcionam

### Compilar
```bash
# Opção 1: Makefile
make build

# Opção 2: Script
./scripts/build.sh

# Opção 3: Go direto
go build -o build/bagus
```

### Executar
```bash
./build/bagus
# ou
make run
```

### Sites Compatíveis
- ✅ DuckDuckGo
- ✅ Wikipedia
- ✅ Stack Overflow
- ✅ YouTube
- ✅ Reddit
- ✅ Medium
- ✅ Maioria dos blogs

### Sites Incompatíveis
- ❌ Google (use DuckDuckGo)
- ❌ Facebook
- ❌ Twitter
- ❌ GitHub (algumas páginas)
- ❌ Instagram
- ❌ LinkedIn

---

## 🚀 Versão CEF (100% Compatibilidade)

### Características
- ✅ **100% compatibilidade**: TODOS os sites
- ✅ **Chromium completo**: Mesma engine do Chrome
- ✅ **DevTools**: F12 para debug
- ⚠️ **Maior**: ~165MB
- ⚠️ **Instalação**: ~30 minutos (primeira vez)

### Pré-requisitos
```bash
# Ubuntu/Debian
sudo apt-get install build-essential cmake libgtk-3-dev

# Fedora
sudo dnf install gcc-c++ cmake gtk3-devel

# Arch
sudo pacman -S base-devel cmake gtk3
```

### Instalar CEF (Uma vez)
```bash
# Opção 1: Makefile
make install-cef

# Opção 2: Script
./scripts/install_cef.sh
```

**Tempo:** 15-30 minutos  
**Download:** ~500MB  
**Espaço:** ~2GB

### Compilar
```bash
# Opção 1: Makefile
make build-cef

# Opção 2: Script
./scripts/build_cef.sh
```

**Tempo:** 2-5 minutos

### Executar
```bash
cd build
./run_bagus_cef.sh

# ou
make run-cef
```

### Sites Compatíveis
- ✅ **TODOS**: 100% de compatibilidade
- ✅ Google
- ✅ Facebook
- ✅ Twitter
- ✅ GitHub
- ✅ Instagram
- ✅ LinkedIn
- ✅ **Qualquer site**

---

## 🔒 Verificar Privacidade

### Zero Telemetria Garantida
```bash
# Opção 1: Makefile
make verify-privacy

# Opção 2: Script
./scripts/verify_privacy.sh
```

**O que verifica:**
- ❌ Google Analytics
- ❌ Analytics de terceiros
- ❌ Telemetria
- ❌ Crash reporting
- ❌ Tracking pixels
- ❌ Fingerprinting
- ❌ Conexões HTTP externas

**Resultado esperado:** ✅ ZERO telemetria

---

## 📊 Comparação Rápida

| Característica | Webview | CEF |
|----------------|---------|-----|
| **Tamanho** | 5MB | 165MB |
| **Compilação** | 1 min | 5 min |
| **Instalação** | Imediata | 30 min |
| **Compatibilidade** | 70% | 100% |
| **Google/Facebook** | ❌ | ✅ |
| **DevTools** | ❌ | ✅ |
| **Memória** | ~50MB | ~200MB |

---

## 🎯 Qual Escolher?

### Use Webview se:
- ✅ Quer algo leve e rápido
- ✅ Não precisa de Google/Facebook
- ✅ DuckDuckGo, Wikipedia são suficientes
- ✅ Quer começar imediatamente

### Use CEF se:
- ✅ Precisa de 100% compatibilidade
- ✅ Usa Google, Facebook, Twitter
- ✅ Quer DevTools (F12)
- ✅ Não se importa com tamanho maior

---

## 🛠️ Comandos Úteis

### Build
```bash
make menu              # Menu interativo
make build             # Webview (padrão)
make build-webview     # Webview explícito
make build-cef         # CEF
make install-cef       # Instalar CEF
```

### Executar
```bash
make run               # Webview
make run-cef           # CEF
```

### Testes
```bash
make test              # Testes unitários
make test-coverage     # Com coverage
make verify-privacy    # Zero telemetria
```

### Utilitários
```bash
make fmt               # Formatar código
make lint              # Linter
make clean             # Limpar build
make help              # Ver todas opções
```

---

## 📝 Notas

### Primeira Compilação
- **Webview**: ~1 minuto
- **CEF**: ~30 minutos (instalação) + 5 minutos (compilação)

### Compilações Subsequentes
- **Webview**: ~30 segundos
- **CEF**: ~2 minutos

### Espaço em Disco
- **Webview**: ~10MB (código + binário)
- **CEF**: ~2GB (CEF + código + binário)

### Dependências
- **Webview**: Go 1.21+, webview libs
- **CEF**: Go 1.21+, GCC 9+, CMake 3.19+, GTK3

---

## 🆘 Problemas Comuns

### "CEF não encontrado"
```bash
make install-cef
```

### "libcef.so not found"
```bash
cd build
./run_bagus_cef.sh  # Use o script
```

### Erro de compilação
```bash
# Instale dependências
sudo apt-get install build-essential cmake libgtk-3-dev
```

### Webview não compila
```bash
# Ubuntu/Debian
sudo apt-get install libwebkit2gtk-4.0-dev

# Fedora
sudo dnf install webkit2gtk3-devel
```

---

## 📖 Documentação Adicional

- [README.md](README.md) - Visão geral
- [INSTALL_CEF.md](INSTALL_CEF.md) - Instalação CEF detalhada
- [QUICKSTART_CEF.md](QUICKSTART_CEF.md) - Guia rápido CEF
- [PRIVACY.md](PRIVACY.md) - Política de privacidade
- [SECURITY.md](SECURITY.md) - Segurança

---

**Pronto para começar? Execute `make menu`!** 🚀
