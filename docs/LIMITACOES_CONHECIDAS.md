# ⚠️ Limitações Conhecidas - Bagus Browser

## 📋 Visão Geral

O Bagus Browser é construído com foco em **segurança e privacidade**, o que significa algumas limitações intencionais e outras técnicas.

---

## 🚫 Limitações Intencionais (Por Design)

### 1. Sem Suporte a Extensões

**Status:** ❌ Não suportado (intencional)

**Motivo:**
- Segurança: Menor superfície de ataque
- Privacidade: Sem telemetria de extensões
- Performance: Mais leve e rápido

**Impacto:**
- Gerenciadores de senha via extensão NÃO funcionam
- Ad blockers via extensão NÃO funcionam
- Qualquer extensão NÃO funciona

**Solução:**
- Use apps desktop (KeePassXC, Bitwarden, etc)
- Use gerenciadores com auto-type global
- Veja: `docs/GERENCIADORES_SENHA.md`

---

### 2. WebGL Bloqueado

**Status:** ❌ Bloqueado (intencional)

**Motivo:**
- Anti-fingerprinting
- Privacidade

**Impacto:**
- Alguns jogos web não funcionam
- Visualizações 3D podem não funcionar
- Google Maps pode ter problemas

**Solução:**
- Use apps nativos quando possível
- Para mapas: OpenStreetMap funciona melhor

---

### 3. WebAudio Bloqueado

**Status:** ❌ Bloqueado (intencional)

**Motivo:**
- Anti-fingerprinting
- Privacidade

**Impacto:**
- Alguns sites de música podem não funcionar
- Efeitos sonoros web podem falhar

**Solução:**
- Use apps nativos (Spotify, etc)
- YouTube funciona normalmente (usa HTML5 video)

---

### 4. Geolocalização Bloqueada

**Status:** ❌ Bloqueada (intencional)

**Motivo:**
- Privacidade

**Impacto:**
- Sites não podem detectar sua localização
- "Encontrar lojas próximas" não funciona automaticamente

**Solução:**
- Digite seu endereço manualmente
- Use apps nativos com GPS

---

### 5. Acesso a Câmera/Microfone Bloqueado

**Status:** ❌ Bloqueado (intencional)

**Motivo:**
- Privacidade
- Segurança

**Impacto:**
- Videochamadas web não funcionam
- Gravação de áudio web não funciona

**Solução:**
- Use apps nativos:
  - Zoom
  - Google Meet (app)
  - Discord (app)
  - Skype

---

## 🔧 Limitações Técnicas (WebKit2GTK)

### 1. Copy/Paste de Screenshots no WhatsApp Web

**Status:** ❌ Não funciona (limitação técnica)

**Causa:**
- WhatsApp Web usa API `FileSystemFileHandle`
- WebKit2GTK 4.0 não suporta essa API
- Erro: `ReferenceError: Can't find variable: FileSystemFileHandle`

**Sites Afetados:**
- ❌ WhatsApp Web
- ❌ Alguns sites modernos com APIs experimentais

**Sites que FUNCIONAM:**
- ✅ Gmail
- ✅ Google Docs
- ✅ Slack
- ✅ Discord
- ✅ Telegram Web
- ✅ Notion

**Soluções:**

#### Opção 1: Upload Manual
```
1. Tirar screenshot
2. Salvar arquivo
3. WhatsApp Web > 📎 (clipe)
4. Selecionar arquivo
5. Enviar
```

#### Opção 2: WhatsApp Desktop
```bash
sudo snap install whatsapp-for-linux
```

#### Opção 3: Usar Telegram Web
- Telegram Web suporta paste de screenshots
- https://web.telegram.org

---

### 2. Signal 'download-started' Inválido

**Status:** ⚠️ Warning (não afeta funcionalidade)

**Logs:**
```
GLib-GObject-WARNING: signal 'download-started' is invalid
```

**Causa:**
- WebKit2GTK mudou API de downloads
- Sinal antigo ainda sendo usado

**Impacto:**
- ❌ Nenhum! Downloads funcionam normalmente
- Apenas warning nos logs

**Solução:**
- Ignorar warning
- Será corrigido em versão futura

---

### 3. Fontes Bloqueadas (Google Docs)

**Status:** ⚠️ Warning (não afeta funcionalidade)

**Logs:**
```
CONSOLE SECURITY WARN [blocked] requested insecure content from filesystem:
```

**Causa:**
- Google Docs tenta carregar fontes via filesystem://
- Bloqueado por segurança

**Impacto:**
- ❌ Nenhum! Google Docs funciona normalmente
- Usa fontes alternativas

