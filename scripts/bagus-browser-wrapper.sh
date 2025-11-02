#!/bin/bash

# Wrapper para executar Bagus Browser com WebKit compilado com WebRTC

# Configurar variáveis de ambiente para WebKit compilado
if [ -d "/opt/webkitgtk-webrtc" ]; then
    export LD_LIBRARY_PATH="/opt/webkitgtk-webrtc/lib:$LD_LIBRARY_PATH"
    export PKG_CONFIG_PATH="/opt/webkitgtk-webrtc/lib/pkgconfig:$PKG_CONFIG_PATH"
fi

# Executar o browser
exec /usr/bin/bagus-browser "$@"
