#!/bin/bash
set -e

# Cores
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

echo -e "${BLUE}🚀 Bagus Browser - GitHub Release Script${NC}"
echo -e "${BLUE}=========================================${NC}"
echo ""

# Este script prepara os arquivos para release manual no GitHub
echo -e "${YELLOW}📝 Preparando release para publicação manual...${NC}"

# Obter versão
VERSION=$(git describe --tags --abbrev=0 2>/dev/null || echo "v4.2.0")
echo -e "${GREEN}Versão: ${VERSION}${NC}"
echo ""

# Verificar se dist existe
if [ ! -d "dist" ]; then
    echo -e "${RED}❌ Diretório dist/ não encontrado${NC}"
    echo -e "${YELLOW}Execute primeiro: ./build.sh${NC}"
    exit 1
fi

# Verificar se tag existe
if ! git rev-parse ${VERSION} >/dev/null 2>&1; then
    echo -e "${RED}❌ Tag ${VERSION} não existe${NC}"
    echo -e "${YELLOW}Crie a tag primeiro com:${NC}"
    echo -e "  git tag -a ${VERSION} -m 'Mensagem da release'"
    echo -e "  git push origin ${VERSION}"
    exit 1
fi

echo -e "${GREEN}✅ Tag ${VERSION} encontrada${NC}"

# Gerar notas de release
echo -e "${YELLOW}📝 Gerando notas de release...${NC}"
RELEASE_NOTES=$(cat <<EOF
# Bagus Browser ${VERSION}

## 🎉 Release

Browser minimalista, seguro e privado construído em Go.

### ✨ Features

- 🌐 WebView completo (WebKit2GTK)
- 📏 Múltiplas abas independentes
- 📝 Menu superior completo (Arquivo, Navegação, Favoritos, Ferramentas)
- 🔍 Buscar na página (Ctrl+F)
- ⭐ Favoritos com criptografia AES-256 (Ctrl+D, Ctrl+Shift+B)
- 📅 Gerenciador de downloads
- 🔍 Zoom (Ctrl++, Ctrl+-, Ctrl+0)
- 🎯 Foco automático na barra de URL ao abrir nova aba
- ⌨️  16 atalhos de teclado

### 🔒 Segurança

- AES-256-GCM para favoritos
- PBKDF2 (100,000 iterações)
- Validação rigorosa de URLs
- Sanitização de input
- WebView hardened

### 🕵️ Privacidade

- **Zero telemetria** (garantido)
- Third-party cookies bloqueados
- WebGL/WebAudio bloqueados (anti-fingerprinting)
- Dados criptografados localmente

### 📊 Estatísticas

- **Tamanho:** 6.4MB
- **Atalhos:** 16
- **Plataforma:** Linux only
- **Licença:** MIT

### 📦 Instalação

#### Ubuntu/Debian (.deb)
\`\`\`bash
wget https://github.com/peder1981/bagus-browser-go/releases/download/${VERSION}/bagus-browser_${VERSION#v}_amd64.deb
sudo dpkg -i bagus-browser_${VERSION#v}_amd64.deb
sudo apt-get install -f  # Instalar dependências
\`\`\`

#### Tarball
\`\`\`bash
wget https://github.com/peder1981/bagus-browser-go/releases/download/${VERSION}/bagus-browser_${VERSION}_linux_amd64.tar.gz
tar -xzf bagus-browser_${VERSION}_linux_amd64.tar.gz
./bagus-browser
\`\`\`

### 🔐 Verificação

\`\`\`bash
wget https://github.com/peder1981/bagus-browser-go/releases/download/${VERSION}/SHA256SUMS
sha256sum -c SHA256SUMS
\`\`\`

### 📚 Documentação

- [README](https://github.com/peder1981/bagus-browser-go/blob/main/README.md)
- [CHANGELOG](https://github.com/peder1981/bagus-browser-go/blob/main/CHANGELOG.md)
- [Segurança](https://github.com/peder1981/bagus-browser-go/blob/main/docs/SECURITY.md)
- [Privacidade](https://github.com/peder1981/bagus-browser-go/blob/main/docs/PRIVACY.md)

---

**Bagus Browser - Navegue com privacidade e segurança** 🌐🔒
EOF
)

# Salvar notas em arquivo
NOTES_FILE="docs/releases/RELEASE_NOTES_${VERSION}.md"
echo "${RELEASE_NOTES}" > ${NOTES_FILE}
echo -e "${GREEN}✅ Notas salvas em: ${NOTES_FILE}${NC}"
echo ""

# Listar arquivos para upload
echo -e "${BLUE}📦 Arquivos prontos para upload:${NC}"
ls -lh dist/
echo ""

# Instruções para publicação manual
echo -e "${BLUE}=========================================${NC}"
echo -e "${GREEN}✅ ARQUIVOS PREPARADOS PARA RELEASE!${NC}"
echo -e "${BLUE}=========================================${NC}"
echo ""
echo -e "${YELLOW}📋 PRÓXIMOS PASSOS (Manual):${NC}"
echo ""
echo -e "${BLUE}1. Acesse:${NC}"
echo -e "   ${GREEN}https://github.com/peder1981/bagus-browser-go/releases/new${NC}"
echo ""
echo -e "${BLUE}2. Preencha:${NC}"
echo -e "   • Tag: ${GREEN}${VERSION}${NC} (já existe)"
echo -e "   • Título: ${GREEN}Bagus Browser ${VERSION} - Menu Superior + Melhorias de UX${NC}"
echo -e "   • Descrição: ${GREEN}Copie o conteúdo de ${NOTES_FILE}${NC}"
echo ""
echo -e "${BLUE}3. Upload de Arquivos:${NC}"
echo -e "   Arraste estes arquivos para a área de upload:"
for file in dist/*.deb dist/*.tar.gz dist/SHA256SUMS; do
    if [ -f "$file" ]; then
        filename=$(basename "$file")
        size=$(du -h "$file" | cut -f1)
        echo -e "   • ${GREEN}${filename}${NC} (${size})"
    fi
done
echo ""
echo -e "${BLUE}4. Publicar:${NC}"
echo -e "   Clique em ${GREEN}'Publish release'${NC}"
echo ""
echo -e "${BLUE}=========================================${NC}"
echo -e "${YELLOW}💡 Dica: Abra ${NOTES_FILE} em um editor para copiar as notas${NC}"
echo -e "${BLUE}=========================================${NC}"
echo ""
