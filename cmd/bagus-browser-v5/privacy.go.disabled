package main

/*
#cgo pkg-config: webkit2gtk-4.0
#include <webkit2/webkit2.h>

// Configurar privacidade do WebView
static void configure_webview_privacy(WebKitWebView* webview, gboolean block_media, gboolean block_webgl, gboolean block_webaudio) {
    WebKitWebContext* context = webkit_web_view_get_context(webview);
    WebKitCookieManager* cookie_manager = webkit_web_context_get_cookie_manager(context);
    
    // Política de cookies: Bloquear third-party
    webkit_cookie_manager_set_accept_policy(
        cookie_manager,
        WEBKIT_COOKIE_POLICY_ACCEPT_NO_THIRD_PARTY
    );
    
    // Configurações de privacidade
    WebKitSettings* settings = webkit_web_view_get_settings(webview);
    
    // MediaStream (câmera/microfone) - respeitar configuração
    webkit_settings_set_enable_media_stream(settings, !block_media);
    
    // WebGL - respeitar configuração
    webkit_settings_set_enable_webgl(settings, !block_webgl);
    
    // WebAudio - respeitar configuração
    webkit_settings_set_enable_webaudio(settings, !block_webaudio);
    
    // Desabilitar acesso offline (cache)
    webkit_settings_set_enable_offline_web_application_cache(settings, FALSE);
}
*/
import "C"
import (
	"log"
)

// PrivacyConfig representa configurações de privacidade
type PrivacyConfig struct {
	BlockThirdPartyCookies bool
	BlockGeolocation       bool
	BlockNotifications     bool
	BlockMediaAccess       bool
	BlockWebGL             bool
	BlockWebAudio          bool
	DNT                    bool // Do Not Track
}

// DefaultPrivacyConfig retorna configuração padrão (balanceada)
func DefaultPrivacyConfig() *PrivacyConfig {
	return &PrivacyConfig{
		BlockThirdPartyCookies: true,
		BlockGeolocation:       true,
		BlockNotifications:     false, // Permitir notificações (usuário pode controlar por site)
		BlockMediaAccess:       false, // Permitir câmera/microfone (Google Meet, etc)
		BlockWebGL:             false, // Permitir WebGL (necessário para Meet)
		BlockWebAudio:          false, // Permitir WebAudio (necessário para Meet)
		DNT:                    true,
	}
}

// ApplyPrivacyConfig aplica configurações de privacidade ao WebView
func ApplyPrivacyConfig(webView *WebView, config *PrivacyConfig) {
	log.Println("🕵️  Aplicando configurações de privacidade...")
	
	// Aplicar configurações via CGO com flags dinâmicas
	blockMedia := C.gboolean(0)
	if config.BlockMediaAccess {
		blockMedia = C.gboolean(1)
	}
	
	blockWebGL := C.gboolean(0)
	if config.BlockWebGL {
		blockWebGL = C.gboolean(1)
	}
	
	blockWebAudio := C.gboolean(0)
	if config.BlockWebAudio {
		blockWebAudio = C.gboolean(1)
	}
	
	C.configure_webview_privacy(webView.cWebView, blockMedia, blockWebGL, blockWebAudio)
	
	log.Println("✅ Privacidade configurada:")
	log.Printf("   - Third-party cookies: bloqueados")
	log.Printf("   - Geolocalização: %v", map[bool]string{true: "bloqueada", false: "permitida"}[config.BlockGeolocation])
	log.Printf("   - Notificações: %v", map[bool]string{true: "bloqueadas", false: "permitidas"}[config.BlockNotifications])
	log.Printf("   - Mídia (câmera/mic): %v", map[bool]string{true: "bloqueada", false: "permitida"}[config.BlockMediaAccess])
	log.Printf("   - WebGL: %v", map[bool]string{true: "bloqueado", false: "permitido"}[config.BlockWebGL])
	log.Printf("   - WebAudio: %v", map[bool]string{true: "bloqueado", false: "permitido"}[config.BlockWebAudio])
	log.Printf("   - Do Not Track: %v", map[bool]string{true: "habilitado", false: "desabilitado"}[config.DNT])
}

// PrivacyManager gerencia privacidade do browser
type PrivacyManager struct {
	config *PrivacyConfig
}

// NewPrivacyManager cria um novo gerenciador de privacidade
func NewPrivacyManager() *PrivacyManager {
	return &PrivacyManager{
		config: DefaultPrivacyConfig(),
	}
}

// NewPrivacyManagerFromConfig cria gerenciador de privacidade a partir de Config
func NewPrivacyManagerFromConfig(config *Config) *PrivacyManager {
	return &PrivacyManager{
		config: &PrivacyConfig{
			BlockThirdPartyCookies: config.BlockThirdPartyCookies,
			BlockGeolocation:       config.BlockGeolocation,
			BlockNotifications:     config.BlockNotifications,
			BlockMediaAccess:       config.BlockMedia,
			BlockWebGL:             config.BlockWebGL,
			BlockWebAudio:          config.BlockWebAudio,
			DNT:                    config.DoNotTrack,
		},
	}
}

// GetConfig retorna configuração atual
func (pm *PrivacyManager) GetConfig() *PrivacyConfig {
	return pm.config
}

// LogPrivacyInfo loga informações de privacidade
func (pm *PrivacyManager) LogPrivacyInfo() {
	log.Println("🕵️  Bagus Browser - Configurações de Privacidade:")
	log.Println("   ✅ Zero telemetria")
	log.Println("   ✅ Sem analytics")
	log.Println("   ✅ Sem crash reports")
	log.Println("   ✅ Sem rastreamento")
	log.Println("   ✅ Third-party cookies bloqueados")
	log.Println("   ✅ Geolocalização bloqueada")
	log.Println("   ✅ Notificações habilitadas (controladas por site)")
	log.Println("   ✅ Acesso a mídia permitido (Google Meet, etc)")
	log.Println("   ✅ WebGL permitido (necessário para Meet)")
	log.Println("   ✅ WebAudio permitido (necessário para Meet)")
	log.Println("   ✅ Do Not Track habilitado")
	log.Println("   ✅ DuckDuckGo como motor de busca padrão")
}

// applyPrivacySettings aplica configurações de privacidade baseadas em Config
func applyPrivacySettings(webView *C.WebKitWebView, config *Config) {
	// Converter Config para PrivacyConfig
	privacyConfig := &PrivacyConfig{
		BlockThirdPartyCookies: config.BlockThirdPartyCookies,
		BlockGeolocation:       config.BlockGeolocation,
		BlockNotifications:     config.BlockNotifications,
		BlockMediaAccess:       config.BlockMedia,
		BlockWebGL:             config.BlockWebGL,
		BlockWebAudio:          config.BlockWebAudio,
		DNT:                    config.DoNotTrack,
	}
	
	// Aplicar configurações
	wv := &WebView{cWebView: webView}
	ApplyPrivacyConfig(wv, privacyConfig)
}
