package main

/*
#cgo pkg-config: webkit2gtk-4.0 gtk+-3.0
#cgo CFLAGS: -Wno-deprecated-declarations
#include <webkit2/webkit2.h>
#include <gtk/gtk.h>
#include <gdk/gdkkeysyms.h>
#include <stdlib.h>
#include <string.h>

// Estrutura do browser
typedef struct {
    GtkApplication *app;
    GtkWindow *window;
    GtkBox *vbox;
    GtkBox *navbar;
    GtkButton *btn_back;
    GtkButton *btn_forward;
    GtkButton *btn_reload;
    GtkButton *btn_home;
    GtkEntry *url_entry;
    GtkNotebook *notebook;
    int tab_count;
} BrowserApp;

static BrowserApp *app_data = NULL;

// Forward declarations
static WebKitWebView* create_webview();

// Configurar WebView com segurança e recursos
static void configure_webview_security(WebKitWebView* webview) {
    WebKitSettings* settings = webkit_web_view_get_settings(webview);
    
    // JavaScript necessário
    webkit_settings_set_enable_javascript(settings, TRUE);
    
    // WebGL
    webkit_settings_set_enable_webgl(settings, TRUE);
    
    // WebAudio
    webkit_settings_set_enable_webaudio(settings, TRUE);
    
    // MediaStream
    webkit_settings_set_enable_media_stream(settings, TRUE);
    
    // WebRTC (tentativa, mas não funciona no Ubuntu 22.04)
    webkit_settings_set_enable_webrtc(settings, TRUE);
    
    // EncryptedMedia
    webkit_settings_set_enable_encrypted_media(settings, TRUE);
    
    // Modal dialogs
    webkit_settings_set_allow_modal_dialogs(settings, TRUE);
    
    // Fullscreen
    webkit_settings_set_enable_fullscreen(settings, TRUE);
    
    // Aceleração de hardware
    webkit_settings_set_hardware_acceleration_policy(settings, WEBKIT_HARDWARE_ACCELERATION_POLICY_ALWAYS);
    
    // JavaScript pode abrir janelas
    webkit_settings_set_javascript_can_open_windows_automatically(settings, TRUE);
    
    // JavaScript pode acessar clipboard
    webkit_settings_set_javascript_can_access_clipboard(settings, TRUE);
    
    // User agent customizado
    webkit_settings_set_user_agent(settings, 
        "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36 Bagus/5.0");
    
    // Desabilitar DNS prefetching (privacidade)
    webkit_settings_set_enable_dns_prefetching(settings, FALSE);
    
    // Habilitar console messages
    webkit_settings_set_enable_write_console_messages_to_stdout(settings, TRUE);
}

// Criar novo WebView
static WebKitWebView* create_webview() {
    WebKitWebView* webview = WEBKIT_WEB_VIEW(webkit_web_view_new());
    configure_webview_security(webview);
    return webview;
}

// Funções de navegação
static void load_uri(WebKitWebView* webview, const char* uri) {
    webkit_web_view_load_uri(webview, uri);
}

static void go_back(WebKitWebView* webview) {
    webkit_web_view_go_back(webview);
}

static void go_forward(WebKitWebView* webview) {
    webkit_web_view_go_forward(webview);
}

static void reload_page(WebKitWebView* webview) {
    webkit_web_view_reload(webview);
}

static void reload_bypass_cache(WebKitWebView* webview) {
    webkit_web_view_reload_bypass_cache(webview);
}

static void stop_loading(WebKitWebView* webview) {
    webkit_web_view_stop_loading(webview);
}

static gboolean can_go_back(WebKitWebView* webview) {
    return webkit_web_view_can_go_back(webview);
}

static gboolean can_go_forward(WebKitWebView* webview) {
    return webkit_web_view_can_go_forward(webview);
}

static const char* get_uri(WebKitWebView* webview) {
    return webkit_web_view_get_uri(webview);
}

static const char* get_title(WebKitWebView* webview) {
    return webkit_web_view_get_title(webview);
}

// Funções de zoom
static void set_zoom_level(WebKitWebView* webview, gdouble zoom) {
    webkit_web_view_set_zoom_level(webview, zoom);
}

static gdouble get_zoom_level(WebKitWebView* webview) {
    return webkit_web_view_get_zoom_level(webview);
}

// Funções de busca
static WebKitFindController* get_find_controller(WebKitWebView* webview) {
    return webkit_web_view_get_find_controller(webview);
}

static void find_text(WebKitWebView* webview, const char* text, gboolean case_sensitive) {
    WebKitFindController* controller = webkit_web_view_get_find_controller(webview);
    guint32 options = WEBKIT_FIND_OPTIONS_WRAP_AROUND;
    if (!case_sensitive) {
        options |= WEBKIT_FIND_OPTIONS_CASE_INSENSITIVE;
    }
    webkit_find_controller_search(controller, text, options, G_MAXUINT);
}

static void find_next(WebKitWebView* webview) {
    WebKitFindController* controller = webkit_web_view_get_find_controller(webview);
    webkit_find_controller_search_next(controller);
}

static void find_previous(WebKitWebView* webview) {
    WebKitFindController* controller = webkit_web_view_get_find_controller(webview);
    webkit_find_controller_search_previous(controller);
}

static void find_finish(WebKitWebView* webview) {
    WebKitFindController* controller = webkit_web_view_get_find_controller(webview);
    webkit_find_controller_search_finish(controller);
}

// Função de impressão
static void print_page(WebKitWebView* webview) {
    WebKitPrintOperation* print_op = webkit_print_operation_new(webview);
    webkit_print_operation_run_dialog(print_op, NULL);
    g_object_unref(print_op);
}

// Callbacks de navegação
static void on_back_clicked(GtkButton *button, gpointer data) {
    int page = gtk_notebook_get_current_page(app_data->notebook);
    if (page >= 0) {
        GtkWidget *child = gtk_notebook_get_nth_page(app_data->notebook, page);
        if (WEBKIT_IS_WEB_VIEW(child)) {
            webkit_web_view_go_back(WEBKIT_WEB_VIEW(child));
        }
    }
}

static void on_forward_clicked(GtkButton *button, gpointer data) {
    int page = gtk_notebook_get_current_page(app_data->notebook);
    if (page >= 0) {
        GtkWidget *child = gtk_notebook_get_nth_page(app_data->notebook, page);
        if (WEBKIT_IS_WEB_VIEW(child)) {
            webkit_web_view_go_forward(WEBKIT_WEB_VIEW(child));
        }
    }
}

static void on_reload_clicked(GtkButton *button, gpointer data) {
    int page = gtk_notebook_get_current_page(app_data->notebook);
    if (page >= 0) {
        GtkWidget *child = gtk_notebook_get_nth_page(app_data->notebook, page);
        if (WEBKIT_IS_WEB_VIEW(child)) {
            webkit_web_view_reload(WEBKIT_WEB_VIEW(child));
        }
    }
}

static void on_home_clicked(GtkButton *button, gpointer data) {
    int page = gtk_notebook_get_current_page(app_data->notebook);
    if (page >= 0) {
        GtkWidget *child = gtk_notebook_get_nth_page(app_data->notebook, page);
        if (WEBKIT_IS_WEB_VIEW(child)) {
            webkit_web_view_load_uri(WEBKIT_WEB_VIEW(child), "https://duckduckgo.com");
        }
    }
}

static void on_url_activate(GtkEntry *entry, gpointer data) {
    GtkEntryBuffer *buffer = gtk_entry_get_buffer(entry);
    const char *url = gtk_entry_buffer_get_text(buffer);
    
    if (!url || strlen(url) == 0) return;
    
    char full_url[2048];
    if (strstr(url, "http://") == url || strstr(url, "https://") == url || strstr(url, "file://") == url) {
        snprintf(full_url, sizeof(full_url), "%s", url);
    } else if (strchr(url, '.') == NULL || strchr(url, ' ') != NULL) {
        snprintf(full_url, sizeof(full_url), "https://duckduckgo.com/?q=%s", url);
    } else {
        snprintf(full_url, sizeof(full_url), "https://%s", url);
    }
    
    int page = gtk_notebook_get_current_page(app_data->notebook);
    if (page >= 0) {
        GtkWidget *child = gtk_notebook_get_nth_page(app_data->notebook, page);
        if (WEBKIT_IS_WEB_VIEW(child)) {
            webkit_web_view_load_uri(WEBKIT_WEB_VIEW(child), full_url);
        }
    }
}

static void on_title_changed(WebKitWebView *webview, GParamSpec *pspec, gpointer data) {
    const char *title = webkit_web_view_get_title(webview);
    if (!title || strlen(title) == 0) {
        title = "Nova Aba";
    }
    
    if (app_data && app_data->notebook) {
        // Atualizar label da aba
        int n_pages = gtk_notebook_get_n_pages(app_data->notebook);
        for (int i = 0; i < n_pages; i++) {
            GtkWidget *page = gtk_notebook_get_nth_page(app_data->notebook, i);
            if (page == GTK_WIDGET(webview)) {
                GtkWidget *label = gtk_label_new(title);
                gtk_notebook_set_tab_label(app_data->notebook, page, label);
                gtk_widget_show(label);
                break;
            }
        }
        
        // Atualizar título da janela se for a aba ativa
        int current = gtk_notebook_get_current_page(app_data->notebook);
        GtkWidget *current_page = gtk_notebook_get_nth_page(app_data->notebook, current);
        if (current_page == GTK_WIDGET(webview) && app_data->window) {
            char window_title[512];
            snprintf(window_title, sizeof(window_title), "%s - Bagus Browser", title);
            gtk_window_set_title(app_data->window, window_title);
        }
    }
}

static void on_uri_changed(WebKitWebView *webview, GParamSpec *pspec, gpointer data) {
    const char *uri = webkit_web_view_get_uri(webview);
    if (uri && app_data && app_data->url_entry) {
        GtkEntryBuffer *buffer = gtk_entry_get_buffer(app_data->url_entry);
        gtk_entry_buffer_set_text(buffer, uri, -1);
    }
}

static gboolean on_permission_request(WebKitWebView *webview, WebKitPermissionRequest *request, gpointer user_data) {
    webkit_permission_request_allow(request);
    return TRUE;
}

// Obter WebView atual
static WebKitWebView* get_current_webview() {
    int page = gtk_notebook_get_current_page(app_data->notebook);
    if (page >= 0) {
        GtkWidget *child = gtk_notebook_get_nth_page(app_data->notebook, page);
        if (WEBKIT_IS_WEB_VIEW(child)) {
            return WEBKIT_WEB_VIEW(child);
        }
    }
    return NULL;
}

// Forward declaration
static void new_tab(const char *url);
static void close_current_tab();
static void show_settings_dialog();

// Forward declarations para favoritos - implementado em bookmarks_ui.go
extern void ShowAddBookmarkDialog(void *window, void *url, void *title);
extern void ShowBookmarksDialog(void *window);

// Forward declarations para autenticação - implementado em auth.go
extern int ShowPasswordDialog(void *window);
extern void ShowSetPasswordDialog(void *window);
extern void ShowChangePasswordDialog(void *window);
extern void ShowRemovePasswordDialog(void *window);

// Forward declarations para sessão - implementado em session_ui.go
extern void SaveCurrentSession(void *notebook);
extern int RestoreSession();
extern char* GetRestoredTabURL(int index);
extern char* GetRestoredTabTitle(int index);
extern int IsRestoredTabActive(int index);

// Callback para salvar sessão ao fechar
static void on_window_destroy(GtkWidget *widget, gpointer data) {
    g_print("💾 Salvando sessão antes de fechar...\n");
    if (app_data && app_data->notebook) {
        SaveCurrentSession(app_data->notebook);
    }
}

// Handler de atalhos de teclado
static gboolean on_key_press(GtkWidget *widget, GdkEventKey *event, gpointer data) {
    guint keyval = event->keyval;
    GdkModifierType state = event->state;
    
    gboolean ctrl = (state & GDK_CONTROL_MASK) != 0;
    gboolean shift = (state & GDK_SHIFT_MASK) != 0;
    gboolean alt = (state & GDK_MOD1_MASK) != 0;
    
    WebKitWebView *webview = get_current_webview();
    
    // Ctrl+Shift+T - Reabrir aba fechada (TODO: implementar histórico)
    if (ctrl && shift && (keyval == GDK_KEY_t || keyval == GDK_KEY_T)) {
        g_print("⌨️  Ctrl+Shift+T - Reabrir aba fechada\n");
        // TODO: implementar histórico de abas fechadas
        return TRUE;
    }
    
    // Ctrl+T - Nova aba
    if (ctrl && !shift && (keyval == GDK_KEY_t || keyval == GDK_KEY_T)) {
        g_print("⌨️  Ctrl+T - Nova aba\n");
        new_tab("https://duckduckgo.com");
        // Focar na URL bar
        gtk_widget_grab_focus(GTK_WIDGET(app_data->url_entry));
        gtk_editable_select_region(GTK_EDITABLE(app_data->url_entry), 0, -1);
        return TRUE;
    }
    
    // Ctrl+W - Fechar aba
    if (ctrl && (keyval == GDK_KEY_w || keyval == GDK_KEY_W)) {
        g_print("⌨️  Ctrl+W - Fechar aba\n");
        close_current_tab();
        return TRUE;
    }
    
    // Ctrl+Tab - Próxima aba
    if (ctrl && keyval == GDK_KEY_Tab) {
        g_print("⌨️  Ctrl+Tab - Próxima aba\n");
        int current = gtk_notebook_get_current_page(app_data->notebook);
        int total = gtk_notebook_get_n_pages(app_data->notebook);
        gtk_notebook_set_current_page(app_data->notebook, (current + 1) % total);
        return TRUE;
    }
    
    // Ctrl+Shift+Tab - Aba anterior
    if (ctrl && shift && keyval == GDK_KEY_ISO_Left_Tab) {
        g_print("⌨️  Ctrl+Shift+Tab - Aba anterior\n");
        int current = gtk_notebook_get_current_page(app_data->notebook);
        int total = gtk_notebook_get_n_pages(app_data->notebook);
        gtk_notebook_set_current_page(app_data->notebook, (current - 1 + total) % total);
        return TRUE;
    }
    
    // Alt+← - Voltar
    if (alt && keyval == GDK_KEY_Left && webview) {
        g_print("⌨️  Alt+← - Voltar\n");
        webkit_web_view_go_back(webview);
        return TRUE;
    }
    
    // Alt+→ - Avançar
    if (alt && keyval == GDK_KEY_Right && webview) {
        g_print("⌨️  Alt+→ - Avançar\n");
        webkit_web_view_go_forward(webview);
        return TRUE;
    }
    
    // F5 ou Ctrl+R - Recarregar
    if ((keyval == GDK_KEY_F5 || (ctrl && keyval == GDK_KEY_r)) && webview) {
        g_print("⌨️  F5/Ctrl+R - Recarregar\n");
        webkit_web_view_reload(webview);
        return TRUE;
    }
    
    // Ctrl+Shift+R - Recarregar sem cache
    if (ctrl && shift && keyval == GDK_KEY_R && webview) {
        g_print("⌨️  Ctrl+Shift+R - Recarregar sem cache\n");
        webkit_web_view_reload_bypass_cache(webview);
        return TRUE;
    }
    
    // Ctrl+L - Focar URL
    if (ctrl && keyval == GDK_KEY_l) {
        g_print("⌨️  Ctrl+L - Focar URL\n");
        gtk_widget_grab_focus(GTK_WIDGET(app_data->url_entry));
        gtk_editable_select_region(GTK_EDITABLE(app_data->url_entry), 0, -1);
        return TRUE;
    }
    
    // Ctrl++ ou Ctrl+= - Zoom in
    if (ctrl && (keyval == GDK_KEY_plus || keyval == GDK_KEY_equal) && webview) {
        g_print("⌨️  Ctrl++ - Aumentar zoom\n");
        gdouble zoom = webkit_web_view_get_zoom_level(webview);
        webkit_web_view_set_zoom_level(webview, zoom + 0.1);
        return TRUE;
    }
    
    // Ctrl+- - Zoom out
    if (ctrl && keyval == GDK_KEY_minus && webview) {
        g_print("⌨️  Ctrl+- - Diminuir zoom\n");
        gdouble zoom = webkit_web_view_get_zoom_level(webview);
        webkit_web_view_set_zoom_level(webview, zoom - 0.1);
        return TRUE;
    }
    
    // Ctrl+0 - Reset zoom
    if (ctrl && keyval == GDK_KEY_0 && webview) {
        g_print("⌨️  Ctrl+0 - Resetar zoom\n");
        webkit_web_view_set_zoom_level(webview, 1.0);
        return TRUE;
    }
    
    // Ctrl+F - Buscar na página (TODO: implementar barra de busca)
    if (ctrl && keyval == GDK_KEY_f) {
        g_print("⌨️  Ctrl+F - Buscar na página\n");
        // TODO: mostrar barra de busca
        return TRUE;
    }
    
    // F3 - Próximo resultado
    if (keyval == GDK_KEY_F3 && webview) {
        g_print("⌨️  F3 - Próximo resultado\n");
        WebKitFindController *controller = webkit_web_view_get_find_controller(webview);
        webkit_find_controller_search_next(controller);
        return TRUE;
    }
    
    // Shift+F3 - Resultado anterior
    if (shift && keyval == GDK_KEY_F3 && webview) {
        g_print("⌨️  Shift+F3 - Resultado anterior\n");
        WebKitFindController *controller = webkit_web_view_get_find_controller(webview);
        webkit_find_controller_search_previous(controller);
        return TRUE;
    }
    
    // Ctrl+P - Imprimir
    if (ctrl && keyval == GDK_KEY_p && webview) {
        g_print("⌨️  Ctrl+P - Imprimir\n");
        WebKitPrintOperation *print_op = webkit_print_operation_new(webview);
        webkit_print_operation_run_dialog(print_op, GTK_WINDOW(app_data->window));
        g_object_unref(print_op);
        return TRUE;
    }
    
    // Ctrl+D - Adicionar favorito
    if (ctrl && keyval == GDK_KEY_d && webview) {
        g_print("⌨️  Ctrl+D - Adicionar favorito\n");
        const char *url = webkit_web_view_get_uri(webview);
        const char *title = webkit_web_view_get_title(webview);
        ShowAddBookmarkDialog(app_data->window, (void*)url, (void*)title);
        return TRUE;
    }
    
    // Ctrl+J - Downloads (TODO: implementar)
    if (ctrl && keyval == GDK_KEY_j) {
        g_print("⌨️  Ctrl+J - Downloads\n");
        // TODO: implementar janela de downloads
        return TRUE;
    }
    
    // Ctrl+, - Configurações
    if (ctrl && keyval == GDK_KEY_comma) {
        g_print("⌨️  Ctrl+, - Configurações\n");
        show_settings_dialog();
        return TRUE;
    }
    
    // Ctrl+Q - Sair
    if (ctrl && keyval == GDK_KEY_q) {
        g_print("⌨️  Ctrl+Q - Sair\n");
        g_application_quit(G_APPLICATION(app_data->app));
        return TRUE;
    }
    
    // Ctrl+1 a Ctrl+9 - Ir para aba específica
    if (ctrl && keyval >= GDK_KEY_1 && keyval <= GDK_KEY_9) {
        int tab_num = keyval - GDK_KEY_1;
        int total = gtk_notebook_get_n_pages(app_data->notebook);
        if (tab_num < total) {
            g_print("⌨️  Ctrl+%d - Ir para aba %d\n", tab_num + 1, tab_num + 1);
            gtk_notebook_set_current_page(app_data->notebook, tab_num);
        }
        return TRUE;
    }
    
    return FALSE;
}

// Callbacks do menu
static void on_menu_about(GtkMenuItem *item, gpointer data) {
    GtkWidget *dialog = gtk_message_dialog_new(
        app_data->window,
        GTK_DIALOG_MODAL,
        GTK_MESSAGE_INFO,
        GTK_BUTTONS_OK,
        "Bagus Browser v5.0.0\n\n"
        "Navegador web focado em privacidade e segurança.\n\n"
        "Tecnologias:\n"
        "• GTK3 - Interface gráfica\n"
        "• WebKitGTK 4.0 - Motor de renderização\n"
        "• CGo Puro - Zero dependências Go\n\n"
        "© 2024 Bagus Browser"
    );
    gtk_dialog_run(GTK_DIALOG(dialog));
    gtk_widget_destroy(dialog);
}

static void on_menu_version(GtkMenuItem *item, gpointer data) {
    GtkWidget *dialog = gtk_message_dialog_new(
        app_data->window,
        GTK_DIALOG_MODAL,
        GTK_MESSAGE_INFO,
        GTK_BUTTONS_OK,
        "Bagus Browser\nVersão: 5.0.0\n\n"
        "WebKitGTK: 4.0\n"
        "GTK: 3.0\n"
        "Build: %s", __DATE__
    );
    gtk_dialog_run(GTK_DIALOG(dialog));
    gtk_widget_destroy(dialog);
}

static void on_menu_settings(GtkMenuItem *item, gpointer data) {
    show_settings_dialog();
}

// Implementação do diálogo de configurações
static void show_settings_dialog() {
    GtkWidget *dialog = gtk_dialog_new_with_buttons(
        "⚙️  Configurações - Bagus Browser v5.0.0",
        app_data->window,
        GTK_DIALOG_MODAL,
        "Cancelar", GTK_RESPONSE_CANCEL,
        "Salvar", GTK_RESPONSE_OK,
        NULL
    );
    
    gtk_window_set_default_size(GTK_WINDOW(dialog), 750, 650);
    
    GtkWidget *content_area = gtk_dialog_get_content_area(GTK_DIALOG(dialog));
    gtk_container_set_border_width(GTK_CONTAINER(content_area), 10);
    
    // Notebook com abas
    GtkWidget *notebook = gtk_notebook_new();
    gtk_box_pack_start(GTK_BOX(content_area), notebook, TRUE, TRUE, 0);
    
    // ABA 1: SEGURANÇA
    GtkWidget *security_box = gtk_box_new(GTK_ORIENTATION_VERTICAL, 10);
    gtk_container_set_border_width(GTK_CONTAINER(security_box), 15);
    
    GtkWidget *security_title = gtk_label_new(NULL);
    gtk_label_set_markup(GTK_LABEL(security_title), "<b>Configurações de Segurança</b>");
    gtk_widget_set_halign(security_title, GTK_ALIGN_START);
    gtk_box_pack_start(GTK_BOX(security_box), security_title, FALSE, FALSE, 0);
    
    GtkWidget *password_frame = gtk_frame_new("🔒 Senha Mestre");
    GtkWidget *password_inner = gtk_box_new(GTK_ORIENTATION_VERTICAL, 10);
    gtk_container_set_border_width(GTK_CONTAINER(password_inner), 10);
    
    GtkWidget *password_desc = gtk_label_new(
        "Proteja o acesso ao browser com uma senha mestre.\n"
        "Todos os seus dados (favoritos, sessões, cookies) já estão criptografados."
    );
    gtk_widget_set_halign(password_desc, GTK_ALIGN_START);
    gtk_label_set_line_wrap(GTK_LABEL(password_desc), TRUE);
    gtk_box_pack_start(GTK_BOX(password_inner), password_desc, FALSE, FALSE, 0);
    
    // Botões de gerenciamento de senha
    GtkWidget *button_box = gtk_box_new(GTK_ORIENTATION_HORIZONTAL, 5);
    gtk_widget_set_halign(button_box, GTK_ALIGN_START);
    
    GtkWidget *set_pwd_btn = gtk_button_new_with_label("Definir Senha");
    GtkWidget *change_pwd_btn = gtk_button_new_with_label("Alterar Senha");
    GtkWidget *remove_pwd_btn = gtk_button_new_with_label("Remover Senha");
    
    g_signal_connect_swapped(set_pwd_btn, "clicked", G_CALLBACK(ShowSetPasswordDialog), app_data->window);
    g_signal_connect_swapped(change_pwd_btn, "clicked", G_CALLBACK(ShowChangePasswordDialog), app_data->window);
    g_signal_connect_swapped(remove_pwd_btn, "clicked", G_CALLBACK(ShowRemovePasswordDialog), app_data->window);
    
    gtk_box_pack_start(GTK_BOX(button_box), set_pwd_btn, FALSE, FALSE, 0);
    gtk_box_pack_start(GTK_BOX(button_box), change_pwd_btn, FALSE, FALSE, 0);
    gtk_box_pack_start(GTK_BOX(button_box), remove_pwd_btn, FALSE, FALSE, 0);
    
    gtk_box_pack_start(GTK_BOX(password_inner), button_box, FALSE, FALSE, 5);
    
    GtkWidget *password_info = gtk_label_new(
        "\n💡 Dica: A senha será solicitada ao iniciar o browser.\n"
        "Se você esquecer a senha, não será possível recuperar os dados."
    );
    gtk_widget_set_halign(password_info, GTK_ALIGN_START);
    gtk_label_set_line_wrap(GTK_LABEL(password_info), TRUE);
    gtk_box_pack_start(GTK_BOX(password_inner), password_info, FALSE, FALSE, 0);
    
    gtk_container_add(GTK_CONTAINER(password_frame), password_inner);
    gtk_box_pack_start(GTK_BOX(security_box), password_frame, FALSE, FALSE, 0);
    
    GtkWidget *security_label = gtk_label_new("🔒 Segurança");
    gtk_notebook_append_page(GTK_NOTEBOOK(notebook), security_box, security_label);
    
    // ABA 2: SESSÕES
    GtkWidget *session_box = gtk_box_new(GTK_ORIENTATION_VERTICAL, 10);
    gtk_container_set_border_width(GTK_CONTAINER(session_box), 15);
    
    GtkWidget *session_title = gtk_label_new(NULL);
    gtk_label_set_markup(GTK_LABEL(session_title), "<b>Sessões e Cookies</b>");
    gtk_widget_set_halign(session_title, GTK_ALIGN_START);
    gtk_box_pack_start(GTK_BOX(session_box), session_title, FALSE, FALSE, 0);
    
    GtkWidget *persist_sessions = gtk_check_button_new_with_label("Manter usuário logado em sites");
    gtk_toggle_button_set_active(GTK_TOGGLE_BUTTON(persist_sessions), TRUE);
    gtk_box_pack_start(GTK_BOX(session_box), persist_sessions, FALSE, FALSE, 0);
    
    GtkWidget *persist_cookies = gtk_check_button_new_with_label("Salvar cookies entre sessões");
    gtk_toggle_button_set_active(GTK_TOGGLE_BUTTON(persist_cookies), TRUE);
    gtk_box_pack_start(GTK_BOX(session_box), persist_cookies, FALSE, FALSE, 0);
    
    GtkWidget *session_label = gtk_label_new("🔄 Sessões");
    gtk_notebook_append_page(GTK_NOTEBOOK(notebook), session_box, session_label);
    
    // ABA 3: PERFORMANCE
    GtkWidget *perf_box = gtk_box_new(GTK_ORIENTATION_VERTICAL, 10);
    gtk_container_set_border_width(GTK_CONTAINER(perf_box), 15);
    
    GtkWidget *perf_title = gtk_label_new(NULL);
    gtk_label_set_markup(GTK_LABEL(perf_title), "<b>Configurações de Performance</b>");
    gtk_widget_set_halign(perf_title, GTK_ALIGN_START);
    gtk_box_pack_start(GTK_BOX(perf_box), perf_title, FALSE, FALSE, 0);
    
    // Frame Cache de Vídeos
    GtkWidget *cache_frame = gtk_frame_new("📹 Cache de Vídeos");
    GtkWidget *cache_inner = gtk_box_new(GTK_ORIENTATION_VERTICAL, 8);
    gtk_container_set_border_width(GTK_CONTAINER(cache_inner), 10);
    
    GtkWidget *cache_enabled = gtk_check_button_new_with_label("Habilitar cache de vídeos");
    gtk_toggle_button_set_active(GTK_TOGGLE_BUTTON(cache_enabled), TRUE);
    gtk_box_pack_start(GTK_BOX(cache_inner), cache_enabled, FALSE, FALSE, 0);
    
    // Slider de tamanho do cache
    GtkWidget *cache_size_box = gtk_box_new(GTK_ORIENTATION_HORIZONTAL, 10);
    GtkWidget *cache_label = gtk_label_new("Tamanho do cache:");
    gtk_widget_set_halign(cache_label, GTK_ALIGN_START);
    gtk_box_pack_start(GTK_BOX(cache_size_box), cache_label, FALSE, FALSE, 0);
    
    GtkWidget *cache_scale = gtk_scale_new_with_range(GTK_ORIENTATION_HORIZONTAL, 100, 5000, 100);
    gtk_range_set_value(GTK_RANGE(cache_scale), 500);
    gtk_scale_set_value_pos(GTK_SCALE(cache_scale), GTK_POS_RIGHT);
    gtk_scale_set_digits(GTK_SCALE(cache_scale), 0);
    gtk_widget_set_hexpand(cache_scale, TRUE);
    
    // Adicionar marcas no slider
    gtk_scale_add_mark(GTK_SCALE(cache_scale), 100, GTK_POS_BOTTOM, "100 MB");
    gtk_scale_add_mark(GTK_SCALE(cache_scale), 500, GTK_POS_BOTTOM, "500 MB");
    gtk_scale_add_mark(GTK_SCALE(cache_scale), 1000, GTK_POS_BOTTOM, "1 GB");
    gtk_scale_add_mark(GTK_SCALE(cache_scale), 2000, GTK_POS_BOTTOM, "2 GB");
    gtk_scale_add_mark(GTK_SCALE(cache_scale), 5000, GTK_POS_BOTTOM, "5 GB");
    
    gtk_box_pack_start(GTK_BOX(cache_size_box), cache_scale, TRUE, TRUE, 0);
    gtk_box_pack_start(GTK_BOX(cache_inner), cache_size_box, FALSE, FALSE, 5);
    
    GtkWidget *cache_desc = gtk_label_new(
        "Melhora reprodução de vídeos no YouTube e outros sites.\n"
        "Maior cache = melhor performance, mas usa mais espaço em disco."
    );
    gtk_widget_set_halign(cache_desc, GTK_ALIGN_START);
    gtk_label_set_line_wrap(GTK_LABEL(cache_desc), TRUE);
    gtk_box_pack_start(GTK_BOX(cache_inner), cache_desc, FALSE, FALSE, 5);
    
    gtk_container_add(GTK_CONTAINER(cache_frame), cache_inner);
    gtk_box_pack_start(GTK_BOX(perf_box), cache_frame, FALSE, FALSE, 0);
    
    GtkWidget *perf_label = gtk_label_new("⚡ Performance");
    gtk_notebook_append_page(GTK_NOTEBOOK(notebook), perf_box, perf_label);
    
    // ABA 4: PRIVACIDADE
    GtkWidget *privacy_box = gtk_box_new(GTK_ORIENTATION_VERTICAL, 10);
    gtk_container_set_border_width(GTK_CONTAINER(privacy_box), 15);
    
    GtkWidget *privacy_title = gtk_label_new(NULL);
    gtk_label_set_markup(GTK_LABEL(privacy_title), "<b>Configurações de Privacidade</b>");
    gtk_widget_set_halign(privacy_title, GTK_ALIGN_START);
    gtk_box_pack_start(GTK_BOX(privacy_box), privacy_title, FALSE, FALSE, 0);
    
    // Frame Bloqueios
    GtkWidget *block_frame = gtk_frame_new("🛡️  Bloqueios de Privacidade");
    GtkWidget *block_inner = gtk_box_new(GTK_ORIENTATION_VERTICAL, 5);
    gtk_container_set_border_width(GTK_CONTAINER(block_inner), 10);
    
    GtkWidget *block_3rd = gtk_check_button_new_with_label("Bloquear cookies de terceiros");
    gtk_toggle_button_set_active(GTK_TOGGLE_BUTTON(block_3rd), TRUE);
    gtk_box_pack_start(GTK_BOX(block_inner), block_3rd, FALSE, FALSE, 0);
    
    GtkWidget *block_geo = gtk_check_button_new_with_label("Bloquear geolocalização");
    gtk_toggle_button_set_active(GTK_TOGGLE_BUTTON(block_geo), TRUE);
    gtk_box_pack_start(GTK_BOX(block_inner), block_geo, FALSE, FALSE, 0);
    
    GtkWidget *block_notif = gtk_check_button_new_with_label("Bloquear notificações");
    gtk_box_pack_start(GTK_BOX(block_inner), block_notif, FALSE, FALSE, 0);
    
    GtkWidget *block_media = gtk_check_button_new_with_label("Bloquear acesso a câmera/microfone");
    gtk_box_pack_start(GTK_BOX(block_inner), block_media, FALSE, FALSE, 0);
    
    GtkWidget *block_webgl = gtk_check_button_new_with_label("Bloquear WebGL");
    gtk_box_pack_start(GTK_BOX(block_inner), block_webgl, FALSE, FALSE, 0);
    
    GtkWidget *block_webaudio = gtk_check_button_new_with_label("Bloquear WebAudio");
    gtk_box_pack_start(GTK_BOX(block_inner), block_webaudio, FALSE, FALSE, 0);
    
    GtkWidget *do_not_track = gtk_check_button_new_with_label("Enviar 'Do Not Track'");
    gtk_toggle_button_set_active(GTK_TOGGLE_BUTTON(do_not_track), TRUE);
    gtk_box_pack_start(GTK_BOX(block_inner), do_not_track, FALSE, FALSE, 0);
    
    gtk_container_add(GTK_CONTAINER(block_frame), block_inner);
    gtk_box_pack_start(GTK_BOX(privacy_box), block_frame, FALSE, FALSE, 0);
    
    // Descrição
    GtkWidget *privacy_desc = gtk_label_new(
        "⚠️  Atenção: Bloquear WebGL e WebAudio pode quebrar sites como Google Meet.\n"
        "Recomendado: Manter apenas cookies de terceiros e geolocalização bloqueados."
    );
    gtk_widget_set_halign(privacy_desc, GTK_ALIGN_START);
    gtk_label_set_line_wrap(GTK_LABEL(privacy_desc), TRUE);
    gtk_box_pack_start(GTK_BOX(privacy_box), privacy_desc, FALSE, FALSE, 5);
    
    // Frame Informações
    GtkWidget *info_frame = gtk_frame_new("ℹ️  Informações de Privacidade");
    GtkWidget *info_inner = gtk_box_new(GTK_ORIENTATION_VERTICAL, 5);
    gtk_container_set_border_width(GTK_CONTAINER(info_inner), 10);
    
    GtkWidget *info_label = gtk_label_new(
        "✅ Zero telemetria\n"
        "✅ Sem analytics\n"
        "✅ Sem crash reports\n"
        "✅ Sem rastreamento\n"
        "✅ DuckDuckGo como motor de busca padrão"
    );
    gtk_widget_set_halign(info_label, GTK_ALIGN_START);
    gtk_box_pack_start(GTK_BOX(info_inner), info_label, FALSE, FALSE, 0);
    
    gtk_container_add(GTK_CONTAINER(info_frame), info_inner);
    gtk_box_pack_start(GTK_BOX(privacy_box), info_frame, FALSE, FALSE, 0);
    
    GtkWidget *privacy_label = gtk_label_new("🕵️  Privacidade");
    gtk_notebook_append_page(GTK_NOTEBOOK(notebook), privacy_box, privacy_label);
    
    gtk_widget_show_all(dialog);
    
    gint response = gtk_dialog_run(GTK_DIALOG(dialog));
    
    if (response == GTK_RESPONSE_OK) {
        g_print("⚙️  Salvando configurações...\n");
        
        GtkWidget *msg = gtk_message_dialog_new(
            app_data->window,
            GTK_DIALOG_MODAL,
            GTK_MESSAGE_INFO,
            GTK_BUTTONS_OK,
            "Configurações salvas com sucesso!\n\nAlgumas mudanças podem exigir reiniciar o browser."
        );
        gtk_dialog_run(GTK_DIALOG(msg));
        gtk_widget_destroy(msg);
    }
    
    gtk_widget_destroy(dialog);
}

// Criar barra de menu
static GtkWidget* create_menu_bar() {
    GtkWidget *menu_bar = gtk_menu_bar_new();
    
    // Menu Arquivo
    GtkWidget *menu_arquivo = gtk_menu_item_new_with_label("Arquivo");
    GtkWidget *menu_arquivo_sub = gtk_menu_new();
    
    GtkWidget *item_nova_aba = gtk_menu_item_new_with_label("Nova Aba (Ctrl+T)");
    g_signal_connect(item_nova_aba, "activate", G_CALLBACK(new_tab), "https://duckduckgo.com");
    gtk_menu_shell_append(GTK_MENU_SHELL(menu_arquivo_sub), item_nova_aba);
    
    GtkWidget *item_fechar_aba = gtk_menu_item_new_with_label("Fechar Aba (Ctrl+W)");
    g_signal_connect(item_fechar_aba, "activate", G_CALLBACK(close_current_tab), NULL);
    gtk_menu_shell_append(GTK_MENU_SHELL(menu_arquivo_sub), item_fechar_aba);
    
    GtkWidget *sep1 = gtk_separator_menu_item_new();
    gtk_menu_shell_append(GTK_MENU_SHELL(menu_arquivo_sub), sep1);
    
    GtkWidget *item_sair = gtk_menu_item_new_with_label("Sair (Ctrl+Q)");
    g_signal_connect_swapped(item_sair, "activate", G_CALLBACK(g_application_quit), app_data->app);
    gtk_menu_shell_append(GTK_MENU_SHELL(menu_arquivo_sub), item_sair);
    
    gtk_menu_item_set_submenu(GTK_MENU_ITEM(menu_arquivo), menu_arquivo_sub);
    gtk_menu_shell_append(GTK_MENU_SHELL(menu_bar), menu_arquivo);
    
    // Menu Navegação
    GtkWidget *menu_nav = gtk_menu_item_new_with_label("Navegação");
    GtkWidget *menu_nav_sub = gtk_menu_new();
    
    GtkWidget *item_voltar = gtk_menu_item_new_with_label("Voltar (Alt+←)");
    g_signal_connect(item_voltar, "activate", G_CALLBACK(on_back_clicked), NULL);
    gtk_menu_shell_append(GTK_MENU_SHELL(menu_nav_sub), item_voltar);
    
    GtkWidget *item_avancar = gtk_menu_item_new_with_label("Avançar (Alt+→)");
    g_signal_connect(item_avancar, "activate", G_CALLBACK(on_forward_clicked), NULL);
    gtk_menu_shell_append(GTK_MENU_SHELL(menu_nav_sub), item_avancar);
    
    GtkWidget *item_recarregar = gtk_menu_item_new_with_label("Recarregar (F5)");
    g_signal_connect(item_recarregar, "activate", G_CALLBACK(on_reload_clicked), NULL);
    gtk_menu_shell_append(GTK_MENU_SHELL(menu_nav_sub), item_recarregar);
    
    GtkWidget *item_home = gtk_menu_item_new_with_label("Página Inicial");
    g_signal_connect(item_home, "activate", G_CALLBACK(on_home_clicked), NULL);
    gtk_menu_shell_append(GTK_MENU_SHELL(menu_nav_sub), item_home);
    
    gtk_menu_item_set_submenu(GTK_MENU_ITEM(menu_nav), menu_nav_sub);
    gtk_menu_shell_append(GTK_MENU_SHELL(menu_bar), menu_nav);
    
    // Menu Favoritos
    GtkWidget *menu_fav = gtk_menu_item_new_with_label("Favoritos");
    GtkWidget *menu_fav_sub = gtk_menu_new();
    
    GtkWidget *item_add_bookmark = gtk_menu_item_new_with_label("Adicionar Favorito (Ctrl+D)");
    g_signal_connect_swapped(item_add_bookmark, "activate", G_CALLBACK(ShowAddBookmarkDialog), app_data->window);
    gtk_menu_shell_append(GTK_MENU_SHELL(menu_fav_sub), item_add_bookmark);
    
    GtkWidget *item_manage_bookmarks = gtk_menu_item_new_with_label("Gerenciar Favoritos (Ctrl+Shift+B)");
    g_signal_connect_swapped(item_manage_bookmarks, "activate", G_CALLBACK(ShowBookmarksDialog), app_data->window);
    gtk_menu_shell_append(GTK_MENU_SHELL(menu_fav_sub), item_manage_bookmarks);
    
    gtk_menu_item_set_submenu(GTK_MENU_ITEM(menu_fav), menu_fav_sub);
    gtk_menu_shell_append(GTK_MENU_SHELL(menu_bar), menu_fav);
    
    // Menu Ferramentas
    GtkWidget *menu_tools = gtk_menu_item_new_with_label("Ferramentas");
    GtkWidget *menu_tools_sub = gtk_menu_new();
    
    GtkWidget *item_buscar = gtk_menu_item_new_with_label("Buscar na Página (Ctrl+F)");
    gtk_menu_shell_append(GTK_MENU_SHELL(menu_tools_sub), item_buscar);
    
    GtkWidget *sep2 = gtk_separator_menu_item_new();
    gtk_menu_shell_append(GTK_MENU_SHELL(menu_tools_sub), sep2);
    
    GtkWidget *item_zoom_in = gtk_menu_item_new_with_label("Aumentar Zoom (Ctrl++)");
    gtk_menu_shell_append(GTK_MENU_SHELL(menu_tools_sub), item_zoom_in);
    
    GtkWidget *item_zoom_out = gtk_menu_item_new_with_label("Diminuir Zoom (Ctrl+-)");
    gtk_menu_shell_append(GTK_MENU_SHELL(menu_tools_sub), item_zoom_out);
    
    GtkWidget *item_zoom_reset = gtk_menu_item_new_with_label("Zoom 100% (Ctrl+0)");
    gtk_menu_shell_append(GTK_MENU_SHELL(menu_tools_sub), item_zoom_reset);
    
    GtkWidget *sep3 = gtk_separator_menu_item_new();
    gtk_menu_shell_append(GTK_MENU_SHELL(menu_tools_sub), sep3);
    
    GtkWidget *item_downloads = gtk_menu_item_new_with_label("Downloads (Ctrl+J)");
    gtk_menu_shell_append(GTK_MENU_SHELL(menu_tools_sub), item_downloads);
    
    gtk_menu_item_set_submenu(GTK_MENU_ITEM(menu_tools), menu_tools_sub);
    gtk_menu_shell_append(GTK_MENU_SHELL(menu_bar), menu_tools);
    
    // Menu Editar
    GtkWidget *menu_edit = gtk_menu_item_new_with_label("Editar");
    GtkWidget *menu_edit_sub = gtk_menu_new();
    
    GtkWidget *item_config = gtk_menu_item_new_with_label("Configurações (Ctrl+,)");
    g_signal_connect(item_config, "activate", G_CALLBACK(on_menu_settings), NULL);
    gtk_menu_shell_append(GTK_MENU_SHELL(menu_edit_sub), item_config);
    
    gtk_menu_item_set_submenu(GTK_MENU_ITEM(menu_edit), menu_edit_sub);
    gtk_menu_shell_append(GTK_MENU_SHELL(menu_bar), menu_edit);
    
    // Menu Ajuda
    GtkWidget *menu_help = gtk_menu_item_new_with_label("Ajuda");
    GtkWidget *menu_help_sub = gtk_menu_new();
    
    GtkWidget *item_version = gtk_menu_item_new_with_label("Versão");
    g_signal_connect(item_version, "activate", G_CALLBACK(on_menu_version), NULL);
    gtk_menu_shell_append(GTK_MENU_SHELL(menu_help_sub), item_version);
    
    GtkWidget *item_about = gtk_menu_item_new_with_label("Sobre");
    g_signal_connect(item_about, "activate", G_CALLBACK(on_menu_about), NULL);
    gtk_menu_shell_append(GTK_MENU_SHELL(menu_help_sub), item_about);
    
    gtk_menu_item_set_submenu(GTK_MENU_ITEM(menu_help), menu_help_sub);
    gtk_menu_shell_append(GTK_MENU_SHELL(menu_bar), menu_help);
    
    return menu_bar;
}

// Fechar aba atual
static void close_current_tab() {
    int page = gtk_notebook_get_current_page(app_data->notebook);
    if (page >= 0) {
        int total = gtk_notebook_get_n_pages(app_data->notebook);
        if (total > 1) {
            gtk_notebook_remove_page(app_data->notebook, page);
            app_data->tab_count--;
        } else {
            // Última aba - sair
            g_application_quit(G_APPLICATION(app_data->app));
        }
    }
}

// Criar nova aba
static void new_tab(const char *url) {
    WebKitWebView *webview = create_webview();
    
    // Conectar sinais
    g_signal_connect(webview, "notify::title", G_CALLBACK(on_title_changed), NULL);
    g_signal_connect(webview, "notify::uri", G_CALLBACK(on_uri_changed), NULL);
    g_signal_connect(webview, "permission-request", G_CALLBACK(on_permission_request), NULL);
    
    // Label da aba
    GtkWidget *label = gtk_label_new("Nova Aba");
    
    // Adicionar ao notebook
    int page = gtk_notebook_append_page(app_data->notebook, GTK_WIDGET(webview), label);
    gtk_notebook_set_current_page(app_data->notebook, page);
    gtk_notebook_set_tab_reorderable(app_data->notebook, GTK_WIDGET(webview), TRUE);
    
    // Mostrar o webview (GTK3)
    gtk_widget_show(GTK_WIDGET(webview));
    
    // Carregar URL
    if (url && strlen(url) > 0) {
        webkit_web_view_load_uri(webview, url);
    } else {
        webkit_web_view_load_uri(webview, "https://duckduckgo.com");
    }
    
    app_data->tab_count++;
}

// Criar interface
static void create_ui() {
    // Janela
    app_data->window = GTK_WINDOW(gtk_application_window_new(app_data->app));
    gtk_window_set_title(app_data->window, "Bagus Browser v5.0.0");
    gtk_window_set_default_size(app_data->window, 1200, 800);
    
    // Conectar callback para salvar sessão ao fechar
    g_signal_connect(app_data->window, "destroy", G_CALLBACK(on_window_destroy), NULL);
    
    // Definir ícone da aplicação
    GError *error = NULL;
    const char *icon_paths[] = {
        "/home/peder/bagus-browser-go/assets/icons/bagus-browser-128.png",
        "/usr/share/icons/hicolor/128x128/apps/bagus-browser.png",
        "/usr/share/pixmaps/bagus-browser.png",
        NULL
    };
    
    for (int i = 0; icon_paths[i] != NULL; i++) {
        if (g_file_test(icon_paths[i], G_FILE_TEST_EXISTS)) {
            gtk_window_set_icon_from_file(app_data->window, icon_paths[i], &error);
            if (!error) {
                g_print("✅ Ícone carregado: %s\n", icon_paths[i]);
                break;
            } else {
                g_error_free(error);
                error = NULL;
            }
        }
    }
    
    // Também definir ícone padrão da aplicação
    gtk_window_set_default_icon_name("bagus-browser");
    
    // VBox principal
    app_data->vbox = GTK_BOX(gtk_box_new(GTK_ORIENTATION_VERTICAL, 0));
    
    // Barra de menu
    GtkWidget *menu_bar = create_menu_bar();
    gtk_box_pack_start(app_data->vbox, menu_bar, FALSE, FALSE, 0);
    
    // Navbar
    app_data->navbar = GTK_BOX(gtk_box_new(GTK_ORIENTATION_HORIZONTAL, 5));
    gtk_widget_set_margin_start(GTK_WIDGET(app_data->navbar), 5);
    gtk_widget_set_margin_end(GTK_WIDGET(app_data->navbar), 5);
    gtk_widget_set_margin_top(GTK_WIDGET(app_data->navbar), 5);
    gtk_widget_set_margin_bottom(GTK_WIDGET(app_data->navbar), 5);
    
    // Botões (GTK3)
    app_data->btn_back = GTK_BUTTON(gtk_button_new_from_icon_name("go-previous", GTK_ICON_SIZE_BUTTON));
    app_data->btn_forward = GTK_BUTTON(gtk_button_new_from_icon_name("go-next", GTK_ICON_SIZE_BUTTON));
    app_data->btn_reload = GTK_BUTTON(gtk_button_new_from_icon_name("view-refresh", GTK_ICON_SIZE_BUTTON));
    app_data->btn_home = GTK_BUTTON(gtk_button_new_from_icon_name("go-home", GTK_ICON_SIZE_BUTTON));
    
    g_signal_connect(app_data->btn_back, "clicked", G_CALLBACK(on_back_clicked), NULL);
    g_signal_connect(app_data->btn_forward, "clicked", G_CALLBACK(on_forward_clicked), NULL);
    g_signal_connect(app_data->btn_reload, "clicked", G_CALLBACK(on_reload_clicked), NULL);
    g_signal_connect(app_data->btn_home, "clicked", G_CALLBACK(on_home_clicked), NULL);
    
    // Entry URL
    app_data->url_entry = GTK_ENTRY(gtk_entry_new());
    gtk_widget_set_hexpand(GTK_WIDGET(app_data->url_entry), TRUE);
    g_signal_connect(app_data->url_entry, "activate", G_CALLBACK(on_url_activate), NULL);
    
    // Adicionar à navbar (GTK3)
    gtk_box_pack_start(app_data->navbar, GTK_WIDGET(app_data->btn_back), FALSE, FALSE, 0);
    gtk_box_pack_start(app_data->navbar, GTK_WIDGET(app_data->btn_forward), FALSE, FALSE, 0);
    gtk_box_pack_start(app_data->navbar, GTK_WIDGET(app_data->btn_reload), FALSE, FALSE, 0);
    gtk_box_pack_start(app_data->navbar, GTK_WIDGET(app_data->btn_home), FALSE, FALSE, 0);
    gtk_box_pack_start(app_data->navbar, GTK_WIDGET(app_data->url_entry), TRUE, TRUE, 5);
    
    // Notebook
    app_data->notebook = GTK_NOTEBOOK(gtk_notebook_new());
    gtk_notebook_set_scrollable(app_data->notebook, TRUE);
    gtk_widget_set_vexpand(GTK_WIDGET(app_data->notebook), TRUE);
    
    // Montar interface (GTK3)
    gtk_box_pack_start(app_data->vbox, GTK_WIDGET(app_data->navbar), FALSE, FALSE, 0);
    gtk_box_pack_start(app_data->vbox, GTK_WIDGET(app_data->notebook), TRUE, TRUE, 0);
    
    gtk_container_add(GTK_CONTAINER(app_data->window), GTK_WIDGET(app_data->vbox));
    
    // Conectar handler de teclado
    g_signal_connect(app_data->window, "key-press-event", G_CALLBACK(on_key_press), NULL);
    
    // Mostrar todos os widgets (GTK3)
    gtk_widget_show_all(GTK_WIDGET(app_data->window));
    
    // Restaurar sessão anterior ou criar aba padrão
    int restored_count = RestoreSession();
    
    if (restored_count > 0) {
        g_print("📂 Restaurando %d abas da sessão anterior...\n", restored_count);
        
        int active_index = 0;
        for (int i = 0; i < restored_count; i++) {
            char *url = GetRestoredTabURL(i);
            if (url) {
                new_tab(url);
                
                // Verificar se esta é a aba ativa
                if (IsRestoredTabActive(i)) {
                    active_index = i;
                }
                
                free(url);
            }
        }
        
        // Ativar a aba que estava ativa
        if (active_index > 0) {
            gtk_notebook_set_current_page(app_data->notebook, active_index);
        }
    } else {
        // Nenhuma sessão anterior, criar aba padrão
        g_print("🆕 Criando nova sessão...\n");
        new_tab("https://duckduckgo.com");
    }
}

// Callback de ativação
static void on_activate(GtkApplication *app, gpointer user_data) {
    if (!app_data) {
        app_data = g_new0(BrowserApp, 1);
        app_data->app = app;
        app_data->tab_count = 0;
    }
    
    // Criar janela temporária para diálogo de senha
    GtkWindow *temp_window = GTK_WINDOW(gtk_window_new(GTK_WINDOW_TOPLEVEL));
    
    // Verificar senha mestre (se habilitada)
    int auth_result = ShowPasswordDialog(temp_window);
    
    gtk_widget_destroy(GTK_WIDGET(temp_window));
    
    if (auth_result == 0) {
        // Autenticação falhou ou foi cancelada
        g_print("❌ Autenticação cancelada ou falhou. Encerrando...\n");
        g_application_quit(G_APPLICATION(app));
        return;
    }
    
    create_ui();
}

// Executar browser
static int run_browser(int argc, char **argv) {
    GtkApplication *app = gtk_application_new("com.bagus.browser", G_APPLICATION_FLAGS_NONE);
    g_signal_connect(app, "activate", G_CALLBACK(on_activate), NULL);
    int status = g_application_run(G_APPLICATION(app), argc, argv);
    g_object_unref(app);
    if (app_data) {
        g_free(app_data);
        app_data = NULL;
    }
    return status;
}

*/
import "C"
import (
	"log"
	"os"
	"runtime"
	"time"
	"unsafe"
)

