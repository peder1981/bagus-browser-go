package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/peder1981/bagus-browser-go/internal/config"
	"github.com/peder1981/bagus-browser-go/internal/ui"
)

func main() {
	log.Println("Iniciando Bagus Browser...")

	// Mostra tela de login
	loginDialog := ui.NewLoginDialog()
	userPath, err := loginDialog.Show()
	if err != nil {
		log.Fatalf("Erro no login: %v", err)
	}

	if userPath == "" {
		log.Println("Login cancelado")
		os.Exit(0)
	}

	log.Printf("Usuário autenticado: %s", userPath)

	// Carrega configuração
	configPath := filepath.Join(userPath, "config.json")
	cfg, err := config.Load(configPath)
	if err != nil {
		log.Printf("Erro ao carregar config, usando padrão: %v", err)
		cfg = config.Default()
	}

	// Cria e inicia browser
	browser, err := ui.NewBrowser(userPath, cfg)
	if err != nil {
		log.Fatalf("Erro ao criar browser: %v", err)
	}

	log.Println("Iniciando interface do browser...")
	if err := browser.Run(); err != nil {
		log.Fatalf("Erro ao executar browser: %v", err)
	}
}
