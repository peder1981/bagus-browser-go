package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

// Bookmark representa um favorito
type Bookmark struct {
	Title   string    `json:"title"`
	URL     string    `json:"url"`
	Added   time.Time `json:"added"`
	Favicon string    `json:"favicon,omitempty"`
}

// BookmarkManager gerencia favoritos com criptografia
type BookmarkManager struct {
	bookmarks   []Bookmark
	crypto      *CryptoManager
	storagePath string
}

// NewBookmarkManager cria um novo gerenciador de favoritos
func NewBookmarkManager(crypto *CryptoManager) (*BookmarkManager, error) {
	// Diretório de configuração
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	
	configDir := filepath.Join(homeDir, ".config", "bagus-browser")
	if err := os.MkdirAll(configDir, 0700); err != nil {
		return nil, err
	}
	
	storagePath := filepath.Join(configDir, "bookmarks.enc")
	
	bm := &BookmarkManager{
		bookmarks:   make([]Bookmark, 0),
		crypto:      crypto,
		storagePath: storagePath,
	}
	
	// Carregar favoritos existentes
	if err := bm.Load(); err != nil {
		log.Printf("⚠️  Não foi possível carregar favoritos: %v", err)
		// Não é erro fatal, apenas começa com lista vazia
	}
	
	return bm, nil
}

// Add adiciona um novo favorito
func (bm *BookmarkManager) Add(title, url string) error {
	// Verificar se já existe
	for _, b := range bm.bookmarks {
		if b.URL == url {
			return fmt.Errorf("favorito já existe")
		}
	}
	
	bookmark := Bookmark{
		Title: title,
		URL:   url,
		Added: time.Now(),
	}
	
	bm.bookmarks = append(bm.bookmarks, bookmark)
	
	log.Printf("⭐ Favorito adicionado: %s", title)
	
	return bm.Save()
}

// Remove remove um favorito
func (bm *BookmarkManager) Remove(url string) error {
	for i, b := range bm.bookmarks {
		if b.URL == url {
			bm.bookmarks = append(bm.bookmarks[:i], bm.bookmarks[i+1:]...)
			log.Printf("🗑️  Favorito removido: %s", b.Title)
			return bm.Save()
		}
	}
	return fmt.Errorf("favorito não encontrado")
}

// GetAll retorna todos os favoritos
func (bm *BookmarkManager) GetAll() []Bookmark {
	return bm.bookmarks
}

// Exists verifica se URL está nos favoritos
func (bm *BookmarkManager) Exists(url string) bool {
	for _, b := range bm.bookmarks {
		if b.URL == url {
			return true
		}
	}
	return false
}

// Save salva favoritos criptografados
func (bm *BookmarkManager) Save() error {
	// Serializar para JSON
	data, err := json.Marshal(bm.bookmarks)
	if err != nil {
		return err
	}
	
	// Criptografar
	encrypted, err := bm.crypto.Encrypt(data)
	if err != nil {
		return err
	}
	
	// Salvar em arquivo
	if err := os.WriteFile(bm.storagePath, []byte(encrypted), 0600); err != nil {
		return err
	}
	
	log.Printf("💾 Favoritos salvos (criptografados): %d itens", len(bm.bookmarks))
	return nil
}

// Load carrega favoritos criptografados
func (bm *BookmarkManager) Load() error {
	// Verificar se arquivo existe
	if _, err := os.Stat(bm.storagePath); os.IsNotExist(err) {
		// Arquivo não existe, começar com lista vazia
		return nil
	}
	
	// Ler arquivo
	encrypted, err := os.ReadFile(bm.storagePath)
	if err != nil {
		return err
	}
	
	// Descriptografar
	data, err := bm.crypto.Decrypt(string(encrypted))
	if err != nil {
		return fmt.Errorf("falha ao descriptografar favoritos: %v", err)
	}
	
	// Deserializar JSON
	if err := json.Unmarshal(data, &bm.bookmarks); err != nil {
		return err
	}
	
	log.Printf("📚 Favoritos carregados: %d itens", len(bm.bookmarks))
	return nil
}

// Clear limpa todos os favoritos
func (bm *BookmarkManager) Clear() error {
	bm.bookmarks = make([]Bookmark, 0)
	return bm.Save()
}

// Export exporta favoritos para JSON (descriptografado)
func (bm *BookmarkManager) Export(path string) error {
	data, err := json.MarshalIndent(bm.bookmarks, "", "  ")
	if err != nil {
		return err
	}
	
	return os.WriteFile(path, data, 0600)
}

// Import importa favoritos de JSON
func (bm *BookmarkManager) Import(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	
	var imported []Bookmark
	if err := json.Unmarshal(data, &imported); err != nil {
		return err
	}
	
	// Adicionar favoritos importados (evitar duplicatas)
	for _, b := range imported {
		if !bm.Exists(b.URL) {
			bm.bookmarks = append(bm.bookmarks, b)
		}
	}
	
	return bm.Save()
}
