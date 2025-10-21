# 🔒 Política de Privacidade - Bagus Browser

## 🎯 Compromisso com Privacidade

**Bagus Browser é 100% privado. ZERO telemetria. ZERO rastreamento.**

## ✅ O que NÃO coletamos

- ❌ **Nenhum dado de navegação** - Histórico, URLs, páginas visitadas
- ❌ **Nenhum dado pessoal** - Nome, email, localização
- ❌ **Nenhuma telemetria** - Métricas de uso, crashes, performance
- ❌ **Nenhum analytics** - Google Analytics, Mixpanel, etc.
- ❌ **Nenhuma conexão externa** - Sem servidores próprios
- ❌ **Nenhum rastreamento** - Cookies, fingerprinting, etc.

## 🔐 O que armazenamos (LOCALMENTE)

### Versão Webview e CEF

Todos os dados são armazenados **APENAS no seu computador**, criptografados:

1. **Histórico de Navegação** (Opcional)
   - Armazenado em: `~/.bagus/history.encrypted`
   - Criptografia: AES-256-GCM
   - Chave: Derivada do seu usuário (PBKDF2-SHA256, 100k iterações)
   - **Você pode desabilitar**: Configuração `history.enabled = false`

2. **Configurações**
   - Armazenado em: `~/.bagus/config.json`
   - Conteúdo: Preferências do navegador (tema, motor de busca, etc.)
   - **Sem dados sensíveis**

3. **Lista de Bloqueio** (Opcional)
   - Armazenado em: `~/.bagus/ad_hosts_block.txt`
   - Conteúdo: Domínios bloqueados (ads, malware)
   - **Você fornece a lista**

4. **Chave de Criptografia**
   - Armazenado em: `~/.bagus/.encryption_key`
   - Permissões: 0600 (apenas você)
   - **Única por usuário**

## 🌐 Conexões de Rede

### O que o Bagus Browser conecta:

1. **Sites que você visita** - Obviamente
2. **Motor de busca padrão** - DuckDuckGo (se você buscar algo)
3. **NADA MAIS** - Zero conexões para nossos servidores

### O que o Bagus Browser NÃO conecta:

- ❌ Servidores de telemetria
- ❌ Servidores de analytics
- ❌ Servidores de atualização automática
- ❌ Servidores de sincronização
- ❌ APIs externas (exceto sites que você visita)

## 🔍 Verificação de Código

### Como verificar que não há telemetria:

```bash
# Buscar por URLs suspeitas
grep -r "http" internal/ | grep -v "localhost" | grep -v "example.com"

# Buscar por telemetria
grep -ri "telemetry\|analytics\|tracking" .

# Buscar por Google Analytics
grep -ri "google-analytics\|ga.js\|gtag" .

# Buscar por conexões externas
grep -ri "http.Post\|http.Get" internal/
```

### Resultado esperado:
**NENHUMA conexão externa além dos sites que você visita.**

## 🛡️ Garantias Técnicas

### Versão Webview

```go
// internal/ui/browser_production.go
// NÃO HÁ:
// - http.Post() para servidores externos
// - http.Get() para APIs de telemetria
// - Conexões WebSocket para tracking
// - Beacons de analytics
```

### Versão CEF

```cpp
// cef/cef_browser.cpp
CefSettings settings;
settings.no_sandbox = true;
settings.windowless_rendering_enabled = false;
settings.multi_threaded_message_loop = false;

// NÃO configuramos:
// - user_agent com tracking
// - remote_debugging_port (exceto se você habilitar)
// - background_color com fingerprinting
```

## 📊 Comparação com Outros Navegadores

| Navegador | Telemetria | Rastreamento | Sincronização | Open Source |
|-----------|------------|--------------|---------------|-------------|
| **Bagus Browser** | ❌ ZERO | ❌ ZERO | ❌ Local | ✅ 100% |
| Chrome | ✅ Sim | ✅ Sim | ✅ Google | ⚠️ Parcial |
| Firefox | ⚠️ Opcional | ⚠️ Opcional | ⚠️ Mozilla | ✅ Sim |
| Brave | ⚠️ Mínima | ❌ Não | ⚠️ Opcional | ✅ Sim |
| Edge | ✅ Sim | ✅ Sim | ✅ Microsoft | ❌ Não |

