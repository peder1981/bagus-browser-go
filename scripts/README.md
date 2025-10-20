# Scripts de Automação

Scripts multiplataforma para build, teste, desenvolvimento e publicação do Bagus Browser Go.

## 📋 Scripts Disponíveis

### Publicação no GitHub

Scripts para publicar rapidamente no GitHub com validação automática.

#### Setup Inicial (Linux/macOS)

```bash
./scripts/setup-github.sh
```

**Funcionalidades:**
- Inicializa repositório Git (se necessário)
- Configura remote (SSH ou HTTPS)
- Cria commit inicial
- Faz primeiro push para GitHub

#### Publicação Rápida

**Linux/macOS (Bash)**
```bash
./scripts/publish.sh ["mensagem do commit"]
```

**Windows (Batch)**
```cmd
scripts\publish.bat [mensagem do commit]
```

**Windows (PowerShell)**
```powershell
.\scripts\publish.ps1 -Message "mensagem do commit"
```

**Funcionalidades:**
- Formata código automaticamente
- Executa lint (se disponível)
- Roda todos os testes
- Compila o projeto
- Cria commit
- Faz push para GitHub

**Exemplos:**
```bash
# Com mensagem personalizada
./scripts/publish.sh "feat: adiciona nova funcionalidade"

# Sem mensagem (usa timestamp)
./scripts/publish.sh
```

### Build Scripts

Compilam o projeto para diferentes plataformas.

#### Linux/macOS (Bash)

```bash
./scripts/build.sh [opção]
```

**Opções:**
- `current` - Compila para plataforma atual (padrão)
- `linux` - Compila para Linux (amd64, arm64)
- `windows` - Compila para Windows (amd64)
- `macos` - Compila para macOS (amd64, arm64)
- `all` - Compila para todas as plataformas

**Exemplos:**
```bash
# Build para plataforma atual
./scripts/build.sh

# Build para todas as plataformas
./scripts/build.sh all

# Build apenas para Linux
./scripts/build.sh linux
```

#### Windows (Batch)

```cmd
scripts\build.bat [opção]
```

**Opções:** Mesmas do script bash

**Exemplos:**
```cmd
REM Build para Windows
scripts\build.bat windows

REM Build para todas as plataformas
scripts\build.bat all
```

#### Windows (PowerShell)

```powershell
.\scripts\build.ps1 -Platform [opção]
```

**Opções:** Mesmas do script bash

**Exemplos:**
```powershell
# Build para Windows
.\scripts\build.ps1 -Platform windows

# Build para todas as plataformas
.\scripts\build.ps1 -Platform all
```

### Test Scripts

Executam testes, linting e formatação.

#### Linux/macOS (Bash)

```bash
./scripts/test.sh [opção]
```

**Opções:**
- `unit` - Executa apenas testes unitários
- `coverage` - Executa testes com coverage
- `lint` - Executa linter (golangci-lint)
- `fmt` - Formata código (gofmt)
- `all` - Executa tudo (padrão)

**Exemplos:**
```bash
# Executar todos os testes
./scripts/test.sh

# Apenas testes unitários
./scripts/test.sh unit

# Gerar coverage
./scripts/test.sh coverage
```

#### Windows (Batch)

```cmd
scripts\test.bat [opção]
```

**Opções:** Mesmas do script bash

**Exemplos:**
```cmd
REM Executar todos os testes
scripts\test.bat all

REM Apenas lint
scripts\test.bat lint
```

#### Windows (PowerShell)

```powershell
.\scripts\test.ps1 -TestType [opção]
```

**Opções:** Mesmas do script bash

**Exemplos:**
```powershell
# Executar todos os testes
.\scripts\test.ps1 -TestType all

# Apenas coverage
.\scripts\test.ps1 -TestType coverage
```

## 🚀 Uso Rápido

### Desenvolvimento Diário

