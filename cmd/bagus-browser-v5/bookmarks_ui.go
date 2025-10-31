package main

/*
#cgo pkg-config: gtk+-3.0
#include <gtk/gtk.h>
#include <stdlib.h>

// Callback para adicionar favorito
static void on_bookmark_add_clicked(GtkButton *button, gpointer data);

// Helper para adicionar item ao store
static void add_bookmark_to_store(GtkListStore *store, const char *title, const char *url) {
    GtkTreeIter iter;
    gtk_list_store_append(store, &iter);
    gtk_list_store_set(store, &iter, 0, title, 1, url, -1);
}

// Criar diálogo para adicionar favorito
static void show_add_bookmark_dialog(GtkWindow *parent, const char *current_url, const char *current_title) {
    GtkWidget *dialog = gtk_dialog_new_with_buttons(
        "⭐ Adicionar Favorito",
        parent,
        GTK_DIALOG_MODAL,
        "Cancelar", GTK_RESPONSE_CANCEL,
        "Adicionar", GTK_RESPONSE_OK,
        NULL
    );
    
    gtk_window_set_default_size(GTK_WINDOW(dialog), 500, 200);
    
    GtkWidget *content_area = gtk_dialog_get_content_area(GTK_DIALOG(dialog));
    gtk_container_set_border_width(GTK_CONTAINER(content_area), 15);
    
    GtkWidget *grid = gtk_grid_new();
    gtk_grid_set_row_spacing(GTK_GRID(grid), 10);
    gtk_grid_set_column_spacing(GTK_GRID(grid), 10);
    gtk_box_pack_start(GTK_BOX(content_area), grid, TRUE, TRUE, 0);
    
    // Título
    GtkWidget *title_label = gtk_label_new("Título:");
    gtk_widget_set_halign(title_label, GTK_ALIGN_END);
    gtk_grid_attach(GTK_GRID(grid), title_label, 0, 0, 1, 1);
    
    GtkWidget *title_entry = gtk_entry_new();
    gtk_entry_set_text(GTK_ENTRY(title_entry), current_title ? current_title : "");
    gtk_widget_set_hexpand(title_entry, TRUE);
    gtk_grid_attach(GTK_GRID(grid), title_entry, 1, 0, 1, 1);
    
    // URL
    GtkWidget *url_label = gtk_label_new("URL:");
    gtk_widget_set_halign(url_label, GTK_ALIGN_END);
    gtk_grid_attach(GTK_GRID(grid), url_label, 0, 1, 1, 1);
    
    GtkWidget *url_entry = gtk_entry_new();
    gtk_entry_set_text(GTK_ENTRY(url_entry), current_url ? current_url : "");
    gtk_widget_set_hexpand(url_entry, TRUE);
    gtk_grid_attach(GTK_GRID(grid), url_entry, 1, 1, 1, 1);
    
    gtk_widget_show_all(dialog);
    
    gint response = gtk_dialog_run(GTK_DIALOG(dialog));
    
    if (response == GTK_RESPONSE_OK) {
        const char *title = gtk_entry_get_text(GTK_ENTRY(title_entry));
        const char *url = gtk_entry_get_text(GTK_ENTRY(url_entry));
        
        // Chamar função Go para adicionar
        AddBookmarkFromC((char*)title, (char*)url);
    }
    
    gtk_widget_destroy(dialog);
}

// Criar diálogo de gerenciamento de favoritos
static void show_bookmarks_dialog(GtkWindow *parent) {
    GtkWidget *dialog = gtk_dialog_new_with_buttons(
        "📚 Gerenciar Favoritos",
        parent,
        GTK_DIALOG_MODAL,
        "Fechar", GTK_RESPONSE_CLOSE,
        NULL
    );
    
    gtk_window_set_default_size(GTK_WINDOW(dialog), 700, 500);
    
    GtkWidget *content_area = gtk_dialog_get_content_area(GTK_DIALOG(dialog));
    gtk_container_set_border_width(GTK_CONTAINER(content_area), 10);
    
    // Scrolled window
    GtkWidget *scrolled = gtk_scrolled_window_new(NULL, NULL);
    gtk_scrolled_window_set_policy(GTK_SCROLLED_WINDOW(scrolled), 
                                   GTK_POLICY_AUTOMATIC, GTK_POLICY_AUTOMATIC);
    gtk_box_pack_start(GTK_BOX(content_area), scrolled, TRUE, TRUE, 0);
    
    // TreeView para lista de favoritos
    GtkListStore *store = gtk_list_store_new(2, G_TYPE_STRING, G_TYPE_STRING);
    
    // Preencher com favoritos (será chamado do Go)
    PopulateBookmarksStore(store);
    
    GtkWidget *tree_view = gtk_tree_view_new_with_model(GTK_TREE_MODEL(store));
    
    // Coluna Título
    GtkCellRenderer *renderer = gtk_cell_renderer_text_new();
    GtkTreeViewColumn *column = gtk_tree_view_column_new_with_attributes(
        "Título", renderer, "text", 0, NULL);
    gtk_tree_view_column_set_expand(column, TRUE);
    gtk_tree_view_append_column(GTK_TREE_VIEW(tree_view), column);
    
    // Coluna URL
    renderer = gtk_cell_renderer_text_new();
    column = gtk_tree_view_column_new_with_attributes(
        "URL", renderer, "text", 1, NULL);
    gtk_tree_view_column_set_expand(column, TRUE);
    gtk_tree_view_append_column(GTK_TREE_VIEW(tree_view), column);
    
    gtk_container_add(GTK_CONTAINER(scrolled), tree_view);
    
    // Botões de ação
    GtkWidget *button_box = gtk_box_new(GTK_ORIENTATION_HORIZONTAL, 5);
    gtk_box_pack_start(GTK_BOX(content_area), button_box, FALSE, FALSE, 5);
    
    GtkWidget *btn_open = gtk_button_new_with_label("🌐 Abrir");
    GtkWidget *btn_delete = gtk_button_new_with_label("🗑️  Remover");
    GtkWidget *btn_export = gtk_button_new_with_label("📤 Exportar");
    GtkWidget *btn_import = gtk_button_new_with_label("📥 Importar");
    
    gtk_box_pack_start(GTK_BOX(button_box), btn_open, FALSE, FALSE, 0);
    gtk_box_pack_start(GTK_BOX(button_box), btn_delete, FALSE, FALSE, 0);
    gtk_box_pack_start(GTK_BOX(button_box), btn_export, FALSE, FALSE, 0);
    gtk_box_pack_start(GTK_BOX(button_box), btn_import, FALSE, FALSE, 0);
    
    gtk_widget_show_all(dialog);
    gtk_dialog_run(GTK_DIALOG(dialog));
    gtk_widget_destroy(dialog);
}

*/
import "C"
import (
	"log"
	"unsafe"
)

