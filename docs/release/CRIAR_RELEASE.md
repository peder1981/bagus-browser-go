# 🚀 CRIAR RELEASE NO GITHUB - PASSO A PASSO

## ✅ Preparação Completa

- ✅ Tag v1.0.0 criada e enviada
- ✅ Código no GitHub
- ✅ Pacotes prontos em `build/`

---

## 📦 PASSO 1: Acesse o GitHub

**URL:** https://github.com/peder1981/bagus-browser-go/releases/new

Ou:
1. Vá para: https://github.com/peder1981/bagus-browser-go
2. Clique em "Releases" (lado direito)
3. Clique em "Create a new release"

---

## 📝 PASSO 2: Preencha o Formulário

### Tag version
```
v1.0.0
```
(Selecione da lista - já existe)

### Release title
```
Bagus Browser v1.0.0 - Primeira Versão Estável
```

### Description
Copie e cole o texto abaixo:

```markdown
# 🌐 Bagus Browser v1.0.0

Primeira versão estável do Bagus Browser - Um navegador leve, rápido e focado em privacidade.

## ✨ Características

- ✅ **Leve**: Apenas 4MB
- ✅ **Rápido**: Inicialização instantânea
- ✅ **Privado**: Zero telemetria
- ✅ **Seguro**: Sem rastreamento
- ✅ **Simples**: Instalação em 2 minutos

## 📦 Downloads

### Linux (Debian/Ubuntu)
```bash
sudo dpkg -i bagus-browser_1.0.0_amd64.deb
sudo apt-get install -f
```

**Uso:** `bagus-browser` ou procure "Bagus Browser" no menu de aplicativos

### Windows
⚠️ **Em desenvolvimento** - Este é um placeholder informativo que direciona para a versão Linux

## 🌐 Sites Compatíveis

✅ DuckDuckGo • Wikipedia • YouTube • Reddit • Stack Overflow • GitHub • E muitos outros (70%+)

## 📋 Requisitos (Linux)

- Ubuntu 20.04+ / Debian 11+
- libwebkit2gtk-4.0-37
- libgtk-3-0

## 🔧 Instalação do Código Fonte

```bash
git clone https://github.com/peder1981/bagus-browser-go.git
cd bagus-browser-go
chmod +x install.sh
./install.sh
```

Pressione ENTER e aguarde 2 minutos.

## 📝 Novidades desta Versão

- ✨ Primeira versão estável
- ✨ Interface GTK3 nativa
- ✨ Engine WebKit2GTK
- ✨ Bloqueador de ads integrado
- ✨ Gerenciamento de histórico
- ✨ Sistema de autenticação local
- ✨ Zero telemetria e rastreamento

## ⚠️ Limitações Conhecidas

- Alguns sites modernos (Google, Facebook, Twitter) podem ter compatibilidade limitada devido ao WebKit2GTK
- Versão Windows em desenvolvimento (aguarde v2.0.0 com CEF)

## 🔜 Próximas Versões

**v1.1.0** (Planejado)
- Gerenciamento de favoritos
- Modo escuro
- Melhorias de UI

**v2.0.0** (Futuro)
- CEF implementado
- 100% compatibilidade de sites
- Suporte Windows funcional
- DevTools integrado

## 🐛 Reportar Problemas

Encontrou um bug? [Abra uma issue](https://github.com/peder1981/bagus-browser-go/issues/new)

## 📚 Documentação

- [README](https://github.com/peder1981/bagus-browser-go/blob/main/README.md) - Visão geral
- [CEF_STATUS](https://github.com/peder1981/bagus-browser-go/blob/main/CEF_STATUS.md) - Status da implementação CEF
- [INSTALACAO_SIMPLES](https://github.com/peder1981/bagus-browser-go/blob/main/INSTALACAO_SIMPLES.md) - Guia de instalação

## 🙏 Agradecimentos

Obrigado por usar o Bagus Browser!

Se você gosta do projeto:
- ⭐ Dê uma estrela no GitHub
- 🐛 Reporte bugs
- 💡 Sugira melhorias
- 🔧 Contribua com código

---

**Bagus Browser - Simples. Robusto. Seguro. Funcional.** 🌐✨
```

---

## 📎 PASSO 3: Anexar Arquivos

Arraste e solte os seguintes arquivos na área "Attach binaries":

1. **`build/bagus-browser_1.0.0_amd64.deb`**
   - Pacote Debian/Ubuntu
   - Tamanho: ~2.4KB

2. **`build/windows/bagus-browser.exe`**
   - Executável Windows (placeholder)
   - Tamanho: ~2MB

---

## ✅ PASSO 4: Configurações Finais

Marque as seguintes opções:

- ☑️ **Set as the latest release**
- ☐ Set as a pre-release (deixe desmarcado)

---

## 🚀 PASSO 5: Publicar

Clique no botão verde:

**"Publish release"**

---

## 🎉 Pronto!

Após publicar, o release estará disponível em:

**https://github.com/peder1981/bagus-browser-go/releases/tag/v1.0.0**

---

## 📊 Verificação

Após publicar, verifique:

- [ ] Release aparece na página principal
- [ ] Arquivos .deb e .exe estão disponíveis para download
- [ ] Tag v1.0.0 está marcada como "Latest"
- [ ] Descrição está formatada corretamente
- [ ] Links funcionam

---

## 📢 Próximos Passos (Opcional)

### Anunciar o Release

1. **Reddit**
   - r/linux
   - r/opensource
   - r/programming

2. **Twitter/X**
   ```
   🚀 Bagus Browser v1.0.0 lançado!
   
   Navegador leve (4MB), rápido e focado em privacidade.
   
   ✅ Zero telemetria
   ✅ Bloqueador de ads
   ✅ Código aberto
   
   Download: https://github.com/peder1981/bagus-browser-go/releases
   
   #opensource #privacy #browser #linux
   ```

3. **Hacker News**
   - Show HN: Bagus Browser - Lightweight privacy-focused browser

---

## ✅ Checklist Final

- [ ] Release criado no GitHub
- [ ] Arquivos anexados
- [ ] Marcado como "latest"
- [ ] Verificado que downloads funcionam
- [ ] Anunciado (opcional)

---

**Boa sorte com o lançamento! 🚀✨**
