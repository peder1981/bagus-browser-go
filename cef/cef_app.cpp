#include "cef_app.h"
#include "cef_browser.h"
#include <cef_browser.h>
#include <cef_command_line.h>
#include <views/cef_browser_view.h>
#include <views/cef_window.h>
#include <wrapper/cef_helpers.h>

BagusCefApp::BagusCefApp() {}

void BagusCefApp::OnContextInitialized() {
  CEF_REQUIRE_UI_THREAD();
  
  // Contexto inicializado
  // O browser será criado quando cef_navigate for chamado
}
