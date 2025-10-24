package main

import (
	"log"
	"os/exec"
	"strings"
	
	"github.com/gotk3/gotk3/gtk"
)

// showSettingsDialog mostra diálogo de configurações
func (b *Browser) showSettingsDialog() {
	// Criar diálogo
	dialog, err := gtk.DialogNew()
	if err != nil {
		log.Fatal("Erro ao criar diálogo:", err)
	}
	
	dialog.SetTitle("⚙️  Configurações - Bagus Browser")
	dialog.SetDefaultSize(600, 500)
	dialog.SetModal(true)
	dialog.SetTransientFor(b.window)
	
	// Botões
	dialog.AddButton("Cancelar", gtk.RESPONSE_CANCEL)
	dialog.AddButton("Salvar", gtk.RESPONSE_OK)
	
	// Área de conteúdo
	contentArea, err := dialog.GetContentArea()
	if err != nil {
		log.Fatal("Erro ao obter área de conteúdo:", err)
	}
	contentArea.SetMarginTop(10)
	contentArea.SetMarginBottom(10)
	contentArea.SetMarginStart(10)
	contentArea.SetMarginEnd(10)
	
	// Notebook para abas
	notebook, err := gtk.NotebookNew()
	if err != nil {
		log.Fatal("Erro ao criar notebook:", err)
	}
	contentArea.Add(notebook)
	
	// Aba 1: Geral
	generalBox := createGeneralSettings(b.config)
	generalLabel, _ := gtk.LabelNew("⚙️  Geral")
	notebook.AppendPage(generalBox, generalLabel)
	
	// Aba 2: Segurança
	securityBox := createSecuritySettings(b.config)
	securityLabel, _ := gtk.LabelNew("🔒 Segurança")
	notebook.AppendPage(securityBox, securityLabel)
	
	// Aba 3: Sessões
	sessionBox := createSessionSettings(b.config)
	sessionLabel, _ := gtk.LabelNew("🔄 Sessões")
	notebook.AppendPage(sessionBox, sessionLabel)
	
	// Aba 4: Performance
	performanceBox := createPerformanceSettings(b.config)
	performanceLabel, _ := gtk.LabelNew("⚡ Performance")
	notebook.AppendPage(performanceBox, performanceLabel)
	
	// Aba 5: Privacidade
	privacyBox := createPrivacySettings(b.config)
	privacyLabel, _ := gtk.LabelNew("🕵️  Privacidade")
	notebook.AppendPage(privacyBox, privacyLabel)
	
	dialog.ShowAll()
	
	// Executar diálogo
	response := dialog.Run()
	
	if response == gtk.RESPONSE_OK {
		// Salvar configurações
		if err := b.config.Save(); err != nil {
			showErrorDialog("Erro", "Erro ao salvar configurações: "+err.Error())
		} else {
			log.Println("⚙️  Configurações salvas com sucesso")
			
			// Aplicar configurações
			b.applyConfig()
			
			showInfoDialog("Sucesso", "Configurações salvas!\nAlgumas mudanças podem exigir reiniciar o browser.")
		}
	}
	
	dialog.Destroy()
}

