package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"time"
	
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

// DownloadItem representa um download em andamento
type DownloadItem struct {
	ID            string
	URL           string
	Filename      string
	Destination   string
	BytesReceived uint64
	TotalBytes    uint64
	Progress      float64
	Status        string // "downloading", "completed", "failed", "cancelled"
	StartTime     time.Time
	EndTime       time.Time
	Error         error
	
	// UI elements
	listBoxRow    *gtk.ListBoxRow
	progressBar   *gtk.ProgressBar
	statusLabel   *gtk.Label
	openBtn       *gtk.Button
	cancelBtn     *gtk.Button
}

// DownloadManager gerencia downloads
type DownloadManager struct {
	downloadPath string
	downloads    map[string]*DownloadItem
	mu           sync.RWMutex
	
	// UI
	window      *gtk.Window
	listBox     *gtk.ListBox
	visible     bool
}

// NewDownloadManager cria um novo gerenciador de downloads
func NewDownloadManager() (*DownloadManager, error) {
	// Diretório padrão de downloads
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	
	downloadPath := filepath.Join(homeDir, "Downloads")
	
	// Criar diretório se não existir
	if err := os.MkdirAll(downloadPath, 0755); err != nil {
		return nil, err
	}
	
	dm := &DownloadManager{
		downloadPath: downloadPath,
		downloads:    make(map[string]*DownloadItem),
	}
	
	// Criar janela de downloads
	if err := dm.createDownloadWindow(); err != nil {
		log.Printf("⚠️  Erro ao criar janela de downloads: %v", err)
	}
	
	return dm, nil
}

// createDownloadWindow cria a janela de gerenciamento de downloads
func (dm *DownloadManager) createDownloadWindow() error {
	window, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		return err
	}
	
	window.SetTitle("Downloads - Bagus Browser")
	window.SetDefaultSize(600, 400)
	window.SetPosition(gtk.WIN_POS_CENTER)
	
	// Não destruir ao fechar, apenas ocultar
	window.Connect("delete-event", func() bool {
		window.Hide()
		dm.visible = false
		return true // Previne destruição
	})
	
	// Container principal
	vbox, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	
	// Header bar
	headerBar, _ := gtk.HeaderBarNew()
	headerBar.SetTitle("Downloads")
	headerBar.SetSubtitle(dm.downloadPath)
	headerBar.SetShowCloseButton(true)
	
	// Botão para abrir pasta
	openFolderBtn, _ := gtk.ButtonNewWithLabel("📁 Abrir Pasta")
	openFolderBtn.Connect("clicked", func() {
		dm.OpenDownloadFolder()
	})
	headerBar.PackStart(openFolderBtn)
	
	// Botão para limpar concluídos
	clearBtn, _ := gtk.ButtonNewWithLabel("🗑️  Limpar Concluídos")
	clearBtn.Connect("clicked", func() {
		dm.ClearCompleted()
	})
	headerBar.PackEnd(clearBtn)
	
	window.SetTitlebar(headerBar)
	
	// Scrolled window para lista
	scrolled, _ := gtk.ScrolledWindowNew(nil, nil)
	scrolled.SetPolicy(gtk.POLICY_NEVER, gtk.POLICY_AUTOMATIC)
	
	// ListBox para downloads
	listBox, _ := gtk.ListBoxNew()
	listBox.SetSelectionMode(gtk.SELECTION_NONE)
	scrolled.Add(listBox)
	
	vbox.PackStart(scrolled, true, true, 0)
	
	// Label quando vazio
	emptyLabel, _ := gtk.LabelNew("")
	emptyLabel.SetMarkup("<span size='large'>📥 Nenhum download</span>\n<span size='small'>Os downloads aparecerão aqui</span>")
	emptyLabel.SetVAlign(gtk.ALIGN_CENTER)
	emptyLabel.SetMarginTop(50)
	emptyLabel.SetMarginBottom(50)
	listBox.SetPlaceholder(emptyLabel)
	
	window.Add(vbox)
	
	dm.window = window
	dm.listBox = listBox
	
	return nil
}

