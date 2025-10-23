# ⚙️  Configurações - Bagus Browser v4.5.0

## 📋 Como Acessar

**Atalho:** `Ctrl+,` (Ctrl + vírgula)

Ou pelo menu (futuro)

---

## 🔒 Aba 1: Segurança

### Senha Mestre

Protege o acesso ao browser com senha.

**Funcionalidades:**
- ✅ Login obrigatório ao abrir o browser
- ✅ Máximo 3 tentativas
- ✅ Todos os dados criptografados (AES-256-GCM)
- ✅ Senha hash com bcrypt

**Como Usar:**

#### Definir Senha
1. Abrir Configurações (Ctrl+,)
2. Aba "Segurança"
3. Marcar "Exigir senha ao abrir o browser"
4. Digitar senha (2x para confirmar)
5. Clicar "Salvar"

#### Alterar Senha
1. Abrir Configurações
2. Clicar "Alterar Senha"
3. Digitar senha atual
4. Digitar nova senha (2x)
5. Salvar

#### Remover Senha
1. Abrir Configurações
2. Clicar "Remover Senha"
3. Confirmar
4. Salvar

**⚠️ Importante:**
- Não há recuperação de senha!
- Se esquecer, precisará deletar `~/.config/bagus-browser/config.enc`
- Isso apagará TODAS as configurações

---

## 🔄 Aba 2: Sessões

### Manter Usuário Logado

Controla se você permanece logado em sites ao fechar e reabrir o browser.

**Opções:**

#### ✅ Habilitado (Padrão)
- Permanece logado em Gmail, GitHub, etc
- Cookies salvos em disco
- Sessões restauradas automaticamente

**Exemplo:**
```
1. Fazer login no Gmail
2. Fechar browser
3. Reabrir browser
4. Gmail ainda logado ✅
```

#### ❌ Desabilitado
- Deslogado de todos os sites ao fechar
- Cookies apagados
- Máxima privacidade

**Exemplo:**
```
1. Fazer login no Gmail
2. Fechar browser
3. Reabrir browser
4. Gmail pede login novamente ❌
```

### Salvar Cookies Entre Sessões

Controla se cookies são salvos no disco.

**Habilitado:**
- Cookies em `~/.config/bagus-browser/cookies.sqlite`
- Necessário para manter login

**Desabilitado:**
- Cookies apenas em memória
- Apagados ao fechar browser

### Limpar Todos os Cookies

Botão para limpar cookies manualmente.

**Quando usar:**
- Deslogar de todos os sites
- Limpar dados de navegação
- Resolver problemas de login

**⚠️ Aviso:** Você será deslogado de TODOS os sites!

---

## ⚡ Aba 3: Performance

### Cache de Vídeos

Melhora reprodução de vídeos (YouTube, etc).

**Como Funciona:**
- Armazena partes do vídeo em disco
- Vídeos carregam mais rápido
- Menos consumo de banda
- Melhor experiência em conexões lentas

**Configurações:**

#### Habilitar Cache
- ✅ Habilitado por padrão
- ❌ Desabilitar para economizar espaço

#### Tamanho do Cache
- **Mínimo:** 100 MB
- **Padrão:** 500 MB
- **Máximo:** 5000 MB (5 GB)

**Recomendações:**
- **Conexão lenta:** 1000-2000 MB
- **Conexão rápida:** 500 MB
- **Pouco espaço:** 100-300 MB

#### Limpar Cache

Botão para limpar cache manualmente.

**Quando usar:**
- Liberar espaço em disco
- Vídeos não carregam
- Resolver problemas de reprodução

**Localização:** `~/.cache/bagus-browser/video/`

---

## 🕵️  Aba 4: Privacidade

### Configurações Avançadas

Controle fino sobre privacidade e segurança.

#### Bloquear Cookies de Terceiros
- ✅ Habilitado por padrão
- Bloqueia rastreadores
- Mantém funcionalidade de sites

#### Bloquear Geolocalização
- ✅ Habilitado por padrão
- Sites não podem detectar sua localização
- Digite endereço manualmente quando necessário

#### Bloquear Notificações
- ❌ Desabilitado por padrão
- Permite sites pedirem permissão
- Você controla por site

#### Bloquear Câmera/Microfone
- ✅ Habilitado por padrão
- Videochamadas web não funcionam
- Use apps nativos (Zoom, Meet, etc)

#### Bloquear WebGL
- ✅ Habilitado por padrão
- Anti-fingerprinting
- Alguns jogos/mapas podem não funcionar

#### Bloquear WebAudio
- ✅ Habilitado por padrão
- Anti-fingerprinting
- YouTube funciona normalmente

#### Do Not Track
- ✅ Habilitado por padrão
- Envia cabeçalho DNT
- Sites podem ignorar

---

## 💾 Onde São Salvos os Dados?

### Configurações
```
~/.config/bagus-browser/config.enc
```
- Criptografado (AES-256-GCM)
- Backup recomendado

### Cookies
```
~/.config/bagus-browser/cookies.sqlite
```
- SQLite database
- Apenas se "Salvar cookies" habilitado