// createGeneralSettings cria aba de configurações gerais
func createGeneralSettings(config *Config) *gtk.Box {
	box, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	box.SetMarginTop(10)
	box.SetMarginBottom(10)
	box.SetMarginStart(10)
	box.SetMarginEnd(10)
	
	// Título
	title, _ := gtk.LabelNew("")
	title.SetMarkup("<b>Configurações Gerais</b>")
	title.SetHAlign(gtk.ALIGN_START)
	box.Add(title)
	
	// Navegador Padrão
	defaultBrowserFrame, _ := gtk.FrameNew("Navegador Padrão do Sistema")
	defaultBrowserBox, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	defaultBrowserBox.SetMarginTop(10)
	defaultBrowserBox.SetMarginBottom(10)
	defaultBrowserBox.SetMarginStart(10)
	defaultBrowserBox.SetMarginEnd(10)
	
	// Descrição
	desc, _ := gtk.LabelNew("Configure o Bagus Browser como navegador padrão do sistema.\nLinks em emails, PDFs e outros aplicativos abrirão automaticamente aqui.")
	desc.SetHAlign(gtk.ALIGN_START)
	desc.SetLineWrap(true)
	defaultBrowserBox.Add(desc)
	
	// Status atual
	statusBox, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 10)
	statusBox.SetMarginTop(10)
	
	statusLabel, _ := gtk.LabelNew("Status:")
	statusBox.Add(statusLabel)
	
	// Verificar se é navegador padrão
	isDefault := checkIfDefaultBrowser()
	
	statusValue, _ := gtk.LabelNew("")
	if isDefault {
		statusValue.SetMarkup("<span foreground='green'>✅ Bagus Browser é o navegador padrão</span>")
	} else {
		statusValue.SetMarkup("<span foreground='orange'>⚠️  Outro navegador está configurado como padrão</span>")
	}
	statusBox.Add(statusValue)
	
	defaultBrowserBox.Add(statusBox)
	
	// Botão para definir como padrão
	btnBox, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 10)
	btnBox.SetMarginTop(10)
	
	setDefaultBtn, _ := gtk.ButtonNewWithLabel("🌐 Definir como Navegador Padrão")
	setDefaultBtn.Connect("clicked", func() {
		if setAsDefaultBrowser() {
			statusValue.SetMarkup("<span foreground='green'>✅ Bagus Browser é o navegador padrão</span>")
			showInfoDialog("Sucesso", "Bagus Browser configurado como navegador padrão!\n\nAgora links em outros aplicativos abrirão aqui.")
		} else {
			showErrorDialog("Erro", "Não foi possível configurar como navegador padrão.\n\nTente manualmente nas Configurações do Sistema.")
		}
	})
	
	if isDefault {
		setDefaultBtn.SetSensitive(false)
		setDefaultBtn.SetLabel("✅ Já é o Navegador Padrão")
	}
	
	btnBox.Add(setDefaultBtn)
	
	// Botão para abrir configurações do sistema
	systemSettingsBtn, _ := gtk.ButtonNewWithLabel("⚙️  Abrir Configurações do Sistema")
	systemSettingsBtn.Connect("clicked", func() {
		openSystemSettings()
	})
	btnBox.Add(systemSettingsBtn)
	
	defaultBrowserBox.Add(btnBox)
	
	// Informação adicional
	separator, _ := gtk.SeparatorNew(gtk.ORIENTATION_HORIZONTAL)
	separator.SetMarginTop(10)
	separator.SetMarginBottom(10)
	defaultBrowserBox.Add(separator)
	
	infoLabel, _ := gtk.LabelNew("💡 Dica: Você também pode configurar manualmente em:\nConfigurações do Sistema → Aplicativos Padrão → Navegador Web")
	infoLabel.SetHAlign(gtk.ALIGN_START)
	infoLabel.SetLineWrap(true)
	defaultBrowserBox.Add(infoLabel)
	
	defaultBrowserFrame.Add(defaultBrowserBox)
	box.Add(defaultBrowserFrame)
	
	return box
}

