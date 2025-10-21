# 🚀 Guia Rápido - Bagus Browser

## Qual versão usar?

### 🏃 Versão Rápida (Recomendado)
```bash
./install.sh  # Escolha opção 1
bagus
```
- ✅ 2 minutos de instalação
- ✅ 70%+ dos sites (DuckDuckGo, Wikipedia, YouTube)
- ⚠️ Google, Facebook não funcionam

### 🌐 Versão Completa (100% Sites)
```bash
./install.sh  # Escolha opção 2 ou 3
bagus --full
```
- ✅ 100% dos sites (Google, Facebook, Twitter)
- ⚠️ 40 minutos de instalação

---

## Instalação Rápida

### 1. Instalar

```bash
./install.sh
```

### 2. Executar

```bash
bagus
# ou
./build/bagus
```

### 3. Login

1. Digite um username válido (ex: `peder`)
2. Clique em "Iniciar Browser"

**Requisitos do username:**
- 3-32 caracteres
- Apenas letras, números, `_` ou `-`
- Sem espaços ou caracteres especiais

### 4. Navegar

1. Digite uma URL na barra de endereço
2. Pressione Enter
3. Use os botões ◀ ▶ ⟳ para navegar

### 5. Abas

- **Nova aba**: Clique em "+ Nova Aba"
- **Trocar aba**: Clique na aba desejada
- **Fechar aba**: (implementar Ctrl+W)

---

## 📁 Estrutura de Dados

Seus dados ficam em `/tmp/{username}/`:

```
/tmp/peder/
├── config.json       # Configurações do browser
├── history.json      # Histórico de navegação
├── default/          # Perfil padrão
├── log/              # Logs
├── cache/            # Cache
└── downloads/        # Downloads
```

---

## 🔧 Configuração

Edite `/tmp/{username}/config.json` para personalizar:

```json
{
  "default": {
    "url": "https://duckduckgo.com/"
  },
  "settings": {
    "JavascriptEnabled": true,
    "AutoLoadImages": true,
    "ForceDarkMode": false,
    ...
  }
}
```

---

## 🛡️ Bloqueador de Domínios

Crie `/tmp/{username}/ad_hosts_block.txt` com domínios a bloquear:

```
ads.example.com
tracker.example.com
analytics.example.com
```

Um domínio por linha. Linhas começando com `#` são ignoradas.

---

## ⌨️ Atalhos (Planejados)

- `Ctrl+T` - Nova aba
- `Ctrl+W` - Fechar aba
- `Ctrl+R` - Recarregar
- `Ctrl+L` - Focar barra de URL
- `Ctrl+Q` - Sair

---

## 🧪 Testar

```bash
# Todos os testes
go test ./...

# Com cobertura
go test -cover ./...

# Verbose
go test -v ./...
```

---

## 🐛 Troubleshooting

### Erro: "Username inválido"
- Verifique se o username tem 3-32 caracteres
- Use apenas letras, números, `_` ou `-`

### Erro: "Erro ao criar browser"
- Verifique permissões em `/tmp`
- Certifique-se que webview está instalado

### Webview não funciona
**Linux**: Instale webkit2gtk
```bash
sudo apt install libwebkit2gtk-4.0-dev  # Debian/Ubuntu
sudo dnf install webkit2gtk3-devel      # Fedora
```

**Windows**: WebView2 já incluído no Windows 10+

**macOS**: WKWebView já incluído

---

## 📝 Logs

Logs ficam em `/tmp/{username}/log/`:
- `block.log` - Domínios bloqueados

---

## 🚀 Build para Produção

```bash
# Linux
CGO_ENABLED=1 go build -ldflags="-s -w" -o bagus-browser main.go

# Windows (cross-compile)
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 go build -o bagus-browser.exe main.go

# macOS (cross-compile)
GOOS=darwin GOARCH=amd64 CGO_ENABLED=1 go build -o bagus-browser-mac main.go
```

---

## 📚 Mais Informações

- `IMPLEMENTATION_COMPLETE.md` - Documentação completa
- `README.md` - Visão geral
- `docs/` - Documentação técnica

---

**Pronto para navegar! 🌐**
