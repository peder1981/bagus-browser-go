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

static const char* get_download_suggested_filename(WebKitDownload* download) {
    WebKitURIResponse* response = webkit_download_get_response(download);
    if (response) {
        const char* suggested = webkit_uri_response_get_suggested_filename(response);
        if (suggested && suggested[0] != '\0') {
            return suggested;
        }
    }
    return NULL;
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

// Callbacks Go exportados
extern void goDownloadReceivedData(WebKitDownload* download, guint64 data_length);
extern void goDownloadFinished(WebKitDownload* download);
extern void goDownloadFailed(WebKitDownload* download);

// Callback para progresso
static void on_download_received_data(WebKitDownload* download, guint64 data_length, gpointer user_data) {
    goDownloadReceivedData(download, data_length);
}

// Callback para conclusão
static void on_download_finished(WebKitDownload* download, gpointer user_data) {
    goDownloadFinished(download);
}

// Callback para falha
static void on_download_failed(WebKitDownload* download, GError* error, gpointer user_data) {
    goDownloadFailed(download);
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
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"
	"unsafe"
	
	"github.com/gotk3/gotk3/gtk"
)

// Variável global para armazenar o handler
var globalDownloadHandler *DownloadHandler

// goDownloadStartedCallback é chamada pelo C quando um download inicia
//export goDownloadStartedCallback
func goDownloadStartedCallback(download *C.WebKitDownload) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("❌ PANIC no callback de download: %v", r)
		}
	}()
	
	log.Println("📥 Download detectado via callback C!")
	
	if globalDownloadHandler == nil {
		log.Println("❌ globalDownloadHandler é NIL!")
		return
	}
	
	log.Println("✅ globalDownloadHandler existe, chamando HandleDownload...")
	globalDownloadHandler.HandleDownload(unsafe.Pointer(download))
	log.Println("✅ HandleDownload retornou")
}

// goDownloadReceivedData é chamada quando há progresso no download
//export goDownloadReceivedData
func goDownloadReceivedData(download *C.WebKitDownload, dataLength C.guint64) {
	if globalDownloadHandler == nil {
		return
	}
	
	ptr := uintptr(unsafe.Pointer(download))
	globalDownloadHandler.mu.RLock()
	item, exists := globalDownloadHandler.activeDownloads[ptr]
	globalDownloadHandler.mu.RUnlock()
	
	if !exists {
		return
	}
	
	// Obter tamanho total e recebido
	received := uint64(C.get_download_received(download))
	total := uint64(C.get_download_total(download))
	
	// Atualizar progresso
	item.UpdateProgress(received, total)
}

// goDownloadFinished é chamada quando o download termina
//export goDownloadFinished
func goDownloadFinished(download *C.WebKitDownload) {
	if globalDownloadHandler == nil {
		return
	}
	
	ptr := uintptr(unsafe.Pointer(download))
	globalDownloadHandler.mu.RLock()
	item, exists := globalDownloadHandler.activeDownloads[ptr]
	globalDownloadHandler.mu.RUnlock()
	
	if exists {
		item.Complete()
		log.Printf("✅ Download concluído: %s", item.Filename)
		
		// Notificação desktop
		globalDownloadHandler.showNotification(
			"Download Concluído",
			fmt.Sprintf("✅ %s foi baixado com sucesso!", item.Filename),
		)
		
		// Remover da lista de ativos
		globalDownloadHandler.mu.Lock()
		delete(globalDownloadHandler.activeDownloads, ptr)
		globalDownloadHandler.mu.Unlock()
	}
}

// goDownloadFailed é chamada quando o download falha
//export goDownloadFailed
func goDownloadFailed(download *C.WebKitDownload) {
	if globalDownloadHandler == nil {
		return
	}
	
	ptr := uintptr(unsafe.Pointer(download))
	globalDownloadHandler.mu.RLock()
	item, exists := globalDownloadHandler.activeDownloads[ptr]
	globalDownloadHandler.mu.RUnlock()
	
	if exists {
		item.Fail()
		log.Printf("❌ Download falhou: %s", item.Filename)
		
		// Notificação desktop
		globalDownloadHandler.showNotification(
			"Download Falhou",
			fmt.Sprintf("❌ Erro ao baixar %s", item.Filename),
		)
		
		// Remover da lista de ativos
		globalDownloadHandler.mu.Lock()
		delete(globalDownloadHandler.activeDownloads, ptr)
		globalDownloadHandler.mu.Unlock()
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
	log.Println("🚀 HandleDownload INICIADO")
	cDownload := (*C.WebKitDownload)(download)
	
	// Obter URI do download
	cURI := C.get_download_uri(cDownload)
	uri := C.GoString(cURI)
	log.Printf("🔗 URI do download: %s", uri)
	
	// Tentar obter nome sugerido do WebKit (Content-Disposition)
	var filename string
	cSuggested := C.get_download_suggested_filename(cDownload)
	if cSuggested != nil {
		filename = C.GoString(cSuggested)
		log.Printf("✨ Nome sugerido pelo WebKit: %s", filename)
	}
	
	// Se não tiver nome sugerido, extrair da URI
	if filename == "" {
		filename = extractFilename(uri)
		log.Printf("📝 Nome extraído da URI: %s", filename)
	}
	
	// Fallback para nome genérico
	if filename == "" || len(filename) > 100 {
		timestamp := fmt.Sprintf("%d", time.Now().Unix())
		filename = fmt.Sprintf("download_%s", timestamp)
		log.Printf("⚠️  Usando nome genérico: %s", filename)
	}
	
	log.Printf("📥 Novo download: %s", filename)
	
	// Gerar caminho de destino automaticamente
	destination := dh.downloadManager.GetUniqueFilename(filename)
	destinationURI := fmt.Sprintf("file://%s", destination)
	
	// Configurar destino IMEDIATAMENTE (antes de retornar do callback)
	cDestination := C.CString(destinationURI)
	defer C.free(unsafe.Pointer(cDestination))
	C.set_download_destination(cDownload, cDestination)
	
	// Gerar ID único para o download
	downloadID := fmt.Sprintf("%p", download)
	
	// Adicionar ao gerenciador
	item := dh.downloadManager.AddDownloadWithDestination(downloadID, uri, filepath.Base(destination), destination)
	
	// Guardar referência
	dh.mu.Lock()
	dh.activeDownloads[uintptr(download)] = item
	dh.mu.Unlock()
	
	// Conectar sinais
	C.connect_download_signals(cDownload, C.gpointer(download))
	
	log.Printf("✅ Download iniciado: %s → %s", filename, destination)
	
	// Notificação de início
	dh.showNotification(
		"Download Iniciado",
		fmt.Sprintf("📥 Baixando %s...", filename),
	)
}

// showNotification mostra notificação desktop
func (dh *DownloadHandler) showNotification(title, message string) {
	// Usar notify-send do sistema (disponível na maioria das distros Linux)
	go func() {
		// Executar notify-send de forma assíncrona
		// Ignora erros se notify-send não estiver disponível
		_ = exec.Command("notify-send", "-i", "download", title, message).Run()
	}()
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
