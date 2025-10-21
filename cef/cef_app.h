#ifndef CEF_APP_H_
#define CEF_APP_H_

#include <cef_app.h>

// Implementação do CefApp
class BagusCefApp : public CefApp,
                    public CefBrowserProcessHandler {
 public:
  BagusCefApp();

  // CefApp methods
  virtual CefRefPtr<CefBrowserProcessHandler> GetBrowserProcessHandler() override {
    return this;
  }

  // CefBrowserProcessHandler methods
  virtual void OnContextInitialized() override;

 private:
  IMPLEMENT_REFCOUNTING(BagusCefApp);
};

#endif  // CEF_APP_H_
