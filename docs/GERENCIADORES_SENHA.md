# 🔐 Usando Gerenciadores de Senha com Bagus Browser

## ⚠️ IMPORTANTE: Extensões NÃO Funcionam

O Bagus Browser **NÃO suporta extensões** por design (segurança e privacidade).

**Isso significa:**
- ❌ Extensão do Proton Pass NÃO funciona
- ❌ Extensão do Bitwarden NÃO funciona
- ❌ Extensão do 1Password NÃO funciona
- ❌ Auto-fill via extensão NÃO funciona

**Mas você PODE usar:**
- ✅ App desktop do gerenciador
- ✅ Copiar e colar (Ctrl+C / Ctrl+V)
- ✅ Auto-type (KeePassXC)

---

## 📋 Compatibilidade

O Bagus Browser é compatível com gerenciadores de senha que funcionam através de:

1. **App Desktop** (não extensão!)
2. **Copiar e colar (Ctrl+C / Ctrl+V)**
3. **Auto-type global (KeePassXC)**

---

## ✅ Gerenciadores Testados

### Proton Pass

**Status:** ✅ Compatível (via clipboard)

**⚠️ ATENÇÃO:** Use o **app desktop**, NÃO a extensão!

**Workflow Correto:**

#### Opção 1: App Desktop Proton Pass (Recomendado)
```
1. Instalar app desktop Proton Pass
   Download: https://proton.me/pass/download

2. Abrir Proton Pass desktop
3. Buscar credencial
4. Clicar em "Copy username" ou "Copy password"
5. Voltar ao Bagus Browser
6. Clicar no campo de login/senha
7. Ctrl+V (colar)
```

#### Opção 2: Extensão em Outro Navegador
```
1. Manter Chrome/Firefox aberto com Proton Pass
2. No Bagus Browser, abrir site de login
3. Alt+Tab para Chrome/Firefox
4. Abrir Proton Pass (extensão)
5. Copiar senha (Ctrl+C)
6. Alt+Tab para Bagus Browser
7. Ctrl+V no campo
```

**Vantagens:**
- ✅ Segurança mantida (E2E)
- ✅ Sem armazenamento local
- ✅ Workflow simples (2-3 cliques)
- ✅ Funciona com qualquer site

---

### Outros Gerenciadores

#### 1Password
- ✅ Compatível via clipboard
- ✅ App desktop disponível
- ✅ Atalho global Ctrl+\\ para buscar senha

#### Bitwarden
- ✅ Compatível via clipboard
- ✅ App desktop disponível
- ✅ Atalho global Ctrl+Shift+L para auto-fill

#### KeePassXC ⭐ RECOMENDADO

**Status:** ✅ Melhor compatibilidade (auto-type)

**Por que é o melhor:**
- ✅ Auto-type funciona PERFEITAMENTE
- ✅ Atalho global (Ctrl+Alt+A)
- ✅ Preenche usuário E senha automaticamente
- ✅ Não precisa copiar/colar
- ✅ Open source e gratuito

**Workflow:**
```
1. Instalar KeePassXC
   sudo apt install keepassxc

2. Configurar auto-type:
   - Tools > Settings > Browser Integration
   - Enable "Auto-Type"
   - Configure shortcut: Ctrl+Alt+A

3. Usar:
   - Abrir site de login no Bagus Browser
   - Clicar no campo de usuário
   - Pressionar Ctrl+Alt+A
   - Selecionar credencial
   - KeePassXC preenche tudo automaticamente!
```

#### Bitwarden
- ✅ Compatível via clipboard
- ✅ App desktop disponível
- ⚠️ Auto-fill NÃO funciona (precisa extensão)

#### LastPass
- ✅ Compatível via clipboard
- ✅ App desktop disponível
- ⚠️ Auto-fill NÃO funciona (precisa extensão)

---

## 🎯 Workflow Recomendado

### Opção 1: KeePassXC Auto-Type (MELHOR) ⭐

**Não precisa copiar/colar!**

```
1. Abrir site de login no Bagus Browser
2. Clicar no campo de usuário
3. Pressionar Ctrl+Alt+A
4. Selecionar credencial no KeePassXC
5. Pronto! Usuário e senha preenchidos automaticamente
```

**Vantagens:**
- ✅ Mais rápido (1 atalho)
- ✅ Mais seguro (não usa clipboard)
- ✅ Funciona em qualquer site
- ✅ Preenche múltiplos campos

---

### Opção 2: Clipboard (Proton Pass, Bitwarden, etc)

```
1. Abrir app desktop do gerenciador
2. Buscar credencial
3. Copiar usuário (Ctrl+C)
4. Voltar ao Bagus Browser
5. Colar no campo (Ctrl+V)
6. Repetir para senha
```

**Atalhos:**
- Copiar: `Ctrl+C` ou `Ctrl+Ins`
- Colar: `Ctrl+V` ou `Shift+Ins`

