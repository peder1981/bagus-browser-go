package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// Config representa as configurações do usuário
type Config struct {
	// Autenticação
	RequirePassword   bool   `json:"require_password"`
	PasswordHash      string `json:"password_hash"`      // Hash bcrypt da senha
	PasswordSalt      string `json:"password_salt"`      // Salt para derivação de chave
	
	// Sessões e Cookies
	PersistSessions   bool   `json:"persist_sessions"`   // Manter usuário logado em sites
	PersistCookies    bool   `json:"persist_cookies"`    // Salvar cookies entre sessões
	
	// Cache e Performance
	VideoCacheEnabled bool   `json:"video_cache_enabled"` // Cache de vídeos
	VideoCacheSize    int    `json:"video_cache_size"`    // Tamanho em MB (padrão: 500)
	
	// Privacidade (mantém configurações existentes)
	BlockThirdPartyCookies bool `json:"block_third_party_cookies"`
	BlockGeolocation       bool `json:"block_geolocation"`
	BlockNotifications     bool `json:"block_notifications"`
	BlockMedia             bool `json:"block_media"`
	BlockWebGL             bool `json:"block_webgl"`
	BlockWebAudio          bool `json:"block_webaudio"`
	DoNotTrack             bool `json:"do_not_track"`
}

// DefaultConfig retorna configurações padrão
func DefaultConfig() *Config {
	return &Config{
		// Autenticação desabilitada por padrão
		RequirePassword:   false,
		PasswordHash:      "",
		PasswordSalt:      "",
		
		// Sessões: manter usuário logado por padrão
		PersistSessions:   true,
		PersistCookies:    true,
		
		// Cache de vídeos habilitado
		VideoCacheEnabled: true,
		VideoCacheSize:    500, // 500 MB
		
		// Privacidade (valores seguros por padrão)
		BlockThirdPartyCookies: true,
		BlockGeolocation:       true,
		BlockNotifications:     false, // Permitir notificações
		BlockMedia:             false, // Permitir câmera/microfone (Google Meet, etc)
		BlockWebGL:             false, // Permitir WebGL (necessário para Meet)
		BlockWebAudio:          false, // Permitir WebAudio (necessário para Meet)
		DoNotTrack:             true,
	}
}

// LoadConfig carrega configurações do arquivo
func LoadConfig() (*Config, error) {
	configPath := getConfigPath()
	
	// Se não existe, criar com padrões
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		config := DefaultConfig()
		if err := config.Save(); err != nil {
			return nil, err
		}
		return config, nil
	}
	
	// Ler arquivo
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	
	// Descriptografar se necessário
	cm, err := NewCryptoManager("")
	if err != nil {
		return nil, err
	}
	
	decrypted, err := cm.Decrypt(string(data))
	if err != nil {
		// Se falhar, pode ser arquivo antigo não criptografado
		decrypted = data
	}
	
	// Parse JSON
	var config Config
	if err := json.Unmarshal(decrypted, &config); err != nil {
		return nil, err
	}
	
	return &config, nil
}

// Save salva configurações no arquivo
func (c *Config) Save() error {
	configPath := getConfigPath()
	
	// Criar diretório se não existe
	configDir := filepath.Dir(configPath)
	if err := os.MkdirAll(configDir, 0700); err != nil {
		return err
	}
	
	// Serializar para JSON
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	
	// Criptografar
	cm, err := NewCryptoManager("")
	if err != nil {
		return err
	}
	
	encryptedStr, err := cm.Encrypt(data)
	if err != nil {
		return err
	}
	
	encrypted := []byte(encryptedStr)
	
	// Salvar arquivo
	if err := ioutil.WriteFile(configPath, encrypted, 0600); err != nil {
		return err
	}
	
	log.Println("⚙️  Configurações salvas")
	return nil
}

// getConfigPath retorna caminho do arquivo de configuração
func getConfigPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(homeDir, ".config", "bagus-browser", "config.enc")
}

// SetPassword define senha mestre
func (c *Config) SetPassword(password string) error {
	// Gerar salt
	salt, err := generateSalt()
	if err != nil {
		return err
	}
	
	// Hash da senha com bcrypt
	hash, err := hashPassword(password)
	if err != nil {
		return err
	}
	
	c.RequirePassword = true
	c.PasswordHash = hash
	c.PasswordSalt = salt
	
	return c.Save()
}

// VerifyPassword verifica senha mestre
func (c *Config) VerifyPassword(password string) bool {
	if !c.RequirePassword || c.PasswordHash == "" {
		return true // Sem senha configurada
	}
	
	return verifyPassword(password, c.PasswordHash)
}

// RemovePassword remove senha mestre
func (c *Config) RemovePassword() error {
	c.RequirePassword = false
	c.PasswordHash = ""
	c.PasswordSalt = ""
	
	return c.Save()
}