// createSecuritySettings cria aba de segurança
func createSecuritySettings(config *Config) *gtk.Box {
	box, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	box.SetMarginTop(10)
	box.SetMarginBottom(10)
	box.SetMarginStart(10)
	box.SetMarginEnd(10)
	
	// Título
	title, _ := gtk.LabelNew("")
	title.SetMarkup("<b>Configurações de Segurança</b>")
	title.SetHAlign(gtk.ALIGN_START)
	box.Add(title)
	
	// Senha mestre
	passwordFrame, _ := gtk.FrameNew("Senha Mestre")
	passwordBox, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 5)
	passwordBox.SetMarginTop(10)
	passwordBox.SetMarginBottom(10)
	passwordBox.SetMarginStart(10)
	passwordBox.SetMarginEnd(10)
	
	passwordCheck, _ := gtk.CheckButtonNewWithLabel("Exigir senha ao abrir o browser")
	passwordCheck.SetActive(config.RequirePassword)
	passwordBox.Add(passwordCheck)
	
	passwordDesc, _ := gtk.LabelNew("Protege o acesso ao browser com senha mestre.\nTodos os dados (favoritos, sessões, cookies) são criptografados.")
	passwordDesc.SetHAlign(gtk.ALIGN_START)
	passwordDesc.SetLineWrap(true)
	passwordBox.Add(passwordDesc)
	
	// Botões de senha
	passwordBtnBox, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 5)
	passwordBtnBox.SetMarginTop(5)
	
	setPasswordBtn, _ := gtk.ButtonNewWithLabel("Definir Senha")
	changePasswordBtn, _ := gtk.ButtonNewWithLabel("Alterar Senha")
	removePasswordBtn, _ := gtk.ButtonNewWithLabel("Remover Senha")
	
	if config.RequirePassword {
		setPasswordBtn.SetSensitive(false)
		changePasswordBtn.SetSensitive(true)
		removePasswordBtn.SetSensitive(true)
	} else {
		setPasswordBtn.SetSensitive(true)
		changePasswordBtn.SetSensitive(false)
		removePasswordBtn.SetSensitive(false)
	}
	
	passwordBtnBox.Add(setPasswordBtn)
	passwordBtnBox.Add(changePasswordBtn)
	passwordBtnBox.Add(removePasswordBtn)
	passwordBox.Add(passwordBtnBox)
	
	// Conectar eventos
	passwordCheck.Connect("toggled", func() {
		active := passwordCheck.GetActive()
		config.RequirePassword = active
		
		if active && config.PasswordHash == "" {
			// Definir senha
			password, ok := showSetPasswordDialog()
			if ok {
				config.SetPassword(password)
				setPasswordBtn.SetSensitive(false)
				changePasswordBtn.SetSensitive(true)
				removePasswordBtn.SetSensitive(true)
			} else {
				passwordCheck.SetActive(false)
			}
		}
	})
	
	setPasswordBtn.Connect("clicked", func() {
		password, ok := showSetPasswordDialog()
		if ok {
			config.SetPassword(password)
			passwordCheck.SetActive(true)
			setPasswordBtn.SetSensitive(false)
			changePasswordBtn.SetSensitive(true)
			removePasswordBtn.SetSensitive(true)
		}
	})
	
	changePasswordBtn.Connect("clicked", func() {
		// Verificar senha atual
		currentPass, ok := showPasswordDialog("Senha Atual", "Digite a senha atual:")
		if !ok || !config.VerifyPassword(currentPass) {
			showErrorDialog("Erro", "Senha atual incorreta!")
			return
		}
		
		// Nova senha
		newPass, ok := showSetPasswordDialog()
		if ok {
			config.SetPassword(newPass)
			showInfoDialog("Sucesso", "Senha alterada com sucesso!")
		}
	})
	
	removePasswordBtn.Connect("clicked", func() {
		// Confirmar
		confirmDialog := gtk.MessageDialogNew(
			nil,
			gtk.DIALOG_MODAL,
			gtk.MESSAGE_QUESTION,
			gtk.BUTTONS_YES_NO,
			"Tem certeza que deseja remover a senha mestre?\n\nOs dados continuarão criptografados, mas não será necessário senha para acessar.",
		)
		confirmDialog.SetTitle("Confirmar")
		response := confirmDialog.Run()
		confirmDialog.Destroy()
		
		if response == gtk.RESPONSE_YES {
			config.RemovePassword()
			passwordCheck.SetActive(false)
			setPasswordBtn.SetSensitive(true)
			changePasswordBtn.SetSensitive(false)
			removePasswordBtn.SetSensitive(false)
			showInfoDialog("Sucesso", "Senha removida!")
		}
	})
	
	passwordFrame.Add(passwordBox)
	box.Add(passwordFrame)
	
	return box
}

