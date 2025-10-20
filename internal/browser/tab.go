package browser

import (
	"fmt"
	"log"
)

// Tab representa uma aba do navegador
type Tab struct {
	URL     string
	Title   string
	engine  *Engine
	history []string
}

// Navigate navega para uma URL
func (t *Tab) Navigate(url string) error {
	if url == "" {
		return fmt.Errorf("URL não pode ser vazia")
	}

	// Adiciona ao histórico
	t.history = append(t.history, t.URL)
	t.URL = url

	if t.engine.debug {
		log.Printf("Navegando para: %s", url)
	}

	return nil
}

// Back volta para a página anterior
func (t *Tab) Back() error {
	if len(t.history) == 0 {
		return fmt.Errorf("não há histórico para voltar")
	}

	// Pega a última URL do histórico
	lastURL := t.history[len(t.history)-1]
	t.history = t.history[:len(t.history)-1]
	t.URL = lastURL

	if t.engine.debug {
		log.Printf("Voltando para: %s", lastURL)
	}

	return nil
}

// Reload recarrega a página atual
func (t *Tab) Reload() error {
	if t.engine.debug {
		log.Printf("Recarregando: %s", t.URL)
	}

	// TODO: Implementar reload real
	return nil
}

// GetHistory retorna o histórico da aba
func (t *Tab) GetHistory() []string {
	return t.history
}
