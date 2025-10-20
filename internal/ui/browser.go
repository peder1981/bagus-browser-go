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

// Browser representa a janela principal do navegador
type Browser struct {
	w        webview.WebView
	engine   *browser.Engine
	config   *config.Config
	blocker  *security.Blocker
	history  *browser.History
	tabs     []*BrowserTab
	current  int
	basePath string
}

// BrowserTab representa uma aba do navegador
type BrowserTab struct {
	URL   string
	Title string
}

// NewBrowser cria nova instância do browser
func NewBrowser(userPath string, cfg *config.Config) (*Browser, error) {
	// Cria storage manager
	storageManager, err := browser.NewStorageManager(userPath)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar storage: %w", err)
	}

	// Cria engine
	eng, err := browser.NewEngine(storageManager, false)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar engine: %w", err)
	}

	// Cria histórico
	hist, err := browser.NewHistory(userPath)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar histórico: %w", err)
	}

	// Cria bloqueador
	blockListPath := filepath.Join(userPath, "ad_hosts_block.txt")
	logger := log.Default()
	blocker, err := security.NewBlocker(blockListPath, logger)
	if err != nil {
		log.Printf("Aviso: erro ao criar bloqueador: %v", err)
		blocker = nil
	}

	b := &Browser{
		engine:   eng,
		config:   cfg,
		blocker:  blocker,
		history:  hist,
		tabs:     make([]*BrowserTab, 0),
		current:  -1,
		basePath: userPath,
	}

	return b, nil
}

// Run inicia o browser com UI
func (b *Browser) Run() error {
	debug := true
	w := webview.New(debug)
	defer w.Destroy()

	b.w = w

	w.SetTitle("Bagus Browser")
	w.SetSize(1200, 800, webview.HintNone)

	// Bind funções Go para JavaScript
	b.bindFunctions()

	// Navega diretamente para URL inicial
	defaultURL := b.config.Default.URL
	log.Printf("Navegando para: %s", defaultURL)
	w.Navigate(defaultURL)

	w.Run()
	return nil
}

// bindFunctions registra funções Go para serem chamadas do JavaScript
func (b *Browser) bindFunctions() {
	// Nova aba
	b.w.Bind("newTabGo", func(url string) {
		b.newTab(url)
	})

	// Navegar
	b.w.Bind("navigateGo", func(url string) {
		b.navigate(url)
	})

	// Voltar
	b.w.Bind("goBackGo", func() {
		b.goBack()
	})

	// Avançar
	b.w.Bind("goForwardGo", func() {
		b.goForward()
	})

	// Recarregar
	b.w.Bind("reloadGo", func() {
		b.reload()
	})

	// Fechar aba
	b.w.Bind("closeTabGo", func(index int) {
		b.closeTab(index)
	})

	// Buscar histórico
	b.w.Bind("searchHistoryGo", func(query string) string {
		results := b.history.Search(query)
		data, _ := json.Marshal(results)
		return string(data)
	})

	// Obter tabs
	b.w.Bind("getTabsGo", func() string {
		data, _ := json.Marshal(b.tabs)
		return string(data)
	})

	// Processar input da barra de navegação
	b.w.Bind("processInputGo", func(input string) string {
		return b.processInput(input)
	})
}

// newTab cria nova aba (simplificado - navega diretamente)
func (b *Browser) newTab(url string) {
	if url == "" {
		url = b.config.Default.URL
	}

	b.navigate(url)
}