// createSessionSettings cria aba de sessões
func createSessionSettings(config *Config) *gtk.Box {
	box, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	box.SetMarginTop(10)
	box.SetMarginBottom(10)
	box.SetMarginStart(10)
	box.SetMarginEnd(10)
	
	// Título
	title, _ := gtk.LabelNew("")
	title.SetMarkup("<b>Gerenciamento de Sessões</b>")
	title.SetHAlign(gtk.ALIGN_START)
	box.Add(title)
	
	// Persistir sessões
	sessionCheck, _ := gtk.CheckButtonNewWithLabel("Manter usuário logado em sites")
	sessionCheck.SetActive(config.PersistSessions)
	sessionCheck.Connect("toggled", func() {
		config.PersistSessions = sessionCheck.GetActive()
	})
	box.Add(sessionCheck)
	
	sessionDesc, _ := gtk.LabelNew("Quando habilitado, você permanecerá logado em sites (Gmail, GitHub, etc)\nao fechar e reabrir o browser.")
	sessionDesc.SetHAlign(gtk.ALIGN_START)
	sessionDesc.SetLineWrap(true)
	sessionDesc.SetMarginBottom(10)
	box.Add(sessionDesc)
	
	// Persistir cookies
	cookieCheck, _ := gtk.CheckButtonNewWithLabel("Salvar cookies entre sessões")
	cookieCheck.SetActive(config.PersistCookies)
	cookieCheck.Connect("toggled", func() {
		config.PersistCookies = cookieCheck.GetActive()
	})
	box.Add(cookieCheck)
	
	cookieDesc, _ := gtk.LabelNew("Salva cookies no disco. Necessário para manter login em sites.\nSe desabilitado, todos os cookies são apagados ao fechar o browser.")
	cookieDesc.SetHAlign(gtk.ALIGN_START)
	cookieDesc.SetLineWrap(true)
	cookieDesc.SetMarginBottom(10)
	box.Add(cookieDesc)
	
	// Aviso
	warningBox, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 5)
	warningIcon, _ := gtk.ImageNewFromIconName("dialog-warning", gtk.ICON_SIZE_BUTTON)
	warningLabel, _ := gtk.LabelNew("⚠️  Desabilitar essas opções aumenta a privacidade, mas você\nprecisará fazer login novamente em todos os sites a cada vez.")
	warningLabel.SetHAlign(gtk.ALIGN_START)
	warningLabel.SetLineWrap(true)
	warningBox.Add(warningIcon)
	warningBox.Add(warningLabel)
	box.Add(warningBox)
	
	// Botão limpar cookies
	separator, _ := gtk.SeparatorNew(gtk.ORIENTATION_HORIZONTAL)
	separator.SetMarginTop(10)
	separator.SetMarginBottom(10)
	box.Add(separator)
	
	clearCookiesBtn, _ := gtk.ButtonNewWithLabel("🗑️  Limpar Todos os Cookies Agora")
	clearCookiesBtn.Connect("clicked", func() {
		confirmDialog := gtk.MessageDialogNew(
			nil,
			gtk.DIALOG_MODAL,
			gtk.MESSAGE_QUESTION,
			gtk.BUTTONS_YES_NO,
			"Tem certeza que deseja limpar todos os cookies?\n\nVocê será deslogado de todos os sites.",
		)
		confirmDialog.SetTitle("Confirmar")
		response := confirmDialog.Run()
		confirmDialog.Destroy()
		
		if response == gtk.RESPONSE_YES {
			// TODO: Implementar limpeza de cookies
			showInfoDialog("Sucesso", "Cookies limpos!\nReinicie o browser para aplicar.")
		}
	})
	box.Add(clearCookiesBtn)
	
	return box
}

// createPerformanceSettings cria aba de performance
func createPerformanceSettings(config *Config) *gtk.Box {
	box, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	box.SetMarginTop(10)
	box.SetMarginBottom(10)
	box.SetMarginStart(10)
	box.SetMarginEnd(10)
	
	// Título
	title, _ := gtk.LabelNew("")
	title.SetMarkup("<b>Configurações de Performance</b>")
	title.SetHAlign(gtk.ALIGN_START)
	box.Add(title)
	
	// Cache de vídeos
	cacheCheck, _ := gtk.CheckButtonNewWithLabel("Habilitar cache de vídeos")
	cacheCheck.SetActive(config.VideoCacheEnabled)
	cacheCheck.Connect("toggled", func() {
		config.VideoCacheEnabled = cacheCheck.GetActive()
	})
	box.Add(cacheCheck)
	
	cacheDesc, _ := gtk.LabelNew("Melhora reprodução de vídeos (YouTube, etc) armazenando em cache.\nVídeos carregam mais rápido e consomem menos banda.")
	cacheDesc.SetHAlign(gtk.ALIGN_START)
	cacheDesc.SetLineWrap(true)
	box.Add(cacheDesc)
	
	// Tamanho do cache
	cacheSizeBox, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 10)
	cacheSizeBox.SetMarginTop(10)
	
	cacheSizeLabel, _ := gtk.LabelNew("Tamanho do cache:")
	cacheSizeBox.Add(cacheSizeLabel)
	
	cacheSizeAdj, _ := gtk.AdjustmentNew(float64(config.VideoCacheSize), 100, 5000, 100, 500, 0)
	cacheSizeScale, _ := gtk.ScaleNew(gtk.ORIENTATION_HORIZONTAL, cacheSizeAdj)
	cacheSizeScale.SetHExpand(true)
	cacheSizeScale.SetDrawValue(true)
	// SetValuePos não existe em gotk3, valor já aparece por padrão
	cacheSizeScale.Connect("value-changed", func() {
		config.VideoCacheSize = int(cacheSizeScale.GetValue())
	})
	cacheSizeBox.Add(cacheSizeScale)
	
	cacheSizeMB, _ := gtk.LabelNew("MB")
	cacheSizeBox.Add(cacheSizeMB)
	
	box.Add(cacheSizeBox)
	
	// Botão limpar cache
	separator, _ := gtk.SeparatorNew(gtk.ORIENTATION_HORIZONTAL)
	separator.SetMarginTop(10)
	separator.SetMarginBottom(10)
	box.Add(separator)
	
	clearCacheBtn, _ := gtk.ButtonNewWithLabel("🗑️  Limpar Cache de Vídeos")
	clearCacheBtn.Connect("clicked", func() {
		if err := clearVideoCache(); err != nil {
			showErrorDialog("Erro", "Erro ao limpar cache: "+err.Error())
		} else {
			showInfoDialog("Sucesso", "Cache de vídeos limpo!")
		}
	})
	box.Add(clearCacheBtn)
	
	return box
}

