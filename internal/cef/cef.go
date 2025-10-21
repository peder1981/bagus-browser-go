//go:build cef
// +build cef

package cef

/*
#cgo LDFLAGS: -L${SRCDIR}/../../build -L/opt/cef/Release -L/opt/cef/build/libcef_dll_wrapper -lcef_wrapper -lcef -lcef_dll_wrapper -Wl,-rpath,/opt/cef/Release -lstdc++
#include <stdlib.h>

int cef_initialize_wrapper(int argc, char** argv);
void cef_run_message_loop_wrapper();
void cef_shutdown_wrapper();
void cef_navigate_wrapper(const char* url);
*/
import "C"
import (
	"log"
	"os"
	"unsafe"
)

// CEFBrowser representa o navegador CEF
type CEFBrowser struct {
	initialized bool
}

// NewCEFBrowser cria nova instância do navegador CEF
func NewCEFBrowser() *CEFBrowser {
	return &CEFBrowser{
		initialized: false,
	}
}

// Initialize inicializa o CEF
func (b *CEFBrowser) Initialize() error {
	log.Println("Inicializando CEF...")

	// Converte argumentos para C
	argc := C.int(len(os.Args))
	argv := make([]*C.char, len(os.Args))
	for i, arg := range os.Args {
		argv[i] = C.CString(arg)
		defer C.free(unsafe.Pointer(argv[i]))
	}

	// Inicializa CEF
	result := C.cef_initialize_wrapper(argc, &argv[0])
	if result == 0 {
		return os.ErrInvalid
	}

	b.initialized = true
	log.Println("CEF inicializado com sucesso")
	return nil
}

// Run executa o message loop do CEF
func (b *CEFBrowser) Run() {
	if !b.initialized {
		log.Fatal("CEF não foi inicializado")
		return
	}

	log.Println("Iniciando message loop do CEF...")
	C.cef_run_message_loop_wrapper()
	log.Println("Message loop finalizado")
}

// Shutdown finaliza o CEF
func (b *CEFBrowser) Shutdown() {
	if !b.initialized {
		return
	}

	log.Println("Finalizando CEF...")
	C.cef_shutdown_wrapper()
	b.initialized = false
	log.Println("CEF finalizado")
}

// Navigate navega para uma URL
func (b *CEFBrowser) Navigate(url string) {
	if !b.initialized {
		log.Println("CEF não inicializado, não é possível navegar")
		return
	}

	cURL := C.CString(url)
	defer C.free(unsafe.Pointer(cURL))

	C.cef_navigate_wrapper(cURL)
	log.Printf("Navegando para: %s", url)
}
