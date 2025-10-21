# 🔑 Configurar GitHub Token para Releases

## 📝 Criar Token

1. **Ir para:** https://github.com/settings/tokens

2. **Clicar em:** "Generate new token" → "Generate new token (classic)"

3. **Configurar:**
   - **Note:** `Bagus Browser Releases`
   - **Expiration:** 90 days (ou No expiration)
   - **Scopes:** Marcar:
     - ✅ `repo` (Full control of private repositories)
       - ✅ `repo:status`
       - ✅ `repo_deployment`
       - ✅ `public_repo`
       - ✅ `repo:invite`

4. **Gerar token** e **COPIAR** (só aparece uma vez!)

---

## 🔧 Configurar Token

### Opção 1: Variável de Ambiente (Temporário)
```bash
export GITHUB_TOKEN=ghp_seu_token_aqui_1234567890
```

### Opção 2: Arquivo .bashrc (Permanente)
```bash
echo 'export GITHUB_TOKEN=ghp_seu_token_aqui_1234567890' >> ~/.bashrc
source ~/.bashrc
```

### Opção 3: Arquivo .env (Projeto)
```bash
# Criar arquivo .env
echo 'GITHUB_TOKEN=ghp_seu_token_aqui_1234567890' > .env

# Carregar antes de usar
source .env
```

---

## 🚀 Usar Release Script

### Com Token
```bash
# Definir token
export GITHUB_TOKEN=ghp_seu_token_aqui

# Executar release
./release.sh
```

### Ou Tudo em Uma Linha
```bash
GITHUB_TOKEN=ghp_seu_token_aqui ./release.sh
```

---

## ✅ Verificar Token

```bash
# Ver se está definido
echo $GITHUB_TOKEN

# Testar com API
curl -H "Authorization: token $GITHUB_TOKEN" \
     https://api.github.com/user
```

---

## 🔒 Segurança

### ⚠️ NUNCA:
- ❌ Commitar token no git
- ❌ Compartilhar token
- ❌ Deixar em arquivo público

### ✅ SEMPRE:
- ✅ Usar variável de ambiente
- ✅ Adicionar .env ao .gitignore
- ✅ Revogar tokens não usados
- ✅ Usar tokens com escopo mínimo

---

## 🗑️ Revogar Token

Se comprometer o token:
1. Ir para: https://github.com/settings/tokens
2. Clicar em "Delete" no token
3. Gerar novo token

---

## 📦 Exemplo Completo

```bash
# 1. Criar token no GitHub
# https://github.com/settings/tokens

# 2. Exportar token
export GITHUB_TOKEN=ghp_1234567890abcdef

# 3. Build
./build.sh

# 4. Release
./release.sh

# 5. Verificar
# https://github.com/peder1981/bagus-browser-go/releases
```

---

## 🔄 Alternativa: gh CLI

Se preferir não usar token diretamente:

```bash
# Instalar
sudo apt install gh

# Autenticar
gh auth login

# Usar
./release.sh  # Detecta gh CLI automaticamente
```

---

## ✅ Script Atualizado

O `release.sh` agora suporta:
1. **GITHUB_TOKEN** (preferido)
2. **gh CLI** (fallback automático)

**Prioridade:**
1. Se `GITHUB_TOKEN` definido → Usa API REST
2. Se não, tenta `gh CLI` → Usa gh CLI
3. Se nenhum → Erro com instruções

---

**Pronto para publicar releases com token!** 🚀🔑
