package ui

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/peder1981/bagus-browser-go/internal/browser"
	"github.com/peder1981/bagus-browser-go/internal/config"
	"github.com/peder1981/bagus-browser-go/internal/security"
	webview "github.com/webview/webview_go"
)

// BrowserMultiProcess representa o navegador com múltiplos processos
type BrowserMultiProcess struct {
	w           webview.WebView
	config      *config.Config
	history     *browser.History
	blocker     *security.Blocker
	userPath    string
	controlFile string
	contentCmd  *exec.Cmd
}

// NewBrowserMultiProcess cria nova instância do browser
func NewBrowserMultiProcess(userPath string, cfg *config.Config) (*BrowserMultiProcess, error) {
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

	// Arquivo de controle para comunicação
	controlFile := filepath.Join(os.TempDir(), fmt.Sprintf("bagus-control-%d.json", os.Getpid()))

	b := &BrowserMultiProcess{
		config:      cfg,
		history:     hist,
		blocker:     blocker,
		userPath:    userPath,
		controlFile: controlFile,
	}

	return b, nil
}

// Run inicia o navegador
func (b *BrowserMultiProcess) Run() error {
	// Limpa arquivo de controle ao sair
	defer func() {
		os.Remove(b.controlFile)
		if err := b.history.Save(); err != nil {
			log.Printf("Erro ao salvar histórico: %v", err)
		} else {
			log.Println("Histórico salvo com sucesso")
		}
	}()

	log.Println("Iniciando interface do browser...")

	// Inicia janela de conteúdo em processo separado
	go b.startContentWindow()

	// Aguarda um pouco para janela de conteúdo iniciar
	time.Sleep(500 * time.Millisecond)

	// Inicia janela de controle (bloqueia aqui)
	return b.runControlWindow()
}

// startContentWindow inicia janela de conteúdo em processo separado
func (b *BrowserMultiProcess) startContentWindow() {
	// Escreve configuração inicial
	config := map[string]string{
		"url": b.config.Default.URL,
	}
	b.writeControlFile(config)

	// Inicia processo da janela de conteúdo
	// Nota: Em produção, isso seria um executável separado
	// Por ora, vamos usar a mesma abordagem de janela única mas sem iframe
	go b.runContentWindowInline()
}

// runContentWindowInline executa janela de conteúdo inline
func (b *BrowserMultiProcess) runContentWindowInline() {
	debug := false
	w := webview.New(debug)
	defer w.Destroy()

	w.SetTitle("Bagus Browser - Conteúdo")
	w.SetSize(1200, 720, webview.HintNone)

	// Navega para URL inicial
	initialURL := b.config.Default.URL
	log.Printf("Janela de conteúdo: Navegando para %s", initialURL)
	w.Navigate(initialURL)

	// Monitor de mudanças no arquivo de controle
	go func() {
		lastURL := initialURL
		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()

		for range ticker.C {
			config := b.readControlFile()
			if newURL, ok := config["url"]; ok && newURL != lastURL {
				lastURL = newURL
				log.Printf("Janela de conteúdo: Navegando para %s", newURL)
				w.Navigate(newURL)
			}
		}
	}()

	w.Run()
}

// runControlWindow executa janela de controle
func (b *BrowserMultiProcess) runControlWindow() error {
	debug := false
	w := webview.New(debug)
	defer w.Destroy()

	b.w = w

	w.SetTitle("Bagus Browser - Controles")
	w.SetSize(1200, 80, webview.HintNone)

	// Bind funções
	b.bindFunctions()

	// Carrega HTML da barra de controle
	html := b.generateControlHTML()
	w.SetHtml(html)

	w.Run()
	return nil
}

// bindFunctions registra funções Go para JavaScript
func (b *BrowserMultiProcess) bindFunctions() {
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
}

// navigate processa e navega para URL
func (b *BrowserMultiProcess) navigate(input string) {
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

	// Atualiza arquivo de controle para janela de conteúdo
	log.Printf("Navegando para: %s", finalURL)
	config := map[string]string{
		"url": finalURL,
	}
	b.writeControlFile(config)

	// Atualiza campo de URL na barra
	if b.w != nil {
		escapedURL := strings.ReplaceAll(finalURL, "'", "\\'")
		js := fmt.Sprintf("document.getElementById('url-bar').value = '%s';", escapedURL)
		b.w.Eval(js)
	}
}

// processInput processa input da barra de navegação
func (b *BrowserMultiProcess) processInput(input string) string {
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

// writeControlFile escreve configuração no arquivo de controle
func (b *BrowserMultiProcess) writeControlFile(config map[string]string) {
	data, err := json.Marshal(config)
	if err != nil {
		log.Printf("Erro ao serializar config: %v", err)
		return
	}

	if err := os.WriteFile(b.controlFile, data, 0600); err != nil {
		log.Printf("Erro ao escrever arquivo de controle: %v", err)
	}
}

// readControlFile lê configuração do arquivo de controle
func (b *BrowserMultiProcess) readControlFile() map[string]string {
	data, err := os.ReadFile(b.controlFile)
	if err != nil {
		return make(map[string]string)
	}

	var config map[string]string
	if err := json.Unmarshal(data, &config); err != nil {
		return make(map[string]string)
	}

	return config
}

// generateControlHTML gera HTML da barra de controle
func (b *BrowserMultiProcess) generateControlHTML() string {
	defaultURL := b.config.Default.URL

	return fmt.Sprintf(`<!DOCTYPE html>
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
            overflow: hidden;
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

    <script>
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
            navigateGo(input);
        }
        
        function goBack() {
            console.log('Voltar não implementado');
        }
        
        function goForward() {
            console.log('Avançar não implementado');
        }
        
        function reload() {
            var url = document.getElementById('url-bar').value;
            if (url) {
                navigateGo(url);
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
