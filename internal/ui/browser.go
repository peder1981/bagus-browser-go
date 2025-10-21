package ui

import (
	"log"
	"net/url"
	"path/filepath"
	"strings"
	"sync"

	"github.com/peder1981/bagus-browser-go/internal/browser"
	"github.com/peder1981/bagus-browser-go/internal/config"
	"github.com/peder1981/bagus-browser-go/internal/security"
)

// BrowserDualWindow representa o navegador com duas janelas independentes
type BrowserDualWindow struct {
	config        *config.Config
	history       *browser.History
	blocker       *security.Blocker
	controlWindow *ControlWindow
	contentWindow *ContentWindow
	currentURL    string
	mu            sync.RWMutex
	wg            sync.WaitGroup
}

// NewBrowserDualWindow cria nova instância do browser com duas janelas
func NewBrowserDualWindow(userPath string, cfg *config.Config) (*BrowserDualWindow, error) {
	// Cria histórico
	hist, err := browser.NewHistory(filepath.Join(userPath, "history"))
	if err != nil {
		return nil, err
	}

	// Cria blocker
	blocker, err := security.NewBlocker(filepath.Join(userPath, "ad_hosts_block.txt"), log.Default())
	if err != nil {
		log.Printf("Lista de bloqueio não encontrada: %s", filepath.Join(userPath, "ad_hosts_block.txt"))
	}

	b := &BrowserDualWindow{
		config:        cfg,
		history:       hist,
		blocker:       blocker,
		controlWindow: NewControlWindow(),
		contentWindow: NewContentWindow(),
		currentURL:    cfg.Default.URL,
	}

	return b, nil
}

// Run inicia o navegador com duas janelas
func (b *BrowserDualWindow) Run() error {
	// Salva histórico ao fechar
	defer func() {
		if err := b.history.Save(); err != nil {
			log.Printf("Erro ao salvar histórico: %v", err)
		} else {
			log.Println("Histórico salvo com sucesso")
		}
	}()

	log.Println("Iniciando interface do browser...")

	// Inicia janela de conteúdo em goroutine
	b.wg.Add(1)
	go func() {
		defer b.wg.Done()
		b.contentWindow.Run(b.currentURL, func(url string) {
			b.mu.Lock()
			b.currentURL = url
			b.mu.Unlock()
			
			if b.controlWindow != nil {
				b.controlWindow.UpdateURL(url)
			}
		})
	}()

	// Inicia janela de controle (bloqueia até fechar)
	b.controlWindow.Run(
		// onNavigate
		func(url string) {
			b.navigate(url)
		},
		// onSearch
		func(query string) []interface{} {
			results := b.history.Search(query)
			// Converte para []interface{}
			interfaceResults := make([]interface{}, len(results))
			for i, r := range results {
				interfaceResults[i] = r
			}
			return interfaceResults
		},
	)

	// Aguarda janela de conteúdo fechar
	b.wg.Wait()

	return nil
}

// navigate processa e navega para URL
func (b *BrowserDualWindow) navigate(input string) {
	// Processa input
	finalURL := b.processInput(input)
	if finalURL == "" {
		return
	}

	// Valida URL
	finalURL = security.SanitizeURL(finalURL)
	if err := security.ValidateURL(finalURL); err != nil {
		log.Printf("URL inválida: %v", err)
		return
	}

	// Verifica bloqueio
	if b.blocker != nil {
		domain, err := security.ExtractDomain(finalURL)
		if err == nil && b.blocker.IsBlocked(domain) {
			log.Printf("Domínio bloqueado: %s", domain)
			return
		}
	}

	// Adiciona ao histórico
	b.history.Add(finalURL, finalURL)

	// Atualiza URL atual
	b.mu.Lock()
	b.currentURL = finalURL
	b.mu.Unlock()

	// Navega na janela de conteúdo
	if b.contentWindow != nil {
		b.contentWindow.Navigate(finalURL)
	}

	// Atualiza barra de controle
	if b.controlWindow != nil {
		b.controlWindow.UpdateURL(finalURL)
	}
}

// processInput processa input da barra de navegação
func (b *BrowserDualWindow) processInput(input string) string {
	input = strings.TrimSpace(input)
	if input == "" {
		return ""
	}

	// Verifica se parece com uma URL
	isURL := strings.Contains(input, ".") &&
		(strings.HasPrefix(input, "http://") ||
			strings.HasPrefix(input, "https://") ||
			!strings.Contains(input, " "))

	if isURL {
		// Adiciona protocolo se necessário
		if !strings.HasPrefix(input, "http://") && !strings.HasPrefix(input, "https://") {
			input = "https://" + input
		}
		// Valida URL
		if _, err := url.Parse(input); err == nil {
			return input
		}
	}

	// Se não for URL, faz busca no DuckDuckGo
	searchURL := "https://duckduckgo.com/?q=" + url.QueryEscape(input)
	return searchURL
}
