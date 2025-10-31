# 🎥 Google Meet - Limitações e Status

## ❌ Status Atual: NÃO SUPORTADO

O Google Meet **não funciona** no Bagus Browser devido a limitações do WebKit2GTK 4.0 no Ubuntu 22.04.

## 🔍 Problema Técnico

### Erro Identificado
```javascript
TypeError: undefined is not an object (evaluating 'a.RTCPeerConnection.prototype')
```

### Causa Raiz
O **WebKit2GTK 4.0** no Ubuntu 22.04 foi compilado **sem suporte completo a WebRTC**, especificamente:
- `RTCPeerConnection` não está disponível
- API WebRTC não é exposta para JavaScript
- Mesmo com `webkit_settings_set_enable_webrtc(TRUE)`, a API não existe

### Por que isso acontece?
O WebKit2GTK 4.0 em algumas distribuições Linux (incluindo Ubuntu 22.04) é compilado sem as bibliotecas WebRTC necessárias, tornando impossível usar aplicações de videoconferência que dependem dessa tecnologia.

## ✅ O que FUNCIONA no Bagus Browser

### Multimídia Geral
- ✅ **YouTube** - Vídeos funcionam perfeitamente
- ✅ **YouTube Music** - Streaming de áudio
- ✅ **Vimeo, Dailymotion** - Outros players de vídeo
- ✅ **MediaStream** - Câmera/microfone (para apps que não usam WebRTC)
- ✅ **WebGL** - Renderização 3D
- ✅ **WebAudio** - Processamento de áudio avançado

### Recursos Implementados
- ✅ Permissões de mídia (câmera/microfone)
- ✅ Aceleração de hardware
- ✅ Fullscreen
- ✅ Popups e janelas
- ✅ EncryptedMedia (DRM)

## ❌ O que NÃO funciona

### Aplicações que requerem WebRTC
- ❌ **Google Meet**
- ❌ **Microsoft Teams** (web)
- ❌ **Zoom** (web)
- ❌ **Discord** (chamadas de voz/vídeo)
- ❌ **Jitsi Meet**
- ❌ Qualquer app de videoconferência baseada em WebRTC

## 🔧 Possíveis Soluções

### Opção 1: Migrar para WebKitGTK 6.0 ⚠️
**Status**: Tentado, mas incompatível

WebKitGTK 6.0 tem suporte completo a WebRTC, mas:
- Usa GTK4 ao invés de GTK3
- API completamente diferente
- Requer reescrever grande parte do código
- Quebra compatibilidade com sistemas mais antigos

**Instalação**:
```bash
sudo apt install libwebkitgtk-6.0-dev
```

**Problemas encontrados**:
- Funções mudaram de nome
- Estruturas de dados diferentes
- Callbacks incompatíveis
- Requer migração completa GTK3 → GTK4

### Opção 2: Compilar WebKit2GTK com WebRTC 🔨
**Status**: Muito complexo

- Requer compilar WebKit do zero
- Várias horas de compilação
- Dependências complexas
- Não é prático para distribuição

### Opção 3: Usar aplicativos nativos ✅
**Status**: RECOMENDADO

Para videoconferências, use aplicativos nativos:
- **Google Meet**: Não tem app nativo, use Chrome/Chromium
- **Zoom**: `sudo apt install zoom`
- **Microsoft Teams**: `sudo snap install teams-for-linux`
- **Discord**: `sudo snap install discord`

## 📊 Comparação de Browsers

| Browser | WebRTC | Google Meet |
|---------|--------|-------------|
| Chrome/Chromium | ✅ | ✅ |
| Firefox | ✅ | ✅ |
| Epiphany (GNOME Web) | ❌ | ❌ |
| Bagus Browser | ❌ | ❌ |

**Nota**: Todos os browsers baseados em WebKit2GTK 4.0 no Ubuntu 22.04 têm a mesma limitação.

## 🎯 Recomendação

Para usar Google Meet:
1. Use **Google Chrome** ou **Chromium** para videoconferências
2. Use **Bagus Browser** para navegação geral, YouTube, etc.
3. Aguarde migração futura para WebKitGTK 6.0 (quando estável)

## 📝 Notas Técnicas

### Configurações Implementadas
```c
webkit_settings_set_enable_media_stream(settings, TRUE);
webkit_settings_set_enable_webrtc(settings, TRUE);  // Não funciona no WebKit2GTK 4.0
webkit_settings_set_enable_webgl(settings, TRUE);
webkit_settings_set_enable_webaudio(settings, TRUE);
```

### Verificação de Suporte
```bash
# Verificar versão do WebKit
pkg-config --modversion webkit2gtk-4.0

# Verificar se WebRTC está disponível
# (não há comando direto, apenas testando no browser)
```

## 🔮 Futuro

Possíveis melhorias:
1. Migração para WebKitGTK 6.0 quando GTK4 estiver mais maduro
2. Suporte a múltiplos backends (WebKit 4.0 + 6.0)
3. Detecção automática de capacidades do sistema

## 📚 Referências

- [WebKitGTK Bug Tracker](https://bugs.webkit.org/)
- [WebRTC Support in WebKit](https://webkit.org/blog/7763/a-closer-look-into-webrtc/)
- [Ubuntu WebKit2GTK Package](https://packages.ubuntu.com/jammy/libwebkit2gtk-4.0-37)
