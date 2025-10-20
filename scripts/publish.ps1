# Bagus Browser Go - Script de Publicação no GitHub (PowerShell)
# Uso: .\scripts\publish.ps1 [-Message "mensagem do commit"]

param(
    [string]$Message = ""
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
Write-ColorOutput Cyan "║  Bagus Browser Go - Publish Script   ║"
Write-ColorOutput Cyan "╚════════════════════════════════════════╝"
Write-Output ""

# Verificar se é um repositório git
if (-not (Test-Path ".git")) {
    Write-ColorOutput Red "✗ Erro: Não é um repositório Git"
    Write-Output "Execute 'git init' primeiro"
    exit 1
}

# Mensagem de commit
if ($Message -eq "") {
    $timestamp = Get-Date -Format "yyyy-MM-dd HH:mm"
    $Message = "Update: $timestamp"
}

Write-ColorOutput Blue "═══ Preparando Publicação ═══"
Write-Output ""

# 1. Verificar qualidade do código
Write-ColorOutput Yellow "► Verificando qualidade do código..."

# Formatar
Write-Output "  • Formatando código..."
go fmt .\... | Out-Null
Write-ColorOutput Green "    ✓ Código formatado"

# Lint (opcional)
if (Get-Command golangci-lint -ErrorAction SilentlyContinue) {
    Write-Output "  • Executando lint..."
    try {
        golangci-lint run .\... | Out-Null
        Write-ColorOutput Green "    ✓ Lint passou"
    }
    catch {
        Write-ColorOutput Yellow "    ⚠ Lint encontrou problemas (não bloqueante)"
    }
}

# Testes
Write-Output "  • Executando testes..."
try {
    go test .\... | Out-Null
    Write-ColorOutput Green "    ✓ Todos os testes passaram"
}
catch {
    Write-ColorOutput Red "    ✗ Alguns testes falharam"
    $continue = Read-Host "Deseja continuar mesmo assim? (s/N)"
    if ($continue -ne "s" -and $continue -ne "S") {
        Write-ColorOutput Red "Publicação cancelada"
        exit 1
    }
}

# Build
Write-Output "  • Compilando..."
try {
    go build -o build\bagus.exe .\cmd\bagus | Out-Null
    Write-ColorOutput Green "    ✓ Build bem-sucedido"
}
catch {
    Write-ColorOutput Red "    ✗ Build falhou"
    exit 1
}

Write-Output ""

# 2. Git status
Write-ColorOutput Blue "═══ Status do Git ═══"
git status --short
Write-Output ""

# 3. Adicionar arquivos
Write-ColorOutput Yellow "► Adicionando arquivos..."
git add .
Write-ColorOutput Green "✓ Arquivos adicionados"
Write-Output ""

# 4. Commit
Write-ColorOutput Yellow "► Criando commit..."
Write-ColorOutput Cyan "  Mensagem: `"$Message`""
git commit -m $Message
Write-ColorOutput Green "✓ Commit criado"
Write-Output ""

# 5. Verificar remote
Write-ColorOutput Yellow "► Verificando remote..."
$remotes = git remote
if ($remotes -notcontains "origin") {
    Write-ColorOutput Yellow "⚠ Remote 'origin' não configurado"
    Write-Output ""
    Write-Output "Configure o remote com:"
    Write-ColorOutput Cyan "git remote add origin git@github.com:peder1981/bagus-browser-go.git"
    Write-Output "ou"
    Write-ColorOutput Cyan "git remote add origin https://github.com/peder1981/bagus-browser-go.git"
    exit 1
}

$remoteUrl = git remote get-url origin
Write-ColorOutput Cyan "  Remote: $remoteUrl"
Write-Output ""

# 6. Push
Write-ColorOutput Yellow "► Publicando no GitHub..."
$branch = git branch --show-current
Write-ColorOutput Cyan "  Branch: $branch"

try {
    git push origin $branch
    Write-ColorOutput Green "✓ Push bem-sucedido"
}
catch {
    Write-ColorOutput Red "✗ Push falhou"
    Write-Output ""
    Write-Output "Se for o primeiro push, tente:"
    Write-ColorOutput Cyan "git push -u origin $branch"
    exit 1
}

Write-Output ""
Write-ColorOutput Green "╔════════════════════════════════════════╗"
Write-ColorOutput Green "║     Publicação Concluída! 🚀          ║"
Write-ColorOutput Green "╚════════════════════════════════════════╝"
Write-Output ""
Write-ColorOutput Cyan "Branch: $branch"
Write-ColorOutput Cyan "Commit: $Message"
Write-Output ""
