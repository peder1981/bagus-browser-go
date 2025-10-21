# 🧪 INSTRUÇÕES DE TESTE - IMPORTANTE

## ⚠️ ATENÇÃO

O binário foi **RECOMPILADO AGORA** (18:16).

---

## 🚀 Como Testar Corretamente

### 1. Fechar Qualquer Instância Antiga

```bash
# Feche todas as janelas do Bagus Browser que estiverem abertas
# Pressione Ctrl+C em qualquer terminal rodando o browser
```

### 2. Executar o Novo Binário

```bash
cd /home/peder/bagus-browser-go
./build/bagus
```

### 3. Testar a Tela de Login

**Username válido**: `peder`

**O que deve acontecer**:
1. ✅ Tela de login abre
2. ✅ Digite: `peder`
3. ✅ **NÃO deve aparecer** `[object Promise]`
4. ✅ Clique "Iniciar Browser"
5. ✅ Browser deve abrir

---

## ❌ Se Ainda Mostrar [object Promise]

### Isso significa que o cache do WebView está sendo usado

**Solução**: Limpar cache e tentar novamente

```bash
# 1. Fechar browser
Ctrl+C

# 2. Limpar cache do WebView
rm -rf ~/.cache/webkitgtk
rm -rf /tmp/peder

# 3. Recompilar forçando rebuild
cd /home/peder/bagus-browser-go
rm -f build/bagus
go clean -cache
go build -o build/bagus ./cmd/bagus

# 4. Executar novamente
./build/bagus
```

---

## ✅ Teste Definitivo

### Username: `peder`

**Resultado Esperado**:
```
1. Tela de login abre
2. Campo username aceita digitação
3. Validação em tempo real (sem [object Promise])
4. Clique "Iniciar Browser"
5. Janela do browser abre
6. Console mostra apenas:
   2025/10/20 18:XX:XX Iniciando Bagus Browser...
   Overriding existing handler for signal 10...
```

**NÃO deve aparecer**:
- ❌ `[object Promise]`
- ❌ `ReferenceError`
- ❌ Erros JavaScript

---

## 🐛 Debug

Se o problema persistir, execute:

```bash
# Verificar se o binário é novo
stat build/bagus | grep Modify
# Deve mostrar: 18:16 ou mais recente

# Verificar código compilado
strings build/bagus | grep "startBrowserClick"
# Deve encontrar a função

# Executar com debug
./build/bagus 2>&1 | tee debug.log
```

---

## 📞 Reporte

Após testar, reporte:

1. **Versão do binário**: `stat build/bagus | grep Modify`
2. **Erro ainda aparece?**: Sim/Não
3. **Console output**: Cole os logs
4. **Screenshot**: Se possível

---

**IMPORTANTE**: Execute `./build/bagus` (não `bagus-browser` ou outro binário antigo)

**Data do Build**: 2025-10-20 18:16
**Status**: Pronto para teste
