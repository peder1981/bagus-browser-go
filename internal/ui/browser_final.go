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

// BrowserFinal representa o navegador com interface nativa
type BrowserFinal struct {
	w       webview.WebView
	config  *config.Config
	history *browser.History
	blocker *security.Blocker
}

// NewBrowserFinal cria nova instância do browser
func NewBrowserFinal(userPath string, cfg *config.Config) (*BrowserFinal, error) {
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

	b := &BrowserFinal{
		config:  cfg,
		history: hist,
		blocker: blocker,
	}

	return b, nil
}

// Run inicia o navegador
func (b *BrowserFinal) Run() error {
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

	// Carrega interface HTML com controles integrados
	html := b.generateHTML()
	w.SetHtml(html)

	w.Run()
	return nil
}

// bindFunctions registra funções Go para JavaScript
func (b *BrowserFinal) bindFunctions() {
	// Navegar
	b.w.Bind("navigateGo", func(input string) {
		b.navigate(input)
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

	// Voltar
	b.w.Bind("goBackGo", func() {
		js := "if (window.browserHistory && window.browserHistory.length > 1) { window.browserHistory.pop(); var url = window.browserHistory[window.browserHistory.length - 1]; if (url) loadInIframe(url); }"
		b.w.Eval(js)
	})

	// Avançar (placeholder)
	b.w.Bind("goForwardGo", func() {
		log.Println("Avançar não implementado")
	})

	// Recarregar
	b.w.Bind("reloadGo", func() {
		js := "var iframe = document.getElementById('content-frame'); if (iframe && iframe.src) iframe.src = iframe.src;"
		b.w.Eval(js)
	})
}

// navigate processa e navega para URL
func (b *BrowserFinal) navigate(input string) {
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

	// Navega no iframe
	log.Printf("Navegando para: %s", finalURL)
	if b.w != nil {
		escapedURL := strings.ReplaceAll(finalURL, "'", "\\'")
		js := fmt.Sprintf("loadInIframe('%s');", escapedURL)
		b.w.Eval(js)
	}
}

// processInput processa input da barra de navegação
func (b *BrowserFinal) processInput(input string) string {
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

// generateHTML gera HTML da interface completa
func (b *BrowserFinal) generateHTML() string {
	defaultURL := b.config.Default.URL

	return fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Bagus Browser</title>
    <style>
        * { margin: 0; padding: 0; box-sizing: border-box; }
        html, body { 
            height: 100%%; 
            overflow: hidden;
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
        }
        
        .browser-container {
            display: flex;
            flex-direction: column;
            height: 100vh;
        }
        
        .toolbar {
            background: #2d2d2d;
            padding: 10px;
            display: flex;
            gap: 10px;
            align-items: center;
            border-bottom: 1px solid #444;
            flex-shrink: 0;
        }
        
        .toolbar button {
            padding: 10px 20px;
            background: #0078d4;
            color: white;
            border: none;
            border-radius: 4px;
            cursor: pointer;
            font-size: 16px;
            font-weight: 500;
            transition: background 0.2s;
        }
        
        .toolbar button:hover { background: #106ebe; }
        .toolbar button:active { background: #005a9e; }
        
        .toolbar input {
            flex: 1;
            padding: 10px 15px;
            background: #3c3c3c;
            border: 1px solid #555;
            border-radius: 4px;
            color: #fff;
            font-size: 15px;
        }
        
        .toolbar input:focus {
            outline: none;
            border-color: #0078d4;
            background: #454545;
        }
        
        .content-area {
            flex: 1;
            position: relative;
            overflow: hidden;
        }
        
        #content-frame {
            width: 100%%;
            height: 100%%;
            border: none;
            background: white;
        }
        
        .loading {
            position: absolute;
            top: 50%%;
            left: 50%%;
            transform: translate(-50%%, -50%%);
            color: #666;
            font-size: 18px;
        }
        
        .suggestions {
            position: absolute;
            top: 60px;
            left: 10px;
            right: 10px;
            background: #2d2d2d;
            border: 1px solid #555;
            border-radius: 4px;
            max-height: 300px;
            overflow-y: auto;
            display: none;
            z-index: 1000;
            box-shadow: 0 4px 8px rgba(0,0,0,0.3);
        }
        
        .suggestion {
            padding: 12px;
            cursor: pointer;
            border-bottom: 1px solid #444;
            color: #fff;
        }
        
        .suggestion:hover { background: #3c3c3c; }
        .suggestion:last-child { border-bottom: none; }
    </style>
</head>
<body>
    <div class="browser-container">
        <div class="toolbar">
            <button onclick="goBack()" title="Voltar">◀</button>
            <button onclick="goForward()" title="Avançar">▶</button>
            <button onclick="reload()" title="Recarregar">⟳</button>
            <input type="text" 
                   id="url-bar" 
                   placeholder="Digite uma URL ou termo de busca..." 
                   onkeydown="handleKeyDown(event)"
                   onkeyup="handleKeyUp(event)"
                   onfocus="hideSuggestions()"
                   value="%s">
            <button onclick="openSettings()" title="Configurações">⚙</button>
        </div>
        <div class="suggestions" id="suggestions"></div>
        <div class="content-area">
            <div class="loading" id="loading">Carregando...</div>
            <iframe id="content-frame" sandbox="allow-same-origin allow-scripts allow-forms allow-popups allow-popups-to-escape-sandbox"></iframe>
        </div>
    </div>

    <script>
        // Histórico de navegação
        window.browserHistory = [];
        
        // Carrega URL inicial
        window.addEventListener('DOMContentLoaded', function() {
            var initialUrl = document.getElementById('url-bar').value;
            if (initialUrl) {
                loadInIframe(initialUrl);
            }
        });
        
        function loadInIframe(url) {
            var iframe = document.getElementById('content-frame');
            var loading = document.getElementById('loading');
            var urlBar = document.getElementById('url-bar');
            
            loading.style.display = 'block';
            iframe.style.display = 'none';
            
            iframe.src = url;
            urlBar.value = url;
            
            // Adiciona ao histórico
            if (window.browserHistory.length === 0 || window.browserHistory[window.browserHistory.length - 1] !== url) {
                window.browserHistory.push(url);
            }
            
            iframe.onload = function() {
                loading.style.display = 'none';
                iframe.style.display = 'block';
            };
        }
        
        function handleKeyDown(event) {
            if (event.key === 'Enter') {
                navigate();
            }
        }
        
        function handleKeyUp(event) {
            if (event.key === 'Enter' || event.key === 'Escape') {
                return;
            }
            
            var query = event.target.value.trim();
            if (query.length < 2) {
                hideSuggestions();
                return;
            }
            
            try {
                var resultsJSON = searchHistoryGo(query);
                var results = JSON.parse(resultsJSON);
                showSuggestions(results);
            } catch (e) {
                console.error('Erro ao buscar histórico:', e);
            }
        }
        
        function navigate() {
            var input = document.getElementById('url-bar').value.trim();
            if (!input) return;
            
            hideSuggestions();
            
            var finalUrl = processInputGo(input);
            if (finalUrl) {
                navigateGo(finalUrl);
            }
        }
        
        function goBack() {
            if (typeof goBackGo === 'function') {
                goBackGo();
            }
        }
        
        function goForward() {
            if (typeof goForwardGo === 'function') {
                goForwardGo();
            }
        }
        
        function reload() {
            if (typeof reloadGo === 'function') {
                reloadGo();
            }
        }
        
        function openSettings() {
            alert('Configurações não implementadas ainda');
        }
        
        function showSuggestions(results) {
            var suggestionsDiv = document.getElementById('suggestions');
            if (!results || results.length === 0) {
                hideSuggestions();
                return;
            }
            
            suggestionsDiv.innerHTML = '';
            results.slice(0, 10).forEach(function(entry) {
                var div = document.createElement('div');
                div.className = 'suggestion';
                var url = entry.URL || entry.url;
                var title = entry.Title || entry.title || url;
                
                div.innerHTML = 
                    '<div style="font-weight: bold; margin-bottom: 4px;">' + escapeHtml(title) + '</div>' +
                    '<div style="font-size: 13px; color: #aaa;">' + escapeHtml(url) + '</div>';
                
                div.onclick = function() {
                    document.getElementById('url-bar').value = url;
                    navigate();
                };
                suggestionsDiv.appendChild(div);
            });
            suggestionsDiv.style.display = 'block';
        }
        
        function hideSuggestions() {
            document.getElementById('suggestions').style.display = 'none';
        }
        
        function escapeHtml(text) {
            var div = document.createElement('div');
            div.textContent = text;
            return div.innerHTML;
        }
    </script>
</body>
</html>`, defaultURL)
}
