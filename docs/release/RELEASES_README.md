# 📦 Bagus Browser - Downloads

## 🚀 Versão Atual: v1.0.0

**Status:** ✅ Estável e Funcional

---

## 💻 Linux (Recomendado)

### Debian/Ubuntu (.deb)

**Download:** [bagus-browser_1.0.0_amd64.deb](https://github.com/peder1981/bagus-browser-go/releases/latest)

**Instalação:**
```bash
sudo dpkg -i bagus-browser_1.0.0_amd64.deb
sudo apt-get install -f
```

**Uso:**
```bash
bagus-browser
```

Ou procure "Bagus Browser" no menu de aplicativos.

### Requisitos:
- Ubuntu 20.04+ ou Debian 11+
- libwebkit2gtk-4.0-37
- libgtk-3-0

---

## 🪟 Windows

**Status:** ⚠️ Em Desenvolvimento

A versão Windows está em desenvolvimento. Por enquanto, oferecemos:

**Download:** [bagus-browser.exe](https://github.com/peder1981/bagus-browser-go/releases/latest) (placeholder)

Este executável informa sobre o status do desenvolvimento e direciona para a versão Linux.

### Por que não há versão Windows funcional?

A versão atual usa **WebKit2GTK**, que é específico para Linux.

Para Windows, estamos trabalhando em:
- ✨ Integração com Microsoft Edge WebView2
- ✨ Ou CEF (Chromium Embedded Framework)

**Previsão:** v2.0.0 (quando CEF estiver estável)

---

## 🔧 Instalação do Código Fonte

Para qualquer plataforma Linux:

```bash
git clone https://github.com/peder1981/bagus-browser-go.git
cd bagus-browser-go
chmod +x install.sh
./install.sh
```

Pressione ENTER e aguarde 2 minutos.

---

## ✨ Características

| Característica | Status |
|----------------|--------|
| **Tamanho** | 4MB |
| **Inicialização** | Instantânea |
| **Telemetria** | Zero |
| **Rastreamento** | Nenhum |
| **Ads** | Bloqueados |
| **Código** | Aberto |

---

## 🌐 Compatibilidade de Sites

### ✅ Funciona Perfeitamente:
- DuckDuckGo
- Wikipedia
- YouTube
- Reddit
- Stack Overflow
- GitHub
- Hacker News
- E muitos outros

### ⚠️ Compatibilidade Limitada:
- Google (algumas funcionalidades)
- Facebook (algumas funcionalidades)
- Twitter (algumas funcionalidades)

**Nota:** Para 100% de compatibilidade, aguarde a versão CEF (v2.0.0)

---

## 📊 Comparação de Versões

| Versão | Engine | Tamanho | Compatibilidade | Status |
|--------|--------|---------|-----------------|--------|
| **v1.x (Atual)** | WebKit2GTK | 4MB | 70%+ sites | ✅ Estável |
| **v2.x (Futura)** | CEF | ~50MB | 100% sites | 🔜 Em desenvolvimento |

---

## 🔄 Atualizações

### Como atualizar:

**Debian/Ubuntu:**
```bash
sudo dpkg -i bagus-browser_[nova-versao]_amd64.deb
```

**Código Fonte:**
```bash
cd bagus-browser-go
git pull
./install.sh
```

---

## 📝 Notas de Versão

### v1.0.0 (Atual)
- ✨ Primeira versão estável
- ✨ Interface GTK3 nativa
- ✨ Engine WebKit2GTK
- ✨ Bloqueador de ads
- ✨ Gerenciamento de histórico
- ✨ Zero telemetria

### Próximas Versões

**v1.1.0** (Planejado)
- 🔜 Gerenciamento de favoritos
- 🔜 Modo escuro
- 🔜 Melhorias de UI

**v2.0.0** (Futuro)
- 🔜 CEF implementado
- 🔜 100% compatibilidade
- 🔜 Suporte Windows

---

## 🐛 Problemas Conhecidos

### v1.0.0

1. **Alguns sites modernos não funcionam perfeitamente**
   - **Causa:** WebKit2GTK tem limitações
   - **Solução:** Aguardar v2.0.0 com CEF

2. **Sem suporte Windows**
   - **Causa:** WebKit2GTK é Linux-only
   - **Solução:** Em desenvolvimento

---

## 🆘 Suporte

### Encontrou um problema?

1. **Verifique:** [Issues conhecidas](https://github.com/peder1981/bagus-browser-go/issues)
2. **Reporte:** [Abrir nova issue](https://github.com/peder1981/bagus-browser-go/issues/new)

### Precisa de ajuda?

- 📖 [Documentação](README.md)
- 💬 [Discussões](https://github.com/peder1981/bagus-browser-go/discussions)
- 🐛 [Issues](https://github.com/peder1981/bagus-browser-go/issues)

---

## 🔐 Verificação de Integridade

### Checksums SHA256:

```
# Será atualizado a cada release
bagus-browser_1.0.0_amd64.deb: [checksum]
bagus-browser.exe: [checksum]
```

### Como verificar:

```bash
sha256sum bagus-browser_1.0.0_amd64.deb
```

Compare com o checksum oficial no release.

---

## 📜 Licença

MIT License - Veja [LICENSE](LICENSE) para detalhes.

---

## 🙏 Contribuindo

Quer ajudar? Veja [CONTRIBUTING.md](CONTRIBUTING.md)

Áreas que precisam de ajuda:
- 🔧 Implementação CEF
- 🪟 Suporte Windows
- 🎨 Melhorias de UI
- 📝 Documentação
- 🧪 Testes

---

## ⭐ Apoie o Projeto

Se você gosta do Bagus Browser:

- ⭐ Dê uma estrela no GitHub
- 🐛 Reporte bugs
- 💡 Sugira melhorias
- 🔧 Contribua com código
- 📢 Compartilhe com amigos

---

**Obrigado por usar o Bagus Browser! 🌐✨**
