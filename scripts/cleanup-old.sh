#!/bin/bash
# Script de limpeza do código antigo para v3.0.0

set -e

echo "🧹 Iniciando limpeza do código antigo..."

# Remover diretório CEF completo
echo "📁 Removendo diretório cef/..."
rm -rf cef/

# Remover internal/cef
echo "📁 Removendo internal/cef/..."
rm -rf internal/cef/

# Remover internal/ui antigo
echo "📁 Removendo internal/ui/..."
rm -rf internal/ui/

# Remover main_cef.go
echo "📄 Removendo main_cef.go..."
rm -f main_cef.go

# Remover scripts CEF
echo "📄 Removendo scripts CEF..."
rm -f scripts/build/build_cef.sh
rm -f scripts/setup/install_cef.sh

# Remover documentação CEF
echo "📁 Removendo docs/cef/..."
rm -rf docs/cef/

# Remover documentação antiga
echo "📄 Removendo documentação antiga..."
rm -f docs/KNOWN_LIMITATIONS.md
rm -f docs/KEYBOARD_SHORTCUTS.md
rm -f QUICK_USAGE.md

# Remover docs de instalação CEF
echo "📄 Removendo docs de instalação CEF..."
rm -f docs/install/INSTALL_CEF.md
rm -f docs/getting-started/QUICKSTART_CEF.md

# Remover builds antigos
echo "📁 Limpando diretório build/..."
rm -rf build/*
rm -rf dist/*

echo "✅ Limpeza concluída!"
echo ""
echo "📊 Arquivos removidos:"
echo "  - cef/ (código C++)"
echo "  - internal/cef/ (bindings Go)"
echo "  - internal/ui/ (UI antiga)"
echo "  - main_cef.go"
echo "  - Scripts e docs CEF"
echo "  - Documentação antiga"
echo ""
echo "🚀 Pronto para reconstrução!"
