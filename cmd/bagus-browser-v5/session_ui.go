package main

/*
#cgo pkg-config: webkit2gtk-4.0 gtk+-3.0
#include <webkit2/webkit2.h>
#include <gtk/gtk.h>
#include <stdlib.h>

// Obter número de abas
static int get_tab_count(GtkNotebook *notebook) {
    return gtk_notebook_get_n_pages(notebook);
}

// Obter URL de uma aba
static const char* get_tab_url(GtkNotebook *notebook, int index) {
    GtkWidget *page = gtk_notebook_get_nth_page(notebook, index);
    if (!page || !WEBKIT_IS_WEB_VIEW(page)) {
        return NULL;
    }
    return webkit_web_view_get_uri(WEBKIT_WEB_VIEW(page));
}

// Obter título de uma aba
static const char* get_tab_title(GtkNotebook *notebook, int index) {
    GtkWidget *page = gtk_notebook_get_nth_page(notebook, index);
    if (!page || !WEBKIT_IS_WEB_VIEW(page)) {
        return NULL;
    }
    const char *title = webkit_web_view_get_title(WEBKIT_WEB_VIEW(page));
    return title ? title : "Nova Aba";
}

// Obter índice da aba ativa
static int get_current_tab_index(GtkNotebook *notebook) {
    return gtk_notebook_get_current_page(notebook);
}
*/
import "C"

import (
	"log"
	"unsafe"
)

//export SaveCurrentSession
func SaveCurrentSession(notebook unsafe.Pointer) {
	if globalSessionManager == nil {
		log.Println("⚠️  SessionManager não inicializado")
		return
	}

	cNotebook := (*C.GtkNotebook)(notebook)
	tabCount := int(C.get_tab_count(cNotebook))
	
	if tabCount == 0 {
		log.Println("💾 Nenhuma aba para salvar")
		return
	}

	var tabs []SessionTab
	currentIndex := int(C.get_current_tab_index(cNotebook))

	for i := 0; i < tabCount; i++ {
		cURL := C.get_tab_url(cNotebook, C.int(i))
		cTitle := C.get_tab_title(cNotebook, C.int(i))

		if cURL == nil {
			continue
		}

		url := C.GoString(cURL)
		title := C.GoString(cTitle)

		// Não salvar abas vazias ou about:blank
		if url == "" || url == "about:blank" {
			continue
		}

		tab := SessionTab{
			URL:    url,
			Title:  title,
			Active: i == currentIndex,
		}
		tabs = append(tabs, tab)
	}

	if len(tabs) == 0 {
		log.Println("💾 Nenhuma aba válida para salvar")
		return
	}

	if err := globalSessionManager.Save(tabs); err != nil {
		log.Printf("❌ Erro ao salvar sessão: %v", err)
	} else {
		log.Printf("💾 Sessão salva: %d abas", len(tabs))
	}
}

//export RestoreSession
func RestoreSession() int {
	if globalSessionManager == nil {
		log.Println("⚠️  SessionManager não inicializado")
		return 0
	}

	session, err := globalSessionManager.Load()
	if err != nil {
		log.Printf("⚠️  Erro ao carregar sessão: %v", err)
		return 0
	}

	if len(session.Tabs) == 0 {
		log.Println("📂 Nenhuma sessão anterior encontrada")
		return 0
	}

	log.Printf("📂 Restaurando sessão: %d abas", len(session.Tabs))
	
	// Retornar número de abas para restaurar
	// As abas serão criadas pelo código C
	return len(session.Tabs)
}

//export GetRestoredTabURL
func GetRestoredTabURL(index int) *C.char {
	if globalSessionManager == nil {
		return nil
	}

	session, err := globalSessionManager.Load()
	if err != nil || index >= len(session.Tabs) {
		return nil
	}

	return C.CString(session.Tabs[index].URL)
}

//export GetRestoredTabTitle
func GetRestoredTabTitle(index int) *C.char {
	if globalSessionManager == nil {
		return nil
	}

	session, err := globalSessionManager.Load()
	if err != nil || index >= len(session.Tabs) {
		return nil
	}

	return C.CString(session.Tabs[index].Title)
}

//export IsRestoredTabActive
func IsRestoredTabActive(index int) C.int {
	if globalSessionManager == nil {
		return 0
	}

	session, err := globalSessionManager.Load()
	if err != nil || index >= len(session.Tabs) {
		return 0
	}

	if session.Tabs[index].Active {
		return 1
	}
	return 0
}
