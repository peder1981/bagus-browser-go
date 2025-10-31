# ✅ Checklist de Integração v4.x → v5.0.0

**Data**: 30/10/2025  
**Objetivo**: Migrar TODAS as funcionalidades v4.x para v5.0.0

## 📋 Estrutura Browser

### Componentes Principais
- [ ] window (GtkWindow)
- [ ] notebook (GtkNotebook) - Sistema de abas
- [ ] urlEntry (GtkEntry) - Barra de URL
- [ ] tabs ([]Tab) - Lista de abas
- [ ] closedTabs ([]ClosedTab) - Histórico de abas fechadas
- [ ] validator (URLValidator)
- [ ] privacyManager (PrivacyManager)
- [ ] bookmarkManager (BookmarkManager)
- [ ] downloadManager (DownloadManager)
- [ ] downloadHandler (DownloadHandler)
- [ ] sessionManager (SessionManager)
- [ ] config (Config)
- [ ] findBar (GtkBox) - Barra de busca
- [ ] findEntry (GtkEntry) - Campo de busca
- [ ] findBarVisible (bool)

## 🔧 Funções CGo (WebKit)

### Navegação Básica ✅
- [x] create_webview()
- [x] load_uri()
- [x] go_back()
- [x] go_forward()
- [x] reload_page()
- [x] get_uri()
- [x] get_title()
- [x] can_go_back()
- [x] can_go_forward()

### Zoom
- [ ] set_zoom_level()
- [ ] get_zoom_level()

### Busca na Página
- [ ] get_find_controller()
- [ ] find_text()
- [ ] find_next()
- [ ] find_previous()
- [ ] find_finish()

### Impressão
- [ ] print_page()

## 🎨 Interface GTK4

### Janela Principal
- [ ] Criar janela com ícone
- [ ] Tamanho padrão 1200x800
- [ ] Título dinâmico

### Menu Bar
- [ ] Arquivo
  - [ ] Nova Aba (Ctrl+T)
  - [ ] Nova Janela (Ctrl+N)
  - [ ] Reabrir Aba (Ctrl+Shift+T)
  - [ ] Imprimir (Ctrl+P)
  - [ ] Sair (Ctrl+Q)
- [ ] Editar
  - [ ] Copiar (Ctrl+C)
  - [ ] Colar (Ctrl+V)
  - [ ] Buscar (Ctrl+F)
  - [ ] Configurações (Ctrl+,)
- [ ] Ver
  - [ ] Zoom In (Ctrl++)
  - [ ] Zoom Out (Ctrl+-)
  - [ ] Zoom Reset (Ctrl+0)
  - [ ] Tela Cheia (F11)
- [ ] Favoritos
  - [ ] Adicionar Favorito (Ctrl+D)
  - [ ] Gerenciar Favoritos (Ctrl+B)
  - [ ] Barra de Favoritos
- [ ] Histórico
  - [ ] Mostrar Histórico (Ctrl+H)
  - [ ] Limpar Histórico
- [ ] Ferramentas
  - [ ] Downloads (Ctrl+J)
  - [ ] Modo Privado
- [ ] Ajuda
  - [ ] Sobre

### Toolbar
- [ ] Botão Voltar
- [ ] Botão Avançar
- [ ] Botão Recarregar
- [ ] Botão Home
- [ ] Entry URL
- [ ] Botão Favoritos
- [ ] Botão Menu

### Notebook (Abas)
- [ ] Criar aba
- [ ] Fechar aba
- [ ] Alternar abas
- [ ] Mover abas
- [ ] Ícone de carregamento
- [ ] Favicon
- [ ] Botão fechar na aba

### Barra de Busca
- [ ] Entry de busca
- [ ] Botão anterior
- [ ] Botão próximo
- [ ] Contador de resultados
- [ ] Fechar barra

## ⌨️ Atalhos de Teclado (31 total)

