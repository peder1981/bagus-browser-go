package main

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	
	"github.com/gotk3/gotk3/gtk"
	"golang.org/x/crypto/bcrypt"
)

// generateSalt gera salt aleatório
func generateSalt() (string, error) {
	salt := make([]byte, 32)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(salt), nil
}

// hashPassword cria hash bcrypt da senha
func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// verifyPassword verifica senha contra hash
func verifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// showPasswordDialog mostra diálogo para senha mestre
func showPasswordDialog(title, message string) (string, bool) {
	// Criar diálogo
	dialog, err := gtk.DialogNew()
	if err != nil {
		log.Fatal("Erro ao criar diálogo:", err)
	}
	defer dialog.Destroy()
	
	dialog.SetTitle(title)
	dialog.SetDefaultSize(400, 150)
	dialog.SetModal(true)
	
	// Adicionar botões
	dialog.AddButton("Cancelar", gtk.RESPONSE_CANCEL)
	dialog.AddButton("OK", gtk.RESPONSE_OK)
	
	// Área de conteúdo
	contentArea, err := dialog.GetContentArea()
	if err != nil {
		log.Fatal("Erro ao obter área de conteúdo:", err)
	}
	
	// Label com mensagem
	label, err := gtk.LabelNew(message)
	if err != nil {
		log.Fatal("Erro ao criar label:", err)
	}
	label.SetMarginTop(10)
	label.SetMarginBottom(10)
	contentArea.Add(label)
	
	// Entry para senha
	entry, err := gtk.EntryNew()
	if err != nil {
		log.Fatal("Erro ao criar entry:", err)
	}
	entry.SetVisibility(false) // Ocultar senha
	entry.SetInputPurpose(gtk.INPUT_PURPOSE_PASSWORD)
	entry.SetMarginStart(10)
	entry.SetMarginEnd(10)
	entry.SetMarginBottom(10)
	contentArea.Add(entry)
	
	// Focar no entry
	entry.GrabFocus()
	
	// Enter ativa OK
	entry.Connect("activate", func() {
		dialog.Response(gtk.RESPONSE_OK)
	})
	
	dialog.ShowAll()
	
	// Executar diálogo
	response := dialog.Run()
	
	if response == gtk.RESPONSE_OK {
		text, err := entry.GetText()
		if err != nil {
			return "", false
		}
		return text, true
	}
	
	return "", false
}

// showSetPasswordDialog mostra diálogo para definir senha
func showSetPasswordDialog() (string, bool) {
	// Criar diálogo
	dialog, err := gtk.DialogNew()
	if err != nil {
		log.Fatal("Erro ao criar diálogo:", err)
	}
	defer dialog.Destroy()
	
	dialog.SetTitle("Definir Senha Mestre")
	dialog.SetDefaultSize(400, 200)
	dialog.SetModal(true)
	
	// Adicionar botões
	dialog.AddButton("Cancelar", gtk.RESPONSE_CANCEL)
	dialog.AddButton("Definir", gtk.RESPONSE_OK)
	
	// Área de conteúdo
	contentArea, err := dialog.GetContentArea()
	if err != nil {
		log.Fatal("Erro ao obter área de conteúdo:", err)
	}
	
	// Label
	label, err := gtk.LabelNew("Digite a senha mestre:")
	if err != nil {
		log.Fatal("Erro ao criar label:", err)
	}
	label.SetMarginTop(10)
	label.SetMarginBottom(5)
	contentArea.Add(label)
	
	// Entry 1
	entry1, err := gtk.EntryNew()
	if err != nil {
		log.Fatal("Erro ao criar entry:", err)
	}
	entry1.SetVisibility(false)
	entry1.SetInputPurpose(gtk.INPUT_PURPOSE_PASSWORD)
	entry1.SetMarginStart(10)
	entry1.SetMarginEnd(10)
	entry1.SetMarginBottom(5)
	contentArea.Add(entry1)
	
	// Label 2
	label2, err := gtk.LabelNew("Confirme a senha:")
	if err != nil {
		log.Fatal("Erro ao criar label:", err)
	}
	label2.SetMarginBottom(5)
	contentArea.Add(label2)
	
	// Entry 2
	entry2, err := gtk.EntryNew()
	if err != nil {
		log.Fatal("Erro ao criar entry:", err)
	}
	entry2.SetVisibility(false)
	entry2.SetInputPurpose(gtk.INPUT_PURPOSE_PASSWORD)
	entry2.SetMarginStart(10)
	entry2.SetMarginEnd(10)
	entry2.SetMarginBottom(10)
	contentArea.Add(entry2)
	
	entry1.GrabFocus()
	
	dialog.ShowAll()
	
	// Loop até senhas coincidirem
	for {
		response := dialog.Run()
		
		if response != gtk.RESPONSE_OK {
			return "", false
		}
		
		pass1, _ := entry1.GetText()
		pass2, _ := entry2.GetText()
		
		if pass1 == "" {
			showErrorDialog("Erro", "Senha não pode ser vazia!")
			continue
		}
		
		if pass1 != pass2 {
			showErrorDialog("Erro", "Senhas não coincidem!")
			entry2.SetText("")
			entry2.GrabFocus()
			continue
		}
		
		return pass1, true
	}
}

// authenticateUser autentica usuário com senha mestre
func authenticateUser(config *Config) bool {
	if !config.RequirePassword {
		return true // Sem senha configurada
	}
	
	maxAttempts := 3
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		password, ok := showPasswordDialog(
			"Bagus Browser - Autenticação",
			"Digite a senha mestre:",
		)
		
		if !ok {
			return false // Usuário cancelou
		}
		
		if config.VerifyPassword(password) {
			log.Println("🔓 Autenticação bem-sucedida")
			return true
		}
		
		if attempt < maxAttempts {
			showErrorDialog(
				"Senha Incorreta",
				"Senha incorreta. Tente novamente.\n" +
				"Tentativas restantes: " + string(rune('0'+maxAttempts-attempt)),
			)
		}
	}
	
	showErrorDialog(
		"Acesso Negado",
		"Número máximo de tentativas excedido.\nO browser será fechado.",
	)
	
	return false
}

// showErrorDialog mostra diálogo de erro
func showErrorDialog(title, message string) {
	dialog := gtk.MessageDialogNew(
		nil,
		gtk.DIALOG_MODAL,
		gtk.MESSAGE_ERROR,
		gtk.BUTTONS_OK,
		message,
	)
	dialog.SetTitle(title)
	dialog.Run()
	dialog.Destroy()
}
