package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Config representa a configuração do browser
type Config struct {
	Default  DefaultConfig  `json:"default"`
	Settings SettingsConfig `json:"settings"`
}

// DefaultConfig configurações padrão
type DefaultConfig struct {
	URL string `json:"url"`
}

// SettingsConfig todas as configurações de segurança e comportamento
type SettingsConfig struct {
	LocalStorageEnabled                 bool `json:"LocalStorageEnabled"`
	XSSAuditingEnabled                  bool `json:"XSSAuditingEnabled"`
	HyperlinkAuditingEnabled            bool `json:"HyperlinkAuditingEnabled"`
	FullScreenSupportEnabled            bool `json:"FullScreenSupportEnabled"`
	JavascriptCanAccessClipboard        bool `json:"JavascriptCanAccessClipboard"`
	PluginsEnabled                      bool `json:"PluginsEnabled"`
	AutoLoadImages                      bool `json:"AutoLoadImages"`
	JavascriptEnabled                   bool `json:"JavascriptEnabled"`
	JavascriptCanOpenWindows            bool `json:"JavascriptCanOpenWindows"`
	LinksIncludedInFocusChain           bool `json:"LinksIncludedInFocusChain"`
	LocalContentCanAccessRemoteUrls     bool `json:"LocalContentCanAccessRemoteUrls"`
	SpatialNavigationEnabled            bool `json:"SpatialNavigationEnabled"`
	LocalContentCanAccessFileUrls       bool `json:"LocalContentCanAccessFileUrls"`
	ScrollAnimatorEnabled               bool `json:"ScrollAnimatorEnabled"`
	ErrorPageEnabled                    bool `json:"ErrorPageEnabled"`
	ScreenCaptureEnabled                bool `json:"ScreenCaptureEnabled"`
	WebGLEnabled                        bool `json:"WebGLEnabled"`
	Accelerated2dCanvasEnabled          bool `json:"Accelerated2dCanvasEnabled"`
	AutoLoadIconsForPage                bool `json:"AutoLoadIconsForPage"`
	TouchIconsEnabled                   bool `json:"TouchIconsEnabled"`
	FocusOnNavigationEnabled            bool `json:"FocusOnNavigationEnabled"`
	PrintElementBackgrounds             bool `json:"PrintElementBackgrounds"`
	AllowRunningInsecureContent         bool `json:"AllowRunningInsecureContent"`
	AllowGeolocationOnInsecureOrigins   bool `json:"AllowGeolocationOnInsecureOrigins"`
	AllowWindowActivationFromJavaScript bool `json:"AllowWindowActivationFromJavaScript"`
	ShowScrollBars                      bool `json:"ShowScrollBars"`
	PlaybackRequiresUserGesture         bool `json:"PlaybackRequiresUserGesture"`
	JavascriptCanPaste                  bool `json:"JavascriptCanPaste"`
	WebRTCPublicInterfacesOnly          bool `json:"WebRTCPublicInterfacesOnly"`
	DnsPrefetchEnabled                  bool `json:"DnsPrefetchEnabled"`
	PdfViewerEnabled                    bool `json:"PdfViewerEnabled"`
	NavigateOnDropEnabled               bool `json:"NavigateOnDropEnabled"`
	ReadingFromCanvasEnabled            bool `json:"ReadingFromCanvasEnabled"`
	ForceDarkMode                       bool `json:"ForceDarkMode"`
	PrintHeaderAndFooter                bool `json:"PrintHeaderAndFooter"`
	PreferCSSMarginsForPrinting         bool `json:"PreferCSSMarginsForPrinting"`
	TouchEventsApiEnabled               bool `json:"TouchEventsApiEnabled"`
}

// Load carrega configuração de um arquivo
func Load(path string) (*Config, error) {
	// Valida que o arquivo existe
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("arquivo de configuração não encontrado: %s", path)
	}

	// Valida tamanho do arquivo (max 1MB)
	info, err := os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("erro ao verificar arquivo: %w", err)
	}
	if info.Size() > 1024*1024 {
		return nil, fmt.Errorf("arquivo de configuração muito grande")
	}

	// Lê arquivo
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler configuração: %w", err)
	}

	// Parse JSON
	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("JSON inválido: %w", err)
	}

	return &config, nil
}

// Save salva configuração em arquivo
func (c *Config) Save(path string) error {
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Errorf("erro ao serializar configuração: %w", err)
	}

	// Cria diretório se não existir
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0700); err != nil {
		return fmt.Errorf("erro ao criar diretório: %w", err)
	}

	// Escreve arquivo com permissões restritas
	if err := os.WriteFile(path, data, 0600); err != nil {
		return fmt.Errorf("erro ao salvar configuração: %w", err)
	}

	return nil
}

// Default retorna configuração padrão
func Default() *Config {
	return &Config{
		Default: DefaultConfig{
			URL: "https://duckduckgo.com/",
		},
		Settings: SettingsConfig{
			LocalStorageEnabled:                 true,
			XSSAuditingEnabled:                  true,
			HyperlinkAuditingEnabled:            false,
			FullScreenSupportEnabled:            true,
			JavascriptCanAccessClipboard:        false,
			PluginsEnabled:                      false,
			AutoLoadImages:                      true,
			JavascriptEnabled:                   true,
			JavascriptCanOpenWindows:            false,
			LinksIncludedInFocusChain:           true,
			LocalContentCanAccessRemoteUrls:     false,
			SpatialNavigationEnabled:            false,
			LocalContentCanAccessFileUrls:       false,
			ScrollAnimatorEnabled:               true,
			ErrorPageEnabled:                    true,
			ScreenCaptureEnabled:                false,
			WebGLEnabled:                        true,
			Accelerated2dCanvasEnabled:          true,
			AutoLoadIconsForPage:                true,
			TouchIconsEnabled:                   true,
			FocusOnNavigationEnabled:            true,
			PrintElementBackgrounds:             true,
			AllowRunningInsecureContent:         false,
			AllowGeolocationOnInsecureOrigins:   false,
			AllowWindowActivationFromJavaScript: false,
			ShowScrollBars:                      true,
			PlaybackRequiresUserGesture:         true,
			JavascriptCanPaste:                  false,
			WebRTCPublicInterfacesOnly:          true,
			DnsPrefetchEnabled:                  true,
			PdfViewerEnabled:                    true,
			NavigateOnDropEnabled:               false,
			ReadingFromCanvasEnabled:            false,
			ForceDarkMode:                       false,
			PrintHeaderAndFooter:                true,
			PreferCSSMarginsForPrinting:         true,
			TouchEventsApiEnabled:               true,
		},
	}
}
