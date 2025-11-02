#!/bin/bash

echo "🧪 Teste Google Meet - Bagus Browser v5.0.0"
echo ""
echo "Instruções:"
echo "1. O browser vai abrir"
echo "2. Navegue para https://meet.google.com"
echo "3. Aguarde carregar"
echo "4. Clique em 'Nova reunião'"
echo "5. Observe os logs abaixo"
echo ""
echo "═══════════════════════════════════════════════════════════"
echo ""

# Configurar variáveis de ambiente para WebKit compilado
if [ -d "/opt/webkitgtk-webrtc" ]; then
    export LD_LIBRARY_PATH="/opt/webkitgtk-webrtc/lib:$LD_LIBRARY_PATH"
    export PKG_CONFIG_PATH="/opt/webkitgtk-webrtc/lib/pkgconfig:$PKG_CONFIG_PATH"
fi

./build/bagus-browser
