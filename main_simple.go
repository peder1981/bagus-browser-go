package main

/*
#cgo pkg-config: webkit2gtk-4.0
#include <webkit2/webkit2.h>

static WebKitWebView* create_webview() {
    return WEBKIT_WEB_VIEW(webkit_web_view_new());
}

static void load_uri(WebKitWebView* webview, const char* uri) {
    webkit_web_view_load_uri(webview, uri);
}

static void go_back(WebKitWebView* webview) {
    webkit_web_view_go_back(webview);
}

static void go_forward(WebKitWebView* webview) {
    webkit_web_view_go_forward(webview);
}

static void reload_page(WebKitWebView* webview) {
    webkit_web_view_reload(webview);
}

static gboolean can_go_back(WebKitWebView* webview) {
    return webkit_web_view_can_go_back(webview);
}

static gboolean can_go_forward(WebKitWebView* webview) {
    return webkit_web_view_can_go_forward(webview);
}

static const char* get_uri(WebKitWebView* webview) {
    return webkit_web_view_get_uri(webview);
}

static const char* get_title(WebKitWebView* webview) {
    return webkit_web_view_get_title(webview);
}
*/
import "C"
import (
	"fmt"
	"log"
	"runtime"
	"unsafe"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

// WebView encapsula WebKitWebView
type WebView struct {
	cWebView *C.WebKitWebView
	widget   *gtk.Widget
}

// NewWebView cria um novo WebView
func NewWebView() (*WebView, error) {
	cWebView := C.create_webview()
	if cWebView == nil {
		return nil, fmt.Errorf("falha ao criar WebView")
	}

	// Converter para GtkWidget
	gtkWidget := (*C.GtkWidget)(unsafe.Pointer(cWebView))
	
	// Criar gtk.Widget do gotk3
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(gtkWidget))}
	widget := &gtk.Widget{glib.InitiallyUnowned{obj}}

	return &WebView{
		cWebView: cWebView,
		widget:   widget,
	}, nil
}

// Widget retorna o gtk.Widget
func (w *WebView) Widget() *gtk.Widget {
	return w.widget
}

// LoadURI carrega uma URL
func (w *WebView) LoadURI(uri string) {
	cURI := C.CString(uri)
	defer C.free(unsafe.Pointer(cURI))
	C.load_uri(w.cWebView, cURI)
}

// GoBack volta na história
func (w *WebView) GoBack() {
	if C.can_go_back(w.cWebView) != 0 {
		C.go_back(w.cWebView)
	}
}

// GoForward avança na história
func (w *WebView) GoForward() {
	if C.can_go_forward(w.cWebView) != 0 {
		C.go_forward(w.cWebView)
	}
}

// Reload recarrega a página
func (w *WebView) Reload() {
	C.reload_page(w.cWebView)
}

// GetURI retorna a URI atual
func (w *WebView) GetURI() string {
	cURI := C.get_uri(w.cWebView)
	if cURI == nil {
		return ""
	}
	return C.GoString(cURI)
}

// GetTitle retorna o título da página
func (w *WebView) GetTitle() string {
	cTitle := C.get_title(w.cWebView)
	if cTitle == nil {
		return ""
	}
	return C.GoString(cTitle)
}

// Tab representa uma aba com WebView e label
type Tab struct {
	webView *WebView
	label   *gtk.Label
}

// Browser representa o navegador
type Browser struct {
	window   *gtk.Window
	notebook *gtk.Notebook
	urlEntry *gtk.Entry
	tabs     []*Tab
}

func main() {
	runtime.LockOSThread()
	gtk.Init(nil)

	log.Println("🌐 Iniciando Bagus Browser POC - WebKit CGO...")

	browser := NewBrowser()
	browser.Show()

	log.Println("✅ Browser iniciado com WebView!")
	log.Println("📝 Navegue para diferentes sites")
	log.Println("⌨️  Atalhos: Ctrl+T (nova aba), Ctrl+W (fechar), Alt+← (voltar), Alt+→ (avançar)")

	gtk.Main()
}

