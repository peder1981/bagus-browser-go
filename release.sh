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

# Verificar token do GitHub
if [ -z "$GITHUB_TOKEN" ]; then
    echo -e "${YELLOW}⚠️  GITHUB_TOKEN não definido${NC}"
    echo -e "${YELLOW}Opções:${NC}"
    echo -e "  1. Exportar: export GITHUB_TOKEN=seu_token_aqui"
    echo -e "  2. Criar em: https://github.com/settings/tokens"
    echo -e "  3. Ou usar gh CLI: gh auth login"
    echo ""
    
    # Tentar usar gh CLI como fallback
    if command -v gh &> /dev/null && gh auth status &> /dev/null; then
        echo -e "${GREEN}✅ Usando gh CLI autenticado${NC}"
        USE_GH_CLI=true
    else
        echo -e "${RED}❌ Nenhuma autenticação disponível${NC}"
        exit 1
    fi
else
    echo -e "${GREEN}✅ GITHUB_TOKEN encontrado${NC}"
    USE_GH_CLI=false
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

if [ "$USE_GH_CLI" = true ]; then
    # Usar gh CLI
    gh release create ${VERSION} \
        --title "Bagus Browser ${VERSION}" \
        --notes "${RELEASE_NOTES}" \
        dist/*.deb \
        dist/*.tar.gz \
        dist/SHA256SUMS
    
    RESULT=$?
else
    # Usar API REST com token
    REPO="peder1981/bagus-browser-go"
    API_URL="https://api.github.com/repos/${REPO}/releases"
    
    # Criar release
    RESPONSE=$(curl -s -X POST \
        -H "Authorization: token ${GITHUB_TOKEN}" \
        -H "Accept: application/vnd.github.v3+json" \
        "${API_URL}" \
        -d "{
            \"tag_name\": \"${VERSION}\",
            \"name\": \"Bagus Browser ${VERSION}\",
            \"body\": $(echo "${RELEASE_NOTES}" | jq -Rs .),
            \"draft\": false,
            \"prerelease\": false
        }")
    
    RELEASE_ID=$(echo "${RESPONSE}" | jq -r '.id')
    
    if [ "$RELEASE_ID" = "null" ] || [ -z "$RELEASE_ID" ]; then
        echo -e "${RED}❌ Erro ao criar release${NC}"
        echo "${RESPONSE}" | jq .
        exit 1
    fi
    
    echo -e "${GREEN}✅ Release criada (ID: ${RELEASE_ID})${NC}"
    
    # Upload de arquivos
    UPLOAD_URL="https://uploads.github.com/repos/${REPO}/releases/${RELEASE_ID}/assets"
    
    for file in dist/*.deb dist/*.tar.gz dist/SHA256SUMS; do
        if [ -f "$file" ]; then
            filename=$(basename "$file")
            echo -e "${YELLOW}📤 Uploading ${filename}...${NC}"
            
            curl -s -X POST \
                -H "Authorization: token ${GITHUB_TOKEN}" \
                -H "Content-Type: application/octet-stream" \
                "${UPLOAD_URL}?name=${filename}" \
                --data-binary "@${file}" > /dev/null
            
            echo -e "${GREEN}   ✅ ${filename} uploaded${NC}"
        fi
    done
    
    RESULT=0
fi

if [ $RESULT -eq 0 ]; then
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
