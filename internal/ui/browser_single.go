package ui

import (
	"encoding/json"
	"fmt"
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

	// Carrega interface HTML com iframe
	html := b.generateHTML()
	w.SetHtml(html)

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

// navigate processa e navega para URL (retorna URL processada)
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

	// Retorna URL para navegação via iframe
	log.Printf("Navegando para: %s", finalURL)
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

// generateHTML gera HTML com barra de navegação e iframe
func (b *BrowserSingleWindow) generateHTML() string {
	defaultURL := b.config.Default.URL
	if defaultURL == "" {
		defaultURL = "https://duckduckgo.com"
	}
	
	return fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>Bagus Browser</title>
    <style>
        * { margin: 0; padding: 0; box-sizing: border-box; }
        body { font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif; overflow: hidden; }
        #nav-bar {
            position: fixed;
            top: 0;
            left: 0;
            right: 0;
            height: 50px;
            background: #2d2d2d;
            display: flex;
            align-items: center;
            gap: 10px;
            padding: 0 10px;
            box-shadow: 0 2px 5px rgba(0,0,0,0.3);
            z-index: 1000;
        }
        #content-frame {
            position: fixed;
            top: 50px;
            left: 0;
            right: 0;
            bottom: 0;
            border: none;
            width: 100%%;
            height: calc(100%% - 50px);
        }
        .nav-btn {
            padding: 8px 16px;
            background: #0078d4;
            color: white;
            border: none;
            border-radius: 4px;
            cursor: pointer;
            font-size: 16px;
            min-width: 40px;
        }
        .nav-btn:hover { background: #106ebe; }
        #url-input {
            flex: 1;
            padding: 10px 15px;
            background: #3c3c3c;
            border: 1px solid #555;
            border-radius: 4px;
            color: #fff;
            font-size: 15px;
        }
        #url-input:focus {
            outline: none;
            border-color: #0078d4;
        }
    </style>
    <script>
        function navigate(url) {
            var frame = document.getElementById('content-frame');
            var urlInput = document.getElementById('url-input');
            if (frame && url) {
                frame.src = url;
                if (urlInput) urlInput.value = url;
            }
        }

        async function handleUrlInput(event) {
            if (event.key === 'Enter') {
                var input = event.target.value.trim();
                if (input && window.navigateGo) {
                    try {
                        var url = await navigateGo(input);
                        if (url) navigate(url);
                    } catch (e) {
                        console.error('Erro ao navegar:', e);
                    }
                }
            }
        }

        window.addEventListener('DOMContentLoaded', function() {
            // Navega para URL inicial
            navigate('%s');
        });
    </script>
</head>
<body>
    <div id="nav-bar">
        <button class="nav-btn" onclick="document.getElementById('content-frame').contentWindow.history.back()" title="Voltar">◀</button>
        <button class="nav-btn" onclick="document.getElementById('content-frame').contentWindow.history.forward()" title="Avançar">▶</button>
        <button class="nav-btn" onclick="document.getElementById('content-frame').contentWindow.location.reload()" title="Recarregar">⟳</button>
        <input type="text" id="url-input" placeholder="Digite uma URL ou termo de busca..." onkeydown="handleUrlInput(event)" value="%s">
    </div>
    <iframe id="content-frame" src="%s"></iframe>
</body>
</html>`, defaultURL, defaultURL, defaultURL)
}
