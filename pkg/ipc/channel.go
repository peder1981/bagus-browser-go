package ipc

import (
	"sync"
)

// MessageType define os tipos de mensagens entre janelas
type MessageType string

const (
	// Control → Content
	MsgNavigate MessageType = "navigate"
	MsgBack     MessageType = "back"
	MsgForward  MessageType = "forward"
	MsgReload   MessageType = "reload"
	MsgStop     MessageType = "stop"
	
	// Content → Control
	MsgURLChanged      MessageType = "url_changed"
	MsgTitleChanged    MessageType = "title_changed"
	MsgLoadingStarted  MessageType = "loading_started"
	MsgLoadingFinished MessageType = "loading_finished"
	MsgLoadingFailed   MessageType = "loading_failed"
)

// Message representa uma mensagem entre janelas
type Message struct {
	Type    MessageType
	Payload interface{}
}

// Channel gerencia comunicação bidirecional entre janelas
type Channel struct {
	controlToContent chan Message
	contentToControl chan Message
	mu               sync.RWMutex
	closed           bool
}

// NewChannel cria um novo canal de comunicação
func NewChannel() *Channel {
	return &Channel{
		controlToContent: make(chan Message, 10),
		contentToControl: make(chan Message, 10),
	}
}

// SendToContent envia mensagem da janela de controle para conteúdo
func (c *Channel) SendToContent(msg Message) error {
	c.mu.RLock()
	defer c.mu.RUnlock()
	
	if c.closed {
		return ErrChannelClosed
	}
	
	select {
	case c.controlToContent <- msg:
		return nil
	default:
		return ErrChannelFull
	}
}

// SendToControl envia mensagem da janela de conteúdo para controle
func (c *Channel) SendToControl(msg Message) error {
	c.mu.RLock()
	defer c.mu.RUnlock()
	
	if c.closed {
		return ErrChannelClosed
	}
	
	select {
	case c.contentToControl <- msg:
		return nil
	default:
		return ErrChannelFull
	}
}

// ReceiveFromControl recebe mensagens na janela de conteúdo
func (c *Channel) ReceiveFromControl() <-chan Message {
	return c.controlToContent
}

// ReceiveFromContent recebe mensagens na janela de controle
func (c *Channel) ReceiveFromContent() <-chan Message {
	return c.contentToControl
}

// Close fecha o canal
func (c *Channel) Close() {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	if !c.closed {
		c.closed = true
		close(c.controlToContent)
		close(c.contentToControl)
	}
}

// Errors
var (
	ErrChannelClosed = &ChannelError{"channel is closed"}
	ErrChannelFull   = &ChannelError{"channel is full"}
)

// ChannelError representa um erro do canal
type ChannelError struct {
	msg string
}

func (e *ChannelError) Error() string {
	return e.msg
}
