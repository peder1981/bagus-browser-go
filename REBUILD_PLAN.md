# 🏗️ Plano de Reconstrução - Bagus Browser v3.0.0

## 🎯 Objetivo

Reconstruir o Bagus Browser do zero com arquitetura de **2 janelas**:
- **Janela de Controle**: Barra de navegação, botões, histórico
- **Janela de Conteúdo**: WebView para renderizar sites

## 🏛️ Pilares do Projeto

1. **Segurança** - Validação de URLs, sanitização, proteção
2. **Robustez** - Tratamento de erros, recuperação de falhas
3. **Leveza** - ~4MB, sem dependências pesadas
4. **Privacidade** - Zero telemetria, dados locais, bloqueio de ads

## 📐 Nova Arquitetura

```
┌─────────────────────────────────────────┐
│  Janela de Controle (GTK)               │
│  ┌───┐┌───┐┌───┐┌──────────────────┐   │
│  │ ← ││ → ││ ⟳ ││ URL Input        │   │
│  └───┘└───┘└───┘└──────────────────┘   │
└─────────────────────────────────────────┘
              ↓ controla
┌─────────────────────────────────────────┐
│  Janela de Conteúdo (WebView)           │
│                                          │
│  [Site renderizado aqui]                │
│                                          │
└─────────────────────────────────────────┘
```

## 📁 Nova Estrutura de Diretórios

```
bagus-browser-go/
├── cmd/
│   └── bagus/
│       └── main.go              # Entry point
├── internal/
│   ├── browser/
│   │   ├── browser.go           # Orquestrador principal
│   │   ├── control_window.go   # Janela de controle (GTK)
│   │   ├── content_window.go   # Janela de conteúdo (WebView)
│   │   └── coordinator.go      # Sincronização entre janelas
│   ├── navigation/
│   │   ├── history.go           # Histórico de navegação
│   │   ├── validator.go         # Validação de URLs
│   │   └── processor.go         # Processamento de input
│   ├── security/
│   │   ├── blocker.go           # Bloqueador de ads/trackers
│   │   ├── sanitizer.go         # Sanitização de URLs
│   │   └── validator.go         # Validação de segurança
│   ├── storage/
│   │   ├── config.go            # Configurações
│   │   ├── history.go           # Persistência de histórico
│   │   └── cache.go             # Cache local
│   └── ui/
│       ├── gtk_helpers.go       # Helpers GTK
│       └── styles.go            # Estilos e temas
├── pkg/
│   └── ipc/
│       └── channel.go           # Comunicação entre janelas
├── configs/
│   └── default.json             # Configuração padrão
├── scripts/
│   ├── build.sh                 # Build script
│   └── install.sh               # Instalador
├── docs/
│   ├── ARCHITECTURE.md          # Documentação da arquitetura
│   ├── SECURITY.md              # Documentação de segurança
│   └── USAGE.md                 # Guia de uso
├── go.mod
├── go.sum
├── Makefile
├── README.md
└── LICENSE
```

## 🔧 Tecnologias

### Core
- **Go 1.21+** - Linguagem principal
- **gotk3** - Interface GTK para janela de controle
- **webview** - WebKit2GTK para janela de conteúdo

### Bibliotecas
- **gotk3/gotk3** - Bindings GTK3
- **webview/webview_go** - WebView
- Sem dependências externas pesadas

## 🎨 Features da Janela de Controle

### Barra de Navegação
- ✅ Botão Voltar (←)
- ✅ Botão Avançar (→)
- ✅ Botão Recarregar (⟳)
- ✅ Campo de URL com autocompletar
- ✅ Botão Ir/Enter

### Menu
- ✅ Histórico
- ✅ Configurações
- ✅ Sobre
- ✅ Sair

### Indicadores
- ✅ Status de carregamento
- ✅ Indicador de segurança (HTTPS)
- ✅ Contador de bloqueios (ads/trackers)

## 🔒 Features de Segurança

### Validação
- ✅ Validação de URLs
- ✅ Sanitização de input
- ✅ Proteção contra XSS
- ✅ Proteção contra injeção

### Privacidade
- ✅ Zero telemetria
- ✅ Bloqueador de ads integrado
- ✅ Bloqueador de trackers
- ✅ Dados armazenados localmente
- ✅ Sem cookies de terceiros

