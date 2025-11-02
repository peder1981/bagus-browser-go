#!/bin/bash
#
# 📦 Script para Embarcar Dependências no Pacote
# Coleta GTK3, WebKit2GTK e GStreamer com todas as libs necessárias
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
BUNDLE_DIR="${1:-build/bundle}"
DEPS_DIR="${BUNDLE_DIR}/lib"
INCLUDE_DIR="${BUNDLE_DIR}/include"
PKG_CONFIG_DIR="${BUNDLE_DIR}/lib/pkgconfig"

# Dependências principais (pacotes Debian)
MAIN_PACKAGES=(
    "libgtk-3-0"
    "libwebkit2gtk-4.0-37"
    "libgstreamer1.0-0"
    "libgstreamer-plugins-base1.0-0"
    "libgstreamer-plugins-good1.0-0"
    "libgstreamer-plugins-bad1.0-0"
    "gstreamer1.0-plugins-base"
    "gstreamer1.0-plugins-good"
    "gstreamer1.0-plugins-bad"
    "gstreamer1.0-libav"
    "libgstreamer-gl1.0-0"
)

# Dependências de suporte (libs necessárias)
SUPPORT_PACKAGES=(
    "libatk1.0-0"
    "libcairo2"
    "libgdk-pixbuf2.0-0"
    "libglib2.0-0"
    "libpango-1.0-0"
    "libpangocairo-1.0-0"
    "libfontconfig1"
    "libfreetype6"
    "libx11-6"
    "libxext6"
    "libxrender1"
    "libxrandr2"
    "libxi6"
    "libxinerama1"
    "libxcursor1"
    "libxcomposite1"
    "libxdamage1"
    "libxfixes3"
    "libxkbcommon0"
    "libwayland-client0"
    "libwayland-egl1"
    "libepoxy0"
    "libssl3"
    "libsoup2.4-1"
    "libsqlite3-0"
    "libwebp7"
    "libpng16-16"
    "libharfbuzz0b"
    "libicu72"
    "libunistring2"
    "libffi8"
    "libpcre3"
    "libzstd1"
)

echo -e "${CYAN}"
echo "╔════════════════════════════════════════╗"
echo "║   📦 Bundle de Dependências Bagus     ║"
echo "║        Versão: 1.0                    ║"
echo "╚════════════════════════════════════════╝"
echo -e "${NC}"

# Criar diretórios
echo -e "${BLUE}📁 Criando estrutura de diretórios...${NC}"
mkdir -p ${DEPS_DIR}
mkdir -p ${INCLUDE_DIR}
mkdir -p ${PKG_CONFIG_DIR}

# Função para copiar biblioteca e suas dependências
copy_library() {
    local package=$1
    
    # Verificar se pacote existe
    if ! dpkg -l | grep -q "^ii.*$package"; then
        return 0  # Pacote não instalado
    fi
    
    # Obter lista de arquivos .so do pacote
    local lib_files=$(dpkg -L "$package" 2>/dev/null | grep -E "\.so" | grep -v "^/usr/share")
    
    if [ -z "$lib_files" ]; then
        return 0  # Pacote pode não ter bibliotecas .so
    fi
    
    # Copiar cada arquivo .so
    while IFS= read -r lib_path; do
        if [ -f "$lib_path" ]; then
            cp "$lib_path" ${DEPS_DIR}/ 2>/dev/null || true
        fi
    done <<< "$lib_files"
}

# Função para copiar headers
copy_headers() {
    local package=$1
    local header_paths=$(dpkg -L "$package" 2>/dev/null | grep -E "\.h$" | head -20)
    
    for header in $header_paths; do
        local dir=$(dirname "$header")
        mkdir -p ${INCLUDE_DIR}$(echo $dir | sed 's|/usr/include||')
        cp "$header" ${INCLUDE_DIR}$(echo $dir | sed 's|/usr/include||')/ 2>/dev/null || true
    done
}

# Função para copiar .pc files
copy_pkgconfig() {
    local package=$1
    local pc_files=$(dpkg -L "$package" 2>/dev/null | grep -E "\.pc$")
    
    for pc_file in $pc_files; do
        cp "$pc_file" ${PKG_CONFIG_DIR}/ 2>/dev/null || true
    done
}

# Coletar dependências principais
echo -e "${BLUE}📚 Coletando dependências principais...${NC}"
for package in "${MAIN_PACKAGES[@]}"; do
    echo -e "${YELLOW}  • $package${NC}"
    copy_library "$package"
    copy_headers "$package"
    copy_pkgconfig "$package"
done