// ShowDownloadWindow mostra a janela de downloads
func (dm *DownloadManager) ShowDownloadWindow() {
	if dm.window != nil {
		dm.window.ShowAll()
		dm.window.Present()
		dm.visible = true
	}
}

// IsVisible retorna se a janela está visível
func (dm *DownloadManager) IsVisible() bool {
	return dm.visible
}

// GetDownloadPath retorna o caminho de downloads
func (dm *DownloadManager) GetDownloadPath() string {
	return dm.downloadPath
}

// SetDownloadPath define um novo caminho de downloads
func (dm *DownloadManager) SetDownloadPath(path string) error {
	// Verificar se diretório existe
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// Criar se não existir
		if err := os.MkdirAll(path, 0755); err != nil {
			return err
		}
	}
	
	dm.downloadPath = path
	log.Printf("📁 Pasta de downloads: %s", path)
	
	// Atualizar subtitle da janela
	if dm.window != nil {
		glib.IdleAdd(func() bool {
			headerBar, _ := dm.window.GetTitlebar()
			if hb, ok := headerBar.(*gtk.HeaderBar); ok {
				hb.SetSubtitle(path)
			}
			return false
		})
	}
	
	return nil
}

// GetUniqueFilename retorna um nome de arquivo único
func (dm *DownloadManager) GetUniqueFilename(filename string) string {
	fullPath := filepath.Join(dm.downloadPath, filename)
	
	// Se arquivo não existe, retornar como está
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return fullPath
	}
	
	// Arquivo existe, adicionar número
	ext := filepath.Ext(filename)
	name := filename[:len(filename)-len(ext)]
	
	for i := 1; i < 1000; i++ {
		newFilename := fmt.Sprintf("%s (%d)%s", name, i, ext)
		fullPath = filepath.Join(dm.downloadPath, newFilename)
		
		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			return fullPath
		}
	}
	
	// Fallback
	return filepath.Join(dm.downloadPath, fmt.Sprintf("%s_%d%s", name, os.Getpid(), ext))
}

// AddDownload adiciona um novo download
func (dm *DownloadManager) AddDownload(id, url, filename string) *DownloadItem {
	dm.mu.Lock()
	defer dm.mu.Unlock()
	
	destination := dm.GetUniqueFilename(filename)
	
	item := &DownloadItem{
		ID:          id,
		URL:         url,
		Filename:    filepath.Base(destination),
		Destination: destination,
		Status:      "downloading",
		StartTime:   time.Now(),
	}
	
	dm.downloads[id] = item
	
	// Criar UI para o download
	glib.IdleAdd(func() bool {
		dm.createDownloadRow(item)
		return false
	})
	
	// Mostrar janela se não estiver visível
	if !dm.visible {
		glib.IdleAdd(func() bool {
			dm.ShowDownloadWindow()
			return false
		})
	}
	
	log.Printf("📥 Download iniciado: %s", filename)
	
	return item
}

