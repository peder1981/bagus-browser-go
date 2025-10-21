# 🧪 Guia de Testes

## 📊 Visão Geral

O Bagus Browser possui uma suíte de testes abrangente com **33.9% de cobertura** nos pacotes testáveis.

### Cobertura por Pacote

| Pacote | Coverage | Status |
|--------|----------|--------|
| `internal/storage` | 78.8% | ✅ Excelente |
| `pkg/utils` | 82.5% | ✅ Excelente |
| `internal/security` | 39.9% | ✅ Bom |
| `internal/browser` | 19.6% | ⚠️ Aceitável |
| `internal/config` | 0.0% | ⚠️ Sem testes |
| `internal/lockfile` | 0.0% | ⚠️ Sem testes |
| `internal/ui` | N/A | ⚠️ Requer CGO |

## 🚀 Executando Testes

### Testes Rápidos (Sem CGO)

Executa testes em todos os pacotes **exceto** aqueles que dependem de GTK/WebKit:

```bash
make test
```

**Pacotes testados:**
- `internal/browser`
- `internal/config`
- `internal/lockfile`
- `internal/security`
- `internal/storage`
- `pkg/utils`

**Pacotes pulados:**
- `main` (requer webview)
- `cmd/bagus` (requer webview)
- `internal/ui` (requer webview)

### Testes com Coverage

Gera relatório HTML de cobertura:

```bash
make test-coverage
```

**Saída:**
- `coverage.out` - Dados de cobertura
- `coverage.html` - Relatório visual (abra no browser)

### Testes Completos (Com CGO)

Executa **TODOS** os testes, incluindo UI:

```bash
make test-all
```

**Requer:**
```bash
sudo apt install libwebkit2gtk-4.0-dev libgtk-3-dev
```

## 📝 Estrutura de Testes

### Testes Unitários

```
internal/
├── browser/
│   ├── engine_test.go      # Testes do engine de abas
│   └── tab_test.go         # Testes de abas individuais
├── security/
│   └── encryption_test.go  # Testes de criptografia
├── storage/
│   └── manager_test.go     # Testes de gerenciamento
└── pkg/utils/
    └── utils_test.go       # Testes de utilitários
```

### Convenções

1. **Nomenclatura:** `*_test.go`
2. **Funções:** `Test<Nome>(t *testing.T)`
3. **Subtestes:** `t.Run("caso", func(t *testing.T) {...})`
4. **Table-driven:** Usar quando múltiplos casos

### Exemplo de Teste

```go
func TestMinhaFuncao(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected string
    }{
        {"caso 1", "input1", "output1"},
        {"caso 2", "input2", "output2"},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := MinhaFuncao(tt.input)
            if result != tt.expected {
                t.Errorf("got %v, want %v", result, tt.expected)
            }
        })
    }
}
```

## 🎯 Comandos Úteis

### Testar Pacote Específico

```bash
go test -v ./internal/browser/...
```

### Testar com Verbose

```bash
go test -v ./...
```

### Coverage de Pacote Específico

```bash
go test -coverprofile=cover.out ./internal/security/...
go tool cover -html=cover.out
```

### Executar Teste Específico

```bash
go test -v -run TestNomeDaFuncao ./internal/browser/...
```

### Benchmark

```bash
go test -bench=. ./internal/security/...
```

## 🔍 Análise de Coverage

### Ver Coverage no Terminal

```bash
go tool cover -func=coverage.out
```

### Ver Coverage HTML

```bash
go tool cover -html=coverage.out
```

### Coverage por Linha

```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html
# Abra coverage.html no browser
```

## ⚠️ Limitações

### Pacotes Sem Testes

**`internal/config`**
- Motivo: Funções simples de I/O
- Prioridade: Baixa

**`internal/lockfile`**
- Motivo: Depende de syscalls do sistema
- Prioridade: Média

**`internal/ui`**
- Motivo: Requer GTK/WebKit (CGO)
- Prioridade: Baixa (testes manuais)

### Testes de UI

Testes de interface gráfica são **complexos** e requerem:
- Display server (X11/Wayland)
- GTK/WebKit instalados
- Ambiente gráfico

**Alternativa:** Testes manuais e E2E

## 🚀 Melhorando Coverage

### Prioridades

1. **Alta:** `internal/browser` (19.6% → 60%+)
2. **Média:** `internal/config` (0% → 50%+)
3. **Média:** `internal/lockfile` (0% → 40%+)
4. **Baixa:** `internal/security` (39.9% → 60%+)

### Como Adicionar Testes

1. Crie arquivo `*_test.go` no mesmo pacote
2. Importe `testing`
3. Escreva funções `Test*`
4. Execute `go test -v`

### Exemplo Rápido

```bash
# Criar teste
cat > internal/config/config_test.go << 'EOF'
package config

import "testing"

func TestDefault(t *testing.T) {
    cfg := Default()
    if cfg.Default.URL == "" {
        t.Error("URL padrão não pode ser vazia")
    }
}
EOF

# Executar
go test -v ./internal/config/...
```

## 📚 Recursos

- [Go Testing](https://golang.org/pkg/testing/)
- [Table Driven Tests](https://github.com/golang/go/wiki/TableDrivenTests)
- [Go Coverage](https://blog.golang.org/cover)
- [Testing Best Practices](https://github.com/golang/go/wiki/CodeReviewComments#testing)

## 🤝 Contribuindo

Ao adicionar código:

1. ✅ Escreva testes para novas funcionalidades
2. ✅ Mantenha coverage acima de 30%
3. ✅ Execute `make test` antes de commit
4. ✅ Documente casos de teste complexos

---

**Última Atualização:** 2025-10-21  
**Coverage Total:** 33.9%  
**Testes Passando:** ✅ 100%