var (
	globalBookmarkManager *BookmarkManager
	globalCrypto          *CryptoManager
)

// InitBookmarks inicializa o sistema de favoritos
func InitBookmarks() error {
	var err error
	
	// Criar crypto manager
	globalCrypto, err = NewCryptoManager("")
	if err != nil {
		return err
	}
	
	// Criar bookmark manager
	globalBookmarkManager, err = NewBookmarkManager(globalCrypto)
	if err != nil {
		return err
	}
	
	log.Printf("📚 Sistema de favoritos inicializado: %d favoritos", len(globalBookmarkManager.GetAll()))
	return nil
}

//export AddBookmarkFromC
func AddBookmarkFromC(title *C.char, url *C.char) {
	goTitle := C.GoString(title)
	goURL := C.GoString(url)
	
	if globalBookmarkManager == nil {
		log.Println("⚠️  BookmarkManager não inicializado")
		return
	}
	
	err := globalBookmarkManager.Add(goTitle, goURL)
	if err != nil {
		log.Printf("❌ Erro ao adicionar favorito: %v", err)
	} else {
		log.Printf("⭐ Favorito adicionado: %s", goTitle)
	}
}

//export PopulateBookmarksStore
func PopulateBookmarksStore(store *C.GtkListStore) {
	if globalBookmarkManager == nil {
		return
	}
	
	bookmarks := globalBookmarkManager.GetAll()
	
	for _, b := range bookmarks {
		cTitle := C.CString(b.Title)
		cURL := C.CString(b.URL)
		
		// Usar função C helper para adicionar item
		C.add_bookmark_to_store(store, cTitle, cURL)
		
		C.free(unsafe.Pointer(cTitle))
		C.free(unsafe.Pointer(cURL))
	}
}

//export ShowAddBookmarkDialog
func ShowAddBookmarkDialog(window, url, title unsafe.Pointer) {
	C.show_add_bookmark_dialog(
		(*C.GtkWindow)(window),
		(*C.char)(url),
		(*C.char)(title),
	)
}

//export ShowBookmarksDialog
func ShowBookmarksDialog(window unsafe.Pointer) {
	C.show_bookmarks_dialog((*C.GtkWindow)(window))
}
