# Guia de Contribuição

Obrigado por considerar contribuir com o Bagus Browser Go! 🎉

## 📋 Índice

- [Código de Conduta](#código-de-conduta)
- [Como Contribuir](#como-contribuir)
- [Configuração do Ambiente](#configuração-do-ambiente)
- [Processo de Desenvolvimento](#processo-de-desenvolvimento)
- [Padrões de Código](#padrões-de-código)
- [Testes](#testes)
- [Commits e Pull Requests](#commits-e-pull-requests)
- [Reportar Bugs](#reportar-bugs)
- [Sugerir Melhorias](#sugerir-melhorias)

## Código de Conduta

Este projeto adere a um código de conduta. Ao participar, você concorda em manter um ambiente respeitoso e inclusivo.

### Nossos Valores

- **Respeito**: Trate todos com respeito e empatia
- **Inclusão**: Seja acolhedor com contribuidores de todos os níveis
- **Colaboração**: Trabalhe junto para melhorar o projeto
- **Qualidade**: Mantenha altos padrões de código e documentação

## Como Contribuir

### Tipos de Contribuição

Aceitamos diversos tipos de contribuição:

- 🐛 **Correção de Bugs**: Corrija bugs existentes
- ✨ **Novas Features**: Implemente novas funcionalidades
- 📝 **Documentação**: Melhore ou traduza documentação
- 🧪 **Testes**: Adicione ou melhore testes
- 🎨 **UI/UX**: Melhore a interface e experiência
- 🔒 **Segurança**: Reporte ou corrija vulnerabilidades
- 🌍 **Tradução**: Traduza para outros idiomas

## Configuração do Ambiente

### Pré-requisitos

- **Go**: 1.21 ou superior
- **Git**: Para controle de versão
- **Make**: Para automação de builds
- **golangci-lint**: Para linting (instalado automaticamente)

### Setup Inicial

```bash
# 1. Fork o repositório no GitHub

# 2. Clone seu fork
git clone https://github.com/SEU_USUARIO/bagus-browser-go.git
cd bagus-browser-go

# 3. Adicione o upstream
git remote add upstream https://github.com/peder1981/bagus-browser-go.git

# 4. Instale dependências
make deps

# 5. Verifique se tudo funciona
make test
make build
```

### Estrutura do Projeto

```
bagus-browser-go/
├── cmd/bagus/          # Entry point
├── internal/           # Código privado
│   ├── browser/        # Engine do navegador
│   ├── crypto/         # Criptografia
│   ├── security/       # Segurança
│   ├── ui/             # Interface
│   └── storage/        # Armazenamento
├── pkg/                # Código público
├── tests/              # Testes
├── docs/               # Documentação
└── scripts/            # Scripts de automação
```

## Processo de Desenvolvimento

### Workflow

1. **Crie uma Issue**: Descreva o que pretende fazer
2. **Aguarde Feedback**: Espere aprovação/discussão
3. **Crie uma Branch**: `git checkout -b feature/minha-feature`
4. **Desenvolva**: Implemente sua mudança
5. **Teste**: Execute testes e lint
6. **Commit**: Faça commits semânticos
7. **Push**: Envie para seu fork
8. **Pull Request**: Abra um PR para o upstream

### Branches

- `main`: Branch principal (produção)
- `develop`: Branch de desenvolvimento
- `feature/*`: Novas funcionalidades
- `fix/*`: Correções de bugs
- `docs/*`: Atualizações de documentação
- `refactor/*`: Refatorações
- `test/*`: Adição de testes

### Exemplo de Workflow

```bash
# Atualizar main
git checkout main
git pull upstream main

# Criar branch de feature
git checkout -b feature/nova-funcionalidade

# Fazer mudanças
# ... editar arquivos ...

# Testar
make test
make lint

# Commit
git add .
git commit -m "feat: adiciona nova funcionalidade X"

# Push
git push origin feature/nova-funcionalidade

# Abrir PR no GitHub
```

## Padrões de Código

### Estilo Go

Seguimos os padrões oficiais do Go:

- **gofmt**: Formatação automática
- **golangci-lint**: Linting rigoroso
- **Effective Go**: Práticas recomendadas

### Convenções

#### Nomenclatura

```go
// Pacotes: lowercase, sem underscores
package browser

// Funções exportadas: PascalCase
func NewEngine() *Engine {}

// Funções privadas: camelCase
func validateURL(url string) error {}

// Constantes: PascalCase ou SCREAMING_SNAKE_CASE
const MaxTabs = 100
const DEFAULT_TIMEOUT = 30

// Variáveis: camelCase
var currentTab *Tab
```

#### Comentários

```go
// Package browser implementa o motor do navegador.
// 
// Este pacote fornece funcionalidades core para navegação,
// gerenciamento de abas e integração com o motor de renderização.
package browser

// Engine representa o motor principal do navegador.
// Coordena todos os componentes e gerencia o ciclo de vida.
type Engine struct {
    tabs []*Tab
}

// NewEngine cria uma nova instância do motor.
// Retorna erro se a inicialização falhar.
func NewEngine() (*Engine, error) {
    // Implementação
}
```

#### Tratamento de Erros

```go
// ✅ BOM: Retornar erros, não panic
func LoadPage(url string) error {
    if !isValid(url) {
        return fmt.Errorf("URL inválida: %s", url)
    }
    return nil
}

// ❌ RUIM: Usar panic para erros esperados
func LoadPage(url string) {
    if !isValid(url) {
        panic("URL inválida")
    }
}

// ✅ BOM: Wrap errors com contexto
if err := loadResource(url); err != nil {
    return fmt.Errorf("falha ao carregar recurso %s: %w", url, err)
}
```

#### Estruturas

```go
// ✅ BOM: Documentar campos importantes
type Config struct {
    // MaxTabs define o número máximo de abas abertas
    MaxTabs int `yaml:"max_tabs"`
    
    // Timeout em segundos para requisições
    Timeout int `yaml:"timeout"`
}

// ✅ BOM: Usar construtores
func NewConfig() *Config {
    return &Config{
        MaxTabs: 10,
        Timeout: 30,
    }
}
```

### Organização de Imports

```go
import (
    // 1. Standard library
    "context"
    "fmt"
    "time"
    
    // 2. External packages
    "github.com/webview/webview"
    
    // 3. Internal packages
    "github.com/peder1981/bagus-browser-go/internal/browser"
    "github.com/peder1981/bagus-browser-go/pkg/utils"
)
```

## Testes

### Estrutura de Testes

```go
// internal/browser/engine_test.go
package browser

import (
    "testing"
)

func TestNewEngine(t *testing.T) {
    engine, err := NewEngine()
    if err != nil {
        t.Fatalf("NewEngine() erro = %v", err)
    }
    if engine == nil {
        t.Error("NewEngine() retornou nil")
    }
}

func TestEngine_OpenTab(t *testing.T) {
    tests := []struct {
        name    string
        url     string
        wantErr bool
    }{
        {"URL válida", "https://example.com", false},
        {"URL inválida", "invalid", true},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            engine := &Engine{}
            err := engine.OpenTab(tt.url)
            if (err != nil) != tt.wantErr {
                t.Errorf("OpenTab() erro = %v, wantErr %v", err, tt.wantErr)
            }
        })
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
go test -v ./internal/browser -run TestNewEngine

# Benchmark
go test -bench=. ./...
```

### Cobertura de Testes

- **Mínimo**: 70% de cobertura
- **Recomendado**: 80%+
- **Crítico**: 90%+ para código de segurança

## Commits e Pull Requests

### Mensagens de Commit

Usamos [Conventional Commits](https://www.conventionalcommits.org/):

```
<tipo>(<escopo>): <descrição>

[corpo opcional]

[rodapé opcional]
```

#### Tipos

- `feat`: Nova funcionalidade
- `fix`: Correção de bug
- `docs`: Documentação
- `style`: Formatação (não afeta código)
- `refactor`: Refatoração
- `test`: Adicionar/modificar testes
- `chore`: Tarefas de manutenção
- `perf`: Melhoria de performance
- `ci`: CI/CD
- `build`: Sistema de build

#### Exemplos

```bash
# Feature
git commit -m "feat(browser): adiciona suporte a múltiplas abas"

# Bug fix
git commit -m "fix(security): corrige validação de URL"

# Documentação
git commit -m "docs(readme): atualiza instruções de instalação"

# Breaking change
git commit -m "feat(api)!: muda assinatura de NewEngine

BREAKING CHANGE: NewEngine agora requer parâmetro de config"
```

### Pull Requests

#### Template de PR

```markdown
## Descrição
Breve descrição das mudanças.

## Tipo de Mudança
- [ ] Bug fix
- [ ] Nova feature
- [ ] Breaking change
- [ ] Documentação

## Checklist
- [ ] Código segue os padrões do projeto
- [ ] Testes adicionados/atualizados
- [ ] Documentação atualizada
- [ ] Lint passa sem erros
- [ ] Todos os testes passam

## Testes
Como testar as mudanças.

## Screenshots (se aplicável)
```

#### Processo de Review

1. **CI passa**: Todos os checks devem passar
2. **Review**: Pelo menos 1 aprovação necessária
3. **Discussão**: Responda a comentários
4. **Atualização**: Faça mudanças solicitadas
5. **Merge**: Mantenedor fará o merge

## Reportar Bugs

### Template de Bug Report

```markdown
**Descrição do Bug**
Descrição clara e concisa do bug.

**Passos para Reproduzir**
1. Vá para '...'
2. Clique em '...'
3. Veja o erro

**Comportamento Esperado**
O que deveria acontecer.

**Comportamento Atual**
O que está acontecendo.

**Screenshots**
Se aplicável.

**Ambiente**
- OS: [ex: Ubuntu 22.04]
- Versão: [ex: 2.0.0-alpha]
- Go Version: [ex: 1.21.0]

**Logs**
```
Cole logs relevantes aqui
```

**Contexto Adicional**
Qualquer outra informação relevante.
```

## Sugerir Melhorias

### Template de Feature Request

```markdown
**Problema a Resolver**
Descrição clara do problema.

**Solução Proposta**
Como você sugere resolver.

**Alternativas Consideradas**
Outras soluções que você considerou.

**Contexto Adicional**
Informações adicionais, mockups, etc.
```

## Recursos Adicionais

### Documentação

- [Arquitetura](docs/ARCHITECTURE.md)
- [Guia de Desenvolvimento](docs/DEVELOPMENT.md)
- [API Reference](docs/API.md)

### Comunicação

- **Issues**: Para bugs e features
- **Discussions**: Para perguntas e ideias
- **Email**: Para questões privadas

### Ferramentas Úteis

- [Go Playground](https://go.dev/play/)
- [golangci-lint](https://golangci-lint.run/)
- [Effective Go](https://go.dev/doc/effective_go)

## Agradecimentos

Obrigado por contribuir! Cada contribuição, por menor que seja, é valiosa. 🙏

---

**Dúvidas?** Abra uma issue ou discussion!
