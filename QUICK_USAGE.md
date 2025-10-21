# 🚀 Guia Rápido de Uso - Bagus Browser

## ▶️ Iniciar o Browser

```bash
./build/bagus
```

O browser abrirá automaticamente no DuckDuckGo.

## 🌐 Como Navegar

### Método 1: Atalho de Teclado (Recomendado)

1. Pressione **`Ctrl+L`**
2. Digite a URL ou termo de busca
3. Pressione **Enter**

**Exemplos:**
- `github.com` → Abre GitHub
- `golang tutorial` → Busca no DuckDuckGo

### Método 2: Console JavaScript

1. Pressione **`F12`** para abrir o console
2. Digite: `navigateGo("github.com")`
3. Pressione **Enter**

### Método 3: Links na Página

Clique normalmente em qualquer link na página atual.

## ⌨️ Atalhos Principais

| Atalho | Ação |
|--------|------|
| `Ctrl+L` | Navegar para nova URL |
| `Alt+←` | Voltar |
| `Alt+→` | Avançar |
| `F5` | Recarregar |
| `F12` | Abrir console |

## 💡 Dicas Rápidas

### Navegar Rapidamente
```
Ctrl+L → Digite URL → Enter
```

### Ver Atalhos Disponíveis
```
F12 (abre console)
```
Você verá uma mensagem de boas-vindas com todos os atalhos.

### Navegar via Código
```javascript
// No console (F12)
navigateGo("https://exemplo.com")
```

## 🔍 Buscar vs Navegar

O browser detecta automaticamente se você digitou uma URL ou termo de busca:

**URLs (navega diretamente):**
- `google.com`
- `github.com/user/repo`
- `https://exemplo.com`

**Buscas (usa DuckDuckGo):**
- `golang tutorial`
- `como fazer X`
- `notícias brasil`

## 🐛 Problemas Comuns

### "Atalho não funciona"

**Causa:** Alguns sites interceptam atalhos (Gmail, Google Docs, etc.)

**Solução:** Use o console JavaScript:
```javascript
navigateGo("https://outro-site.com")
```

### "Página em branco"

**Causa:** Site pode estar bloqueando o carregamento.

**Solução:** 
1. Pressione `F5` para recarregar
2. Ou tente outro site

### "Como voltar para página anterior?"

```
Alt+← (voltar)
Alt+→ (avançar)
```

## 📚 Documentação Completa

- [Todos os Atalhos](docs/KEYBOARD_SHORTCUTS.md)
- [Limitações Conhecidas](docs/KNOWN_LIMITATIONS.md)
- [Guia de Instalação](docs/install/INSTALL.md)

## 🎯 Fluxo de Uso Típico

```
1. Abrir browser
   ./build/bagus

2. Navegar
   Ctrl+L → Digite URL → Enter

3. Navegar em links
   Clique normalmente

4. Voltar
   Alt+←

5. Nova navegação
   Ctrl+L → Nova URL → Enter
```

## 🔧 Comandos Avançados

### Ver Histórico
```
Ctrl+H (depois F12 para ver console)
```

### Buscar no Histórico
```javascript
// No console (F12)
searchHistoryGo("github")
```

### Processar Input
```javascript
// No console (F12)
processInputGo("golang tutorial")
```

## 💻 Múltiplas Janelas

Para abrir múltiplas janelas:

```bash
# Terminal 1
./build/bagus

# Terminal 2  
./build/bagus

# Cada janela é independente
```

## 🎨 Personalização

### Alterar URL Inicial

Edite `~/.config/bagus-browser/config.json`:

```json
{
  "default": {
    "url": "https://seu-site-favorito.com"
  }
}
```

## 📊 Status

- ✅ **Funcional**: 100% dos sites carregam
- ✅ **Atalhos**: Todos funcionando
- ✅ **Navegação**: Via teclado e código
- ⚠️ **Barra visual**: Não disponível (limitação técnica)

## 🚀 Próxima Versão (v2.1.0)

- Janela de controle separada
- Mais atalhos
- Melhor UX

---

**Versão:** 2.0.0  
**Última Atualização:** 2025-10-21

**Dúvidas?** Abra uma issue no GitHub!
