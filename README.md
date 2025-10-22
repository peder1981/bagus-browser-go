# Bagus Browser v4.4.0

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
```bash
wget https://github.com/peder1981/bagus-browser-go/releases/download/v4.4.0/bagus-browser_4.4.0_amd64.deb
sudo dpkg -i bagus-browser_4.4.0_amd64.deb
sudo apt-get install -f
```

#### Tarball
```bash
wget https://github.com/peder1981/bagus-browser-go/releases/download/v4.4.0/bagus-browser_4.4.0_linux_amd64.tar.gz
tar -xzf bagus-browser_4.4.0_linux_amd64.tar.gz
sudo mv bagus-browser /usr/local/bin/
```

### 🔗 Links

- **Código:** https://github.com/peder1981/bagus-browser-go
- **Issues:** https://github.com/peder1981/bagus-browser-go/issues
- **Documentação:** https://github.com/peder1981/bagus-browser-go/tree/main/docs

---

**Data:** 2025-10-22  
**Versão:** 4.4.0
