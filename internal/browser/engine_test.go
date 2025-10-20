package browser

import (
	"github.com/peder1981/bagus-browser-go/internal/storage"
	"testing"
)

func setupTestEngine(t *testing.T) *Engine {
	manager, err := storage.NewManager("default", "")
	if err != nil {
		t.Fatalf("Erro ao criar storage manager: %v", err)
	}
	engine, err := NewEngine(manager, false)
	if err != nil {
		t.Fatalf("Erro ao criar engine: %v", err)
	}
	return engine
}

func TestNewEngine(t *testing.T) {
	manager, _ := storage.NewManager("default", "")
	engine, err := NewEngine(manager, false)

	if err != nil {
		t.Fatalf("NewEngine() erro = %v, esperado nil", err)
	}

	if engine == nil {
		t.Fatal("NewEngine() retornou nil")
	}

	if engine.tabs == nil {
		t.Error("engine.tabs não foi inicializado")
	}
}

func TestEngine_NewTab(t *testing.T) {
	engine := setupTestEngine(t)

	tests := []struct {
		name    string
		url     string
		wantErr bool
	}{
		{
			name:    "URL válida HTTPS",
			url:     "https://example.com",
			wantErr: false,
		},
		{
			name:    "URL válida HTTP",
			url:     "http://example.com",
			wantErr: false,
		},
		{
			name:    "URL vazia",
			url:     "",
			wantErr: false, // Aba em branco é permitida
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

func TestEngine_CloseTab(t *testing.T) {
	engine := setupTestEngine(t)

	// Abrir uma aba primeiro
	_, _ = engine.NewTab("https://example.com")

	if len(engine.tabs) == 0 {
		t.Fatal("Nenhuma aba foi aberta")
	}

	err := engine.CloseTab(0)
	if err != nil {
		t.Errorf("CloseTab() erro = %v", err)
	}

	if len(engine.tabs) != 0 {
		t.Errorf("Aba não foi fechada, tabs = %d", len(engine.tabs))
	}
}

func TestEngine_GetTabs(t *testing.T) {
	engine := setupTestEngine(t)

	// Sem abas
	tabs := engine.GetTabs()
	if len(tabs) != 0 {
		t.Errorf("GetTabs() = %d abas, esperado 0", len(tabs))
	}

	// Com aba
	_, _ = engine.NewTab("https://example.com")
	tabs = engine.GetTabs()
	if len(tabs) != 1 {
		t.Errorf("GetTabs() = %d abas, esperado 1", len(tabs))
	}
}

func TestEngine_MultipleTabsManagement(t *testing.T) {
	engine := setupTestEngine(t)

	if len(engine.GetTabs()) != 0 {
		t.Errorf("Tabs iniciais = %d, esperado 0", len(engine.GetTabs()))
	}

	_, _ = engine.NewTab("https://example.com")
	if len(engine.GetTabs()) != 1 {
		t.Errorf("Tabs após primeira aba = %d, esperado 1", len(engine.GetTabs()))
	}

	_, _ = engine.NewTab("https://example.org")
	if len(engine.GetTabs()) != 2 {
		t.Errorf("Tabs após segunda aba = %d, esperado 2", len(engine.GetTabs()))
	}
}
