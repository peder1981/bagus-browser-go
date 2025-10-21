# 🔍 Análise: Wails para Múltiplas Abas

## 📊 Resumo Executivo

**Wails v2:** ❌ **NÃO suporta múltiplas janelas nativamente**  
**Wails v3:** ✅ **SUPORTA múltiplas janelas** (Alpha desde 2024)

---

## 🎯 Descobertas Principais

### Wails v2 (Atual Estável)
- **Status:** Estável, produção
- **Múltiplas Janelas:** ❌ NÃO (issue #1480 aberta desde 2022)
- **Arquitetura:** 1 janela apenas
- **Limitação:** API declarativa single-window

### Wails v3 (Alpha)
- **Status:** ⚠️ Alpha (não produção)
- **Múltiplas Janelas:** ✅ SIM (feature principal)
- **Arquitetura:** API procedural multi-window
- **Exemplo:**
```go
app := application.New(application.Options{
    Name: "Multi Window Demo",
})

window1 := app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
    Title: "Window 1",
})

window2 := app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
    Title: "Window 2",
})

window1.SetURL("/")
window2.SetURL("https://wails.app")

app.Run()
```

---

## 🏗️ Arquitetura para Browser com Abas

### Opção A: Wails v3 - Múltiplas Janelas (1 janela por aba)
```
┌─────────────────────────────────────┐
│  Janela Principal (Wails)           │
│  ┌─────┬─────┬─────┐                │
│  │ Aba1│ Aba2│ Aba3│ (HTML/CSS)     │
│  └─────┴─────┴─────┘                │
│                                      │
│  [Controles de navegação]           │
└─────────────────────────────────────┘
         ↓ Cria novas janelas
┌─────────────────────────────────────┐
│  Janela Aba 1 (WebView)             │
│  [Conteúdo do site]                 │
└─────────────────────────────────────┘
┌─────────────────────────────────────┐
│  Janela Aba 2 (WebView)             │
│  [Conteúdo do site]                 │
└─────────────────────────────────────┘
```

**Prós:**
- ✅ Go puro
- ✅ Múltiplas janelas nativas
- ✅ Cada aba = janela independente

**Contras:**
- ❌ Wails v3 está em Alpha
- ❌ Não é "abas" tradicionais (são janelas separadas)
- ⚠️ Pode ser confuso para usuário

---

### Opção B: Wails v2/v3 - Abas HTML (1 janela, múltiplos iframes)
```
┌─────────────────────────────────────┐
│  Janela Wails                       │
│  ┌─────┬─────┬─────┐                │
│  │ Aba1│ Aba2│ Aba3│ (HTML tabs)    │
│  └─────┴─────┴─────┘                │
│  ┌─────────────────────────────┐    │
│  │ <iframe src="site1.com">    │    │
│  │ <iframe src="site2.com">    │    │
│  │ (troca visibilidade)        │    │
│  └─────────────────────────────┘    │
└─────────────────────────────────────┘
```

**Prós:**
- ✅ Abas visuais tradicionais
- ✅ Funciona em Wails v2 e v3
- ✅ UI familiar

**Contras:**
- ❌ `<iframe>` tem limitações (X-Frame-Options)
- ❌ Muitos sites bloqueiam iframes
- ❌ Problemas de segurança (CSP)
- ❌ **NÃO FUNCIONA para browser genérico**

---

### Opção C: Wails v3 - Abas com WebView Embarcado
```
┌─────────────────────────────────────┐
│  Janela Wails (Frontend HTML)       │
│  ┌─────┬─────┬─────┐                │
│  │ Aba1│ Aba2│ Aba3│                │
│  └─────┴─────┴─────┘                │
│  ┌─────────────────────────────┐    │
│  │ <webview> tag (Electron-like)│   │
│  │ (se Wails suportar)         │    │
│  └─────────────────────────────┘    │
└─────────────────────────────────────┘
```

**Status:** ⚠️ **PRECISA VERIFICAR** se Wails v3 suporta `<webview>` tag

---

## 🔬 Limitações Críticas do Wails

### 1. Wails v2
- ❌ **NÃO suporta múltiplas janelas**
- ❌ Issue #1480 aberta desde 2022
- ❌ Sem previsão de implementação em v2

### 2. Wails v3
- ⚠️ **Alpha** - não recomendado para produção
- ✅ Suporta múltiplas janelas
- ❌ Mas são **janelas separadas**, não abas

### 3. Problema do `<iframe>`
- ❌ Maioria dos sites bloqueia iframe (X-Frame-Options)
- ❌ Google, Facebook, Twitter, etc. não funcionam
- ❌ **Inviável para browser genérico**

---

## 📊 Comparação: Wails vs Outras Soluções

| Solução | Abas Nativas | Go | Estável | Tamanho | Viável |
|---------|--------------|-----|---------|---------|--------|
| **Wails v2** | ❌ | ✅ | ✅ | ~8MB | ❌ |
| **Wails v3** | ⚠️ (janelas) | ✅ | ❌ (Alpha) | ~8MB | ⚠️ |
| **Electron** | ✅ | ❌ | ✅ | ~100MB | ✅ |
| **Tauri** | ✅ | ❌ | ✅ | ~5MB | ✅ |
| **GTK+WebKit** | ✅ | ✅ | ✅ | ~10MB | ✅ |

---

## 🎯 Conclusão sobre Wails

### ❌ Wails NÃO é a solução ideal para browser com abas porque:

1. **Wails v2:** Não suporta múltiplas janelas
2. **Wails v3:** Alpha, instável, e cria janelas separadas (não abas)
3. **`<iframe>`:** Não funciona para sites genéricos
4. **Sem `<webview>` tag:** Wails não tem equivalente ao Electron

### ✅ Wails É BOM para:
- Aplicações single-window
- Dashboards
- Ferramentas administrativas
- Apps com UI customizada

### ❌ Wails NÃO É BOM para:
- **Browsers com abas** ← Nosso caso
- Apps que precisam renderizar sites externos
- Apps que precisam múltiplas views de conteúdo web

---

## 🚀 Alternativas Recomendadas

### Opção 1: GTK3 + WebKit2GTK (Go Puro) ⭐ RECOMENDADO
**Arquitetura:**
```go
// Usar GTK Notebook + WebKit2GTK diretamente
window := gtk.WindowNew()
notebook := gtk.NotebookNew()

for i := 0; i < numAbas; i++ {
    webview := webkit2.WebViewNew()
    scrolled := gtk.ScrolledWindowNew()
    scrolled.Add(webview)
    notebook.AppendPage(scrolled, label)
}
```

**Prós:**
- ✅ Go puro
- ✅ Abas nativas (GTK Notebook)
- ✅ WebKit2GTK completo
- ✅ Sem limitações de iframe
- ✅ ~10-15MB

**Contras:**
- ⚠️ Requer bindings GTK/WebKit (gotk3)
- ⚠️ Mais complexo que Wails
- ⚠️ Precisa gerenciar GTK manualmente

**Viabilidade:** ✅ **ALTA** - Solução mais adequada para Go

---

### Opção 2: Electron (JavaScript) ⭐ MAIS FÁCIL
**Prós:**
- ✅ Abas triviais
- ✅ `<webview>` tag dedicada
- ✅ Documentação extensa
- ✅ Muitos exemplos de browsers

**Contras:**
- ❌ Não é Go
- ❌ ~100-150MB
- ❌ Alto uso de memória

**Viabilidade:** ✅ **ALTA** - Mas abandona Go

---

### Opção 3: Tauri (Rust)
**Prós:**
- ✅ Leve (~5-8MB)
- ✅ Múltiplas janelas
- ✅ Moderno

**Contras:**
- ❌ Não é Go (Rust)
- ⚠️ Curva de aprendizado

**Viabilidade:** ✅ **MÉDIA** - Mas abandona Go

---

## 🎯 Recomendação Final

### Para manter Go:
**Usar GTK3 + WebKit2GTK diretamente** (sem Wails)

### Para ter abas facilmente:
**Migrar para Electron** (abandona Go)

### Wails:
❌ **NÃO RECOMENDADO** para browser com abas

---

## 📝 Próximos Passos

**Se escolher GTK + WebKit2GTK (Go):**
1. Pesquisar bindings `gotk3/gotk3`
2. Verificar suporte a `webkit2gtk`
3. Criar POC com 1 aba
4. Implementar sistema de abas

**Se escolher Electron (JavaScript):**
1. Criar projeto Electron
2. Implementar abas com `<webview>`
3. Migrar lógica de navegação
4. Aceitar ~100MB de overhead

**Qual caminho você prefere?**