### Navegação
- [ ] Ctrl+T - Nova aba
- [ ] Ctrl+W - Fechar aba
- [ ] Ctrl+Shift+T - Reabrir aba
- [ ] Ctrl+Tab - Próxima aba
- [ ] Ctrl+Shift+Tab - Aba anterior
- [ ] Ctrl+1-9 - Ir para aba N
- [ ] Alt+Left - Voltar
- [ ] Alt+Right - Avançar
- [ ] Ctrl+R / F5 - Recarregar
- [ ] Ctrl+Shift+R - Recarregar (limpar cache)
- [ ] Ctrl+L - Focar URL
- [ ] Ctrl+N - Nova janela

### Edição
- [ ] Ctrl+C - Copiar
- [ ] Ctrl+V - Colar
- [ ] Ctrl+X - Cortar
- [ ] Ctrl+A - Selecionar tudo
- [ ] Ctrl+F - Buscar na página
- [ ] Ctrl+G / F3 - Próximo resultado
- [ ] Ctrl+Shift+G / Shift+F3 - Resultado anterior

### Zoom
- [ ] Ctrl++ - Zoom in
- [ ] Ctrl+- - Zoom out
- [ ] Ctrl+0 - Reset zoom

### Favoritos
- [ ] Ctrl+D - Adicionar favorito
- [ ] Ctrl+B - Gerenciar favoritos

### Outros
- [ ] Ctrl+H - Histórico
- [ ] Ctrl+J - Downloads
- [ ] Ctrl+P - Imprimir
- [ ] Ctrl+, - Configurações
- [ ] Ctrl+Q - Sair
- [ ] F11 - Tela cheia
- [ ] Esc - Parar carregamento / Sair tela cheia

## 📦 Gerenciadores

### BookmarkManager
- [ ] Adicionar favorito
- [ ] Remover favorito
- [ ] Listar favoritos
- [ ] Buscar favoritos
- [ ] Salvar (criptografado)
- [ ] Carregar

### DownloadManager
- [ ] Iniciar download
- [ ] Pausar download
- [ ] Cancelar download
- [ ] Listar downloads
- [ ] Abrir arquivo
- [ ] Abrir pasta

### SessionManager
- [ ] Salvar sessão
- [ ] Carregar sessão
- [ ] Restaurar abas

### PrivacyManager
- [ ] Limpar cookies
- [ ] Limpar cache
- [ ] Limpar histórico
- [ ] Modo privado

### Config
- [ ] Carregar configurações
- [ ] Salvar configurações
- [ ] Senha mestre
- [ ] Configurações de privacidade
- [ ] Configurações de cache

## 🎨 Diálogos GTK4

### Favoritos
- [ ] Adicionar favorito
- [ ] Gerenciar favoritos
- [ ] Editar favorito
- [ ] Remover favorito

### Downloads
- [ ] Gerenciador de downloads
- [ ] Progresso de download
- [ ] Notificação de conclusão

### Configurações
- [ ] Aba Segurança
- [ ] Aba Sessões
- [ ] Aba Performance
- [ ] Aba Privacidade

### Busca
- [ ] Barra de busca
- [ ] Resultados
- [ ] Navegação

### Outros
- [ ] Sobre
- [ ] Confirmação de saída
- [ ] Autenticação

## 🖼️ Recursos

### Ícone
- [ ] Ícone da aplicação (window icon)
- [ ] Ícone do .desktop
- [ ] Ícones da toolbar
- [ ] Favicons

### Temas
- [ ] CSS customizado
- [ ] Cores
- [ ] Fontes

## 🔐 Segurança

### Criptografia
- [ ] AES-256-GCM para sessões
- [ ] AES-256-GCM para favoritos
- [ ] AES-256-GCM para configurações
- [ ] bcrypt para senha mestre

### Privacidade
- [ ] Bloqueio de trackers
- [ ] Bloqueio de cookies third-party
- [ ] User agent customizado
- [ ] DNS prefetching desabilitado

## 📊 Progresso

```
[░░░░░░░░░░░░░░░░░░░░] 0% - Checklist criado
```

---

**Próximo passo**: Começar implementação da estrutura Browser em v5
