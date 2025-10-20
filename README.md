# Bagus Browser Go

Browser seguro e multiplataforma focado em privacidade, escrito em Go.

## 🚀 Status do Projeto

**Versão:** 2.0.0-alpha  
**Status:** ✅ Implementação Completa - Pronto para Uso

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