// createDownloadRow cria a linha de UI para um download
func (dm *DownloadManager) createDownloadRow(item *DownloadItem) {
	if dm.listBox == nil {
		return
	}
	
	// Container principal
	box, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 5)
	box.SetMarginTop(10)
	box.SetMarginBottom(10)
	box.SetMarginStart(10)
	box.SetMarginEnd(10)
	
	// Header: nome do arquivo e botões
	headerBox, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 10)
	
	// Ícone
	icon, _ := gtk.LabelNew("📄")
	headerBox.PackStart(icon, false, false, 0)
	
	// Nome do arquivo
	filenameLabel, _ := gtk.LabelNew(item.Filename)
	filenameLabel.SetEllipsize(3) // PANGO_ELLIPSIZE_END
	filenameLabel.SetXAlign(0)
	headerBox.PackStart(filenameLabel, true, true, 0)
	
	// Botão abrir
	openBtn, _ := gtk.ButtonNewWithLabel("📂")
	openBtn.SetTooltipText("Abrir arquivo")
	openBtn.SetSensitive(false)
	openBtn.Connect("clicked", func() {
		dm.OpenFile(item.Destination)
	})
	headerBox.PackEnd(openBtn, false, false, 0)
	
	// Botão cancelar
	cancelBtn, _ := gtk.ButtonNewWithLabel("❌")
	cancelBtn.SetTooltipText("Cancelar download")
	cancelBtn.Connect("clicked", func() {
		dm.CancelDownload(item.ID)
	})
	headerBox.PackEnd(cancelBtn, false, false, 0)
	
	box.PackStart(headerBox, false, false, 0)
	
	// Progress bar
	progressBar, _ := gtk.ProgressBarNew()
	progressBar.SetShowText(true)
	progressBar.SetText("Iniciando...")
	box.PackStart(progressBar, false, false, 0)
	
	// Status label
	statusLabel, _ := gtk.LabelNew("")
	statusLabel.SetXAlign(0)
	statusLabel.SetMarkup("<small>Aguardando...</small>")
	box.PackStart(statusLabel, false, false, 0)
	
	// Criar row
	row, _ := gtk.ListBoxRowNew()
	row.Add(box)
	row.ShowAll()
	
	// Adicionar à lista
	dm.listBox.Add(row)
	
	// Salvar referências
	item.listBoxRow = row
	item.progressBar = progressBar
	item.statusLabel = statusLabel
	item.openBtn = openBtn
	item.cancelBtn = cancelBtn
}

// UpdateProgress atualiza o progresso de um download
func (dm *DownloadManager) UpdateProgress(id string, received, total uint64) {
	dm.mu.RLock()
	item, exists := dm.downloads[id]
	dm.mu.RUnlock()
	
	if !exists {
		return
	}
	
	item.BytesReceived = received
	item.TotalBytes = total
	
	if total > 0 {
		item.Progress = float64(received) / float64(total)
	} else {
		item.Progress = 0
	}
	
	// Atualizar UI
	glib.IdleAdd(func() bool {
		if item.progressBar != nil {
			item.progressBar.SetFraction(item.Progress)
			
			if total > 0 {
				item.progressBar.SetText(fmt.Sprintf("%.1f%% - %s de %s",
					item.Progress*100,
					formatBytes(received),
					formatBytes(total)))
			} else {
				item.progressBar.SetText(fmt.Sprintf("%s baixados", formatBytes(received)))
			}
		}
		
		if item.statusLabel != nil {
			elapsed := time.Since(item.StartTime)
			speed := float64(received) / elapsed.Seconds()
			
			var eta string
			if total > 0 && speed > 0 {
				remaining := float64(total-received) / speed
				eta = fmt.Sprintf(" - %s restantes", formatDuration(time.Duration(remaining)*time.Second))
			}
			
			item.statusLabel.SetMarkup(fmt.Sprintf("<small>%s/s%s</small>",
				formatBytes(uint64(speed)), eta))
		}
		
		return false
	})
}

// CompleteDownload marca um download como concluído
func (dm *DownloadManager) CompleteDownload(id string) {
	dm.mu.RLock()
	item, exists := dm.downloads[id]
	dm.mu.RUnlock()
	
	if !exists {
		return
	}
	
	item.Status = "completed"
	item.EndTime = time.Now()
	item.Progress = 1.0
	
	log.Printf("✅ Download concluído: %s", item.Filename)
	
	// Atualizar UI
	glib.IdleAdd(func() bool {
		if item.progressBar != nil {
			item.progressBar.SetFraction(1.0)
			item.progressBar.SetText("✅ Concluído")
		}
		
		if item.statusLabel != nil {
			duration := item.EndTime.Sub(item.StartTime)
			item.statusLabel.SetMarkup(fmt.Sprintf("<small>Concluído em %s - %s</small>",
				formatDuration(duration),
				formatBytes(item.BytesReceived)))
		}
		
		// Habilitar botão abrir, desabilitar cancelar
		if item.openBtn != nil {
			item.openBtn.SetSensitive(true)
		}
		if item.cancelBtn != nil {
			item.cancelBtn.SetSensitive(false)
		}
		
		return false
	})
	
	// Notificação do sistema
	dm.sendNotification("Download Concluído", item.Filename)
}

