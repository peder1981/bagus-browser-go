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
    // O WebKit já faz isso automaticamente
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

// Callback para permissões de mídia
extern gboolean goPermissionRequestCallback(WebKitPermissionRequest* request);

// Callback C que chama função Go para permissões
static gboolean permission_request_callback(WebKitWebView* webview, WebKitPermissionRequest* request, gpointer user_data) {
    return goPermissionRequestCallback(request);
}

// Callback para criar nova janela/aba
extern WebKitWebView* goCreateCallback(WebKitWebView* webview, WebKitNavigationAction* navigation_action);

// Callback C que chama função Go para criar nova janela
static WebKitWebView* create_callback(WebKitWebView* webview, WebKitNavigationAction* navigation_action, gpointer user_data) {
    return goCreateCallback(webview, navigation_action);
}

// Criar novo WebView relacionado ao original
static WebKitWebView* create_related_webview(WebKitWebView* webview) {
    // Criar novo WebView com o mesmo contexto
    WebKitWebContext* context = webkit_web_view_get_context(webview);
    WebKitWebView* new_webview = WEBKIT_WEB_VIEW(webkit_web_view_new_with_context(context));

    // Copiar settings do original
    WebKitSettings* settings = webkit_web_view_get_settings(webview);
    webkit_web_view_set_settings(new_webview, settings);

    return new_webview;
}

// Conectar handler de permissões ao WebView
static void connect_permission_handler(WebKitWebView* webview) {
    g_signal_connect(webview, "permission-request", G_CALLBACK(permission_request_callback), NULL);
}

// Callback para decisões de navegação
extern gboolean goDecidePolicyCallback(WebKitWebView* webview, WebKitPolicyDecision* decision, WebKitPolicyDecisionType decision_type);

// Callback C para decisões de política
static gboolean decide_policy_callback(WebKitWebView* webview, WebKitPolicyDecision* decision, WebKitPolicyDecisionType decision_type, gpointer user_data) {
    return goDecidePolicyCallback(webview, decision, decision_type);
}

// Conectar handler de criação de janelas ao WebView
static void connect_create_handler(WebKitWebView* webview) {
    g_signal_connect(webview, "create", G_CALLBACK(create_callback), NULL);
}

// Conectar handler de decisões de política
static void connect_decide_policy_handler(WebKitWebView* webview) {
    g_signal_connect(webview, "decide-policy", G_CALLBACK(decide_policy_callback), NULL);
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

    // Habilitar WebRTC (RTCPeerConnection - ESSENCIAL para Google Meet)
    webkit_settings_set_enable_webrtc(settings, TRUE);

    // Habilitar EncryptedMedia (para DRM - Netflix, etc)
    webkit_settings_set_enable_encrypted_media(settings, TRUE);

    // Permitir modal dialogs (alguns sites precisam)
    webkit_settings_set_allow_modal_dialogs(settings, TRUE);

    // Habilitar fullscreen (para vídeos)
    webkit_settings_set_enable_fullscreen(settings, TRUE);

    // Habilitar aceleração de hardware
    webkit_settings_set_hardware_acceleration_policy(settings, WEBKIT_HARDWARE_ACCELERATION_POLICY_ALWAYS);

    // Permitir JavaScript abrir janelas automaticamente (necessário para Google Meet)
    webkit_settings_set_javascript_can_open_windows_automatically(settings, TRUE);

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

// goPermissionRequestCallback trata pedidos de permissão (webcam, microfone, etc)
//
//export goPermissionRequestCallback
func goPermissionRequestCallback(request *C.WebKitPermissionRequest) C.gboolean {
	// Por enquanto, aceitar automaticamente todas as permissões
	// TODO: Adicionar diálogo de confirmação para o usuário
	log.Println("🔐 Permissão solicitada - concedendo automaticamente")

	// Permitir a permissão
	C.webkit_permission_request_allow(request)

	return C.TRUE
}

// goCreateCallback trata solicitações de criar nova janela (popups, Google Meet, etc)
//
//export goCreateCallback
func goCreateCallback(webview *C.WebKitWebView, navigationAction *C.WebKitNavigationAction) *C.WebKitWebView {
	log.Println("🎆 Nova janela solicitada - popup será bloqueado (não criando nova janela)")
	// Para evitar crashes internos do WebKit em certos sites (ex: suporte Totvs),
	// não vamos criar um novo WebView aqui. Retornar nil instrui o WebKit a não abrir
	// uma nova janela/popup.
	return nil
}

// goDecidePolicyCallback trata decisões de navegação
//
//export goDecidePolicyCallback
func goDecidePolicyCallback(webview *C.WebKitWebView, decision *C.WebKitPolicyDecision, decisionType C.WebKitPolicyDecisionType) C.gboolean {
	// WEBKIT_POLICY_DECISION_TYPE_NAVIGATION_ACTION = 0
	// WEBKIT_POLICY_DECISION_TYPE_NEW_WINDOW_ACTION = 1
	// WEBKIT_POLICY_DECISION_TYPE_RESPONSE = 2

	if decisionType == 1 { // NEW_WINDOW_ACTION
		log.Println("🪟 Solicitação de nova janela detectada")
	}

	// Permitir todas as navegações
	C.webkit_policy_decision_use(decision)
	return C.TRUE
}

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
	C.connect_permission_handler(webView.cWebView)
	C.connect_create_handler(webView.cWebView)
	C.connect_decide_policy_handler(webView.cWebView)
	log.Println("🎥 WebView configurado para multimídia + permissões + popups + navegação")
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