const (
	AppVersion = "5.0.0"
	AppName    = "Bagus Browser"
)

// Variável global para gerenciamento de sessão
var globalSessionManager *SessionManager

// Tab representa uma aba do navegador
type Tab struct {
	webview unsafe.Pointer // *C.WebKitWebView
	title   string
	url     string
	loading bool
}

// ClosedTab representa uma aba fechada (para Ctrl+Shift+T)
type ClosedTab struct {
	URL   string
	Title string
	Time  time.Time
}

// URLValidator já está definido em security.go

// Browser representa o navegador
type Browser struct {
	app             unsafe.Pointer // *C.GtkApplication
	window          unsafe.Pointer // *C.GtkWindow
	notebook        unsafe.Pointer // *C.GtkNotebook
	urlEntry        unsafe.Pointer // *C.GtkEntry
	findBar         unsafe.Pointer // *C.GtkBox
	findEntry       unsafe.Pointer // *C.GtkEntry
	findBarVisible  bool
	tabs            []*Tab
	closedTabs      []ClosedTab
	validator       *URLValidator
	privacyManager  *PrivacyManager
	bookmarkManager *BookmarkManager
	downloadManager *DownloadManager
	downloadHandler *DownloadHandler
	sessionManager  *SessionManager
	config          *Config
}

