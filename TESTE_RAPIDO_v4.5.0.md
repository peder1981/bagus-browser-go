# 🧪 Teste Rápido - v4.5.0

## ⚙️ Como Acessar as Configurações

### Método 1: Atalho de Teclado (Recomendado)
```
1. Abrir Bagus Browser
2. Pressionar: Ctrl+, (Ctrl + vírgula)
3. Janela de configurações deve abrir ✅
```

### Método 2: Via Terminal (Debug)
```bash
# Abrir browser com logs
bagus-browser 2>&1 | grep -i "config\|senha\|settings"
```

---

## 🔒 Testar Senha Mestre

### Passo a Passo:

1. **Abrir Configurações:**
   ```
   Ctrl+, (Ctrl + vírgula)
   ```

2. **Ir para Aba "Segurança":**
   - Deve ser a primeira aba
   - Ícone: 🔒

3. **Definir Senha:**
   ```
   ☑ Marcar: "Exigir senha ao abrir o browser"
   
   Diálogo deve aparecer:
   - Digite senha: teste123
   - Confirme senha: teste123
   - Clicar "Definir"
   ```

4. **Salvar:**
   ```
   Clicar "Salvar" no diálogo de configurações
   ```

5. **Testar:**
   ```
   Ctrl+Q (fechar browser)
   bagus-browser (reabrir)
   
   Deve pedir senha ✅
   ```

---

## 🔄 Testar Manter Logado

### Passo a Passo:

1. **Abrir Configurações:**
   ```
   Ctrl+,
   ```

2. **Ir para Aba "Sessões":**
   - Segunda aba
   - Ícone: 🔄

3. **Verificar Opções:**
   ```
   ✅ Manter usuário logado em sites (deve estar marcado)
   ✅ Salvar cookies entre sessões (deve estar marcado)
   ```

4. **Testar:**
   ```
   1. Abrir Gmail (mail.google.com)
   2. Fazer login
   3. Fechar browser (Ctrl+Q)
   4. Reabrir browser
   5. Abrir Gmail novamente
   
   Deve estar logado ✅
   ```

---

## 📹 Testar Cache de Vídeos

### Passo a Passo:

1. **Abrir Configurações:**
   ```
   Ctrl+,
   ```

2. **Ir para Aba "Performance":**
   - Terceira aba
   - Ícone: ⚡

3. **Verificar Cache:**
   ```
   ✅ Habilitar cache de vídeos (deve estar marcado)
   Slider: 500 MB (padrão)
   ```

4. **Testar:**
   ```
   1. Abrir YouTube
   2. Assistir vídeo
   3. Fechar aba (Ctrl+W)
   
   No terminal deve aparecer:
   🛑 Carregamento parado
   ```

---

## ❌ Problemas Comuns

### 1. Ctrl+, Não Funciona

**Verificar:**
```bash
# Ver se atalho está registrado
bagus-browser 2>&1 | grep "Ctrl+,"
```

**Deve aparecer:**
```
⌨️  Atalhos:
   ...
   Ctrl+, (configurações)
```

**Solução:**
```bash
# Recompilar e reinstalar
cd ~/bagus-browser-go
make build
make install
```

### 2. Diálogo Não Abre

**Verificar logs:**
```bash
bagus-browser 2>&1 | grep -i "settings\|config\|dialog"
```

**Quando pressionar Ctrl+, deve aparecer:**
```
⌨️  Ctrl+, - Configurações
```

### 3. Senha Não Funciona

**Verificar arquivo:**
```bash
ls -lh ~/.config/bagus-browser/config.enc
```

**Deve existir e ter ~1KB**

**Se não existir:**
```bash
# Remover e recriar
rm -f ~/.config/bagus-browser/config.enc
bagus-browser
# Configurar senha novamente
```

---

## 🔍 Debug Completo

### Ver Todos os Logs:

```bash
bagus-browser 2>&1 | tee bagus-debug.log
```

**Pressionar Ctrl+, e verificar no log:**
```
⌨️  Ctrl+, - Configurações
⚙️  Configurações carregadas
```

**Se aparecer erro, copiar e reportar**

---

## ✅ Checklist de Teste

- [ ] Ctrl+, abre janela de configurações
- [ ] Aba "Segurança" visível
- [ ] Aba "Sessões" visível
- [ ] Aba "Performance" visível
- [ ] Aba "Privacidade" visível
- [ ] Consegue definir senha
- [ ] Senha funciona ao reabrir
- [ ] Manter logado funciona
- [ ] Cache de vídeos funciona

---

## 📞 Se Nada Funcionar

### Reinstalação Completa:

```bash
# 1. Remover versão antiga
sudo dpkg -r bagus-browser

# 2. Limpar configurações (CUIDADO: apaga tudo)
rm -rf ~/.config/bagus-browser/

# 3. Recompilar
cd ~/bagus-browser-go
make clean
make build

# 4. Reinstalar
make install

# 5. Testar
bagus-browser
Ctrl+,
```

---

## 📊 Versão Instalada

**Verificar:**
```bash
bagus-browser --version
# Ou
dpkg -l | grep bagus-browser
```

**Deve mostrar:**
```
4.5.0
```

---

## 🎯 Teste Mínimo (30 segundos)

```bash
# 1. Abrir browser
bagus-browser

# 2. Pressionar Ctrl+,
# Deve abrir janela de configurações ✅

# 3. Ver 4 abas:
# - 🔒 Segurança
# - 🔄 Sessões
# - ⚡ Performance
# - 🕵️ Privacidade

# 4. Fechar (Esc ou X)
```

**Se isso funcionar, está tudo OK!** ✅

---

**Data:** 23/10/2025  
**Versão:** 4.5.0  
**Status:** Instalado
