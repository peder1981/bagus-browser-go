# 🌐 Bagus Browser

**Browser minimalista, seguro e privado para Linux**

[![Versão](https://img.shields.io/badge/versão-4.6.3-blue.svg)](CHANGELOG.md)
[![Licença](https://img.shields.io/badge/licença-MIT-green.svg)](LICENSE)
[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8.svg)](https://go.dev/)
[![WebKit](https://img.shields.io/badge/WebKit-2.0-orange.svg)](https://webkitgtk.org/)

---

## ✨ Características

- 🔒 **Zero Telemetria** - Sem rastreamento, sem analytics
- 🛡️ **Privacidade Máxima** - Anti-fingerprinting, bloqueio de trackers
- ⚡ **Leve e Rápido** - ~7MB, baixo uso de memória
- 🎨 **Interface Limpa** - GTK+ nativo, integração perfeita com Linux
- 📥 **Downloads Robustos** - Gerenciador visual com progresso em tempo real
- 🔐 **Criptografia** - AES-256-GCM para dados sensíveis
- 🎥 **Multimídia** - Suporte a webcam/microfone (Meet, YouTube Music)

---

## 🚀 Instalação Rápida

### Ubuntu/Debian
```bash
# Download do .deb
wget https://github.com/peder1981/bagus-browser-go/releases/latest/download/bagus-browser_amd64.deb

# Instalar
sudo dpkg -i bagus-browser_amd64.deb
sudo apt-get install -f  # Resolver dependências
```

### Compilar do Fonte
```bash
# Dependências
sudo apt-get install golang libwebkit2gtk-4.0-dev libgtk-3-dev

# Clonar
git clone https://github.com/peder1981/bagus-browser-go.git
cd bagus-browser-go

# Compilar
go build -o build/bagus-browser ./cmd/bagus-browser

# Executar
./build/bagus-browser
```

---

## ⌨️ Atalhos Principais

| Atalho | Ação |
|--------|------|
| `Ctrl+T` | Nova aba |
| `Ctrl+W` | Fechar aba |
| `Ctrl+L` | Focar barra de URL |
| `Ctrl+F` | Buscar na página |
| `Ctrl+J` | Gerenciador de downloads |
| `Ctrl+,` | Configurações |
| `Ctrl+Shift+T` | Reabrir aba fechada |
| `Alt+←/→` | Voltar/Avançar |
| `F5` | Recarregar |

[Ver todos os atalhos →](docs/ATALHOS.md)

---

## 📁 Estrutura do Projeto

```
bagus-browser-go/
├── cmd/bagus-browser/    # Aplicação principal
├── pkg/                  # Pacotes públicos
│   ├── browser/         # Core do navegador
│   ├── download/        # Sistema de downloads
│   ├── security/        # Segurança e privacidade
│   └── ui/              # Interface gráfica
├── internal/config/      # Configurações internas
├── assets/              # Recursos (ícones, etc)
├── scripts/             # Scripts de build
└── docs/                # Documentação
```

[Ver estrutura completa →](ESTRUTURA_PROJETO.md)

---

## 🔧 Desenvolvimento

### Requisitos
- Go 1.21+
- WebKit2GTK 4.0
- GTK+ 3.0
- Linux (Ubuntu 20.04+, Debian 11+, Fedora 35+)

### Build
```bash
# Desenvolvimento
go build -o build/bagus-browser ./cmd/bagus-browser

# Build completo (cria .deb e tarball)
./scripts/bagus build

# Instalar localmente
./scripts/bagus install
```

### Testes
```bash
# Executar testes
go test ./...

# Teste rápido
./build/bagus-browser
```

---

## 📚 Documentação

- [Guia de Configurações](docs/CONFIGURACOES.md)
- [Lista de Atalhos](docs/ATALHOS.md)
- [Changelog](CHANGELOG.md)
- [Como Compilar](BUILD.md)
- [Estrutura do Projeto](ESTRUTURA_PROJETO.md)

---

## 🛡️ Segurança e Privacidade

### O que bloqueamos
- ✅ Telemetria e analytics
- ✅ Third-party cookies
- ✅ Geolocalização
- ✅ WebGL (anti-fingerprinting)
- ✅ WebAudio (anti-fingerprinting)
- ✅ DNS prefetching
- ✅ Plugins (Flash, Java)

### O que criptografamos
- ✅ Favoritos (AES-256-GCM)
- ✅ Sessões (AES-256-GCM)
- ✅ Configurações (AES-256-GCM)
- ✅ Senha mestre (bcrypt cost 10)

---

## 🤝 Contribuindo

Contribuições são bem-vindas! Por favor:

1. Fork o projeto
2. Crie uma branch (`git checkout -b feature/MinhaFeature`)
3. Commit suas mudanças (`git commit -m 'Adiciona MinhaFeature'`)
4. Push para a branch (`git push origin feature/MinhaFeature`)
5. Abra um Pull Request

---

## 📝 Licença

Este projeto está licenciado sob a Licença MIT - veja o arquivo [LICENSE](LICENSE) para detalhes.

---

## 🙏 Agradecimentos

- [WebKitGTK](https://webkitgtk.org/) - Engine de renderização
- [gotk3](https://github.com/gotk3/gotk3) - Bindings Go para GTK
- Comunidade Go e Linux

---

## 📧 Contato

- **Issues**: [GitHub Issues](https://github.com/peder1981/bagus-browser-go/issues)
- **Discussões**: [GitHub Discussions](https://github.com/peder1981/bagus-browser-go/discussions)

---

<p align="center">
  <b>Feito com ❤️ para a comunidade Linux</b>
</p>
