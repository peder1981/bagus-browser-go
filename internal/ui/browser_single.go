package ui

import (
	"encoding/json"
	"log"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/peder1981/bagus-browser-go/internal/browser"
	"github.com/peder1981/bagus-browser-go/internal/config"
	"github.com/peder1981/bagus-browser-go/internal/security"
	webview "github.com/webview/webview_go"
)

// BrowserSingleWindow representa o navegador com janela única
type BrowserSingleWindow struct {
	w       webview.WebView
	config  *config.Config
	history *browser.History
	blocker *security.Blocker
}

// NewBrowserSingleWindow cria nova instância do browser
func NewBrowserSingleWindow(userPath string, cfg *config.Config) (*BrowserSingleWindow, error) {
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

	b := &BrowserSingleWindow{
		config:  cfg,
		history: hist,
		blocker: blocker,
	}

	return b, nil
}

// Run inicia o navegador
func (b *BrowserSingleWindow) Run() error {
	// Salva histórico ao fechar
	defer func() {
		if err := b.history.Save(); err != nil {
			log.Printf("Erro ao salvar histórico: %v", err)
		} else {
			log.Println("Histórico salvo com sucesso")
		}
	}()

	log.Println("Iniciando interface do browser...")

	debug := false
	w := webview.New(debug)
	defer w.Destroy()

	b.w = w

	w.SetTitle("Bagus Browser")
	w.SetSize(1200, 800, webview.HintNone)

	// Bind funções
	b.bindFunctions()

	// Injeta controles que serão executados em TODAS as páginas
	w.Init(b.getControlsScript())

	// Navega diretamente para URL inicial
	initialURL := b.config.Default.URL
	if initialURL == "" {
		initialURL = "https://duckduckgo.com"
	}
	
	log.Printf("Navegando para URL inicial: %s", initialURL)
	w.Navigate(initialURL)

	w.Run()
	return nil
}

// bindFunctions registra funções Go para JavaScript
func (b *BrowserSingleWindow) bindFunctions() {
	// Navegar (retorna URL processada)
	b.w.Bind("navigateGo", func(url string) string {
		return b.navigate(url)
	})

	// Buscar histórico
	b.w.Bind("searchHistoryGo", func(query string) string {
		results := b.history.Search(query)
		data, _ := json.Marshal(results)
		return string(data)
	})

	// Processar input
	b.w.Bind("processInputGo", func(input string) string {
		return b.processInput(input)
	})
}

// getControlsScript retorna o script de controles para ser injetado em todas as páginas
func (b *BrowserSingleWindow) getControlsScript() string {
	return `
		(function() {
			// Previne múltiplas injeções
			if (window.bagusControlsInjected) return;
			window.bagusControlsInjected = true;

			// Atalhos de teclado
			document.addEventListener('keydown', function(e) {
				// Ctrl+L - Abrir diálogo de navegação
				if (e.ctrlKey && e.key === 'l') {
					e.preventDefault();
					var url = prompt('Digite uma URL ou termo de busca:', window.location.href);
					if (url && window.navigateGo) {
						navigateGo(url);
					}
				}
				
				// Ctrl+R ou F5 - Recarregar
				if ((e.ctrlKey && e.key === 'r') || e.key === 'F5') {
					e.preventDefault();
					window.location.reload();
				}
				
				// Alt+← - Voltar
				if (e.altKey && e.key === 'ArrowLeft') {
					e.preventDefault();
					window.history.back();
				}
				
				// Alt+→ - Avançar
				if (e.altKey && e.key === 'ArrowRight') {
					e.preventDefault();
					window.history.forward();
				}
				
				// Ctrl+H - Mostrar histórico
				if (e.ctrlKey && e.key === 'h') {
					e.preventDefault();
					if (window.searchHistoryGo) {
						var results = searchHistoryGo('');
						console.log('Histórico:', results);
						alert('Histórico disponível no console (F12)');
					}
				}
			});

			// Mensagem de boas-vindas
			console.log('%c🌐 Bagus Browser v2.0.0', 'font-size: 20px; font-weight: bold; color: #0078d4;');
			console.log('%cAtalhos de Teclado:', 'font-size: 14px; font-weight: bold; margin-top: 10px;');
			console.log('  Ctrl+L     - Navegar para URL');
			console.log('  Ctrl+R/F5  - Recarregar página');
			console.log('  Alt+←      - Voltar');
			console.log('  Alt+→      - Avançar');
			console.log('  Ctrl+H     - Ver histórico');
			console.log('%cNavegação via código:', 'font-size: 14px; font-weight: bold; margin-top: 10px;');
			console.log('  navigateGo("https://exemplo.com")');
		})();
	`
}

// navigate processa e navega para URL
func (b *BrowserSingleWindow) navigate(input string) string {
	// Processa input
	finalURL := b.processInput(input)
	if finalURL == "" {
		return ""
	}

	// Valida URL
	finalURL = security.SanitizeURL(finalURL)
	if err := security.ValidateURL(finalURL); err != nil {
		log.Printf("URL inválida: %v", err)
		return ""
	}

	// Verifica bloqueio
	if b.blocker != nil {
		domain, err := security.ExtractDomain(finalURL)
		if err == nil && b.blocker.IsBlocked(domain) {
			log.Printf("Domínio bloqueado: %s", domain)
			return ""
		}
	}

	// Adiciona ao histórico
	b.history.Add(finalURL, finalURL)

	// Navega diretamente
	log.Printf("Navegando para: %s", finalURL)
	if b.w != nil {
		b.w.Navigate(finalURL)
	}
	return finalURL
}

// processInput processa input da barra de navegação
func (b *BrowserSingleWindow) processInput(input string) string {
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
