package utils

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// GetGoVersion retorna a versão do Go
func GetGoVersion() string {
	return runtime.Version()
}

// GetOS retorna o sistema operacional
func GetOS() string {
	return runtime.GOOS
}

// GetArch retorna a arquitetura
func GetArch() string {
	return runtime.GOARCH
}

// IsLinux verifica se está rodando em Linux
func IsLinux() bool {
	return runtime.GOOS == "linux"
}

// IsWindows verifica se está rodando em Windows
func IsWindows() bool {
	return runtime.GOOS == "windows"
}

// IsMacOS verifica se está rodando em macOS
func IsMacOS() bool {
	return runtime.GOOS == "darwin"
}

// GetHomeDir retorna o diretório home do usuário
func GetHomeDir() (string, error) {
	return os.UserHomeDir()
}

// GetConfigDir retorna o diretório de configuração do Bagus Browser
func GetConfigDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	var configDir string
	switch runtime.GOOS {
	case "windows":
		configDir = filepath.Join(os.Getenv("APPDATA"), "bagus")
	case "darwin":
		configDir = filepath.Join(home, "Library", "Application Support", "bagus")
	default: // linux e outros
		configDir = filepath.Join(home, ".bagus")
	}

	return configDir, nil
}

// GetCacheDir retorna o diretório de cache do Bagus Browser
func GetCacheDir() (string, error) {
	configDir, err := GetConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(configDir, "cache"), nil
}

// GetDataDir retorna o diretório de dados do Bagus Browser
func GetDataDir() (string, error) {
	configDir, err := GetConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(configDir, "data"), nil
}

// EnsureDir garante que um diretório existe, criando-o se necessário
func EnsureDir(path string) error {
	return os.MkdirAll(path, 0700)
}

// FileExists verifica se um arquivo existe
func FileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// DirExists verifica se um diretório existe
func DirExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

// ExpandPath expande ~ para o diretório home
func ExpandPath(path string) string {
	if strings.HasPrefix(path, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			return path
		}
		return filepath.Join(home, path[1:])
	}
	return path
}

// JoinPath junta elementos de caminho
func JoinPath(elem ...string) string {
	return filepath.Join(elem...)
}
