package main

/*
#cgo pkg-config: webkit2gtk-4.0
#include <webkit2/webkit2.h>
#include <stdlib.h>

// Obter informações do download
static const char* get_download_uri(WebKitDownload* download) {
    WebKitURIRequest* request = webkit_download_get_request(download);
    return webkit_uri_request_get_uri(request);
}

static const char* get_download_destination(WebKitDownload* download) {
    return webkit_download_get_destination(download);
}

static guint64 get_download_received(WebKitDownload* download) {
    return webkit_download_get_received_data_length(download);
}

static guint64 get_download_total(WebKitDownload* download) {
    WebKitURIResponse* response = webkit_download_get_response(download);
    if (response) {
        return webkit_uri_response_get_content_length(response);
    }
    return 0;
}

static void set_download_destination(WebKitDownload* download, const char* dest) {
    webkit_download_set_destination(download, dest);
}

static void cancel_download(WebKitDownload* download) {
    webkit_download_cancel(download);
}

// Callback para progresso
static void on_download_received_data(WebKitDownload* download, guint64 data_length, gpointer user_data) {
    // Callback será tratado em Go
}

// Callback para conclusão
static void on_download_finished(WebKitDownload* download, gpointer user_data) {
    // Callback será tratado em Go
}

// Callback para falha
static void on_download_failed(WebKitDownload* download, GError* error, gpointer user_data) {
    // Callback será tratado em Go
}

// Conectar sinais do download
static void connect_download_signals(WebKitDownload* download, gpointer user_data) {
    g_signal_connect(download, "received-data", G_CALLBACK(on_download_received_data), user_data);
    g_signal_connect(download, "finished", G_CALLBACK(on_download_finished), user_data);
    g_signal_connect(download, "failed", G_CALLBACK(on_download_failed), user_data);
}

// Variável global para armazenar callback Go
extern void goDownloadStartedCallback(WebKitDownload* download);

// Callback C que chama função Go
static void download_started_callback(WebKitWebContext* context, WebKitDownload* download, gpointer user_data) {
    goDownloadStartedCallback(download);
}

// Conectar handler de downloads ao contexto
static void connect_download_handler_to_context(WebKitWebContext* context) {
    g_signal_connect(context, "download-started", G_CALLBACK(download_started_callback), NULL);
}

*/
import "C"
import (
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"sync"
	"unsafe"
	
	"github.com/gotk3/gotk3/gtk"
)

// Variável global para armazenar o handler
var globalDownloadHandler *DownloadHandler

// goDownloadStartedCallback é chamada pelo C quando um download inicia
//export goDownloadStartedCallback
func goDownloadStartedCallback(download *C.WebKitDownload) {
	if globalDownloadHandler != nil {
		log.Println("📥 Download detectado via callback C!")
		globalDownloadHandler.HandleDownload(unsafe.Pointer(download))
	}
}

// DownloadHandler gerencia downloads do WebKit
type DownloadHandler struct {
	browser         *Browser
	downloadManager *DownloadManager
	activeDownloads map[uintptr]*DownloadItem
	mu              sync.RWMutex
}

// NewDownloadHandler cria um novo handler de downloads
func NewDownloadHandler(browser *Browser, dm *DownloadManager) *DownloadHandler {
	return &DownloadHandler{
		browser:         browser,
		downloadManager: dm,
		activeDownloads: make(map[uintptr]*DownloadItem),
	}
}

