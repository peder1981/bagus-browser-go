#!/bin/bash
set -e

echo "🎨 Gerando ícones em múltiplos tamanhos..."

# Verificar se inkscape ou imagemagick está instalado
if command -v inkscape &> /dev/null; then
    CONVERTER="inkscape"
    echo "✅ Usando Inkscape"
elif command -v convert &> /dev/null; then
    CONVERTER="imagemagick"
    echo "✅ Usando ImageMagick"
else
    echo "❌ Erro: Instale inkscape ou imagemagick"
    echo "   sudo apt-get install inkscape"
    echo "   ou"
    echo "   sudo apt-get install imagemagick"
    exit 1
fi

# Tamanhos necessários
SIZES=(16 22 24 32 48 64 128 256 512)

for size in "${SIZES[@]}"; do
    echo "  Gerando ${size}x${size}..."
    
    if [ "$CONVERTER" = "inkscape" ]; then
        inkscape bagus-browser.svg \
            --export-filename="bagus-browser-${size}.png" \
            --export-width=${size} \
            --export-height=${size} \
            2>/dev/null
    else
        convert bagus-browser.svg \
            -resize ${size}x${size} \
            bagus-browser-${size}.png
    fi
done

echo "✅ Ícones gerados com sucesso!"
ls -lh bagus-browser-*.png
