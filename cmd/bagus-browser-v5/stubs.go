package main

import (
	"crypto/rand"
	"golang.org/x/crypto/bcrypt"
)

/*
#include <webkit2/webkit2.h>
*/
import "C"

// Stubs temporários para tipos desabilitados
type DownloadManager struct{}
type DownloadHandler struct{}

// WebView stub para compatibilidade
type WebView struct {
	cWebView *C.WebKitWebView
}

// Funções de senha
func generateSalt() ([]byte, error) {
	salt := make([]byte, 32)
	_, err := rand.Read(salt)
	return salt, err
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func verifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// showErrorDialog stub
func showErrorDialog(title, message string) {
	// TODO: Implementar diálogo GTK4
}
