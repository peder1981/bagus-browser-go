# 🌐 Bagus Browser v3.0.0

**Browser leve, seguro e focado em privacidade com arquitetura de 2 janelas**, escrito em Go.**

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go)](https://golang.org)
[![Release](https://img.shields.io/github/v/release/peder1981/bagus-browser-go)](https://github.com/peder1981/bagus-browser-go/releases)

## ✨ Características

- 🚀 **Leve e Rápido**: ~4MB, inicialização instantânea
- 🔒 **Zero Telemetria**: 100% privado, sem rastreamento
- 🛡️ **Bloqueador de Ads**: Integrado e configurável
- 🌍 **Compatibilidade Total**: Suporta todos os sites modernos
- 💻 **Multiplataforma**: Linux e Windows
- 🎨 **Interface Moderna**: Design limpo e intuitivo

## 📦 Instalação Rápida

### Linux (Debian/Ubuntu)

```bash
# Download da última versão
wget https://github.com/peder1981/bagus-browser-go/releases/latest/download/bagus-browser_1.0.0_amd64.deb

# Instalar
sudo dpkg -i bagus-browser_1.0.0_amd64.deb

# Executar
bagus-browser
```

### Linux (Outras Distribuições)

```bash
# Download
wget https://github.com/peder1981/bagus-browser-go/releases/latest/download/bagus-2.0.0-linux-amd64.tar.gz

# Extrair e instalar
tar -xzf bagus-2.0.0-linux-amd64.tar.gz
sudo mv bagus-linux-amd64 /usr/local/bin/bagus
sudo chmod +x /usr/local/bin/bagus

# Executar
bagus
```

### Windows

1. Baixe `bagus-2.0.0-windows-amd64.zip` da [página de releases](https://github.com/peder1981/bagus-browser-go/releases)
2. Extraia o arquivo
3. Execute `bagus-windows-amd64.exe`

## 🛠️ Compilar do Código Fonte

### Pré-requisitos

- Go 1.21 ou superior
- GCC (para CGO)
- WebKit2GTK (Linux) ou WebView2 (Windows)

### Linux

```bash
# Instalar dependências
sudo apt install golang gcc libwebkit2gtk-4.1-dev

# Clonar repositório
git clone https://github.com/peder1981/bagus-browser-go.git
cd bagus-browser-go

# Compilar
./build-all.sh

# Executar
./build/bagus-linux-amd64
```

## 🎯 Uso

Após instalação, execute:

```bash
bagus-browser  # Se instalou via .deb
# ou
bagus          # Se instalou manualmente
```

### Primeira Execução

1. Digite um username (3-32 caracteres, apenas letras, números, _ ou -)
2. O browser abre automaticamente
3. Digite uma URL ou termo de busca na barra
4. Navegue livremente!

## 🔧 Recursos

### Navegação
- Barra de endereço inteligente (URL ou busca)
- Histórico de navegação com busca
- Sugestões automáticas
- Botões voltar/avançar/recarregar

### Segurança
- Bloqueador de anúncios integrado
- Validação de URLs
- Sanitização de entrada
- Proteção contra múltiplas instâncias

### Privacidade
- Zero telemetria
- Sem rastreamento
- Dados armazenados localmente
- Código 100% open source

## ⌨️ Atalhos de Teclado

O Bagus Browser suporta atalhos de teclado para navegação rápida:

- **Ctrl+L** - Navegar para URL
- **Alt+←** - Voltar
- **Alt+→** - Avançar  
- **Ctrl+R / F5** - Recarregar
- **Ctrl+H** - Ver histórico

[Ver todos os atalhos](docs/KEYBOARD_SHORTCUTS.md)

## 📚 Documentação

### Guias de Instalação
- [Instalação Completa](docs/install/INSTALL.md) - Guia detalhado de instalação
- [Instalação CEF](docs/install/INSTALL_CEF.md) - Versão com 100% compatibilidade
- [Instalação de Ícone](docs/install/INSTALACAO_ICONE.md) - Integração com desktop

### Primeiros Passos
- [Quick Start](docs/getting-started/QUICK_START.md) - Comece rapidamente
- [Quick Start CEF](docs/getting-started/QUICKSTART_CEF.md) - Versão completa
- [Quick Start GitHub](docs/getting-started/QUICK_START_GITHUB.md) - Para desenvolvedores

### Desenvolvimento
- [Build Guide](docs/build/BUILD.md) - Como compilar o projeto
- [Build Tags](docs/build/BUILD_TAGS.md) - Tags de compilação
- [Testing](TESTING.md) - Guia de testes
- [Linting](LINTING.md) - Qualidade de código
- [Contributing](CONTRIBUTING.md) - Como contribuir

### Uso
- [Como Usar](docs/usage/COMO_USAR.md) - Guia de uso
- [Exemplos](docs/usage/EXEMPLOS.md) - Exemplos práticos
- [Comandos Rápidos](docs/usage/COMANDOS_RAPIDOS.md) - Referência rápida

### Release
- [Release Guide](docs/release/RELEASE_GUIDE.md) - Como criar releases
- [Release Instructions](docs/release/RELEASE_INSTRUCTIONS.md) - Instruções detalhadas

### Segurança
- [Security Policy](SECURITY.md) - Política de segurança
- [Privacy](docs/security/PRIVACY.md) - Privacidade e dados

### CEF (Chromium Embedded Framework)
- [CEF README](docs/cef/README_CEF.md) - Documentação CEF
- [CEF Status](docs/cef/CEF_STATUS.md) - Status da implementação

## 📊 Status do Projeto

**Versão Atual:** 2.0.0  
**Status:** ✅ Estável e Pronto para Produção  
**Compatibilidade:** 100% dos sites modernos

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Security](https://img.shields.io/badge/security-hardened-green.svg)](SECURITY.md)
[![Platform](https://img.shields.io/badge/platform-Linux%20%7C%20Windows%20%7C%20macOS-lightgrey.svg)](https://github.com/peder1981/bagus-browser-go)

## 🎯 Sobre o Projeto

Esta é a versão **multiplataforma** do Bagus Browser, reescrita em Go para oferecer:

- ✅ **Performance Superior**: Compilado nativamente para cada plataforma
- ✅ **Multiplataforma**: Linux, Windows, macOS
- ✅ **Binário Único**: Sem dependências externas
- ✅ **Menor Consumo**: Footprint de memória reduzido
- ✅ **Segurança Nativa**: Criptografia e isolamento em nível de sistema

## 📋 Funcionalidades

### ✅ Implementadas
- 🔒 **Segurança e Privacidade** - Validação de entrada, bloqueador de domínios
- 🌐 **Multiplataforma** - Linux, Windows, macOS (webview nativo)
- ⚡ **Performance** - Escrito em Go, binário único de 3.7MB
- 🗂️ **Sistema de Abas** - Múltiplas abas de navegação
- 📜 **Histórico** - Busca e sugestões automáticas
- ⚙️ **Configurável** - 40+ configurações de segurança
- 🛡️ **Bloqueador** - Bloqueio de domínios maliciosos
- 💾 **Persistência** - Histórico e configurações salvos

### 🚧 Planejadas (Futuro)
- 🎨 **Interface Avançada** - Painéis laterais customizados
- 🔧 **Extensível** - Sistema de plugins e extensões
- 🔐 **Proxy** - Suporte SOCKS5/HTTP configurável
- 🛠️ **DevTools** - Ferramentas de desenvolvedor integradas

## 🔄 Relação com a Versão Python

Este projeto é uma **reescrita completa** do [Bagus Browser Python](https://github.com/peder1981/bagus_browser), mantendo as mesmas funcionalidades de segurança e privacidade, mas com melhorias significativas de performance e portabilidade.

## 🔒 Características de Segurança

- **Criptografia AES-256**: Proteção de dados do usuário
- **Isolamento de Processos**: Sandboxing nativo por plataforma
- **Bloqueio de Conteúdo**: Lista de domínios maliciosos e rastreadores
- **Validação de Entrada**: Proteção contra injeções e path traversal
- **Configurações Seguras**: Proteção contra XSS, fingerprinting e vazamentos
- **Zero Trust**: Nenhum dado é enviado para servidores externos

## 📋 Requisitos

- **Go**: 1.21 ou superior (apenas para compilação)
- **Sistema Operacional**: 
  - Linux (Ubuntu 20.04+, Fedora 35+, Arch Linux)
  - Windows 10/11
  - macOS 11+ (Big Sur ou superior)

## 🚀 Instalação

### Binários Pré-compilados (Recomendado)

```bash
# Linux
wget https://github.com/peder1981/bagus-browser-go/releases/latest/download/bagus-linux-amd64
chmod +x bagus-linux-amd64
./bagus-linux-amd64

# Windows (PowerShell)
Invoke-WebRequest -Uri "https://github.com/peder1981/bagus-browser-go/releases/latest/download/bagus-windows-amd64.exe" -OutFile "bagus.exe"
.\bagus.exe

# macOS
wget https://github.com/peder1981/bagus-browser-go/releases/latest/download/bagus-darwin-amd64
chmod +x bagus-darwin-amd64
./bagus-darwin-amd64
```

### Compilar do Código-Fonte

```bash
# Clone o repositório
git clone https://github.com/peder1981/bagus-browser-go.git
cd bagus-browser-go

# Compile
make build

# Ou compile para plataforma específica
make build-linux
make build-windows
make build-macos

# Execute
./build/bagus
```

## 📖 Uso

```bash
# Iniciar o browser
./bagus

# Iniciar com perfil específico
./bagus --profile=work

# Modo debug
./bagus --debug

# Configuração customizada
./bagus --config=/path/to/config.yaml
```

## 🏗️ Arquitetura

```
bagus-browser-go/
├── cmd/bagus/          # Entry point da aplicação
├── internal/           # Código privado
│   ├── browser/        # Engine do navegador
│   ├── crypto/         # Criptografia e segurança
│   ├── security/       # Bloqueio e validação
│   ├── ui/             # Interface gráfica
│   └── storage/        # Persistência de dados
├── pkg/                # Pacotes reutilizáveis
└── assets/             # Recursos estáticos
```

## 🛠️ Desenvolvimento

```bash
# Instalar dependências
make deps

# Executar testes
make test

# Executar testes com coverage
make test-coverage

# Lint
make lint

# Formatar código
make fmt

# Build para todas as plataformas
make build-all
```

## 🔐 Segurança

Para informações detalhadas sobre segurança, consulte [SECURITY.md](SECURITY.md).

### Reportar Vulnerabilidades

**Não abra issues públicas para vulnerabilidades de segurança.**

Entre em contato diretamente através de: security@bagus-browser.dev

## 📚 Documentação

- [Arquitetura](docs/ARCHITECTURE.md)
- [Guia de Desenvolvimento](docs/DEVELOPMENT.md)
- [API Reference](docs/API.md)
- [Guia de Migração Python→Go](docs/MIGRATION.md)

## 🤝 Contribuindo

Contribuições são bem-vindas! Por favor:

1. Fork o projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

Veja [CONTRIBUTING.md](CONTRIBUTING.md) para mais detalhes.

## 📊 Roadmap

- [x] Estrutura base do projeto
- [ ] Engine de navegação (WebView)
- [ ] Sistema de criptografia
- [ ] Bloqueador de conteúdo
- [ ] Interface gráfica multiplataforma
- [ ] Sistema de extensões
- [ ] Sincronização segura (opcional)
- [ ] Builds automatizados
- [ ] Testes de integração
- [ ] Documentação completa

## 📝 Licença

Este projeto está sob licença MIT. Veja o arquivo [LICENSE](LICENSE) para detalhes.

## 🙏 Agradecimentos

- Comunidade Go
- Projeto [webview](https://github.com/webview/webview)
- Mantenedores de listas de bloqueio
- Versão original em Python

## 🔗 Links

- **Versão Python**: [bagus_browser](https://github.com/peder1981/bagus_browser)
- **Website**: [bagus-browser.dev](https://bagus-browser.dev)
- **Documentação**: [docs.bagus-browser.dev](https://docs.bagus-browser.dev)

---

**Desenvolvido com ❤️ para privacidade e segurança**
