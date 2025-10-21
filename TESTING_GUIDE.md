# 🧪 Guia de Testes - Bagus Browser v4.1

## 🚀 Como Executar

```bash
cd /home/peder/bagus-webkit-cgo
./bagus-webkit
```

---

## ✅ CHECKLIST DE TESTES

### 1. Navegação Básica
- [ ] Browser abre com DuckDuckGo
- [ ] Digite "github.com" e pressione Enter
- [ ] Página carrega corretamente
- [ ] Título da aba atualiza
- [ ] URL atualiza no campo

### 2. Múltiplas Abas
- [ ] Pressione Ctrl+T (nova aba)
- [ ] Digite "google.com"
- [ ] Pressione Ctrl+T novamente
- [ ] Digite "wikipedia.org"
- [ ] Alterne entre abas
- [ ] Feche aba com Ctrl+W
- [ ] Tente fechar última aba (deve avisar)

### 3. Navegação (Histórico)
- [ ] Navegue: google.com → github.com → youtube.com
- [ ] Pressione Alt+← (voltar)
- [ ] Pressione Alt+→ (avançar)
- [ ] Verifique se funciona

### 4. Zoom
- [ ] Pressione Ctrl++ (aumentar)
- [ ] Pressione Ctrl+- (diminuir)
- [ ] Pressione Ctrl+0 (resetar)
- [ ] Verifique logs mostrando porcentagem

### 5. Buscar na Página
- [ ] Navegue para wikipedia.org
- [ ] Pressione Ctrl+F
- [ ] Digite "linux"
- [ ] Verifique highlight
- [ ] Pressione F3 (próximo)
- [ ] Pressione Shift+F3 (anterior)
- [ ] Pressione Esc ou feche dialog

### 6. Favoritos
- [ ] Navegue para github.com
- [ ] Pressione Ctrl+D
- [ ] Verifique mensagem de confirmação
- [ ] Pressione Ctrl+D novamente (deve avisar que já existe)
- [ ] Pressione Ctrl+Shift+B
- [ ] Verifique lista de favoritos
- [ ] Clique em "Abrir" (deve abrir em nova aba)
- [ ] Clique em "Remover"
- [ ] Verifique que favorito foi removido

### 7. Favoritos Criptografados
- [ ] Adicione alguns favoritos
- [ ] Feche o browser
- [ ] Verifique arquivo: `~/.config/bagus-browser/bookmarks.enc`
- [ ] Tente abrir com editor de texto (deve estar criptografado)
- [ ] Abra o browser novamente
- [ ] Pressione Ctrl+Shift+B
- [ ] Verifique que favoritos foram carregados

### 8. Downloads
- [ ] Navegue para um site com download
- [ ] Inicie um download
- [ ] Verifique pasta ~/Downloads
- [ ] Verifique logs mostrando pasta

### 9. Privacidade
- [ ] Verifique logs ao iniciar
- [ ] Confirme mensagens de privacidade
- [ ] Verifique "Zero telemetria"
- [ ] Verifique "Third-party cookies bloqueados"

### 10. Segurança
- [ ] Digite "javascript:alert('test')" na URL
- [ ] Deve ser bloqueado (protocolo inválido)
- [ ] Digite URL com espaços/caracteres especiais
- [ ] Deve ser sanitizado

### 11. Atalhos de Teclado
- [ ] Ctrl+T - Nova aba ✓
- [ ] Ctrl+W - Fechar aba ✓
- [ ] Alt+← - Voltar ✓
- [ ] Alt+→ - Avançar ✓
- [ ] F5 - Recarregar ✓
- [ ] Ctrl+R - Recarregar ✓
- [ ] Ctrl+L - Focar URL ✓
- [ ] Ctrl++ - Zoom in ✓
- [ ] Ctrl+- - Zoom out ✓
- [ ] Ctrl+0 - Reset zoom ✓
- [ ] Ctrl+F - Buscar ✓
- [ ] F3 - Próximo resultado ✓
- [ ] Shift+F3 - Resultado anterior ✓
- [ ] Ctrl+D - Adicionar favorito ✓
- [ ] Ctrl+Shift+B - Gerenciar favoritos ✓

---

## 🐛 PROBLEMAS CONHECIDOS

### Para Reportar:
1. Descreva o problema
2. Passos para reproduzir
3. Comportamento esperado
4. Comportamento atual
5. Logs relevantes

---

## ✅ RESULTADO ESPERADO

Todos os testes devem passar sem crashes ou erros críticos.

---

**Execute os testes e reporte qualquer problema!**
