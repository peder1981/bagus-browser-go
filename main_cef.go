//go:build cef
// +build cef

package main

import (
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/peder1981/bagus-browser-go/internal/browser"
	"github.com/peder1981/bagus-browser-go/internal/cef"
	"github.com/peder1981/bagus-browser-go/internal/config"
	"github.com/peder1981/bagus-browser-go/internal/ui"
)

func main() {
	log.Println("Iniciando Bagus Browser (CEF)...")

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

	// Cria histórico
	hist, err := browser.NewHistory(filepath.Join(userPath, "history"))
	if err != nil {
		log.Fatalf("Erro ao criar histórico: %v", err)
	}
	defer func() {
		if err := hist.Save(); err != nil {
			log.Printf("Erro ao salvar histórico: %v", err)
		} else {
			log.Println("Histórico salvo com sucesso")
		}
	}()

	// Cria browser CEF
	browser := cef.NewCEFBrowser()

	// Inicializa CEF
	if err := browser.Initialize(); err != nil {
		log.Fatalf("Erro ao inicializar CEF: %v", err)
	}
	defer browser.Shutdown()

	// Captura sinais para shutdown gracioso
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigChan
		log.Println("Sinal de interrupção recebido, finalizando...")
		browser.Shutdown()
		os.Exit(0)
	}()

	// Navega para URL inicial
	browser.Navigate(cfg.Default.URL)

	// Executa message loop (bloqueia aqui)
	browser.Run()

	log.Println("Bagus Browser finalizado")
}
