package main

/*
#cgo pkg-config: gtk+-3.0
#cgo CFLAGS: -Wno-deprecated-declarations
#include <gtk/gtk.h>
#include <stdlib.h>
#include <string.h>

// Diálogo de senha
static GtkWidget* create_password_dialog(GtkWindow *parent, const char *title, const char *message) {
    GtkWidget *dialog = gtk_dialog_new_with_buttons(
        title,
        parent,
        GTK_DIALOG_MODAL | GTK_DIALOG_DESTROY_WITH_PARENT,
        "Cancelar", GTK_RESPONSE_CANCEL,
        "OK", GTK_RESPONSE_OK,
        NULL
    );
    
    GtkWidget *content = gtk_dialog_get_content_area(GTK_DIALOG(dialog));
    gtk_container_set_border_width(GTK_CONTAINER(content), 15);
    
    // Mensagem
    GtkWidget *label = gtk_label_new(message);
    gtk_label_set_line_wrap(GTK_LABEL(label), TRUE);
    gtk_box_pack_start(GTK_BOX(content), label, FALSE, FALSE, 10);
    
    // Campo de senha
    GtkWidget *entry = gtk_entry_new();
    gtk_entry_set_visibility(GTK_ENTRY(entry), FALSE);
    gtk_entry_set_activates_default(GTK_ENTRY(entry), TRUE);
    gtk_box_pack_start(GTK_BOX(content), entry, FALSE, FALSE, 5);
    
    // Armazenar entry como data do dialog
    g_object_set_data(G_OBJECT(dialog), "password_entry", entry);
    
    gtk_dialog_set_default_response(GTK_DIALOG(dialog), GTK_RESPONSE_OK);
    gtk_widget_show_all(content);
    
    return dialog;
}

// Obter senha do diálogo
static const char* get_password_from_dialog(GtkWidget *dialog) {
    GtkWidget *entry = g_object_get_data(G_OBJECT(dialog), "password_entry");
    if (entry) {
        return gtk_entry_get_text(GTK_ENTRY(entry));
    }
    return NULL;
}

// Wrapper para gtk_dialog_run
static gint run_dialog(GtkWidget *dialog) {
    return gtk_dialog_run(GTK_DIALOG(dialog));
}

// Wrapper para gtk_window_set_title
static void set_window_title(GtkWidget *window, const char *title) {
    gtk_window_set_title(GTK_WINDOW(window), title);
}

// Wrapper para gtk_message_dialog_new (error)
static GtkWidget* create_error_dialog(GtkWindow *parent, const char *message) {
    return gtk_message_dialog_new(
        parent,
        GTK_DIALOG_MODAL,
        GTK_MESSAGE_ERROR,
        GTK_BUTTONS_OK,
        "%s", message
    );
}

// Wrapper para gtk_message_dialog_new (info)
static GtkWidget* create_info_dialog(GtkWindow *parent, const char *message) {
    return gtk_message_dialog_new(
        parent,
        GTK_DIALOG_MODAL,
        GTK_MESSAGE_INFO,
        GTK_BUTTONS_OK,
        "%s", message
    );
}
*/
import "C"

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"unsafe"

	"golang.org/x/crypto/bcrypt"
)

// AuthConfig armazena a configuração de autenticação
type AuthConfig struct {
	Enabled      bool   `json:"enabled"`
	PasswordHash string `json:"password_hash"`
}

// AuthManager gerencia autenticação
type AuthManager struct {
	configPath string
	config     *AuthConfig
}

var globalAuthManager *AuthManager

// NewAuthManager cria um novo gerenciador de autenticação
func NewAuthManager() (*AuthManager, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	configDir := filepath.Join(homeDir, ".config", "bagus-browser")
	if err := os.MkdirAll(configDir, 0700); err != nil {
		return nil, err
	}

	configPath := filepath.Join(configDir, "auth.json")

	am := &AuthManager{
		configPath: configPath,
	}

	// Carregar configuração
	if err := am.load(); err != nil {
		// Se não existe, criar config padrão
		am.config = &AuthConfig{
			Enabled:      false,
			PasswordHash: "",
		}
		am.save()
	}

	return am, nil
}

// load carrega a configuração
func (am *AuthManager) load() error {
	data, err := os.ReadFile(am.configPath)
	if err != nil {
		return err
	}

	var config AuthConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return err
	}

	am.config = &config
	return nil
}

// save salva a configuração
func (am *AuthManager) save() error {
	data, err := json.Marshal(am.config)
	if err != nil {
		return err
	}

	return os.WriteFile(am.configPath, data, 0600)
}

// IsEnabled retorna se a senha está habilitada
func (am *AuthManager) IsEnabled() bool {
	return am.config.Enabled
}

// HasPassword retorna se há senha definida
func (am *AuthManager) HasPassword() bool {
	return am.config.PasswordHash != ""
}

