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

// Start inicia o coordenador e ambas as janelas
func (c *Coordinator) Start() error {
	c.mu.Lock()
	if c.running {
		c.mu.Unlock()
		return nil
	}
	c.running = true
	c.mu.Unlock()

	// Criar janela de conteúdo
	contentWindow, err := NewContentWindow(c.channel)
	if err != nil {
		return err
	}
	c.contentWindow = contentWindow

	// Criar janela de controle
	controlWindow, err := NewControlWindow(c.channel)
	if err != nil {
		return err
	}
	c.controlWindow = controlWindow

	// Iniciar processamento de mensagens
	go c.processMessages()

	// Iniciar janela de conteúdo em goroutine
	go func() {
		if err := c.contentWindow.Run(); err != nil {
			log.Printf("Erro na janela de conteúdo: %v", err)
		}
	}()

	// Iniciar janela de controle (bloqueia até fechar)
	if err := c.controlWindow.Run(); err != nil {
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
			c.controlWindow.UpdateURL(url)
		}
	case ipc.MsgTitleChanged:
		if title, ok := msg.Payload.(string); ok {
			c.controlWindow.UpdateTitle(title)
		}
	case ipc.MsgLoadingStarted:
		c.controlWindow.SetLoading(true)
	case ipc.MsgLoadingFinished:
		c.controlWindow.SetLoading(false)
	case ipc.MsgLoadingFailed:
		c.controlWindow.SetLoading(false)
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
