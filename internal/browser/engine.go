package browser

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/peder1981/bagus-browser-go/internal/storage"
)

// Engine representa o motor principal do navegador
type Engine struct {
	storage *storage.Manager
	debug   bool
	tabs    []*Tab
}

// NewEngine cria uma nova instância do motor do navegador
func NewEngine(storageManager *storage.Manager, debug bool) (*Engine, error) {
	if storageManager == nil {
		return nil, fmt.Errorf("storage manager não pode ser nil")
	}

	engine := &Engine{
		storage: storageManager,
		debug:   debug,
		tabs:    make([]*Tab, 0),
	}

	if debug {
		log.Println("Browser engine inicializado em modo debug")
	}

	return engine, nil
}

// Run inicia o navegador
func (e *Engine) Run() error {
	log.Println("Browser engine executando...")

	// TODO: Implementar UI e WebView
	// Por enquanto, apenas um placeholder

	fmt.Println("🚀 Bagus Browser está rodando!")
	fmt.Println("📁 Dados armazenados em:", e.storage.GetDataDir())
	fmt.Println()
	fmt.Println("⚠️  Esta é uma versão alpha - UI será implementada em breve")
	fmt.Println()
	fmt.Println("Pressione Ctrl+C para sair")

	// Configurar captura de sinais para shutdown gracioso
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Aguardar sinal de interrupção
	<-sigChan

	fmt.Println()
	log.Println("Encerrando Bagus Browser...")

	// Cleanup
	if err := e.storage.Close(); err != nil {
		log.Printf("Erro ao fechar storage: %v", err)
	}

	log.Println("Browser encerrado com sucesso")
	return nil
}

// NewTab cria uma nova aba
func (e *Engine) NewTab(url string) (*Tab, error) {
	tab := &Tab{
		URL:    url,
		engine: e,
	}

	e.tabs = append(e.tabs, tab)

	if e.debug {
		log.Printf("Nova aba criada: %s", url)
	}

	return tab, nil
}

// GetTabs retorna todas as abas abertas
func (e *Engine) GetTabs() []*Tab {
	return e.tabs
}

// CloseTab fecha uma aba específica
func (e *Engine) CloseTab(index int) error {
	if index < 0 || index >= len(e.tabs) {
		return fmt.Errorf("índice de aba inválido: %d", index)
	}

	// Remove a aba
	e.tabs = append(e.tabs[:index], e.tabs[index+1:]...)

	if e.debug {
		log.Printf("Aba %d fechada", index)
	}

	return nil
}
