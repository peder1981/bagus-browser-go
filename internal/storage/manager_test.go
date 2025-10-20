package storage

import (
	"testing"
)

func TestNewManager(t *testing.T) {
	manager, err := NewManager("test-profile", "")
	if err != nil {
		t.Fatalf("NewManager() erro = %v", err)
	}

	if manager == nil {
		t.Fatal("NewManager() retornou nil")
	}

	if manager.GetProfile() != "test-profile" {
		t.Errorf("GetProfile() = %s, esperado test-profile", manager.GetProfile())
	}
}

func TestManager_GetDataDir(t *testing.T) {
	manager, _ := NewManager("test-profile", "")

	dataDir := manager.GetDataDir()
	if dataDir == "" {
		t.Error("GetDataDir() retornou string vazia")
	}
}

func TestManager_GetHistoryDir(t *testing.T) {
	manager, _ := NewManager("test-profile", "")

	historyDir := manager.GetHistoryDir()
	if historyDir == "" {
		t.Error("GetHistoryDir() retornou string vazia")
	}
}

func TestManager_GetBookmarksDir(t *testing.T) {
	manager, _ := NewManager("test-profile", "")

	bookmarksDir := manager.GetBookmarksDir()
	if bookmarksDir == "" {
		t.Error("GetBookmarksDir() retornou string vazia")
	}
}

func TestManager_GetCacheDir(t *testing.T) {
	manager, _ := NewManager("test-profile", "")

	cacheDir := manager.GetCacheDir()
	if cacheDir == "" {
		t.Error("GetCacheDir() retornou string vazia")
	}
}

func TestManager_GetDownloadsDir(t *testing.T) {
	manager, _ := NewManager("test-profile", "")

	downloadsDir := manager.GetDownloadsDir()
	if downloadsDir == "" {
		t.Error("GetDownloadsDir() retornou string vazia")
	}
}

func TestManager_GetProfile(t *testing.T) {
	manager, _ := NewManager("my-profile", "")

	profile := manager.GetProfile()
	if profile != "my-profile" {
		t.Errorf("GetProfile() = %s, esperado my-profile", profile)
	}
}

func TestManager_ClearCache(t *testing.T) {
	manager, _ := NewManager("test-profile", "")

	err := manager.ClearCache()
	if err != nil {
		t.Errorf("ClearCache() erro = %v", err)
	}
}

func TestManager_ClearHistory(t *testing.T) {
	manager, _ := NewManager("test-profile", "")

	err := manager.ClearHistory()
	if err != nil {
		t.Errorf("ClearHistory() erro = %v", err)
	}
}

func TestManager_Close(t *testing.T) {
	manager, _ := NewManager("test-profile", "")

	err := manager.Close()
	if err != nil {
		t.Errorf("Close() erro = %v", err)
	}
}
