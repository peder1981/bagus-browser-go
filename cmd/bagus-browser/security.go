package main

import (
	"fmt"
	"net/url"
	"strings"
)

// URLValidator valida e sanitiza URLs
type URLValidator struct {
	blockedDomains []string
}

// NewURLValidator cria um novo validador
func NewURLValidator() *URLValidator {
	return &URLValidator{
		blockedDomains: []string{
			// Lista de domínios bloqueados (exemplo)
			// Pode ser expandida conforme necessário
		},
	}
}

// ValidateAndSanitize valida e sanitiza uma URL ou termo de busca
func (v *URLValidator) ValidateAndSanitize(input string) (string, error) {
	// Remover espaços em branco
	input = strings.TrimSpace(input)
	
	if input == "" {
		return "", fmt.Errorf("entrada vazia")
	}
	
	// Detectar se é URL ou termo de busca
	if v.isURL(input) {
		return v.validateURL(input)
	}
	
	// Se não é URL, criar busca no DuckDuckGo
	return v.createSearchURL(input), nil
}

// isURL detecta se o input parece uma URL
func (v *URLValidator) isURL(input string) bool {
	// Tem protocolo
	if strings.HasPrefix(input, "http://") || strings.HasPrefix(input, "https://") {
		return true
	}
	
	// Tem ponto e não tem espaços (provavelmente domínio)
	if strings.Contains(input, ".") && !strings.Contains(input, " ") {
		return true
	}
	
	// Começa com localhost
	if strings.HasPrefix(input, "localhost") {
		return true
	}
	
	return false
}

// validateURL valida uma URL
func (v *URLValidator) validateURL(input string) (string, error) {
	// Adicionar https:// se não tiver protocolo
	if !strings.HasPrefix(input, "http://") && !strings.HasPrefix(input, "https://") {
		input = "https://" + input
	}
	
	// Parse URL
	u, err := url.Parse(input)
	if err != nil {
		return "", fmt.Errorf("URL inválida: %v", err)
	}
	
	// Validar scheme
	if u.Scheme != "http" && u.Scheme != "https" {
		return "", fmt.Errorf("protocolo não suportado: %s (use http ou https)", u.Scheme)
	}
	
	// Validar host
	if u.Host == "" {
		return "", fmt.Errorf("host inválido")
	}
	
	// Verificar se domínio está bloqueado
	if v.isDomainBlocked(u.Host) {
		return "", fmt.Errorf("domínio bloqueado: %s", u.Host)
	}
	
	return u.String(), nil
}

// createSearchURL cria URL de busca no DuckDuckGo
func (v *URLValidator) createSearchURL(query string) string {
	// Escapar query
	escaped := url.QueryEscape(query)
	return fmt.Sprintf("https://duckduckgo.com/?q=%s", escaped)
}

// isDomainBlocked verifica se domínio está bloqueado
func (v *URLValidator) isDomainBlocked(host string) bool {
	// Remover porta se houver
	if idx := strings.Index(host, ":"); idx != -1 {
		host = host[:idx]
	}
	
	// Verificar lista de bloqueio
	for _, blocked := range v.blockedDomains {
		if host == blocked || strings.HasSuffix(host, "."+blocked) {
			return true
		}
	}
	
	return false
}

// AddBlockedDomain adiciona domínio à lista de bloqueio
func (v *URLValidator) AddBlockedDomain(domain string) {
	v.blockedDomains = append(v.blockedDomains, domain)
}

// SanitizeInput sanitiza input do usuário
func SanitizeInput(input string) string {
	// Remover caracteres perigosos
	input = strings.TrimSpace(input)
	
	// Remover quebras de linha
	input = strings.ReplaceAll(input, "\n", "")
	input = strings.ReplaceAll(input, "\r", "")
	
	// Remover tabs
	input = strings.ReplaceAll(input, "\t", "")
	
	// Limitar tamanho (proteção contra input muito grande)
	if len(input) > 2048 {
		input = input[:2048]
	}
	
	return input
}
