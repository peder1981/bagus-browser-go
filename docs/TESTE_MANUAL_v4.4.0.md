# 🧪 Teste Manual v4.4.0 - Bagus Browser

## ⚠️ IMPORTANTE: Execute TODOS os testes

Por favor, execute cada teste abaixo e marque o resultado.

---

## 📋 Preparação

```bash
# 1. Compilar versão mais recente
cd ~/bagus-browser-go
make build

# 2. Instalar
make install

# 3. Abrir terminal para ver logs
# Terminal 1: Executar browser
bagus-browser

# Terminal 2: Ver logs em tempo real
# (Os logs aparecem no Terminal 1)
```

---

## 🧪 Teste 1: Ctrl+T (Foco Imediato)

**Objetivo:** Verificar se cursor vai imediatamente para URL entry

**Passos:**
1. Abrir Bagus Browser
2. Pressionar `Ctrl+T`
3. **IMEDIATAMENTE** digitar "teste"
4. Verificar onde o texto apareceu

**Resultado Esperado:**
- [ ] Texto "teste" aparece na URL entry (barra de endereço)
- [ ] Cursor está piscando na URL entry
- [ ] Texto está selecionado (azul)

**Resultado Atual:**
- [ ] ✅ Funcionou
- [ ] ❌ Não funcionou - Descrever: _________________

**Logs Esperados:**
```
⌨️  Ctrl+T - Nova aba
✅ Aba 2 criada - Carregando: https://duckduckgo.com
```

---

## 🧪 Teste 2: Ctrl+Shift+T (Reabrir Aba)

**Objetivo:** Verificar se reabre última aba fechada

**Passos:**
1. Abrir 3 abas diferentes:
   - Aba 1: google.com
   - Aba 2: github.com
   - Aba 3: duckduckgo.com

2. Fechar Aba 3 (Ctrl+W)
3. Fechar Aba 2 (Ctrl+W)
4. Pressionar `Ctrl+Shift+T`
5. Verificar qual aba abriu

**Resultado Esperado:**
- [ ] Aba 2 (github.com) reabre primeiro
- [ ] Pressionar Ctrl+Shift+T novamente
- [ ] Aba 3 (duckduckgo.com) reabre

**Resultado Atual:**
- [ ] ✅ Funcionou
- [ ] ❌ Não funcionou - Descrever: _________________

**Logs Esperados:**
```
🔍 DEBUG: keyVal=116 (t=116, T=84), Ctrl=true, Shift=true
⌨️  Ctrl+Shift+T - Reabrir última aba fechada
♻️  Reabrindo aba: GitHub
```

**⚠️ ATENÇÃO:** Se aparecer apenas o log DEBUG mas NÃO aparecer "Reabrir última aba fechada", copie o log completo!

---

## 🧪 Teste 3: Copy/Paste de Screenshot

**Objetivo:** Verificar se consegue colar screenshot

**Passos:**
1. Tirar screenshot (Print Screen)
2. Abrir Gmail no Bagus Browser
3. Fazer login
4. Clicar em "Compose" (Escrever)
5. Clicar no corpo do email
6. Pressionar `Ctrl+V`

**Resultado Esperado:**
- [ ] Imagem do screenshot aparece no email
- [ ] Ou: aparece opção para fazer upload

**Resultado Atual:**
- [ ] ✅ Funcionou no Gmail
- [ ] ✅ Funcionou em outro site: _________________
- [ ] ❌ Não funcionou em nenhum site

**Sites para Testar:**
- [ ] Gmail (https://mail.google.com)
- [ ] Google Docs (https://docs.google.com)
- [ ] WhatsApp Web (https://web.whatsapp.com)

---

## 🧪 Teste 4: Histórico de Abas Fechadas

**Objetivo:** Verificar se mantém histórico de 10 abas

**Passos:**
1. Abrir 12 abas diferentes
2. Fechar todas (Ctrl+W)
3. Pressionar Ctrl+Shift+T 11 vezes
4. Contar quantas abas reabriram

**Resultado Esperado:**
- [ ] Apenas 10 abas reabrem (limite)
- [ ] Ordem inversa (LIFO)

**Resultado Atual:**
- [ ] ✅ Funcionou (10 abas)
- [ ] ❌ Reabriu menos: _____ abas
- [ ] ❌ Reabriu mais: _____ abas

---

## 🧪 Teste 5: Foco com Múltiplas Abas

**Objetivo:** Verificar se foco funciona com várias abas abertas

**Passos:**
1. Abrir 5 abas
2. Ir para aba 3
3. Pressionar Ctrl+T
4. Verificar foco

**Resultado Esperado:**
- [ ] Nova aba abre
- [ ] Cursor na URL entry
- [ ] Pode digitar imediatamente

**Resultado Atual:**
- [ ] ✅ Funcionou
- [ ] ❌ Não funcionou

---

## 📊 Resumo dos Testes

| Teste | Status | Observações |
|-------|--------|-------------|
| 1. Ctrl+T (foco) | ⬜ | |
| 2. Ctrl+Shift+T | ⬜ | |
| 3. Copy/Paste screenshot | ⬜ | |
| 4. Histórico 10 abas | ⬜ | |
| 5. Foco múltiplas abas | ⬜ | |

---

## 🐛 Reportar Problemas

Se algum teste falhar, forneça:

### 1. Logs do Terminal

Copie TODOS os logs que aparecerem, especialmente:
- Logs com 🔍 DEBUG
- Logs com ⌨️ (atalhos)
- Logs com ❌ (erros)

### 2. Comportamento Observado

Descreva exatamente o que aconteceu:
- Onde o cursor foi parar?
- O que apareceu na tela?
- Algum erro visível?

### 3. Passos para Reproduzir

Liste exatamente o que você fez:
1. Passo 1
2. Passo 2
3. etc

---

## ✅ Checklist Final

Antes de confirmar que está OK:

- [ ] Executei TODOS os 5 testes
- [ ] Marquei resultado de cada teste
- [ ] Copiei logs se houver problemas
- [ ] Testei em sites diferentes (Gmail, etc)
- [ ] Testei com múltiplas abas abertas

---

## 📝 Notas Adicionais

### Logs de Debug

Os logs com 🔍 DEBUG mostram:
- `keyVal`: código da tecla pressionada
- `t=116`: código da tecla 't' minúscula
- `T=84`: código da tecla 'T' maiúscula
- `Ctrl=true/false`: se Ctrl está pressionado
- `Shift=true/false`: se Shift está pressionado

**Exemplo de log correto:**
```
🔍 DEBUG: keyVal=116 (t=116, T=84), Ctrl=true, Shift=true
⌨️  Ctrl+Shift+T - Reabrir última aba fechada
```

**Se aparecer apenas o DEBUG mas não a ação, há um problema!**

---

## 🎯 Próximos Passos

Após executar todos os testes:

### Se TUDO OK:
1. Marcar todos como ✅
2. Confirmar que pode criar release v4.4.0

### Se ALGUM PROBLEMA:
1. Copiar logs completos
2. Descrever comportamento observado
3. Reportar para correção
4. Aguardar nova versão para testar

---

**Data:** 22/10/2025  
**Versão:** 4.4.0  
**Commit:** (após compilar)  
**Testador:** _________________  
**Data do Teste:** _________________
