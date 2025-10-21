package browser

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
*/
import "C"
import (
	"fmt"
	"unsafe"

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
