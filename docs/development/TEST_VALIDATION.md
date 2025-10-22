# 🧪 Validação Completa do Bagus Browser

## ✅ Checklist de Validação

### 1. Título da Janela
- [ ] Título deve ser "Bagus Browser"
- [ ] Verificar na barra de título
- [ ] Verificar no gerenciador de janelas

### 2. Navegação Básica
- [ ] Abrir browser (deve iniciar com DuckDuckGo)
- [ ] Digitar "github.com" e pressionar Enter
- [ ] Verificar se carrega corretamente
- [ ] Digitar "youtube.com" e pressionar Enter
- [ ] Verificar se carrega corretamente

### 3. Atalhos de Abas
- [ ] **Ctrl+T** - Nova aba (deve abrir DuckDuckGo)
- [ ] **Ctrl+W** - Fechar aba (não deve fechar se for a última)
- [ ] Abrir 3 abas e fechar a do meio
- [ ] Verificar que as outras abas continuam funcionando

### 4. Atalhos de Navegação
- [ ] **Alt+←** - Voltar (navegar em 3 sites e voltar)
- [ ] **Alt+→** - Avançar (após voltar, avançar novamente)
- [ ] **F5** - Recarregar página
- [ ] **Ctrl+R** - Recarregar página (alternativo)
- [ ] **Ctrl+L** - Focar barra de URL (deve selecionar todo o texto)

### 5. Atalhos de Zoom
- [ ] **Ctrl++** - Aumentar zoom (verificar visualmente)
- [ ] **Ctrl+-** - Diminuir zoom (verificar visualmente)
- [ ] **Ctrl+0** - Resetar zoom para 100%
- [ ] Verificar logs no terminal mostrando porcentagem

### 6. Busca na Página
- [ ] **Ctrl+F** - Abrir busca
- [ ] Digitar texto para buscar
- [ ] Verificar que destaca na página
- [ ] **F3** - Próximo resultado
- [ ] **Shift+F3** - Resultado anterior
- [ ] **Esc** - Fechar busca

### 7. Favoritos
- [ ] Navegar para um site (ex: github.com)
- [ ] **Ctrl+D** - Adicionar favorito
- [ ] Verificar mensagem de confirmação
- [ ] **Ctrl+D** novamente - Deve avisar que já existe
- [ ] **Ctrl+Shift+B** - Gerenciar favoritos (CRÍTICO - estava quebrado)
- [ ] Verificar que abre janela de favoritos
- [ ] Clicar em "Abrir" - Deve abrir em nova aba
- [ ] Clicar em "Remover" - Deve remover da lista
- [ ] Fechar janela de favoritos

### 8. Privacidade
- [ ] Verificar logs ao iniciar:
  - ✅ Zero telemetria
  - ✅ Third-party cookies bloqueados
  - ✅ WebGL bloqueado
  - ✅ WebAudio bloqueado
  - ✅ DuckDuckGo como motor de busca

### 9. Múltiplas Abas Simultâneas
- [ ] Abrir 5 abas com sites diferentes
- [ ] Alternar entre abas
- [ ] Verificar que cada aba mantém seu conteúdo
- [ ] Verificar que URL na barra muda ao trocar de aba
- [ ] Fechar abas aleatoriamente
- [ ] Verificar que não há crashes

### 10. Validação de URL
- [ ] Digitar URL sem protocolo (ex: "google.com")
- [ ] Verificar que adiciona "https://" automaticamente
- [ ] Digitar termo de busca (ex: "golang tutorial")
- [ ] Verificar que busca no DuckDuckGo
- [ ] Digitar URL inválida
- [ ] Verificar tratamento de erro

### 11. Persistência de Favoritos
- [ ] Adicionar 3 favoritos
- [ ] Fechar browser
- [ ] Abrir browser novamente
- [ ] **Ctrl+Shift+B** - Verificar que favoritos foram carregados
- [ ] Verificar arquivo: `~/.config/bagus-browser/bookmarks.enc`

### 12. Interface
- [ ] Botões da toolbar funcionam:
  - [ ] ← (Voltar)
  - [ ] → (Avançar)
  - [ ] ⟳ (Recarregar)
  - [ ] + (Nova aba)
- [ ] Campo URL aceita entrada
- [ ] Tooltips aparecem ao passar mouse

## 📊 Resultado Final

**Total de Testes:** 50+

**Aprovados:** _____

**Falhados:** _____

**Críticos Corrigidos:**
- ✅ Título "POC" removido
- ✅ Ctrl+Shift+B (Gerenciar Favoritos) corrigido

## 🐛 Bugs Encontrados

_(Listar aqui qualquer bug encontrado durante os testes)_

---

## 📝 Notas de Teste

_(Adicionar observações durante os testes)_
