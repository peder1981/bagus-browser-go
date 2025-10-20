package security

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

// Blocker gerencia bloqueio de domínios
type Blocker struct {
	domains map[string]bool
	mu      sync.RWMutex
	logger  *log.Logger
}

// NewBlocker cria novo bloqueador
func NewBlocker(blockListPath string, logger *log.Logger) (*Blocker, error) {
	b := &Blocker{
		domains: make(map[string]bool),
		logger:  logger,
	}

	if err := b.LoadBlockList(blockListPath); err != nil {
		return nil, err
	}

	return b, nil
}

// LoadBlockList carrega lista de domínios bloqueados
func (b *Blocker) LoadBlockList(path string) error {
	if path == "" {
		return nil
	}

	// Verifica se arquivo existe
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Printf("Lista de bloqueio não encontrada: %s", path)
		return nil
	}

	// Valida tamanho (max 10MB)
	info, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("erro ao verificar arquivo: %w", err)
	}
	if info.Size() > 10*1024*1024 {
		return fmt.Errorf("arquivo de bloqueio muito grande")
	}

	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("erro ao abrir lista de bloqueio: %w", err)
	}
	defer file.Close()

	b.mu.Lock()
	defer b.mu.Unlock()

	scanner := bufio.NewScanner(file)
	count := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Ignora linhas vazias e comentários
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Adiciona domínio em lowercase
		domain := strings.ToLower(line)
		b.domains[domain] = true
		count++
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("erro ao ler lista de bloqueio: %w", err)
	}

	log.Printf("Carregados %d domínios bloqueados", count)
	return nil
}

// IsBlocked verifica se domínio está bloqueado
func (b *Blocker) IsBlocked(domain string) bool {
	b.mu.RLock()
	defer b.mu.RUnlock()

	domain = strings.ToLower(domain)

	// Verifica domínio completo
	if b.domains[domain] {
		if b.logger != nil {
			b.logger.Printf("Bloqueado: %s", domain)
		}
		return true
	}

	// Verifica subdomínios (ex: ads.example.com -> example.com)
	parts := strings.Split(domain, ".")
	if len(parts) > 2 {
		// Tenta domínio base
		baseDomain := strings.Join(parts[len(parts)-2:], ".")
		if b.domains[baseDomain] {
			if b.logger != nil {
				b.logger.Printf("Bloqueado (base): %s", domain)
			}
			return true
		}
	}

	return false
}

// AddDomain adiciona domínio à lista de bloqueio
func (b *Blocker) AddDomain(domain string) {
	b.mu.Lock()
	defer b.mu.Unlock()

	domain = strings.ToLower(strings.TrimSpace(domain))
	if domain != "" {
		b.domains[domain] = true
	}
}

// RemoveDomain remove domínio da lista de bloqueio
func (b *Blocker) RemoveDomain(domain string) {
	b.mu.Lock()
	defer b.mu.Unlock()

	domain = strings.ToLower(strings.TrimSpace(domain))
	delete(b.domains, domain)
}

// Count retorna número de domínios bloqueados
func (b *Blocker) Count() int {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return len(b.domains)
}

// Clear limpa lista de bloqueio
func (b *Blocker) Clear() {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.domains = make(map[string]bool)
}
