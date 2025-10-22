# 🚀 Guia para Publicar Release v4.2.0 no GitHub

## ✅ Já Concluído

- [x] Código compilado
- [x] Pacotes criados (.deb e .tar.gz)
- [x] Checksums gerados
- [x] Commit realizado
- [x] Tag v4.2.0 criada
- [x] Push para GitHub concluído

## 📦 Arquivos Prontos para Upload

Os seguintes arquivos estão em `dist/`:

1. **bagus-browser_v4.1.0_amd64.deb** (1.3MB)
2. **bagus-browser_v4.1.0_linux_amd64.tar.gz** (1.4MB)
3. **SHA256SUMS** (checksums)

## 🌐 Passos para Criar Release no GitHub

### Opção 1: Via Interface Web (Recomendado)

1. **Acesse:** https://github.com/peder1981/bagus-browser-go/releases/new

2. **Preencha:**
   - **Tag:** v4.2.0 (já existe)
   - **Release title:** Bagus Browser v4.2.0 - Menu Superior + Melhorias de UX
   - **Description:** Copie o conteúdo de `RELEASE_NOTES_v4.2.0.md`

3. **Upload de Arquivos:**
   - Arraste os 3 arquivos de `dist/` para a área de upload
   - Ou clique em "Attach binaries" e selecione os arquivos

4. **Publicar:**
   - Clique em "Publish release"

### Opção 2: Via GitHub CLI (se disponível)

```bash
# Instalar GitHub CLI
sudo apt install gh

# Autenticar
gh auth login

# Criar release
gh release create v4.2.0 \
  --title "Bagus Browser v4.2.0 - Menu Superior + Melhorias de UX" \
  --notes-file RELEASE_NOTES_v4.2.0.md \
  dist/*.deb \
  dist/*.tar.gz \
  dist/SHA256SUMS
```

### Opção 3: Via Script Automatizado

```bash
# O script já está corrigido, mas precisa de autenticação
bash scripts/release.sh
```

## 📝 Notas de Release (Copiar para GitHub)

```markdown
# Bagus Browser v4.2.0

## 🎉 Release - Menu Superior + Melhorias de UX

Browser minimalista, seguro e privado construído em Go.

### ✨ Novidades

- **📋 Menu Superior Completo** - Menu bar profissional com 4 seções organizadas
  - Menu Arquivo (Nova Aba, Fechar Aba, Sair)
  - Menu Navegação (Voltar, Avançar, Recarregar)
  - Menu Favoritos (Adicionar, Gerenciar)
  - Menu Ferramentas (Buscar, Zoom)
- **🎯 Foco Automático na Barra de URL** - Ao abrir nova aba, cursor vai direto para a barra
- **⌨️ Novo Atalho Ctrl+Q** - Sair do navegador
- Texto selecionado automaticamente ao focar barra de URL

### 🐛 Correções

- **Ctrl+Shift+B** - Gerenciar Favoritos agora funciona corretamente
  - Problema: GTK retorna KEY_B (maiúscula) quando Shift pressionado
  - Solução: Aceita tanto KEY_b quanto KEY_B
- Nome simplificado para "Bagus Browser" (removido "POC" e "WebKit CGO")

### 🎨 Melhorias

- Interface mais intuitiva e profissional
- Descoberta de funcionalidades facilitada via menu
- UX de navegação mais rápida e natural
- Acessibilidade melhorada com múltiplas formas de acesso

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
wget https://github.com/peder1981/bagus-browser-go/releases/download/v4.2.0/bagus-browser_4.2.0_amd64.deb
sudo dpkg -i bagus-browser_4.2.0_amd64.deb
sudo apt-get install -f  # Instalar dependências
\`\`\`

#### Tarball
\`\`\`bash
wget https://github.com/peder1981/bagus-browser-go/releases/download/v4.2.0/bagus-browser_v4.2.0_linux_amd64.tar.gz
tar -xzf bagus-browser_v4.2.0_linux_amd64.tar.gz
./bagus-browser
\`\`\`

### 🔐 Verificação

\`\`\`bash
wget https://github.com/peder1981/bagus-browser-go/releases/download/v4.2.0/SHA256SUMS
sha256sum -c SHA256SUMS
\`\`\`

---

**Bagus Browser - Navegue com privacidade e segurança** 🌐🔒
```

## ✅ Checklist Final

- [ ] Acessar https://github.com/peder1981/bagus-browser-go/releases/new
- [ ] Selecionar tag v4.2.0
- [ ] Adicionar título e descrição
- [ ] Fazer upload dos 3 arquivos de dist/
- [ ] Publicar release
- [ ] Verificar que os downloads funcionam
- [ ] Compartilhar release

## 🎊 Após Publicar

1. Teste a instalação do .deb
2. Teste o tarball
3. Verifique os checksums
4. Compartilhe nas redes sociais
5. Atualize README se necessário

---

**Status:** Pronto para publicação manual no GitHub!
