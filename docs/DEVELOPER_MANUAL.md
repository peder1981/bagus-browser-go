# Manual do Desenvolvedor - Bagus Browser Go

Guia completo para desenvolvedores que desejam contribuir ou estender o Bagus Browser Go.

## 📋 Índice

- [Visão Geral](#visão-geral)
- [Arquitetura](#arquitetura)
- [Setup do Ambiente](#setup-do-ambiente)
- [Estrutura de Código](#estrutura-de-código)
- [Módulos Principais](#módulos-principais)
- [Desenvolvimento de Features](#desenvolvimento-de-features)
- [Testes](#testes)
- [Build e Deploy](#build-e-deploy)
- [Debugging](#debugging)
- [Performance](#performance)
- [Segurança](#segurança)
- [Contribuindo](#contribuindo)

## Visão Geral

### Tecnologias Utilizadas

- **Linguagem**: Go 1.21+
- **UI Framework**: webview/webview (planejado)
- **Criptografia**: golang.org/x/crypto
- **Build**: Make, Scripts multiplataforma
- **CI/CD**: GitHub Actions (planejado)
- **Testes**: Go testing, golangci-lint

### Princípios de Design

1. **Simplicidade**: Código simples e legível
2. **Segurança**: Security by design
3. **Performance**: Otimização constante
4. **Modularidade**: Componentes desacoplados
5. **Testabilidade**: Código testável por design

## Arquitetura

### Visão Geral da Arquitetura

```
┌─────────────────────────────────────────────────────┐
│                   UI Layer                          │
│  (Janelas, WebView, Componentes)                   │
└────────────────┬────────────────────────────────────┘
                 │
┌────────────────▼────────────────────────────────────┐
│              Browser Engine                         │
│  (Navegação, Abas, Histórico)                      │
└────┬──────────┬──────────┬──────────┬──────────────┘
     │          │          │          │
┌────▼────┐ ┌──▼─────┐ ┌──▼──────┐ ┌▼──────────┐
│Security │ │ Crypto │ │ Storage │ │  Utils    │
│ Layer   │ │ Module │ │ Manager │ │           │
└─────────┘ └────────┘ └─────────┘ └───────────┘
```

### Fluxo de Dados

```
User Input → UI → Engine → Security Validation
                    ↓
              Storage/Crypto
                    ↓
              Response → UI
```

### Padrões Arquiteturais

#### 1. Dependency Injection

```go
// Interface define contrato
type Storage interface {
    Save(key string, value []byte) error
    Load(key string) ([]byte, error)
}

// Implementação concreta
type FileStorage struct {
    basePath string
}

// Injeção no construtor
func NewEngine(storage Storage) *Engine {
    return &Engine{storage: storage}
}
```

#### 2. Repository Pattern

```go
type HistoryRepository interface {
    Add(entry *HistoryEntry) error
    Get(id string) (*HistoryEntry, error)
    List(filter Filter) ([]*HistoryEntry, error)
    Delete(id string) error
}
```

#### 3. Factory Pattern

```go
type EngineFactory struct {
    config *Config
}

func (f *EngineFactory) Create() (*Engine, error) {
    storage := f.createStorage()
    security := f.createSecurity()
    return NewEngine(storage, security), nil
}
```

## Setup do Ambiente

### Requisitos

```bash
# Verificar versões
go version  # >= 1.21
git --version
make --version
```

### Instalação

```bash
# 1. Clone o repositório
git clone https://github.com/peder1981/bagus-browser-go.git
cd bagus-browser-go

# 2. Instale dependências
make deps

# 3. Configure pre-commit hooks (opcional)
cat > .git/hooks/pre-commit << 'EOF'
#!/bin/bash
make fmt
make lint
make test
EOF
chmod +x .git/hooks/pre-commit

# 4. Verifique setup
make test
make build
```

### Configuração do Editor

#### VS Code (.vscode/settings.json)

```json
{
  "go.useLanguageServer": true,
  "go.lintTool": "golangci-lint",
  "go.lintOnSave": "workspace",
  "go.formatTool": "gofmt",
  "go.formatFlags": ["-s"],
  "editor.formatOnSave": true,
  "go.testFlags": ["-v", "-race"],
  "go.coverOnSave": true,
  "go.coverageDecorator": {
    "type": "gutter"
  },
  "go.toolsManagement.autoUpdate": true
}
```

#### GoLand

1. Settings → Go → GOROOT: Selecione Go 1.21+
2. Settings → Tools → File Watchers: Adicione gofmt
3. Settings → Editor → Code Style → Go: Use gofmt
4. Settings → Tools → External Tools: Configure golangci-lint

## Estrutura de Código

### Organização de Pacotes

```
bagus-browser-go/
├── cmd/                    # Executáveis
│   └── bagus/
│       └── main.go        # Entry point
│
├── internal/              # Código privado
│   ├── browser/           # Motor do navegador
│   │   ├── engine.go
│   │   ├── engine_test.go
│   │   ├── tab.go
│   │   └── navigation.go
│   │
│   ├── crypto/            # Criptografia
│   │   ├── vault.go
│   │   ├── vault_test.go
│   │   └── keymanager.go
│   │
│   ├── security/          # Segurança
│   │   ├── validator.go
│   │   ├── blocker.go
│   │   └── sandbox.go
│   │
│   ├── storage/           # Armazenamento
│   │   ├── manager.go
│   │   ├── history.go
│   │   └── bookmarks.go
│   │
│   └── ui/                # Interface
│       ├── window.go
│       └── webview.go
│
├── pkg/                   # Código público
│   └── utils/
│       └── system.go
│
├── tests/                 # Testes
│   ├── integration/
│   └── e2e/
│
├── docs/                  # Documentação
├── scripts/               # Scripts
├── configs/               # Configurações
└── assets/                # Recursos
```

### Convenções de Nomenclatura

#### Arquivos

```
engine.go           # Implementação
engine_test.go      # Testes
engine_linux.go     # Específico de plataforma
engine_windows.go
engine_darwin.go
```

#### Pacotes

```go
package browser     // ✅ Lowercase, singular
package browsers    // ❌ Plural
package Browser     // ❌ Uppercase
```

#### Tipos

```go
type Engine struct {}        // ✅ PascalCase (exportado)
type tabManager struct {}    // ✅ camelCase (privado)
type URL_Validator struct {} // ❌ Underscores
```

#### Funções

```go
func NewEngine() *Engine {}           // ✅ Construtor
func (e *Engine) Navigate() error {}  // ✅ Método exportado
func (e *Engine) validateURL() bool{} // ✅ Método privado
func get_url() string {}              // ❌ Underscores
```

## Módulos Principais

### Browser Engine

**Localização**: `internal/browser/`

#### Responsabilidades

- Gerenciamento de abas
- Navegação e histórico
- Coordenação de componentes
- Lifecycle management

#### Estrutura

```go
// engine.go
package browser

import (
    "github.com/peder1981/bagus-browser-go/internal/storage"
    "github.com/peder1981/bagus-browser-go/internal/security"
)

type Engine struct {
    tabs       []*Tab
    currentTab *Tab
    storage    storage.Manager
    security   security.Validator
}

func NewEngine(storage storage.Manager, security security.Validator) (*Engine, error) {
    return &Engine{
        tabs:     make([]*Tab, 0),
        storage:  storage,
        security: security,
    }, nil
}

func (e *Engine) Navigate(url string) error {
    // 1. Validar URL
    if err := e.security.ValidateURL(url); err != nil {
        return fmt.Errorf("URL inválida: %w", err)
    }
    
    // 2. Navegar
    if err := e.currentTab.Navigate(url); err != nil {
        return fmt.Errorf("falha na navegação: %w", err)
    }
    
    // 3. Salvar no histórico
    if err := e.storage.AddHistory(url); err != nil {
        // Log, mas não falhe
        log.Printf("falha ao salvar histórico: %v", err)
    }
    
    return nil
}
```

#### Testes

```go
// engine_test.go
package browser

import (
    "testing"
)

type mockStorage struct {
    addHistoryCalled bool
}

func (m *mockStorage) AddHistory(url string) error {
    m.addHistoryCalled = true
    return nil
}

func TestEngine_Navigate(t *testing.T) {
    storage := &mockStorage{}
    security := &mockSecurity{}
    engine, _ := NewEngine(storage, security)
    
    err := engine.Navigate("https://example.com")
    
    if err != nil {
        t.Errorf("Navigate() erro = %v", err)
    }
    
    if !storage.addHistoryCalled {
        t.Error("AddHistory não foi chamado")
    }
}
```

### Storage Manager

**Localização**: `internal/storage/`

#### Responsabilidades

- Persistência de dados
- Gerenciamento de perfis
- Cache
- Histórico e favoritos

#### Implementação

```go
// manager.go
package storage

import (
    "encoding/json"
    "os"
    "path/filepath"
)

type Manager struct {
    basePath string
}

func NewManager(basePath string) (*Manager, error) {
    // Expandir ~ para home
    if basePath[0] == '~' {
        home, _ := os.UserHomeDir()
        basePath = filepath.Join(home, basePath[1:])
    }
    
    // Criar diretório se não existir
    if err := os.MkdirAll(basePath, 0700); err != nil {
        return nil, fmt.Errorf("falha ao criar diretório: %w", err)
    }
    
    return &Manager{basePath: basePath}, nil
}

func (m *Manager) Save(key string, value interface{}) error {
    // Serializar
    data, err := json.Marshal(value)
    if err != nil {
        return fmt.Errorf("falha ao serializar: %w", err)
    }
    
    // Salvar
    path := filepath.Join(m.basePath, key+".json")
    if err := os.WriteFile(path, data, 0600); err != nil {
        return fmt.Errorf("falha ao escrever: %w", err)
    }
    
    return nil
}

func (m *Manager) Load(key string, dest interface{}) error {
    path := filepath.Join(m.basePath, key+".json")
    
    data, err := os.ReadFile(path)
    if err != nil {
        return fmt.Errorf("falha ao ler: %w", err)
    }
    
    if err := json.Unmarshal(data, dest); err != nil {
        return fmt.Errorf("falha ao deserializar: %w", err)
    }
    
    return nil
}
```

### Security Module

**Localização**: `internal/security/`

#### Validação de Entrada

```go
// validator.go
package security

import (
    "net/url"
    "path/filepath"
    "strings"
)

type Validator struct {
    allowedSchemes []string
}

func NewValidator() *Validator {
    return &Validator{
        allowedSchemes: []string{"http", "https"},
    }
}

func (v *Validator) ValidateURL(rawURL string) error {
    // Parse URL
    u, err := url.Parse(rawURL)
    if err != nil {
        return ErrInvalidURL
    }
    
    // Verificar scheme
    if !v.isAllowedScheme(u.Scheme) {
        return ErrInvalidScheme
    }
    
    // Verificar host
    if u.Host == "" {
        return ErrMissingHost
    }
    
    return nil
}

func (v *Validator) ValidatePath(path string) error {
    // Limpar path
    clean := filepath.Clean(path)
    
    // Verificar path traversal
    if strings.Contains(clean, "..") {
        return ErrPathTraversal
    }
    
    return nil
}

func (v *Validator) SanitizeInput(input string) string {
    // Remover caracteres perigosos
    dangerous := []string{"<", ">", "&", "\"", "'"}
    result := input
    
    for _, char := range dangerous {
        result = strings.ReplaceAll(result, char, "")
    }
    
    return result
}
```

#### Bloqueador de Conteúdo

```go
// blocker.go
package security

import (
    "bufio"
    "os"
    "strings"
)

type Blocker struct {
    blocklist map[string]bool
}

func NewBlocker(blocklistPath string) (*Blocker, error) {
    b := &Blocker{
        blocklist: make(map[string]bool),
    }
    
    if err := b.loadBlocklist(blocklistPath); err != nil {
        return nil, err
    }
    
    return b, nil
}

func (b *Blocker) IsBlocked(domain string) bool {
    // Normalizar domínio
    domain = strings.ToLower(domain)
    domain = strings.TrimPrefix(domain, "www.")
    
    return b.blocklist[domain]
}

func (b *Blocker) loadBlocklist(path string) error {
    file, err := os.Open(path)
    if err != nil {
        return err
    }
    defer file.Close()
    
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        if line != "" && !strings.HasPrefix(line, "#") {
            b.blocklist[line] = true
        }
    }
    
    return scanner.Err()
}
```

### Crypto Module

**Localização**: `internal/crypto/`

#### Vault (Cofre Criptografado)

```go
// vault.go
package crypto

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "crypto/sha256"
    "io"
    
    "golang.org/x/crypto/pbkdf2"
)

type Vault struct {
    key []byte
}

func NewVault(password string) (*Vault, error) {
    // Derivar chave da senha usando PBKDF2
    salt := []byte("bagus-browser-salt") // Em produção, use salt aleatório
    key := pbkdf2.Key([]byte(password), salt, 100000, 32, sha256.New)
    
    return &Vault{key: key}, nil
}

func (v *Vault) Encrypt(plaintext []byte) ([]byte, error) {
    block, err := aes.NewCipher(v.key)
    if err != nil {
        return nil, err
    }
    
    // GCM mode
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return nil, err
    }
    
    // Gerar nonce
    nonce := make([]byte, gcm.NonceSize())
    if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
        return nil, err
    }
    
    // Criptografar
    ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
    return ciphertext, nil
}

func (v *Vault) Decrypt(ciphertext []byte) ([]byte, error) {
    block, err := aes.NewCipher(v.key)
    if err != nil {
        return nil, err
    }
    
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return nil, err
    }
    
    nonceSize := gcm.NonceSize()
    if len(ciphertext) < nonceSize {
        return nil, ErrInvalidCiphertext
    }
    
    nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
    plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
    if err != nil {
        return nil, err
    }
    
    return plaintext, nil
}
```

## Desenvolvimento de Features

### Workflow

1. **Planejamento**
   - Criar issue descrevendo a feature
   - Discutir abordagem
   - Obter aprovação

2. **Implementação**
   - Criar branch: `feature/nome-da-feature`
   - Implementar código
   - Escrever testes
   - Documentar

3. **Review**
   - Abrir Pull Request
   - Passar por code review
   - Fazer ajustes solicitados

4. **Merge**
   - Aprovação de mantenedor
   - Merge para main/develop

### Exemplo: Adicionar Nova Feature

#### 1. Criar Issue

```markdown
## Feature: Modo Leitura

### Descrição
Adicionar modo de leitura que remove elementos desnecessários da página.

### Requisitos
- Botão na barra de ferramentas
- Detectar artigos automaticamente
- Configurações de fonte e tema
- Salvar preferências

### Implementação Proposta
- Novo módulo `internal/reader/`
- Integração com Engine
- UI para configurações
```

#### 2. Implementar

```go
// internal/reader/reader.go
package reader

type Reader struct {
    config *Config
}

type Config struct {
    FontSize  int
    FontFamily string
    Theme     string
}

func NewReader(config *Config) *Reader {
    return &Reader{config: config}
}

func (r *Reader) ExtractContent(html string) (string, error) {
    // Implementação de extração de conteúdo
    return "", nil
}
```

#### 3. Testes

```go
// internal/reader/reader_test.go
package reader

import "testing"

func TestReader_ExtractContent(t *testing.T) {
    reader := NewReader(&Config{})
    
    html := `<html><body><article>Conteúdo</article></body></html>`
    content, err := reader.ExtractContent(html)
    
    if err != nil {
        t.Fatalf("erro inesperado: %v", err)
    }
    
    if content == "" {
        t.Error("conteúdo vazio")
    }
}
```

#### 4. Integração

```go
// internal/browser/engine.go
func (e *Engine) EnableReaderMode() error {
    reader := reader.NewReader(e.readerConfig)
    content, err := reader.ExtractContent(e.currentTab.HTML())
    if err != nil {
        return err
    }
    
    e.currentTab.SetContent(content)
    return nil
}
```

## Testes

### Estratégia de Testes

```
Pirâmide de Testes:
     /\
    /E2E\      ← Poucos, críticos
   /──────\
  /Integr.\   ← Moderados
 /──────────\
/  Unitários \  ← Muitos, rápidos
──────────────
```

### Testes Unitários

```go
func TestEngine_Navigate_ValidURL(t *testing.T) {
    // Arrange
    engine := setupTestEngine()
    url := "https://example.com"
    
    // Act
    err := engine.Navigate(url)
    
    // Assert
    if err != nil {
        t.Errorf("esperado nil, obtido %v", err)
    }
}

func TestEngine_Navigate_InvalidURL(t *testing.T) {
    engine := setupTestEngine()
    
    err := engine.Navigate("invalid-url")
    
    if err == nil {
        t.Error("esperado erro, obtido nil")
    }
}
```

### Table-Driven Tests

```go
func TestValidator_ValidateURL(t *testing.T) {
    tests := []struct {
        name    string
        url     string
        wantErr bool
    }{
        {"URL válida HTTPS", "https://example.com", false},
        {"URL válida HTTP", "http://example.com", false},
        {"URL inválida", "invalid", true},
        {"Scheme inválido", "ftp://example.com", true},
        {"URL vazia", "", true},
    }
    
    validator := NewValidator()
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := validator.ValidateURL(tt.url)
            if (err != nil) != tt.wantErr {
                t.Errorf("ValidateURL() erro = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}
```

### Mocks e Stubs

```go
// Mock de Storage
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

// Uso
func TestWithMock(t *testing.T) {
    mock := &MockStorage{
        SaveFunc: func(key string, value []byte) error {
            if key == "invalid" {
                return errors.New("erro")
            }
            return nil
        },
    }
    
    engine := NewEngine(mock, nil)
    // ... teste
}
```

### Benchmarks

```go
func BenchmarkEngine_Navigate(b *testing.B) {
    engine := setupTestEngine()
    url := "https://example.com"
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = engine.Navigate(url)
    }
}

// Executar: go test -bench=. -benchmem
```

### Coverage

```bash
# Gerar coverage
go test -coverprofile=coverage.out ./...

# Ver relatório HTML
go tool cover -html=coverage.out

# Ver resumo
go tool cover -func=coverage.out
```

## Build e Deploy

### Build Local

```bash
# Build simples
go build -o bagus ./cmd/bagus

# Build otimizado
go build -ldflags="-s -w" -o bagus ./cmd/bagus

# Build com versão
VERSION=2.0.0
go build -ldflags="-s -w -X main.version=$VERSION" -o bagus ./cmd/bagus
```

### Cross-Compilation

```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o bagus-linux-amd64 ./cmd/bagus
GOOS=linux GOARCH=arm64 go build -o bagus-linux-arm64 ./cmd/bagus

# Windows
GOOS=windows GOARCH=amd64 go build -o bagus-windows-amd64.exe ./cmd/bagus

# macOS
GOOS=darwin GOARCH=amd64 go build -o bagus-darwin-amd64 ./cmd/bagus
GOOS=darwin GOARCH=arm64 go build -o bagus-darwin-arm64 ./cmd/bagus
```

### Scripts de Build

```bash
# Usar scripts multiplataforma
./scripts/build.sh all      # Linux/macOS
scripts\build.bat all       # Windows CMD
.\scripts\build.ps1 -Platform all  # PowerShell
```

## Debugging

### Delve

```bash
# Instalar
go install github.com/go-delve/delve/cmd/dlv@latest

# Debug
dlv debug ./cmd/bagus

# Comandos
(dlv) break main.main
(dlv) continue
(dlv) print variavel
(dlv) step
(dlv) next
```

### Logging

```go
import "log"

// Debug logging
log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
log.Printf("DEBUG: valor = %v", valor)

// Conditional logging
if debug {
    log.Printf("DEBUG: %v", data)
}
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
# Analisar
go tool pprof cpu.prof
go tool pprof -http=:8080 mem.prof
```

## Performance

### Otimizações

```go
// ✅ Pré-alocar slices
items := make([]Item, 0, expectedSize)

// ✅ Usar strings.Builder
var sb strings.Builder
sb.WriteString("hello")
sb.WriteString(" world")
result := sb.String()

// ✅ Sync.Pool para objetos reutilizáveis
var bufferPool = sync.Pool{
    New: func() interface{} {
        return new(bytes.Buffer)
    },
}

buf := bufferPool.Get().(*bytes.Buffer)
defer bufferPool.Put(buf)
buf.Reset()
```

### Concorrência

```go
// Worker pool
func processURLs(urls []string) {
    const numWorkers = 10
    jobs := make(chan string, len(urls))
    results := make(chan Result, len(urls))
    
    // Start workers
    for w := 0; w < numWorkers; w++ {
        go worker(jobs, results)
    }
    
    // Send jobs
    for _, url := range urls {
        jobs <- url
    }
    close(jobs)
    
    // Collect results
    for range urls {
        <-results
    }
}
```

## Segurança

### Checklist

- [ ] Validar todas as entradas
- [ ] Sanitizar dados do usuário
- [ ] Usar criptografia forte (AES-256)
- [ ] Evitar SQL injection (usar prepared statements)
- [ ] Prevenir path traversal
- [ ] Implementar rate limiting
- [ ] Logs não devem conter dados sensíveis
- [ ] Usar HTTPS apenas
- [ ] Validar certificados SSL

### Análise de Segurança

```bash
# Scan de vulnerabilidades
go list -json -m all | nancy sleuth

# Análise estática
gosec ./...

# Verificar dependências
go mod verify
```

## Contribuindo

Veja [CONTRIBUTING.md](../CONTRIBUTING.md) para guidelines completos.

### Quick Start

```bash
# 1. Fork e clone
git clone https://github.com/SEU_USUARIO/bagus-browser-go.git

# 2. Criar branch
git checkout -b feature/minha-feature

# 3. Desenvolver
# ... código ...

# 4. Testar
make test
make lint

# 5. Commit
git commit -m "feat: adiciona minha feature"

# 6. Push e PR
git push origin feature/minha-feature
```

---

**Versão**: 2.0.0-alpha  
**Última Atualização**: 2024-10-20

**Happy Coding! 🚀**
