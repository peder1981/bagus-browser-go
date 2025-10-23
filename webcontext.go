package main

/*
#cgo pkg-config: webkit2gtk-4.0
#include <webkit2/webkit2.h>
#include <stdlib.h>

// Obter contexto web padrão
static WebKitWebContext* get_default_context() {
    return webkit_web_context_get_default();
}

// Configurar gerenciamento de recursos para liberar quando aba fechar
static void setup_resource_management(WebKitWebView* webview) {
    // Conectar sinal para limpar recursos quando página for destruída
    // O WebKit já faz isso automaticamente, mas podemos forçar
    g_signal_connect(webview, "close", G_CALLBACK(gtk_widget_destroy), NULL);
}

// Parar carregamento de recursos
static void stop_loading(WebKitWebView* webview) {
    webkit_web_view_stop_loading(webview);
}

*/
import "C"
import (
	"log"
	"unsafe"
)

// WebContext gerencia contexto global do WebKit
type WebContext struct {
	cContext *C.WebKitWebContext
}

// GetDefaultWebContext retorna contexto web padrão
func GetDefaultWebContext() *WebContext {
	return &WebContext{
		cContext: C.get_default_context(),
	}
}

// Initialize inicializa contexto com configurações
func (wc *WebContext) Initialize(config *Config) {
	log.Println("🌐 Inicializando WebContext...")
	
	// Configurar persistência de cookies
	setupCookiePersistence(unsafe.Pointer(wc.cContext), config)
	
	// Configurar cache de vídeos
	setupVideoCache(unsafe.Pointer(wc.cContext), config)
	
	log.Println("✅ WebContext inicializado")
}

// setupResourceManagement configura gerenciamento de recursos para um WebView
func setupResourceManagement(webView *WebView) {
	C.setup_resource_management(webView.cWebView)
}

// stopLoading para carregamento de recursos
func stopLoading(webView *WebView) {
	C.stop_loading(webView.cWebView)
	log.Println("🛑 Carregamento parado")
}
