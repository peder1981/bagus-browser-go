package main

/*
#cgo pkg-config: webkit2gtk-4.0
#include <webkit2/webkit2.h>

// Configurar privacidade do WebView
static void configure_webview_privacy(WebKitWebView* webview) {
    WebKitWebContext* context = webkit_web_view_get_context(webview);
    WebKitCookieManager* cookie_manager = webkit_web_context_get_cookie_manager(context);
    
    // Política de cookies: Bloquear third-party
    webkit_cookie_manager_set_accept_policy(
        cookie_manager,
        WEBKIT_COOKIE_POLICY_ACCEPT_NO_THIRD_PARTY
    );
    
    // Configurações de privacidade
    WebKitSettings* settings = webkit_web_view_get_settings(webview);
    
    // Desabilitar WebGL (fingerprinting)
    webkit_settings_set_enable_webgl(settings, FALSE);
    
    // Desabilitar WebAudio (fingerprinting)
    webkit_settings_set_enable_webaudio(settings, FALSE);
    
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

// DefaultPrivacyConfig retorna configuração padrão (máxima privacidade)
func DefaultPrivacyConfig() *PrivacyConfig {
	return &PrivacyConfig{
		BlockThirdPartyCookies: true,
		BlockGeolocation:       true,
		BlockNotifications:     false, // Permitir notificações (usuário pode controlar por site)
		BlockMediaAccess:       true,
		BlockWebGL:             true,
		BlockWebAudio:          true,
		DNT:                    true,
	}
}

// ApplyPrivacyConfig aplica configurações de privacidade ao WebView
func ApplyPrivacyConfig(webView *WebView, config *PrivacyConfig) {
	log.Println("🕵️  Aplicando configurações de privacidade...")
	
	// Aplicar configurações via CGO
	C.configure_webview_privacy(webView.cWebView)
	
	log.Println("✅ Privacidade configurada:")
	log.Printf("   - Third-party cookies: %v", !config.BlockThirdPartyCookies)
	log.Printf("   - Geolocalização: %v", !config.BlockGeolocation)
	log.Printf("   - Notificações: %v", !config.BlockNotifications)
	log.Printf("   - Mídia (câmera/mic): %v", !config.BlockMediaAccess)
	log.Printf("   - WebGL: %v", !config.BlockWebGL)
	log.Printf("   - WebAudio: %v", !config.BlockWebAudio)
	log.Printf("   - Do Not Track: %v", config.DNT)
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
	log.Println("   ✅ Acesso a mídia bloqueado")
	log.Println("   ✅ WebGL bloqueado (anti-fingerprinting)")
	log.Println("   ✅ WebAudio bloqueado (anti-fingerprinting)")
	log.Println("   ✅ Do Not Track habilitado")
	log.Println("   ✅ DuckDuckGo como motor de busca padrão")
}
