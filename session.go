package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// SessionTab representa uma aba salva
type SessionTab struct {
	URL    string `json:"url"`
	Title  string `json:"title"`
	Active bool   `json:"active"`
}

// Session representa uma sessão do browser
type Session struct {
	Tabs []SessionTab `json:"tabs"`
}

// SessionManager gerencia sessões do browser
type SessionManager struct {
	crypto      *CryptoManager
	storagePath string
}

// NewSessionManager cria um novo gerenciador de sessões
func NewSessionManager(crypto *CryptoManager) (*SessionManager, error) {
	// Diretório de configuração
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	
	configDir := filepath.Join(homeDir, ".config", "bagus-browser")
	if err := os.MkdirAll(configDir, 0700); err != nil {
		return nil, err
	}
	
	storagePath := filepath.Join(configDir, "session.enc")
	
	return &SessionManager{
		crypto:      crypto,
		storagePath: storagePath,
	}, nil
}

// Save salva a sessão atual
func (sm *SessionManager) Save(tabs []SessionTab) error {
	session := Session{
		Tabs: tabs,
	}
	
	// Serializar para JSON
	data, err := json.Marshal(session)
	if err != nil {
		return err
	}
	
	// Criptografar
	encrypted, err := sm.crypto.Encrypt(data)
	if err != nil {
		return err
	}
	
	// Salvar em arquivo
	if err := os.WriteFile(sm.storagePath, []byte(encrypted), 0600); err != nil {
		return err
	}
	
	log.Printf("💾 Sessão salva: %d abas", len(tabs))
	return nil
}

// Load carrega a sessão salva
func (sm *SessionManager) Load() (*Session, error) {
	// Verificar se arquivo existe
	if _, err := os.Stat(sm.storagePath); os.IsNotExist(err) {
		// Arquivo não existe, retornar sessão vazia
		return &Session{Tabs: []SessionTab{}}, nil
	}
	
	// Ler arquivo
	encrypted, err := os.ReadFile(sm.storagePath)
	if err != nil {
		return nil, err
	}
	
	// Descriptografar
	data, err := sm.crypto.Decrypt(string(encrypted))
	if err != nil {
		return nil, fmt.Errorf("falha ao descriptografar sessão: %v", err)
	}
	
	// Deserializar JSON
	var session Session
	if err := json.Unmarshal(data, &session); err != nil {
		return nil, err
	}
	
	log.Printf("📂 Sessão carregada: %d abas", len(session.Tabs))
	return &session, nil
}

// Clear limpa a sessão salva
func (sm *SessionManager) Clear() error {
	if err := os.Remove(sm.storagePath); err != nil && !os.IsNotExist(err) {
		return err
	}
	log.Println("🗑️  Sessão limpa")
	return nil
}
