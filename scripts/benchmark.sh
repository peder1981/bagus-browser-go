#!/bin/bash
#
# 🔬 Bagus Browser - Script de Benchmark Automático
# Compara performance com outros navegadores
#

set -e

# Cores
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
CYAN='\033[0;36m'
NC='\033[0m'

# Configurações
ITERATIONS=5
SLEEP_TIME=2

echo -e "${CYAN}"
echo "╔════════════════════════════════════════╗"
echo "║   🔬 Bagus Browser Benchmark Suite    ║"
echo "╚════════════════════════════════════════╝"
echo -e "${NC}"
echo ""

# Verificar se bagus-browser está instalado
if ! command -v bagus-browser &> /dev/null; then
    echo -e "${RED}❌ bagus-browser não encontrado${NC}"
    echo "Instale primeiro: ./scripts/bagus install"
    exit 1
fi

# Função para medir tempo de inicialização
benchmark_startup() {
    local browser=$1
    local command=$2
    
    echo -e "${YELLOW}📊 Testando inicialização: $browser${NC}"
    
    local total=0
    for i in $(seq 1 $ITERATIONS); do
        echo -n "  Tentativa $i/$ITERATIONS... "
        
        # Medir tempo
        local start=$(date +%s%N)
        $command &> /dev/null &
        local pid=$!
        sleep $SLEEP_TIME
        kill $pid 2>/dev/null || true
        local end=$(date +%s%N)
        
        local elapsed=$(( (end - start) / 1000000 ))
        total=$(( total + elapsed ))
        
        echo "${elapsed}ms"
        
        # Limpar processos restantes
        killall -9 $(basename $command) 2>/dev/null || true
        sleep 1
    done
    
    local avg=$(( total / ITERATIONS ))
    echo -e "${GREEN}  ✅ Média: ${avg}ms${NC}"
    echo ""
    
    echo $avg
}

# Função para medir uso de RAM
benchmark_memory() {
    local browser=$1
    local command=$2
    
    echo -e "${YELLOW}💾 Testando uso de RAM: $browser${NC}"
    
    # Iniciar browser
    $command &> /dev/null &
    local pid=$!
    sleep 3
    
    # Medir RAM
    local mem=$(ps aux | grep -E "$browser|$pid" | grep -v grep | awk '{sum+=$6} END {print sum/1024}')
    
    echo -e "${GREEN}  ✅ RAM: ${mem} MB${NC}"
    
    # Fechar
    kill $pid 2>/dev/null || true
    killall -9 $(basename $command) 2>/dev/null || true
    sleep 1
    
    echo ""
    echo $mem
}

# Função para medir uso de CPU
benchmark_cpu() {
    local browser=$1
    local command=$2
    
    echo -e "${YELLOW}🖥️  Testando uso de CPU: $browser${NC}"
    
    # Iniciar browser
    $command &> /dev/null &
    local pid=$!
    sleep 2
    
    # Medir CPU por 10 segundos
    local cpu=$(top -b -n 10 -d 1 -p $pid 2>/dev/null | grep $pid | awk '{sum+=$9; count++} END {print sum/count}')
    
    echo -e "${GREEN}  ✅ CPU (idle): ${cpu}%${NC}"
    
    # Fechar
    kill $pid 2>/dev/null || true
    killall -9 $(basename $command) 2>/dev/null || true
    sleep 1
    
    echo ""
    echo $cpu
}

# Função para medir tamanho do binário
benchmark_size() {
    local browser=$1
    local path=$2
    
    echo -e "${YELLOW}📦 Testando tamanho: $browser${NC}"
    
    if [ -f "$path" ]; then
        local size=$(du -h "$path" | cut -f1)
        echo -e "${GREEN}  ✅ Tamanho: ${size}${NC}"
        echo ""
        echo $size
    else
        echo -e "${RED}  ❌ Binário não encontrado${NC}"
        echo ""
        echo "N/A"
    fi
}

# Resultados
declare -A STARTUP_TIMES
declare -A MEMORY_USAGE
declare -A CPU_USAGE
declare -A BINARY_SIZES

# Testar Bagus Browser
echo -e "${BLUE}═══════════════════════════════════════${NC}"
echo -e "${BLUE}  Testando: Bagus Browser${NC}"
echo -e "${BLUE}═══════════════════════════════════════${NC}"
echo ""

STARTUP_TIMES["Bagus"]=$(benchmark_startup "bagus-browser" "bagus-browser")
MEMORY_USAGE["Bagus"]=$(benchmark_memory "bagus-browser" "bagus-browser")
CPU_USAGE["Bagus"]=$(benchmark_cpu "bagus-browser" "bagus-browser")
BINARY_SIZES["Bagus"]=$(benchmark_size "bagus-browser" "/usr/bin/bagus-browser")

# Testar Firefox (se disponível)
if command -v firefox &> /dev/null; then
    echo -e "${BLUE}═══════════════════════════════════════${NC}"
    echo -e "${BLUE}  Testando: Firefox${NC}"
    echo -e "${BLUE}═══════════════════════════════════════${NC}"
    echo ""
    
    STARTUP_TIMES["Firefox"]=$(benchmark_startup "firefox" "firefox --new-window about:blank")
    MEMORY_USAGE["Firefox"]=$(benchmark_memory "firefox" "firefox --new-window about:blank")
    CPU_USAGE["Firefox"]=$(benchmark_cpu "firefox" "firefox --new-window about:blank")
    BINARY_SIZES["Firefox"]=$(benchmark_size "firefox" "/usr/bin/firefox")
