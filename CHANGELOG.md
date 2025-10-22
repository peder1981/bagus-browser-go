# Changelog

Todas as mudanças notáveis neste projeto serão documentadas neste arquivo.

O formato é baseado em [Keep a Changelog](https://keepachangelog.com/pt-BR/1.0.0/),
e este projeto adere ao [Semantic Versioning](https://semver.org/lang/pt-BR/).

## [4.0.0] - 2025-10-21

### 🎊 Release Oficial

Primeira versão de produção do Bagus Browser com todos os 4 pilares implementados!

### ✨ Adicionado
- WebView funcionando via CGO (WebKit2GTK-4.0)
- Sistema de múltiplas abas independentes
- Navegação completa (voltar, avançar, recarregar)
- Títulos dinâmicos nas abas (URL ou título da página)
- Funcionalidade de zoom (Ctrl++, Ctrl+-, Ctrl+0)
- 9 atalhos de teclado
- Busca integrada com DuckDuckGo
- Detecção automática URL vs termo de busca

### 🔒 Segurança
- Validação rigorosa de URLs (HTTP/HTTPS apenas)
- Sanitização de input (proteção XSS)
- WebView configurado com segurança máxima
- Plugins desabilitados
- Java desabilitado
- Lista de bloqueio de domínios
- User-Agent customizado (Bagus/4.0)

### 🕵️ Privacidade
- **Zero telemetria** (garantido)
- **Zero rastreamento** (garantido)
- Third-party cookies bloqueados
- WebGL bloqueado (anti-fingerprinting)
- WebAudio bloqueado (anti-fingerprinting)
- Cache offline desabilitado
- Do Not Track habilitado
- DuckDuckGo como motor de busca padrão

### 📚 Documentação
- README.md completo
- SECURITY.md - Documentação de segurança
- PRIVACY.md - Política de privacidade
- LICENSE (MIT)
- CHANGELOG.md (este arquivo)

### 🎯 Pilares
- ✅ Segurança - Implementado
- ✅ Privacidade - Implementado
- ✅ Robustez - Implementado
- ✅ Leveza - Implementado (5.5MB)

### ⌨️ Atalhos de Teclado
- Ctrl+T - Nova aba
- Ctrl+W - Fechar aba
- Alt+← - Voltar
- Alt+→ - Avançar
- F5 / Ctrl+R - Recarregar
- Ctrl+L - Focar URL
- Ctrl++ - Aumentar zoom
- Ctrl+- - Diminuir zoom
- Ctrl+0 - Resetar zoom

---

## [4.4.0] - 2025-10-22

### ✨ Novas Funcionalidades

#### Impressão (Ctrl+P)
- ✅ Diálogo nativo de impressão
- ✅ Imprimir para PDF
- ✅ Imprimir para impressoras físicas
- ✅ WebKit2GTK PrintOperation

#### Restauração de Sessão
- ✅ Salva todas as abas ao fechar
- ✅ Restaura abas ao reabrir
- ✅ Arquivo criptografado: `~/.config/bagus-browser/session.enc`
- ✅ AES-256-GCM + PBKDF2

#### Ctrl+Shift+T - Reabrir Abas Fechadas
- ✅ Histórico de até 10 abas fechadas
- ✅ LIFO (Last In, First Out)
- ✅ Não salva abas vazias

#### Downloads
- ✅ Handler de downloads conectado
- ✅ Downloads automáticos para ~/Downloads
- ✅ Logs informativos

#### Notificações Web
- ✅ Habilitadas por padrão
- ✅ Controladas por site
- ✅ Permissões gerenciadas pelo WebKit

#### Copiar/Colar Avançado
- ✅ Ctrl+Ins - Copiar (alternativo)
- ✅ Shift+Ins - Colar (alternativo)
- ✅ Suporte a imagens e HTML

### 🐛 Correções

#### Ctrl+T Melhorado
- ✅ Foco imediato na nova aba (glib.IdleAdd)
- ✅ Cursor na barra de URL garantido

### 📚 Documentação

#### Gerenciadores de Senha
- ✅ Guia completo: `docs/GERENCIADORES_SENHA.md`
- ✅ Compatibilidade: Proton Pass, KeePassXC, Bitwarden, 1Password
- ✅ Workflows e exemplos

#### Análise Técnica
- ✅ `docs/development/MELHORIAS_v4.4.0.md`
- ✅ Detalhes de implementação
- ✅ Testes e validações

### 📊 Estatísticas
- **Atalhos:** 30 (antes: 27)
- **Arquivos novos:** 2 (session.go, docs/GERENCIADORES_SENHA.md)
- **Linhas adicionadas:** ~850

---

## [4.3.0] - 2025-10-22

### 🐛 Correções Críticas

#### Favoritos Criptografados - CORRIGIDO
- **Problema:** Favoritos não carregavam entre sessões
- **Causa:** Chave de criptografia era gerada aleatoriamente a cada execução
- **Solução:** Chave derivada de hostname + username (persistente)
- ✅ Favoritos agora carregam corretamente ao reabrir browser
- ✅ Persistência entre sessões garantida
- ✅ Segurança mantida (AES-256-GCM + PBKDF2)

### ⌨️ Navegação Entre Abas

#### Novos Atalhos (11 adicionados)
- **Ctrl+Tab** - Próxima aba (circular)
- **Ctrl+Shift+Tab** - Aba anterior (circular)
- **Ctrl+1 a Ctrl+9** - Ir direto para aba específica (9 atalhos)

#### Funcionalidades
- ✅ Navegação circular entre abas
- ✅ Atualização automática da URL entry
- ✅ Funções NextTab(), PreviousTab(), GoToTab()
- ✅ Validação de índices
- ✅ Logs informativos

### 📊 Estatísticas
- **Atalhos:** 27 (antes: 16)
- **Funções novas:** 3
- **Linhas adicionadas:** ~80

---

## [4.2.0] - 2025-10-21

### 🎨 Melhorias de UX

#### Adicionado
- **Menu Superior Completo** - Menu bar profissional com 4 seções organizadas
  - Menu Arquivo (Nova Aba, Fechar Aba, Sair)
  - Menu Navegação (Voltar, Avançar, Recarregar)
  - Menu Favoritos (Adicionar, Gerenciar)
  - Menu Ferramentas (Buscar, Zoom)
- **Foco Automático na Barra de URL** - Ao abrir nova aba, cursor vai direto para a barra
- **Novo Atalho Ctrl+Q** - Sair do navegador
- Texto selecionado automaticamente ao focar barra de URL

#### Corrigido
- **Ctrl+Shift+B** - Gerenciar Favoritos agora funciona corretamente
  - Problema: GTK retorna KEY_B (maiúscula) quando Shift pressionado
  - Solução: Aceita tanto KEY_b quanto KEY_B
- Nome do projeto simplificado para "Bagus Browser" (removido "POC" e "WebKit CGO")

#### Melhorado
- Interface mais intuitiva e profissional
- Descoberta de funcionalidades facilitada via menu
- UX de navegação mais rápida e natural
- Acessibilidade melhorada com múltiplas formas de acesso

### ⌨️ Atalhos Atualizados (16 total)
- Ctrl+Q - Sair (NOVO)
- Ctrl+Shift+B - Gerenciar Favoritos (CORRIGIDO)
- Todos os atalhos anteriores mantidos

---

## [4.1.0] - 2025-10-21

### ✨ Features Implementadas

#### Adicionado
- **Favoritos com Criptografia AES-256-GCM**
  - Ctrl+D para adicionar
  - Ctrl+Shift+B para gerenciar
  - PBKDF2 com 100,000 iterações
  - Storage: ~/.config/bagus-browser/bookmarks.enc
- **Buscar na Página**
  - Ctrl+F para abrir busca
  - F3 para próximo resultado
  - Shift+F3 para resultado anterior
  - Busca incremental case-insensitive
- **Gerenciador de Downloads**
  - Pasta padrão: ~/Downloads
  - Nomes únicos automáticos
- **Sistema de Zoom Completo**
  - Ctrl++ para aumentar
  - Ctrl+- para diminuir
  - Ctrl+0 para resetar
  - Zoom independente por aba

#### Segurança
- Criptografia AES-256-GCM para dados sensíveis
- PBKDF2 para derivação de chaves
- SHA-256 para hashing
- Modo autenticado (GCM) para proteção contra adulteração

### ⌨️ Atalhos (15 total)
- Navegação: 6 atalhos
- Zoom: 3 atalhos
- Busca: 4 atalhos
- Favoritos: 2 atalhos

---

## [Unreleased] - v5.0

### Planejado
- Sistema de extensões
- Temas (claro/escuro)
- Sincronização (opcional)
- Modo super-privado

---

## Notas de Versão

### Versão 4.0.0
Esta é a primeira versão de produção do Bagus Browser. Todos os 4 pilares fundamentais (Segurança, Privacidade, Robustez, Leveza) foram implementados e testados.

O browser está pronto para uso real e oferece uma experiência de navegação segura e privada.

**Destaques:**
- Browser completo e funcional
- Zero telemetria garantido
- Código limpo e documentado
- Binário de apenas 5.5MB
- Performance nativa via CGO

---

[4.0.0]: https://github.com/peder1981/bagus-browser-go/releases/tag/v4.0.0
