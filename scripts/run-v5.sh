#!/bin/bash

echo "🚀 Iniciando Bagus Browser v5.0.0..."
echo ""

# Usar WebKit compilado com WebRTC se disponível
if [ -d "/opt/webkitgtk-webrtc" ]; then
    export LD_LIBRARY_PATH="/opt/webkitgtk-webrtc/lib:$LD_LIBRARY_PATH"
    export PKG_CONFIG_PATH="/opt/webkitgtk-webrtc/lib/pkgconfig:$PKG_CONFIG_PATH"
    
    # Detectar qual versão foi compilada
    if [ -f "/opt/webkitgtk-webrtc/lib/pkgconfig/webkitgtk-6.0.pc" ]; then
        echo "🔧 Usando WebKit GTK4 compilado com WebRTC"
    elif [ -f "/opt/webkitgtk-webrtc/lib/pkgconfig/webkit2gtk-4.0.pc" ]; then
        echo "🔧 Usando WebKit GTK3 compilado com WebRTC"
    fi
fi

./build/bagus-browser 2>&1 | tee /tmp/bagus-v5-log.txt
