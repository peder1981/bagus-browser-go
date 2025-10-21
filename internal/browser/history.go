package browser

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/peder1981/bagus-browser-go/internal/security"
)

const (
	maxHistorySize = 10000 // Máximo de entradas no histórico
)

// HistoryEntry representa uma entrada no histórico
type HistoryEntry struct {
	URL       string    `json:"url"`
	Title     string    `json:"title"`
	Timestamp time.Time `json:"timestamp"`
}

// History gerencia o histórico de navegação
type History struct {
	entries   []HistoryEntry
	path      string
	encryptor *security.Encryptor
	mu        sync.RWMutex
}

// NewHistory cria novo gerenciador de histórico
func NewHistory(storagePath string) (*History, error) {
	// Cria encryptor para este usuário
	encryptor, err := security.NewEncryptor(storagePath)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar encryptor: %w", err)
	}

	h := &History{
		entries:   make([]HistoryEntry, 0),
		path:      filepath.Join(storagePath, "history.encrypted"),
		encryptor: encryptor,
	}

	// Tenta migrar histórico antigo não criptografado
	oldPath := filepath.Join(storagePath, "history.json")
	if _, err := os.Stat(oldPath); err == nil {
		if err := h.migrateFromUnencrypted(oldPath); err != nil {
			fmt.Printf("Aviso: erro ao migrar histórico: %v\n", err)
		}
	}

	// Carrega histórico existente
	if err := h.Load(); err != nil {
		// Não é erro fatal se não conseguir carregar
		fmt.Printf("Aviso: não foi possível carregar histórico: %v\n", err)
	}

	return h, nil
}

// Add adiciona entrada ao histórico
func (h *History) Add(url, title string) {
	h.mu.Lock()
	defer h.mu.Unlock()

	// Não adiciona URLs vazias
	if url == "" {
		return
	}

	// Não adiciona duplicatas consecutivas
	if len(h.entries) > 0 && h.entries[len(h.entries)-1].URL == url {
		return
	}

	entry := HistoryEntry{
		URL:       url,
		Title:     title,
		Timestamp: time.Now(),
	}

	h.entries = append(h.entries, entry)

	// Limita tamanho do histórico
	if len(h.entries) > maxHistorySize {
		h.entries = h.entries[len(h.entries)-maxHistorySize:]
	}
}

// Search busca no histórico
func (h *History) Search(query string) []HistoryEntry {
	h.mu.RLock()
	defer h.mu.RUnlock()

	if query == "" {
		return nil
	}

	results := make([]HistoryEntry, 0)
	query = toLower(query)

	// Busca nas últimas 1000 entradas
	start := 0
	if len(h.entries) > 1000 {
		start = len(h.entries) - 1000
	}

	for i := len(h.entries) - 1; i >= start && len(results) < 50; i-- {
		entry := h.entries[i]
		if contains(toLower(entry.URL), query) || contains(toLower(entry.Title), query) {
			results = append(results, entry)
		}
	}

	return results
}

// GetRecent retorna entradas recentes
func (h *History) GetRecent(limit int) []HistoryEntry {
	h.mu.RLock()
	defer h.mu.RUnlock()

	if limit <= 0 {
		limit = 50
	}

	start := len(h.entries) - limit
	if start < 0 {
		start = 0
	}

	// Retorna em ordem reversa (mais recente primeiro)
	result := make([]HistoryEntry, 0, limit)
	for i := len(h.entries) - 1; i >= start; i-- {
		result = append(result, h.entries[i])
	}

	return result
}

// Clear limpa todo o histórico
func (h *History) Clear() error {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.entries = make([]HistoryEntry, 0)
	return h.save()
}

// Save salva histórico em disco
func (h *History) Save() error {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.save()
}

// save (interno, sem lock)
func (h *History) save() error {
	// Limita tamanho antes de salvar
	entriesToSave := h.entries
	if len(entriesToSave) > maxHistorySize {
		entriesToSave = entriesToSave[len(entriesToSave)-maxHistorySize:]
	}

	// Serializa para JSON
	data, err := json.Marshal(entriesToSave)
	if err != nil {
		return fmt.Errorf("erro ao serializar histórico: %w", err)
	}

	// Criptografa dados
	encryptedData, err := h.encryptor.Encrypt(data)
	if err != nil {
		return fmt.Errorf("erro ao criptografar histórico: %w", err)
	}

	// Cria diretório se não existir
	dir := filepath.Dir(h.path)
	if err := os.MkdirAll(dir, 0700); err != nil {
		return fmt.Errorf("erro ao criar diretório: %w", err)
	}

	// Escreve arquivo criptografado
	if err := os.WriteFile(h.path, []byte(encryptedData), 0600); err != nil {
		return fmt.Errorf("erro ao salvar histórico: %w", err)
	}

	return nil
}

// Load carrega histórico do disco
func (h *History) Load() error {
	// Verifica se arquivo existe
	if _, err := os.Stat(h.path); os.IsNotExist(err) {
		return nil // Não é erro se não existir
	}

	// Valida tamanho (max 20MB criptografado)
	info, err := os.Stat(h.path)
	if err != nil {
		return err
	}
	if info.Size() > 20*1024*1024 {
		return fmt.Errorf("arquivo de histórico muito grande")
	}

	// Lê arquivo criptografado
	encryptedData, err := os.ReadFile(h.path)
	if err != nil {
		return fmt.Errorf("erro ao ler histórico: %w", err)
	}

	// Descriptografa
	data, err := h.encryptor.Decrypt(string(encryptedData))
	if err != nil {
		return fmt.Errorf("erro ao descriptografar histórico: %w", err)
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	var entries []HistoryEntry
	if err := json.Unmarshal(data, &entries); err != nil {
		return fmt.Errorf("erro ao parsear histórico: %w", err)
	}

	// Limita tamanho
	if len(entries) > maxHistorySize {
		entries = entries[len(entries)-maxHistorySize:]
	}

	h.entries = entries
	return nil
}

// migrateFromUnencrypted migra histórico antigo não criptografado
func (h *History) migrateFromUnencrypted(oldPath string) error {
	// Lê arquivo antigo
	data, err := os.ReadFile(oldPath)
	if err != nil {
		return fmt.Errorf("erro ao ler histórico antigo: %w", err)
	}

	// Parse JSON
	var entries []HistoryEntry
	if err := json.Unmarshal(data, &entries); err != nil {
		return fmt.Errorf("erro ao parsear histórico antigo: %w", err)
	}

	// Salva em formato criptografado
	h.mu.Lock()
	h.entries = entries
	h.mu.Unlock()

	if err := h.Save(); err != nil {
		return fmt.Errorf("erro ao salvar histórico criptografado: %w", err)
	}

	// Remove arquivo antigo
	if err := os.Remove(oldPath); err != nil {
		fmt.Printf("Aviso: não foi possível remover histórico antigo: %v\n", err)
	}

	fmt.Println("Histórico migrado com sucesso para formato criptografado")
	return nil
}

// Count retorna número de entradas
func (h *History) Count() int {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return len(h.entries)
}

// Helper functions
func toLower(s string) string {
	return string([]rune(s)) // Simplificado
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || findSubstring(s, substr))
}

func findSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
