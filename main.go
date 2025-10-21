package main

import (
	"log"

	"github.com/peder1981/bagus-browser-go/internal/browser"
	"github.com/peder1981/bagus-browser-go/internal/lockfile"
)

func main() {
	log.Println("🌐 Iniciando Bagus Browser v3.0.0...")

	// Verifica se já existe outra instância rodando
	lock := lockfile.New("bagus-browser")
	if err := lock.TryLock(); err != nil {
		log.Fatalf("❌ Erro: %v", err)
	}
	defer lock.Unlock()

	// Cria coordenador (gerencia ambas as janelas)
	coordinator := browser.NewCoordinator()

	// Inicia browser (bloqueia até fechar)
	log.Println("✅ Browser iniciado com sucesso!")
	if err := coordinator.Start(); err != nil {
		log.Fatalf("❌ Erro ao executar browser: %v", err)
	}

	log.Println("👋 Browser encerrado")
}
