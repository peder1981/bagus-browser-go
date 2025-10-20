# Guia de Desenvolvimento

Este guia fornece informações detalhadas para desenvolvedores que desejam contribuir ou entender o Bagus Browser Go.

## 📋 Índice

- [Configuração do Ambiente](#configuração-do-ambiente)
- [Estrutura do Projeto](#estrutura-do-projeto)
- [Desenvolvimento](#desenvolvimento)
- [Testes](#testes)
- [Build e Deploy](#build-e-deploy)
- [Debugging](#debugging)
- [Segurança](#segurança)
- [Performance](#performance)

## Configuração do Ambiente

### Requisitos

#### Obrigatórios
- **Go**: 1.21+ ([Download](https://go.dev/dl/))
- **Git**: Para controle de versão
- **Make**: Para automação (Linux/macOS)

#### Opcionais
- **golangci-lint**: Linting (instalado automaticamente)
- **Docker**: Para builds isolados
- **VS Code**: Editor recomendado com extensão Go

### Instalação

```bash
# Clone o repositório
git clone https://github.com/peder1981/bagus-browser-go.git
cd bagus-browser-go

# Instale dependências
make deps

# Verifique instalação
go version
make help
```

### Configuração do Editor

#### VS Code

Extensões recomendadas:
- Go (golang.go)
- Go Test Explorer
- Error Lens

`.vscode/settings.json`:
```json
{
  "go.useLanguageServer": true,
  "go.lintTool": "golangci-lint",
  "go.lintOnSave": "workspace",
  "go.formatTool": "gofmt",
  "editor.formatOnSave": true,
  "go.testFlags": ["-v"],
  "go.coverOnSave": true
}
```

#### GoLand/IntelliJ

- Enable Go Modules
- Configure golangci-lint
- Enable auto-format on save

## Estrutura do Projeto

### Organização de Diretórios

```
bagus-browser-go/
├── cmd/                    # Aplicações executáveis
│   └── bagus/             # Entry point principal
│       └── main.go        # Inicialização
│
├── internal/              # Código privado (não exportável)
│   ├── browser/           # Motor do navegador
│   │   ├── engine.go     # Engine principal
│   │   ├── tab.go        # Gerenciamento de abas
│   │   └── navigation.go # Sistema de navegação
│   │
│   ├── crypto/            # Criptografia
│   │   ├── aes.go        # AES-256-GCM
│   │   ├── keymanager.go # Gerenciamento de chaves
│   │   └── vault.go      # Cofre de dados
│   │
│   ├── security/          # Segurança
│   │   ├── blocker.go    # Bloqueador de conteúdo
│   │   ├── validator.go  # Validação de entrada
│   │   └── sandbox.go    # Sandboxing
│   │
│   ├── ui/                # Interface gráfica
│   │   ├── window.go     # Janela principal
│   │   ├── webview.go    # WebView wrapper
│   │   └── components/   # Componentes UI
│   │
│   └── storage/           # Armazenamento
│       ├── manager.go    # Gerenciador de storage
│       ├── history.go    # Histórico
│       └── bookmarks.go  # Favoritos
│
├── pkg/                   # Código público (exportável)
│   └── utils/            # Utilitários
│       └── system.go     # Funções de sistema
│
├── tests/                 # Testes de integração
│   ├── integration/      # Testes de integração
│   └── e2e/              # Testes end-to-end
│
├── docs/                  # Documentação
├── scripts/               # Scripts de automação
├── assets/                # Recursos estáticos
├── configs/               # Configurações
└── build/                 # Binários compilados
```

### Convenções de Nomenclatura

#### Pacotes
- **Lowercase**: Sempre minúsculo
- **Sem underscores**: Use nomes compostos simples
- **Descritivo**: Nome deve indicar propósito

```go
// ✅ BOM
package browser
package storage

// ❌ RUIM
package Browser
package storage_manager
```

#### Arquivos
- **Lowercase com underscores**: Para separar palavras
- **Sufixo _test.go**: Para testes
- **Sufixo _platform.go**: Para código específico de plataforma

```
engine.go
tab_manager.go
engine_test.go
sandbox_linux.go
sandbox_windows.go
```

#### Funções e Métodos
- **PascalCase**: Funções exportadas
- **camelCase**: Funções privadas

```go
// Exportada (pública)
func NewEngine() *Engine {}

// Privada
func validateURL(url string) error {}
```

## Desenvolvimento

### Workflow de Desenvolvimento

```bash
# 1. Criar branch
git checkout -b feature/minha-feature

# 2. Desenvolver
# ... editar código ...

# 3. Formatar
make fmt

# 4. Lint
make lint

# 5. Testar
make test

# 6. Build
make build

# 7. Commit
git add .
git commit -m "feat: adiciona minha feature"

# 8. Push
git push origin feature/minha-feature
```

### Adicionando Nova Funcionalidade

#### 1. Planejamento
- Crie uma issue descrevendo a feature
- Discuta a abordagem
- Obtenha aprovação

#### 2. Implementação

```go
// internal/browser/feature.go
package browser

// Feature representa uma nova funcionalidade.
type Feature struct {
    name    string
    enabled bool
}

// NewFeature cria uma nova instância de Feature.
func NewFeature(name string) *Feature {
    return &Feature{
        name:    name,
        enabled: true,
    }
}

// Enable ativa a feature.
func (f *Feature) Enable() {
    f.enabled = true
}

// Disable desativa a feature.
func (f *Feature) Disable() {
    f.enabled = false
}
```

#### 3. Testes

```go
// internal/browser/feature_test.go
package browser

import "testing"

func TestNewFeature(t *testing.T) {
    f := NewFeature("test")
    
    if f.name != "test" {
        t.Errorf("esperado 'test', obtido '%s'", f.name)
    }
    
    if !f.enabled {
        t.Error("feature deveria estar habilitada por padrão")
    }
}

func TestFeature_Disable(t *testing.T) {
    f := NewFeature("test")
    f.Disable()
    
    if f.enabled {
        t.Error("feature deveria estar desabilitada")
    }
}
```

#### 4. Documentação

```go
// Package browser implementa o motor do navegador.
//
// Este pacote fornece funcionalidades core para navegação,
// gerenciamento de abas e integração com o motor de renderização.
//
// Exemplo de uso:
//
//     engine, err := browser.NewEngine()
//     if err != nil {
//         log.Fatal(err)
//     }
//     
//     if err := engine.Navigate("https://example.com"); err != nil {
//         log.Fatal(err)
//     }
//
package browser
```

### Padrões de Design

#### Dependency Injection

```go
// ✅ BOM: Injetar dependências
type Engine struct {
    storage  Storage
    security Security
}

func NewEngine(storage Storage, security Security) *Engine {
    return &Engine{
        storage:  storage,
        security: security,
    }
}

// ❌ RUIM: Dependências hard-coded
type Engine struct {
    storage *FileStorage
}

func NewEngine() *Engine {
    return &Engine{
        storage: &FileStorage{},
    }
}
```

#### Interfaces

```go
// ✅ BOM: Definir interfaces pequenas
type Storage interface {
    Save(key string, value []byte) error
    Load(key string) ([]byte, error)
}

type Security interface {
    Validate(input string) error
}

// ❌ RUIM: Interfaces grandes
type Everything interface {
    Save(key string, value []byte) error
    Load(key string) ([]byte, error)
    Validate(input string) error
    Encrypt(data []byte) ([]byte, error)
    Decrypt(data []byte) ([]byte, error)
}
```

#### Error Handling

```go
// ✅ BOM: Erros customizados
var (
    ErrInvalidURL = errors.New("URL inválida")
    ErrNotFound   = errors.New("não encontrado")
)

type ValidationError struct {
    Field string
    Err   error
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validação falhou no campo %s: %v", e.Field, e.Err)
}

// ✅ BOM: Wrap errors
if err := loadConfig(); err != nil {
    return fmt.Errorf("falha ao carregar config: %w", err)
}
```

## Testes

### Tipos de Testes

#### 1. Testes Unitários

```go
func TestEngine_Navigate(t *testing.T) {
    engine := NewEngine()
    
    tests := []struct {
        name    string
        url     string
        wantErr bool
    }{
        {"URL válida", "https://example.com", false},
        {"URL inválida", "invalid", true},
        {"URL vazia", "", true},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := engine.Navigate(tt.url)
            if (err != nil) != tt.wantErr {
                t.Errorf("Navigate() erro = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}
```

#### 2. Testes de Integração

```go
// tests/integration/browser_test.go
func TestBrowserIntegration(t *testing.T) {
    // Setup
    storage := storage.NewManager()
    security := security.NewValidator()
    engine := browser.NewEngine(storage, security)
    
    // Test
    err := engine.Navigate("https://example.com")
    if err != nil {
        t.Fatalf("falha na navegação: %v", err)
    }
    
    // Verify
    history := storage.GetHistory()
    if len(history) != 1 {
        t.Errorf("esperado 1 entrada no histórico, obtido %d", len(history))
    }
}
```

#### 3. Benchmarks

```go
func BenchmarkEngine_Navigate(b *testing.B) {
    engine := NewEngine()
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = engine.Navigate("https://example.com")
    }
}
```

### Executar Testes

```bash
# Todos os testes
make test

# Com coverage
make test-coverage

# Teste específico
go test -v ./internal/browser -run TestEngine_Navigate

# Benchmarks
go test -bench=. ./internal/browser

# Race detector
go test -race ./...
```

### Mocks e Stubs

```go
// Usar interfaces para facilitar mocking
type MockStorage struct {
    SaveFunc func(key string, value []byte) error
    LoadFunc func(key string) ([]byte, error)
}

func (m *MockStorage) Save(key string, value []byte) error {
    if m.SaveFunc != nil {
        return m.SaveFunc(key, value)
    }
    return nil
}

func (m *MockStorage) Load(key string) ([]byte, error) {
    if m.LoadFunc != nil {
        return m.LoadFunc(key)
    }
    return nil, nil
}

// Uso em testes
func TestWithMock(t *testing.T) {
    mock := &MockStorage{
        SaveFunc: func(key string, value []byte) error {
            return errors.New("erro simulado")
        },
    }
    
    engine := NewEngine(mock, nil)
    // ... teste ...
}
```

## Build e Deploy

### Build Local

```bash
# Plataforma atual
make build
./build/bagus

# Plataforma específica
make build-linux
make build-windows
make build-macos

# Todas as plataformas
make build-all
```

### Build com Scripts

```bash
# Linux/macOS
./scripts/build.sh all

# Windows (CMD)
scripts\build.bat all

# Windows (PowerShell)
.\scripts\build.ps1 -Platform all
```

### Cross-Compilation

```bash
# Linux para Windows
GOOS=windows GOARCH=amd64 go build -o bagus.exe ./cmd/bagus

# Linux para macOS
GOOS=darwin GOARCH=arm64 go build -o bagus-macos ./cmd/bagus
```

### Otimizações de Build

```bash
# Build otimizado (menor tamanho)
go build -ldflags="-s -w" -o bagus ./cmd/bagus

# Com informações de versão
go build -ldflags="-X main.version=2.0.0" -o bagus ./cmd/bagus

# Build estático (sem dependências)
CGO_ENABLED=0 go build -o bagus ./cmd/bagus
```

## Debugging

### Logs

```go
import "log"

// Debug logging
log.Printf("DEBUG: navegando para %s", url)

// Error logging
log.Printf("ERROR: falha ao carregar: %v", err)
```

### Delve Debugger

```bash
# Instalar
go install github.com/go-delve/delve/cmd/dlv@latest

# Debug
dlv debug ./cmd/bagus

# Breakpoint
(dlv) break main.main
(dlv) continue
(dlv) print variavel
```

### Profiling

```go
import (
    "runtime/pprof"
    "os"
)

// CPU profiling
f, _ := os.Create("cpu.prof")
pprof.StartCPUProfile(f)
defer pprof.StopCPUProfile()

// Memory profiling
f, _ := os.Create("mem.prof")
pprof.WriteHeapProfile(f)
f.Close()
```

```bash
# Analisar profile
go tool pprof cpu.prof
go tool pprof -http=:8080 cpu.prof
```

## Segurança

### Práticas Seguras

```go
// ✅ BOM: Validar entrada
func Navigate(url string) error {
    if !isValidURL(url) {
        return ErrInvalidURL
    }
    // ...
}

// ✅ BOM: Sanitizar paths
func LoadFile(path string) error {
    cleanPath := filepath.Clean(path)
    if !strings.HasPrefix(cleanPath, basePath) {
        return ErrInvalidPath
    }
    // ...
}

// ✅ BOM: Usar crypto/rand
import "crypto/rand"

func generateKey() ([]byte, error) {
    key := make([]byte, 32)
    _, err := rand.Read(key)
    return key, err
}
```

### Análise de Segurança

```bash
# Scan de vulnerabilidades
go list -json -m all | nancy sleuth

# Análise estática
gosec ./...

# Verificar dependências
go mod verify
```

## Performance

### Profiling

```bash
# CPU profile
go test -cpuprofile=cpu.prof -bench=.
go tool pprof cpu.prof

# Memory profile
go test -memprofile=mem.prof -bench=.
go tool pprof mem.prof
```

### Otimizações

```go
// ✅ BOM: Usar sync.Pool para objetos reutilizáveis
var bufferPool = sync.Pool{
    New: func() interface{} {
        return new(bytes.Buffer)
    },
}

func process() {
    buf := bufferPool.Get().(*bytes.Buffer)
    defer bufferPool.Put(buf)
    buf.Reset()
    // usar buffer
}

// ✅ BOM: Pré-alocar slices
items := make([]Item, 0, expectedSize)

// ✅ BOM: Usar strings.Builder
var sb strings.Builder
sb.WriteString("hello")
sb.WriteString(" world")
result := sb.String()
```

## Recursos Adicionais

- [Effective Go](https://go.dev/doc/effective_go)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md)

---

**Dúvidas?** Consulte a [documentação](README.md) ou abra uma issue!
