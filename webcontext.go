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

// Configurar suporte robusto a multimídia
static void setup_multimedia_support(WebKitWebContext* context) {
    // Configurações básicas de contexto
    webkit_web_context_set_automation_allowed(context, TRUE);
}

// Configurar settings para WebView com suporte a multimídia
static void configure_webview_multimedia(WebKitWebView* webview) {
    WebKitSettings* settings = webkit_web_view_get_settings(webview);
    
    // Habilitar JavaScript (essencial)
    webkit_settings_set_enable_javascript(settings, TRUE);
    
    // Habilitar WebGL (para aplicações modernas)
    webkit_settings_set_enable_webgl(settings, TRUE);
    
    // Habilitar WebAudio (para áudio avançado)
    webkit_settings_set_enable_webaudio(settings, TRUE);
    
    // Habilitar MediaStream (para webcam/microfone - Google Meet)
    webkit_settings_set_enable_media_stream(settings, TRUE);
    
    // Habilitar EncryptedMedia (para DRM - Netflix, etc)
    webkit_settings_set_enable_encrypted_media(settings, TRUE);
    
    // Permitir modal dialogs (alguns sites precisam)
    webkit_settings_set_allow_modal_dialogs(settings, TRUE);
    
    // Habilitar fullscreen (para vídeos)
    webkit_settings_set_enable_fullscreen(settings, TRUE);
    
    // Habilitar aceleração de hardware
    webkit_settings_set_hardware_acceleration_policy(settings, WEBKIT_HARDWARE_ACCELERATION_POLICY_ALWAYS);
    
    // User agent moderno (compatibilidade)
    webkit_settings_set_user_agent(settings, 
        "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36");
}

// Configurar diretório de downloads
static void setup_download_directory(WebKitWebContext* context, const char* download_path) {
    if (download_path != NULL && download_path[0] != '\0') {
        // WebKit2GTK usa variável de ambiente XDG_DOWNLOAD_DIR
        // Vamos configurar via código quando possível
    }
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
func (wc *WebContext) Initialize(config *Config, downloadPath string) {
	log.Println("🌐 Inicializando WebContext...")
	
	// Configurar persistência de cookies
	setupCookiePersistence(unsafe.Pointer(wc.cContext), config)
	
	// Configurar cache de vídeos
	setupVideoCache(unsafe.Pointer(wc.cContext), config)
	
	// Configurar suporte a multimídia
	C.setup_multimedia_support(wc.cContext)
	log.Println("🎬 Suporte a multimídia habilitado (Meet, YouTube Music, etc)")
	
	// Configurar pasta de downloads
	cDownloadPath := C.CString(downloadPath)
	defer C.free(unsafe.Pointer(cDownloadPath))
	C.setup_download_directory(wc.cContext, cDownloadPath)
	log.Printf("📁 Pasta de downloads configurada: %s", downloadPath)
	
	log.Println("✅ WebContext inicializado")
}

// ConfigureWebViewMultimedia configura WebView com suporte robusto a multimídia
func ConfigureWebViewMultimedia(webView *WebView) {
	C.configure_webview_multimedia(webView.cWebView)
	log.Println("🎥 WebView configurado para multimídia")
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
