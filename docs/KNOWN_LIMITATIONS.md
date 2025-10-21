# ⚠️ Limitações Conhecidas do Bagus Browser

## 🌐 Versão Webview (Atual)

### Sem Barra de Navegação Nativa

**Problema:**
O webview nativo (WebKit2GTK no Linux) **não possui barra de navegação própria**. Quando usamos `w.Navigate()` para carregar um site, ele substitui completamente o conteúdo da janela.

**Tentativas de Solução:**

1. **❌ Injeção de HTML + window.location.href**
   - Barra desaparece ao navegar
   - `window.location.href` substitui toda a página

2. **❌ Barra fixa + iframe**
   - Muitos sites bloqueiam iframe (`X-Frame-Options: DENY`)
   - DuckDuckGo, Google, Facebook, etc. não carregam
   - Resultado: tela em branco

3. **✅ Navegação direta com w.Navigate() (ATUAL)**
   - Sites carregam 100%
   - Mas sem barra de navegação visível
   - Navegação via atalhos de teclado ou menu

### Navegação Atual

**Como navegar sem barra:**

1. **Atalhos de Teclado:**
   - `Ctrl+L` - Focar barra de endereço (se implementado)
   - `Alt+←` - Voltar
   - `Alt+→` - Avançar
   - `F5` - Recarregar

2. **Via Console JavaScript:**
   ```javascript
   navigateGo("https://exemplo.com")
   ```

3. **Abrir nova instância:**
   ```bash
   ./build/bagus
   # Digite URL no terminal antes de abrir
   ```

### Soluções Possíveis

#### Opção 1: Duas Janelas (Implementação Futura)
```
┌─────────────────────────┐
│ Janela de Controle      │
│ [←][→][⟳] [URL______]   │
└─────────────────────────┘
┌─────────────────────────┐
│                         │
│   Janela de Conteúdo    │
│                         │
└─────────────────────────┘
```

**Vantagens:**
- ✅ Barra sempre visível
- ✅ Sites carregam 100%
- ✅ Controles sempre acessíveis

**Desvantagens:**
- ❌ Duas janelas separadas
- ❌ Mais complexo de gerenciar

#### Opção 2: Versão CEF (Chromium Embedded Framework)
```bash
./scripts/setup/install_cef.sh
./scripts/build/build_cef.sh
```

**Vantagens:**
- ✅ Barra de navegação integrada
- ✅ 100% compatibilidade
- ✅ DevTools (F12)
- ✅ Todas as features do Chrome

**Desvantagens:**
- ❌ ~165MB (vs 4MB webview)
- ❌ Instalação mais complexa
- ❌ Maior consumo de recursos

#### Opção 3: Browser Nativo com Wrapper
Usar browser nativo do sistema e apenas gerenciar:
- Histórico
- Bloqueio de ads
- Configurações

## 📊 Comparação de Versões

| Feature | Webview | CEF | Browser Nativo |
|---------|---------|-----|----------------|
| Tamanho | 4MB | 165MB | 0MB (usa sistema) |
| Barra Nav | ❌ | ✅ | ✅ |
| Compatibilidade | 100% | 100% | 100% |
| DevTools | ❌ | ✅ | ✅ |
| Privacidade | ✅ | ✅ | ⚠️ |
| Controle | ⚠️ | ✅ | ❌ |

## 🎯 Recomendação Atual

**Para uso diário:**
- Use a **versão CEF** para experiência completa

**Para desenvolvimento/testes:**
- Use a **versão webview** (mais leve e rápida)

**Para máxima privacidade:**
- Use a **versão webview** com navegação via código

## 🔄 Roadmap

### v2.1.0 (Próxima Versão)
- [ ] Implementar janela de controle separada
- [ ] Atalhos de teclado para navegação
- [ ] Menu de contexto com opções

### v2.2.0
- [ ] Versão CEF como padrão
- [ ] Webview como alternativa leve
- [ ] Seleção de versão no instalador

### v3.0.0
- [ ] Arquitetura multi-processo
- [ ] Abas nativas
- [ ] Extensões

## 📝 Notas Técnicas

### Por que iframe não funciona?

Muitos sites enviam o header HTTP:
```
X-Frame-Options: DENY
```

Ou CSP:
```
Content-Security-Policy: frame-ancestors 'none'
```

Isso **impede** que o site seja carregado em iframe por questões de segurança (previne clickjacking).

### Por que não usar Electron?

Electron seria uma solução, mas:
- ❌ ~100MB+ de overhead
- ❌ Baseado em Node.js (não Go)
- ❌ Mais complexo

CEF oferece funcionalidade similar com melhor integração Go.

---

**Última Atualização:** 2025-10-21  
**Versão:** 2.0.0  
**Status:** Documentado
