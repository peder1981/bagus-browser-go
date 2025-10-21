package main

/*
#cgo pkg-config: webkit2gtk-4.0
#include <webkit2/webkit2.h>

// Configurar WebView com segurança
static void configure_webview_security(WebKitWebView* webview) {
    WebKitSettings* settings = webkit_web_view_get_settings(webview);
    
    // JavaScript necessário para sites modernos
    webkit_settings_set_enable_javascript(settings, TRUE);
    
    // Desabilitar plugins (segurança)
    webkit_settings_set_enable_plugins(settings, FALSE);
    
    // Desabilitar Java (segurança)
    webkit_settings_set_enable_java(settings, FALSE);
    
    // User agent customizado
    webkit_settings_set_user_agent(settings, 
        "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Bagus/4.0");
    
    // Habilitar proteções
    webkit_settings_set_enable_dns_prefetching(settings, FALSE);
    webkit_settings_set_enable_page_cache(settings, FALSE);
}

static WebKitWebView* create_webview() {
    WebKitWebView* webview = WEBKIT_WEB_VIEW(webkit_web_view_new());
    configure_webview_security(webview);
    return webview;
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

static void set_zoom_level(WebKitWebView* webview, gdouble zoom) {
    webkit_web_view_set_zoom_level(webview, zoom);
}

static gdouble get_zoom_level(WebKitWebView* webview) {
    return webkit_web_view_get_zoom_level(webview);
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

// Funções de busca
static WebKitFindController* get_find_controller(WebKitWebView* webview) {
    return webkit_web_view_get_find_controller(webview);
}

static void find_text(WebKitWebView* webview, const char* text, gboolean case_sensitive) {
    WebKitFindController* controller = webkit_web_view_get_find_controller(webview);
    guint32 options = WEBKIT_FIND_OPTIONS_WRAP_AROUND;
    if (!case_sensitive) {
        options |= WEBKIT_FIND_OPTIONS_CASE_INSENSITIVE;
    }
    webkit_find_controller_search(controller, text, options, G_MAXUINT);
}

static void find_next(WebKitWebView* webview) {
    WebKitFindController* controller = webkit_web_view_get_find_controller(webview);
    webkit_find_controller_search_next(controller);
}

static void find_previous(WebKitWebView* webview) {
    WebKitFindController* controller = webkit_web_view_get_find_controller(webview);
    webkit_find_controller_search_previous(controller);
}

static void find_finish(WebKitWebView* webview) {
    WebKitFindController* controller = webkit_web_view_get_find_controller(webview);
    webkit_find_controller_search_finish(controller);
}

// Função para conectar sinal de download (será chamada do Go)
// O handler real será implementado em Go
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

// SetZoomLevel define o nível de zoom
func (w *WebView) SetZoomLevel(zoom float64) {
	C.set_zoom_level(w.cWebView, C.gdouble(zoom))
}

// GetZoomLevel retorna o nível de zoom atual
func (w *WebView) GetZoomLevel() float64 {
	return float64(C.get_zoom_level(w.cWebView))
}

// ZoomIn aumenta o zoom
func (w *WebView) ZoomIn() {
	currentZoom := w.GetZoomLevel()
	w.SetZoomLevel(currentZoom + 0.1)
}

// ZoomOut diminui o zoom
func (w *WebView) ZoomOut() {
	currentZoom := w.GetZoomLevel()
	if currentZoom > 0.2 {
		w.SetZoomLevel(currentZoom - 0.1)
	}
}

// ZoomReset reseta o zoom para 100%
func (w *WebView) ZoomReset() {
	w.SetZoomLevel(1.0)
}

// FindText busca texto na página
func (w *WebView) FindText(text string, caseSensitive bool) {
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))
	
	var cCaseSensitive C.gboolean
	if caseSensitive {
		cCaseSensitive = C.gboolean(1)
	} else {
		cCaseSensitive = C.gboolean(0)
	}
	
	C.find_text(w.cWebView, cText, cCaseSensitive)
}

// FindNext busca próxima ocorrência
func (w *WebView) FindNext() {
	C.find_next(w.cWebView)
}

// FindPrevious busca ocorrência anterior
func (w *WebView) FindPrevious() {
	C.find_previous(w.cWebView)
}

// FindFinish finaliza busca
func (w *WebView) FindFinish() {
	C.find_finish(w.cWebView)
}

// Tab representa uma aba com WebView e label
type Tab struct {
	webView *WebView
	label   *gtk.Label
}

// Browser representa o navegador
type Browser struct {
	window          *gtk.Window
	notebook        *gtk.Notebook
	urlEntry        *gtk.Entry
	tabs            []*Tab
	validator       *URLValidator
	privacyManager  *PrivacyManager
	bookmarkManager *BookmarkManager
	downloadManager *DownloadManager
	findBar         *gtk.Box
	findEntry       *gtk.Entry
	findBarVisible  bool
}

func main() {
	runtime.LockOSThread()
	gtk.Init(nil)

	log.Println("🌐 Iniciando Bagus Browser POC - WebKit CGO...")

	browser := NewBrowser()
	browser.Show()

	log.Println("✅ Browser iniciado com WebView!")
	log.Println("📝 Navegue para diferentes sites")
	log.Println("⌨️  Atalhos:")
	log.Println("   Ctrl+T (nova aba), Ctrl+W (fechar)")
	log.Println("   Alt+← (voltar), Alt+→ (avançar), F5 (recarregar)")
	log.Println("   Ctrl++ (zoom in), Ctrl+- (zoom out), Ctrl+0 (reset zoom)")
	log.Println("   Ctrl+F (buscar), F3 (próximo), Shift+F3 (anterior)")
	log.Println("   Ctrl+L (focar URL)")

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

	// Criar CryptoManager
	crypto, err := NewCryptoManager("")
	if err != nil {
		log.Fatal("Erro ao criar crypto:", err)
	}
	
	// Criar BookmarkManager
	bookmarkManager, err := NewBookmarkManager(crypto)
	if err != nil {
		log.Printf("⚠️  Erro ao criar bookmark manager: %v", err)
	}
	
	// Criar DownloadManager
	downloadManager, err := NewDownloadManager()
	if err != nil {
		log.Printf("⚠️  Erro ao criar download manager: %v", err)
	} else {
		log.Printf("📁 Downloads: %s", downloadManager.GetDownloadPath())
	}
	
	browser := &Browser{
		window:          win,
		notebook:        notebook,
		urlEntry:        urlEntry,
		tabs:            make([]*Tab, 0),
		validator:       NewURLValidator(),
		privacyManager:  NewPrivacyManager(),
		bookmarkManager: bookmarkManager,
		downloadManager: downloadManager,
	}
	
	// Logar informações de privacidade
	browser.privacyManager.LogPrivacyInfo()

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

		// Ctrl++ ou Ctrl+= - Zoom in
		if ctrlPressed && (keyVal == gdk.KEY_plus || keyVal == gdk.KEY_equal) {
			log.Println("⌨️  Ctrl++ - Aumentar zoom")
			b.ZoomIn()
			return true
		}

		// Ctrl+- - Zoom out
		if ctrlPressed && keyVal == gdk.KEY_minus {
			log.Println("⌨️  Ctrl+- - Diminuir zoom")
			b.ZoomOut()
			return true
		}

		// Ctrl+0 - Reset zoom
		if ctrlPressed && keyVal == gdk.KEY_0 {
			log.Println("⌨️  Ctrl+0 - Resetar zoom")
			b.ZoomReset()
			return true
		}

		// Ctrl+F - Buscar na página
		if ctrlPressed && keyVal == gdk.KEY_f {
			log.Println("⌨️  Ctrl+F - Buscar na página")
			b.ShowFindBar()
			return true
		}

		// Esc - Fechar barra de busca
		if keyVal == gdk.KEY_Escape && b.findBarVisible {
			log.Println("⌨️  Esc - Fechar busca")
			b.HideFindBar()
			return true
		}

		// F3 - Próximo resultado
		if keyVal == gdk.KEY_F3 {
			log.Println("⌨️  F3 - Próximo resultado")
			b.FindNext()
			return true
		}

		// Shift+F3 - Resultado anterior
		shiftPressed := (state & uint(gdk.SHIFT_MASK)) != 0
		if shiftPressed && keyVal == gdk.KEY_F3 {
			log.Println("⌨️  Shift+F3 - Resultado anterior")
			b.FindPrevious()
			return true
		}

		// Ctrl+D - Adicionar favorito
		if ctrlPressed && keyVal == gdk.KEY_d {
			log.Println("⌨️  Ctrl+D - Adicionar favorito")
			b.AddBookmark()
			return true
		}

		// Ctrl+Shift+B - Gerenciar favoritos
		if ctrlPressed && shiftPressed && keyVal == gdk.KEY_b {
			log.Println("⌨️  Ctrl+Shift+B - Gerenciar favoritos")
			b.ShowBookmarks()
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
	
	// Aplicar configurações de privacidade
	ApplyPrivacyConfig(webView, b.privacyManager.GetConfig())

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
func (b *Browser) Navigate(input string) {
	webView := b.getCurrentWebView()
	if webView == nil {
		return
	}

	// Sanitizar input
	input = SanitizeInput(input)
	
	// Validar e processar URL
	validURL, err := b.validator.ValidateAndSanitize(input)
	if err != nil {
		log.Printf("⚠️  Erro ao validar URL: %v", err)
		// Mostrar erro ao usuário (pode adicionar dialog depois)
		return
	}

	log.Printf("🌐 Navegando para: %s", validURL)
	webView.LoadURI(validURL)
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

// ZoomIn aumenta o zoom da aba atual
func (b *Browser) ZoomIn() {
	webView := b.getCurrentWebView()
	if webView != nil {
		webView.ZoomIn()
		zoom := webView.GetZoomLevel()
		log.Printf("🔍 Zoom: %.0f%%", zoom*100)
	}
}

// ZoomOut diminui o zoom da aba atual
func (b *Browser) ZoomOut() {
	webView := b.getCurrentWebView()
	if webView != nil {
		webView.ZoomOut()
		zoom := webView.GetZoomLevel()
		log.Printf("🔍 Zoom: %.0f%%", zoom*100)
	}
}

// ZoomReset reseta o zoom da aba atual
func (b *Browser) ZoomReset() {
	webView := b.getCurrentWebView()
	if webView != nil {
		webView.ZoomReset()
		log.Println("🔍 Zoom: 100%")
	}
}

// ShowFindBar mostra a barra de busca
func (b *Browser) ShowFindBar() {
	// Por enquanto, usar um dialog simples
	dialog, _ := gtk.DialogNew()
	dialog.SetTitle("Buscar na Página")
	dialog.SetTransientFor(b.window)
	dialog.SetModal(true)
	dialog.SetDefaultSize(400, 100)
	
	// Entry para busca
	entry, _ := gtk.EntryNew()
	entry.SetPlaceholderText("Digite o texto para buscar...")
	
	// Botões
	dialog.AddButton("Próximo", gtk.RESPONSE_ACCEPT)
	dialog.AddButton("Anterior", gtk.RESPONSE_REJECT)
	dialog.AddButton("Fechar", gtk.RESPONSE_CLOSE)
	
	// Adicionar entry ao dialog
	box, _ := dialog.GetContentArea()
	box.SetMarginTop(10)
	box.SetMarginBottom(10)
	box.SetMarginStart(10)
	box.SetMarginEnd(10)
	box.PackStart(entry, true, true, 5)
	
	dialog.ShowAll()
	
	// Handler para buscar ao digitar
	entry.Connect("changed", func() {
		text, _ := entry.GetText()
		if text != "" {
			webView := b.getCurrentWebView()
			if webView != nil {
				webView.FindText(text, false)
			}
		}
	})
	
	// Handler para resposta
	dialog.Connect("response", func(d *gtk.Dialog, response gtk.ResponseType) {
		webView := b.getCurrentWebView()
		
		if webView != nil {
			switch response {
			case gtk.RESPONSE_ACCEPT: // Próximo
				webView.FindNext()
			case gtk.RESPONSE_REJECT: // Anterior
				webView.FindPrevious()
			case gtk.RESPONSE_CLOSE: // Fechar
				webView.FindFinish()
				dialog.Close()
			}
		}
	})
	
	dialog.Run()
	dialog.Destroy()
}

// HideFindBar esconde a barra de busca
func (b *Browser) HideFindBar() {
	webView := b.getCurrentWebView()
	if webView != nil {
		webView.FindFinish()
	}
	b.findBarVisible = false
}

// FindNext busca próxima ocorrência
func (b *Browser) FindNext() {
	webView := b.getCurrentWebView()
	if webView != nil {
		webView.FindNext()
	}
}

// FindPrevious busca ocorrência anterior
func (b *Browser) FindPrevious() {
	webView := b.getCurrentWebView()
	if webView != nil {
		webView.FindPrevious()
	}
}

// AddBookmark adiciona página atual aos favoritos
func (b *Browser) AddBookmark() {
	if b.bookmarkManager == nil {
		log.Println("⚠️  Bookmark manager não disponível")
		return
	}
	
	webView := b.getCurrentWebView()
	if webView == nil {
		return
	}
	
	url := webView.GetURI()
	title := webView.GetTitle()
	
	if title == "" {
		title = url
	}
	
	// Verificar se já existe
	if b.bookmarkManager.Exists(url) {
		log.Println("⚠️  Favorito já existe")
		
		// Mostrar mensagem
		dialog := gtk.MessageDialogNew(b.window, gtk.DIALOG_MODAL,
			gtk.MESSAGE_INFO, gtk.BUTTONS_OK, "Este site já está nos favoritos!")
		dialog.Run()
		dialog.Destroy()
		return
	}
	
	// Adicionar
	if err := b.bookmarkManager.Add(title, url); err != nil {
		log.Printf("❌ Erro ao adicionar favorito: %v", err)
		return
	}
	
	// Mostrar confirmação
	dialog := gtk.MessageDialogNew(b.window, gtk.DIALOG_MODAL,
		gtk.MESSAGE_INFO, gtk.BUTTONS_OK,
		fmt.Sprintf("⭐ Favorito adicionado!\n\n%s", title))
	dialog.Run()
	dialog.Destroy()
}

// ShowBookmarks mostra gerenciador de favoritos
func (b *Browser) ShowBookmarks() {
	if b.bookmarkManager == nil {
		log.Println("⚠️  Bookmark manager não disponível")
		return
	}
	
	dialog, _ := gtk.DialogNew()
	dialog.SetTitle("Favoritos")
	dialog.SetTransientFor(b.window)
	dialog.SetModal(true)
	dialog.SetDefaultSize(600, 400)
	
	// Criar lista
	scrolled, _ := gtk.ScrolledWindowNew(nil, nil)
	listBox, _ := gtk.ListBoxNew()
	scrolled.Add(listBox)
	
	// Adicionar favoritos
	bookmarks := b.bookmarkManager.GetAll()
	for _, bookmark := range bookmarks {
		row, _ := gtk.ListBoxRowNew()
		box, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 10)
		
		// Título
		label, _ := gtk.LabelNew(bookmark.Title)
		label.SetHAlign(gtk.ALIGN_START)
		box.PackStart(label, true, true, 5)
		
		// Botão abrir
		btnOpen, _ := gtk.ButtonNewWithLabel("Abrir")
		url := bookmark.URL
		btnOpen.Connect("clicked", func() {
			b.NewTab(url)
			dialog.Close()
		})
		box.PackStart(btnOpen, false, false, 5)
		
		// Botão remover
		btnRemove, _ := gtk.ButtonNewWithLabel("Remover")
		btnRemove.Connect("clicked", func() {
			if err := b.bookmarkManager.Remove(url); err != nil {
				log.Printf("❌ Erro ao remover: %v", err)
			}
			row.Destroy()
		})
		box.PackStart(btnRemove, false, false, 5)
		
		row.Add(box)
		listBox.Add(row)
	}
	
	// Adicionar ao dialog
	box, _ := dialog.GetContentArea()
	box.PackStart(scrolled, true, true, 10)
	
	// Botão fechar
	dialog.AddButton("Fechar", gtk.RESPONSE_CLOSE)
	
	dialog.ShowAll()
	dialog.Run()
	dialog.Destroy()
}

// Show mostra a janela
func (b *Browser) Show() {
	b.window.ShowAll()
}
