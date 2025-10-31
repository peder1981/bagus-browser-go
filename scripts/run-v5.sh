#!/bin/bash

echo "🚀 Iniciando Bagus Browser v5.0.0 com WebRTC..."
echo ""

export LD_LIBRARY_PATH=/opt/webkitgtk-webrtc/lib:$LD_LIBRARY_PATH

./build/bagus-browser-v5 2>&1 | tee /tmp/bagus-v5-log.txt
