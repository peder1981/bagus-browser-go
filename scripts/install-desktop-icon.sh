#!/bin/bash
set -e

# Cores
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

echo -e "${BLUE}🎨 Instalador de Ícone - Bagus Browser${NC}"
echo -e "${BLUE}======================================${NC}"
echo ""

# Verificar se está instalado
if ! command -v bagus-browser &> /dev/null; then
    echo -e "${RED}❌ Bagus Browser não está instalado${NC}"
    echo -e "${YELLOW}Instale primeiro: sudo dpkg -i bagus-browser_*.deb${NC}"
    exit 1
fi

echo -e "${GREEN}✅ Bagus Browser encontrado${NC}"
echo ""

# Atualizar caches do sistema
echo -e "${YELLOW}📦 Atualizando caches do sistema...${NC}"

# Cache de ícones
if command -v gtk-update-icon-cache &> /dev/null; then
    echo -e "  ${BLUE}→${NC} Atualizando cache de ícones..."
    sudo gtk-update-icon-cache -f -t /usr/share/icons/hicolor 2>/dev/null || true
fi

# Banco de dados de aplicações
if command -v update-desktop-database &> /dev/null; then
    echo -e "  ${BLUE}→${NC} Atualizando banco de dados de aplicações..."
    sudo update-desktop-database /usr/share/applications 2>/dev/null || true
fi

# Cache MIME
if command -v update-mime-database &> /dev/null; then
    echo -e "  ${BLUE}→${NC} Atualizando cache MIME..."
    sudo update-mime-database /usr/share/mime 2>/dev/null || true
fi

# XDG desktop menu
if command -v xdg-desktop-menu &> /dev/null; then
    echo -e "  ${BLUE}→${NC} Forçando atualização do menu..."
    xdg-desktop-menu forceupdate 2>/dev/null || true
fi

echo ""
echo -e "${YELLOW}🖥️  Criando atalho na área de trabalho...${NC}"

# Detectar área de trabalho do usuário
DESKTOP_DIR="$HOME/Desktop"
if [ ! -d "$DESKTOP_DIR" ]; then
    DESKTOP_DIR="$HOME/Área de Trabalho"
fi
if [ ! -d "$DESKTOP_DIR" ]; then
    DESKTOP_DIR="$HOME/Área de trabalho"
fi

if [ -d "$DESKTOP_DIR" ]; then
    # Copiar .desktop para área de trabalho
    cp /usr/share/applications/bagus-browser.desktop "$DESKTOP_DIR/"
    chmod +x "$DESKTOP_DIR/bagus-browser.desktop"
    
    # Marcar como confiável (GNOME)
    if command -v gio &> /dev/null; then
        gio set "$DESKTOP_DIR/bagus-browser.desktop" metadata::trusted true 2>/dev/null || true
    fi
    
    echo -e "${GREEN}✅ Atalho criado em: $DESKTOP_DIR${NC}"
else
    echo -e "${YELLOW}⚠️  Área de trabalho não encontrada${NC}"
fi

echo ""
echo -e "${YELLOW}🔄 Reiniciando interface...${NC}"

# Detectar e reiniciar DE
if pgrep -x "gnome-shell" > /dev/null; then
    echo -e "  ${BLUE}→${NC} GNOME detectado"
    killall -SIGQUIT gnome-shell 2>/dev/null || true
elif pgrep -x "plasmashell" > /dev/null; then
    echo -e "  ${BLUE}→${NC} KDE Plasma detectado"
    killall plasmashell 2>/dev/null && kstart5 plasmashell 2>/dev/null &
elif pgrep -x "xfce4-panel" > /dev/null; then
    echo -e "  ${BLUE}→${NC} XFCE detectado"
    xfce4-panel -r 2>/dev/null || true
elif pgrep -x "nautilus" > /dev/null; then
    echo -e "  ${BLUE}→${NC} Nautilus detectado"
    killall nautilus 2>/dev/null || true
fi

# Limpar cache do usuário
echo -e "  ${BLUE}→${NC} Limpando cache do usuário..."
rm -f "$HOME/.cache/icon-theme.cache" 2>/dev/null || true
rm -rf "$HOME/.cache/thumbnails" 2>/dev/null || true

echo ""
echo -e "${GREEN}✅ Configuração completa!${NC}"
echo ""
echo -e "${BLUE}📍 Onde encontrar o Bagus Browser:${NC}"
echo -e "  ${GREEN}•${NC} Menu de aplicações → Internet → Bagus Browser"
echo -e "  ${GREEN}•${NC} Área de trabalho (atalho criado)"
echo -e "  ${GREEN}•${NC} Terminal: ${YELLOW}bagus-browser${NC}"
echo ""
echo -e "${YELLOW}💡 Se ainda não aparecer:${NC}"
echo -e "  ${GREEN}1.${NC} Faça logout e login novamente"
echo -e "  ${GREEN}2.${NC} Ou reinicie o computador"
echo ""
