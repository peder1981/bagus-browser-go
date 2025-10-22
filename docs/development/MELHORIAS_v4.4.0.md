# 🎯 Melhorias v4.4.0 - Funcionalidades Avançadas

**Data:** 22/10/2025 09:18 BRT  
**Versão:** 4.4.0

---

## 📋 Pontos Identificados

### 1. ❌ Copiar/Colar Imagens da Área de Transferência
**Status:** NÃO IMPLEMENTADO  
**Prioridade:** MÉDIA

**Requisitos:**
- Ctrl+C / Ctrl+Ins - Copiar imagem
- Ctrl+V / Shift+Ins - Colar imagem
- Suporte a área de transferência do sistema

**Implementação:**
- WebKit2GTK suporta clipboard nativamente
- Precisa habilitar contexto de edição
- Adicionar atalhos de teclado

---

### 2. ❌ Downloads Não Funcionam
**Status:** PARCIALMENTE IMPLEMENTADO  
**Prioridade:** CRÍTICA

**Problema:**
- DownloadManager existe mas não está conectado ao WebView
- Sinal `download-started` não está conectado
- Falta tratamento de downloads

**Solução:**
```go
// Conectar sinal de download no WebView
webView.widget.Connect("download-started", func(wv *webkit.WebView, download *webkit.Download) {
    // Tratar download
})
```

---

### 3. ❌ Restauração de Sessão
**Status:** NÃO IMPLEMENTADO  
**Prioridade:** ALTA

**Requisitos:**
- Salvar abas abertas ao fechar
- Restaurar abas ao reabrir
- Salvar URLs e posição de scroll
- Arquivo: `~/.config/bagus-browser/session.enc`

**Implementação:**
```go
type Session struct {
    Tabs []SessionTab `json:"tabs"`
}

type SessionTab struct {
    URL    string `json:"url"`
    Title  string `json:"title"`
    Active bool   `json:"active"`
}
```

---

### 4. ✅ Ctrl+T Não Foca Nova Aba
**Status:** JÁ IMPLEMENTADO (mas precisa verificar)  
**Prioridade:** BAIXA

**Código Atual:**
```go
// Linha 736: b.notebook.SetCurrentPage(pageNum)
// Linha 780-781: b.urlEntry.GrabFocus()
```

**Possível Problema:**
- Timing: LoadURI pode estar atrasando o foco
- Solução: Usar glib.IdleAdd para garantir execução após renderização

---

### 5. ❌ Notificações Web
**Status:** BLOQUEADO (por privacidade)  
**Prioridade:** MÉDIA

**Problema:**
- Atualmente bloqueado em `privacy.go`
- Linha 23: `webkit_settings_set_enable_notifications(settings, FALSE)`

**Solução:**
- Criar opção de configuração
- Permitir notificações por site (whitelist)
- Mostrar permissão ao usuário

---

### 6. ❌ Gerenciadores de Senha (Proton Pass)
**Status:** NÃO IMPLEMENTADO  
**Prioridade:** ALTA

**Requisitos:**
- Suporte a extensões de navegador
- WebExtensions API
- Ou: Permitir injeção de scripts

**Problema:**
- WebKit2GTK não suporta extensões Chrome/Firefox nativamente
- Proton Pass requer WebExtensions API

**Soluções Possíveis:**
1. **UserScripts:** Permitir scripts customizados
2. **Native Messaging:** Comunicação com app nativo Proton Pass
3. **Clipboard:** Copiar/colar senhas manualmente

**Recomendação:**
- Implementar UserScripts para permitir injeção de código
- Documentar como usar com gerenciadores de senha

---

### 7. ❌ Impressão (Ctrl+P)
**Status:** NÃO IMPLEMENTADO  
**Prioridade:** MÉDIA

**Requisitos:**
- Ctrl+P - Abrir diálogo de impressão
- Imprimir para PDF
- Imprimir para impressora física
- Preview de impressão

**Implementação:**
```go
// WebKit2GTK tem suporte nativo
webView.widget.Connect("print", func() {
    // Criar WebKitPrintOperation
    printOp := webkit.WebKitPrintOperationNew(webView.widget)
    printOp.RunDialog(b.window)
})
```

---

## 🎯 Plano de Implementação

### Fase 1: Críticas (v4.4.0)
1. ✅ Downloads funcionando
2. ✅ Restauração de sessão
3. ✅ Impressão (Ctrl+P)

### Fase 2: Importantes (v4.5.0)
4. ✅ Copiar/Colar imagens
5. ✅ Notificações (com controle)
6. ✅ Verificar Ctrl+T

### Fase 3: Avançadas (v4.6.0)
7. ✅ Suporte a gerenciadores de senha (UserScripts)

---

## 📊 Análise Técnica

### Downloads
**Complexidade:** BAIXA  
**Tempo estimado:** 30 minutos  
**Arquivos:** `main.go`, `downloads.go`

### Restauração de Sessão
**Complexidade:** MÉDIA  
**Tempo estimado:** 1 hora  
**Arquivos:** `session.go` (novo), `main.go`

### Impressão
**Complexidade:** BAIXA  
**Tempo estimado:** 20 minutos  
**Arquivos:** `main.go`

### Copiar/Colar Imagens
**Complexidade:** MÉDIA  
**Tempo estimado:** 45 minutos  
**Arquivos:** `main.go`

### Notificações
**Complexidade:** MÉDIA  
**Tempo estimado:** 30 minutos  
**Arquivos:** `privacy.go`, `main.go`

### Gerenciadores de Senha
**Complexidade:** ALTA  
**Tempo estimado:** 2-3 horas  
**Arquivos:** `userscripts.go` (novo), `main.go`

**Total Estimado:** 5-6 horas

---

## 🚀 Próximos Passos

1. Implementar downloads (CRÍTICO)
2. Implementar restauração de sessão (IMPORTANTE)
3. Implementar impressão (RÁPIDO)
4. Verificar/corrigir Ctrl+T
5. Implementar notificações com controle
6. Implementar copiar/colar imagens
7. Documentar uso com gerenciadores de senha

---

**Criado em:** 22/10/2025 09:18 BRT  
**Versão alvo:** 4.4.0  
**Status:** Planejamento
