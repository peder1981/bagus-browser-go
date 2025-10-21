#!/bin/bash
# Script para compilar Bagus Browser com CEF

set -e

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

echo -e "${BLUE}╔════════════════════════════════════════╗${NC}"
echo -e "${BLUE}║   Bagus Browser CEF - Build           ║${NC}"
echo -e "${BLUE}╚════════════════════════════════════════╝${NC}"
echo ""

# Verifica se CEF está instalado
if [ ! -d "/opt/cef" ]; then
    echo -e "${RED}✗ CEF não encontrado em /opt/cef${NC}"
    echo -e "${YELLOW}Execute primeiro: ./scripts/install_cef.sh${NC}"
    exit 1
fi

# Verifica se libcef_dll_wrapper foi compilado
if [ ! -f "/opt/cef/build/libcef_dll_wrapper/libcef_dll_wrapper.a" ]; then
    echo -e "${RED}✗ libcef_dll_wrapper não encontrado${NC}"
    echo -e "${YELLOW}Execute: cd /opt/cef/build && make libcef_dll_wrapper${NC}"
    exit 1
fi

echo -e "${GREEN}✓ CEF encontrado${NC}"
echo ""

# Cria diretório de build
mkdir -p build

# Compila código C++
echo -e "${YELLOW}► Compilando código C++...${NC}"
cd cef

# Compila cef_app.cpp
g++ -c \
    -std=c++17 \
    -I/opt/cef \
    -I/opt/cef/include \
    -DUSE_SANDBOX=0 \
    cef_app.cpp \
    -o ../build/cef_app.o

# Compila cef_browser.cpp
g++ -c \
    -std=c++17 \
    -I/opt/cef \
    -I/opt/cef/include \
    -DUSE_SANDBOX=0 \
    cef_browser.cpp \
    -o ../build/cef_browser.o

# Compila cef_wrapper.cc
g++ -c \
    -std=c++17 \
    -I/opt/cef \
    -I/opt/cef/include \
    -DUSE_SANDBOX=0 \
    cef_wrapper.cc \
    -o ../build/cef_wrapper.o

# Cria biblioteca estática
ar rcs ../build/libcef_wrapper.a ../build/cef_app.o ../build/cef_browser.o ../build/cef_wrapper.o

echo -e "${GREEN}✓ C++ compilado${NC}"
echo ""

# Compila Go com CGO
echo -e "${YELLOW}► Compilando Go com CEF...${NC}"
cd ..

CGO_ENABLED=1 \
CC="gcc" \
CXX="g++" \
CGO_CXXFLAGS="-std=c++17 -I${PWD}/cef -I/opt/cef -I/opt/cef/include" \
CGO_LDFLAGS="-L/opt/cef/build/libcef_dll_wrapper -L/opt/cef/Release -L/opt/cef/build/tests/cefsimple/Release -lcef_dll_wrapper -lcef -lstdc++ -lX11 -lpthread -ldl" \
go build \
    -o build/bagus-cef \
    -tags cef \
    main_cef.go

echo -e "${GREEN}✓ Go compilado${NC}"
echo ""

# Copia recursos do CEF
echo -e "${YELLOW}► Copiando recursos do CEF...${NC}"
cp -r /opt/cef/Release/* build/ 2>/dev/null || true
cp -r /opt/cef/Resources/* build/ 2>/dev/null || true

echo -e "${GREEN}✓ Recursos copiados${NC}"
echo ""

# Cria script de execução
cat > build/run_bagus_cef.sh << 'EOF'
#!/bin/bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
export LD_LIBRARY_PATH="$DIR:$LD_LIBRARY_PATH"
"$DIR/bagus-cef" "$@"
EOF

chmod +x build/run_bagus_cef.sh

echo -e "${GREEN}╔════════════════════════════════════════╗${NC}"
echo -e "${GREEN}║   Build Concluído com Sucesso!        ║${NC}"
echo -e "${GREEN}╚════════════════════════════════════════╝${NC}"
echo ""
echo -e "${BLUE}Executável:${NC} build/bagus-cef"
echo -e "${BLUE}Script:${NC} build/run_bagus_cef.sh"
echo ""
echo -e "${YELLOW}Para executar:${NC}"
echo -e "  cd build && ./run_bagus_cef.sh"
echo ""
