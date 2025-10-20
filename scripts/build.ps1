# Bagus Browser Go - Script de Build (PowerShell)
# Uso: .\scripts\build.ps1 [windows|linux|macos|all]

param(
    [string]$Platform = "current"
)

$ErrorActionPreference = "Stop"

# Variáveis
$APP_NAME = "bagus"
$VERSION = "2.0.0-alpha"
$BUILD_DIR = "build"
$CMD_DIR = "cmd/bagus"

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
Write-ColorOutput Cyan "║   Bagus Browser Go - Build Script    ║"
Write-ColorOutput Cyan "╚════════════════════════════════════════╝"
Write-Output ""

# Função para build
function Build-Platform {
    param(
        [string]$OS,
        [string]$Arch,
        [string]$Output
    )
    
    Write-ColorOutput Yellow "► Compilando para $OS/$Arch..."
    
    $env:GOOS = $OS
    $env:GOARCH = $Arch
    
    $ldflags = "-s -w -X main.appVersion=$VERSION"
    
    try {
        go build -ldflags $ldflags -o "$BUILD_DIR/$Output" ".\$CMD_DIR"
        Write-ColorOutput Green "✓ Build concluído: $BUILD_DIR/$Output"
    }
    catch {
        Write-ColorOutput Red "✗ Erro ao compilar"
        exit 1
    }
}

# Criar diretório de build
if (-not (Test-Path $BUILD_DIR)) {
    New-Item -ItemType Directory -Path $BUILD_DIR | Out-Null
}

# Instalar dependências
Write-ColorOutput Yellow "► Instalando dependências..."
go mod download
go mod tidy
Write-ColorOutput Green "✓ Dependências instaladas"
Write-Output ""

# Build por plataforma
switch ($Platform.ToLower()) {
    "windows" {
        Write-ColorOutput Cyan "═══ Build Windows ═══"
        Build-Platform "windows" "amd64" "$APP_NAME-windows-amd64.exe"
    }
    "linux" {
        Write-ColorOutput Cyan "═══ Build Linux ═══"
        Build-Platform "linux" "amd64" "$APP_NAME-linux-amd64"
        Build-Platform "linux" "arm64" "$APP_NAME-linux-arm64"
    }
    "macos" {
        Write-ColorOutput Cyan "═══ Build macOS ═══"
        Build-Platform "darwin" "amd64" "$APP_NAME-darwin-amd64"
        Build-Platform "darwin" "arm64" "$APP_NAME-darwin-arm64"
    }
    "all" {
        Write-ColorOutput Cyan "═══ Build Todas as Plataformas ═══"
        Build-Platform "linux" "amd64" "$APP_NAME-linux-amd64"
        Build-Platform "linux" "arm64" "$APP_NAME-linux-arm64"
        Build-Platform "windows" "amd64" "$APP_NAME-windows-amd64.exe"
        Build-Platform "darwin" "amd64" "$APP_NAME-darwin-amd64"
        Build-Platform "darwin" "arm64" "$APP_NAME-darwin-arm64"
    }
    "current" {
        Write-ColorOutput Cyan "═══ Build Plataforma Atual ═══"
        $ldflags = "-s -w -X main.appVersion=$VERSION"
        go build -ldflags $ldflags -o "$BUILD_DIR/$APP_NAME.exe" ".\$CMD_DIR"
        Write-ColorOutput Green "✓ Build concluído: $BUILD_DIR/$APP_NAME.exe"
    }
    default {
        Write-ColorOutput Red "Plataforma inválida: $Platform"
        Write-Output "Uso: .\build.ps1 [windows|linux|macos|all|current]"
        exit 1
    }
}

Write-Output ""
Write-ColorOutput Green "╔════════════════════════════════════════╗"
Write-ColorOutput Green "║      Build Concluído com Sucesso!     ║"
Write-ColorOutput Green "╚════════════════════════════════════════╝"