// FailDownload marca um download como falho
func (dm *DownloadManager) FailDownload(id string, err error) {
	dm.mu.RLock()
	item, exists := dm.downloads[id]
	dm.mu.RUnlock()
	
	if !exists {
		return
	}
	
	item.Status = "failed"
	item.Error = err
	item.EndTime = time.Now()
	
	log.Printf("❌ Download falhou: %s - %v", item.Filename, err)
	
	// Atualizar UI
	glib.IdleAdd(func() bool {
		if item.progressBar != nil {
			item.progressBar.SetText("❌ Falhou")
		}
		
		if item.statusLabel != nil {
			item.statusLabel.SetMarkup(fmt.Sprintf("<small><span foreground='red'>Erro: %v</span></small>", err))
		}
		
		// Desabilitar botão cancelar
		if item.cancelBtn != nil {
			item.cancelBtn.SetSensitive(false)
		}
		
		return false
	})
}

// CancelDownload cancela um download
func (dm *DownloadManager) CancelDownload(id string) {
	dm.mu.RLock()
	item, exists := dm.downloads[id]
	dm.mu.RUnlock()
	
	if !exists || item.Status != "downloading" {
		return
	}
	
	item.Status = "cancelled"
	item.EndTime = time.Now()
	
	log.Printf("🚫 Download cancelado: %s", item.Filename)
	
	// Remover arquivo parcial
	if _, err := os.Stat(item.Destination); err == nil {
		os.Remove(item.Destination)
	}
	
	// Atualizar UI
	glib.IdleAdd(func() bool {
		if item.progressBar != nil {
			item.progressBar.SetText("🚫 Cancelado")
		}
		
		if item.statusLabel != nil {
			item.statusLabel.SetMarkup("<small>Cancelado pelo usuário</small>")
		}
		
		// Remover da lista
		if item.listBoxRow != nil {
			dm.listBox.Remove(item.listBoxRow)
		}
		
		return false
	})
	
	// Remover do mapa
	dm.mu.Lock()
	delete(dm.downloads, id)
	dm.mu.Unlock()
}

// ClearCompleted remove downloads concluídos da lista
func (dm *DownloadManager) ClearCompleted() {
	dm.mu.Lock()
	defer dm.mu.Unlock()
	
	for id, item := range dm.downloads {
		if item.Status == "completed" {
			glib.IdleAdd(func() bool {
				if item.listBoxRow != nil {
					dm.listBox.Remove(item.listBoxRow)
				}
				return false
			})
			
			delete(dm.downloads, id)
		}
	}
	
	log.Println("🗑️  Downloads concluídos removidos")
}

// OpenFile abre um arquivo baixado
func (dm *DownloadManager) OpenFile(path string) {
	cmd := exec.Command("xdg-open", path)
	if err := cmd.Start(); err != nil {
		log.Printf("❌ Erro ao abrir arquivo: %v", err)
		showErrorDialog("Erro", fmt.Sprintf("Não foi possível abrir o arquivo:\n%v", err))
	}
}

// OpenDownloadFolder abre a pasta de downloads
func (dm *DownloadManager) OpenDownloadFolder() {
	cmd := exec.Command("xdg-open", dm.downloadPath)
	if err := cmd.Start(); err != nil {
		log.Printf("❌ Erro ao abrir pasta: %v", err)
		showErrorDialog("Erro", fmt.Sprintf("Não foi possível abrir a pasta:\n%v", err))
	}
}

// sendNotification envia notificação do sistema
func (dm *DownloadManager) sendNotification(title, message string) {
	cmd := exec.Command("notify-send", "-i", "download", title, message)
	cmd.Run()
}

// Funções auxiliares

func formatBytes(bytes uint64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	
	div, exp := uint64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

func formatDuration(d time.Duration) string {
	if d < time.Minute {
		return fmt.Sprintf("%ds", int(d.Seconds()))
	}
	if d < time.Hour {
		return fmt.Sprintf("%dm %ds", int(d.Minutes()), int(d.Seconds())%60)
	}
	return fmt.Sprintf("%dh %dm", int(d.Hours()), int(d.Minutes())%60)
}
