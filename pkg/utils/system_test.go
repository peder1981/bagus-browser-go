package utils

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestGetHomeDir(t *testing.T) {
	home, err := GetHomeDir()

	if err != nil {
		t.Fatalf("GetHomeDir() erro = %v", err)
	}

	if home == "" {
		t.Error("GetHomeDir() retornou string vazia")
	}

	// Verificar que o diretório existe
	if _, err := os.Stat(home); os.IsNotExist(err) {
		t.Errorf("Diretório home não existe: %s", home)
	}
}

func TestGetConfigDir(t *testing.T) {
	configDir, err := GetConfigDir()

	if err != nil {
		t.Fatalf("GetConfigDir() erro = %v", err)
	}

	if configDir == "" {
		t.Error("GetConfigDir() retornou string vazia")
	}

	// Verificar que contém "bagus" no caminho
	if !contains(configDir, "bagus") && !contains(configDir, ".bagus") {
		t.Errorf("GetConfigDir() não contém 'bagus': %s", configDir)
	}
}

func TestGetCacheDir(t *testing.T) {
	cacheDir, err := GetCacheDir()

	if err != nil {
		t.Fatalf("GetCacheDir() erro = %v", err)
	}

	if cacheDir == "" {
		t.Error("GetCacheDir() retornou string vazia")
	}
}

func TestGetDataDir(t *testing.T) {
	dataDir, err := GetDataDir()

	if err != nil {
		t.Fatalf("GetDataDir() erro = %v", err)
	}

	if dataDir == "" {
		t.Error("GetDataDir() retornou string vazia")
	}
}

func TestEnsureDir(t *testing.T) {
	tmpDir := t.TempDir()
	testPath := filepath.Join(tmpDir, "test", "nested", "dir")

	err := EnsureDir(testPath)
	if err != nil {
		t.Fatalf("EnsureDir() erro = %v", err)
	}

	// Verificar que o diretório foi criado
	if _, err := os.Stat(testPath); os.IsNotExist(err) {
		t.Error("EnsureDir() não criou o diretório")
	}

	// Chamar novamente deve funcionar (idempotente)
	err = EnsureDir(testPath)
	if err != nil {
		t.Errorf("EnsureDir() em diretório existente erro = %v", err)
	}
}

func TestFileExists(t *testing.T) {
	tmpDir := t.TempDir()

	// Arquivo que existe
	existingFile := filepath.Join(tmpDir, "exists.txt")
	_ = os.WriteFile(existingFile, []byte("test"), 0644)

	if !FileExists(existingFile) {
		t.Error("FileExists() = false para arquivo existente")
	}

	// Arquivo que não existe
	nonExistingFile := filepath.Join(tmpDir, "nao-existe.txt")
	if FileExists(nonExistingFile) {
		t.Error("FileExists() = true para arquivo inexistente")
	}
}

func TestDirExists(t *testing.T) {
	tmpDir := t.TempDir()

	// Diretório que existe
	if !DirExists(tmpDir) {
		t.Error("DirExists() = false para diretório existente")
	}

	// Diretório que não existe
	nonExistingDir := filepath.Join(tmpDir, "nao-existe")
	if DirExists(nonExistingDir) {
		t.Error("DirExists() = true para diretório inexistente")
	}
}

func TestGetOS(t *testing.T) {
	os := GetOS()

	if os == "" {
		t.Error("GetOS() retornou string vazia")
	}

	validOS := []string{"linux", "windows", "darwin"}
	if !containsString(validOS, os) {
		t.Errorf("GetOS() = %s, esperado um de %v", os, validOS)
	}
}

func TestGetArch(t *testing.T) {
	arch := GetArch()

	if arch == "" {
		t.Error("GetArch() retornou string vazia")
	}

	validArch := []string{"amd64", "arm64", "386", "arm"}
	if !containsString(validArch, arch) {
		t.Errorf("GetArch() = %s, esperado um de %v", arch, validArch)
	}
}

func TestIsLinux(t *testing.T) {
	result := IsLinux()
	expected := runtime.GOOS == "linux"

	if result != expected {
		t.Errorf("IsLinux() = %v, esperado %v", result, expected)
	}
}

func TestIsWindows(t *testing.T) {
	result := IsWindows()
	expected := runtime.GOOS == "windows"

	if result != expected {
		t.Errorf("IsWindows() = %v, esperado %v", result, expected)
	}
}

func TestIsMacOS(t *testing.T) {
	result := IsMacOS()
	expected := runtime.GOOS == "darwin"

	if result != expected {
		t.Errorf("IsMacOS() = %v, esperado %v", result, expected)
	}
}

func TestExpandPath(t *testing.T) {
	tests := []struct {
		name string
		path string
	}{
		{
			name: "Path com ~",
			path: "~/test",
		},
		{
			name: "Path absoluto",
			path: "/tmp/test",
		},
		{
			name: "Path relativo",
			path: "test/path",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ExpandPath(tt.path)
			if result == "" {
				t.Error("ExpandPath() retornou string vazia")
			}
		})
	}
}

func TestJoinPath(t *testing.T) {
	result := JoinPath("a", "b", "c")
	expected := filepath.Join("a", "b", "c")

	if result != expected {
		t.Errorf("JoinPath() = %s, esperado %s", result, expected)
	}
}

// Helper functions
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) &&
		(s[:len(substr)] == substr || s[len(s)-len(substr):] == substr ||
			len(s) > len(substr)*2 && s[len(s)/2-len(substr)/2:len(s)/2+len(substr)/2+len(substr)%2] == substr))
}

func containsString(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}
