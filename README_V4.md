# 🌐 Bagus Browser v4.0

**Browser minimalista, seguro e privado construído em Go**

---

## 🎯 Pilares Fundamentais

### 🔒 Segurança
- Validação rigorosa de URLs
- Sanitização de input
- WebView configurado com segurança máxima
- Plugins e Java desabilitados
- Lista de bloqueio de domínios

### 🕵️ Privacidade  
- **Zero telemetria** (garantido)
- **Zero rastreamento** (garantido)
- Third-party cookies bloqueados
- WebGL/WebAudio bloqueados (anti-fingerprinting)
- DuckDuckGo como motor de busca padrão

### 💪 Robustez
- WebView via CGO (WebKit2GTK)
- Múltiplas abas independentes
- Navegação completa
- 9 atalhos de teclado
- Sem crashes

### 🪶 Leveza
- **5.5MB** binário
- WebKit do sistema (não embarcado)
- Go puro + CGO
- Rápido e eficiente

---

## ✨ Funcionalidades

### Navegação
- ✅ Campo URL com validação
- ✅ Botões: ←, →, ⟳
- ✅ Histórico por aba
- ✅ Busca integrada (DuckDuckGo)
- ✅ Detecção automática URL vs busca

### Abas
- ✅ Múltiplas abas simultâneas
- ✅ Títulos dinâmicos (URL ou título da página)
- ✅ WebView independente por aba
- ✅ Criar/fechar abas
- ✅ Proteção última aba

### Zoom
- ✅ Aumentar (Ctrl++)
- ✅ Diminuir (Ctrl+-)
- ✅ Resetar (Ctrl+0)
- ✅ Independente por aba

---

## ⌨️ Atalhos de Teclado

| Atalho | Ação |
|--------|------|
| **Ctrl+T** | Nova aba |
| **Ctrl+W** | Fechar aba |
| **Alt+←** | Voltar |
| **Alt+→** | Avançar |
| **F5 / Ctrl+R** | Recarregar |
| **Ctrl+L** | Focar URL |
| **Ctrl++** | Aumentar zoom |
| **Ctrl+-** | Diminuir zoom |
| **Ctrl+0** | Resetar zoom |

---

## 🚀 Instalação

### Dependências (Ubuntu/Debian)
```bash
sudo apt-get install -y \
    libgtk-3-dev \
    libwebkit2gtk-4.0-dev \
    pkg-config \
    build-essential
```

### Compilar
```bash
cd /home/peder/bagus-webkit-cgo
go build -o bagus-webkit .
```

### Executar
```bash
./bagus-webkit
```

---

## 📊 Especificações Técnicas

### Stack
- **Linguagem:** Go 1.23+
- **UI:** GTK3 (via gotk3)
- **Renderização:** WebKit2GTK-4.0 (via CGO)
- **Binário:** 5.5MB

### Arquitetura
```
Browser
├── WebView (CGO → WebKit2GTK)
├── Security (Validação + Sanitização)
├── Privacy (Zero telemetria + Anti-fingerprinting)
└── UI (GTK3 + GtkNotebook)
```

### Segurança
- Validação de URLs (HTTP/HTTPS apenas)
- Sanitização de input (XSS protection)
- WebView hardened (plugins/Java desabilitados)
- User-Agent customizado

### Privacidade
- Zero telemetria
- Third-party cookies bloqueados
- WebGL bloqueado (fingerprinting)
- WebAudio bloqueado (fingerprinting)
- DuckDuckGo padrão

---

## 📚 Documentação

- **[SECURITY.md](docs/SECURITY.md)** - Documentação de segurança
- **[PRIVACY.md](docs/PRIVACY.md)** - Política de privacidade
- **[LESSONS_LEARNED.md](LESSONS_LEARNED.md)** - Lições técnicas

---

## 🎯 Roadmap

### v4.0 (Atual) ✅
- [x] WebView funcionando
- [x] Múltiplas abas
- [x] Navegação completa
- [x] Segurança implementada
- [x] Privacidade implementada
- [x] Zoom

### v4.1 (Planejado)
- [ ] Favoritos (Ctrl+D)
- [ ] Buscar na página (Ctrl+F)
- [ ] Downloads
- [ ] Histórico global

### v4.2 (Futuro)
- [ ] Extensões
- [ ] Temas
- [ ] Sincronização (opcional)

---

## 🤝 Contribuindo

Este é um projeto pessoal focado em privacidade e minimalismo.

**Princípios:**
1. **Privacidade primeiro** - Zero telemetria, sempre
2. **Segurança** - Validação rigorosa
3. **Simplicidade** - Código limpo e manutenível
4. **Leveza** - Binário pequeno

---

## 📜 Licença

MIT License - Veja LICENSE para detalhes

---

## 🙏 Agradecimentos

- **WebKit** - Motor de renderização
- **GTK** - Toolkit de UI
- **DuckDuckGo** - Motor de busca privado
- **Go** - Linguagem incrível

---

## 📞 Contato

**Issues:** GitHub Issues  
**Privacidade:** Leia PRIVACY.md  
**Segurança:** Leia SECURITY.md  

---

## 🎊 Status do Projeto

**Status:** ✅ Funcional e estável  
**Versão:** 4.0.0  
**Data:** Outubro 2025  
**Pilares:** 🔒 Segurança | 🕵️ Privacidade | 💪 Robustez | 🪶 Leveza  

---

**Bagus Browser - Navegue com privacidade e segurança** 🌐🔒
