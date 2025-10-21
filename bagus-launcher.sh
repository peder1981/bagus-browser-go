#!/bin/bash
# BAGUS BROWSER - LAUNCHER
# Uso: bagus [--fast|--full|--help]

GREEN='\033[0;32m'
BLUE='\033[0;34m'
CYAN='\033[0;36m'
YELLOW='\033[1;33m'
NC='\033[0m'

# Ajuda
if [ "$1" = "--help" ] || [ "$1" = "-h" ]; then
    echo "Bagus Browser - Launcher"
    echo ""
    echo "Uso:"
    echo "  bagus           - Menu interativo"
    echo "  bagus --fast    - Versão Rápida (sem menu)"
    echo "  bagus --full    - Versão Completa (sem menu)"
    echo "  bagus --help    - Esta ajuda"
    echo ""
    echo "Exemplos:"
    echo "  bagus                    # Abre menu"
    echo "  bagus --fast             # Abre versão rápida direto"
    echo "  bagus --full             # Abre versão completa direto"
    echo "  bagus --fast google.com  # Abre URL na versão rápida"
    exit 0
fi

# Detecta versões instaladas
RAPIDA=$([ -f "build/bagus" ] && echo "true" || echo "false")
COMPLETA=$([ -f "build/run_bagus_cef.sh" ] && echo "true" || echo "false")

# Nenhuma instalada
if [ "$RAPIDA" = "false" ] && [ "$COMPLETA" = "false" ]; then
    echo -e "${YELLOW}⚠ Nenhuma versão instalada!${NC}"
    echo -e "${BLUE}Execute:${NC} ./install.sh"
    exit 1
fi

# Argumentos de linha de comando
if [ "$1" = "--fast" ]; then
    if [ "$RAPIDA" = "false" ]; then
        echo -e "${YELLOW}⚠ Versão Rápida não instalada!${NC}"
        exit 1
    fi
    shift
    exec ./build/bagus "$@"
fi

if [ "$1" = "--full" ]; then
    if [ "$COMPLETA" = "false" ]; then
        echo -e "${YELLOW}⚠ Versão Completa não instalada!${NC}"
        exit 1
    fi
    shift
    cd build && exec ./run_bagus_cef.sh "$@"
fi

# Só uma instalada - executa direto
if [ "$RAPIDA" = "true" ] && [ "$COMPLETA" = "false" ]; then
    exec ./build/bagus "$@"
fi

if [ "$RAPIDA" = "false" ] && [ "$COMPLETA" = "true" ]; then
    cd build && exec ./run_bagus_cef.sh "$@"
fi

# Ambas instaladas - menu de escolha
clear
echo -e "${CYAN}╔═══════════════════════════════════════╗${NC}"
echo -e "${CYAN}║     🌐 BAGUS BROWSER - LAUNCHER      ║${NC}"
echo -e "${CYAN}╚═══════════════════════════════════════╝${NC}"
echo ""
echo -e "${BLUE}Qual versão você quer usar?${NC}"
echo ""
echo -e "${GREEN}1)${NC} Versão Rápida ⚡"
echo -e "   • Abre instantaneamente"
echo -e "   • DuckDuckGo, Wikipedia, YouTube"
echo ""
echo -e "${GREEN}2)${NC} Versão Completa 🌐"
echo -e "   • Google, Facebook, Twitter"
echo -e "   • 100% dos sites"
echo ""
echo -ne "${CYAN}Escolha [1-2]:${NC} "
read -r choice

case $choice in
    1)
        echo -e "${GREEN}✓ Abrindo Versão Rápida...${NC}"
        exec ./build/bagus "$@"
        ;;
    2)
        echo -e "${GREEN}✓ Abrindo Versão Completa...${NC}"
        cd build && exec ./run_bagus_cef.sh "$@"
        ;;
    *)
        echo -e "${YELLOW}Opção inválida. Abrindo Versão Rápida...${NC}"
        exec ./build/bagus "$@"
        ;;
esac
