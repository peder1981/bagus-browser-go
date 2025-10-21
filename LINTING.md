# 🔍 Linting e Qualidade de Código

## Configuração do Linter

O projeto usa `golangci-lint` com configuração personalizada em `.golangci.yml`.

### Executar Linter

```bash
# Via Makefile
make lint

# Diretamente
golangci-lint run --timeout=5m
```

## ⚠️ Issues Conhecidos

O linter atualmente reporta alguns issues que precisam ser corrigidos:

### 1. Erros Não Verificados (errcheck)

**Arquivos afetados:**
- `cmd/bagus/main.go:21` - `defer lock.Unlock()`
- `internal/lockfile/lockfile.go` - Múltiplos erros não verificados
- `internal/security/blocker.go:58` - `defer file.Close()`
- `internal/ui/browser_final.go` - `b.w.Bind()` retornos não verificados

**Prioridade:** Alta  
**Impacto:** Potenciais erros silenciosos

**Solução sugerida:**
```go
// Antes
defer lock.Unlock()

// Depois
defer func() {
    if err := lock.Unlock(); err != nil {
        log.Printf("Erro ao liberar lock: %v", err)
    }
}()
```

### 2. Erros de Ortografia (misspell)

**Arquivos afetados:**
- `internal/browser/history.go:74` - "duplicatas" → "duplicates"
- `internal/ui/browser_final.go:73` - "controles" → "controls"

**Prioridade:** Baixa  
**Impacto:** Apenas comentários

**Solução sugerida:**
```go
// Antes
// Não adiciona duplicatas consecutivas

// Depois  
// Não adiciona duplicates consecutivas
// OU manter em português se for intencional
```

## 📋 Plano de Correção

### Fase 1: Erros Críticos (Alta Prioridade)
- [ ] Corrigir todos os `errcheck` em `internal/lockfile/lockfile.go`
- [ ] Corrigir `defer lock.Unlock()` em `cmd/bagus/main.go`
- [ ] Corrigir `defer file.Close()` em `internal/security/blocker.go`

### Fase 2: Erros Médios (Média Prioridade)
- [ ] Corrigir `b.w.Bind()` retornos não verificados
- [ ] Revisar outros erros de `errcheck`

### Fase 3: Melhorias (Baixa Prioridade)
- [ ] Corrigir erros de ortografia em comentários
- [ ] Adicionar mais linters (gosec, gocyclo, etc.)

## 🚀 Como Contribuir

### Antes de Fazer PR

1. Execute o linter:
```bash
make lint
```

2. Corrija todos os issues reportados no seu código

3. Execute testes:
```bash
make test
```

4. Formate o código:
```bash
make fmt
```

### Workflow Recomendado

```bash
# 1. Fazer mudanças
vim internal/meu_arquivo.go

# 2. Formatar
make fmt

# 3. Lint
make lint

# 4. Testar
make test

# 5. Commit
git add .
git commit -m "feat: minha feature"
```

## 🔧 Configuração do CI/CD

O GitHub Actions está configurado para executar o linter automaticamente.

**Arquivo:** `.github/workflows/lint.yml` (se existir)

### Ignorar Pacotes CGO

A configuração atual ignora erros de typecheck em pacotes que dependem de GTK/WebKit:
- `main.go`
- `cmd/bagus/main.go`
- `internal/ui/`

Isso é necessário porque o CI não tem GTK/WebKit instalado.

## 📚 Recursos

- [golangci-lint Documentation](https://golangci-lint.run/)
- [Enabled Linters](https://golangci-lint.run/usage/linters/)
- [Configuration](https://golangci-lint.run/usage/configuration/)

## 🎯 Meta de Qualidade

**Objetivo:** Zero issues do linter antes da v2.1.0

**Status Atual:**
- ❌ ~15 issues reportados
- ⚠️ Maioria são `errcheck` (erros não verificados)
- ✅ Configuração do linter funcionando

---

**Última Atualização:** 2025-10-21  
**Issues Pendentes:** ~15  
**Prioridade:** Média (não bloqueia release v2.0.0)
