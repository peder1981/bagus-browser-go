package browser

import (
	"log"
	"sync"

	"github.com/peder1981/bagus-browser-go/pkg/ipc"
)

// Coordinator gerencia a comunicação entre janelas de controle e conteúdo
type Coordinator struct {
	channel        *ipc.Channel
	controlWindow  *ControlWindow
	contentWindow  *ContentWindow
	mu             sync.RWMutex
	running        bool
}

// NewCoordinator cria um novo coordenador
func NewCoordinator() *Coordinator {
	return &Coordinator{
		channel: ipc.NewChannel(),
	}
}

// Start inicia o coordenador - apenas WebView (sem janela de controle por enquanto)
func (c *Coordinator) Start() error {
	c.mu.Lock()
	if c.running {
		c.mu.Unlock()
		return nil
	}
	c.running = true
	c.mu.Unlock()

	// Criar janela de conteúdo (WebView simples)
	contentWindow, err := NewContentWindow(c.channel)
	if err != nil {
		return err
	}
	c.contentWindow = contentWindow

	// Iniciar processamento de mensagens
	go c.processMessages()

	// Iniciar janela de conteúdo (bloqueia até fechar)
	if err := c.contentWindow.Run(); err != nil {
		return err
	}

	return nil
}

// processMessages processa mensagens entre janelas
func (c *Coordinator) processMessages() {
	for {
		select {
		case msg, ok := <-c.channel.ReceiveFromContent():
			if !ok {
				return
			}
			c.handleContentMessage(msg)
			
		case msg, ok := <-c.channel.ReceiveFromControl():
			if !ok {
				return
			}
			c.handleControlMessage(msg)
		}
	}
}

// handleContentMessage processa mensagens da janela de conteúdo
func (c *Coordinator) handleContentMessage(msg ipc.Message) {
	switch msg.Type {
	case ipc.MsgURLChanged:
		if url, ok := msg.Payload.(string); ok {
			log.Printf("URL alterada: %s", url)
		}
	case ipc.MsgTitleChanged:
		if title, ok := msg.Payload.(string); ok {
			log.Printf("Título: %s", title)
		}
	case ipc.MsgLoadingStarted:
		log.Printf("Carregando...")
	case ipc.MsgLoadingFinished:
		log.Printf("Carregamento concluído")
	case ipc.MsgLoadingFailed:
		if err, ok := msg.Payload.(error); ok {
			log.Printf("Erro ao carregar: %v", err)
		}
	}
}

// handleControlMessage processa mensagens da janela de controle
func (c *Coordinator) handleControlMessage(msg ipc.Message) {
	// Mensagens já são enviadas diretamente para o canal
	// Este método pode ser usado para logging ou processamento adicional
	log.Printf("Comando de controle: %s", msg.Type)
}

// Stop para o coordenador
func (c *Coordinator) Stop() {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	if !c.running {
		return
	}
	
	c.running = false
	c.channel.Close()
	
	if c.controlWindow != nil {
		c.controlWindow.Close()
	}
	if c.contentWindow != nil {
		c.contentWindow.Close()
	}
}