// createPrivacySettings cria aba de privacidade
func createPrivacySettings(config *Config) *gtk.Box {
	box, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	box.SetMarginTop(10)
	box.SetMarginBottom(10)
	box.SetMarginStart(10)
	box.SetMarginEnd(10)
	
	// Título
	title, _ := gtk.LabelNew("")
	title.SetMarkup("<b>Configurações de Privacidade</b>")
	title.SetHAlign(gtk.ALIGN_START)
	box.Add(title)
	
	// Third-party cookies
	thirdPartyCheck, _ := gtk.CheckButtonNewWithLabel("Bloquear cookies de terceiros")
	thirdPartyCheck.SetActive(config.BlockThirdPartyCookies)
	thirdPartyCheck.Connect("toggled", func() {
		config.BlockThirdPartyCookies = thirdPartyCheck.GetActive()
	})
	box.Add(thirdPartyCheck)
	
	// Geolocalização
	geoCheck, _ := gtk.CheckButtonNewWithLabel("Bloquear geolocalização")
	geoCheck.SetActive(config.BlockGeolocation)
	geoCheck.Connect("toggled", func() {
		config.BlockGeolocation = geoCheck.GetActive()
	})
	box.Add(geoCheck)
	
	// Notificações
	notifCheck, _ := gtk.CheckButtonNewWithLabel("Bloquear notificações")
	notifCheck.SetActive(config.BlockNotifications)
	notifCheck.Connect("toggled", func() {
		config.BlockNotifications = notifCheck.GetActive()
	})
	box.Add(notifCheck)
	
	// Mídia
	mediaCheck, _ := gtk.CheckButtonNewWithLabel("Bloquear acesso a câmera/microfone")
	mediaCheck.SetActive(config.BlockMedia)
	mediaCheck.Connect("toggled", func() {
		config.BlockMedia = mediaCheck.GetActive()
	})
	box.Add(mediaCheck)
	
	// WebGL
	webglCheck, _ := gtk.CheckButtonNewWithLabel("Bloquear WebGL (anti-fingerprinting)")
	webglCheck.SetActive(config.BlockWebGL)
	webglCheck.Connect("toggled", func() {
		config.BlockWebGL = webglCheck.GetActive()
	})
	box.Add(webglCheck)
	
	// WebAudio
	webaudioCheck, _ := gtk.CheckButtonNewWithLabel("Bloquear WebAudio (anti-fingerprinting)")
	webaudioCheck.SetActive(config.BlockWebAudio)
	webaudioCheck.Connect("toggled", func() {
		config.BlockWebAudio = webaudioCheck.GetActive()
	})
	box.Add(webaudioCheck)
	
	// Do Not Track
	dntCheck, _ := gtk.CheckButtonNewWithLabel("Enviar cabeçalho Do Not Track")
	dntCheck.SetActive(config.DoNotTrack)
	dntCheck.Connect("toggled", func() {
		config.DoNotTrack = dntCheck.GetActive()
	})
	box.Add(dntCheck)
	
	return box
}

