# 🚀 RELEASE PRONTO - ÚLTIMOS PASSOS

## ✅ BUILD COMPLETO!

Os pacotes estão prontos em `dist/`:
- ✅ `bagus-browser_v4.1.0_amd64.deb` (1.3MB)
- ✅ `bagus-browser_v4.1.0_linux_amd64.tar.gz` (1.4MB)
- ✅ `SHA256SUMS` (checksums)

---

## 📤 PUBLICAR NO GITHUB

### Opção 1: GitHub CLI (Automático)

```bash
# 1. Instalar GitHub CLI (se não tiver)
sudo apt install gh

# 2. Autenticar
gh auth login
# Escolha: GitHub.com → HTTPS → Login via browser

# 3. Publicar release
./release.sh
```

### Opção 2: Interface Web (Manual)

1. **Ir para:** https://github.com/peder1981/bagus-browser-go/releases/new

2. **Tag:** v4.1.0

3. **Título:** Bagus Browser v4.1.0

4. **Descrição:**
```markdown
# Bagus Browser v4.1.0

Browser minimalista, seguro e privado construído em Go.

## ✨ Features
- 🌐 WebView completo (WebKit2GTK)
- 📑 Múltiplas abas independentes
- 🔍 Buscar na página (Ctrl+F)
- ⭐ Favoritos com AES-256 (Ctrl+D)
- 📥 Downloads
- 🔍 Zoom (Ctrl++, Ctrl+-, Ctrl+0)
- ⌨️  15 atalhos de teclado
- 🎨 Ícone profissional

## 🔒 Segurança
- AES-256-GCM para favoritos
- PBKDF2 (100,000 iterações)
- Validação de URLs
- Zero telemetria

## 📦 Instalação

### Ubuntu/Debian (.deb)
\`\`\`bash
wget https://github.com/peder1981/bagus-browser-go/releases/download/v4.1.0/bagus-browser_4.1.0_amd64.deb
sudo dpkg -i bagus-browser_4.1.0_amd64.deb
sudo apt-get install -f
\`\`\`

### Tarball
\`\`\`bash
wget https://github.com/peder1981/bagus-browser-go/releases/download/v4.1.0/bagus-browser_v4.1.0_linux_amd64.tar.gz
tar -xzf bagus-browser_v4.1.0_linux_amd64.tar.gz
./bagus-browser
\`\`\`

## 🔐 Verificação
\`\`\`bash
wget https://github.com/peder1981/bagus-browser-go/releases/download/v4.1.0/SHA256SUMS
sha256sum -c SHA256SUMS
\`\`\`
```

5. **Anexar arquivos:**
   - Arraste `dist/bagus-browser_v4.1.0_amd64.deb`
   - Arraste `dist/bagus-browser_v4.1.0_linux_amd64.tar.gz`
   - Arraste `dist/SHA256SUMS`

6. **Publicar release**

---

## 🎯 APÓS PUBLICAR

### Testar Download
```bash
# Baixar do GitHub
wget https://github.com/peder1981/bagus-browser-go/releases/download/v4.1.0/bagus-browser_4.1.0_amd64.deb

# Instalar
sudo dpkg -i bagus-browser_4.1.0_amd64.deb
sudo apt-get install -f

# Executar
bagus-browser
```

### Compartilhar
- Reddit: r/golang, r/privacy, r/linux
- Hacker News
- Twitter/X
- Dev.to

---

## 📊 CHECKSUMS

```
5256de1af2ef45f596307f6a95242860ab375e20a6a31baf854c9339b22bee2d  bagus-browser_v4.1.0_amd64.deb
2aedf6c8434827871fe598ce920e1839e2b6ef7727aebf5ee3f60f3a64910e30  bagus-browser_v4.1.0_linux_amd64.tar.gz
```

---

## ✅ TUDO PRONTO!

Seu browser está compilado, empacotado e pronto para ser publicado! 🚀