## 🔐 Criptografia

### Histórico de Navegação

```go
// internal/security/encryption.go

// Algoritmo: AES-256-GCM
// - Chave: 256 bits (32 bytes)
// - Nonce: 96 bits (12 bytes) - único por operação
// - Tag de autenticação: 128 bits (16 bytes)

// Derivação de chave: PBKDF2-SHA256
// - Iterações: 100.000
// - Salt: 32 bytes aleatórios
// - Output: 32 bytes (256 bits)
```

**Resultado:** Mesmo que alguém acesse seu computador, não consegue ler seu histórico sem sua senha.

## 🚫 O que desabilitamos no CEF

```cpp
// Desabilitado por padrão:
settings.remote_debugging_port = 0;  // Sem debug remoto
settings.log_severity = LOGSEVERITY_DISABLE;  // Sem logs
settings.uncaught_exception_stack_size = 0;  // Sem stack traces

// Não habilitamos:
// - Crash reporting
// - Usage statistics
// - Update checks
// - Safe browsing (você pode habilitar localmente)
```

## 📝 Auditoria de Código

### Arquivos a verificar:

1. **`internal/ui/browser_production.go`** - Navegador webview
   - ✅ Sem http.Post/Get externos
   - ✅ Sem websockets
   - ✅ Sem beacons

2. **`internal/cef/cef.go`** - Bindings CEF
   - ✅ Sem telemetria
   - ✅ Sem analytics

3. **`cef/cef_browser.cpp`** - Browser CEF
   - ✅ Sem remote debugging
   - ✅ Sem crash reporting
   - ✅ Sem update checks

4. **`internal/security/encryption.go`** - Criptografia
   - ✅ AES-256-GCM
   - ✅ PBKDF2-SHA256
   - ✅ Sem backdoors

## 🔓 Código Aberto

**100% do código é open source:**
- Você pode auditar cada linha
- Você pode compilar do zero
- Você pode verificar binários

```bash
# Compile você mesmo
git clone https://github.com/peder1981/bagus-browser-go
cd bagus-browser-go
go build -o bagus

# Verifique o binário
strings bagus | grep -i "telemetry\|analytics\|tracking"
# Resultado: NADA
```

## 🛠️ Configurações de Privacidade

### Desabilitar Histórico

```json
// ~/.bagus/config.json
{
  "history": {
    "enabled": false  // Desabilita completamente
  }
}
```

### Desabilitar Cache

```json
{
  "cache": {
    "enabled": false
  }
}
```

### Modo Privado (Planejado)

```bash
bagus --private
# Nenhum dado salvo em disco
```

## 📞 Contato

Encontrou telemetria? **Reporte imediatamente:**
- GitHub Issues: https://github.com/peder1981/bagus-browser-go/issues
- Email: security@bagus-browser.dev (planejado)

**Oferecemos recompensa para quem encontrar telemetria não documentada.**

## ✅ Checklist de Privacidade

- [x] Zero telemetria
- [x] Zero rastreamento
- [x] Zero conexões externas (exceto sites visitados)
- [x] Histórico criptografado (AES-256-GCM)
- [x] Dados armazenados localmente
- [x] Código 100% open source
- [x] Auditável
- [x] Compilável do zero
- [ ] Modo privado (planejado)
- [ ] VPN integrada (planejado)
- [ ] Tor integration (planejado)

## 🎯 Resumo

**Bagus Browser respeita sua privacidade:**

1. ✅ **ZERO telemetria** - Nenhum dado enviado
2. ✅ **ZERO rastreamento** - Nenhum tracking
3. ✅ **Dados locais** - Tudo no seu computador
4. ✅ **Criptografia forte** - AES-256-GCM
5. ✅ **Open source** - Auditável
6. ✅ **Sem backdoors** - Código limpo

**Sua privacidade é nossa prioridade #1.** 🔒

---

*Última atualização: 2025-10-20*  
*Versão: 2.0.0*
