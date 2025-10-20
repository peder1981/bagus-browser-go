package browser

import (
	"github.com/peder1981/bagus-browser-go/internal/storage"
	"testing"
)

func setupTestTab(t *testing.T) (*Engine, *Tab) {
	manager, _ := storage.NewManager("default", "")
	engine, _ := NewEngine(manager, false)
	tab, _ := engine.NewTab("")
	return engine, tab
}

func TestTab_Creation(t *testing.T) {
	engine, _ := setupTestTab(t)

	tests := []struct {
		name    string
		url     string
		wantErr bool
	}{
		{
			name:    "URL válida",
			url:     "https://example.com",
			wantErr: false,
		},
		{
			name:    "URL vazia",
			url:     "",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tab, err := engine.NewTab(tt.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTab() erro = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && tab == nil {
				t.Error("NewTab() retornou nil sem erro")
			}
		})
	}
}

func TestTab_Navigate(t *testing.T) {
	_, tab := setupTestTab(t)

	tests := []struct {
		name    string
		url     string
		wantErr bool
	}{
		{
			name:    "URL válida",
			url:     "https://example.com",
			wantErr: false,
		},
		{
			name:    "URL vazia",
			url:     "",
			wantErr: true, // URL vazia não é permitida em Navigate
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tab.Navigate(tt.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("Navigate() erro = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTab_URL(t *testing.T) {
	engine, _ := setupTestTab(t)
	url := "https://example.com"
	tab, _ := engine.NewTab(url)

	if tab.URL != url {
		t.Errorf("URL = %s, esperado %s", tab.URL, url)
	}
}

func TestTab_Back(t *testing.T) {
	_, tab := setupTestTab(t)

	// Sem histórico
	err := tab.Back()
	if err == nil {
		t.Error("Back() deveria retornar erro sem histórico")
	}

	// Com histórico
	_ = tab.Navigate("https://example.com")
	_ = tab.Navigate("https://example.org")
	err = tab.Back()
	if err != nil {
		t.Errorf("Back() erro = %v", err)
	}
}

func TestTab_Reload(t *testing.T) {
	_, tab := setupTestTab(t)
	_ = tab.Navigate("https://example.com")

	err := tab.Reload()
	if err != nil {
		t.Errorf("Reload() erro = %v", err)
	}
}

func TestTab_GetHistory(t *testing.T) {
	_, tab := setupTestTab(t)

	// Histórico inicial vazio
	history := tab.GetHistory()
	if len(history) != 0 {
		t.Errorf("GetHistory() = %d itens, esperado 0", len(history))
	}

	// Navegar cria histórico
	// Primeira navegação adiciona URL inicial ao histórico
	_ = tab.Navigate("https://example.com")
	// Segunda navegação adiciona a primeira URL ao histórico
	_ = tab.Navigate("https://example.org")
	history = tab.GetHistory()
	// Esperamos 2 itens: URL inicial vazia + primeira navegação
	if len(history) != 2 {
		t.Errorf("GetHistory() = %d itens, esperado 2", len(history))
	}
}
