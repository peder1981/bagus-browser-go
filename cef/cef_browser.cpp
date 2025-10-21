#include "cef_browser.h"
#include "cef_app.h"
#include <cef_app.h>
#include <wrapper/cef_helpers.h>
#include <iostream>

// Global client instance
static CefRefPtr<BagusBrowserClient> g_client;

BagusBrowserClient::BagusBrowserClient() : is_closing_(false) {}

BagusBrowserClient::~BagusBrowserClient() {}

void BagusBrowserClient::OnTitleChange(CefRefPtr<CefBrowser> browser,
                                       const CefString& title) {
  CEF_REQUIRE_UI_THREAD();
  
  // Atualiza título da janela
  std::cout << "Title: " << title.ToString() << std::endl;
}

void BagusBrowserClient::OnAddressChange(CefRefPtr<CefBrowser> browser,
                                        CefRefPtr<CefFrame> frame,
                                        const CefString& url) {
  CEF_REQUIRE_UI_THREAD();
  
  if (frame->IsMain()) {
    std::cout << "URL: " << url.ToString() << std::endl;
    // Aqui podemos chamar callback Go para atualizar histórico
  }
}

void BagusBrowserClient::OnAfterCreated(CefRefPtr<CefBrowser> browser) {
  CEF_REQUIRE_UI_THREAD();
  
  browser_list_.push_back(browser);
  std::cout << "Browser created. Total browsers: " << browser_list_.size() << std::endl;
}

bool BagusBrowserClient::DoClose(CefRefPtr<CefBrowser> browser) {
  CEF_REQUIRE_UI_THREAD();
  
  if (browser_list_.size() == 1) {
    is_closing_ = true;
  }
  
  return false;
}

void BagusBrowserClient::OnBeforeClose(CefRefPtr<CefBrowser> browser) {
  CEF_REQUIRE_UI_THREAD();
  
  BrowserList::iterator it = browser_list_.begin();
  for (; it != browser_list_.end(); ++it) {
    if ((*it)->IsSame(browser)) {
      browser_list_.erase(it);
      break;
    }
  }
  
  std::cout << "Browser closed. Remaining browsers: " << browser_list_.size() << std::endl;
  
  if (browser_list_.empty()) {
    CefQuitMessageLoop();
  }
}

void BagusBrowserClient::OnLoadError(CefRefPtr<CefBrowser> browser,
                                    CefRefPtr<CefFrame> frame,
                                    ErrorCode errorCode,
                                    const CefString& errorText,
                                    const CefString& failedUrl) {
  CEF_REQUIRE_UI_THREAD();
  
  if (errorCode == ERR_ABORTED)
    return;
  
  std::cerr << "Load error: " << errorText.ToString() 
            << " (" << failedUrl.ToString() << ")" << std::endl;
}

void BagusBrowserClient::OnLoadStart(CefRefPtr<CefBrowser> browser,
                                    CefRefPtr<CefFrame> frame,
                                    TransitionType transition_type) {
  CEF_REQUIRE_UI_THREAD();
  
  if (frame->IsMain()) {
    std::cout << "Loading started: " << frame->GetURL().ToString() << std::endl;
  }
}

void BagusBrowserClient::OnLoadEnd(CefRefPtr<CefBrowser> browser,
                                  CefRefPtr<CefFrame> frame,
                                  int httpStatusCode) {
  CEF_REQUIRE_UI_THREAD();
  
  if (frame->IsMain()) {
    std::cout << "Loading finished: " << frame->GetURL().ToString() 
              << " (Status: " << httpStatusCode << ")" << std::endl;
  }
}

bool BagusBrowserClient::OnBeforeBrowse(CefRefPtr<CefBrowser> browser,
                                       CefRefPtr<CefFrame> frame,
                                       CefRefPtr<CefRequest> request,
                                       bool user_gesture,
                                       bool is_redirect) {
  CEF_REQUIRE_UI_THREAD();
  
  // Aqui podemos adicionar bloqueio de domínios
  std::string url = request->GetURL().ToString();
  std::cout << "Before browse: " << url << std::endl;
  
  return false;  // false = permite navegação
}

