# API Reference

Documentação completa da API do Bagus Browser Go.

## 📋 Índice

- [Browser Engine](#browser-engine)
- [Storage](#storage)
- [Security](#security)
- [Crypto](#crypto)
- [UI](#ui)
- [Utils](#utils)

## Browser Engine

### Package: `internal/browser`

Motor principal do navegador.

#### Engine

```go
type Engine struct {
    // Campos privados
}
```

##### Funções

###### NewEngine

```go
func NewEngine() (*Engine, error)
```

Cria uma nova instância do motor do navegador.

**Retorna:**
- `*Engine`: Instância do motor
- `error`: Erro se a inicialização falhar

**Exemplo:**
```go
engine, err := browser.NewEngine()
if err != nil {
    log.Fatal(err)
}
defer engine.Close()
```

###### Navigate

```go
func (e *Engine) Navigate(url string) error
```

Navega para a URL especificada.

**Parâmetros:**
- `url` (string): URL de destino

**Retorna:**
- `error`: Erro se a navegação falhar

**Exemplo:**
```go
if err := engine.Navigate("https://example.com"); err != nil {
    log.Printf("Erro ao navegar: %v", err)
}
```

###### Back

```go
func (e *Engine) Back() error
```

Volta para a página anterior no histórico.

**Retorna:**
- `error`: Erro se não houver página anterior

**Exemplo:**
```go
if err := engine.Back(); err != nil {
    log.Printf("Não há página anterior: %v", err)
}
```

###### Forward

```go
func (e *Engine) Forward() error
```

Avança para a próxima página no histórico.

**Retorna:**
- `error`: Erro se não houver próxima página

###### Reload

```go
func (e *Engine) Reload() error
```

Recarrega a página atual.

**Retorna:**
- `error`: Erro se o reload falhar

###### Close

```go
func (e *Engine) Close() error
```

Fecha o motor e libera recursos.

**Retorna:**
- `error`: Erro se o fechamento falhar

#### Tab

```go
type Tab struct {
    ID      string
    URL     string
    Title   string
    Active  bool
}
```

Representa uma aba do navegador.

##### Funções

###### NewTab

```go
func NewTab(url string) (*Tab, error)
```

Cria uma nova aba.

**Parâmetros:**
- `url` (string): URL inicial

**Retorna:**
- `*Tab`: Nova aba
- `error`: Erro se a criação falhar

**Exemplo:**
```go
tab, err := browser.NewTab("https://example.com")
if err != nil {
    log.Fatal(err)
}
```

###### Close

```go
func (t *Tab) Close() error
```

Fecha a aba.

**Retorna:**
- `error`: Erro se o fechamento falhar

## Storage

### Package: `internal/storage`

Sistema de armazenamento e persistência.

#### Manager

```go
type Manager struct {
    // Campos privados
}
```

Gerenciador de armazenamento.

##### Funções

###### NewManager

```go
func NewManager(basePath string) (*Manager, error)
```

Cria um novo gerenciador de armazenamento.

**Parâmetros:**
- `basePath` (string): Caminho base para armazenamento

**Retorna:**
- `*Manager`: Gerenciador de storage
- `error`: Erro se a inicialização falhar

**Exemplo:**
```go
manager, err := storage.NewManager("~/.bagus")
if err != nil {
    log.Fatal(err)
}
```

###### Save

```go
func (m *Manager) Save(key string, value []byte) error
```

Salva dados no storage.

**Parâmetros:**
- `key` (string): Chave de identificação
- `value` ([]byte): Dados a salvar

**Retorna:**
- `error`: Erro se o salvamento falhar

**Exemplo:**
```go
data := []byte("conteúdo")
if err := manager.Save("config", data); err != nil {
    log.Printf("Erro ao salvar: %v", err)
}
```

###### Load

```go
func (m *Manager) Load(key string) ([]byte, error)
```

Carrega dados do storage.

**Parâmetros:**
- `key` (string): Chave de identificação

**Retorna:**
- `[]byte`: Dados carregados
- `error`: Erro se o carregamento falhar

**Exemplo:**
```go
data, err := manager.Load("config")
if err != nil {
    log.Printf("Erro ao carregar: %v", err)
}
```

###### Delete

```go
func (m *Manager) Delete(key string) error
```

Remove dados do storage.

**Parâmetros:**
- `key` (string): Chave de identificação

**Retorna:**
- `error`: Erro se a remoção falhar

###### List

```go
func (m *Manager) List() ([]string, error)
```

Lista todas as chaves armazenadas.

**Retorna:**
- `[]string`: Lista de chaves
- `error`: Erro se a listagem falhar

## Security

### Package: `internal/security`

Sistema de segurança e validação.

#### Validator

```go
type Validator struct {
    // Campos privados
}
```

Validador de entrada.

##### Funções

###### NewValidator

```go
func NewValidator() *Validator
```

Cria um novo validador.

**Retorna:**
- `*Validator`: Instância do validador

**Exemplo:**
```go
validator := security.NewValidator()
```

###### ValidateURL

```go
func (v *Validator) ValidateURL(url string) error
```

Valida uma URL.

**Parâmetros:**
- `url` (string): URL a validar

**Retorna:**
- `error`: Erro se a URL for inválida

**Exemplo:**
```go
if err := validator.ValidateURL("https://example.com"); err != nil {
    log.Printf("URL inválida: %v", err)
}
```

###### ValidatePath

```go
func (v *Validator) ValidatePath(path string) error
```

Valida um caminho de arquivo.

**Parâmetros:**
- `path` (string): Caminho a validar

**Retorna:**
- `error`: Erro se o caminho for inválido

###### SanitizeInput

```go
func (v *Validator) SanitizeInput(input string) string
```

Sanitiza entrada do usuário.

**Parâmetros:**
- `input` (string): Entrada a sanitizar

**Retorna:**
- `string`: Entrada sanitizada

**Exemplo:**
```go
safe := validator.SanitizeInput(userInput)
```

#### Blocker

```go
type Blocker struct {
    // Campos privados
}
```

Bloqueador de conteúdo.

##### Funções

###### NewBlocker

```go
func NewBlocker() (*Blocker, error)
```

Cria um novo bloqueador.

**Retorna:**
- `*Blocker`: Instância do bloqueador
- `error`: Erro se a inicialização falhar

###### IsBlocked

```go
func (b *Blocker) IsBlocked(url string) bool
```

Verifica se uma URL está bloqueada.

**Parâmetros:**
- `url` (string): URL a verificar

**Retorna:**
- `bool`: true se bloqueada, false caso contrário

**Exemplo:**
```go
if blocker.IsBlocked("https://malware.com") {
    log.Println("URL bloqueada")
}
```

###### AddToBlocklist

```go
func (b *Blocker) AddToBlocklist(domain string) error
```

Adiciona domínio à lista de bloqueio.

**Parâmetros:**
- `domain` (string): Domínio a bloquear

**Retorna:**
- `error`: Erro se a adição falhar

## Crypto

### Package: `internal/crypto`

Sistema de criptografia.

#### Vault

```go
type Vault struct {
    // Campos privados
}
```

Cofre de dados criptografados.

##### Funções

###### NewVault

```go
func NewVault(password string) (*Vault, error)
```

Cria um novo cofre.

**Parâmetros:**
- `password` (string): Senha mestra

**Retorna:**
- `*Vault`: Instância do cofre
- `error`: Erro se a criação falhar

**Exemplo:**
```go
vault, err := crypto.NewVault("senha-forte")
if err != nil {
    log.Fatal(err)
}
```

###### Encrypt

```go
func (v *Vault) Encrypt(plaintext []byte) ([]byte, error)
```

Criptografa dados.

**Parâmetros:**
- `plaintext` ([]byte): Dados a criptografar

**Retorna:**
- `[]byte`: Dados criptografados
- `error`: Erro se a criptografia falhar

**Exemplo:**
```go
encrypted, err := vault.Encrypt([]byte("dados secretos"))
if err != nil {
    log.Fatal(err)
}
```

###### Decrypt

```go
func (v *Vault) Decrypt(ciphertext []byte) ([]byte, error)
```

Descriptografa dados.

**Parâmetros:**
- `ciphertext` ([]byte): Dados criptografados

**Retorna:**
- `[]byte`: Dados descriptografados
- `error`: Erro se a descriptografia falhar

**Exemplo:**
```go
decrypted, err := vault.Decrypt(encrypted)
if err != nil {
    log.Fatal(err)
}
```

## UI

### Package: `internal/ui`

Interface gráfica.

#### Window

```go
type Window struct {
    // Campos privados
}
```

Janela principal do navegador.

##### Funções

###### NewWindow

```go
func NewWindow(title string, width, height int) (*Window, error)
```

Cria uma nova janela.

**Parâmetros:**
- `title` (string): Título da janela
- `width` (int): Largura em pixels
- `height` (int): Altura em pixels

**Retorna:**
- `*Window`: Instância da janela
- `error`: Erro se a criação falhar

**Exemplo:**
```go
window, err := ui.NewWindow("Bagus Browser", 1024, 768)
if err != nil {
    log.Fatal(err)
}
```

###### Show

```go
func (w *Window) Show() error
```

Exibe a janela.

**Retorna:**
- `error`: Erro se a exibição falhar

###### Hide

```go
func (w *Window) Hide() error
```

Oculta a janela.

**Retorna:**
- `error`: Erro se a ocultação falhar

###### Close

```go
func (w *Window) Close() error
```

Fecha a janela.

**Retorna:**
- `error`: Erro se o fechamento falhar

## Utils

### Package: `pkg/utils`

Utilitários gerais.

#### System

##### Funções

###### GetHomeDir

```go
func GetHomeDir() (string, error)
```

Retorna o diretório home do usuário.

**Retorna:**
- `string`: Caminho do diretório home
- `error`: Erro se não conseguir determinar

**Exemplo:**
```go
home, err := utils.GetHomeDir()
if err != nil {
    log.Fatal(err)
}
```

###### GetConfigDir

```go
func GetConfigDir() (string, error)
```

Retorna o diretório de configuração.

**Retorna:**
- `string`: Caminho do diretório de config
- `error`: Erro se não conseguir determinar

###### EnsureDir

```go
func EnsureDir(path string) error
```

Garante que um diretório existe.

**Parâmetros:**
- `path` (string): Caminho do diretório

**Retorna:**
- `error`: Erro se não conseguir criar

**Exemplo:**
```go
if err := utils.EnsureDir("~/.bagus/data"); err != nil {
    log.Fatal(err)
}
```

## Erros Comuns

### Browser

```go
var (
    ErrInvalidURL     = errors.New("URL inválida")
    ErrNavigationFailed = errors.New("falha na navegação")
    ErrNoHistory      = errors.New("sem histórico")
)
```

### Storage

```go
var (
    ErrNotFound       = errors.New("não encontrado")
    ErrInvalidKey     = errors.New("chave inválida")
    ErrStorageFull    = errors.New("armazenamento cheio")
)
```

### Security

```go
var (
    ErrBlocked        = errors.New("conteúdo bloqueado")
    ErrInvalidInput   = errors.New("entrada inválida")
    ErrPathTraversal  = errors.New("tentativa de path traversal")
)
```

### Crypto

```go
var (
    ErrEncryptionFailed = errors.New("falha na criptografia")
    ErrDecryptionFailed = errors.New("falha na descriptografia")
    ErrInvalidKey       = errors.New("chave inválida")
)
```

## Exemplos Completos

### Exemplo 1: Navegação Básica

```go
package main

import (
    "log"
    "github.com/peder1981/bagus-browser-go/internal/browser"
)

func main() {
    // Criar engine
    engine, err := browser.NewEngine()
    if err != nil {
        log.Fatal(err)
    }
    defer engine.Close()
    
    // Navegar
    if err := engine.Navigate("https://example.com"); err != nil {
        log.Fatal(err)
    }
    
    // Voltar
    if err := engine.Back(); err != nil {
        log.Printf("Não há página anterior: %v", err)
    }
}
```

### Exemplo 2: Storage com Criptografia

```go
package main

import (
    "log"
    "github.com/peder1981/bagus-browser-go/internal/storage"
    "github.com/peder1981/bagus-browser-go/internal/crypto"
)

func main() {
    // Criar vault
    vault, err := crypto.NewVault("senha-forte")
    if err != nil {
        log.Fatal(err)
    }
    
    // Criar storage
    manager, err := storage.NewManager("~/.bagus")
    if err != nil {
        log.Fatal(err)
    }
    
    // Criptografar dados
    data := []byte("dados secretos")
    encrypted, err := vault.Encrypt(data)
    if err != nil {
        log.Fatal(err)
    }
    
    // Salvar
    if err := manager.Save("secret", encrypted); err != nil {
        log.Fatal(err)
    }
    
    // Carregar
    loaded, err := manager.Load("secret")
    if err != nil {
        log.Fatal(err)
    }
    
    // Descriptografar
    decrypted, err := vault.Decrypt(loaded)
    if err != nil {
        log.Fatal(err)
    }
    
    log.Printf("Dados: %s", decrypted)
}
```

### Exemplo 3: Validação e Bloqueio

```go
package main

import (
    "log"
    "github.com/peder1981/bagus-browser-go/internal/security"
)

func main() {
    // Criar validator
    validator := security.NewValidator()
    
    // Criar blocker
    blocker, err := security.NewBlocker()
    if err != nil {
        log.Fatal(err)
    }
    
    url := "https://example.com"
    
    // Validar
    if err := validator.ValidateURL(url); err != nil {
        log.Printf("URL inválida: %v", err)
        return
    }
    
    // Verificar bloqueio
    if blocker.IsBlocked(url) {
        log.Println("URL bloqueada")
        return
    }
    
    log.Println("URL válida e não bloqueada")
}
```

---

**Versão da API**: 2.0.0-alpha  
**Última atualização**: 2024-10-20