// SetPassword define uma nova senha
func (am *AuthManager) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	am.config.PasswordHash = string(hash)
	am.config.Enabled = true

	return am.save()
}

// VerifyPassword verifica se a senha está correta
func (am *AuthManager) VerifyPassword(password string) bool {
	if am.config.PasswordHash == "" {
		return false
	}

	err := bcrypt.CompareHashAndPassword([]byte(am.config.PasswordHash), []byte(password))
	return err == nil
}

// RemovePassword remove a senha
func (am *AuthManager) RemovePassword() error {
	am.config.PasswordHash = ""
	am.config.Enabled = false
	return am.save()
}

// ChangePassword altera a senha (verifica a antiga primeiro)
func (am *AuthManager) ChangePassword(oldPassword, newPassword string) error {
	if !am.VerifyPassword(oldPassword) {
		return fmt.Errorf("senha atual incorreta")
	}

	return am.SetPassword(newPassword)
}

//export ShowPasswordDialog
func ShowPasswordDialog(parent unsafe.Pointer) int {
	if globalAuthManager == nil || !globalAuthManager.IsEnabled() {
		return 1 // Sucesso (sem senha)
	}

	cTitle := C.CString("🔒 Senha Mestre - Bagus Browser")
	cMessage := C.CString("Digite a senha mestre para acessar o browser:")
	defer C.free(unsafe.Pointer(cTitle))
	defer C.free(unsafe.Pointer(cMessage))

	attempts := 0
	maxAttempts := 3

	for attempts < maxAttempts {
		dialog := C.create_password_dialog((*C.GtkWindow)(parent), cTitle, cMessage)
		response := C.run_dialog(dialog)

		if response == C.GTK_RESPONSE_OK {
			cPassword := C.get_password_from_dialog(dialog)
			password := C.GoString(cPassword)

			C.gtk_widget_destroy(dialog)

			if globalAuthManager.VerifyPassword(password) {
				log.Println("✅ Autenticação bem-sucedida")
				return 1 // Sucesso
			}

			attempts++
			remaining := maxAttempts - attempts

			if remaining > 0 {
				cTitle = C.CString("❌ Senha Incorreta")
				cMessage = C.CString(fmt.Sprintf("Senha incorreta! Tentativas restantes: %d", remaining))
			}
		} else {
			C.gtk_widget_destroy(dialog)
			return 0 // Cancelado
		}
	}

	log.Println("❌ Máximo de tentativas excedido")
	return 0 // Falha
}

//export ShowSetPasswordDialog
func ShowSetPasswordDialog(parent unsafe.Pointer) {
	if globalAuthManager == nil {
		return
	}

	cTitle := C.CString("🔒 Definir Senha Mestre")
	cMessage1 := C.CString("Digite a nova senha mestre:")
	cMessage2 := C.CString("Confirme a senha mestre:")
	defer C.free(unsafe.Pointer(cTitle))
	defer C.free(unsafe.Pointer(cMessage1))
	defer C.free(unsafe.Pointer(cMessage2))

	// Primeira senha
	dialog1 := C.create_password_dialog((*C.GtkWindow)(parent), cTitle, cMessage1)
	response1 := C.run_dialog(dialog1)

	if response1 != C.GTK_RESPONSE_OK {
		C.gtk_widget_destroy(dialog1)
		return
	}

	cPassword1 := C.get_password_from_dialog(dialog1)
	password1 := C.GoString(cPassword1)
	C.gtk_widget_destroy(dialog1)

	if len(password1) < 4 {
		showError(parent, "Senha muito curta", "A senha deve ter pelo menos 4 caracteres.")
		return
	}

	// Confirmar senha
	dialog2 := C.create_password_dialog((*C.GtkWindow)(parent), cTitle, cMessage2)
	response2 := C.run_dialog(dialog2)

	if response2 != C.GTK_RESPONSE_OK {
		C.gtk_widget_destroy(dialog2)
		return
	}

	cPassword2 := C.get_password_from_dialog(dialog2)
	password2 := C.GoString(cPassword2)
	C.gtk_widget_destroy(dialog2)

	if password1 != password2 {
		showError(parent, "Senhas não conferem", "As senhas digitadas não são iguais.")
		return
	}

	if err := globalAuthManager.SetPassword(password1); err != nil {
		showError(parent, "Erro", fmt.Sprintf("Erro ao definir senha: %v", err))
		return
	}

	showInfo(parent, "Sucesso", "Senha mestre definida com sucesso!\nO browser solicitará a senha na próxima inicialização.")
	log.Println("✅ Senha mestre definida")
}

