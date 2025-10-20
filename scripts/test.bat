@echo off
REM Bagus Browser Go - Script de Testes (Windows)
REM Uso: scripts\test.bat [unit|coverage|all]

setlocal enabledelayedexpansion

echo ========================================
echo   Bagus Browser Go - Test Script
echo ========================================
echo.

set TEST_TYPE=%1
if "%TEST_TYPE%"=="" set TEST_TYPE=all

if "%TEST_TYPE%"=="unit" goto run_unit
if "%TEST_TYPE%"=="coverage" goto run_coverage
if "%TEST_TYPE%"=="lint" goto run_lint
if "%TEST_TYPE%"=="fmt" goto run_fmt
if "%TEST_TYPE%"=="all" goto run_all
goto invalid_type

:run_unit
echo [*] Executando testes unitarios...
go test -v .\...
if errorlevel 1 goto error
echo [OK] Todos os testes passaram
goto end

:run_coverage
echo [*] Executando testes com coverage...
go test -v -coverprofile=coverage.out .\...
if errorlevel 1 goto error
echo [OK] Testes concluidos

echo [*] Gerando relatorio HTML...
go tool cover -html=coverage.out -o coverage.html
echo [OK] Relatorio gerado: coverage.html

echo.
echo === Coverage Summary ===
go tool cover -func=coverage.out | findstr /C:"total:"
goto end

:run_lint
echo [*] Executando linter...
where golangci-lint >nul 2>&1
if errorlevel 1 (
    echo [*] golangci-lint nao encontrado. Instalando...
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
)
golangci-lint run .\...
if errorlevel 1 goto error
echo [OK] Lint passou
goto end

:run_fmt
echo [*] Formatando codigo...
go fmt .\...
echo [OK] Codigo formatado
goto end

:run_all
echo === Executando Todos os Testes ===
call :run_fmt
call :run_lint
call :run_unit
call :run_coverage
goto end

:invalid_type
echo [X] Tipo de teste invalido: %TEST_TYPE%
echo Uso: %0 [unit^|coverage^|lint^|fmt^|all]
exit /b 1

:error
echo [X] Alguns testes falharam
exit /b 1

:end
echo.
echo ========================================
echo    Testes Concluidos com Sucesso!
echo ========================================
endlocal
