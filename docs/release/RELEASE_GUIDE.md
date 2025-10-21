# 📦 Guia de Release - Bagus Browser

## 🚀 Como Criar um Release no GitHub

### 1. Preparar os Pacotes

```bash
# Criar pacote .deb (Linux)
./build-deb.sh

# Criar executável Windows (placeholder)
./build-windows.sh
```

### 2. Arquivos Gerados

Após executar os scripts, você terá:

```
build/
├── bagus-browser_1.0.0_amd64.deb    # Pacote Debian/Ubuntu
└── windows/
    └── bagus-browser.exe             # Executável Windows (placeholder)
```

### 3. Criar Release no GitHub

#### Via Interface Web:

1. Acesse: `https://github.com/peder1981/bagus-browser-go/releases/new`

2. Preencha:
   - **Tag version:** `v1.0.0`
   - **Release title:** `Bagus Browser v1.0.0 - Primeira Versão Estável`

3. **Descrição do Release:**

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
- **Arquivo:** `bagus-browser_1.0.0_amd64.deb`
- **Instalação:**
  ```bash
  sudo dpkg -i bagus-browser_1.0.0_amd64.deb
  sudo apt-get install -f
  ```
- **Uso:** `bagus-browser` ou procure "Bagus Browser" no menu

### Windows
- **Status:** ⚠️ Em desenvolvimento
- **Arquivo:** `bagus-browser.exe` (placeholder informativo)
- **Nota:** A versão funcional está disponível apenas para Linux

## 🌐 Sites Compatíveis

✅ DuckDuckGo • Wikipedia • YouTube • Reddit • Stack Overflow • GitHub • E muitos outros

## 📋 Requisitos

### Linux
- Ubuntu 20.04+ / Debian 11+
- libwebkit2gtk-4.0-37
- libgtk-3-0

## 🔧 Instalação Manual (Código Fonte)

```bash
git clone https://github.com/peder1981/bagus-browser-go.git
cd bagus-browser-go
chmod +x install.sh
./install.sh
```

## 📝 Notas da Versão

### Novidades
- ✨ Primeira versão estável
- ✨ Interface GTK3 nativa
- ✨ Engine WebKit2GTK
- ✨ Bloqueador de ads integrado
- ✨ Gerenciamento de histórico
- ✨ Sistema de autenticação local

### Limitações Conhecidas
- ⚠️ Alguns sites modernos (Google, Facebook, Twitter) podem não funcionar perfeitamente
- ⚠️ Versão Windows em desenvolvimento (CEF)

### Próximas Versões
- 🔜 Implementação CEF para melhor compatibilidade
- 🔜 Suporte a Windows
- 🔜 Gerenciamento de favoritos
- 🔜 Modo escuro

## 🐛 Reportar Problemas

Encontrou um bug? [Abra uma issue](https://github.com/peder1981/bagus-browser-go/issues/new)

## 📚 Documentação

- [README](README.md) - Visão geral
- [CEF_STATUS](CEF_STATUS.md) - Status da implementação CEF
- [RESULTADO_FINAL](RESULTADO_FINAL.md) - Especificações técnicas

## 🙏 Agradecimentos

Obrigado por usar o Bagus Browser!

---

**Checksums (SHA256):**
```
# Será gerado automaticamente após upload
```
```

4. **Anexar Arquivos:**
   - Arraste `build/bagus-browser_1.0.0_amd64.deb`
   - Arraste `build/windows/bagus-browser.exe`

5. **Marcar como:**
   - ☑️ Set as the latest release
   - ☐ Set as a pre-release

6. Clique em **Publish release**

---

## 📝 Checklist de Release

Antes de publicar, verifique:

- [ ] Código compilando sem erros
- [ ] Testes básicos executados
- [ ] Pacote .deb criado
- [ ] Executável Windows criado
- [ ] README.md atualizado
- [ ] CEF_STATUS.md atualizado
- [ ] CHANGELOG.md criado (se aplicável)
- [ ] Tag de versão criada
- [ ] Release notes escritas
- [ ] Arquivos anexados ao release

---

## 🔄 Versionamento

Seguimos [Semantic Versioning](https://semver.org/):

- **MAJOR.MINOR.PATCH** (ex: 1.0.0)
  - **MAJOR**: Mudanças incompatíveis
  - **MINOR**: Novas funcionalidades compatíveis
  - **PATCH**: Correções de bugs

### Exemplos:
- `v1.0.0` - Primeira versão estável
- `v1.1.0` - Nova funcionalidade (ex: favoritos)
- `v1.0.1` - Correção de bug
- `v2.0.0` - Mudança grande (ex: CEF implementado)

---

## 📊 Métricas de Release

Após publicar, monitore:

- Downloads por plataforma
- Issues reportadas
- Feedback dos usuários
- Performance em diferentes sistemas

---

## 🔐 Assinatura de Releases (Opcional)

Para releases assinados:

```bash
# Gerar checksum
sha256sum build/bagus-browser_1.0.0_amd64.deb > checksums.txt

# Assinar com GPG
gpg --armor --detach-sign checksums.txt

# Anexar checksums.txt e checksums.txt.asc ao release
```

---

## 📢 Divulgação

Após o release:

1. **Anunciar em:**
   - README do repositório
   - Twitter/X
   - Reddit (r/linux, r/opensource)
   - Fóruns Linux

2. **Atualizar:**
   - Site do projeto (se houver)
   - Documentação
   - Wiki

---

## 🆘 Troubleshooting

### Erro ao criar .deb
```bash
# Instalar ferramentas necessárias
sudo apt-get install dpkg-dev
```

### Erro ao cross-compile Windows
```bash
# Instalar mingw-w64
sudo apt-get install mingw-w64
```

### Arquivo muito grande
```bash
# Comprimir antes de fazer upload
gzip build/bagus-browser_1.0.0_amd64.deb
```

---

## ✅ Exemplo de Release Completo

```bash
# 1. Atualizar versão no código
VERSION="1.0.0"

# 2. Criar pacotes
./build-deb.sh
./build-windows.sh

# 3. Criar tag
git tag -a v1.0.0 -m "Release v1.0.0"
git push origin v1.0.0

# 4. Criar release no GitHub (via web)

# 5. Anunciar
echo "Release v1.0.0 publicado!"
```

---

## 📅 Cronograma Sugerido

- **v1.0.0** - Versão inicial (WebView)
- **v1.1.0** - Favoritos e melhorias UI
- **v1.2.0** - Modo escuro
- **v2.0.0** - CEF implementado (quando estável)
- **v2.1.0** - Suporte Windows completo

---

**Boa sorte com o release! 🚀**