// NewBrowser cria uma nova instância do browser
func NewBrowser() *Browser {
	// Criar janela
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Erro ao criar janela:", err)
	}

	win.SetTitle("Bagus Browser POC - WebKit CGO")
	win.SetDefaultSize(1200, 800)
	win.Connect("destroy", gtk.MainQuit)

	// Criar notebook (abas)
	notebook, err := gtk.NotebookNew()
	if err != nil {
		log.Fatal("Erro ao criar notebook:", err)
	}
	notebook.SetScrollable(true)

	// Criar entry URL
	urlEntry, err := gtk.EntryNew()
	if err != nil {
		log.Fatal("Erro ao criar entry:", err)
	}

	browser := &Browser{
		window:   win,
		notebook: notebook,
		urlEntry: urlEntry,
		tabs:     make([]*Tab, 0),
	}

	// Criar toolbar
	toolbar := browser.createToolbar()

	// Layout
	vbox, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	vbox.PackStart(toolbar, false, false, 0)
	vbox.PackStart(notebook, true, true, 0)

	win.Add(vbox)

	// Conectar atalhos
	browser.setupKeyboardShortcuts()

	// Criar primeira aba
	browser.NewTab("https://duckduckgo.com")

	return browser
}

// createToolbar cria a barra de ferramentas
func (b *Browser) createToolbar() *gtk.Box {
	toolbar, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 5)
	toolbar.SetMarginTop(5)
	toolbar.SetMarginBottom(5)
	toolbar.SetMarginStart(5)
	toolbar.SetMarginEnd(5)

	// Botão Voltar
	btnBack, _ := gtk.ButtonNewWithLabel("←")
	btnBack.Connect("clicked", func() {
		b.GoBack()
	})
	toolbar.PackStart(btnBack, false, false, 0)

	// Botão Avançar
	btnForward, _ := gtk.ButtonNewWithLabel("→")
	btnForward.Connect("clicked", func() {
		b.GoForward()
	})
	toolbar.PackStart(btnForward, false, false, 0)

	// Botão Recarregar
	btnReload, _ := gtk.ButtonNewWithLabel("⟳")
	btnReload.Connect("clicked", func() {
		b.Reload()
	})
	toolbar.PackStart(btnReload, false, false, 0)

	// Campo URL
	b.urlEntry.SetPlaceholderText("Digite URL ou busque...")
	b.urlEntry.Connect("activate", func() {
		text, _ := b.urlEntry.GetText()
		b.Navigate(text)
	})
	toolbar.PackStart(b.urlEntry, true, true, 0)

	// Botão Nova Aba
	btnNewTab, _ := gtk.ButtonNewWithLabel("+")
	btnNewTab.Connect("clicked", func() {
		b.NewTab("https://duckduckgo.com")
	})
	toolbar.PackStart(btnNewTab, false, false, 0)

	return toolbar
}

// setupKeyboardShortcuts configura atalhos de teclado
func (b *Browser) setupKeyboardShortcuts() {
	b.window.Connect("key-press-event", func(win *gtk.Window, event *gdk.Event) bool {
		keyEvent := gdk.EventKeyNewFromEvent(event)
		keyVal := keyEvent.KeyVal()
		state := keyEvent.State()

		ctrlPressed := (state & uint(gdk.CONTROL_MASK)) != 0
		altPressed := (state & uint(gdk.MOD1_MASK)) != 0

		// Ctrl+T - Nova aba
		if ctrlPressed && keyVal == gdk.KEY_t {
			log.Println("⌨️  Ctrl+T - Nova aba")
			b.NewTab("https://duckduckgo.com")
			return true
		}

		// Ctrl+W - Fechar aba
		if ctrlPressed && keyVal == gdk.KEY_w {
			log.Println("⌨️  Ctrl+W - Fechar aba")
			b.CloseCurrentTab()
			return true
		}

		// Alt+← - Voltar
		if altPressed && keyVal == gdk.KEY_Left {
			log.Println("⌨️  Alt+← - Voltar")
			b.GoBack()
			return true
		}

		// Alt+→ - Avançar
		if altPressed && keyVal == gdk.KEY_Right {
			log.Println("⌨️  Alt+→ - Avançar")
			b.GoForward()
			return true
		}

		// F5 - Recarregar
		if keyVal == gdk.KEY_F5 {
			log.Println("⌨️  F5 - Recarregar")
			b.Reload()
			return true
		}

		// Ctrl+R - Recarregar
		if ctrlPressed && keyVal == gdk.KEY_r {
			log.Println("⌨️  Ctrl+R - Recarregar")
			b.Reload()
			return true
		}

		// Ctrl+L - Focar URL
		if ctrlPressed && keyVal == gdk.KEY_l {
			log.Println("⌨️  Ctrl+L - Focar URL")
			b.urlEntry.GrabFocus()
			b.urlEntry.SelectRegion(0, -1)
			return true
		}

		return false
	})
}

