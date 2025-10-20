package utils

import (
	"runtime"
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
