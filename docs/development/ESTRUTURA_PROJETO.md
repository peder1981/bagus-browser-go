# 📁 Estrutura do Projeto - Bagus Browser

## 🏗️ Organização de Diretórios

```
bagus-browser-go/
│
├── cmd/                          # Aplicações executáveis
│   └── bagus-browser/           # Aplicação principal
│       └── main.go              # Entry point
│
├── pkg/                          # Pacotes públicos (podem ser importados)
│   ├── browser/                 # Funcionalidades do navegador
│   │   ├── bookmarks.go        # Gerenciamento de favoritos
│   │   └── webcontext.go       # Contexto WebKit e multimídia
│   │
│   ├── download/                # Sistema de downloads
│   │   ├── download_handler.go # Handler CGo + callbacks
│   │   └── downloads.go        # Gerenciador de downloads
│   │
│   ├── security/                # Segurança e privacidade
│   │   ├── auth.go             # Autenticação com senha mestre
│   │   ├── crypto.go           # Criptografia (AES-256-GCM)
│   │   ├── privacy.go          # Configurações de privacidade
│   │   └── security.go         # Políticas de segurança
│   │
│   └── ui/                      # Interface gráfica
│       └── settings.go          # Janela de configurações
│
├── internal/                     # Pacotes internos (não exportáveis)
│   └── config/                  # Configurações internas
│       ├── config.go           # Gerenciamento de config
│       ├── cookies.go          # Cookies e cache
│       └── session.go          # Sessões do browser
│
├── assets/                       # Recursos estáticos
│   ├── icon.png                # Ícone do aplicativo
│   └── bagus-browser.desktop   # Desktop entry
│
├── scripts/                      # Scripts auxiliares
│   ├── bagus                   # Script de build
│   └── install.sh              # Script de instalação
│
├── build/                        # Binários compilados
│   └── bagus-browser           # Executável
│
├── dist/                         # Pacotes de distribuição
│   ├── *.deb                   # Pacotes Debian
│   └── *.tar.gz                # Tarballs
│
├── docs/                         # Documentação
│   ├── CONFIGURACOES.md        # Guia de configurações
│   ├── ATALHOS.md              # Lista de atalhos
│   └── ...
│
├── go.mod                        # Dependências Go
├── go.sum                        # Checksums das dependências
├── README.md                     # Documentação principal
├── CHANGELOG.md                  # Histórico de versões
├── LICENSE                       # Licença MIT
└── .gitignore                    # Arquivos ignorados pelo Git
```

---

## 📦 Descrição dos Pacotes

### `cmd/bagus-browser/`
**Entry point da aplicação.**
- Inicialização do GTK
- Setup da janela principal
- Gerenciamento de abas
- Atalhos de teclado

### `pkg/browser/`
**Funcionalidades core do navegador.**
- `bookmarks.go`: Sistema de favoritos criptografados
- `webcontext.go`: WebKit context, multimídia, permissões

### `pkg/download/`
**Sistema completo de downloads.**
- `download_handler.go`: Callbacks CGo, detecção de downloads
- `downloads.go`: UI do gerenciador, progresso, notificações

### `pkg/security/`
**Segurança e privacidade.**
- `auth.go`: Senha mestre com bcrypt
- `crypto.go`: AES-256-GCM + PBKDF2
- `privacy.go`: Bloqueios e anti-fingerprinting
- `security.go`: Políticas de segurança WebKit

### `pkg/ui/`
**Componentes de interface.**
- `settings.go`: Janela de configurações (Ctrl+,)

### `internal/config/`
**Configurações internas (não exportáveis).**
- `config.go`: Persistência de configurações
- `cookies.go`: Gerenciamento de cookies e cache
- `session.go`: Restauração de sessões

---

## 🔧 Como Compilar

```bash
# Desenvolvimento
go build -o build/bagus-browser ./cmd/bagus-browser

# Build completo com pacotes
./scripts/bagus build

# Instalação local
./scripts/bagus install
```

---

## 📝 Convenções

### Nomenclatura
- **Pacotes**: lowercase, sem underscores
- **Arquivos**: snake_case.go
- **Funções exportadas**: PascalCase
- **Funções internas**: camelCase

### Estrutura de Código
- **CGo**: Sempre no topo do arquivo
- **Imports**: Agrupados (stdlib, external, internal)
- **Logs**: Usar emojis para facilitar debug
- **Comentários**: Documentar funções exportadas

### Segurança
- Nunca commitar senhas ou chaves
- Usar criptografia para dados sensíveis
- Validar todas as entradas do usuário

---

## 🎯 Benefícios da Estrutura

✅ **Organização Clara**: Fácil encontrar código
✅ **Modularidade**: Pacotes independentes
✅ **Testabilidade**: Fácil criar testes unitários
✅ **Escalabilidade**: Adicionar features sem bagunça
✅ **Manutenibilidade**: Código limpo e organizado
✅ **Padrão Go**: Segue convenções da comunidade

---

## 📚 Referências

- [Go Project Layout](https://github.com/golang-standards/project-layout)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