---

## 🔒 Segurança

### Por que não extensões?

O Bagus Browser **não suporta extensões** por design:

**Motivos:**
- ✅ Menor superfície de ataque
- ✅ Sem permissões excessivas
- ✅ Sem telemetria de extensões
- ✅ Mais leve e rápido
- ✅ Privacidade garantida

### Clipboard é seguro?

**Sim, quando usado corretamente:**

1. **Gerenciadores modernos** limpam o clipboard após uso
2. **Timeout automático** (30-60 segundos)
3. **Criptografia** dos dados copiados
4. **Sem histórico** persistente

**Dicas:**
- Use gerenciadores com auto-clear do clipboard
- Não deixe senhas copiadas por muito tempo
- Feche o gerenciador após uso

---

## 📱 Configuração por Gerenciador

### Proton Pass

```bash
# Instalar app desktop (opcional)
# Disponível em: https://proton.me/pass/download

# Usar extensão em outro navegador
# Chrome: https://chrome.google.com/webstore
# Firefox: https://addons.mozilla.org
```

**Workflow:**
1. Abra Proton Pass
2. Busque a credencial
3. Clique em "Copy password"
4. Cole no Bagus Browser

---

### KeePassXC

```bash
# Instalar no Ubuntu/Debian
sudo apt install keepassxc

# Configurar auto-type
# 1. Abra KeePassXC
# 2. Tools > Settings > Browser Integration
# 3. Enable "Auto-Type"
# 4. Configure global shortcut (Ctrl+Alt+A)
```

**Workflow:**
1. Foque no campo de login
2. Pressione `Ctrl+Alt+A`
3. Selecione a credencial
4. KeePassXC preenche automaticamente

---

### Bitwarden

```bash
# Instalar app desktop
# Download: https://bitwarden.com/download/

# Ou usar CLI
sudo snap install bw
```

**Workflow:**
1. Abra Bitwarden desktop
2. Busque a credencial
3. Clique em "Copy password"
4. Cole no Bagus Browser

---

## 🎨 Atalhos de Teclado

### Copiar/Colar

| Ação | Atalho Principal | Atalho Alternativo |
|------|------------------|-------------------|
| Copiar | `Ctrl+C` | `Ctrl+Ins` |
| Colar | `Ctrl+V` | `Shift+Ins` |
| Selecionar Tudo | `Ctrl+A` | - |

### Navegação

| Ação | Atalho |
|------|--------|
| Focar URL | `Ctrl+L` |
| Nova Aba | `Ctrl+T` |
| Fechar Aba | `Ctrl+W` |

---

## ❓ FAQ

### P: Posso instalar extensões no Bagus Browser?
**R:** Não. O Bagus Browser não suporta extensões por design (segurança e privacidade).

### P: Proton Pass funciona?
**R:** Sim, via clipboard. Copie a senha do Proton Pass e cole no Bagus Browser.

### P: Qual gerenciador é mais compatível?
**R:** KeePassXC (auto-type) e Bitwarden (clipboard) são os mais práticos.

### P: É seguro usar clipboard?
**R:** Sim, gerenciadores modernos limpam o clipboard automaticamente após 30-60 segundos.

### P: Posso usar o preenchimento automático do sistema?
**R:** Sim, se seu sistema operacional tiver um gerenciador de senhas integrado (como GNOME Keyring).

### P: Funciona com 2FA/TOTP?
**R:** Sim, copie o código TOTP do seu gerenciador e cole no site.

---

## 🔮 Futuro

### Planejado para v5.0

- [ ] **UserScripts:** Permitir scripts customizados
- [ ] **Native Messaging:** Comunicação com apps nativos
- [ ] **Clipboard Monitor:** Detectar senhas copiadas
- [ ] **Auto-fill básico:** Preencher campos detectados

---

## 📚 Recursos

### Documentação Oficial

- **Proton Pass:** https://proton.me/support/pass
- **KeePassXC:** https://keepassxc.org/docs/
- **Bitwarden:** https://bitwarden.com/help/
- **1Password:** https://support.1password.com/

### Tutoriais

- [Como usar KeePassXC com auto-type](https://keepassxc.org/docs/KeePassXC_GettingStarted.html#_auto_type)
- [Configurar Bitwarden CLI](https://bitwarden.com/help/cli/)
- [Proton Pass: Primeiros passos](https://proton.me/support/pass-get-started)

---

## ✅ Conclusão

O Bagus Browser é **totalmente compatível** com gerenciadores de senha através de:

1. ✅ **Clipboard** (todos os gerenciadores)
2. ✅ **Auto-type** (KeePassXC)
3. ✅ **Atalhos alternativos** (Ctrl+Ins / Shift+Ins)

**Recomendação:** Use KeePassXC para melhor experiência (auto-type) ou Proton Pass para máxima segurança (E2E).

---

**Atualizado em:** 22/10/2025  
**Versão:** 4.4.0
