# 🌐 Bagus Browser v5.0.0

**Browser minimalista, seguro e privado - Reescrito em CGo Puro**

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Version](https://img.shields.io/badge/version-5.0.0-blue.svg)](https://github.com/peder1981/bagus-browser-go/releases)
[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8.svg)](https://golang.org/)
[![GTK](https://img.shields.io/badge/GTK-3.0+-green.svg)](https://www.gtk.org/)

---

## 🎊 Novidade: Versão 5.0.0!

**Reescrita completa em CGo puro!** Eliminamos 100% das dependências Go problemáticas (gotk3) e implementamos tudo em **C nativo** com GTK3 e WebKitGTK 4.0.

### ✨ Destaques v5.0.0

- ✅ **100% CGo Puro** - Zero dependências Go externas
- ✅ **Binário 2.3 MB** - Extremamente leve
- ✅ **Estabilidade Máxima** - Sem crashes do gotk3
- ✅ **Interface Completa** - Menu, toolbar, abas, ícone
- ✅ **Configurações Avançadas** - 4 abas (Segurança, Sessões, Performance, Privacidade)
- ✅ **20+ Atalhos** - Produtividade máxima
- ✅ **Favoritos Criptografados** - AES-256-GCM
- ✅ **Migração Automática** - 100% compatível com v4

---

## 🚀 Instalação Rápida

### Via .deb (Recomendado)

```bash
# Download
wget https://github.com/peder1981/bagus-browser-go/releases/download/v5.0.0/bagus-browser_v5.0.0_amd64.deb

# Instalar
sudo dpkg -i bagus-browser_v5.0.0_amd64.deb
sudo apt-get install -f

# Executar
bagus-browser
```

### Via Tarball

```bash
# Download e extrair
wget https://github.com/peder1981/bagus-browser-go/releases/download/v5.0.0/bagus-browser_v5.0.0_linux_amd64.tar.gz
tar -xzf bagus-browser_v5.0.0_linux_amd64.tar.gz

# Instalar
sudo mv bagus-browser /usr/local/bin/
sudo chmod +x /usr/local/bin/bagus-browser

# Executar
bagus-browser
```

---

## 🎯 Funcionalidades

### Interface
- ✅ Barra de menu completa (Arquivo, Navegação, Favoritos, Ferramentas, Editar, Ajuda)
- ✅ Toolbar com botões de navegação
- ✅ Sistema de abas reordenáveis
- ✅ Ícone na barra de tarefas
- ✅ URL bar inteligente (detecta URL vs busca)

### Configurações (Ctrl+,)
- ✅ **Segurança** - Senha mestre com bcrypt
- ✅ **Sessões** - Manter logado, cookies persistentes
- ✅ **Performance** - Slider de cache (100 MB - 5 GB)
- ✅ **Privacidade** - 7 opções de bloqueio

### Favoritos
- ✅ Adicionar com Ctrl+D
- ✅ Gerenciar via menu
- ✅ Criptografia AES-256-GCM
- ✅ Importação automática do v4

### Atalhos (20+)
```
Abas:           Ctrl+T, Ctrl+W, Ctrl+Tab, Ctrl+Shift+Tab, Ctrl+1-9
Navegação:      Alt+←, Alt+→, F5, Ctrl+R, Ctrl+Shift+R, Ctrl+L
Zoom:           Ctrl++, Ctrl+-, Ctrl+0
Busca:          Ctrl+F, F3, Shift+F3
Outros:         Ctrl+P, Ctrl+D, Ctrl+,, Ctrl+Q
```

---

## 🔒 Privacidade e Segurança

### Garantias Absolutas
- ✅ **Zero telemetria** (código aberto, verificável)
- ✅ **Zero rastreamento**
- ✅ **Zero analytics**
- ✅ **Zero crash reports**
- ✅ **DuckDuckGo** como motor de busca padrão

### Criptografia
- ✅ **AES-256-GCM** - Favoritos e sessões
- ✅ **bcrypt** - Senhas (cost 10)
- ✅ **PBKDF2** - Derivação de chaves (100,000 iterações)
- ✅ **Salt aleatório** - 32 bytes

### Bloqueios de Privacidade
- ✅ Cookies de terceiros
- ✅ Geolocalização
- ✅ Notificações
- ✅ Câmera/microfone
- ✅ WebGL (anti-fingerprinting)
- ✅ WebAudio (anti-fingerprinting)
- ✅ Do Not Track

---

## 📊 Requisitos do Sistema

- **Sistema Operacional:** Linux (testado em Ubuntu 22.04+)
- **GTK:** 3.0 ou superior
- **WebKitGTK:** 4.0 (webkit2gtk-4.0-37)
- **Go:** 1.21+ (apenas para compilação)
- **Arquitetura:** amd64 (x86_64)

---

## 🛠️ Compilar do Código-Fonte

```bash
# Clonar repositório
git clone https://github.com/peder1981/bagus-browser-go.git
cd bagus-browser-go

# Instalar dependências
sudo apt-get install -y \
    golang-go \
    libgtk-3-dev \
    libwebkit2gtk-4.0-dev \
    pkg-config

# Compilar e empacotar
./scripts/bagus build

# Instalar
./scripts/bagus install
```

---

## 📁 Estrutura de Dados

Todos os dados são armazenados em `~/.config/bagus-browser/`:

```
~/.config/bagus-browser/
├── bookmarks.enc       # Favoritos criptografados
├── config.enc          # Configurações criptografadas
├── session.enc         # Sessões criptografadas
└── cookies.sqlite      # Cookies (SQLite)

~/.cache/bagus-browser/
└── video/              # Cache de vídeos (configurável)
```

---

## 🔄 Migração do v4

A migração é **100% automática**! Basta instalar o v5 e todos os seus dados serão carregados:

- ✅ Favoritos
- ✅ Configurações
- ✅ Sessões
- ✅ Cookies

---

## 📚 Documentação

- [Release Notes v5.0.0](docs/releases/v5.0.0.md)
- [CHANGELOG](CHANGELOG.md)
- [Guia de Compilação](docs/development/BUILD.md)
- [Segurança](docs/SECURITY.md)
- [Privacidade](docs/PRIVACY.md)

---

## 🤝 Contribuindo

Contribuições são bem-vindas! Por favor:

1. Fork o projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

---

## 📝 Licença

Este projeto está licenciado sob a Licença MIT - veja o arquivo [LICENSE](LICENSE) para detalhes.

---

## 🙏 Agradecimentos

- **GTK Team** - Toolkit gráfico
- **WebKitGTK Team** - Motor de renderização
- **Go Team** - Linguagem de programação
- **DuckDuckGo** - Motor de busca privado
- **Comunidade Open Source** - Suporte e feedback

---

## 📞 Contato

- **Issues:** https://github.com/peder1981/bagus-browser-go/issues
- **Discussões:** https://github.com/peder1981/bagus-browser-go/discussions

---

## ⭐ Estrelas

Se você gostou do Bagus Browser, considere dar uma estrela no GitHub! ⭐

---

**Feito com ❤️ e foco em privacidade**

**Bagus Browser v5.0.0** - Browser minimalista, seguro e privado
