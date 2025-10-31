#!/bin/bash

# Matar qualquer instância antiga
pkill -9 -f bagus-browser-v5 2>/dev/null

echo "🔨 Recompilando..."
go build -o build/bagus-browser-v5 ./cmd/bagus-browser-v5

if [ $? -ne 0 ]; then
    echo "❌ Erro na compilação!"
    exit 1
fi

echo "✅ Compilado com sucesso!"
echo ""
echo "🧪 Teste WebRTC Local"
echo ""
echo "Instruções:"
echo "1. O browser vai abrir"
echo "2. Digite na barra de endereços: file:///tmp/test-webrtc.html"
echo "3. Pressione Enter"
echo "4. Veja os resultados do teste"
echo ""
echo "═══════════════════════════════════════════════════════════"
echo ""

export LD_LIBRARY_PATH=/opt/webkitgtk-webrtc/lib:$LD_LIBRARY_PATH

# Abrir direto no arquivo de teste
./build/bagus-browser-v5
