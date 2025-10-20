package browser

import (
	"fmt"
	"os"
	"path/filepath"
)

// StorageManager gerencia armazenamento do browser
type StorageManager struct {
	basePath string
}

// NewStorageManager cria novo gerenciador de armazenamento
func NewStorageManager(basePath string) (*StorageManager, error) {
	// Valida path
	if basePath == "" {
		return nil, fmt.Errorf("basePath não pode ser vazio")
	}

	// Cria diretório base se não existir
	if err := os.MkdirAll(basePath, 0700); err != nil {
		return nil, fmt.Errorf("erro ao criar diretório base: %w", err)
	}

	sm := &StorageManager{
		basePath: basePath,
	}

	// Cria subdiretórios necessários
	dirs := []string{
		"default",
		"log",
		"analyze",
		filepath.Join("analyze", "pending"),
		"cache",
		"downloads",
	}

	for _, dir := range dirs {
		path := filepath.Join(basePath, dir)
		if err := os.MkdirAll(path, 0700); err != nil {
			return nil, fmt.Errorf("erro ao criar diretório %s: %w", dir, err)
		}
	}

	return sm, nil
}

// GetBasePath retorna caminho base
func (sm *StorageManager) GetBasePath() string {
	return sm.basePath
}

// GetDataDir retorna diretório de dados
func (sm *StorageManager) GetDataDir() string {
	return filepath.Join(sm.basePath, "default")
}

// GetCacheDir retorna diretório de cache
func (sm *StorageManager) GetCacheDir() string {
	return filepath.Join(sm.basePath, "cache")
}

// GetLogDir retorna diretório de logs
func (sm *StorageManager) GetLogDir() string {
	return filepath.Join(sm.basePath, "log")
}

// GetDownloadsDir retorna diretório de downloads
func (sm *StorageManager) GetDownloadsDir() string {
	return filepath.Join(sm.basePath, "downloads")
}

// GetAnalyzeDir retorna diretório de análise
func (sm *StorageManager) GetAnalyzeDir() string {
	return filepath.Join(sm.basePath, "analyze")
}

// ClearCache limpa cache
func (sm *StorageManager) ClearCache() error {
	cacheDir := sm.GetCacheDir()
	if err := os.RemoveAll(cacheDir); err != nil {
		return fmt.Errorf("erro ao limpar cache: %w", err)
	}
	return os.MkdirAll(cacheDir, 0700)
}

// Close fecha storage manager
func (sm *StorageManager) Close() error {
	// Nada a fazer por enquanto
	return nil
}
