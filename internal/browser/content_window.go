package browser

import (
	"log"

	"github.com/peder1981/bagus-browser-go/pkg/ipc"
	webview "github.com/webview/webview_go"
)

// ContentWindow representa a janela de conteúdo (WebView)
type ContentWindow struct {
	webview webview.WebView
	channel *ipc.Channel
	currentURL string
}

// NewContentWindow cria uma nova janela de conteúdo
func NewContentWindow(channel *ipc.Channel) (*ContentWindow, error) {
	w := webview.New(false) // false = não debug
	if w == nil {
		return nil, ErrWebViewInit
	}

	cw := &ContentWindow{
		webview: w,
		channel: channel,
	}

	// Configurar janela
	w.SetTitle("Bagus Browser v3.0.0")
	w.SetSize(1200, 800, webview.HintNone)

	// Injetar controles de navegação
	w.Init(cw.getNavigationScript())

	return cw, nil
}

// Run inicia a janela de conteúdo
func (cw *ContentWindow) Run() error {
	defer cw.webview.Destroy()

	// Processar mensagens de controle em goroutine
	go cw.processControlMessages()

	// Navegar para página inicial
	cw.Navigate("https://duckduckgo.com")

	// Iniciar loop da janela (bloqueia)
	cw.webview.Run()
	
	return nil
}

// processControlMessages processa mensagens da janela de controle
func (cw *ContentWindow) processControlMessages() {
	for msg := range cw.channel.ReceiveFromControl() {
		switch msg.Type {
		case ipc.MsgNavigate:
			if url, ok := msg.Payload.(string); ok {
				cw.Navigate(url)
			}
		case ipc.MsgBack:
			cw.Back()
		case ipc.MsgForward:
			cw.Forward()
		case ipc.MsgReload:
			cw.Reload()
		case ipc.MsgStop:
			cw.Stop()
		}
	}
}

// Navigate navega para uma URL
func (cw *ContentWindow) Navigate(url string) {
	log.Printf("Navegando para: %s", url)
	cw.currentURL = url
	
	// Notificar início do carregamento
	cw.channel.SendToControl(ipc.Message{
		Type: ipc.MsgLoadingStarted,
	})
	
	// Navegar
	cw.webview.Navigate(url)
	
	// Notificar mudança de URL
	cw.channel.SendToControl(ipc.Message{
		Type:    ipc.MsgURLChanged,
		Payload: url,
	})
	
	// Notificar fim do carregamento (simplificado)
	// Em produção, usar callbacks do webview
	go func() {
		cw.channel.SendToControl(ipc.Message{
			Type: ipc.MsgLoadingFinished,
		})
	}()
}

// Back volta para página anterior
func (cw *ContentWindow) Back() {
	cw.webview.Eval("window.history.back()")
}

// Forward avança para próxima página
func (cw *ContentWindow) Forward() {
	cw.webview.Eval("window.history.forward()")
}

// Reload recarrega a página atual
func (cw *ContentWindow) Reload() {
	cw.webview.Eval("window.location.reload()")
}

// Stop para o carregamento
func (cw *ContentWindow) Stop() {
	cw.webview.Eval("window.stop()")
}

// Close fecha a janela
func (cw *ContentWindow) Close() {
	cw.webview.Terminate()
}

// getNavigationScript retorna o script de controles de navegação
func (cw *ContentWindow) getNavigationScript() string {
	return `
		(function() {
			// Previne múltiplas injeções
			if (window.bagusControlsInjected) return;
			window.bagusControlsInjected = true;

			// Atalhos de teclado
			document.addEventListener('keydown', function(e) {
				// Ctrl+L - Navegar
				if (e.ctrlKey && e.key === 'l') {
					e.preventDefault();
					var url = prompt('Digite uma URL ou termo de busca:', window.location.href);
					if (url) {
						window.location.href = url.includes(' ') 
							? 'https://duckduckgo.com/?q=' + encodeURIComponent(url)
							: (url.startsWith('http') ? url : 'https://' + url);
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
			});

			// Mensagem de boas-vindas
			console.log('%c🌐 Bagus Browser v3.0.0', 'font-size: 20px; font-weight: bold; color: #0078d4;');
			console.log('%cAtalhos de Teclado:', 'font-size: 14px; font-weight: bold; margin-top: 10px;');
			console.log('  Ctrl+L     - Navegar para URL');
			console.log('  Ctrl+R/F5  - Recarregar página');
			console.log('  Alt+←      - Voltar');
			console.log('  Alt+→      - Avançar');
			console.log('%cNavegação Simples e Privada! 🔒', 'font-size: 14px; font-weight: bold; margin-top: 10px; color: #28a745;');
		})();
	`
}

// Errors
var (
	ErrWebViewInit = &BrowserError{"failed to initialize webview"}
)

// BrowserError representa um erro do browser
type BrowserError struct {
	msg string
}

func (e *BrowserError) Error() string {
	return e.msg
}
