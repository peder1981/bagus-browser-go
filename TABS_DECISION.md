# 🎯 Decisão: Múltiplas Abas no Bagus Browser

## 📊 Resumo da Investigação

Você solicitou **múltiplas abas** mantendo **Go** como linguagem.

**Investigação realizada:**
1. ✅ Wails v2/v3 analisado
2. ✅ GTK + WebKit2GTK pesquisado
3. ✅ Alternativas comparadas

---

## 🔍 Descobertas Principais

### ❌ Wails NÃO é adequado
- **Wails v2:** Não suporta múltiplas janelas
- **Wails v3:** Alpha, cria janelas separadas (não abas)
- **`<iframe>`:** Bloqueado pela maioria dos sites
- **Conclusão:** Não serve para browser com abas

### ✅ GTK + WebKit2GTK É ADEQUADO
- **Go puro:** Mantém linguagem
- **Abas nativas:** GtkNotebook
- **WebKit completo:** Sem limitações
- **Leve:** ~10-15MB
- **Conclusão:** Solução ideal para Go

---

## 🏆 Solução Recomendada

### GTK3 + WebKit2GTK (Go Puro)

**Stack:**
```
Go + gotk3/gotk3 + sourcegraph/go-webkit2
```

**Arquitetura:**
```
┌─────────────────────────────────────┐
│  GtkWindow                          │
│  ┌─────────────────────────────┐    │
│  │ Toolbar: [←][→][⟳][URL][+] │    │
│  └─────────────────────────────┘    │
│  ┌─────────────────────────────┐    │
│  │ GtkNotebook                 │    │
│  │ ┌─────┬─────┬─────┐         │    │
│  │ │ Aba1│ Aba2│ Aba3│         │    │
│  │ └─────┴─────┴─────┘         │    │
│  │ ┌─────────────────────┐     │    │
│  │ │ WebKitWebView       │     │    │
│  │ │ (site renderizado)  │     │    │
│  │ └─────────────────────┘     │    │
│  └─────────────────────────────┘    │
└─────────────────────────────────────┘
```

**Prós:**
- ✅ Go puro (mantém stack)
- ✅ Abas nativas (GtkNotebook)
- ✅ WebKit completo (sem limitações)
- ✅ Leve (~10-15MB)
- ✅ Escalável (pode adicionar features)

**Contras:**
- ⚠️ Requer reescrita (~2-3 dias)
- ⚠️ Mais complexo que solução atual
- ⚠️ Portabilidade média (melhor em Linux)

---

## 📋 Plano de Implementação

### Opção A: POC Primeiro (Recomendado)
**Tempo:** 4-6 horas

1. Criar projeto novo `bagus-gtk-poc/`
2. Implementar janela GTK básica
3. Adicionar 1 WebView
4. Testar navegação
5. Adicionar 2-3 abas
6. Validar viabilidade

**Resultado:** Decisão informada antes de reescrever tudo

---

### Opção B: Implementação Completa
**Tempo:** 2-3 dias

**Fase 1: Base (6-8h)**
- Estrutura do projeto
- Janela GTK + Notebook
- 1 WebView funcional

**Fase 2: Abas (4-6h)**
- Sistema de abas
- Criar/fechar abas
- Navegação por aba

**Fase 3: UI (4-6h)**
- Barra de ferramentas
- Botões navegação
- Campo URL

**Fase 4: Features (6-8h)**
- Histórico por aba
- Atalhos de teclado
- Configurações

**Fase 5: Polimento (4-6h)**
- Ícones
- Temas
- Otimizações

**Total:** ~24-34 horas (~3-4 dias)

---

## 🎯 Alternativas (Se Não Quiser Go)

### Electron (JavaScript)
**Prós:**
- ✅ Abas triviais
- ✅ Documentação extensa
- ✅ Muitos exemplos

**Contras:**
- ❌ Não é Go
- ❌ ~100-150MB
- ❌ Alto uso de memória

**Tempo:** ~1-2 dias (mais fácil)

---

### Tauri (Rust)
**Prós:**
- ✅ Leve (~5-8MB)
- ✅ Moderno
- ✅ Seguro

**Contras:**
- ❌ Não é Go (Rust)
- ⚠️ Curva de aprendizado

**Tempo:** ~3-5 dias (aprender Rust)

---

## 📊 Comparação Final

| Solução | Linguagem | Abas | Tamanho | Tempo | Viável |
|---------|-----------|------|---------|-------|--------|
| **GTK+WebKit** | ✅ Go | ✅ Sim | ~12MB | 3-4d | ✅ |
| **Electron** | ❌ JS | ✅ Sim | ~120MB | 1-2d | ✅ |
| **Tauri** | ❌ Rust | ✅ Sim | ~6MB | 3-5d | ✅ |
| **Wails** | ✅ Go | ❌ Não | ~8MB | - | ❌ |
| **Atual** | ✅ Go | ❌ Não | ~7MB | - | ⚠️ |

---

## 🤔 Próxima Decisão

**Você precisa escolher:**

### A) Criar POC GTK+WebKit (4-6h)
- Validar solução antes de comprometer
- Testar viabilidade técnica
- Ver se atende necessidades
- **Recomendado:** Menos risco

### B) Implementar GTK+WebKit completo (3-4d)
- Ir direto para implementação
- Confiar na análise
- Começar do zero
- **Mais rápido:** Se tiver certeza

### C) Migrar para Electron (1-2d)
- Aceitar JavaScript
- Ganhar abas facilmente
- Perder Go
- **Mais fácil:** Mas muda stack

### D) Manter como está
- Aceitar sem abas
- Focar em outras features
- Browser minimalista
- **Mais simples:** Mas limitado

---

## 💡 Minha Recomendação

**Opção A: Criar POC GTK+WebKit**

**Por quê:**
1. ✅ Valida solução (4-6h investidas)
2. ✅ Mantém Go
3. ✅ Decisão informada
4. ✅ Baixo risco

**Se POC funcionar:**
→ Implementar completo (3-4 dias)

**Se POC não funcionar:**
→ Avaliar Electron ou aceitar sem abas

---

## 🚀 Ação Imediata

**Posso criar POC agora?**

**Se sim:**
1. Criar `bagus-gtk-poc/`
2. Implementar janela + 1 WebView
3. Adicionar 2-3 abas
4. Testar navegação
5. Mostrar resultado

**Tempo estimado:** 4-6 horas

**O que você decide?**
