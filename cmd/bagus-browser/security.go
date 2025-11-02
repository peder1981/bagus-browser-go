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
	// Protocolos suportados (RFC 3986 + protocolos históricos)
	supportedProtocols := []string{
		// Protocolos modernos
		"http://", "https://", "file://", "ftp://", "ftps://",
		"ws://", "wss://", "data:", "mailto:", "tel:",
		// Protocolos históricos e alternativos
		"gopher://", "gemini://", "ipfs://", "ipns://",
		"sftp://", "ssh://", "git://", "svn://",
		"rtmp://", "rtmps://", "rtsp://", "mms://",
		"news://", "nntp://", "telnet://", "irc://",
		"ircs://", "xmpp://", "sip://", "sips://",
		"magnet:", "bitcoin:", "ethereum:",
	}
	
	for _, proto := range supportedProtocols {
		if strings.HasPrefix(input, proto) {
			return true
		}
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
	// Protocolos suportados (RFC 3986 + protocolos históricos)
	supportedSchemes := map[string]bool{
		// Protocolos modernos
		"http": true, "https": true, "file": true, "ftp": true, "ftps": true,
		"ws": true, "wss": true, "data": true, "mailto": true, "tel": true,
		// Protocolos históricos e alternativos
		"gopher": true, "gemini": true, "ipfs": true, "ipns": true,
		"sftp": true, "ssh": true, "git": true, "svn": true,
		"rtmp": true, "rtmps": true, "rtsp": true, "mms": true,
		"news": true, "nntp": true, "telnet": true, "irc": true,
		"ircs": true, "xmpp": true, "sip": true, "sips": true,
		"magnet": true, "bitcoin": true, "ethereum": true,
	}
	
	// Detectar protocolo existente
	var detectedScheme string
	for proto := range supportedSchemes {
		if strings.HasPrefix(input, proto+"://") {
			detectedScheme = proto
			break
		}
		if strings.HasPrefix(input, proto+":") && proto != "http" && proto != "https" {
			detectedScheme = proto
			break
		}
	}
	
	// Se não tem protocolo, adicionar https://
	if detectedScheme == "" {
		input = "https://" + input
	}
	
	// Parse URL
	u, err := url.Parse(input)
	if err != nil {
		return "", fmt.Errorf("URL inválida: %v", err)
	}
	
	// Validar scheme
	if !supportedSchemes[u.Scheme] {
		return "", fmt.Errorf("protocolo não suportado: %s", u.Scheme)
	}
	
	// Validar host (não é necessário para file://, data:, mailto:, tel:, magnet:, bitcoin:, ethereum:)
	requiresHost := map[string]bool{
		"http": true, "https": true, "ftp": true, "ftps": true, "ws": true, "wss": true,
		"gopher": true, "gemini": true, "sftp": true, "ssh": true, "git": true, "svn": true,
		"rtmp": true, "rtmps": true, "rtsp": true, "mms": true, "news": true, "nntp": true,
		"telnet": true, "irc": true, "ircs": true, "xmpp": true, "sip": true, "sips": true,
		"ipfs": true, "ipns": true,
	}
	
	if requiresHost[u.Scheme] {
		if u.Host == "" {
			return "", fmt.Errorf("host inválido para protocolo %s", u.Scheme)
		}
		
		// Verificar se domínio está bloqueado
		if v.isDomainBlocked(u.Host) {
			return "", fmt.Errorf("domínio bloqueado: %s", u.Host)
		}
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
