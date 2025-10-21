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

	// Carrega interface HTML
	html := b.generateHTML()
	w.SetHtml(html)

	w.Run()
	return nil
}

// bindFunctions registra funções Go para JavaScript
func (b *BrowserSingleWindow) bindFunctions() {
	// Navegar
	b.w.Bind("navigateGo", func(url string) {
		b.navigate(url)
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

// navigate processa e navega para URL
func (b *BrowserSingleWindow) navigate(input string) {
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

	// Navega diretamente
	log.Printf("Navegando para: %s", finalURL)
	if b.w != nil {
		b.w.Navigate(finalURL)
	}
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

// generateHTML gera HTML com barra de navegação injetada
func (b *BrowserSingleWindow) generateHTML() string {
	return fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>Bagus Browser</title>
    <script>
        // Injeta barra de navegação assim que a página carregar
        window.addEventListener('DOMContentLoaded', function() {
            injectNavBar();
            // Navega para URL inicial
            setTimeout(function() {
                window.location.href = '%s';
            }, 100);
        });

        // Reinjeta barra após navegação
        window.addEventListener('load', function() {
            setTimeout(injectNavBar, 500);
        });

        function injectNavBar() {
            // Remove barra existente
            var existing = document.getElementById('bagus-nav-bar');
            if (existing) existing.remove();

            // Cria barra
            var nav = document.createElement('div');
            nav.id = 'bagus-nav-bar';
            nav.style.cssText = 'position: fixed; top: 0; left: 0; right: 0; z-index: 999999; background: #2d2d2d; padding: 10px; display: flex; gap: 10px; box-shadow: 0 2px 5px rgba(0,0,0,0.3);';

            // Botão Voltar
            var btnBack = createButton('◀', 'Voltar', function() { history.back(); });
            nav.appendChild(btnBack);

            // Botão Avançar
            var btnForward = createButton('▶', 'Avançar', function() { history.forward(); });
            nav.appendChild(btnForward);

            // Botão Recarregar
            var btnReload = createButton('⟳', 'Recarregar', function() { location.reload(); });
            nav.appendChild(btnReload);

            // Campo URL
            var urlInput = document.createElement('input');
            urlInput.type = 'text';
            urlInput.id = 'bagus-url-input';
            urlInput.value = window.location.href;
            urlInput.placeholder = 'Digite uma URL ou termo de busca...';
            urlInput.style.cssText = 'flex: 1; padding: 10px 15px; background: #3c3c3c; border: 1px solid #555; border-radius: 4px; color: #fff; font-size: 15px;';
            urlInput.onkeydown = function(e) {
                if (e.key === 'Enter') {
                    var input = this.value.trim();
                    if (input && window.processInputGo && window.navigateGo) {
                        navigateGo(input);
                    }
                }
            };
            nav.appendChild(urlInput);

            // Insere no body
            if (document.body) {
                document.body.insertBefore(nav, document.body.firstChild);
                // Ajusta padding
                document.body.style.paddingTop = (nav.offsetHeight) + 'px';
            }
        }

        function createButton(text, title, onclick) {
            var btn = document.createElement('button');
            btn.textContent = text;
            btn.title = title;
            btn.style.cssText = 'padding: 10px 20px; background: #0078d4; color: white; border: none; border-radius: 4px; cursor: pointer; font-size: 16px;';
            btn.onmouseover = function() { this.style.background = '#106ebe'; };
            btn.onmouseout = function() { this.style.background = '#0078d4'; };
            btn.onclick = onclick;
            return btn;
        }
    </script>
</head>
<body>
    <div style="display: flex; align-items: center; justify-content: center; height: 100vh; font-family: Arial;">
        <p>Carregando Bagus Browser...</p>
    </div>
</body>
</html>`, b.config.Default.URL)
}