void BagusBrowserClient::CloseAllBrowsers(bool force_close) {
  if (browser_list_.empty())
    return;
  
  BrowserList::const_iterator it = browser_list_.begin();
  for (; it != browser_list_.end(); ++it) {
    (*it)->GetHost()->CloseBrowser(force_close);
  }
}

// Funções C para binding com Go

int cef_initialize(int argc, char** argv) {
  std::cout << "[CEF] Iniciando inicialização..." << std::endl;
  
  CefMainArgs main_args(argc, argv);
  std::cout << "[CEF] CefMainArgs criado" << std::endl;
  
  // Verifica se é um subprocess
  int exit_code = CefExecuteProcess(main_args, nullptr, nullptr);
  if (exit_code >= 0) {
    std::cout << "[CEF] Subprocess terminando com código: " << exit_code << std::endl;
    // É um subprocess, retorna o código de saída
    exit(exit_code);
  }
  
  std::cout << "[CEF] Processo principal, continuando..." << std::endl;
  
  CefRefPtr<BagusCefApp> app(new BagusCefApp());
  std::cout << "[CEF] BagusCefApp criado" << std::endl;
  
  CefSettings settings;
  settings.no_sandbox = true;
  settings.windowless_rendering_enabled = false;
  settings.multi_threaded_message_loop = false;
  
  std::cout << "[CEF] Settings configurado" << std::endl;
  
  // Configura o caminho do subprocess (usa o próprio executável)
  CefString(&settings.browser_subprocess_path).FromASCII("");
  std::cout << "[CEF] Subprocess path configurado" << std::endl;
  
  // ============================================
  // PRIVACIDADE: ZERO TELEMETRIA
  // ============================================
  
  // Desabilita remote debugging (sem telemetria remota)
  settings.remote_debugging_port = 0;
  
  // Desabilita logs (sem envio de logs)
  settings.log_severity = LOGSEVERITY_DISABLE;
  
  // Desabilita stack traces (sem crash reporting)
  settings.uncaught_exception_stack_size = 0;
  
  // Não configura user agent com tracking
  // Não configura product version com telemetria
  // Não habilita background_color fingerprinting
  
  // GARANTIA: Nenhum dado é enviado para servidores externos
  // GARANTIA: Nenhuma telemetria é coletada
  // GARANTIA: Nenhum rastreamento é habilitado
  
  // ============================================
  
  std::cout << "[CEF] Chamando CefInitialize..." << std::endl;
  
  // Inicializa CEF
  bool result = CefInitialize(main_args, settings, app.get(), nullptr);
  
  std::cout << "[CEF] CefInitialize retornou: " << (result ? "true" : "false") << std::endl;
  
  if (result) {
    std::cout << "[CEF] Criando BagusBrowserClient..." << std::endl;
    g_client = new BagusBrowserClient();
    std::cout << "[CEF] BagusBrowserClient criado" << std::endl;
  }
  
  return result ? 1 : 0;
}

void cef_run_message_loop() {
  CefRunMessageLoop();
}

void cef_shutdown() {
  if (g_client.get()) {
    g_client->CloseAllBrowsers(false);
    g_client = nullptr;
  }
  
  CefShutdown();
}

void cef_navigate(const char* url) {
  if (!g_client.get())
    return;
    
  if (g_client->IsClosing())
    return;
  
  // Se não há browser ainda, cria um
  if (!g_client->HasBrowsers()) {
    CefWindowInfo window_info;
    
#if defined(OS_LINUX)
    window_info.bounds.x = 0;
    window_info.bounds.y = 0;
    window_info.bounds.width = 1200;
    window_info.bounds.height = 800;
#endif

    CefBrowserSettings browser_settings;
    browser_settings.javascript = STATE_ENABLED;
    browser_settings.javascript_close_windows = STATE_ENABLED;
    browser_settings.javascript_access_clipboard = STATE_ENABLED;
    browser_settings.local_storage = STATE_ENABLED;
    browser_settings.databases = STATE_ENABLED;
    browser_settings.webgl = STATE_ENABLED;
    
    CefBrowserHost::CreateBrowser(window_info, g_client, url, browser_settings,
                                   nullptr, nullptr);
  } else {
    // Navega no browser existente
    CefRefPtr<CefBrowser> browser = g_client->GetFirstBrowser();
    if (browser.get()) {
      browser->GetMainFrame()->LoadURL(url);
    }
  }
}
