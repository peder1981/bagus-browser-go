@echo off
REM Bagus Browser Go - Script de Build (Windows)
REM Uso: scripts\build.bat [windows|linux|macos|all]

setlocal enabledelayedexpansion

set APP_NAME=bagus
set VERSION=2.0.0-alpha
set BUILD_DIR=build
set CMD_DIR=cmd/bagus

echo ========================================
echo   Bagus Browser Go - Build Script
echo ========================================
echo.

REM Criar diretório de build
if not exist %BUILD_DIR% mkdir %BUILD_DIR%

REM Instalar dependências
echo [*] Instalando dependencias...
go mod download
go mod tidy
if errorlevel 1 (
    echo [X] Erro ao instalar dependencias
    exit /b 1
)
echo [OK] Dependencias instaladas
echo.

REM Determinar plataforma
set PLATFORM=%1
if "%PLATFORM%"=="" set PLATFORM=current

if "%PLATFORM%"=="windows" goto build_windows
if "%PLATFORM%"=="linux" goto build_linux
if "%PLATFORM%"=="macos" goto build_macos
if "%PLATFORM%"=="all" goto build_all
if "%PLATFORM%"=="current" goto build_current
goto invalid_platform

:build_windows
echo === Build Windows ===
echo [*] Compilando para Windows/amd64...
set GOOS=windows
set GOARCH=amd64
go build -ldflags "-s -w -X main.appVersion=%VERSION%" -o %BUILD_DIR%\%APP_NAME%-windows-amd64.exe .\%CMD_DIR%
if errorlevel 1 goto error
echo [OK] Build concluido: %BUILD_DIR%\%APP_NAME%-windows-amd64.exe
goto end

:build_linux
echo === Build Linux ===
echo [*] Compilando para Linux/amd64...
set GOOS=linux
set GOARCH=amd64
go build -ldflags "-s -w -X main.appVersion=%VERSION%" -o %BUILD_DIR%\%APP_NAME%-linux-amd64 .\%CMD_DIR%
if errorlevel 1 goto error
echo [OK] Build concluido: %BUILD_DIR%\%APP_NAME%-linux-amd64

echo [*] Compilando para Linux/arm64...
set GOARCH=arm64
go build -ldflags "-s -w -X main.appVersion=%VERSION%" -o %BUILD_DIR%\%APP_NAME%-linux-arm64 .\%CMD_DIR%
if errorlevel 1 goto error
echo [OK] Build concluido: %BUILD_DIR%\%APP_NAME%-linux-arm64
goto end

:build_macos
echo === Build macOS ===
echo [*] Compilando para macOS/amd64...
set GOOS=darwin
set GOARCH=amd64
go build -ldflags "-s -w -X main.appVersion=%VERSION%" -o %BUILD_DIR%\%APP_NAME%-darwin-amd64 .\%CMD_DIR%
if errorlevel 1 goto error
echo [OK] Build concluido: %BUILD_DIR%\%APP_NAME%-darwin-amd64

echo [*] Compilando para macOS/arm64...
set GOARCH=arm64
go build -ldflags "-s -w -X main.appVersion=%VERSION%" -o %BUILD_DIR%\%APP_NAME%-darwin-arm64 .\%CMD_DIR%
if errorlevel 1 goto error
echo [OK] Build concluido: %BUILD_DIR%\%APP_NAME%-darwin-arm64
goto end

:build_all
echo === Build Todas as Plataformas ===
call :build_windows
call :build_linux
call :build_macos
goto end

:build_current
echo === Build Plataforma Atual ===
go build -ldflags "-s -w -X main.appVersion=%VERSION%" -o %BUILD_DIR%\%APP_NAME%.exe .\%CMD_DIR%
if errorlevel 1 goto error
echo [OK] Build concluido: %BUILD_DIR%\%APP_NAME%.exe
goto end

:invalid_platform
echo [X] Plataforma invalida: %PLATFORM%
echo Uso: %0 [windows^|linux^|macos^|all^|current]
exit /b 1

:error
echo [X] Erro ao compilar
exit /b 1

:end
echo.
echo ========================================
echo      Build Concluido com Sucesso!
echo ========================================
endlocal
