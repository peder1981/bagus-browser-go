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

# Verificar se gh CLI está instalado
if ! command -v gh &> /dev/null; then
    echo -e "${RED}❌ GitHub CLI (gh) não encontrado${NC}"
    echo -e "${YELLOW}Instale com: sudo apt install gh${NC}"
    echo -e "${YELLOW}Ou: https://cli.github.com/${NC}"
    exit 1
fi

# Verificar autenticação
if ! gh auth status &> /dev/null; then
    echo -e "${RED}❌ Não autenticado no GitHub${NC}"
    echo -e "${YELLOW}Execute: gh auth login${NC}"
    exit 1
fi

# Obter versão
VERSION=$(git describe --tags --abbrev=0 2>/dev/null || echo "v4.1.0")
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
    echo -e "${YELLOW}⚠️  Tag ${VERSION} não existe${NC}"
    read -p "Criar tag agora? (y/n) " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        echo -e "${YELLOW}📝 Digite a mensagem da release:${NC}"
        read -p "> " RELEASE_MSG
        git tag -a ${VERSION} -m "${RELEASE_MSG}"
        git push origin ${VERSION}
        echo -e "${GREEN}✅ Tag criada e enviada${NC}"
    else
        exit 1
    fi
fi

# Gerar notas de release
echo -e "${YELLOW}📝 Gerando notas de release...${NC}"
RELEASE_NOTES=$(cat <<EOF
# Bagus Browser ${VERSION}

## 🎉 Release

Browser minimalista, seguro e privado construído em Go.

### ✨ Features

- 🌐 WebView completo (WebKit2GTK)
- 📑 Múltiplas abas independentes
- 🔍 Buscar na página (Ctrl+F)
- ⭐ Favoritos com criptografia AES-256 (Ctrl+D)
- 📥 Gerenciador de downloads
- 🔍 Zoom (Ctrl++, Ctrl+-, Ctrl+0)
- ⌨️  15 atalhos de teclado

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
- **Atalhos:** 15
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

# Criar release
echo -e "${YELLOW}🚀 Criando release no GitHub...${NC}"
echo ""

gh release create ${VERSION} \
    --title "Bagus Browser ${VERSION}" \
    --notes "${RELEASE_NOTES}" \
    dist/*.deb \
    dist/*.tar.gz \
    dist/SHA256SUMS

if [ $? -eq 0 ]; then
    echo ""
    echo -e "${GREEN}✅ Release criada com sucesso!${NC}"
    echo ""
    echo -e "${BLUE}🔗 URL: https://github.com/peder1981/bagus-browser-go/releases/tag/${VERSION}${NC}"
    echo ""
    echo -e "${GREEN}📦 Arquivos publicados:${NC}"
    ls -lh dist/
    echo ""
    echo -e "${GREEN}🎉 Release completa!${NC}"
else
    echo -e "${RED}❌ Erro ao criar release${NC}"
    exit 1
fi
