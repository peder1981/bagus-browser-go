#!/bin/bash
# ╔════════════════════════════════════════════════════════╗
# ║                                                        ║
# ║        BAGUS BROWSER - INSTALADOR ÚNICO               ║
# ║        Simples, Rápido e Sem Complicação              ║
# ║                                                        ║
# ╚════════════════════════════════════════════════════════╝
#
# Este é o ÚNICO script que você precisa executar!
# Ele detecta seu sistema, instala dependências e compila o navegador.

set -e

GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
CYAN='\033[0;36m'
RED='\033[0;31m'
NC='\033[0m'

# Função para instalar versão rápida
install_fast() {
    echo -e "${YELLOW}► Detectando sistema...${NC}"
    if [ -f /etc/os-release ]; then
        . /etc/os-release
        OS=$ID
    else
        OS=$(uname -s)
    fi
    
    echo -e "${YELLOW}► Instalando dependências...${NC}"
    case $OS in
        ubuntu|debian|pop|linuxmint)
            sudo apt-get update -qq
            sudo apt-get install -y -qq golang git build-essential libwebkit2gtk-4.0-dev libgtk-3-dev > /dev/null 2>&1
            ;;
        fedora)
            sudo dnf install -y -q golang git gcc-c++ webkit2gtk3-devel gtk3-devel > /dev/null 2>&1
            ;;
        arch|manjaro)
            sudo pacman -S --noconfirm --quiet go git base-devel webkit2gtk gtk3 > /dev/null 2>&1
            ;;
    esac
    
    echo -e "${YELLOW}► Baixando dependências Go...${NC}"
    go mod download > /dev/null 2>&1
    go mod tidy > /dev/null 2>&1
    
    echo -e "${YELLOW}► Compilando Bagus Browser...${NC}"
    mkdir -p build
    go build -ldflags "-s -w" -o build/bagus . > /dev/null 2>&1
    
    if [ -f "build/bagus" ]; then
        SIZE=$(du -h build/bagus | cut -f1)
        echo -e "${GREEN}✓ Compilado com sucesso! (${SIZE})${NC}"
    else
        echo -e "${RED}✗ Erro na compilação${NC}"
        exit 1
    fi
    
    mkdir -p ~/.bagus
    chmod 700 ~/.bagus
}

clear
echo -e "${CYAN}╔════════════════════════════════════════════════════════╗${NC}"
echo -e "${CYAN}║                                                        ║${NC}"
echo -e "${CYAN}║           🌐 BAGUS BROWSER - INSTALAÇÃO 🌐            ║${NC}"
echo -e "${CYAN}║                                                        ║${NC}"
echo -e "${CYAN}╚════════════════════════════════════════════════════════╝${NC}"
echo ""
echo -e "${BLUE}Instalação do Bagus Browser${NC}"
echo ""
echo -e "${GREEN}✓ Versão Rápida e Funcional${NC}"
echo "   • Instalação: 2 minutos"
echo "   • Tamanho: ~4MB"
echo "   • Sites: DuckDuckGo, Wikipedia, YouTube, Reddit, Stack Overflow"
echo "   • Status: 100% funcional e testado"
echo ""
echo -e "${YELLOW}Pressione ENTER para instalar...${NC}"
read -r choice

echo ""
echo -e "${BLUE}╔════════════════════════════════════════╗${NC}"
echo -e "${BLUE}║   Instalando Bagus Browser            ║${NC}"
echo -e "${BLUE}╚════════════════════════════════════════╝${NC}"
echo ""
install_fast

# Criar comando bagus
mkdir -p ~/.local/bin
INSTALL_DIR="$(pwd)"
cat > ~/.local/bin/bagus << EOF
#!/bin/bash
cd "$INSTALL_DIR"
./build/bagus "\$@"
EOF
chmod +x ~/.local/bin/bagus

echo ""
echo -e "${GREEN}╔════════════════════════════════════════╗${NC}"
echo -e "${GREEN}║   ✓ Instalação Concluída!             ║${NC}"
echo -e "${GREEN}╚════════════════════════════════════════╝${NC}"
echo ""
echo -e "${CYAN}Para usar, digite:${NC} ${GREEN}bagus${NC}"
echo ""

exit 0
