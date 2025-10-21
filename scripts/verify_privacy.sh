#!/bin/bash
# Script para verificar que não há telemetria no código

set -e

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

echo -e "${BLUE}╔════════════════════════════════════════╗${NC}"
echo -e "${BLUE}║   Verificação de Privacidade          ║${NC}"
echo -e "${BLUE}╚════════════════════════════════════════╝${NC}"
echo ""

ISSUES=0

# Função para verificar
check() {
    local pattern=$1
    local description=$2
    
    echo -e "${YELLOW}► Verificando: ${description}${NC}"
    
    result=$(grep -ri "$pattern" internal/ cef/ 2>/dev/null | \
        grep -v "Binary file" | \
        grep -v ".git" | \
        grep -v "// Não" | \
        grep -v "// Desabilita" | \
        grep -v "// GARANTIA" | \
        grep -v "// PRIVACIDADE" || true)
    
    if [ -z "$result" ]; then
        echo -e "${GREEN}✓ OK - Nenhum resultado encontrado${NC}"
    else
        echo -e "${RED}✗ ALERTA - Encontrado:${NC}"
        echo "$result"
        ISSUES=$((ISSUES + 1))
    fi
    echo ""
}

# Verificações

check "google-analytics\|ga\.js\|gtag" "Google Analytics"
check "mixpanel\|amplitude\|segment\.io" "Analytics de terceiros"
check "telemetry\|usage.statistics" "Telemetria explícita"
check "crash.report\|sentry\.io\|bugsnag" "Crash reporting"
check "track\|beacon\|pixel" "Tracking pixels"
check "fingerprint" "Fingerprinting"

# Verifica conexões HTTP suspeitas
echo -e "${YELLOW}► Verificando conexões HTTP externas${NC}"

# Busca por http.Post e http.Get
http_calls=$(grep -rn "http\.Post\|http\.Get" internal/ 2>/dev/null | grep -v "Binary file" | grep -v ".git" || true)

if [ -z "$http_calls" ]; then
    echo -e "${GREEN}✓ OK - Nenhuma conexão HTTP externa encontrada${NC}"
else
    echo -e "${YELLOW}⚠ Conexões HTTP encontradas (verificar manualmente):${NC}"
    echo "$http_calls"
    echo -e "${BLUE}Nota: Conexões para sites visitados são esperadas${NC}"
fi
echo ""

# Verifica URLs hardcoded suspeitas
echo -e "${YELLOW}► Verificando URLs hardcoded${NC}"

urls=$(grep -rn "https\?://" internal/ cef/ 2>/dev/null | \
    grep -v "Binary file" | \
    grep -v ".git" | \
    grep -v "duckduckgo.com" | \
    grep -v "example.com" | \
    grep -v "localhost" | \
    grep -v "127.0.0.1" | \
    grep -v "github.com/peder1981" || true)

if [ -z "$urls" ]; then
    echo -e "${GREEN}✓ OK - Apenas URLs esperadas${NC}"
else
    echo -e "${YELLOW}⚠ URLs encontradas (verificar manualmente):${NC}"
    echo "$urls"
fi
echo ""

# Verifica CEF settings
if [ -f "cef/cef_browser.cpp" ]; then
    echo -e "${YELLOW}► Verificando configurações CEF${NC}"
    
    if grep -q "remote_debugging_port = 0" cef/cef_browser.cpp; then
        echo -e "${GREEN}✓ Remote debugging desabilitado${NC}"
    else
        echo -e "${RED}✗ Remote debugging pode estar habilitado${NC}"
        ISSUES=$((ISSUES + 1))
    fi
    
    if grep -q "log_severity = LOGSEVERITY_DISABLE" cef/cef_browser.cpp; then
        echo -e "${GREEN}✓ Logs desabilitados${NC}"
    else
        echo -e "${YELLOW}⚠ Logs podem estar habilitados${NC}"
    fi
    
    if grep -q "uncaught_exception_stack_size = 0" cef/cef_browser.cpp; then
        echo -e "${GREEN}✓ Stack traces desabilitados${NC}"
    else
        echo -e "${YELLOW}⚠ Stack traces podem estar habilitados${NC}"
    fi
    echo ""
fi

# Verifica binário (se existir)
if [ -f "build/bagus" ]; then
    echo -e "${YELLOW}► Verificando binário${NC}"
    
    suspicious=$(strings build/bagus | grep -i "telemetry\|analytics\|tracking\|google-analytics" || true)
    
    if [ -z "$suspicious" ]; then
        echo -e "${GREEN}✓ Binário limpo${NC}"
    else
        echo -e "${RED}✗ Strings suspeitas encontradas no binário:${NC}"
        echo "$suspicious"
        ISSUES=$((ISSUES + 1))
    fi
    echo ""
fi

# Resultado final
echo -e "${BLUE}╔════════════════════════════════════════╗${NC}"
if [ $ISSUES -eq 0 ]; then
    echo -e "${GREEN}║   ✓ VERIFICAÇÃO PASSOU                ║${NC}"
    echo -e "${GREEN}║   Zero telemetria detectada           ║${NC}"
else
    echo -e "${RED}║   ✗ VERIFICAÇÃO FALHOU                ║${NC}"
    echo -e "${RED}║   $ISSUES problema(s) encontrado(s)      ║${NC}"
fi
echo -e "${BLUE}╚════════════════════════════════════════╝${NC}"
echo ""

if [ $ISSUES -eq 0 ]; then
    echo -e "${GREEN}Bagus Browser está livre de telemetria! 🔒${NC}"
    exit 0
else
    echo -e "${RED}Por favor, revise os problemas acima.${NC}"
    exit 1
fi