// processInput processa input da barra de navegação
// Retorna a URL final para navegação
func (b *Browser) processInput(input string) string {
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

// navigate navega para URL
func (b *Browser) navigate(url string) {
	url = security.SanitizeURL(url)

	if err := security.ValidateURL(url); err != nil {
		log.Printf("URL inválida: %v", err)
		return
	}

	// Verifica bloqueio
	if b.blocker != nil {
		domain, err := security.ExtractDomain(url)
		if err == nil && b.blocker.IsBlocked(domain) {
			log.Printf("Domínio bloqueado: %s", domain)
			return
		}
	}

	log.Printf("Navegando para: %s", url)
	b.history.Add(url, url)
	
	// Navega diretamente no webview
	if b.w != nil {
		b.w.Navigate(url)
	}
}

// goBack volta na navegação
func (b *Browser) goBack() {
	// Implementação simplificada
	log.Println("Voltar")
}

// goForward avança na navegação
func (b *Browser) goForward() {
	log.Println("Avançar")
}

// reload recarrega página
func (b *Browser) reload() {
	if b.current >= 0 && b.current < len(b.tabs) {
		log.Printf("Recarregando: %s", b.tabs[b.current].URL)
	}
}

// closeTab fecha aba
func (b *Browser) closeTab(index int) {
	if index < 0 || index >= len(b.tabs) {
		return
	}

	if len(b.tabs) == 1 {
		// Não fecha última aba
		return
	}

	b.tabs = append(b.tabs[:index], b.tabs[index+1:]...)

	if b.current >= len(b.tabs) {
		b.current = len(b.tabs) - 1
	}

	b.updateUI()
}

// updateUI atualiza interface
func (b *Browser) updateUI() {
	if b.w == nil {
		return
	}

	// Atualiza lista de tabs
	tabsJSON, _ := json.Marshal(b.tabs)
	js := fmt.Sprintf("updateTabs(%s, %d)", string(tabsJSON), b.current)
	b.w.Eval(js)

	// Atualiza iframe com URL atual
	if b.current >= 0 && b.current < len(b.tabs) {
		url := b.tabs[b.current].URL
		js = fmt.Sprintf("updateIframe('%s')", url)
		b.w.Eval(js)
	}
}

// generateHTML gera HTML da interface
func (b *Browser) generateHTML() string {
	return `<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>Bagus Browser</title>
    <style>
        * { margin: 0; padding: 0; box-sizing: border-box; }
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
            display: flex;
            flex-direction: column;
            height: 100vh;
            background: #1e1e1e;
            color: #fff;
        }
        .toolbar {
            display: flex;
            gap: 10px;
            padding: 10px;
            background: #2d2d2d;
            border-bottom: 1px solid #444;
        }
        .toolbar button {
            padding: 8px 16px;
            background: #0078d4;
            color: white;
            border: none;
            border-radius: 4px;
            cursor: pointer;
            font-size: 14px;
        }
        .toolbar button:hover { background: #106ebe; }
        .toolbar input {
            flex: 1;
            padding: 8px 12px;
            background: #3c3c3c;
            border: 1px solid #555;
            border-radius: 4px;
            color: #fff;
            font-size: 14px;
        }
        .tabs {
            display: flex;
            gap: 5px;
            padding: 5px 10px;
            background: #252525;
            border-bottom: 1px solid #444;
            overflow-x: auto;
        }
        .tab {
            padding: 8px 16px;
            background: #3c3c3c;
            border-radius: 4px 4px 0 0;
            cursor: pointer;
            white-space: nowrap;
            font-size: 13px;
        }
        .tab.active { background: #0078d4; }
        .tab:hover { background: #4a4a4a; }
        .tab.active:hover { background: #106ebe; }
        .content {
            flex: 1;
            display: flex;
            position: relative;
        }
        iframe {
            width: 100%;
            height: 100%;
            border: none;
            background: white;
        }
        .suggestions {
            position: absolute;
            top: 50px;
            left: 10px;
            right: 10px;
            background: #2d2d2d;
            border: 1px solid #555;
            border-radius: 4px;
            max-height: 300px;
            overflow-y: auto;
            display: none;
            z-index: 1000;
        }
        .suggestion {
            padding: 10px;
            cursor: pointer;
            border-bottom: 1px solid #444;
        }
        .suggestion:hover { background: #3c3c3c; }
    </style>
</head>
<body>
    <div class="toolbar">
        <button onclick="goBack()">◀</button>
        <button onclick="goForward()">▶</button>
        <button onclick="reload()">⟳</button>
        <input type="text" id="urlBar" placeholder="Digite uma URL ou termo de busca..." 
               onkeyup="handleUrlInput(event)" onkeydown="if(event.key==='Enter') navigate()"
               onfocus="hideSuggestions()">
        <button onclick="newTab()">+ Nova Aba</button>
    </div>
    <div class="tabs" id="tabs"></div>
    <div class="suggestions" id="suggestions"></div>
    <div class="content">
        <iframe id="webview" sandbox="allow-same-origin allow-scripts allow-forms allow-popups"></iframe>
    </div>

    <script>
        let currentTabs = [];
        let currentIndex = 0;

        function newTab() {
            const url = document.getElementById('urlBar').value || 'https://duckduckgo.com/';
            newTabGo(url);
        }

        function navigate() {
            const input = document.getElementById('urlBar').value.trim();
            if (!input) return;
            
            // Oculta sugestões
            hideSuggestions();
            
            // Processa input (URL ou busca)
            const finalUrl = processInputGo(input);
            if (finalUrl) {
                navigateGo(finalUrl);
            }
        }
        
        function hideSuggestions() {
            document.getElementById('suggestions').style.display = 'none';
        }

        function goBack() { goBackGo(); }
        function goForward() { goForwardGo(); }
        function reload() { reloadGo(); }

        function updateTabs(tabs, current) {
            currentTabs = tabs;
            currentIndex = current;
            
            const tabsDiv = document.getElementById('tabs');
            tabsDiv.innerHTML = '';
            
            tabs.forEach((tab, index) => {
                const tabEl = document.createElement('div');
                tabEl.className = 'tab' + (index === current ? ' active' : '');
                tabEl.textContent = tab.Title || 'Nova Aba';
                tabEl.onclick = () => switchTab(index);
                tabsDiv.appendChild(tabEl);
            });
        }

        function updateIframe(url) {
            const iframe = document.getElementById('webview');
            iframe.src = url;
            document.getElementById('urlBar').value = url;
        }

        function switchTab(index) {
            if (index >= 0 && index < currentTabs.length) {
                currentIndex = index;
                updateTabs(currentTabs, index);
                updateIframe(currentTabs[index].URL);
            }
        }

        function handleUrlInput(event) {
            // Se pressionar Enter, navega
            if (event.key === 'Enter') {
                return;
            }
            
            // Se pressionar Escape, oculta sugestões
            if (event.key === 'Escape') {
                hideSuggestions();
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
                    '<div style="font-weight: bold; margin-bottom: 2px;">' + escapeHtml(title) + '</div>' +
                    '<div style="font-size: 12px; color: #aaa;">' + escapeHtml(url) + '</div>';
                
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

        // Inicialização
        window.addEventListener('load', () => {
            console.log('Bagus Browser carregado');
            // Carrega primeira aba
            newTabGo('https://duckduckgo.com/');
        });
    </script>
</body>
</html>`
}
