# 📦 Instruções para Publicar Release

## Arquivos Prontos para Publicação

Os seguintes arquivos foram gerados e estão prontos para upload:

### Binários
- `build/bagus-linux-amd64` - Executável Linux 64-bit
- `build/bagus-browser_1.0.0_amd64.deb` - Pacote Debian/Ubuntu

### Pacotes Compactados
- `dist/bagus-2.0.0-linux-amd64.tar.gz` - Linux 64-bit (tar.gz)

## Como Publicar no GitHub

### Opção 1: Via Interface Web (Recomendado)

1. Acesse: https://github.com/peder1981/bagus-browser-go/releases/new

2. Preencha os campos:
   - **Tag version:** `v2.0.0`
   - **Release title:** `Bagus Browser 2.0.0`
   - **Description:** (copie o texto abaixo)

```markdown
## Bagus Browser 2.0.0

### 🚀 Novidades
- ✅ Navegação direta com webview nativo (sem iframe)
- ✅ Proteção contra múltiplas instâncias
- ✅ Suporte completo a todos os sites
- ✅ Interface moderna e responsiva
- ✅ Histórico de navegação com busca
- ✅ Bloqueador de anúncios integrado

### 📦 Downloads

**Linux:**
- `bagus-2.0.0-linux-amd64.tar.gz` - Para sistemas Linux 64-bit

**Debian/Ubuntu:**
- `bagus-browser_1.0.0_amd64.deb` - Pacote .deb para instalação fácil

### 🔧 Instalação

**Linux (método rápido):**
```bash
tar -xzf bagus-2.0.0-linux-amd64.tar.gz
sudo mv bagus-linux-amd64 /usr/local/bin/bagus
sudo chmod +x /usr/local/bin/bagus
bagus
```

**Debian/Ubuntu:**
```bash
sudo dpkg -i bagus-browser_1.0.0_amd64.deb
bagus-browser
```

### 🐛 Correções
- ✅ Corrigido problema de múltiplas instâncias
- ✅ Corrigido problema de sites não carregarem
- ✅ Melhorada estabilidade geral
- ✅ Desativado modo debug no login

### 📝 Notas
- Primeira versão estável
- Zero telemetria
- 100% open source
- Compatibilidade total com sites modernos
```

3. Faça upload dos arquivos:
   - `dist/bagus-2.0.0-linux-amd64.tar.gz`
   - `build/bagus-browser_1.0.0_amd64.deb`

4. Clique em **Publish release**

### Opção 2: Via GitHub CLI (se disponível)

```bash
# Instalar GitHub CLI (se necessário)
sudo apt install gh

# Autenticar
gh auth login

# Criar release
gh release create v2.0.0 \
  --title "Bagus Browser 2.0.0" \
  --notes-file RELEASE_NOTES.md \
  dist/bagus-2.0.0-linux-amd64.tar.gz \
  build/bagus-browser_1.0.0_amd64.deb
```

### Opção 3: Via Git Tags

```bash
# Criar tag
git tag -a v2.0.0 -m "Release 2.0.0"

# Push da tag
git push origin v2.0.0

# Depois faça upload manual dos arquivos na interface web
```

## Verificação Pós-Release

Após publicar, verifique:

1. ✅ Release aparece em: https://github.com/peder1981/bagus-browser-go/releases
2. ✅ Arquivos estão disponíveis para download
3. ✅ Links no README funcionam
4. ✅ Badges no README estão atualizados

## Próximos Passos

1. Anunciar release nas redes sociais
2. Atualizar documentação se necessário
3. Monitorar issues e feedback dos usuários
4. Planejar próxima versão

---

**Versão:** 2.0.0  
**Data:** $(date +%Y-%m-%d)  
**Status:** ✅ Pronto para Publicação
