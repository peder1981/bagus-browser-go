# ✅ Solução: GTK3 + WebKit2GTK para Múltiplas Abas

## 🎯 Solução Recomendada: Go Puro com Abas Nativas

**Stack:**
- **Go** - Linguagem principal
- **GTK3** - Interface nativa (via `gotk3/gotk3`)
- **WebKit2GTK** - Motor de renderização (via `sourcegraph/go-webkit2`)
- **GtkNotebook** - Widget de abas nativo

---

## 🏗️ Arquitetura

```
┌─────────────────────────────────────────────┐
│  GtkWindow (Janela Principal)               │
│  ┌─────────────────────────────────────┐    │
│  │  GtkHeaderBar                       │    │
│  │  [←] [→] [⟳] [URL Input] [+]       │    │
│  └─────────────────────────────────────┘    │
│  ┌─────────────────────────────────────┐    │
│  │  GtkNotebook (Abas)                 │    │
│  │  ┌─────┬─────┬─────┐               │    │
│  │  │ Aba1│ Aba2│ Aba3│               │    │
│  │  └─────┴─────┴─────┘               │    │
│  │  ┌─────────────────────────────┐   │    │
│  │  │ GtkScrolledWindow           │   │    │
│  │  │  ┌─────────────────────┐    │   │    │
│  │  │  │ WebKitWebView       │    │   │    │
│  │  │  │ (renderiza site)    │    │   │    │
│  │  │  └─────────────────────┘    │   │    │
│  │  └─────────────────────────────┘   │    │
│  └─────────────────────────────────────┘    │
└─────────────────────────────────────────────┘
```

---

## 📦 Dependências

### Go Packages
```bash
go get github.com/gotk3/gotk3/gtk
go get github.com/gotk3/gotk3/glib
go get github.com/sourcegraph/go-webkit2/webkit2
```

### Sistema (Ubuntu/Debian)
```bash
sudo apt-get install -y \
    libgtk-3-dev \
    libwebkit2gtk-4.0-dev \
    pkg-config
```

---

## 💻 Exemplo de Código

### Estrutura Básica

```go
package main

import (
    "log"
    "runtime"
    
    "github.com/gotk3/gotk3/glib"
    "github.com/gotk3/gotk3/gtk"
    "github.com/sourcegraph/go-webkit2/webkit2"
)

type Browser struct {
    window   *gtk.Window
    notebook *gtk.Notebook
    urlEntry *gtk.Entry
    tabs     []*Tab
}

type Tab struct {
    webView   *webkit2.WebView
    container *gtk.ScrolledWindow
    label     *gtk.Label
    url       string
}

func main() {
    runtime.LockOSThread()
    gtk.Init(nil)
    
    browser := NewBrowser()
    browser.Show()
    
    gtk.Main()
}

func NewBrowser() *Browser {
    // Criar janela principal
    win, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
    win.SetTitle("Bagus Browser v4.0.0")
    win.SetDefaultSize(1200, 800)
    win.Connect("destroy", gtk.MainQuit)
    
    // Criar notebook (abas)
    notebook, _ := gtk.NotebookNew()
    notebook.SetScrollable(true)
    
    // Criar barra de ferramentas
    toolbar := createToolbar()
    
    // Layout vertical
    vbox, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
    vbox.PackStart(toolbar, false, false, 0)
    vbox.PackStart(notebook, true, true, 0)
    
    win.Add(vbox)
    
    browser := &Browser{
        window:   win,
        notebook: notebook,
        tabs:     make([]*Tab, 0),
    }
    
    // Criar primeira aba
    browser.NewTab("https://duckduckgo.com")
    
    return browser
}

func (b *Browser) NewTab(url string) {
    // Criar WebView
    webView := webkit2.NewWebView()
    
    // Criar container scrollable
    scrolled, _ := gtk.ScrolledWindowNew(nil, nil)
    scrolled.Add(webView)
    
    // Criar label da aba
    label, _ := gtk.LabelNew("Nova Aba")
    
    // Adicionar aba ao notebook
    pageNum := b.notebook.AppendPage(scrolled, label)
    b.notebook.SetCurrentPage(pageNum)
    
    // Criar tab struct
    tab := &Tab{
        webView:   webView,
        container: scrolled,
        label:     label,
        url:       url,
    }
    
    // Conectar eventos
    webView.Connect("load-changed", func(_ *glib.Object, i int) {
        loadEvent := webkit2.LoadEvent(i)
        if loadEvent == webkit2.LoadFinished {
            title := webView.Title()
            label.SetText(title)
        }
    })
    
    // Carregar URL
    webView.LoadURI(url)
    
    b.tabs = append(b.tabs, tab)
    b.window.ShowAll()
}

func (b *Browser) Navigate(url string) {
    pageNum := b.notebook.GetCurrentPage()
    if pageNum >= 0 && pageNum < len(b.tabs) {
        tab := b.tabs[pageNum]
        tab.webView.LoadURI(url)
    }
}

func (b *Browser) GoBack() {
    pageNum := b.notebook.GetCurrentPage()
    if pageNum >= 0 && pageNum < len(b.tabs) {
        tab := b.tabs[pageNum]
        if tab.webView.CanGoBack() {
            tab.webView.GoBack()
        }
    }
}

func (b *Browser) GoForward() {
    pageNum := b.notebook.GetCurrentPage()
    if pageNum >= 0 && pageNum < len(b.tabs) {
        tab := b.tabs[pageNum]
        if tab.webView.CanGoForward() {
            tab.webView.GoForward()
        }
    }
}

func (b *Browser) Reload() {
    pageNum := b.notebook.GetCurrentPage()
    if pageNum >= 0 && pageNum < len(b.tabs) {
        tab := b.tabs[pageNum]
        tab.webView.Reload()
    }
}

func (b *Browser) Show() {
    b.window.ShowAll()
}

func createToolbar() *gtk.Box {
    toolbar, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 5)
    toolbar.SetMarginTop(5)
    toolbar.SetMarginBottom(5)
    toolbar.SetMarginStart(5)
    toolbar.SetMarginEnd(5)
    
    // Botão Voltar
    btnBack, _ := gtk.ButtonNewFromIconName("go-previous", gtk.ICON_SIZE_BUTTON)
    toolbar.PackStart(btnBack, false, false, 0)
    
    // Botão Avançar
    btnForward, _ := gtk.ButtonNewFromIconName("go-next", gtk.ICON_SIZE_BUTTON)
    toolbar.PackStart(btnForward, false, false, 0)
    
    // Botão Recarregar
    btnReload, _ := gtk.ButtonNewFromIconName("view-refresh", gtk.ICON_SIZE_BUTTON)
    toolbar.PackStart(btnReload, false, false, 0)
    
    // Campo URL
    urlEntry, _ := gtk.EntryNew()
    urlEntry.SetPlaceholderText("Digite URL ou busque...")
    toolbar.PackStart(urlEntry, true, true, 0)
    
    // Botão Nova Aba
    btnNewTab, _ := gtk.ButtonNewFromIconName("tab-new", gtk.ICON_SIZE_BUTTON)
    toolbar.PackStart(btnNewTab, false, false, 0)
    
    return toolbar
}
```

