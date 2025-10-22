# 🔐 Usando Gerenciadores de Senha com Bagus Browser

## 📋 Compatibilidade

O Bagus Browser é compatível com gerenciadores de senha que funcionam através de:

1. **Preenchimento automático nativo do sistema**
2. **Copiar e colar (Ctrl+C / Ctrl+V)**
3. **Atalhos alternativos (Ctrl+Ins / Shift+Ins)**

---

## ✅ Gerenciadores Testados

### Proton Pass

**Status:** ✅ Compatível (via clipboard)

**Como usar:**
1. Instale a extensão Proton Pass no navegador principal (Chrome/Firefox)
2. Ou use o app desktop Proton Pass
3. Copie a senha do Proton Pass (Ctrl+C)
4. Cole no Bagus Browser (Ctrl+V ou Shift+Ins)

**Vantagens:**
- ✅ Segurança mantida
- ✅ Sem armazenamento local de senhas
- ✅ Criptografia E2E do Proton Pass
- ✅ Workflow simples

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

#### KeePassXC
- ✅ Compatível via clipboard
- ✅ Auto-type disponível
- ✅ Atalho global Ctrl+Alt+A para auto-type

#### LastPass
- ✅ Compatível via clipboard
- ✅ App desktop disponível

---

## 🎯 Workflow Recomendado

### Opção 1: Clipboard (Mais Seguro)

1. Abra seu gerenciador de senha
2. Busque a credencial desejada
3. Copie o usuário/senha
4. Cole no Bagus Browser

**Atalhos:**
- Copiar: `Ctrl+C` ou `Ctrl+Ins`
- Colar: `Ctrl+V` ou `Shift+Ins`

### Opção 2: Auto-Type (KeePassXC)

1. Configure auto-type no KeePassXC
2. Foque no campo de login
3. Pressione `Ctrl+Alt+A`
4. KeePassXC preenche automaticamente

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
