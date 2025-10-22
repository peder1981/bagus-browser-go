# Bagus Browser v4.2.0

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
```bash
wget https://github.com/peder1981/bagus-browser-go/releases/download/v4.2.0/bagus-browser_4.2.0_amd64.deb
sudo dpkg -i bagus-browser_4.2.0_amd64.deb
sudo apt-get install -f  # Instalar dependências
```

#### Tarball
```bash
wget https://github.com/peder1981/bagus-browser-go/releases/download/v4.2.0/bagus-browser_v4.2.0_linux_amd64.tar.gz
tar -xzf bagus-browser_v4.2.0_linux_amd64.tar.gz
./bagus-browser
```

### 🔐 Verificação

```bash
wget https://github.com/peder1981/bagus-browser-go/releases/download/v4.2.0/SHA256SUMS
sha256sum -c SHA256SUMS
```

### 📚 Documentação

- [README](https://github.com/peder1981/bagus-browser-go/blob/main/README.md)
- [CHANGELOG](https://github.com/peder1981/bagus-browser-go/blob/main/CHANGELOG.md)
- [Segurança](https://github.com/peder1981/bagus-browser-go/blob/main/docs/SECURITY.md)
- [Privacidade](https://github.com/peder1981/bagus-browser-go/blob/main/docs/PRIVACY.md)

---

**Bagus Browser - Navegue com privacidade e segurança** 🌐🔒
