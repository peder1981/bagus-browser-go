package browser

import (
	"log"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	"github.com/peder1981/bagus-browser-go/pkg/ipc"
)

// ControlWindow representa a janela de controle (GTK)
type ControlWindow struct {
	window    *gtk.Window
	urlEntry  *gtk.Entry
	backBtn   *gtk.Button
	forwardBtn *gtk.Button
	reloadBtn *gtk.Button
	stopBtn   *gtk.Button
	spinner   *gtk.Spinner
	channel   *ipc.Channel
}

// NewControlWindow cria uma nova janela de controle
func NewControlWindow(channel *ipc.Channel) (*ControlWindow, error) {
	gtk.Init(nil)

	// Criar janela
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		return nil, err
	}

	cw := &ControlWindow{
		window:  win,
		channel: channel,
	}

	// Configurar janela
	win.SetTitle("Bagus Browser")
	win.SetDefaultSize(800, 60)
	win.SetResizable(false)
	win.SetPosition(gtk.WIN_POS_CENTER)

	// Criar UI
	if err := cw.buildUI(); err != nil {
		return nil, err
	}

	// Conectar eventos
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	return cw, nil
}

// buildUI constrói a interface da janela de controle
func (cw *ControlWindow) buildUI() error {
	// Container principal
	box, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 5)
	if err != nil {
		return err
	}
	box.SetMarginTop(10)
	box.SetMarginBottom(10)
	box.SetMarginStart(10)
	box.SetMarginEnd(10)

	// Botão Voltar
	backBtn, err := gtk.ButtonNewWithLabel("←")
	if err != nil {
		return err
	}
	backBtn.SetTooltipText("Voltar (Alt+←)")
	backBtn.Connect("clicked", func() {
		cw.channel.SendToContent(ipc.Message{Type: ipc.MsgBack})
	})
	cw.backBtn = backBtn
	box.PackStart(backBtn, false, false, 0)

	// Botão Avançar
	forwardBtn, err := gtk.ButtonNewWithLabel("→")
	if err != nil {
		return err
	}
	forwardBtn.SetTooltipText("Avançar (Alt+→)")
	forwardBtn.Connect("clicked", func() {
		cw.channel.SendToContent(ipc.Message{Type: ipc.MsgForward})
	})
	cw.forwardBtn = forwardBtn
	box.PackStart(forwardBtn, false, false, 0)

	// Botão Recarregar
	reloadBtn, err := gtk.ButtonNewWithLabel("⟳")
	if err != nil {
		return err
	}
	reloadBtn.SetTooltipText("Recarregar (F5)")
	reloadBtn.Connect("clicked", func() {
		cw.channel.SendToContent(ipc.Message{Type: ipc.MsgReload})
	})
	cw.reloadBtn = reloadBtn
	box.PackStart(reloadBtn, false, false, 0)

	// Botão Parar
	stopBtn, err := gtk.ButtonNewWithLabel("✕")
	if err != nil {
		return err
	}
	stopBtn.SetTooltipText("Parar")
	stopBtn.Connect("clicked", func() {
		cw.channel.SendToContent(ipc.Message{Type: ipc.MsgStop})
	})
	cw.stopBtn = stopBtn
	box.PackStart(stopBtn, false, false, 0)

	// Campo de URL
	urlEntry, err := gtk.EntryNew()
	if err != nil {
		return err
	}
	urlEntry.SetPlaceholderText("Digite uma URL ou termo de busca...")
	urlEntry.Connect("activate", func() {
		text, _ := urlEntry.GetText()
		cw.navigate(text)
	})
	cw.urlEntry = urlEntry
	box.PackStart(urlEntry, true, true, 0)

	// Botão Ir
	goBtn, err := gtk.ButtonNewWithLabel("Ir")
	if err != nil {
		return err
	}
	goBtn.Connect("clicked", func() {
		text, _ := urlEntry.GetText()
		cw.navigate(text)
	})
	box.PackStart(goBtn, false, false, 0)

	// Spinner de carregamento
	spinner, err := gtk.SpinnerNew()
	if err != nil {
		return err
	}
	cw.spinner = spinner
	box.PackStart(spinner, false, false, 0)

	// Adicionar box à janela
	cw.window.Add(box)

	// Conectar atalhos de teclado
	cw.connectKeyboardShortcuts()

	return nil
}

// connectKeyboardShortcuts conecta atalhos de teclado
func (cw *ControlWindow) connectKeyboardShortcuts() {
	cw.window.Connect("key-press-event", func(win *gtk.Window, ev *gdk.Event) bool {
		keyEvent := gdk.EventKeyNewFromEvent(ev)
		keyVal := keyEvent.KeyVal()
		state := keyEvent.State()

		// Alt+← - Voltar (MOD1 = Alt = 1<<3 = 8)
		if keyVal == gdk.KEY_Left && gdk.ModifierType(state)&gdk.ModifierType(8) != 0 {
			cw.channel.SendToContent(ipc.Message{Type: ipc.MsgBack})
			return true
		}

		// Alt+→ - Avançar
		if keyVal == gdk.KEY_Right && gdk.ModifierType(state)&gdk.ModifierType(8) != 0 {
			cw.channel.SendToContent(ipc.Message{Type: ipc.MsgForward})
			return true
		}

		// F5 - Recarregar
		if keyVal == gdk.KEY_F5 {
			cw.channel.SendToContent(ipc.Message{Type: ipc.MsgReload})
			return true
		}

		// Ctrl+L - Focar URL (CONTROL = 1<<2 = 4)
		if keyVal == gdk.KEY_l && gdk.ModifierType(state)&gdk.ModifierType(4) != 0 {
			cw.urlEntry.GrabFocus()
			cw.urlEntry.SelectRegion(0, -1)
			return true
		}

		return false
	})
}

// navigate processa e navega para URL
func (cw *ControlWindow) navigate(input string) {
	if input == "" {
		return
	}

	// Processar input (adicionar https:// se necessário, etc)
	url := processURL(input)

	log.Printf("Navegando para: %s", url)

	// Enviar comando de navegação
	cw.channel.SendToContent(ipc.Message{
		Type:    ipc.MsgNavigate,
		Payload: url,
	})
}

// processURL processa o input do usuário
func processURL(input string) string {
	// Se contém espaço, é uma busca
	if containsSpace(input) {
		return "https://duckduckgo.com/?q=" + input
	}

	// Se não tem protocolo, adicionar https://
	if !hasProtocol(input) {
		return "https://" + input
	}

	return input
}

func containsSpace(s string) bool {
	for _, c := range s {
		if c == ' ' {
			return true
		}
	}
	return false
}

func hasProtocol(s string) bool {
	return len(s) > 7 && (s[:7] == "http://" || s[:8] == "https://")
}

// UpdateURL atualiza o campo de URL
func (cw *ControlWindow) UpdateURL(url string) {
	cw.urlEntry.SetText(url)
}

// UpdateTitle atualiza o título da janela
func (cw *ControlWindow) UpdateTitle(title string) {
	if title != "" {
		cw.window.SetTitle(title + " - Bagus Browser")
	} else {
		cw.window.SetTitle("Bagus Browser")
	}
}

// SetLoading define o estado de carregamento
func (cw *ControlWindow) SetLoading(loading bool) {
	if loading {
		cw.spinner.Start()
	} else {
		cw.spinner.Stop()
	}
}

// Run inicia a janela de controle
func (cw *ControlWindow) Run() error {
	cw.window.ShowAll()
	gtk.Main()
	return nil
}

// Close fecha a janela
func (cw *ControlWindow) Close() {
	cw.window.Close()
}
