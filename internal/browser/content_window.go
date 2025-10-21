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
	w.SetTitle("Bagus Browser - Conteúdo")
	w.SetSize(1200, 800, webview.HintNone)

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
