package ui

import (
	"encoding/json"
	"fmt"
	"sync"

	webview "github.com/webview/webview_go"
)

// ControlWindow gerencia a janela de controle (barra de navegação)
type ControlWindow struct {
	w            webview.WebView
	currentURL   string
	contentReady chan bool
	mu           sync.RWMutex
}

// NewControlWindow cria nova janela de controle
func NewControlWindow() *ControlWindow {
	return &ControlWindow{
		currentURL:   "https://duckduckgo.com/",
		contentReady: make(chan bool, 1),
	}
}

// Run inicia a janela de controle
func (c *ControlWindow) Run(onNavigate func(string), onSearch func(string) []interface{}) {
	debug := false
	w := webview.New(debug)
	defer w.Destroy()

	c.w = w

	w.SetTitle("Bagus Browser - Controles")
	w.SetSize(1200, 80, webview.HintNone)

	// Bind funções
	w.Bind("navigateGo", func(url string) {
		c.mu.Lock()
		c.currentURL = url
		c.mu.Unlock()

		if onNavigate != nil {
			onNavigate(url)
		}
	})

	w.Bind("searchHistoryGo", func(query string) string {
		if onSearch == nil {
			return "[]"
		}
		results := onSearch(query)
		data, _ := json.Marshal(results)
		return string(data)
	})

	w.Bind("processInputGo", func(input string) string {
		// Retorna a URL processada (será implementada pelo Browser)
		return input
	})

	// Carrega HTML da barra de controle
	html := c.generateHTML()
	w.SetHtml(html)

	// Atualiza URL inicial
	go func() {
		c.contentReady <- true
	}()

	w.Run()
}

// UpdateURL atualiza o campo de URL na barra
func (c *ControlWindow) UpdateURL(url string) {
	c.mu.Lock()
	c.currentURL = url
	c.mu.Unlock()

	if c.w != nil {
		escapedURL := escapeJS(url)
		js := fmt.Sprintf(`
			var urlBar = document.getElementById('urlBar');
			if (urlBar) {
				urlBar.value = '%s';
			}
		`, escapedURL)
		c.w.Eval(js)
	}
}

// generateHTML gera HTML da barra de controle
func (c *ControlWindow) generateHTML() string {
	return `<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>Bagus Browser - Controles</title>
    <style>
        * { margin: 0; padding: 0; box-sizing: border-box; }
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
            background: #2d2d2d;
            color: #fff;
            padding: 10px;
        }
        .toolbar {
            display: flex;
            gap: 10px;
            align-items: center;
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
        }
        .suggestion:hover { background: #3c3c3c; }
        .suggestion:last-child { border-bottom: none; }
    </style>
</head>
<body>
    <div class="toolbar">
        <button onclick="goBack()" title="Voltar">◀</button>
        <button onclick="goForward()" title="Avançar">▶</button>
        <button onclick="reload()" title="Recarregar">⟳</button>
        <input type="text" id="urlBar" placeholder="Digite uma URL ou termo de busca..." 
               onkeyup="handleUrlInput(event)" 
               onkeydown="if(event.key==='Enter') navigate()"
               onfocus="hideSuggestions()"
               value="https://duckduckgo.com/">
        <button onclick="openSettings()" title="Configurações">⚙</button>
    </div>
    <div class="suggestions" id="suggestions"></div>

    <script>
        function navigate() {
            const input = document.getElementById('urlBar').value.trim();
            if (!input) return;
            
            hideSuggestions();
            
            // Processa input (URL ou busca)
            const finalUrl = processInput(input);
            if (finalUrl) {
                navigateGo(finalUrl);
            }
        }
        
        function processInput(input) {
            // Verifica se parece com uma URL
            const isURL = input.includes('.') && 
                (input.startsWith('http://') || 
                 input.startsWith('https://') ||
                 !input.includes(' '));

            if (isURL) {
                // Adiciona protocolo se necessário
                if (!input.startsWith('http://') && !input.startsWith('https://')) {
                    return 'https://' + input;
                }
                return input;
            }

            // Se não for URL, faz busca no DuckDuckGo
            return 'https://duckduckgo.com/?q=' + encodeURIComponent(input);
        }
        
        function hideSuggestions() {
            document.getElementById('suggestions').style.display = 'none';
        }

        function goBack() {
            console.log('Voltar não implementado ainda');
        }
        
        function goForward() {
            console.log('Avançar não implementado ainda');
        }
        
        function reload() {
            const url = document.getElementById('urlBar').value;
            if (url) {
                navigateGo(url);
            }
        }
        
        function openSettings() {
            console.log('Configurações não implementadas ainda');
        }

        function handleUrlInput(event) {
            if (event.key === 'Enter' || event.key === 'Escape') {
                return;
            }
            
            const query = event.target.value.trim();
            if (query.length < 2) {
                hideSuggestions();
                return;
            }

            try {
                const resultsJSON = searchHistoryGo(query);
                const results = JSON.parse(resultsJSON);
                showSuggestions(results);
            } catch (e) {
                console.error('Erro ao buscar histórico:', e);
            }
        }

        function showSuggestions(results) {
            const suggestionsDiv = document.getElementById('suggestions');
            if (!results || results.length === 0) {
                hideSuggestions();
                return;
            }

            suggestionsDiv.innerHTML = '';
            results.slice(0, 10).forEach(entry => {
                const div = document.createElement('div');
                div.className = 'suggestion';
                const url = entry.URL || entry.url;
                const title = entry.Title || entry.title || url;
                
                div.innerHTML = 
                    '<div style="font-weight: bold; margin-bottom: 4px;">' + escapeHtml(title) + '</div>' +
                    '<div style="font-size: 13px; color: #aaa;">' + escapeHtml(url) + '</div>';
                
                div.onclick = () => {
                    document.getElementById('urlBar').value = url;
                    navigate();
                };
                suggestionsDiv.appendChild(div);
            });
            suggestionsDiv.style.display = 'block';
        }
        
        function escapeHtml(text) {
            const div = document.createElement('div');
            div.textContent = text;
            return div.innerHTML;
        }
    </script>
</body>
</html>`
}

// escapeJS escapa string para uso em JavaScript
func escapeJS(s string) string {
	s = fmt.Sprintf("%q", s)
	return s[1 : len(s)-1]
}
