#!/bin/bash

# Script de instalação do ícone do Bagus Browser
# Não expõe dados sensíveis - usa variáveis de ambiente padrão

set -e

echo "╔════════════════════════════════════════╗"
echo "║   Bagus Browser - Instalação Ícone   ║"
echo "╚════════════════════════════════════════╝"
echo ""

# Diretórios padrão do Linux (sem dados sensíveis)
ICON_DIR="${XDG_DATA_HOME:-$HOME/.local/share}/icons/hicolor"
DESKTOP_DIR="${XDG_DATA_HOME:-$HOME/.local/share}/applications"
BIN_DIR="${HOME}/.local/bin"

# Cria diretórios se não existirem
mkdir -p "$ICON_DIR/48x48/apps"
mkdir -p "$ICON_DIR/64x64/apps"
mkdir -p "$ICON_DIR/128x128/apps"
mkdir -p "$ICON_DIR/256x256/apps"
mkdir -p "$ICON_DIR/512x512/apps"
mkdir -p "$ICON_DIR/scalable/apps"
mkdir -p "$DESKTOP_DIR"
mkdir -p "$BIN_DIR"

echo "► Copiando ícones..."
cp assets/icons/bagus-48.png "$ICON_DIR/48x48/apps/bagus-browser.png"
cp assets/icons/bagus-64.png "$ICON_DIR/64x64/apps/bagus-browser.png"
cp assets/icons/bagus-128.png "$ICON_DIR/128x128/apps/bagus-browser.png"
cp assets/icons/bagus-256.png "$ICON_DIR/256x256/apps/bagus-browser.png"
cp assets/icons/bagus-512.png "$ICON_DIR/512x512/apps/bagus-browser.png"
cp assets/icons/bagus.svg "$ICON_DIR/scalable/apps/bagus-browser.svg"

echo "► Copiando executável..."
cp build/bagus "$BIN_DIR/bagus-browser"
chmod +x "$BIN_DIR/bagus-browser"

echo "► Instalando arquivo .desktop..."
cp assets/bagus-browser.desktop "$DESKTOP_DIR/bagus-browser.desktop"
chmod +x "$DESKTOP_DIR/bagus-browser.desktop"

echo "► Atualizando cache de ícones..."
if command -v gtk-update-icon-cache &> /dev/null; then
    gtk-update-icon-cache -f -t "$ICON_DIR" 2>/dev/null || true
fi

if command -v update-desktop-database &> /dev/null; then
    update-desktop-database "$DESKTOP_DIR" 2>/dev/null || true
fi

echo ""
echo "✓ Instalação concluída!"
echo ""
echo "Bagus Browser instalado em:"
echo "  • Executável: $BIN_DIR/bagus-browser"
echo "  • Ícones: $ICON_DIR"
echo "  • Desktop: $DESKTOP_DIR"
echo ""
echo "Você pode agora:"
echo "  1. Executar: bagus-browser"
echo "  2. Procurar 'Bagus Browser' no menu de aplicativos"
echo ""