//export ShowChangePasswordDialog
func ShowChangePasswordDialog(parent unsafe.Pointer) {
	if globalAuthManager == nil || !globalAuthManager.HasPassword() {
		showError(parent, "Erro", "Nenhuma senha definida. Use 'Definir Senha' primeiro.")
		return
	}

	cTitle := C.CString("🔒 Alterar Senha Mestre")
	cMessage1 := C.CString("Digite a senha atual:")
	cMessage2 := C.CString("Digite a nova senha:")
	cMessage3 := C.CString("Confirme a nova senha:")
	defer C.free(unsafe.Pointer(cTitle))
	defer C.free(unsafe.Pointer(cMessage1))
	defer C.free(unsafe.Pointer(cMessage2))
	defer C.free(unsafe.Pointer(cMessage3))

	// Senha atual
	dialog1 := C.create_password_dialog((*C.GtkWindow)(parent), cTitle, cMessage1)
	response1 := C.run_dialog(dialog1)

	if response1 != C.GTK_RESPONSE_OK {
		C.gtk_widget_destroy(dialog1)
		return
	}

	cOldPassword := C.get_password_from_dialog(dialog1)
	oldPassword := C.GoString(cOldPassword)
	C.gtk_widget_destroy(dialog1)

	// Nova senha
	dialog2 := C.create_password_dialog((*C.GtkWindow)(parent), cTitle, cMessage2)
	response2 := C.run_dialog(dialog2)

	if response2 != C.GTK_RESPONSE_OK {
		C.gtk_widget_destroy(dialog2)
		return
	}

	cNewPassword1 := C.get_password_from_dialog(dialog2)
	newPassword1 := C.GoString(cNewPassword1)
	C.gtk_widget_destroy(dialog2)

	if len(newPassword1) < 4 {
		showError(parent, "Senha muito curta", "A senha deve ter pelo menos 4 caracteres.")
		return
	}

	// Confirmar nova senha
	dialog3 := C.create_password_dialog((*C.GtkWindow)(parent), cTitle, cMessage3)
	response3 := C.run_dialog(dialog3)

	if response3 != C.GTK_RESPONSE_OK {
		C.gtk_widget_destroy(dialog3)
		return
	}

	cNewPassword2 := C.get_password_from_dialog(dialog3)
	newPassword2 := C.GoString(cNewPassword2)
	C.gtk_widget_destroy(dialog3)

	if newPassword1 != newPassword2 {
		showError(parent, "Senhas não conferem", "As senhas digitadas não são iguais.")
		return
	}

	if err := globalAuthManager.ChangePassword(oldPassword, newPassword1); err != nil {
		showError(parent, "Erro", err.Error())
		return
	}

	showInfo(parent, "Sucesso", "Senha alterada com sucesso!")
	log.Println("✅ Senha mestre alterada")
}

//export ShowRemovePasswordDialog
func ShowRemovePasswordDialog(parent unsafe.Pointer) {
	if globalAuthManager == nil || !globalAuthManager.HasPassword() {
		showError(parent, "Erro", "Nenhuma senha definida.")
		return
	}

	cTitle := C.CString("🔒 Remover Senha Mestre")
	cMessage := C.CString("Digite a senha atual para remover:")
	defer C.free(unsafe.Pointer(cTitle))
	defer C.free(unsafe.Pointer(cMessage))

	dialog := C.create_password_dialog((*C.GtkWindow)(parent), cTitle, cMessage)
	response := C.run_dialog(dialog)

	if response != C.GTK_RESPONSE_OK {
		C.gtk_widget_destroy(dialog)
		return
	}

	cPassword := C.get_password_from_dialog(dialog)
	password := C.GoString(cPassword)
	C.gtk_widget_destroy(dialog)

	if !globalAuthManager.VerifyPassword(password) {
		showError(parent, "Erro", "Senha incorreta!")
		return
	}

	if err := globalAuthManager.RemovePassword(); err != nil {
		showError(parent, "Erro", fmt.Sprintf("Erro ao remover senha: %v", err))
		return
	}

	showInfo(parent, "Sucesso", "Senha mestre removida com sucesso!")
	log.Println("✅ Senha mestre removida")
}

// Funções auxiliares para diálogos
func showError(parent unsafe.Pointer, title, message string) {
	cTitle := C.CString(title)
	cMessage := C.CString(message)
	defer C.free(unsafe.Pointer(cTitle))
	defer C.free(unsafe.Pointer(cMessage))

	dialog := C.create_error_dialog((*C.GtkWindow)(parent), cMessage)
	C.set_window_title(dialog, cTitle)
	C.run_dialog(dialog)
	C.gtk_widget_destroy(dialog)
}

func showInfo(parent unsafe.Pointer, title, message string) {
	cTitle := C.CString(title)
	cMessage := C.CString(message)
	defer C.free(unsafe.Pointer(cTitle))
	defer C.free(unsafe.Pointer(cMessage))

	dialog := C.create_info_dialog((*C.GtkWindow)(parent), cMessage)
	C.set_window_title(dialog, cTitle)
	C.run_dialog(dialog)
	C.gtk_widget_destroy(dialog)
}