fi

# Testar Chrome (se disponível)
if command -v google-chrome &> /dev/null; then
    echo -e "${BLUE}═══════════════════════════════════════${NC}"
    echo -e "${BLUE}  Testando: Google Chrome${NC}"
    echo -e "${BLUE}═══════════════════════════════════════${NC}"
    echo ""
    
    STARTUP_TIMES["Chrome"]=$(benchmark_startup "chrome" "google-chrome --new-window about:blank")
    MEMORY_USAGE["Chrome"]=$(benchmark_memory "chrome" "google-chrome --new-window about:blank")
    CPU_USAGE["Chrome"]=$(benchmark_cpu "chrome" "google-chrome --new-window about:blank")
    BINARY_SIZES["Chrome"]=$(benchmark_size "chrome" "/usr/bin/google-chrome")
fi

# Testar Epiphany (se disponível)
if command -v epiphany &> /dev/null; then
    echo -e "${BLUE}═══════════════════════════════════════${NC}"
    echo -e "${BLUE}  Testando: GNOME Web (Epiphany)${NC}"
    echo -e "${BLUE}═══════════════════════════════════════${NC}"
    echo ""
    
    STARTUP_TIMES["Epiphany"]=$(benchmark_startup "epiphany" "epiphany about:blank")
    MEMORY_USAGE["Epiphany"]=$(benchmark_memory "epiphany" "epiphany about:blank")
    CPU_USAGE["Epiphany"]=$(benchmark_cpu "epiphany" "epiphany about:blank")
    BINARY_SIZES["Epiphany"]=$(benchmark_size "epiphany" "/usr/bin/epiphany")
fi

# Exibir resultados
echo ""
echo -e "${GREEN}╔════════════════════════════════════════╗${NC}"
echo -e "${GREEN}║        📊 Resultados Finais            ║${NC}"
echo -e "${GREEN}╚════════════════════════════════════════╝${NC}"
echo ""

echo -e "${CYAN}🚀 Tempo de Inicialização (menor é melhor):${NC}"
for browser in "${!STARTUP_TIMES[@]}"; do
    printf "  %-15s %s ms\n" "$browser:" "${STARTUP_TIMES[$browser]}"
done
echo ""

echo -e "${CYAN}💾 Uso de RAM (menor é melhor):${NC}"
for browser in "${!MEMORY_USAGE[@]}"; do
    printf "  %-15s %s MB\n" "$browser:" "${MEMORY_USAGE[$browser]}"
done
echo ""

echo -e "${CYAN}🖥️  Uso de CPU em Idle (menor é melhor):${NC}"
for browser in "${!CPU_USAGE[@]}"; do
    printf "  %-15s %s%%\n" "$browser:" "${CPU_USAGE[$browser]}"
done
echo ""

echo -e "${CYAN}📦 Tamanho do Binário (menor é melhor):${NC}"
for browser in "${!BINARY_SIZES[@]}"; do
    printf "  %-15s %s\n" "$browser:" "${BINARY_SIZES[$browser]}"
done
echo ""

# Salvar resultados em arquivo
RESULTS_FILE="benchmark_results_$(date +%Y%m%d_%H%M%S).txt"
{
    echo "Bagus Browser - Benchmark Results"
    echo "Data: $(date)"
    echo "Sistema: $(uname -a)"
    echo ""
    echo "=== Tempo de Inicialização ==="
    for browser in "${!STARTUP_TIMES[@]}"; do
        echo "$browser: ${STARTUP_TIMES[$browser]} ms"
    done
    echo ""
    echo "=== Uso de RAM ==="
    for browser in "${!MEMORY_USAGE[@]}"; do
        echo "$browser: ${MEMORY_USAGE[$browser]} MB"
    done
    echo ""
    echo "=== Uso de CPU ==="
    for browser in "${!CPU_USAGE[@]}"; do
        echo "$browser: ${CPU_USAGE[$browser]}%"
    done
    echo ""
    echo "=== Tamanho do Binário ==="
    for browser in "${!BINARY_SIZES[@]}"; do
        echo "$browser: ${BINARY_SIZES[$browser]}"
    done
} > "$RESULTS_FILE"

echo -e "${GREEN}✅ Resultados salvos em: $RESULTS_FILE${NC}"
echo ""

# Determinar vencedor
echo -e "${YELLOW}🏆 Análise:${NC}"
echo ""

# Comparar com Bagus
bagus_startup=${STARTUP_TIMES["Bagus"]}
bagus_memory=${MEMORY_USAGE["Bagus"]}

for browser in "${!STARTUP_TIMES[@]}"; do
    if [ "$browser" != "Bagus" ]; then
        startup_ratio=$(echo "scale=1; ${STARTUP_TIMES[$browser]} / $bagus_startup" | bc)
        memory_ratio=$(echo "scale=1; ${MEMORY_USAGE[$browser]} / $bagus_memory" | bc)
        
        echo "  vs $browser:"
        echo "    - Bagus é ${startup_ratio}x mais rápido na inicialização"
        echo "    - Bagus usa ${memory_ratio}x menos RAM"
        echo ""
    fi
done

echo -e "${GREEN}╔════════════════════════════════════════╗${NC}"
echo -e "${GREEN}║     ✅ Benchmark Concluído!            ║${NC}"
echo -e "${GREEN}╚════════════════════════════════════════╝${NC}"
echo ""
