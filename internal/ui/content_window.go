package ui

import (
	"log"
	"sync"

	webview "github.com/webview/webview_go"
)

// ContentWindow gerencia a janela de conteúdo (navegação web)
type ContentWindow struct {
	w          webview.WebView
	currentURL string
	mu         sync.RWMutex
}

// NewContentWindow cria nova janela de conteúdo
func NewContentWindow() *ContentWindow {
	return &ContentWindow{
		currentURL: "",
	}
}

// Run inicia a janela de conteúdo
func (c *ContentWindow) Run(initialURL string, onURLChange func(string)) {
	debug := false
	w := webview.New(debug)
	defer w.Destroy()

	c.w = w
	c.currentURL = initialURL

	w.SetTitle("Bagus Browser")
	w.SetSize(1200, 720, webview.HintNone)

	// Navega para URL inicial
	log.Printf("Navegando para: %s", initialURL)
	w.Navigate(initialURL)

	w.Run()
}

// Navigate navega para uma nova URL
func (c *ContentWindow) Navigate(url string) {
	c.mu.Lock()
	c.currentURL = url
	c.mu.Unlock()

	if c.w != nil {
		log.Printf("Navegando para: %s", url)
		c.w.Navigate(url)
	}
}

// GetCurrentURL retorna a URL atual
func (c *ContentWindow) GetCurrentURL() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.currentURL
}