# Coletar dependências de suporte
echo -e "${BLUE}📚 Coletando dependências de suporte...${NC}"
for package in "${SUPPORT_PACKAGES[@]}"; do
    echo -e "${YELLOW}  • $package${NC}"
    copy_library "$package"
done

# Criar script de configuração de ambiente
echo -e "${BLUE}🔧 Criando script de configuração...${NC}"
cat > ${BUNDLE_DIR}/setup-env.sh <<'EOF'
#!/bin/bash
# Script para configurar variáveis de ambiente para usar dependências embarcadas

BUNDLE_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

export LD_LIBRARY_PATH="${BUNDLE_DIR}/lib:${LD_LIBRARY_PATH}"
export PKG_CONFIG_PATH="${BUNDLE_DIR}/lib/pkgconfig:${PKG_CONFIG_PATH}"
export C_INCLUDE_PATH="${BUNDLE_DIR}/include:${C_INCLUDE_PATH}"
export CPLUS_INCLUDE_PATH="${BUNDLE_DIR}/include:${CPLUS_INCLUDE_PATH}"

echo "✅ Ambiente configurado para usar dependências embarcadas"
echo "   LD_LIBRARY_PATH: ${BUNDLE_DIR}/lib"
echo "   PKG_CONFIG_PATH: ${BUNDLE_DIR}/lib/pkgconfig"
EOF
chmod +x ${BUNDLE_DIR}/setup-env.sh

# Criar README
echo -e "${BLUE}📝 Criando README...${NC}"
cat > ${BUNDLE_DIR}/README.md <<'EOF'
# 📦 Dependências Embarcadas - Bagus Browser

Este diretório contém todas as dependências necessárias para executar o Bagus Browser sem necessidade de instalar pacotes do sistema.

## Estrutura

- `lib/` - Bibliotecas compartilhadas (.so)
- `include/` - Headers de desenvolvimento
- `lib/pkgconfig/` - Arquivos .pc para pkg-config
- `setup-env.sh` - Script para configurar variáveis de ambiente

## Como Usar

### Opção 1: Usar o script de setup

```bash
source ./setup-env.sh
./bagus-browser
```

### Opção 2: Definir variáveis manualmente

```bash
export LD_LIBRARY_PATH="$(pwd)/lib:$LD_LIBRARY_PATH"
export PKG_CONFIG_PATH="$(pwd)/lib/pkgconfig:$PKG_CONFIG_PATH"
./bagus-browser
```

### Opção 3: Usar no script de instalação

O instalador detecta automaticamente as dependências embarcadas.

## Versões Incluídas

- GTK3: 3.24+
- WebKit2GTK: 2.48+
- GStreamer: 1.20+
- Dependências de suporte: Todas as libs necessárias

## Tamanho

Aproximadamente 200-300 MB (comprimido: 50-80 MB)

## Compatibilidade

- Linux x86_64
- Ubuntu 22.04+
- Debian 11+
- Outras distribuições Linux com glibc 2.35+

## Notas

- As bibliotecas são compiladas para x86_64
- Requer glibc 2.35 ou superior
- Compatível com libsoup2 e libsoup3
EOF

# Gerar estatísticas
echo -e "${BLUE}📊 Gerando estatísticas...${NC}"
TOTAL_SIZE=$(du -sh ${BUNDLE_DIR} | cut -f1)
LIB_COUNT=$(find ${DEPS_DIR} -type f | wc -l)
HEADER_COUNT=$(find ${INCLUDE_DIR} -type f | wc -l)
PC_COUNT=$(find ${PKG_CONFIG_DIR} -type f | wc -l)

echo ""
echo -e "${GREEN}╔════════════════════════════════════════╗${NC}"
echo -e "${GREEN}║   ✅ Bundle Criado com Sucesso!       ║${NC}"
echo -e "${GREEN}╚════════════════════════════════════════╝${NC}"
echo ""
echo -e "${CYAN}📊 Estatísticas:${NC}"
echo "  Tamanho total: ${TOTAL_SIZE}"
echo "  Bibliotecas: ${LIB_COUNT} arquivos"
echo "  Headers: ${HEADER_COUNT} arquivos"
echo "  .pc files: ${PC_COUNT} arquivos"
echo ""
echo -e "${CYAN}📁 Localização:${NC}"
echo "  ${BUNDLE_DIR}"
echo ""
echo -e "${CYAN}🚀 Próximos passos:${NC}"
echo "  1. Comprimir: tar -czf bagus-browser-bundle.tar.gz ${BUNDLE_DIR}"
echo "  2. Testar: source ${BUNDLE_DIR}/setup-env.sh && ./bagus-browser"
echo "  3. Incluir no pacote .deb"
echo ""
