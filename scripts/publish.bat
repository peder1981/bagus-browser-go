@echo off
REM Bagus Browser Go - Script de Publicacao no GitHub (Windows)
REM Uso: scripts\publish.bat [mensagem do commit]

setlocal enabledelayedexpansion

echo ========================================
echo  Bagus Browser Go - Publish Script
echo ========================================
echo.

REM Verificar se e um repositorio git
if not exist ".git" (
    echo [X] Erro: Nao e um repositorio Git
    echo Execute 'git init' primeiro
    exit /b 1
)

REM Mensagem de commit
set COMMIT_MSG=%*
if "%COMMIT_MSG%"=="" (
    for /f "tokens=1-3 delims=/ " %%a in ('date /t') do set DATE=%%c-%%b-%%a
    for /f "tokens=1-2 delims=: " %%a in ('time /t') do set TIME=%%a:%%b
    set COMMIT_MSG=Update: !DATE! !TIME!
)

echo === Preparando Publicacao ===
echo.

REM 1. Verificar qualidade do codigo
echo [*] Verificando qualidade do codigo...

REM Formatar
echo   - Formatando codigo...
go fmt .\... >nul 2>&1
echo   [OK] Codigo formatado

REM Testes
echo   - Executando testes...
go test .\... >nul 2>&1
if errorlevel 1 (
    echo   [!] Alguns testes falharam
    set /p CONTINUE="Deseja continuar? (s/N): "
    if /i not "!CONTINUE!"=="s" (
        echo [X] Publicacao cancelada
        exit /b 1
    )
) else (
    echo   [OK] Todos os testes passaram
)

REM Build
echo   - Compilando...
go build -o build\bagus.exe .\cmd\bagus >nul 2>&1
if errorlevel 1 (
    echo   [X] Build falhou
    exit /b 1
)
echo   [OK] Build bem-sucedido
echo.

REM 2. Git status
echo === Status do Git ===
git status --short
echo.

REM 3. Adicionar arquivos
echo [*] Adicionando arquivos...
git add .
echo [OK] Arquivos adicionados
echo.

REM 4. Commit
echo [*] Criando commit...
echo   Mensagem: "%COMMIT_MSG%"
git commit -m "%COMMIT_MSG%"
echo [OK] Commit criado
echo.

REM 5. Verificar remote
echo [*] Verificando remote...
git remote | findstr "origin" >nul
if errorlevel 1 (
    echo [!] Remote 'origin' nao configurado
    echo.
    echo Configure o remote com:
    echo git remote add origin git@github.com:peder1981/bagus-browser-go.git
    echo ou
    echo git remote add origin https://github.com/peder1981/bagus-browser-go.git
    exit /b 1
)

for /f "tokens=*" %%i in ('git remote get-url origin') do set REMOTE_URL=%%i
echo   Remote: !REMOTE_URL!
echo.

REM 6. Push
echo [*] Publicando no GitHub...
for /f "tokens=*" %%i in ('git branch --show-current') do set BRANCH=%%i
echo   Branch: !BRANCH!

git push origin !BRANCH!
if errorlevel 1 (
    echo [X] Push falhou
    echo.
    echo Se for o primeiro push, tente:
    echo git push -u origin !BRANCH!
    exit /b 1
)
echo [OK] Push bem-sucedido

echo.
echo ========================================
echo      Publicacao Concluida!
echo ========================================
echo.
echo Branch: !BRANCH!
echo Commit: %COMMIT_MSG%
echo.

endlocal
