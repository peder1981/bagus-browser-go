#ifndef CEF_BROWSER_H_
#define CEF_BROWSER_H_

#include <cef_client.h>
#include <cef_display_handler.h>
#include <cef_life_span_handler.h>
#include <cef_load_handler.h>
#include <cef_request_handler.h>
#include <list>

// Client handler para o browser
class BagusBrowserClient : public CefClient,
                           public CefDisplayHandler,
                           public CefLifeSpanHandler,
                           public CefLoadHandler,
                           public CefRequestHandler {
 public:
  BagusBrowserClient();
  ~BagusBrowserClient();

  // CefClient methods
  virtual CefRefPtr<CefDisplayHandler> GetDisplayHandler() override {
    return this;
  }
  virtual CefRefPtr<CefLifeSpanHandler> GetLifeSpanHandler() override {
    return this;
  }
  virtual CefRefPtr<CefLoadHandler> GetLoadHandler() override {
    return this;
  }
  virtual CefRefPtr<CefRequestHandler> GetRequestHandler() override {
    return this;
  }

  // CefDisplayHandler methods
  virtual void OnTitleChange(CefRefPtr<CefBrowser> browser,
                            const CefString& title) override;
  virtual void OnAddressChange(CefRefPtr<CefBrowser> browser,
                              CefRefPtr<CefFrame> frame,
                              const CefString& url) override;

  // CefLifeSpanHandler methods
  virtual void OnAfterCreated(CefRefPtr<CefBrowser> browser) override;
  virtual bool DoClose(CefRefPtr<CefBrowser> browser) override;
  virtual void OnBeforeClose(CefRefPtr<CefBrowser> browser) override;

  // CefLoadHandler methods
  virtual void OnLoadError(CefRefPtr<CefBrowser> browser,
                          CefRefPtr<CefFrame> frame,
                          ErrorCode errorCode,
                          const CefString& errorText,
                          const CefString& failedUrl) override;
  virtual void OnLoadStart(CefRefPtr<CefBrowser> browser,
                          CefRefPtr<CefFrame> frame,
                          TransitionType transition_type) override;
  virtual void OnLoadEnd(CefRefPtr<CefBrowser> browser,
                        CefRefPtr<CefFrame> frame,
                        int httpStatusCode) override;

  // CefRequestHandler methods
  virtual bool OnBeforeBrowse(CefRefPtr<CefBrowser> browser,
                             CefRefPtr<CefFrame> frame,
                             CefRefPtr<CefRequest> request,
                             bool user_gesture,
                             bool is_redirect) override;

  // Métodos públicos
  void CloseAllBrowsers(bool force_close);
  bool IsClosing() const { return is_closing_; }
  bool HasBrowsers() const { return !browser_list_.empty(); }
  CefRefPtr<CefBrowser> GetFirstBrowser() const {
    return browser_list_.empty() ? nullptr : browser_list_.front();
  }

 private:
  // Lista de browsers
  typedef std::list<CefRefPtr<CefBrowser>> BrowserList;
  BrowserList browser_list_;

  bool is_closing_;

  IMPLEMENT_REFCOUNTING(BagusBrowserClient);
};

// Funções C para binding com Go
#ifdef __cplusplus
extern "C" {
#endif

// Inicializa CEF
int cef_initialize(int argc, char** argv);

// Executa message loop
void cef_run_message_loop();

// Shutdown CEF
void cef_shutdown();

// Navega para URL
void cef_navigate(const char* url);

#ifdef __cplusplus
}
#endif

#endif  // CEF_BROWSER_H_
