package storage

import (
	"fmt"
	"os"
	"path/filepath"
)

// Manager gerencia o armazenamento de dados do navegador
type Manager struct {
	profile    string
	dataDir    string
	configPath string
}

// NewManager cria um novo gerenciador de armazenamento
func NewManager(profile, configPath string) (*Manager, error) {
	if profile == "" {
		profile = "default"
	}

	// Determina o diretório de dados
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("erro ao obter diretório home: %w", err)
	}

	dataDir := filepath.Join(homeDir, ".bagus", profile)

	// Cria o diretório se não existir
	if err := os.MkdirAll(dataDir, 0700); err != nil {
		return nil, fmt.Errorf("erro ao criar diretório de dados: %w", err)
	}

	// Cria subdiretórios necessários
	subdirs := []string{"history", "bookmarks", "downloads", "cache", "sessions"}
	for _, subdir := range subdirs {
		path := filepath.Join(dataDir, subdir)
		if err := os.MkdirAll(path, 0700); err != nil {
			return nil, fmt.Errorf("erro ao criar subdiretório %s: %w", subdir, err)
		}
	}

	// Define caminho de configuração
	if configPath == "" {
		configPath = filepath.Join(dataDir, "config.yaml")
	}

	manager := &Manager{
		profile:    profile,
		dataDir:    dataDir,
		configPath: configPath,
	}

	return manager, nil
}

// GetDataDir retorna o diretório de dados
func (m *Manager) GetDataDir() string {
	return m.dataDir
}

// GetProfile retorna o perfil atual
func (m *Manager) GetProfile() string {
	return m.profile
}

// GetConfigPath retorna o caminho do arquivo de configuração
func (m *Manager) GetConfigPath() string {
	return m.configPath
}

// GetHistoryDir retorna o diretório de histórico
func (m *Manager) GetHistoryDir() string {
	return filepath.Join(m.dataDir, "history")
}

// GetBookmarksDir retorna o diretório de favoritos
func (m *Manager) GetBookmarksDir() string {
	return filepath.Join(m.dataDir, "bookmarks")
}

// GetDownloadsDir retorna o diretório de downloads
func (m *Manager) GetDownloadsDir() string {
	return filepath.Join(m.dataDir, "downloads")
}

// GetCacheDir retorna o diretório de cache
func (m *Manager) GetCacheDir() string {
	return filepath.Join(m.dataDir, "cache")
}

// Close fecha o gerenciador de armazenamento
func (m *Manager) Close() error {
	// TODO: Implementar limpeza de recursos se necessário
	return nil
}

// ClearCache limpa o cache do navegador
func (m *Manager) ClearCache() error {
	cacheDir := m.GetCacheDir()
	if err := os.RemoveAll(cacheDir); err != nil {
		return fmt.Errorf("erro ao limpar cache: %w", err)
	}
	return os.MkdirAll(cacheDir, 0700)
}

// ClearHistory limpa o histórico do navegador
func (m *Manager) ClearHistory() error {
	historyDir := m.GetHistoryDir()
	if err := os.RemoveAll(historyDir); err != nil {
		return fmt.Errorf("erro ao limpar histórico: %w", err)
	}
	return os.MkdirAll(historyDir, 0700)
}