### Cache de Vídeos
```
~/.cache/bagus-browser/video/
```
- Pode ser deletado a qualquer momento
- Recriado automaticamente

### Sessões
```
~/.config/bagus-browser/session.enc
```
- Abas abertas
- Criptografado

### Favoritos
```
~/.config/bagus-browser/bookmarks.enc
```
- Criptografado

---

## 🔐 Segurança dos Dados

### Criptografia

**Algoritmo:** AES-256-GCM

**Chave:** Derivada de:
- Hostname da máquina
- Username do usuário
- PBKDF2 (100,000 iterações)

**Arquivos Criptografados:**
- ✅ config.enc
- ✅ session.enc
- ✅ bookmarks.enc

**Arquivos NÃO Criptografados:**
- ❌ cookies.sqlite (SQLite nativo)
- ❌ Cache de vídeos

### Senha Mestre

**Hash:** bcrypt (cost 10)

**Armazenamento:**
- Hash em config.enc
- Senha NUNCA salva em texto plano

**Segurança:**
- ✅ Resistente a brute-force
- ✅ Salt aleatório (32 bytes)
- ✅ Impossível reverter hash

---

## 📊 Configurações Padrão

Ao instalar pela primeira vez:

```
Segurança:
  Senha mestre: Desabilitada

Sessões:
  Manter logado: Habilitado ✅
  Salvar cookies: Habilitado ✅

Performance:
  Cache vídeos: Habilitado ✅
  Tamanho cache: 500 MB

Privacidade:
  Third-party cookies: Bloqueado ✅
  Geolocalização: Bloqueada ✅
  Notificações: Permitidas ❌
  Câmera/Mic: Bloqueados ✅
  WebGL: Bloqueado ✅
  WebAudio: Bloqueado ✅
  Do Not Track: Habilitado ✅
```

---

## 🎯 Cenários de Uso

### Máxima Privacidade
```
✅ Senha mestre: Habilitada
❌ Manter logado: Desabilitado
❌ Salvar cookies: Desabilitado
✅ Todas as opções de privacidade: Habilitadas
```

**Resultado:**
- Nenhum dado salvo em disco
- Deslogado de tudo ao fechar
- Máxima proteção

### Uso Diário (Recomendado)
```
❌ Senha mestre: Desabilitada
✅ Manter logado: Habilitado
✅ Salvar cookies: Habilitado
✅ Cache vídeos: 500 MB
✅ Privacidade padrão
```

**Resultado:**
- Conveniência
- Boa privacidade
- Performance otimizada

### Computador Compartilhado
```
✅ Senha mestre: Habilitada
✅ Manter logado: Habilitado
✅ Salvar cookies: Habilitado
✅ Privacidade padrão
```

**Resultado:**
- Proteção com senha
- Dados criptografados
- Outros usuários não acessam

---

## ❓ FAQ

### P: Esqueci a senha mestre, o que fazer?

**R:** Não há recuperação. Você precisará:
```bash
rm ~/.config/bagus-browser/config.enc
```
Isso apagará TODAS as configurações.

### P: Posso sincronizar configurações entre computadores?

**R:** Sim, copie os arquivos:
```bash
~/.config/bagus-browser/config.enc
~/.config/bagus-browser/bookmarks.enc
```

**⚠️ Atenção:** A criptografia usa hostname+username, então pode não funcionar entre máquinas diferentes.

### P: Como fazer backup?

**R:**
```bash
# Backup completo
tar -czf bagus-backup.tar.gz ~/.config/bagus-browser/

# Restaurar
tar -xzf bagus-backup.tar.gz -C ~/
```

### P: Cache de vídeos ocupa muito espaço?

**R:** Máximo configurado (padrão 500 MB). O WebKit gerencia automaticamente, removendo vídeos antigos.

### P: Desabilitar cookies quebra sites?

**R:** Sim, muitos sites precisam de cookies para funcionar. Recomendamos manter habilitado.

### P: Posso ter perfis diferentes?

**R:** Não no momento. Futuras versões podem ter essa funcionalidade.

---

## 🔄 Aplicar Mudanças

**Algumas configurações exigem reiniciar o browser:**
- Senha mestre
- Persistência de cookies
- Cache de vídeos

**Outras aplicam imediatamente:**
- Configurações de privacidade
- Limpar cookies/cache

---

## 📝 Notas Técnicas

### Formato do Arquivo config.enc

```json
{
  "require_password": false,
  "password_hash": "",
  "password_salt": "",
  "persist_sessions": true,
  "persist_cookies": true,
  "video_cache_enabled": true,
  "video_cache_size": 500,
  "block_third_party_cookies": true,
  "block_geolocation": true,
  "block_notifications": false,
  "block_media": true,
  "block_webgl": true,
  "block_webaudio": true,
  "do_not_track": true
}
```

Criptografado com AES-256-GCM.

---

**Versão:** 4.5.0  
**Data:** 23/10/2025  
**Autor:** Bagus Browser Team
