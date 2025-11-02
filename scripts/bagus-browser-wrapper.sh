#!/bin/bash
#
# 🌐 Wrapper para Bagus Browser com Dependências Embarcadas
# Detecta e usa dependências embarcadas ou do sistema
#

# Diretórios possíveis de dependências embarcadas
BUNDLE_PATHS=(
    "/opt/bagus-browser/bundle"
    "/usr/lib/bagus-browser/bundle"
    "/usr/local/lib/bagus-browser/bundle"
    "$(dirname "$0")/../lib/bagus-browser/bundle"
)

# Procurar por dependências embarcadas
for bundle_path in "${BUNDLE_PATHS[@]}"; do
    if [ -d "$bundle_path" ] && [ -d "$bundle_path/lib" ]; then
        echo "✅ Usando dependências embarcadas: $bundle_path"
        export LD_LIBRARY_PATH="$bundle_path/lib:$LD_LIBRARY_PATH"
        export PKG_CONFIG_PATH="$bundle_path/lib/pkgconfig:$PKG_CONFIG_PATH"
        export C_INCLUDE_PATH="$bundle_path/include:$C_INCLUDE_PATH"
        break
    fi
done

# Fallback: Tentar usar WebKit compilado com WebRTC
if [ -d "/opt/webkitgtk-webrtc" ]; then
    if [ -z "$LD_LIBRARY_PATH" ] || [[ ! "$LD_LIBRARY_PATH" =~ "webkitgtk-webrtc" ]]; then
        echo "✅ Usando WebKit compilado: /opt/webkitgtk-webrtc"
        export LD_LIBRARY_PATH="/opt/webkitgtk-webrtc/lib:$LD_LIBRARY_PATH"
        export PKG_CONFIG_PATH="/opt/webkitgtk-webrtc/lib/pkgconfig:$PKG_CONFIG_PATH"
    fi
fi

# Executar o browser
exec /usr/bin/bagus-browser "$@"
