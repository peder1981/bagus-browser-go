# Plano de Refatoração Python → Go

## 📊 Análise do Projeto Python

### Estrutura Principal
```
browser/
├── browser.py          # Janela principal com tabs
├── form_login.py       # Tela de login/setup
├── ui/
│   ├── browser_tab.py  # Aba individual do browser
│   ├── private_profile.py  # Perfil privado com bloqueio
│   ├── custom_web_engine_page.py
│   └── table.py
├── panel_navigation.py # Painel de navegação
├── panel_myass.py      # Painel MyAss
├── panel_play.py       # Painel Play
└── api/
    ├── analyze.py      # Análise de segurança
    ├── aes_helper.py   # Criptografia
    ├── logger_helper.py
    ├── myass_helper.py
    └── project_helper.py

```

### Funcionalidades Principais

1. **Form de Login** (`form_login.py`)
   - Validação de username
   - Criação de diretórios
   - Setup inicial

2. **Browser Principal** (`browser.py`)
   - Sistema de tabs laterais (Browser, Download, Navigation, MyAss, etc.)
   - Gerenciamento de abas de navegação
   - Histórico
   - Configurações via JSON
   - Atalhos de teclado

3. **Browser Tab** (`browser_tab.py`)
   - WebView com QtWebEngine
   - Barra de URL com autocomplete
   - Histórico de sugestões
   - Execução de scripts customizados
   - Proxy (SOCKS5/HTTP)
   - Inspector/DevTools

4. **Private Profile** (`private_profile.py`)
   - Bloqueio de domínios (ad_hosts_block.txt)
   - Interceptor de requisições
   - 50+ configurações de WebEngine
   - Cache e storage persistente

5. **Painéis Adicionais**
   - Navigation: Scripts de navegação
   - MyAss: Workflow customizado
   - Play: Player de mídia
   - XMPP Chat
   - Extensions

## 🎯 Implementação em Go

### Fase 1: Core (AGORA)
- [x] Estrutura base
- [ ] Form de login/setup
- [ ] Janela principal
- [ ] WebView básico
- [ ] Sistema de abas
- [ ] Navegação (back, forward, reload)
- [ ] Barra de URL
- [ ] Histórico básico

### Fase 2: Segurança e Perfil
- [ ] Perfil privado
- [ ] Bloqueador de domínios
- [ ] Interceptor de requisições
- [ ] Configurações de segurança
- [ ] Criptografia AES

### Fase 3: Features Avançadas
- [ ] Autocomplete de URL
- [ ] Scripts customizados
- [ ] Proxy support
- [ ] DevTools/Inspector
- [ ] Painéis laterais

### Fase 4: Painéis Específicos
- [ ] Panel Navigation
- [ ] Panel MyAss
- [ ] Panel Play
- [ ] Extensions

## 🔧 Tecnologias Go

### UI Framework
**Opção escolhida**: `webview` (github.com/webview/webview)
- ✅ Multiplataforma
- ✅ Leve e simples
- ✅ Usa engine nativo (WebKit/Edge/GTK)
- ❌ Menos configurável que Qt

**Alternativa**: `go-qt` (therecipe/qt)
- ✅ Port completo do Qt
- ✅ Todas as features do PySide6
- ❌ Mais pesado
- ❌ Setup mais complexo

### Estrutura de Arquivos Go

```
bagus-browser-go/
├── cmd/bagus/
│   └── main.go
├── internal/
│   ├── ui/
│   │   ├── login.go        # Form de login
│   │   ├── mainwindow.go   # Janela principal
│   │   ├── browsertab.go   # Aba de navegação
│   │   └── panels/
│   │       ├── navigation.go
│   │       ├── myass.go
│   │       └── play.go
│   ├── browser/
│   │   ├── engine.go       # Motor do browser
│   │   ├── tab.go          # Gerenciamento de abas
│   │   ├── history.go      # Histórico
│   │   └── profile.go      # Perfil privado
│   ├── security/
│   │   ├── blocker.go      # Bloqueador de domínios
│   │   ├── interceptor.go  # Interceptor de requisições
│   │   ├── validator.go    # Validação de entrada
│   │   └── analyzer.go     # Análise de segurança
│   ├── crypto/
│   │   └── aes.go          # Criptografia AES
│   ├── storage/
│   │   ├── manager.go      # Gerenciamento de storage
│   │   ├── config.go       # Configurações
│   │   └── history.go      # Persistência de histórico
│   └── project/
│       └── helper.go       # Helper de projetos
├── pkg/
│   └── utils/
│       └── system.go
└── configs/
    └── template.json       # Template de configuração
```

## 📝 Mapeamento de Funcionalidades

### Python → Go

| Python | Go | Status |
|--------|-----|--------|
| `FormLogin` | `ui.LoginDialog` | ⏳ Pendente |
| `Browser` | `ui.MainWindow` | ⏳ Pendente |
| `BrowserTab` | `ui.BrowserTab` | ⏳ Pendente |
| `PrivateProfile` | `browser.Profile` | ⏳ Pendente |
| `WebEngineUrlRequestInterceptor` | `security.Interceptor` | ⏳ Pendente |
| `Analyze` | `security.Analyzer` | ⏳ Pendente |
| `ProjectHelper` | `project.Helper` | ⏳ Pendente |

## 🚧 Limitações e Adaptações

### Webview vs QtWebEngine

**QtWebEngine (Python)**:
- 50+ configurações granulares
- DevTools integrado
- Proxy por requisição
- Interceptor de requisições completo

**Webview (Go)**:
- Configurações básicas
- DevTools limitado
- Proxy global apenas
- Interceptor via proxy externo

### Solução Híbrida

Para features avançadas, podemos:
1. Usar proxy local em Go para interceptar
2. Injetar JavaScript para features customizadas
3. Usar comunicação bidirecional (Go ↔ JS)

## 📅 Timeline Estimado

- **Fase 1** (Core): 2-3 dias
- **Fase 2** (Segurança): 1-2 dias
- **Fase 3** (Features): 2-3 dias
- **Fase 4** (Painéis): 3-4 dias

**Total**: ~2 semanas de desenvolvimento

## 🎯 Prioridades

1. ✅ **Alta**: Login, Browser básico, Abas, Navegação
2. ⚠️ **Média**: Histórico, Bloqueador, Configurações
3. ℹ️ **Baixa**: Painéis específicos, Extensions

---

**Status Atual**: Iniciando Fase 1
**Próximo**: Implementar Form de Login
