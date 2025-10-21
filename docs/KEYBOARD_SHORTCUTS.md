# ⌨️ Atalhos de Teclado - Bagus Browser

## 🎯 Navegação Principal

### Ctrl+L - Navegar para URL
Abre um diálogo para digitar uma nova URL ou termo de busca.

**Uso:**
1. Pressione `Ctrl+L`
2. Digite uma URL (ex: `github.com`) ou termo de busca (ex: `golang tutorial`)
3. Pressione Enter

**Exemplos:**
- `google.com` → https://google.com
- `golang tutorial` → Busca no DuckDuckGo
- `https://exemplo.com/pagina` → URL completa

### Alt+← - Voltar
Volta para a página anterior no histórico.

### Alt+→ - Avançar
Avança para a próxima página no histórico.

### Ctrl+R ou F5 - Recarregar
Recarrega a página atual.

### Ctrl+H - Ver Histórico
Mostra o histórico de navegação no console do navegador.

**Nota:** Abra o console com F12 para ver os resultados.

## 💻 Navegação via Console JavaScript

Você também pode navegar programaticamente usando o console do navegador (F12):

```javascript
// Navegar para uma URL
navigateGo("https://github.com")

// Fazer uma busca
navigateGo("golang tutorial")

// Ver histórico
searchHistoryGo("")
```

## 🔍 Tabela de Atalhos

| Atalho | Ação | Descrição |
|--------|------|-----------|
| `Ctrl+L` | Navegar | Abre diálogo para digitar URL |
| `Ctrl+R` | Recarregar | Recarrega a página atual |
| `F5` | Recarregar | Recarrega a página atual |
| `Alt+←` | Voltar | Volta para página anterior |
| `Alt+→` | Avançar | Avança para próxima página |
| `Ctrl+H` | Histórico | Mostra histórico no console |
| `F12` | DevTools | Abre console do navegador |

## 🎨 Mensagem de Boas-Vindas

Ao abrir o Bagus Browser, você verá uma mensagem de boas-vindas no console (F12) com todos os atalhos disponíveis.

Para ver a mensagem:
1. Pressione `F12` para abrir o console
2. Veja a lista completa de atalhos e comandos

## 📝 Dicas de Uso

### Navegação Rápida
```
Ctrl+L → Digite URL → Enter
```

### Voltar e Avançar
```
Alt+← (voltar)
Alt+→ (avançar)
```

### Recarregar Página
```
F5 ou Ctrl+R
```

### Ver Histórico
```
Ctrl+H → F12 (ver console)
```

## 🔧 Configuração

Os atalhos de teclado são injetados automaticamente via JavaScript e funcionam em todas as páginas.

**Limitações:**
- Alguns sites podem interceptar atalhos (ex: Gmail, Google Docs)
- Nesses casos, use o console JavaScript como alternativa

## 🚀 Navegação Avançada

### Via Código JavaScript

Abra o console (F12) e use:

```javascript
// Navegar
navigateGo("https://exemplo.com")

// Processar input (URL ou busca)
processInputGo("golang tutorial")

// Buscar no histórico
searchHistoryGo("github")
```

### Múltiplas Janelas

Para abrir múltiplas janelas:

```bash
# Terminal 1
./build/bagus

# Terminal 2
./build/bagus

# Cada janela é independente
```

## 📚 Próximas Versões

### v2.1.0 (Planejado)
- [ ] `Ctrl+T` - Nova aba/janela
- [ ] `Ctrl+W` - Fechar janela
- [ ] `Ctrl+Tab` - Alternar entre abas
- [ ] `Ctrl++` - Aumentar zoom
- [ ] `Ctrl+-` - Diminuir zoom
- [ ] `Ctrl+0` - Resetar zoom

### v2.2.0 (Planejado)
- [ ] `Ctrl+F` - Buscar na página
- [ ] `Ctrl+D` - Adicionar favorito
- [ ] `Ctrl+Shift+B` - Mostrar favoritos
- [ ] `Ctrl+Shift+Delete` - Limpar dados

## 🐛 Problemas Conhecidos

### Atalhos não funcionam em alguns sites

**Causa:** Alguns sites interceptam eventos de teclado.

**Solução:** Use o console JavaScript:
```javascript
navigateGo("https://outro-site.com")
```

### Diálogo de navegação não abre

**Causa:** Site está bloqueando `prompt()`.

**Solução:** Use o console JavaScript diretamente.

## 💡 Sugestões

Tem sugestões de novos atalhos? Abra uma issue no GitHub!

---

**Última Atualização:** 2025-10-21  
**Versão:** 2.0.0  
**Status:** Ativo
