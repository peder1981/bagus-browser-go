# 🏷️ Build Tags - Bagus Browser

## O que são Build Tags?

Build tags (ou build constraints) são diretivas do Go que controlam **quando** um arquivo deve ser compilado.

No Bagus Browser, usamos build tags para **separar completamente** as versões Webview e CEF.

---

## 📋 Build Tags Implementadas

### Tag `cef`

Arquivos com esta tag **só são compilados** quando você explicitamente pede:

```go
//go:build cef
// +build cef
```

**Arquivos afetados:**
- `internal/cef/cef.go` - Bindings CEF
- `main_cef.go` - Main da versão CEF

---

## 🔨 Como Compilar

### Versão Webview (Padrão)

```bash
# Não usa nenhuma tag - compila apenas Webview
go build -o build/bagus

# Ou
make build

# Ou
./scripts/build.sh
```

**O que acontece:**
- ✅ Compila `main.go`
- ✅ Compila `internal/ui/browser_production.go`
- ❌ **NÃO** compila `internal/cef/cef.go`
- ❌ **NÃO** compila `main_cef.go`

**Resultado:**
- Binário pequeno (~4MB)
- Sem dependências CEF
- Funciona imediatamente

### Versão CEF (Com Tag)

```bash
# Usa tag 'cef' - compila versão CEF
go build -tags cef -o build/bagus-cef main_cef.go

# Ou
make build-cef

# Ou
./scripts/build_cef.sh
```

**O que acontece:**
- ✅ Compila `main_cef.go`
- ✅ Compila `internal/cef/cef.go`
- ✅ Linka com CEF
- ❌ **NÃO** compila `main.go` (evita conflito)

**Resultado:**
- Binário maior (~15MB + 150MB de libs CEF)
- Requer CEF instalado
- 100% compatibilidade

---

## 🎯 Por que usar Build Tags?

### Problema sem Build Tags

```
❌ Sempre tenta compilar CEF
❌ Erro se CEF não estiver instalado
❌ Não pode ter dois main()
❌ Dependências desnecessárias
```

### Solução com Build Tags

```
✅ Compila apenas o que você pede
✅ Funciona sem CEF instalado
✅ Dois binários independentes
✅ Sem dependências extras
```

---

## 📊 Comparação

| Aspecto | Sem Tags | Com Tags |
|---------|----------|----------|
| **Compilação Webview** | ❌ Erro CEF | ✅ Funciona |
| **Compilação CEF** | ⚠️ Sempre tenta | ✅ Sob demanda |
| **Tamanho binário** | Grande | Otimizado |
| **Dependências** | Todas | Apenas necessárias |
| **Manutenção** | Difícil | Fácil |

---

## 🔍 Como Funciona

### 1. Arquivo sem tag (sempre compila)

```go
// main.go
package main

func main() {
    // Versão Webview
}
```

**Compila com:** `go build`

### 2. Arquivo com tag (compila sob demanda)

```go
//go:build cef
// +build cef

// main_cef.go
package main

func main() {
    // Versão CEF
}
```

**Compila com:** `go build -tags cef`

---

## 🛠️ Comandos Úteis

### Ver quais arquivos serão compilados

```bash
# Sem tags (Webview)
go list -f '{{.GoFiles}}' .

# Com tag CEF
go list -tags cef -f '{{.GoFiles}}' .
```

### Verificar build tags

```bash
# Ver todos os arquivos e suas tags
find . -name "*.go" -exec grep -l "//go:build" {} \;
```

### Compilar com múltiplas tags

```bash
# Se tivéssemos mais tags
go build -tags "cef,debug" -o build/bagus-cef-debug
```

---

## 📝 Sintaxe de Build Tags

### Formato Moderno (Go 1.17+)

```go
//go:build cef
// +build cef
```

### Operadores Lógicos

```go
// Compila se CEF OU debug
//go:build cef || debug

// Compila se CEF E debug
//go:build cef && debug

// Compila se NÃO for CEF
//go:build !cef
```

---

## 🚨 Erros Comuns

### Erro: "build constraints exclude all Go files"

**Causa:** Tentou importar pacote com build tag sem usar a tag

```go
// main.go (sem tag)
import "github.com/.../internal/cef" // ❌ Erro!
```

**Solução:** Use a tag ou não importe

```bash
go build -tags cef  # ✅ Funciona
```

### Erro: "main redeclared"

**Causa:** Dois arquivos com `func main()` sendo compilados

**Solução:** Use build tags para separar

```go
// main.go - sem tag
func main() { /* Webview */ }

// main_cef.go - com tag cef
//go:build cef
func main() { /* CEF */ }
```

---

## ✅ Benefícios no Bagus Browser

1. **Compilação Limpa**
   - Webview compila sem CEF instalado
   - CEF compila apenas quando necessário

2. **Binários Otimizados**
   - Webview: ~4MB (sem CEF)
   - CEF: ~15MB (com CEF)

3. **Desenvolvimento Fácil**
   - Desenvolva Webview sem instalar CEF
   - Teste CEF apenas quando necessário

4. **Distribuição Flexível**
   - Distribua versão leve (Webview)
   - Distribua versão completa (CEF)
   - Usuário escolhe

5. **Manutenção Simples**
   - Código separado logicamente
   - Sem #ifdef ou condicionais complexas
   - Go gerencia tudo

---

## 📖 Documentação Oficial

- [Build Constraints](https://pkg.go.dev/cmd/go#hdr-Build_constraints)
- [Go Build Tags](https://go.dev/wiki/BuildTags)

---

## 🎯 Resumo

**Build tags permitem:**
- ✅ Compilar Webview sem CEF instalado
- ✅ Compilar CEF apenas quando solicitado
- ✅ Dois binários independentes
- ✅ Código limpo e organizado
- ✅ Sem conflitos de `main()`

**Comandos:**
```bash
go build              # Webview (padrão)
go build -tags cef    # CEF (explícito)
```

**Pronto!** 🚀
