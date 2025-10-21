package main

/*
#cgo pkg-config: webkit2gtk-4.0
#include <webkit2/webkit2.h>
#include <gtk/gtk.h>

// Wrapper para criar WebView
static WebKitWebView* create_webview() {
    return WEBKIT_WEB_VIEW(webkit_web_view_new());
}

// Wrapper para carregar URI
static void load_uri(WebKitWebView* webview, const char* uri) {
    webkit_web_view_load_uri(webview, uri);
}

// Wrapper para obter URI atual
static const char* get_uri(WebKitWebView* webview) {
    return webkit_web_view_get_uri(webview);
}

// Wrapper para obter título
static const char* get_title(WebKitWebView* webview) {
    return webkit_web_view_get_title(webview);
}

// Wrapper para voltar
static void go_back(WebKitWebView* webview) {
    webkit_web_view_go_back(webview);
}

// Wrapper para avançar
static void go_forward(WebKitWebView* webview) {
    webkit_web_view_go_forward(webview);
}

// Wrapper para recarregar
static void reload_page(WebKitWebView* webview) {
    webkit_web_view_reload(webview);
}

// Wrapper para verificar se pode voltar
static gboolean can_go_back(WebKitWebView* webview) {
    return webkit_web_view_can_go_back(webview);
}

// Wrapper para verificar se pode avançar
static gboolean can_go_forward(WebKitWebView* webview) {
    return webkit_web_view_can_go_forward(webview);
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

// GoBack volta na história
func (w *WebView) GoBack() {
	C.go_back(w.cWebView)
}

// GoForward avança na história
func (w *WebView) GoForward() {
	C.go_forward(w.cWebView)
}

// Reload recarrega a página
func (w *WebView) Reload() {
	C.reload_page(w.cWebView)
}

// CanGoBack verifica se pode voltar
func (w *WebView) CanGoBack() bool {
	return C.can_go_back(w.cWebView) != 0
}

// CanGoForward verifica se pode avançar
func (w *WebView) CanGoForward() bool {
	return C.can_go_forward(w.cWebView) != 0
}

// Browser representa o navegador
type Browser struct {
	window   *gtk.Window
	notebook *gtk.Notebook
	urlEntry *gtk.Entry
	tabCount int
}

// Tab representa uma aba
type Tab struct {
	webView   *WebView
	container *gtk.ScrolledWindow
	label     *gtk.Label
}

func main() {
	runtime.LockOSThread()
	gtk.Init(nil)

	log.Println("🌐 Iniciando Bagus Browser POC - WebKit CGO...")

	browser := NewBrowser()
	browser.Show()

	log.Println("✅ Browser iniciado com WebView!")
	log.Println("📝 Teste: Navegue para diferentes sites")

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
		tabCount: 0,
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
		shiftPressed := (state & uint(gdk.SHIFT_MASK)) != 0
		altPressed := (state & uint(gdk.MOD1_MASK)) != 0

		// Ctrl+T - Nova aba
		if ctrlPressed && !shiftPressed && keyVal == gdk.KEY_t {
			log.Println("⌨️  Ctrl+T - Nova aba")
			b.NewTab("https://duckduckgo.com")
			return true
		}

		// Ctrl+W - Fechar aba
		if ctrlPressed && !shiftPressed && keyVal == gdk.KEY_w {
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
	b.tabCount++

	log.Printf("📑 Criando aba %d com WebView...", b.tabCount)

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

	// Criar label da aba
	tabLabel, _ := gtk.LabelNew(fmt.Sprintf("Aba %d", b.tabCount))

	// Adicionar aba
	pageNum := b.notebook.AppendPage(scrolled, tabLabel)
	b.notebook.SetCurrentPage(pageNum)

	// Carregar URL
	webView.LoadURI(url)

	b.window.ShowAll()

	log.Printf("✅ Aba %d criada - Carregando: %s", b.tabCount, url)
}

// Navigate navega na aba atual
func (b *Browser) Navigate(url string) {
	pageNum := b.notebook.GetCurrentPage()
	if pageNum < 0 {
		return
	}

	page, _ := b.notebook.GetNthPage(pageNum)
	scrolled := &gtk.ScrolledWindow{gtk.Bin{gtk.Container{gtk.Widget{
		glib.InitiallyUnowned{page.Object},
	}}}}

	children := scrolled.GetChildren()
	if children.Length() == 0 {
		return
	}

	// Pegar WebView (primeiro filho do scrolled)
	webViewWidget := children.NthData(0).(*gtk.Widget)
	
	// Converter para WebView
	cWebView := (*C.WebKitWebView)(unsafe.Pointer(webViewWidget.GObject))
	
	// Adicionar https:// se necessário
	if len(url) > 0 && url[0] != 'h' {
		url = "https://" + url
	}

	log.Printf("🌐 Navegando para: %s", url)
	
	cURI := C.CString(url)
	defer C.free(unsafe.Pointer(cURI))
	C.load_uri(cWebView, cURI)
}

// GoBack volta na história da aba atual
func (b *Browser) GoBack() {
	pageNum := b.notebook.GetCurrentPage()
	if pageNum < 0 {
		return
	}

	page, _ := b.notebook.GetNthPage(pageNum)
	scrolled := &gtk.ScrolledWindow{gtk.Bin{gtk.Container{gtk.Widget{
		glib.InitiallyUnowned{page.Object},
	}}}}

	children := scrolled.GetChildren()
	if children.Length() == 0 {
		return
	}

	webViewWidget := children.NthData(0).(*gtk.Widget)
	cWebView := (*C.WebKitWebView)(unsafe.Pointer(webViewWidget.GObject))

	if C.can_go_back(cWebView) != 0 {
		log.Println("← Voltando")
		C.go_back(cWebView)
	}
}

// GoForward avança na história da aba atual
func (b *Browser) GoForward() {
	pageNum := b.notebook.GetCurrentPage()
	if pageNum < 0 {
		return
	}

	page, _ := b.notebook.GetNthPage(pageNum)
	scrolled := &gtk.ScrolledWindow{gtk.Bin{gtk.Container{gtk.Widget{
		glib.InitiallyUnowned{page.Object},
	}}}}

	children := scrolled.GetChildren()
	if children.Length() == 0 {
		return
	}

	webViewWidget := children.NthData(0).(*gtk.Widget)
	cWebView := (*C.WebKitWebView)(unsafe.Pointer(webViewWidget.GObject))

	if C.can_go_forward(cWebView) != 0 {
		log.Println("→ Avançando")
		C.go_forward(cWebView)
	}
}

// Reload recarrega a aba atual
func (b *Browser) Reload() {
	pageNum := b.notebook.GetCurrentPage()
	if pageNum < 0 {
		return
	}

	page, _ := b.notebook.GetNthPage(pageNum)
	scrolled := &gtk.ScrolledWindow{gtk.Bin{gtk.Container{gtk.Widget{
		glib.InitiallyUnowned{page.Object},
	}}}}

	children := scrolled.GetChildren()
	if children.Length() == 0 {
		return
	}

	webViewWidget := children.NthData(0).(*gtk.Widget)
	cWebView := (*C.WebKitWebView)(unsafe.Pointer(webViewWidget.GObject))

	log.Println("⟳ Recarregando")
	C.reload_page(cWebView)
}

// CloseCurrentTab fecha a aba atual
func (b *Browser) CloseCurrentTab() {
	pageNum := b.notebook.GetCurrentPage()
	nPages := b.notebook.GetNPages()

	if nPages <= 1 {
		log.Println("⚠️  Não é possível fechar a última aba")
		return
	}

	if pageNum >= 0 {
		b.notebook.RemovePage(pageNum)
		log.Printf("🗑️  Aba fechada (restam: %d)", b.notebook.GetNPages())
	}
}

// Show mostra a janela
func (b *Browser) Show() {
	b.window.ShowAll()
}
