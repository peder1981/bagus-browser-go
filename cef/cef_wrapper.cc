#include "cef_browser.h"
#include <stdlib.h>

// Wrapper C para as funções C++ do CEF
extern "C" {

int cef_initialize_wrapper(int argc, char** argv) {
    return cef_initialize(argc, argv);
}

void cef_run_message_loop_wrapper() {
    cef_run_message_loop();
}

void cef_shutdown_wrapper() {
    cef_shutdown();
}

void cef_navigate_wrapper(const char* url) {
    cef_navigate(url);
}

}
