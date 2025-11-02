#!/bin/bash

# Script de build multiplataforma para Bagus Browser v5.0.0
# Suporta: Linux, Windows, macOS, FreeBSD

set -e

# Cores
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
CYAN='\033[0;36m'
NC='\033[0m'

# Configuração
PROJECT_NAME="bagus-browser"
VERSION="5.0.0"
BUILD_DIR="build"
DIST_DIR="dist"

# Detectar SO atual
detect_os() {
    case "$(uname -s)" in
        Linux*)     echo "linux";;
        Darwin*)    echo "macos";;
        MINGW*|MSYS*) echo "windows";;
        FreeBSD*)   echo "freebsd";;
        *)          echo "unknown";;
    esac
}

# Detectar arquitetura
detect_arch() {
    case "$(uname -m)" in
        x86_64)     echo "amd64";;
        arm64|aarch64) echo "arm64";;
        armv7l)     echo "arm";;
        i386)       echo "386";;
        *)          echo "amd64";;
    esac
}

# Exibir uso
usage() {
    cat << EOF
Uso: $0 [OPÇÕES]

Opções:
  --os OS              SO alvo: linux, windows, macos, freebsd, all (padrão: SO atual)
  --arch ARCH          Arquitetura: amd64, arm64, arm, 386 (padrão: arquitetura atual)
  --output DIR         Diretório de saída (padrão: dist)
  --help               Exibir esta mensagem

Exemplos:
  $0 --os linux --arch amd64
  $0 --os windows --arch amd64
  $0 --os all
  $0 --os macos --arch arm64

SOs suportados:
  - linux    (x86_64, arm64, arm, i386)
  - windows  (x86_64, i386)
  - macos    (x86_64, arm64)
  - freebsd  (x86_64, arm64)

EOF
    exit 1
}

# Compilar para um SO/arquitetura
build_for_platform() {
    local os=$1
    local arch=$2
    local output_dir=$3
    
    echo -e "${BLUE}🔨 Compilando para $os/$arch...${NC}"
    
    # Definir variáveis de ambiente
    export GOOS=$os
    export GOARCH=$arch
    
    # Definir nome do binário
    local binary_name="$PROJECT_NAME"
    if [ "$os" = "windows" ]; then
        binary_name="${binary_name}.exe"
    fi
    
    # Definir caminho de saída
    local output_path="$output_dir/${binary_name}"
    if [ "$os" != "linux" ] && [ "$os" != "freebsd" ]; then
        output_path="$output_dir/${PROJECT_NAME}-${os}-${arch}/${binary_name}"
        mkdir -p "$(dirname "$output_path")"
    fi
    
    # Compilar
    if go build -ldflags="-s -w" -o "$output_path" ./cmd/bagus-browser; then
        local size=$(du -h "$output_path" | cut -f1)
        echo -e "${GREEN}✅ $os/$arch compilado com sucesso (${size})${NC}"
        echo "$output_path"
    else
        echo -e "${RED}❌ Erro ao compilar para $os/$arch${NC}"
        return 1
    fi
}

# Criar pacote para Linux
package_linux() {
    local arch=$1
    local output_dir=$2
    
    echo -e "${CYAN}📦 Criando pacote Linux ($arch)...${NC}"
    
    # Tarball
    local tarball="$output_dir/${PROJECT_NAME}_v${VERSION}_linux_${arch}.tar.gz"
    tar -czf "$tarball" -C "$output_dir" "$PROJECT_NAME"
    echo -e "${GREEN}✅ Tarball criado: $tarball${NC}"
}

# Criar pacote para Windows
package_windows() {
    local arch=$1
    local output_dir=$2
    
    echo -e "${CYAN}📦 Criando pacote Windows ($arch)...${NC}"
    
    # ZIP
    local zip_file="$output_dir/${PROJECT_NAME}_v${VERSION}_windows_${arch}.zip"
    cd "$output_dir/${PROJECT_NAME}-windows-${arch}"
    zip -r "../../$zip_file" .
    cd - > /dev/null
    
    echo -e "${GREEN}✅ ZIP criado: $zip_file${NC}"
}

# Criar pacote para macOS
package_macos() {
    local arch=$1
    local output_dir=$2
    
    echo -e "${CYAN}📦 Criando pacote macOS ($arch)...${NC}"
    
    # DMG (simplificado - apenas ZIP por enquanto)
    local zip_file="$output_dir/${PROJECT_NAME}_v${VERSION}_macos_${arch}.zip"
    cd "$output_dir/${PROJECT_NAME}-macos-${arch}"
    zip -r "../../$zip_file" .
    cd - > /dev/null
    
    echo -e "${GREEN}✅ ZIP criado: $zip_file${NC}"
}

# Criar pacote para FreeBSD
package_freebsd() {
    local arch=$1
    local output_dir=$2
    
    echo -e "${CYAN}📦 Criando pacote FreeBSD ($arch)...${NC}"
    
    # Tarball
    local tarball="$output_dir/${PROJECT_NAME}_v${VERSION}_freebsd_${arch}.tar.gz"
    tar -czf "$tarball" -C "$output_dir" "$PROJECT_NAME"
    echo -e "${GREEN}✅ Tarball criado: $tarball${NC}"
}

# Main
main() {
    local target_os=""
    local target_arch=""
    local output_dir="$DIST_DIR"
    
    # Parse argumentos
    while [[ $# -gt 0 ]]; do
        case $1 in
            --os)
                target_os="$2"
                shift 2
                ;;
            --arch)
                target_arch="$2"
                shift 2
                ;;
            --output)
                output_dir="$2"
                shift 2
                ;;
            --help)
                usage
                ;;
            *)
                echo "Opção desconhecida: $1"
                usage
                ;;
        esac
    done
    
    # Valores padrão
    if [ -z "$target_os" ]; then
        target_os=$(detect_os)
    fi
    
    if [ -z "$target_arch" ]; then
        target_arch=$(detect_arch)
    fi
    
    # Criar diretório de saída
    mkdir -p "$output_dir"
    
    echo -e "${BLUE}═══════════════════════════════════════════════════════════${NC}"
    echo -e "${BLUE}🌍 Build Multiplataforma - Bagus Browser v${VERSION}${NC}"
    echo -e "${BLUE}═══════════════════════════════════════════════════════════${NC}"
    echo ""
    
    # Compilar
    if [ "$target_os" = "all" ]; then
        # Compilar para todos os SOs
        build_for_platform "linux" "amd64" "$output_dir"
        build_for_platform "linux" "arm64" "$output_dir"
        build_for_platform "windows" "amd64" "$output_dir"
        build_for_platform "windows" "386" "$output_dir"
        build_for_platform "darwin" "amd64" "$output_dir"
        build_for_platform "darwin" "arm64" "$output_dir"
        build_for_platform "freebsd" "amd64" "$output_dir"
    else
        # Compilar para SO específico
        build_for_platform "$target_os" "$target_arch" "$output_dir"
    fi
    
    echo ""
    echo -e "${GREEN}✅ Build concluído!${NC}"
    echo ""
    echo -e "${CYAN}📁 Binários em: $output_dir${NC}"
    ls -lh "$output_dir" | grep -E "bagus|\.exe" || true
}

main "$@"
