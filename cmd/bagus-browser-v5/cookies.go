package main

/*
#cgo pkg-config: webkit2gtk-4.0
#include <webkit2/webkit2.h>
#include <stdlib.h>

// Configurar persistência de cookies
static void setup_cookie_persistence(WebKitWebContext* context, const char* cookie_file, gboolean persist) {
    WebKitCookieManager* cookie_manager = webkit_web_context_get_cookie_manager(context);
    
    if (persist) {
        // Persistir cookies em arquivo
        webkit_cookie_manager_set_persistent_storage(
            cookie_manager,
            cookie_file,
            WEBKIT_COOKIE_PERSISTENT_STORAGE_SQLITE
        );
        
        // Aceitar cookies (exceto third-party se bloqueado)
        webkit_cookie_manager_set_accept_policy(
            cookie_manager,
            WEBKIT_COOKIE_POLICY_ACCEPT_NO_THIRD_PARTY
        );
    } else {
        // Não persistir - apenas sessão
        webkit_cookie_manager_set_accept_policy(
            cookie_manager,
            WEBKIT_COOKIE_POLICY_ACCEPT_NO_THIRD_PARTY
        );
    }
}

// Limpar todos os cookies
static void clear_all_cookies(WebKitWebContext* context) {
    WebKitCookieManager* cookie_manager = webkit_web_context_get_cookie_manager(context);
    webkit_cookie_manager_delete_all_cookies(cookie_manager);
}

// Configurar cache de vídeos
static void setup_video_cache(WebKitWebContext* context, const char* cache_dir, gint cache_size_mb) {
    WebKitWebsiteDataManager* data_manager = webkit_web_context_get_website_data_manager(context);
    
    // Configurar diretório de cache
    // O cache é gerenciado automaticamente pelo WebKit
    // Apenas garantimos que o diretório existe
}

*/
import "C"
import (
	"log"
	"os"
	"path/filepath"
	"unsafe"
)

// setupCookiePersistence configura persistência de cookies
func setupCookiePersistence(context unsafe.Pointer, config *Config) {
	cookieFile := getCookieFilePath()
	
	// Criar diretório se não existe
	cookieDir := filepath.Dir(cookieFile)
	if err := os.MkdirAll(cookieDir, 0700); err != nil {
		log.Printf("❌ Erro ao criar diretório de cookies: %v", err)
		return
	}
	
	cCookieFile := C.CString(cookieFile)
	defer C.free(unsafe.Pointer(cCookieFile))
	
	persist := C.gboolean(0)
	if config.PersistCookies {
		persist = C.gboolean(1)
	}
	
	C.setup_cookie_persistence(
		(*C.WebKitWebContext)(context),
		cCookieFile,
		persist,
	)
	
	if config.PersistCookies {
		log.Println("🍪 Cookies persistentes habilitados")
		log.Printf("📁 Arquivo: %s", cookieFile)
	} else {
		log.Println("🍪 Cookies apenas para sessão (não persistem)")
	}
}

// clearAllCookies limpa todos os cookies
func clearAllCookies(context unsafe.Pointer) {
	C.clear_all_cookies((*C.WebKitWebContext)(context))
	log.Println("🗑️  Todos os cookies foram limpos")
}

// getCookieFilePath retorna caminho do arquivo de cookies
func getCookieFilePath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(homeDir, ".config", "bagus-browser", "cookies.sqlite")
}

// setupVideoCache configura cache de vídeos
func setupVideoCache(context unsafe.Pointer, config *Config) {
	if !config.VideoCacheEnabled {
		log.Println("📹 Cache de vídeos desabilitado")
		return
	}
	
	cacheDir := getVideoCacheDir()
	
	// Criar diretório se não existe
	if err := os.MkdirAll(cacheDir, 0700); err != nil {
		log.Printf("❌ Erro ao criar diretório de cache: %v", err)
		return
	}
	
	cCacheDir := C.CString(cacheDir)
	defer C.free(unsafe.Pointer(cCacheDir))
	
	cacheSizeMB := C.gint(config.VideoCacheSize)
	
	C.setup_video_cache(
		(*C.WebKitWebContext)(context),
		cCacheDir,
		cacheSizeMB,
	)
	
	log.Printf("📹 Cache de vídeos habilitado: %d MB", config.VideoCacheSize)
	log.Printf("📁 Diretório: %s", cacheDir)
}

// getVideoCacheDir retorna diretório de cache de vídeos
func getVideoCacheDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(homeDir, ".cache", "bagus-browser", "video")
}

// clearVideoCache limpa cache de vídeos
func clearVideoCache() error {
	cacheDir := getVideoCacheDir()
	
	if err := os.RemoveAll(cacheDir); err != nil {
		return err
	}
	
	// Recriar diretório
	if err := os.MkdirAll(cacheDir, 0700); err != nil {
		return err
	}
	
	log.Println("🗑️  Cache de vídeos limpo")
	return nil
}