---

## ✅ Vantagens desta Solução

### 1. Go Puro
- ✅ Mantém linguagem Go
- ✅ Sem JavaScript no backend
- ✅ Type-safe

### 2. Abas Nativas
- ✅ GtkNotebook é widget nativo
- ✅ Familiar para usuários Linux
- ✅ Suporta drag-and-drop de abas
- ✅ Scrollable (muitas abas)

### 3. WebKit Completo
- ✅ Motor completo (não iframe)
- ✅ Sem limitações X-Frame-Options
- ✅ Renderiza qualquer site
- ✅ JavaScript completo

### 4. Leve
- ✅ ~10-15MB binário
- ✅ Usa WebKit do sistema
- ✅ Memória eficiente

### 5. Recursos Nativos
- ✅ Histórico por aba
- ✅ Cookies isolados
- ✅ Cache por aba
- ✅ Downloads

---

## ⚠️ Considerações

### Memória
- Cada aba = 1 WebView (~30-50MB)
- 10 abas = ~300-500MB
- Aceitável para browser moderno

### Complexidade
- Mais código que solução atual
- Gerenciamento de estado por aba
- Eventos GTK + WebKit

### Portabilidade
- Linux: ✅ Nativo
- Windows: ⚠️ Requer GTK runtime
- macOS: ⚠️ Requer GTK via Homebrew

---

## 🚀 Plano de Implementação

### Fase 1: POC (Proof of Concept)
- [ ] Criar janela GTK básica
- [ ] Adicionar 1 WebView
- [ ] Testar navegação
- [ ] Verificar memória

### Fase 2: Abas Básicas
- [ ] Implementar GtkNotebook
- [ ] Criar/fechar abas
- [ ] Navegação por aba
- [ ] Atualizar título das abas

### Fase 3: Barra de Ferramentas
- [ ] Botões navegação
- [ ] Campo URL
- [ ] Botão nova aba
- [ ] Conectar eventos

### Fase 4: Features Avançadas
- [ ] Histórico por aba
- [ ] Favoritos
- [ ] Downloads
- [ ] Configurações

### Fase 5: Polimento
- [ ] Atalhos de teclado
- [ ] Ícones personalizados
- [ ] Temas
- [ ] Otimizações

---

## 📊 Comparação: Solução Atual vs GTK+WebKit

| Aspecto | Atual (WebView) | GTK+WebKit |
|---------|-----------------|------------|
| **Abas** | ❌ Não | ✅ Sim |
| **Linguagem** | ✅ Go | ✅ Go |
| **Tamanho** | ✅ 6.6MB | ⚠️ 10-15MB |
| **UI Visual** | ❌ Não | ✅ Sim |
| **Complexidade** | ✅ Simples | ⚠️ Média |
| **Portabilidade** | ✅ Alta | ⚠️ Média |
| **Funcionalidade** | ⚠️ Básica | ✅ Completa |

---

## 🎯 Recomendação

**Migrar para GTK3 + WebKit2GTK** porque:

1. ✅ **Mantém Go** - Não precisa mudar linguagem
2. ✅ **Abas nativas** - GtkNotebook é robusto
3. ✅ **WebKit completo** - Sem limitações
4. ✅ **Leve** - ~10-15MB é aceitável
5. ✅ **Escalável** - Pode adicionar features

**Único contra:**
- ⚠️ Mais complexo que solução atual
- ⚠️ Requer reescrita (~2-3 dias)

---

## 🤔 Decisão

**Você quer:**

**A)** Implementar GTK+WebKit agora (2-3 dias de trabalho)  
**B)** Criar POC primeiro para validar (4-6 horas)  
**C)** Explorar outras alternativas antes  

**Qual opção?**
