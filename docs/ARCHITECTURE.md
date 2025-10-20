# Arquitetura do Bagus Browser Go

## Visão Geral

O Bagus Browser Go é uma reescrita completa do Bagus Browser Python, projetado para ser multiplataforma, performático e seguro.

## Estrutura do Projeto

```
bagus-browser-go/
├── cmd/bagus/              # Entry point da aplicação
│   └── main.go            # Inicialização e CLI
├── internal/              # Código privado (não exportável)
│   ├── browser/           # Motor do navegador
│   │   ├── engine.go     # Engine principal
│   │   ├── tab.go        # Gerenciamento de abas
│   │   └── navigation.go # Sistema de navegação
│   ├── crypto/            # Criptografia
│   │   ├── aes.go        # Criptografia AES-256
│   │   ├── keymanager.go # Gerenciamento de chaves
│   │   └── vault.go      # Cofre de dados
│   ├── security/          # Segurança
│   │   ├── blocker.go    # Bloqueador de conteúdo
│   │   ├── validator.go  # Validação de entrada
│   │   └── sandbox.go    # Sandboxing
│   ├── ui/                # Interface gráfica
│   │   ├── window.go     # Janela principal
│   │   ├── webview.go    # WebView wrapper
│   │   └── components/   # Componentes UI
│   └── storage/           # Armazenamento
│       ├── manager.go    # Gerenciador de storage
│       ├── history.go    # Histórico
│       └── bookmarks.go  # Favoritos
├── pkg/                   # Código público (exportável)
│   └── utils/            # Utilitários
└── assets/               # Recursos estáticos
```

## Componentes Principais

### 1. Browser Engine (`internal/browser`)

Responsável pela lógica principal do navegador:

- **Engine**: Coordena todos os componentes
- **Tab**: Gerencia abas individuais
- **Navigation**: Sistema de navegação e histórico

### 2. Crypto (`internal/crypto`)

Sistema de criptografia:

- **AES-256**: Criptografia de dados
- **Key Manager**: Gerenciamento seguro de chaves
- **Vault**: Armazenamento criptografado

### 3. Security (`internal/security`)

Camada de segurança:

- **Blocker**: Bloqueio de domínios maliciosos
- **Validator**: Validação de entrada (anti-injection)
- **Sandbox**: Isolamento de processos

### 4. UI (`internal/ui`)

Interface gráfica multiplataforma:

- **Window**: Janela principal
- **WebView**: Wrapper para engine de renderização
- **Components**: Componentes reutilizáveis

### 5. Storage (`internal/storage`)

Persistência de dados:

- **Manager**: Gerenciamento de perfis
- **History**: Histórico de navegação
- **Bookmarks**: Favoritos

## Fluxo de Dados

```
┌─────────────┐
│   main.go   │
└──────┬──────┘
       │
       ▼
┌─────────────┐
│   Engine    │◄──────┐
└──────┬──────┘       │
       │              │
       ├──────────────┼─────► Storage Manager
       │              │
       ├──────────────┼─────► Security Layer
       │              │
       ├──────────────┼─────► Crypto Module
       │              │
       └──────────────┴─────► UI Layer
```

## Segurança

### Princípios

1. **Defense in Depth**: Múltiplas camadas de segurança
2. **Least Privilege**: Mínimos privilégios necessários
3. **Zero Trust**: Validação em todos os pontos
4. **Encryption by Default**: Dados sempre criptografados

### Implementações

- Validação de entrada em todas as interfaces
- Sanitização de URLs e paths
- Criptografia AES-256 para dados sensíveis
- Sandboxing de processos
- Bloqueio de conteúdo malicioso

## Performance

### Otimizações

1. **Compilação Nativa**: Binários otimizados por plataforma
2. **Goroutines**: Concorrência eficiente
3. **Memory Pooling**: Reutilização de memória
4. **Lazy Loading**: Carregamento sob demanda
5. **Cache Inteligente**: Cache de recursos frequentes

## Multiplataforma

### Suporte

- **Linux**: Nativo com GTK/Qt
- **Windows**: Nativo com Win32 API
- **macOS**: Nativo com Cocoa

### Abstração

Uso de interfaces para abstrair diferenças entre plataformas:

```go
type WebView interface {
    Navigate(url string) error
    Reload() error
    Back() error
    Forward() error
}
```

## Extensibilidade

### Plugins

Sistema de plugins planejado:

```go
type Plugin interface {
    Name() string
    Version() string
    Init() error
    OnPageLoad(url string) error
}
```

## Roadmap Técnico

### Fase 1: Core (Atual)
- [x] Estrutura base
- [x] Storage manager
- [x] Browser engine básico
- [ ] WebView integration

### Fase 2: Segurança
- [ ] Crypto module completo
- [ ] Content blocker
- [ ] Sandbox implementation

### Fase 3: UI
- [ ] Interface gráfica
- [ ] Tabs management
- [ ] Settings panel

### Fase 4: Features
- [ ] Extensions system
- [ ] Sync (opcional)
- [ ] Advanced privacy features

## Tecnologias

- **Linguagem**: Go 1.21+
- **UI**: webview/webview
- **Crypto**: golang.org/x/crypto
- **Build**: Make, GoReleaser
- **CI/CD**: GitHub Actions

## Contribuindo

Veja [CONTRIBUTING.md](../CONTRIBUTING.md) para guidelines de desenvolvimento.