// NewTab cria uma nova aba
func (b *Browser) NewTab(url string) {
	log.Printf("📑 Criando aba com WebView...")

	// Criar WebView
	webView, err := NewWebView()
	if err != nil {
		log.Printf("❌ Erro ao criar WebView: %v", err)
		return
	}

	// Criar container scrollable
	scrolled, err := gtk.ScrolledWindowNew(nil, nil)
	if err != nil {
		log.Printf("❌ Erro ao criar scrolled window: %v", err)
		return
	}
	scrolled.Add(webView.Widget())

	// Criar label da aba (começa com "Carregando...")
	tabLabel, _ := gtk.LabelNew("Carregando...")

	// Criar Tab
	tab := &Tab{
		webView: webView,
		label:   tabLabel,
	}

	// Adicionar ao slice
	b.tabs = append(b.tabs, tab)
	tabIndex := len(b.tabs) - 1

	// Adicionar aba ao notebook
	pageNum := b.notebook.AppendPage(scrolled, tabLabel)
	b.notebook.SetCurrentPage(pageNum)

	// Conectar sinal notify::title para atualizar label
	webView.widget.Connect("notify::title", func() {
		title := webView.GetTitle()
		uri := webView.GetURI()
		
		// Usar título se disponível, senão usar URI
		if title != "" && title != "about:blank" {
			tabLabel.SetText(title)
		} else if uri != "" && uri != "about:blank" {
			tabLabel.SetText(uri)
		}
		
		// Atualizar URL entry se for a aba atual
		if b.notebook.GetCurrentPage() == pageNum {
			if uri != "" {
				b.urlEntry.SetText(uri)
			}
		}
	})

	// Conectar sinal notify::uri para atualizar quando URL mudar
	webView.widget.Connect("notify::uri", func() {
		uri := webView.GetURI()
		if uri != "" && uri != "about:blank" {
			// Se não tiver título ainda, mostrar URI
			if webView.GetTitle() == "" {
				tabLabel.SetText(uri)
			}
			
			// Atualizar URL entry se for a aba atual
			if b.notebook.GetCurrentPage() == pageNum {
				b.urlEntry.SetText(uri)
			}
		}
	})

	// Carregar URL
	webView.LoadURI(url)

	b.window.ShowAll()

	log.Printf("✅ Aba %d criada - Carregando: %s", tabIndex+1, url)
}

// getCurrentWebView retorna o WebView da aba atual
func (b *Browser) getCurrentWebView() *WebView {
	pageNum := b.notebook.GetCurrentPage()
	if pageNum < 0 || pageNum >= len(b.tabs) {
		return nil
	}
	return b.tabs[pageNum].webView
}

// Navigate navega na aba atual
func (b *Browser) Navigate(url string) {
	webView := b.getCurrentWebView()
	if webView == nil {
		return
	}

	// Adicionar https:// se necessário
	if len(url) > 0 && url[0] != 'h' {
		url = "https://" + url
	}

	log.Printf("🌐 Navegando para: %s", url)
	webView.LoadURI(url)
}

// GoBack volta na história da aba atual
func (b *Browser) GoBack() {
	webView := b.getCurrentWebView()
	if webView != nil {
		log.Println("← Voltando")
		webView.GoBack()
	}
}

// GoForward avança na história da aba atual
func (b *Browser) GoForward() {
	webView := b.getCurrentWebView()
	if webView != nil {
		log.Println("→ Avançando")
		webView.GoForward()
	}
}

// Reload recarrega a aba atual
func (b *Browser) Reload() {
	webView := b.getCurrentWebView()
	if webView != nil {
		log.Println("⟳ Recarregando")
		webView.Reload()
	}
}

// CloseCurrentTab fecha a aba atual
func (b *Browser) CloseCurrentTab() {
	pageNum := b.notebook.GetCurrentPage()
	nPages := b.notebook.GetNPages()

	if nPages <= 1 {
		log.Println("⚠️  Não é possível fechar a última aba")
		return
	}

	if pageNum >= 0 && pageNum < len(b.tabs) {
		// Remover do slice
		b.tabs = append(b.tabs[:pageNum], b.tabs[pageNum+1:]...)
		
		// Remover da notebook
		b.notebook.RemovePage(pageNum)
		
		log.Printf("🗑️  Aba fechada (restam: %d)", b.notebook.GetNPages())
	}
}

// Show mostra a janela
func (b *Browser) Show() {
	b.window.ShowAll()
}