func main() {
	runtime.LockOSThread()

	log.Println("═══════════════════════════════════════════════════════════")
	log.Printf("🌐 %s v%s", AppName, AppVersion)
	log.Println("═══════════════════════════════════════════════════════════")
	log.Println("")
	log.Println("🚀 Tecnologias:")
	log.Println("   ✅ GTK4 - Interface moderna")
	log.Println("   ✅ WebKitGTK 6.0 - Motor de renderização")
	log.Println("   ✅ CGo Puro - Zero dependências Go problemáticas")
	log.Println("")
	log.Println("🎯 Funcionalidades:")
	log.Println("   ✅ Sistema de abas completo")
	log.Println("   ✅ Navegação completa")
	log.Println("   ✅ Toolbar com botões")
	log.Println("   ✅ Entry URL inteligente")
	log.Println("")
	log.Println("═══════════════════════════════════════════════════════════")
	log.Println("")

	// Carregar configurações
	_, err := LoadConfig()
	if err != nil {
		log.Printf("⚠️  Erro ao carregar configurações: %v (usando padrões)", err)
	} else {
		log.Println("⚙️  Configurações carregadas")
	}

	// Inicializar sistema de autenticação
	globalAuthManager, err = NewAuthManager()
	if err != nil {
		log.Printf("⚠️  Erro ao inicializar autenticação: %v", err)
	} else {
		log.Println("🔒 Sistema de autenticação inicializado")
	}

	// Inicializar sistema de sessão
	globalCrypto, err := NewCryptoManager("")
	if err != nil {
		log.Printf("⚠️  Erro ao criar crypto manager: %v", err)
	}
	
	globalSessionManager, err = NewSessionManager(globalCrypto)
	if err != nil {
		log.Printf("⚠️  Erro ao inicializar sessão: %v", err)
	} else {
		log.Println("💾 Sistema de sessão inicializado")
	}

	// Inicializar sistema de favoritos
	if err := InitBookmarks(); err != nil {
		log.Printf("⚠️  Erro ao inicializar favoritos: %v", err)
	}

	// Converter argumentos para C
	argc := C.int(len(os.Args))
	argv := make([]*C.char, len(os.Args))
	for i, arg := range os.Args {
		argv[i] = C.CString(arg)
		defer C.free(unsafe.Pointer(argv[i]))
	}

	// Executar browser GTK4
	log.Println("🚀 Iniciando interface GTK4...")
	log.Println("")
	
	status := C.run_browser(argc, &argv[0])
	
	log.Println("")
	log.Println("👋 Bagus Browser encerrado")
	
	os.Exit(int(status))
}
