#!/bin/bash
# Bagus Browser - Menu de Build
# Escolha qual versão compilar

set -e

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
CYAN='\033[0;36m'
NC='\033[0m'

clear

echo -e "${CYAN}╔════════════════════════════════════════════════════════╗${NC}"
echo -e "${CYAN}║                                                        ║${NC}"
echo -e "${CYAN}║            🌐 BAGUS BROWSER - BUILD MENU 🌐           ║${NC}"
echo -e "${CYAN}║                                                        ║${NC}"
echo -e "${CYAN}╚════════════════════════════════════════════════════════╝${NC}"
echo ""
echo -e "${BLUE}Escolha qual versão deseja compilar:${NC}"
echo ""
echo -e "${GREEN}1)${NC} Versão Webview ${YELLOW}(Recomendado para começar)${NC}"
echo -e "   ${BLUE}├─${NC} Tamanho: ~5MB"
echo -e "   ${BLUE}├─${NC} Compilação: ~1 minuto"
echo -e "   ${BLUE}├─${NC} Compatibilidade: ~70% dos sites"
echo -e "   ${BLUE}├─${NC} ✅ DuckDuckGo, Wikipedia, Stack Overflow"
echo -e "   ${BLUE}└─${NC} ❌ Google, Facebook, Twitter"
echo ""
echo -e "${GREEN}2)${NC} Versão CEF ${YELLOW}(100% Compatibilidade)${NC}"
echo -e "   ${BLUE}├─${NC} Tamanho: ~165MB"
echo -e "   ${BLUE}├─${NC} Compilação: ~5 minutos (+ 30min instalação CEF)"
echo -e "   ${BLUE}├─${NC} Compatibilidade: 100% dos sites"
echo -e "   ${BLUE}├─${NC} ✅ Google, Facebook, Twitter, GitHub"
echo -e "   ${BLUE}└─${NC} ✅ DevTools (F12)"
echo ""
echo -e "${GREEN}3)${NC} Verificar Privacidade ${YELLOW}(Zero Telemetria)${NC}"
echo -e "   ${BLUE}└─${NC} Auditar código para garantir zero telemetria"
echo ""
echo -e "${GREEN}4)${NC} Sair"
echo ""
echo -ne "${CYAN}Escolha uma opção [1-4]: ${NC}"
read -r choice

case $choice in
    1)
        echo ""
        echo -e "${BLUE}╔════════════════════════════════════════╗${NC}"
        echo -e "${BLUE}║   Compilando Versão Webview           ║${NC}"
        echo -e "${BLUE}╚════════════════════════════════════════╝${NC}"
        echo ""
        ./scripts/build.sh
        echo ""
        echo -e "${GREEN}✅ Pronto! Execute: ./build/bagus${NC}"
        ;;
    2)
        echo ""
        # Verifica se CEF está instalado
        if [ ! -d "/opt/cef" ]; then
            echo -e "${YELLOW}CEF não está instalado.${NC}"
            echo ""
            echo -ne "${CYAN}Deseja instalar CEF agora? (s/N): ${NC}"
            read -r install_cef
            
            if [[ $install_cef =~ ^[Ss]$ ]]; then
                echo ""
                echo -e "${BLUE}╔════════════════════════════════════════╗${NC}"
                echo -e "${BLUE}║   Instalando CEF (15-30 minutos)      ║${NC}"
                echo -e "${BLUE}╚════════════════════════════════════════╝${NC}"
                echo ""
                ./scripts/install_cef.sh
            else
                echo ""
                echo -e "${YELLOW}Instalação cancelada.${NC}"
                echo -e "${BLUE}Para instalar depois: ./scripts/install_cef.sh${NC}"
                exit 0
            fi
        fi
        
        echo ""
        echo -e "${BLUE}╔════════════════════════════════════════╗${NC}"
        echo -e "${BLUE}║   Compilando Versão CEF               ║${NC}"
        echo -e "${BLUE}╚════════════════════════════════════════╝${NC}"
        echo ""
        ./scripts/build_cef.sh
        echo ""
        echo -e "${GREEN}✅ Pronto! Execute: cd build && ./run_bagus_cef.sh${NC}"
        ;;
    3)
        echo ""
        echo -e "${BLUE}╔════════════════════════════════════════╗${NC}"
        echo -e "${BLUE}║   Verificando Privacidade             ║${NC}"
        echo -e "${BLUE}╚════════════════════════════════════════╝${NC}"
        echo ""
        ./scripts/verify_privacy.sh
        ;;
    4)
        echo ""
        echo -e "${BLUE}Até logo! 👋${NC}"
        exit 0
        ;;
    *)
        echo ""
        echo -e "${RED}Opção inválida!${NC}"
        exit 1
        ;;
esac

echo ""
echo -e "${CYAN}╔════════════════════════════════════════════════════════╗${NC}"
echo -e "${CYAN}║                   Processo Concluído                   ║${NC}"
echo -e "${CYAN}╚════════════════════════════════════════════════════════╝${NC}"
echo ""