**Solução:**
- Ignorar warning
- Google Docs funciona perfeitamente

---

## 📊 Tabela de Compatibilidade

### Sites Populares

| Site | Status | Observações |
|------|--------|-------------|
| **Gmail** | ✅ Funciona | Paste de screenshots OK |
| **Google Docs** | ✅ Funciona | Paste de screenshots OK |
| **Google Sheets** | ✅ Funciona | Warnings de fontes (ignorar) |
| **GitHub** | ✅ Funciona | Totalmente compatível |
| **YouTube** | ✅ Funciona | Vídeos funcionam |
| **WhatsApp Web** | ⚠️ Parcial | Paste de screenshots NÃO funciona |
| **Telegram Web** | ✅ Funciona | Paste de screenshots OK |
| **Slack** | ✅ Funciona | Paste de screenshots OK |
| **Discord** | ✅ Funciona | Paste de screenshots OK |
| **Twitter/X** | ✅ Funciona | Totalmente compatível |
| **Reddit** | ✅ Funciona | Totalmente compatível |
| **DuckDuckGo** | ✅ Funciona | Motor de busca padrão |

### Funcionalidades

| Funcionalidade | Status | Observações |
|----------------|--------|-------------|
| **Navegação básica** | ✅ | Totalmente funcional |
| **Múltiplas abas** | ✅ | Até 100 abas |
| **Favoritos** | ✅ | Criptografados (AES-256) |
| **Downloads** | ✅ | Para ~/Downloads |
| **Impressão** | ✅ | PDF e impressoras |
| **Buscar na página** | ✅ | Ctrl+F |
| **Zoom** | ✅ | Ctrl+/- |
| **Sessões** | ✅ | Restaura abas |
| **Copy/Paste texto** | ✅ | Funciona em todos os sites |
| **Copy/Paste imagens** | ⚠️ | Depende do site |
| **Extensões** | ❌ | Não suportado (intencional) |
| **WebGL** | ❌ | Bloqueado (privacidade) |
| **WebAudio** | ❌ | Bloqueado (privacidade) |
| **Geolocalização** | ❌ | Bloqueada (privacidade) |
| **Câmera/Mic** | ❌ | Bloqueados (privacidade) |

---

## 🎯 Recomendações

### Para Melhor Experiência:

1. **Gerenciadores de Senha:**
   - Use KeePassXC (auto-type funciona perfeitamente)
   - Veja: `docs/GERENCIADORES_SENHA.md`

2. **WhatsApp:**
   - Use WhatsApp Desktop para paste de imagens
   - Ou use Telegram Web (funciona melhor)

3. **Videochamadas:**
   - Use apps nativos (Zoom, Meet, Discord)
   - Não use versões web

4. **Jogos/3D:**
   - Use apps nativos
   - Bagus Browser não é para jogos

5. **Música:**
   - YouTube funciona
   - Para streaming: use apps nativos

---

## 🔄 Futuras Melhorias

### Em Consideração:

1. **WebKit2GTK 4.1:**
   - Suporte a APIs mais modernas
   - Pode resolver problema do WhatsApp
   - Aguardando estabilidade

2. **Download Handler:**
   - Corrigir warning do signal
   - Melhorar API

3. **Opções de Privacidade:**
   - Permitir habilitar WebGL temporariamente
   - Permitir habilitar câmera por site
   - Controle granular

---

## ❓ FAQ

### P: Por que não suporta extensões?

**R:** Segurança e privacidade. Extensões:
- Podem ter vulnerabilidades
- Coletam telemetria
- Aumentam superfície de ataque
- Diminuem performance

### P: Como usar gerenciador de senha então?

**R:** Use KeePassXC com auto-type (Ctrl+Alt+A). Funciona perfeitamente!

### P: WhatsApp Web vai funcionar algum dia?

**R:** Talvez com WebKit2GTK 4.1, mas sem previsão. Use WhatsApp Desktop.

### P: Posso habilitar WebGL para um site específico?

**R:** Não no momento. Futuras versões podem ter essa opção.

### P: Por que tantas restrições?

**R:** Bagus Browser prioriza **privacidade** sobre conveniência. Se precisa de todas as funcionalidades, use Chrome/Firefox.

---

## 📚 Documentação Relacionada

- `docs/GERENCIADORES_SENHA.md` - Como usar gerenciadores de senha
- `docs/TESTES_v4.4.0.md` - Testes de funcionalidades
- `docs/development/MELHORIAS_v4.4.0.md` - Detalhes técnicos

---

**Atualizado em:** 22/10/2025  
**Versão:** 4.4.0  
**Status:** Documentação completa
