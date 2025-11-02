#!/bin/bash

# Script para monitorar compilação do WebKit

echo "🔍 Monitorando compilação do WebKit com WebRTC..."
echo ""

while true; do
    clear
    echo "═══════════════════════════════════════════════════════════"
    echo "📊 Status da Compilação do WebKit com WebRTC (GTK3)"
    echo "═══════════════════════════════════════════════════════════"
    echo ""
    
    # Verificar se processo está rodando
    if pgrep -f "build-webkit-webrtc.sh" > /dev/null; then
        echo "✅ Status: COMPILANDO"
    elif [ -d "/opt/webkitgtk-webrtc" ]; then
        echo "✅ Status: CONCLUÍDO COM SUCESSO"
        echo ""
        echo "📦 WebKit instalado em: /opt/webkitgtk-webrtc/"
        ls -lh /opt/webkitgtk-webrtc/lib/libwebkit* 2>/dev/null | head -3
        echo ""
        echo "Próximo passo: ./scripts/bagus build"
        break
    else
        echo "⏳ Status: AGUARDANDO OU ERRO"
    fi
    
    echo ""
    echo "📝 Últimas linhas do log:"
    echo "───────────────────────────────────────────────────────────"
    tail -10 /tmp/webkit-build-gtk3.log 2>/dev/null || echo "Log não encontrado"
    echo ""
    echo "💾 Espaço em disco:"
    df -h /tmp | tail -1
    echo ""
    echo "⏱️  Próxima atualização em 30 segundos..."
    echo "   (Pressione Ctrl+C para sair)"
    echo ""
    
    sleep 30
done