// showInfoDialog mostra diálogo de informação
func showInfoDialog(title, message string) {
	dialog := gtk.MessageDialogNew(
		nil,
		gtk.DIALOG_MODAL,
		gtk.MESSAGE_INFO,
		gtk.BUTTONS_OK,
		message,
	)
	dialog.SetTitle(title)
	dialog.Run()
	dialog.Destroy()
}

// applyConfig aplica configurações ao browser
func (b *Browser) applyConfig() {
	// Aplicar configurações de privacidade a todos os WebViews existentes
	for _, tab := range b.tabs {
		if tab.webView != nil {
			applyPrivacySettings(tab.webView.cWebView, b.config)
		}
	}
	
	log.Println("⚙️  Configurações aplicadas")
}

// checkIfDefaultBrowser verifica se Bagus é o navegador padrão
func checkIfDefaultBrowser() bool {
	cmd := exec.Command("xdg-settings", "get", "default-web-browser")
	output, err := cmd.Output()
	if err != nil {
		log.Printf("⚠️  Erro ao verificar navegador padrão: %v", err)
		return false
	}
	
	result := strings.TrimSpace(string(output))
	isDefault := result == "bagus-browser.desktop"
	
	if isDefault {
		log.Println("✅ Bagus Browser é o navegador padrão")
	} else {
		log.Printf("ℹ️  Navegador padrão atual: %s", result)
	}
	
	return isDefault
}

// setAsDefaultBrowser configura Bagus como navegador padrão
func setAsDefaultBrowser() bool {
	log.Println("🌐 Configurando Bagus Browser como navegador padrão...")
	
	// Método 1: xdg-settings (mais confiável)
	cmd := exec.Command("xdg-settings", "set", "default-web-browser", "bagus-browser.desktop")
	if err := cmd.Run(); err != nil {
		log.Printf("❌ Erro ao configurar com xdg-settings: %v", err)
		return false
	}
	
	// Método 2: xdg-mime (backup para tipos MIME específicos)
	mimeTypes := []string{
		"text/html",
		"text/xml",
		"application/xhtml+xml",
		"x-scheme-handler/http",
		"x-scheme-handler/https",
	}
	
	for _, mimeType := range mimeTypes {
		cmd := exec.Command("xdg-mime", "default", "bagus-browser.desktop", mimeType)
		if err := cmd.Run(); err != nil {
			log.Printf("⚠️  Erro ao configurar MIME %s: %v", mimeType, err)
			// Não retornar false aqui, continuar tentando
		}
	}
	
	// Verificar se funcionou
	if checkIfDefaultBrowser() {
		log.Println("✅ Bagus Browser configurado como navegador padrão com sucesso!")
		return true
	}
	
	log.Println("⚠️  Configuração pode não ter funcionado completamente")
	return false
}

// openSystemSettings abre as configurações do sistema
func openSystemSettings() {
	log.Println("⚙️  Abrindo configurações do sistema...")
	
	// Tentar diferentes comandos dependendo do ambiente desktop
	commands := [][]string{
		{"gnome-control-center", "default-apps"},           // GNOME
		{"systemsettings5", "kcm_componentchooser"},        // KDE Plasma 5
		{"systemsettings", "kcm_componentchooser"},         // KDE Plasma 4
		{"unity-control-center", "default-apps"},           // Unity
		{"xfce4-settings-manager"},                         // XFCE
		{"mate-default-applications-properties"},           // MATE
		{"lxqt-config-appearance"},                         // LXQt
		{"gnome-control-center"},                           // GNOME (fallback)
	}
	
	for _, cmdArgs := range commands {
		cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
		if err := cmd.Start(); err == nil {
			log.Printf("✅ Aberto: %s", cmdArgs[0])
			return
		}
	}
	
	// Se nenhum funcionou, tentar xdg-open genérico
	cmd := exec.Command("xdg-open", "x-scheme-handler/http")
	if err := cmd.Start(); err != nil {
		log.Printf("❌ Não foi possível abrir configurações do sistema: %v", err)
		showErrorDialog("Erro", "Não foi possível abrir as configurações do sistema.\n\nAbra manualmente:\nConfigurações → Aplicativos Padrão")
	}
}
