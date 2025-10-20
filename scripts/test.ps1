# Bagus Browser Go - Script de Testes (PowerShell)
# Uso: .\scripts\test.ps1 [unit|coverage|lint|fmt|all]

param(
    [string]$TestType = "all"
)

$ErrorActionPreference = "Stop"

# Cores
function Write-ColorOutput($ForegroundColor) {
    $fc = $host.UI.RawUI.ForegroundColor
    $host.UI.RawUI.ForegroundColor = $ForegroundColor
    if ($args) {
        Write-Output $args
    }
    $host.UI.RawUI.ForegroundColor = $fc
}

# Banner
Write-ColorOutput Cyan "╔════════════════════════════════════════╗"
Write-ColorOutput Cyan "║   Bagus Browser Go - Test Script     ║"
Write-ColorOutput Cyan "╚════════════════════════════════════════╝"
Write-Output ""

# Função para testes unitários
function Run-UnitTests {
    Write-ColorOutput Yellow "► Executando testes unitários..."
    
    try {
        go test -v .\...
        Write-ColorOutput Green "✓ Todos os testes passaram"
    }
    catch {
        Write-ColorOutput Red "✗ Alguns testes falharam"
        exit 1
    }
}

# Função para coverage
function Run-Coverage {
    Write-ColorOutput Yellow "► Executando testes com coverage..."
    
    try {
        go test -v -coverprofile=coverage.out .\...
        Write-ColorOutput Green "✓ Testes concluídos"
        
        Write-ColorOutput Yellow "► Gerando relatório HTML..."
        go tool cover -html=coverage.out -o coverage.html
        Write-ColorOutput Green "✓ Relatório gerado: coverage.html"
        
        Write-Output ""
        Write-ColorOutput Cyan "═══ Coverage Summary ═══"
        $summary = go tool cover -func=coverage.out | Select-Object -Last 1
        Write-Output $summary
    }
    catch {
        Write-ColorOutput Red "✗ Alguns testes falharam"
        exit 1
    }
}

# Função para lint
function Run-Lint {
    Write-ColorOutput Yellow "► Executando linter..."
    
    if (-not (Get-Command golangci-lint -ErrorAction SilentlyContinue)) {
        Write-ColorOutput Yellow "golangci-lint não encontrado. Instalando..."
        go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
    }
    
    try {
        golangci-lint run .\...
        Write-ColorOutput Green "✓ Lint passou"
    }
    catch {
        Write-ColorOutput Red "✗ Problemas encontrados no lint"
        exit 1
    }
}

# Função para formatar
function Run-Fmt {
    Write-ColorOutput Yellow "► Formatando código..."
    go fmt .\...
    Write-ColorOutput Green "✓ Código formatado"
}

# Executar testes
switch ($TestType.ToLower()) {
    "unit" {
        Run-UnitTests
    }
    "coverage" {
        Run-Coverage
    }
    "lint" {
        Run-Lint
    }
    "fmt" {
        Run-Fmt
    }
    "all" {
        Write-ColorOutput Cyan "═══ Executando Todos os Testes ═══"
        Run-Fmt
        Run-Lint
        Run-UnitTests
        Run-Coverage
    }
    default {
        Write-ColorOutput Red "Tipo de teste inválido: $TestType"
        Write-Output "Uso: .\test.ps1 [unit|coverage|lint|fmt|all]"
        exit 1
    }
}

Write-Output ""
Write-ColorOutput Green "╔════════════════════════════════════════╗"
Write-ColorOutput Green "║      Testes Concluídos com Sucesso!   ║"
Write-ColorOutput Green "╚════════════════════════════════════════╝"