```bash
# Linux/macOS
./scripts/test.sh fmt    # Formatar código
./scripts/test.sh lint   # Verificar qualidade
./scripts/test.sh unit   # Testar
./scripts/build.sh       # Compilar

# Windows (PowerShell)
.\scripts\test.ps1 -TestType fmt
.\scripts\test.ps1 -TestType lint
.\scripts\test.ps1 -TestType unit
.\scripts\build.ps1
```

### Antes de Commit

```bash
# Linux/macOS
./scripts/test.sh all

# Windows
.\scripts\test.ps1 -TestType all
```

### Build para Release

```bash
# Linux/macOS
./scripts/build.sh all

# Windows
.\scripts\build.ps1 -Platform all
```

## 📊 Output dos Scripts

### Build Script

```
╔════════════════════════════════════════╗
║   Bagus Browser Go - Build Script    ║
╚════════════════════════════════════════╝

► Instalando dependências...
✓ Dependências instaladas

═══ Build Todas as Plataformas ═══
► Compilando para linux/amd64...
✓ Build concluído: build/bagus-linux-amd64
► Compilando para linux/arm64...
✓ Build concluído: build/bagus-linux-arm64
...

╔════════════════════════════════════════╗
║      Build Concluído com Sucesso!     ║
╚════════════════════════════════════════╝
```

### Test Script

```
╔════════════════════════════════════════╗
║   Bagus Browser Go - Test Script     ║
╚════════════════════════════════════════╝

═══ Executando Todos os Testes ═══
► Formatando código...
✓ Código formatado
► Executando linter...
✓ Lint passou
► Executando testes unitários...
✓ Todos os testes passaram
► Executando testes com coverage...
✓ Testes concluídos
► Gerando relatório HTML...
✓ Relatório gerado: coverage.html

═══ Coverage Summary ═══
total: (statements) 85.2%

╔════════════════════════════════════════╗
║      Testes Concluídos com Sucesso!   ║
╚════════════════════════════════════════╝
```

## 🔧 Requisitos

### Para Build

- **Go**: 1.21 ou superior
- **Make**: Opcional (scripts funcionam independentemente)

### Para Testes

- **Go**: 1.21 ou superior
- **golangci-lint**: Instalado automaticamente se não encontrado

## 📝 Notas

### Permissões (Linux/macOS)

Os scripts bash precisam de permissão de execução:

```bash
chmod +x scripts/*.sh
```

### Execução no Windows

#### PowerShell Execution Policy

Se encontrar erro de política de execução:

```powershell
Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser
```

#### Executar do CMD

```cmd
powershell -ExecutionPolicy Bypass -File scripts\build.ps1 -Platform all
```

### Cross-Compilation

Os scripts suportam cross-compilation. Você pode compilar para qualquer plataforma de qualquer sistema operacional.

**Exemplo:** Compilar para Windows a partir do Linux:

```bash
./scripts/build.sh windows
```

## 🐛 Troubleshooting

### "go: command not found"

**Solução:** Instale Go ou adicione ao PATH

```bash
# Linux/macOS
export PATH=$PATH:/usr/local/go/bin

# Windows
setx PATH "%PATH%;C:\Go\bin"
```

### "golangci-lint: command not found"

**Solução:** O script instalará automaticamente, ou instale manualmente:

```bash
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

### "Permission denied" (Linux/macOS)

**Solução:** Torne o script executável

```bash
chmod +x scripts/build.sh
chmod +x scripts/test.sh
```

### Build falha com erro de memória

**Solução:** Compile plataformas individualmente

```bash
./scripts/build.sh linux
./scripts/build.sh windows
./scripts/build.sh macos
```

## 📚 Recursos Adicionais

- [Makefile](../Makefile) - Alternativa aos scripts
- [DEVELOPMENT.md](../docs/DEVELOPMENT.md) - Guia de desenvolvimento
- [CONTRIBUTING.md](../CONTRIBUTING.md) - Guia de contribuição

## 🤝 Contribuindo

Se você melhorar os scripts, por favor:

1. Teste em todas as plataformas (Linux, Windows, macOS)
2. Mantenha compatibilidade com versões anteriores
3. Atualize esta documentação
4. Abra um Pull Request

---

**Última Atualização:** 2024-10-20
