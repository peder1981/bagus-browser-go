# Guia de Migração Python → Go

Este guia ajuda desenvolvedores a migrar do Bagus Browser Python para a versão Go.

## 📋 Índice

- [Visão Geral](#visão-geral)
- [Diferenças Principais](#diferenças-principais)
- [Mapeamento de Funcionalidades](#mapeamento-de-funcionalidades)
- [Migração de Código](#migração-de-código)
- [Migração de Dados](#migração-de-dados)
- [Performance](#performance)
- [Deployment](#deployment)

## Visão Geral

### Por que Migrar?

| Aspecto | Python | Go |
|---------|--------|-----|
| **Performance** | Interpretado | Compilado nativo |
| **Memória** | ~150MB | ~30MB |
| **Startup** | ~2s | ~0.1s |
| **Multiplataforma** | Linux apenas | Linux, Windows, macOS |
| **Distribuição** | Requer Python | Binário único |
| **Concorrência** | Threading | Goroutines |

### Compatibilidade

- ✅ **Dados**: Compatível (mesmo formato)
- ✅ **Configuração**: Compatível (YAML)
- ✅ **Histórico**: Compatível
- ✅ **Favoritos**: Compatível
- ⚠️ **Extensões**: Não compatível (requer reescrita)

## Diferenças Principais

### Arquitetura

#### Python (PySide6)
```
bagus_browser/
├── main.py
├── browser/
│   ├── engine.py
│   └── window.py
├── security/
│   └── validator.py
└── storage/
    └── manager.py
```

#### Go
```
bagus-browser-go/
├── cmd/bagus/main.go
├── internal/
│   ├── browser/
│   ├── security/
│   └── storage/
└── pkg/utils/
```

### Linguagem

#### Tipagem

**Python (Dinâmica)**
```python
def navigate(url):
    if not url:
        return False
    # ...
    return True
```

**Go (Estática)**
```go
func Navigate(url string) error {
    if url == "" {
        return ErrInvalidURL
    }
    // ...
    return nil
}
```

#### Tratamento de Erros

**Python (Exceções)**
```python
try:
    engine.navigate(url)
except InvalidURLError as e:
    print(f"Erro: {e}")
```

**Go (Valores de Retorno)**
```go
if err := engine.Navigate(url); err != nil {
    log.Printf("Erro: %v", err)
}
```

#### Concorrência

**Python (Threading)**
```python
import threading

def load_page():
    # ...

thread = threading.Thread(target=load_page)
thread.start()
thread.join()
```

**Go (Goroutines)**
```go
func loadPage() {
    // ...
}

go loadPage()
// Não precisa de join explícito
```

## Mapeamento de Funcionalidades

### Browser Engine

| Python | Go | Notas |
|--------|-----|-------|
| `BrowserEngine` | `browser.Engine` | Mesma funcionalidade |
| `engine.navigate(url)` | `engine.Navigate(url)` | Retorna error em vez de exception |
| `engine.back()` | `engine.Back()` | Mesma API |
| `engine.forward()` | `engine.Forward()` | Mesma API |
| `engine.reload()` | `engine.Reload()` | Mesma API |

**Python:**
```python
from browser.engine import BrowserEngine

engine = BrowserEngine()
try:
    engine.navigate("https://example.com")
except Exception as e:
    print(f"Erro: {e}")
```

**Go:**
```go
import "github.com/peder1981/bagus-browser-go/internal/browser"

engine, err := browser.NewEngine()
if err != nil {
    log.Fatal(err)
}
defer engine.Close()

if err := engine.Navigate("https://example.com"); err != nil {
    log.Printf("Erro: %v", err)
}
```

### Storage

| Python | Go | Notas |
|--------|-----|-------|
| `StorageManager` | `storage.Manager` | API similar |
| `manager.save(key, value)` | `manager.Save(key, value)` | Go usa []byte |
| `manager.load(key)` | `manager.Load(key)` | Retorna []byte |
| `manager.delete(key)` | `manager.Delete(key)` | Mesma API |

**Python:**
```python
from storage.manager import StorageManager

manager = StorageManager("~/.bagus")
manager.save("config", {"theme": "dark"})
data = manager.load("config")
```

**Go:**
```go
import (
    "encoding/json"
    "github.com/peder1981/bagus-browser-go/internal/storage"
)

manager, err := storage.NewManager("~/.bagus")
if err != nil {
    log.Fatal(err)
}

config := map[string]string{"theme": "dark"}
data, _ := json.Marshal(config)
manager.Save("config", data)

loaded, _ := manager.Load("config")
var loadedConfig map[string]string
json.Unmarshal(loaded, &loadedConfig)
```

### Security

| Python | Go | Notas |
|--------|-----|-------|
| `Validator` | `security.Validator` | API similar |
| `validator.validate_url(url)` | `validator.ValidateURL(url)` | CamelCase em Go |
| `Blocker` | `security.Blocker` | Mesma funcionalidade |
| `blocker.is_blocked(url)` | `blocker.IsBlocked(url)` | Retorna bool |

**Python:**
```python
from security.validator import Validator

validator = Validator()
if validator.validate_url(url):
    print("URL válida")
```

**Go:**
```go
import "github.com/peder1981/bagus-browser-go/internal/security"

validator := security.NewValidator()
if err := validator.ValidateURL(url); err != nil {
    log.Printf("URL inválida: %v", err)
} else {
    log.Println("URL válida")
}
```

### Crypto

| Python | Go | Notas |
|--------|-----|-------|
| `AESCipher` | `crypto.Vault` | Nome diferente |
| `cipher.encrypt(data)` | `vault.Encrypt(data)` | API similar |
| `cipher.decrypt(data)` | `vault.Decrypt(data)` | API similar |

**Python:**
```python
from crypto.aes import AESCipher

cipher = AESCipher("senha")
encrypted = cipher.encrypt("dados secretos")
decrypted = cipher.decrypt(encrypted)
```

**Go:**
```go
import "github.com/peder1981/bagus-browser-go/internal/crypto"

vault, err := crypto.NewVault("senha")
if err != nil {
    log.Fatal(err)
}

encrypted, _ := vault.Encrypt([]byte("dados secretos"))
decrypted, _ := vault.Decrypt(encrypted)
```

## Migração de Código

### Passo 1: Estrutura do Projeto

**Antes (Python):**
```
bagus_browser/
├── main.py
├── requirements.txt
└── src/
```

**Depois (Go):**
```
bagus-browser-go/
├── cmd/bagus/main.go
├── go.mod
└── internal/
```

### Passo 2: Dependências

**Python (requirements.txt):**
```
PySide6==6.5.0
cryptography==41.0.0
```

**Go (go.mod):**
```go
module github.com/peder1981/bagus-browser-go

go 1.21

require (
    golang.org/x/crypto v0.14.0
)
```

### Passo 3: Entry Point

**Python (main.py):**
```python
#!/usr/bin/env python3
import sys
from PySide6.QtWidgets import QApplication
from browser.window import BrowserWindow

def main():
    app = QApplication(sys.argv)
    window = BrowserWindow()
    window.show()
    sys.exit(app.exec())

if __name__ == "__main__":
    main()
```

**Go (cmd/bagus/main.go):**
```go
package main

import (
    "log"
    "github.com/peder1981/bagus-browser-go/internal/browser"
    "github.com/peder1981/bagus-browser-go/internal/ui"
)

func main() {
    engine, err := browser.NewEngine()
    if err != nil {
        log.Fatal(err)
    }
    defer engine.Close()

    window, err := ui.NewWindow("Bagus Browser", 1024, 768)
    if err != nil {
        log.Fatal(err)
    }

    if err := window.Show(); err != nil {
        log.Fatal(err)
    }
}
```

### Passo 4: Classes → Structs

**Python:**
```python
class BrowserEngine:
    def __init__(self):
        self.tabs = []
        self.current_tab = None
    
    def navigate(self, url):
        if not self._validate_url(url):
            raise ValueError("URL inválida")
        # ...
    
    def _validate_url(self, url):
        return url.startswith("http")
```

**Go:**
```go
type Engine struct {
    tabs       []*Tab
    currentTab *Tab
}

func NewEngine() (*Engine, error) {
    return &Engine{
        tabs: make([]*Tab, 0),
    }, nil
}

func (e *Engine) Navigate(url string) error {
    if !e.validateURL(url) {
        return ErrInvalidURL
    }
    // ...
    return nil
}

func (e *Engine) validateURL(url string) bool {
    return strings.HasPrefix(url, "http")
}
```

### Passo 5: Async → Goroutines

**Python:**
```python
import asyncio

async def load_page(url):
    # carregamento assíncrono
    await asyncio.sleep(1)
    return data

async def main():
    result = await load_page("https://example.com")
```

**Go:**
```go
func loadPage(url string) <-chan Data {
    ch := make(chan Data)
    go func() {
        // carregamento assíncrono
        time.Sleep(1 * time.Second)
        ch <- data
    }()
    return ch
}

func main() {
    result := <-loadPage("https://example.com")
}
```

## Migração de Dados

### Formato de Dados

Ambas as versões usam os mesmos formatos:

- **Configuração**: YAML
- **Histórico**: JSON
- **Favoritos**: JSON
- **Cache**: Arquivos binários

### Script de Migração

```bash
#!/bin/bash
# migrate_data.sh

PYTHON_DIR="$HOME/.bagus_browser"
GO_DIR="$HOME/.bagus"

# Criar diretório Go
mkdir -p "$GO_DIR"

# Copiar dados
cp -r "$PYTHON_DIR/history" "$GO_DIR/"
cp -r "$PYTHON_DIR/bookmarks" "$GO_DIR/"
cp -r "$PYTHON_DIR/config.yaml" "$GO_DIR/"

echo "Migração concluída!"
```

### Validação

```bash
# Verificar dados migrados
ls -la ~/.bagus/

# Testar com versão Go
./bagus --verify-data
```

## Performance

### Comparação

| Operação | Python | Go | Melhoria |
|----------|--------|-----|----------|
| Startup | 2.0s | 0.1s | 20x |
| Navegação | 150ms | 50ms | 3x |
| Memória | 150MB | 30MB | 5x |
| Build | N/A | 5s | - |

### Otimizações em Go

```go
// Pool de buffers
var bufferPool = sync.Pool{
    New: func() interface{} {
        return new(bytes.Buffer)
    },
}

// Pré-alocação
tabs := make([]*Tab, 0, 10)

// Concorrência eficiente
var wg sync.WaitGroup
for _, url := range urls {
    wg.Add(1)
    go func(u string) {
        defer wg.Done()
        loadPage(u)
    }(url)
}
wg.Wait()
```

## Deployment

### Python

```bash
# Requer Python instalado
pip install -r requirements.txt
python main.py
```

### Go

```bash
# Binário único, sem dependências
./bagus

# Ou instalar
sudo cp bagus /usr/local/bin/
```

### Cross-Compilation

```bash
# Python: Não suportado nativamente
# Requer PyInstaller ou similar

# Go: Nativo
GOOS=windows GOARCH=amd64 go build -o bagus.exe
GOOS=darwin GOARCH=arm64 go build -o bagus-macos
```

## Checklist de Migração

### Preparação
- [ ] Backup dos dados Python
- [ ] Instalar Go 1.21+
- [ ] Clonar repositório Go
- [ ] Revisar documentação

### Código
- [ ] Migrar lógica de negócio
- [ ] Converter classes para structs
- [ ] Adaptar tratamento de erros
- [ ] Implementar testes

### Dados
- [ ] Migrar configurações
- [ ] Migrar histórico
- [ ] Migrar favoritos
- [ ] Validar dados migrados

### Testes
- [ ] Testes unitários
- [ ] Testes de integração
- [ ] Testes de performance
- [ ] Validação de segurança

### Deploy
- [ ] Build para todas plataformas
- [ ] Testar em cada plataforma
- [ ] Criar instaladores
- [ ] Documentar processo

## Problemas Comuns

### 1. Tipagem

**Problema:** Python é dinâmico, Go é estático

**Solução:**
```go
// Usar interfaces para flexibilidade
type Storage interface {
    Save(key string, value []byte) error
    Load(key string) ([]byte, error)
}
```

### 2. Nil vs None

**Python:**
```python
if value is None:
    # ...
```

**Go:**
```go
if value == nil {
    // ...
}
```

### 3. Listas vs Slices

**Python:**
```python
items = []
items.append(item)
```

**Go:**
```go
items := make([]Item, 0)
items = append(items, item)
```

### 4. Dicionários vs Maps

**Python:**
```python
config = {"key": "value"}
value = config.get("key", "default")
```

**Go:**
```go
config := map[string]string{"key": "value"}
value, ok := config["key"]
if !ok {
    value = "default"
}
```

## Recursos Adicionais

- [Tour of Go](https://go.dev/tour/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go by Example](https://gobyexample.com/)
- [Python to Go](https://github.com/golang/go/wiki/FromPythonToGo)

## Suporte

- **Issues**: https://github.com/peder1981/bagus-browser-go/issues
- **Discussions**: Para dúvidas sobre migração
- **Email**: Para questões específicas

---

**Boa migração! 🚀**
