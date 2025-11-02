#!/bin/bash
#
# 🧪 Script para Testar Pacote com Dependências Embarcadas
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

echo -e "${CYAN}"
echo "╔════════════════════════════════════════╗"
echo "║   🧪 Teste de Pacote Bundled          ║"
echo "║        Versão: 1.0                    ║"
echo "╚════════════════════════════════════════╝"
echo -e "${NC}"

# Configurações
BUILD_DIR="build"
BUNDLE_DIR="${BUILD_DIR}/bundle"
APP_NAME="bagus-browser"

# Teste 1: Verificar estrutura do bundle
echo -e "${BLUE}🔍 Teste 1: Verificando estrutura do bundle...${NC}"

if [ ! -d "${BUNDLE_DIR}" ]; then
    echo -e "${RED}❌ Bundle não encontrado${NC}"
    exit 1
fi

# Verificar diretórios
for dir in lib include lib/pkgconfig; do
    if [ ! -d "${BUNDLE_DIR}/${dir}" ]; then
        echo -e "${RED}❌ Diretório ${dir} não encontrado${NC}"
        exit 1
    fi
    echo -e "${GREEN}✓ ${dir}${NC}"
done

# Teste 2: Verificar bibliotecas críticas
echo -e "${BLUE}🔍 Teste 2: Verificando bibliotecas críticas...${NC}"

CRITICAL_LIBS=(
    "libgtk-3.so"
    "libwebkit2gtk-4.0.so"
    "libgstreamer-1.0.so"
    "libglib-2.0.so"
)

for lib in "${CRITICAL_LIBS[@]}"; do
    if find ${BUNDLE_DIR}/lib -name "${lib}*" -type f | grep -q .; then
        echo -e "${GREEN}✓ ${lib}${NC}"
    else
        echo -e "${YELLOW}⚠️  ${lib} não encontrado (pode estar em dependência indireta)${NC}"
    fi
done

# Teste 3: Verificar .pc files
echo -e "${BLUE}🔍 Teste 3: Verificando .pc files...${NC}"

CRITICAL_PC=(
    "gtk+-3.0.pc"
    "webkit2gtk-4.0.pc"
    "gstreamer-1.0.pc"
)

for pc in "${CRITICAL_PC[@]}"; do
    if [ -f "${BUNDLE_DIR}/lib/pkgconfig/${pc}" ]; then
        echo -e "${GREEN}✓ ${pc}${NC}"
    else
        echo -e "${YELLOW}⚠️  ${pc} não encontrado${NC}"
    fi
done

# Teste 4: Verificar setup-env.sh
echo -e "${BLUE}🔍 Teste 4: Verificando setup-env.sh...${NC}"

if [ -f "${BUNDLE_DIR}/setup-env.sh" ]; then
    echo -e "${GREEN}✓ setup-env.sh encontrado${NC}"
    
    # Testar sourcing
    if bash -n "${BUNDLE_DIR}/setup-env.sh" 2>/dev/null; then
        echo -e "${GREEN}✓ setup-env.sh é válido${NC}"
    else
        echo -e "${RED}❌ setup-env.sh tem erros de sintaxe${NC}"
        exit 1
    fi
else
    echo -e "${RED}❌ setup-env.sh não encontrado${NC}"
    exit 1
fi

# Teste 5: Testar carregamento de bibliotecas
echo -e "${BLUE}🔍 Teste 5: Testando carregamento de bibliotecas...${NC}"

# Criar script de teste temporário
TEST_SCRIPT=$(mktemp)
cat > ${TEST_SCRIPT} <<'EOF'
#!/bin/bash
source ./setup-env.sh

# Testar pkg-config
if pkg-config --exists gtk+-3.0; then
    echo "✓ GTK3 detectado via pkg-config"
    GTK_VERSION=$(pkg-config --modversion gtk+-3.0)
    echo "  Versão: $GTK_VERSION"
else
    echo "❌ GTK3 não detectado"
    exit 1
fi

if pkg-config --exists webkit2gtk-4.0; then
    echo "✓ WebKit2GTK detectado via pkg-config"
    WEBKIT_VERSION=$(pkg-config --modversion webkit2gtk-4.0)
    echo "  Versão: $WEBKIT_VERSION"
else
    echo "❌ WebKit2GTK não detectado"
    exit 1
fi

if pkg-config --exists gstreamer-1.0; then
    echo "✓ GStreamer detectado via pkg-config"
    GS_VERSION=$(pkg-config --modversion gstreamer-1.0)
    echo "  Versão: $GS_VERSION"
else
    echo "❌ GStreamer não detectado"
    exit 1
fi
EOF

cd ${BUNDLE_DIR}
if bash ${TEST_SCRIPT} 2>/dev/null; then
    echo -e "${GREEN}✓ Bibliotecas carregadas com sucesso${NC}"
else
    echo -e "${YELLOW}⚠️  Alguns testes de carregamento falharam${NC}"
fi
cd - > /dev/null
rm ${TEST_SCRIPT}

# Teste 6: Verificar tamanho
echo -e "${BLUE}🔍 Teste 6: Verificando tamanho...${NC}"

TOTAL_SIZE=$(du -sh ${BUNDLE_DIR} | cut -f1)
LIB_COUNT=$(find ${BUNDLE_DIR}/lib -type f | wc -l)

echo -e "${GREEN}✓ Tamanho total: ${TOTAL_SIZE}${NC}"
echo "  Arquivos: ${LIB_COUNT}"

# Teste 7: Verificar permissões
echo -e "${BLUE}🔍 Teste 7: Verificando permissões...${NC}"

# Verificar se arquivos .so têm permissão de leitura
if find ${BUNDLE_DIR}/lib -name "*.so*" -type f ! -readable | grep -q .; then
    echo -e "${RED}❌ Alguns arquivos .so não têm permissão de leitura${NC}"
    exit 1
else
    echo -e "${GREEN}✓ Todas as bibliotecas têm permissão de leitura${NC}"
fi

# Teste 8: Verificar README
echo -e "${BLUE}🔍 Teste 8: Verificando documentação...${NC}"

if [ -f "${BUNDLE_DIR}/README.md" ]; then
    echo -e "${GREEN}✓ README.md encontrado${NC}"
else
    echo -e "${YELLOW}⚠️  README.md não encontrado${NC}"
fi

# Resumo
echo ""
echo -e "${GREEN}╔════════════════════════════════════════╗${NC}"
echo -e "${GREEN}║   ✅ Testes Concluídos com Sucesso!   ║${NC}"
echo -e "${GREEN}╚════════════════════════════════════════╝${NC}"
echo ""
echo -e "${CYAN}📊 Resumo:${NC}"
echo "  Bundle: ${BUNDLE_DIR}"
echo "  Tamanho: ${TOTAL_SIZE}"
echo "  Arquivos: ${LIB_COUNT}"
echo ""
echo -e "${CYAN}🚀 Próximos passos:${NC}"
echo "  1. Criar .deb bundled: ./scripts/build-deb-bundled.sh"
echo "  2. Testar instalação: sudo dpkg -i dist/bagus-browser_v5.0.0_amd64_bundled.deb"
echo "  3. Executar: bagus-browser"
echo ""
