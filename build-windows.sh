#!/bin/bash
# Script para criar executável Windows do Bagus Browser

set -e

VERSION="1.0.0"
PKG_NAME="bagus-browser"
BUILD_DIR="build/windows"

echo "🔨 Criando executável Windows do Bagus Browser..."

# Instala dependências para cross-compile (se necessário)
if ! command -v x86_64-w64-mingw32-gcc &> /dev/null; then
    echo "⚠️  Instalando mingw-w64 para cross-compile..."
    sudo apt-get update -qq
    sudo apt-get install -y -qq mingw-w64
fi

# Limpa build anterior
rm -rf "$BUILD_DIR"
mkdir -p "$BUILD_DIR"

echo "► Compilando para Windows..."

# Nota: WebKit2GTK não funciona no Windows
# Para Windows, precisaríamos usar webview2 ou CEF
# Por enquanto, criamos um executável básico que informa isso

cat > main_windows.go << 'EOF'
//go:build windows
// +build windows

package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println("╔════════════════════════════════════════╗")
    fmt.Println("║     BAGUS BROWSER - Windows            ║")
    fmt.Println("╚════════════════════════════════════════╝")
    fmt.Println("")
    fmt.Println("⚠️  Versão Windows em desenvolvimento")
    fmt.Println("")
    fmt.Println("A versão atual do Bagus Browser usa WebKit2GTK,")
    fmt.Println("que é específico para Linux.")
    fmt.Println("")
    fmt.Println("Para Windows, estamos trabalhando em:")
    fmt.Println("  • Integração com Microsoft Edge WebView2")
    fmt.Println("  • Ou CEF (Chromium Embedded Framework)")
    fmt.Println("")
    fmt.Println("Por favor, use a versão Linux por enquanto.")
    fmt.Println("")
    fmt.Println("Pressione ENTER para sair...")
    fmt.Scanln()
    os.Exit(0)
}
EOF

# Compila para Windows
GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-s -w -H windowsgui" -o "$BUILD_DIR/${PKG_NAME}.exe" main_windows.go

# Remove arquivo temporário
rm main_windows.go

# Cria README para Windows
cat > "$BUILD_DIR/README.txt" << 'EOF'
BAGUS BROWSER - Windows
=======================

AVISO: Versão em Desenvolvimento
---------------------------------

Esta versão do Bagus Browser para Windows está em desenvolvimento.

A versão atual usa WebKit2GTK, que é específico para Linux.

Para Windows, estamos trabalhando em:
  * Integração com Microsoft Edge WebView2
  * Ou CEF (Chromium Embedded Framework)

VERSÃO LINUX DISPONÍVEL
------------------------

A versão completa e funcional está disponível para Linux:
  - Download: https://github.com/peder1981/bagus-browser-go
  - Instalação: 2 minutos
  - Tamanho: 4MB
  - 100% funcional

CARACTERÍSTICAS (Versão Linux):
  ✓ Zero telemetria
  ✓ Bloqueador de ads integrado
  ✓ Leve e rápido
  ✓ Código aberto

Para mais informações:
  https://github.com/peder1981/bagus-browser-go

EOF

# Cria instalador NSIS (se disponível)
if command -v makensis &> /dev/null; then
    echo "► Criando instalador NSIS..."
    
    cat > "$BUILD_DIR/installer.nsi" << EOF
!define PRODUCT_NAME "Bagus Browser"
!define PRODUCT_VERSION "$VERSION"
!define PRODUCT_PUBLISHER "Bagus Browser Team"
!define PRODUCT_WEB_SITE "https://github.com/peder1981/bagus-browser-go"

Name "\${PRODUCT_NAME} \${PRODUCT_VERSION}"
OutFile "${PKG_NAME}-${VERSION}-setup.exe"
InstallDir "\$PROGRAMFILES64\\Bagus Browser"

Page directory
Page instfiles

Section "MainSection" SEC01
    SetOutPath "\$INSTDIR"
    File "${PKG_NAME}.exe"
    File "README.txt"
    
    CreateDirectory "\$SMPROGRAMS\\Bagus Browser"
    CreateShortCut "\$SMPROGRAMS\\Bagus Browser\\Bagus Browser.lnk" "\$INSTDIR\\${PKG_NAME}.exe"
    CreateShortCut "\$DESKTOP\\Bagus Browser.lnk" "\$INSTDIR\\${PKG_NAME}.exe"
SectionEnd

Section Uninstall
    Delete "\$INSTDIR\\${PKG_NAME}.exe"
    Delete "\$INSTDIR\\README.txt"
    Delete "\$SMPROGRAMS\\Bagus Browser\\Bagus Browser.lnk"
    Delete "\$DESKTOP\\Bagus Browser.lnk"
    RMDir "\$SMPROGRAMS\\Bagus Browser"
    RMDir "\$INSTDIR"
SectionEnd
EOF

    cd "$BUILD_DIR"
    makensis installer.nsi
    cd ../..
fi

echo ""
echo "✅ Executável Windows criado!"
echo "📦 Arquivo: $BUILD_DIR/${PKG_NAME}.exe"
echo ""
echo "⚠️  IMPORTANTE:"
echo "   Esta é uma versão placeholder que informa sobre o desenvolvimento."
echo "   A versão funcional completa está disponível apenas para Linux."
echo ""
