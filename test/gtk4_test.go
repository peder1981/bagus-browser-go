package main

import (
	"log"
	"os"

	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

func main() {
	// Criar aplicação GTK4
	app := gtk.NewApplication("com.bagus.browser.test", 0)
	
	app.ConnectActivate(func() {
		log.Println("🚀 GTK4 Test - Bagus Browser v5.0.0")
		
		// Criar janela
		win := gtk.NewApplicationWindow(app)
		win.SetTitle("Bagus Browser v5.0.0 - GTK4 Test")
		win.SetDefaultSize(800, 600)
		
		// Criar label
		label := gtk.NewLabel("✅ GTK4 funcionando!\n🎉 Bagus Browser v5.0.0 em desenvolvimento")
		label.SetMarginTop(20)
		label.SetMarginBottom(20)
		label.SetMarginStart(20)
		label.SetMarginEnd(20)
		
		// Adicionar à janela
		win.SetChild(label)
		
		// Mostrar
		win.Show()
		
		log.Println("✅ Janela GTK4 criada com sucesso!")
	})
	
	// Executar
	if code := app.Run(os.Args); code > 0 {
		log.Fatalf("Erro ao executar aplicação: %d", code)
	}
}