// HandleDownload processa um novo download
func (dh *DownloadHandler) HandleDownload(download unsafe.Pointer) {
	cDownload := (*C.WebKitDownload)(download)
	
	// Obter URI do download
	cURI := C.get_download_uri(cDownload)
	uri := C.GoString(cURI)
	
	// Extrair nome do arquivo da URI
	filename := extractFilename(uri)
	if filename == "" {
		filename = "download"
	}
	
	log.Printf("📥 Novo download: %s", filename)
	
	// Mostrar diálogo para escolher onde salvar
	done := make(chan string, 1)
	
	glib.IdleAdd(func() bool {
		destination := dh.showSaveDialog(filename)
		done <- destination
		return false
	})
	
	destination := <-done
	
	// Se usuário cancelou, cancelar download
	if destination == "" {
		log.Println("🚫 Download cancelado pelo usuário")
		C.cancel_download(cDownload)
		return
	}
	
	// Gerar ID único para o download
	downloadID := fmt.Sprintf("%p", download)
	
	// Adicionar ao gerenciador com destino escolhido
	item := dh.downloadManager.AddDownloadWithDestination(downloadID, uri, filepath.Base(destination), destination)
	
	// Configurar destino
	destinationURI := fmt.Sprintf("file://%s", destination)
	cDestination := C.CString(destinationURI)
	defer C.free(unsafe.Pointer(cDestination))
	C.set_download_destination(cDownload, cDestination)
	
	// Guardar referência
	dh.mu.Lock()
	dh.activeDownloads[uintptr(download)] = item
	dh.mu.Unlock()
	
	log.Printf("✅ Download configurado: %s → %s", filename, destination)
}

// showSaveDialog mostra diálogo para escolher onde salvar
func (dh *DownloadHandler) showSaveDialog(suggestedFilename string) string {
	dialog, err := gtk.FileChooserDialogNewWith2Buttons(
		"Salvar arquivo",
		dh.browser.window,
		gtk.FILE_CHOOSER_ACTION_SAVE,
		"Cancelar", gtk.RESPONSE_CANCEL,
		"Salvar", gtk.RESPONSE_ACCEPT,
	)
	if err != nil {
		log.Printf("❌ Erro ao criar diálogo: %v", err)
		// Fallback para pasta padrão
		return dh.downloadManager.GetUniqueFilename(suggestedFilename)
	}
	defer dialog.Destroy()
	
	// Configurar diálogo
	dialog.SetCurrentName(suggestedFilename)
	dialog.SetCurrentFolder(dh.downloadManager.GetDownloadPath())
	dialog.SetDoOverwriteConfirmation(true)
	
	// Mostrar e aguardar resposta
	response := dialog.Run()
	
	if response == gtk.RESPONSE_ACCEPT {
		filename := dialog.GetFilename()
		return filename
	}
	
	return ""
}

// CancelDownload cancela um download ativo
func (dh *DownloadHandler) CancelDownload(downloadID string) {
	dh.mu.RLock()
	defer dh.mu.RUnlock()
	
	for ptr, item := range dh.activeDownloads {
		if item.ID == downloadID {
			cDownload := (*C.WebKitDownload)(unsafe.Pointer(ptr))
			C.cancel_download(cDownload)
			log.Printf("🚫 Download cancelado: %s", item.Filename)
			return
		}
	}
}

// extractFilename extrai o nome do arquivo de uma URI
func extractFilename(uri string) string {
	// Remover query string
	if idx := strings.Index(uri, "?"); idx != -1 {
		uri = uri[:idx]
	}
	
	// Remover fragment
	if idx := strings.Index(uri, "#"); idx != -1 {
		uri = uri[:idx]
	}
	
	// Extrair nome do arquivo
	parts := strings.Split(uri, "/")
	if len(parts) > 0 {
		filename := parts[len(parts)-1]
		if filename != "" {
			// Decodificar URL encoding
			filename = strings.ReplaceAll(filename, "%20", " ")
			filename = strings.ReplaceAll(filename, "%2B", "+")
			return filename
		}
	}
	
	return ""
}

// setupGlobalDownloadHandler configura o handler global de downloads no WebContext
func (b *Browser) setupGlobalDownloadHandler(webContext *WebContext) {
	if b.downloadManager == nil {
		return
	}
	
	// Criar handler se não existir
	if b.downloadHandler == nil {
		b.downloadHandler = NewDownloadHandler(b, b.downloadManager)
	}
	
	// Armazenar na variável global para o callback C
	globalDownloadHandler = b.downloadHandler
	
	// Conectar handler ao WebContext via C
	C.connect_download_handler_to_context(webContext.cContext)
	
	log.Println("✅ Handler global de downloads configurado no WebContext")
}

// setupDownloadHandler mantido para compatibilidade (mas não é mais usado)
func (b *Browser) setupDownloadHandler(webView *WebView) {
	// Não faz nada - downloads são gerenciados globalmente pelo WebContext
	log.Println("⚠️  setupDownloadHandler chamado (ignorado - usando handler global)")
}
