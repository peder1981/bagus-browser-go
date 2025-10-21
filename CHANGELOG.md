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

## [Unreleased] - v4.1

### Planejado
- Favoritos (Ctrl+D)
- Buscar na página (Ctrl+F)
- Downloads
- Histórico global

---

## [Unreleased] - v4.2

### Planejado
- Melhorias de UI
- Ícone do aplicativo
- Favicon nas abas
- Indicador de carregamento
- Botão X em cada aba

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
