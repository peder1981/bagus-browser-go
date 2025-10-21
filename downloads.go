package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// DownloadManager gerencia downloads
type DownloadManager struct {
	downloadPath string
}

// NewDownloadManager cria um novo gerenciador de downloads
func NewDownloadManager() (*DownloadManager, error) {
	// Diretório padrão de downloads
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	
	downloadPath := filepath.Join(homeDir, "Downloads")
	
	// Criar diretório se não existir
	if err := os.MkdirAll(downloadPath, 0755); err != nil {
		return nil, err
	}
	
	return &DownloadManager{
		downloadPath: downloadPath,
	}, nil
}

// GetDownloadPath retorna o caminho de downloads
func (dm *DownloadManager) GetDownloadPath() string {
	return dm.downloadPath
}

// SetDownloadPath define um novo caminho de downloads
func (dm *DownloadManager) SetDownloadPath(path string) error {
	// Verificar se diretório existe
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// Criar se não existir
		if err := os.MkdirAll(path, 0755); err != nil {
			return err
		}
	}
	
	dm.downloadPath = path
	log.Printf("📁 Pasta de downloads: %s", path)
	return nil
}

// GetUniqueFilename retorna um nome de arquivo único
func (dm *DownloadManager) GetUniqueFilename(filename string) string {
	fullPath := filepath.Join(dm.downloadPath, filename)
	
	// Se arquivo não existe, retornar como está
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return fullPath
	}
	
	// Arquivo existe, adicionar número
	ext := filepath.Ext(filename)
	name := filename[:len(filename)-len(ext)]
	
	for i := 1; i < 1000; i++ {
		newFilename := fmt.Sprintf("%s (%d)%s", name, i, ext)
		fullPath = filepath.Join(dm.downloadPath, newFilename)
		
		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			return fullPath
		}
	}
	
	// Fallback
	return filepath.Join(dm.downloadPath, fmt.Sprintf("%s_%d%s", name, os.Getpid(), ext))
}
