# 🔐 Sistema OAuth2 com Cache - GitHub

## ✨ Funcionalidades

### 🔑 Autenticação Inteligente
- **Cache criptografado** de tokens (AES-256-CBC)
- **Validação automática** de tokens
- **Múltiplas fontes** de autenticação
- **Renovação automática** quando expirado

### 🔒 Segurança
- Tokens criptografados com AES-256-CBC
- Chave derivada do hostname (única por máquina)
- Permissões 600 nos arquivos de cache
- Validação antes de usar

### 📦 Fontes de Token (Prioridade)
1. **GITHUB_TOKEN** (variável de ambiente)
2. **Cache criptografado** (~/.config/bagus-browser/)
3. **gh CLI** (GitHub CLI autenticado)
4. **Input manual** (solicita se necessário)

---

## 🚀 Uso Automático

### No release.sh
```bash
./release.sh
```

**O script automaticamente:**
1. ✅ Verifica GITHUB_TOKEN
2. ✅ Verifica cache
3. ✅ Verifica gh CLI
4. ✅ Solicita token se necessário
5. ✅ Cacheia para próxima vez

---

## 🔧 Uso Manual

### Obter Token
```bash
./scripts/github-auth.sh get
```

### Ver Status
```bash
./scripts/github-auth.sh status
```

**Mostra:**
- Status do GITHUB_TOKEN
- Status do cache
- Status do gh CLI
- Usuário autenticado

### Limpar Cache
```bash
./scripts/github-auth.sh clear
```

---

## 📝 Exemplos

### Exemplo 1: Primeira Vez (Sem Token)
```bash
$ ./release.sh

🚀 Bagus Browser - GitHub Release Script
=========================================

🔍 Verificando cache de token...
⚠️  Nenhum token válido encontrado
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Por favor, configure um token do GitHub:

Opção 1: Criar token
  1. Ir para: https://github.com/settings/tokens
  2. Generate new token (classic)
  3. Marcar: repo (Full control)
  4. Copiar token

Opção 2: Usar gh CLI
  gh auth login

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Deseja inserir um token agora? (s/n): s
Cole o token do GitHub: ghp_xxxxx
✅ Token válido! (usuário: peder1981)
💾 Token cacheado com segurança
✅ Token obtido com sucesso
```

### Exemplo 2: Com Cache
```bash
$ ./release.sh

🚀 Bagus Browser - GitHub Release Script
=========================================

🔍 Verificando cache de token...
✅ Token em cache válido (usuário: peder1981)
✅ Token obtido com sucesso
Versão: v4.1.0
```

### Exemplo 3: Com GITHUB_TOKEN
```bash
$ export GITHUB_TOKEN=ghp_xxxxx
$ ./release.sh

🚀 Bagus Browser - GitHub Release Script
=========================================

🔑 Usando GITHUB_TOKEN da variável de ambiente
✅ Token válido
✅ Token obtido com sucesso
```

### Exemplo 4: Com gh CLI
```bash
$ gh auth login
$ ./release.sh

🚀 Bagus Browser - GitHub Release Script
=========================================

🔍 Verificando cache de token...
🔍 Verificando gh CLI...
✅ Token do gh CLI válido (usuário: peder1981)
💾 Token cacheado para uso futuro
✅ Token obtido com sucesso
```

---

## 🔐 Segurança

### Onde o Token é Armazenado
```
~/.config/bagus-browser/github_token.cache.enc
```

### Como é Criptografado
- **Algoritmo:** AES-256-CBC
- **Derivação:** PBKDF2
- **Chave:** SHA-256 do hostname
- **Permissões:** 600 (apenas dono)

### Validação
Antes de usar, o token é validado com:
```bash
curl -H "Authorization: token $TOKEN" \
     https://api.github.com/user
```

---

## 🔄 Fluxo de Autenticação

```
┌─────────────────┐
│  release.sh     │
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│ github-auth.sh  │
└────────┬────────┘
         │
         ▼
    ┌────────┐
    │ Token? │
    └───┬────┘
        │
        ├─► GITHUB_TOKEN? ──► Validar ──► ✅ Usar
        │
        ├─► Cache? ──────────► Validar ──► ✅ Usar
        │                          │
        │                          ▼
        │                       ❌ Expirado
        │                          │
        ├─► gh CLI? ─────────► Validar ──► ✅ Usar + Cachear
        │                          │
        │                          ▼
        │                       ❌ Não autenticado
        │
        └─► Solicitar ───────► Validar ──► ✅ Usar + Cachear
                                   │
                                   ▼
                                ❌ Inválido
```

---

## 🛠️ Comandos Úteis

### Ver Status Completo
```bash
./scripts/github-auth.sh status
```

### Forçar Novo Token
```bash
./scripts/github-auth.sh clear
./release.sh
```

### Usar Token Específico
```bash
GITHUB_TOKEN=ghp_xxxxx ./release.sh
```

### Cachear Token do gh CLI
```bash
gh auth login
./scripts/github-auth.sh get
```

---

## ❓ Troubleshooting

### Token Expirado
```bash
./scripts/github-auth.sh clear
./release.sh
# Inserir novo token
```

### Cache Corrompido
```bash
rm ~/.config/bagus-browser/github_token.cache.enc
./release.sh
```

### Permissão Negada
```bash
chmod 700 ~/.config/bagus-browser
chmod 600 ~/.config/bagus-browser/github_token.cache.enc
```

---

## 🎯 Vantagens

### ✅ Conveniência
- Token cacheado entre execuções
- Não precisa inserir toda vez
- Renovação automática

### ✅ Segurança
- Criptografia AES-256
- Validação antes de usar
- Permissões restritas

### ✅ Flexibilidade
- Múltiplas fontes de token
- Fallback automático
- Compatível com gh CLI

### ✅ Transparência
- Logs claros
- Status visível
- Fácil debug

---

## 📊 Comparação

| Método | Segurança | Conveniência | Cache |
|--------|-----------|--------------|-------|
| **GITHUB_TOKEN** | ⚠️ Texto plano | ⭐⭐ | ❌ |
| **gh CLI** | ✅ Seguro | ⭐⭐⭐ | ✅ |
| **OAuth2 Cache** | ✅ Criptografado | ⭐⭐⭐ | ✅ |

---

## ✅ Pronto para Usar!

O sistema OAuth2 com cache está integrado e funcionando automaticamente em:
- `./release.sh`
- `./publish.sh`

**Basta executar e seguir as instruções!** 🚀
