package security

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

var (
	// Regex para validar username
	usernameRegex = regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)
)

// ValidateUsername valida username para prevenir path traversal
func ValidateUsername(username string) error {
	username = strings.TrimSpace(username)

	if len(username) < 3 || len(username) > 32 {
		return fmt.Errorf("username deve ter entre 3 e 32 caracteres")
	}

	if !usernameRegex.MatchString(username) {
		return fmt.Errorf("username deve conter apenas letras, números, _ ou -")
	}

	// Previne path traversal
	if strings.Contains(username, "..") ||
		strings.Contains(username, "/") ||
		strings.Contains(username, "\\") {
		return fmt.Errorf("username contém caracteres inválidos")
	}

	return nil
}

// ValidateURL valida URL para prevenir ataques
func ValidateURL(rawURL string) error {
	rawURL = strings.TrimSpace(rawURL)

	if rawURL == "" {
		return fmt.Errorf("URL vazia")
	}

	// Limita tamanho
	if len(rawURL) > 2048 {
		return fmt.Errorf("URL muito longa")
	}

	// Parse URL
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return fmt.Errorf("URL inválida: %w", err)
	}

	// Valida esquema
	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return fmt.Errorf("protocolo não permitido: %s", parsedURL.Scheme)
	}

	// Previne URLs com credenciais
	if parsedURL.User != nil {
		return fmt.Errorf("URLs com credenciais não são permitidas")
	}

	return nil
}

// SanitizeURL adiciona protocolo se necessário
func SanitizeURL(rawURL string) string {
	rawURL = strings.TrimSpace(rawURL)

	if rawURL == "" {
		return ""
	}

	// Adiciona https:// se não tiver protocolo
	if !strings.HasPrefix(rawURL, "http://") &&
		!strings.HasPrefix(rawURL, "https://") {
		rawURL = "https://" + rawURL
	}

	return rawURL
}

// IsValidURL verifica se URL é válida (versão rápida)
func IsValidURL(rawURL string) bool {
	return ValidateURL(rawURL) == nil
}

// ExtractDomain extrai domínio de uma URL
func ExtractDomain(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	return strings.ToLower(parsedURL.Hostname()), nil
}
