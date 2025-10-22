#!/bin/bash
# Script para gerenciar versões do Bagus Browser

set -e

# Cores
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

# Função para exibir uso
usage() {
    echo "Uso: $0 [comando] [versão]"
    echo ""
    echo "Comandos:"
    echo "  current              - Mostra versão atual"
    echo "  update <versão>      - Atualiza para nova versão (ex: 4.4.0)"
    echo "  release <versão>     - Cria release completa (build + tag + docs)"
    echo ""
    echo "Exemplos:"
    echo "  $0 current"
    echo "  $0 update 4.4.0"
    echo "  $0 release 4.4.0"
    exit 1
}

# Função para obter versão atual
get_current_version() {
    # Pegar a tag mais recente (ordenada por versão)
    git tag -l 'v*' | sort -V | tail -1 | sed 's/^v//' || echo "4.4.0"
}

# Função para atualizar CHANGELOG
update_changelog() {
    local version=$1
    local date=$(date +%Y-%m-%d)
    
    echo -e "${YELLOW}📝 Atualizando CHANGELOG.md...${NC}"
    
    # Verificar se já existe entrada para esta versão
    if grep -q "## \[$version\]" CHANGELOG.md; then
        echo -e "${GREEN}✅ CHANGELOG já contém versão $version${NC}"
    else
        echo -e "${RED}⚠️  Adicione manualmente as mudanças no CHANGELOG.md${NC}"
        echo -e "${YELLOW}   Formato: ## [$version] - $date${NC}"
    fi
}

# Função para criar release notes
create_release_notes() {
    local version=$1
    local notes_file="docs/releases/RELEASE_NOTES_v${version}.md"
    
    echo -e "${YELLOW}📝 Criando release notes...${NC}"
    
    if [ -f "$notes_file" ]; then
        echo -e "${GREEN}✅ Release notes já existem: $notes_file${NC}"
    else
        echo -e "${YELLOW}⚠️  Criando template de release notes...${NC}"
        mkdir -p docs/releases
        
        cat > "$notes_file" << EOF
# Bagus Browser v${version}

## 🎉 Release

Browser minimalista, seguro e privado construído em Go.

### ✨ Features

- 🌐 WebView completo (WebKit2GTK)
- 📏 Múltiplas abas independentes
- 📝 Menu superior completo
- 🔍 Buscar na página
- ⭐ Favoritos criptografados
- 🔄 Restauração de sessão
- 🖨️ Impressão (PDF e impressoras)
- ⌨️  30 atalhos de teclado

### 🔒 Segurança

- AES-256-GCM para favoritos e sessões
- PBKDF2 (100,000 iterações)
- Validação rigorosa de URLs
- WebView hardened

### 🕵️ Privacidade

- **Zero telemetria** (garantido)
- Third-party cookies bloqueados
- WebGL/WebAudio bloqueados (anti-fingerprinting)
- Dados criptografados localmente

### 📊 Estatísticas

- **Tamanho:** 6.4MB
- **Atalhos:** 30
- **Plataforma:** Linux only
- **Licença:** MIT

### 📦 Instalação

#### Ubuntu/Debian (.deb)
\`\`\`bash
wget https://github.com/peder1981/bagus-browser-go/releases/download/v${version}/bagus-browser_${version}_amd64.deb
sudo dpkg -i bagus-browser_${version}_amd64.deb
sudo apt-get install -f
\`\`\`

#### Tarball
\`\`\`bash
wget https://github.com/peder1981/bagus-browser-go/releases/download/v${version}/bagus-browser_${version}_linux_amd64.tar.gz
tar -xzf bagus-browser_${version}_linux_amd64.tar.gz
sudo mv bagus-browser /usr/local/bin/
\`\`\`

### 🔗 Links

- **Código:** https://github.com/peder1981/bagus-browser-go
- **Issues:** https://github.com/peder1981/bagus-browser-go/issues
- **Documentação:** https://github.com/peder1981/bagus-browser-go/tree/main/docs

---

**Data:** $(date +%Y-%m-%d)  
**Versão:** ${version}
EOF
        
        echo -e "${GREEN}✅ Release notes criadas: $notes_file${NC}"
        echo -e "${YELLOW}⚠️  Edite o arquivo para adicionar detalhes específicos${NC}"
    fi
}

# Função para limpar arquivos temporários
cleanup_temp_files() {
    echo -e "${YELLOW}🧹 Limpando arquivos temporários...${NC}"
    
    # Remover builds antigos (manter apenas dist/)
    if [ -d "build" ]; then
        rm -rf build
        echo -e "${GREEN}✅ Removido: build/${NC}"
    fi
    
    # Limpar arquivos de log temporários
    find . -name "*.log" -type f -delete 2>/dev/null || true
    
    echo -e "${GREEN}✅ Limpeza concluída${NC}"
}

# Função para criar release completa
create_release() {
    local version=$1
    
    echo -e "${BLUE}🚀 Criando release v${version}${NC}"
    echo -e "${BLUE}=========================================${NC}"
    echo ""
    
    # 1. Atualizar CHANGELOG
    update_changelog "$version"
    
    # 2. Criar release notes
    create_release_notes "$version"
    
    # 3. Commit mudanças
    echo -e "${YELLOW}📝 Commitando mudanças...${NC}"
    git add CHANGELOG.md docs/releases/
    git commit -m "📝 Preparar release v${version}" || echo "Nada para commitar"
    
    # 4. Criar tag
    echo -e "${YELLOW}🏷️  Criando tag v${version}...${NC}"
    if git tag -l | grep -q "^v${version}$"; then
        echo -e "${YELLOW}⚠️  Tag v${version} já existe${NC}"
    else
        git tag -a "v${version}" -m "Release v${version}"
        echo -e "${GREEN}✅ Tag v${version} criada${NC}"
    fi
    
    # 5. Build
    echo -e "${YELLOW}🔨 Compilando...${NC}"
    bash scripts/build.sh
    
    # 6. Limpar temporários
    cleanup_temp_files
    
    # 7. Push
    echo -e "${YELLOW}📤 Enviando para GitHub...${NC}"
    git push origin main
    git push origin "v${version}"
    
    echo ""
    echo -e "${GREEN}=========================================${NC}"
    echo -e "${GREEN}✅ Release v${version} criada com sucesso!${NC}"
    echo -e "${GREEN}=========================================${NC}"
    echo ""
    echo -e "${BLUE}📋 Próximos passos:${NC}"
    echo -e "1. Acesse: ${YELLOW}https://github.com/peder1981/bagus-browser-go/releases/new${NC}"
    echo -e "2. Selecione tag: ${YELLOW}v${version}${NC}"
    echo -e "3. Copie conteúdo de: ${YELLOW}docs/releases/RELEASE_NOTES_v${version}.md${NC}"
    echo -e "4. Faça upload dos arquivos de: ${YELLOW}dist/${NC}"
    echo -e "5. Publique a release"
    echo ""
}

# Main
case "$1" in
    current)
        version=$(get_current_version)
        echo -e "${GREEN}Versão atual: v${version}${NC}"
        ;;
    update)
        if [ -z "$2" ]; then
            echo -e "${RED}❌ Erro: Especifique a versão${NC}"
            usage
        fi
        update_changelog "$2"
        ;;
    release)
        if [ -z "$2" ]; then
            echo -e "${RED}❌ Erro: Especifique a versão${NC}"
            usage
        fi
        create_release "$2"
        ;;
    *)
        usage
        ;;
esac