### Bloqueio
- ✅ Lista de hosts maliciosos
- ✅ Lista de ads conhecidos
- ✅ Lista de trackers
- ✅ Atualização automática de listas

## 📊 Comunicação Entre Janelas

### IPC (Inter-Process Communication)

```go
type Message struct {
    Type    string      // "navigate", "back", "forward", "reload"
    Payload interface{} // Dados específicos
}

// Canal de comunicação
controlWindow → channel → contentWindow
contentWindow → channel → controlWindow
```

### Eventos

**Control → Content:**
- `Navigate(url string)` - Navegar para URL
- `Back()` - Voltar
- `Forward()` - Avançar
- `Reload()` - Recarregar
- `Stop()` - Parar carregamento

**Content → Control:**
- `URLChanged(url string)` - URL mudou
- `TitleChanged(title string)` - Título mudou
- `LoadingStarted()` - Começou a carregar
- `LoadingFinished()` - Terminou de carregar
- `LoadingFailed(error)` - Falha no carregamento

## 🚀 Fases de Implementação

### Fase 1: Core (1 hora)
- [x] Estrutura de diretórios
- [ ] Janela de controle básica (GTK)
- [ ] Janela de conteúdo básica (WebView)
- [ ] Comunicação IPC
- [ ] Navegação básica

### Fase 2: Navegação (30 min)
- [ ] Histórico
- [ ] Validação de URLs
- [ ] Processamento de input
- [ ] Autocompletar

### Fase 3: Segurança (30 min)
- [ ] Bloqueador de ads
- [ ] Validação de segurança
- [ ] Sanitização

### Fase 4: UI/UX (30 min)
- [ ] Estilos
- [ ] Ícones
- [ ] Feedback visual
- [ ] Indicadores

### Fase 5: Testes (30 min)
- [ ] Testes unitários
- [ ] Testes de integração
- [ ] Validação de segurança
- [ ] Testes de performance

## 📝 Checklist de Remoção

### Arquivos a Remover
- [ ] `internal/ui/browser*.go` (todos os arquivos antigos)
- [ ] `cef/` (diretório completo)
- [ ] `main_cef.go`
- [ ] `internal/cef/`
- [ ] Scripts CEF antigos
- [ ] Documentação CEF
- [ ] `KNOWN_LIMITATIONS.md` (será reescrito)
- [ ] `QUICK_USAGE.md` (será reescrito)

### Arquivos a Manter
- [x] `go.mod` / `go.sum` (atualizar)
- [x] `LICENSE`
- [x] `README.md` (reescrever)
- [x] `.gitignore`
- [x] `internal/browser/history.go` (refatorar)
- [x] `internal/security/` (refatorar)
- [x] `internal/config/` (refatorar)

## 🎯 Critérios de Sucesso

### Funcionalidade
- ✅ Navegação funciona 100%
- ✅ Barra de controle sempre visível
- ✅ Histórico funcional
- ✅ Bloqueio de ads funciona

### Performance
- ✅ Binário < 5MB
- ✅ Uso de RAM < 100MB (idle)
- ✅ Inicialização < 1s

### Segurança
- ✅ Zero telemetria
- ✅ Validação de URLs
- ✅ Bloqueio de trackers
- ✅ Dados locais apenas

### UX
- ✅ Interface responsiva
- ✅ Feedback visual claro
- ✅ Atalhos de teclado
- ✅ Sem travamentos

## 📅 Timeline

**Início:** 2025-10-21 16:30  
**Previsão:** 2025-10-21 19:30 (3 horas)

**Marcos:**
- 17:30 - Core funcional
- 18:00 - Navegação completa
- 18:30 - Segurança implementada
- 19:00 - UI/UX polido
- 19:30 - Testes e validação

## 🔄 Próximos Passos

1. ✅ Criar este documento
2. [ ] Remover código antigo
3. [ ] Criar nova estrutura
4. [ ] Implementar core
5. [ ] Testar e validar
6. [ ] Documentar
7. [ ] Release v3.0.0

---

**Status:** 🚀 Pronto para começar!  
**Versão:** 3.0.0  
**Arquitetura:** 2 Janelas (Control + Content)
