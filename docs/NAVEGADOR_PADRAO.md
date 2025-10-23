# 🌐 Configurar Bagus Browser como Navegador Padrão

## 📋 Índice

1. [O Que Mudou na v4.5.1](#o-que-mudou-na-v451)
2. [Como Funciona](#como-funciona)
3. [Configurar como Padrão](#configurar-como-padrão)
4. [Testar](#testar)
5. [Exemplos de Uso](#exemplos-de-uso)
6. [Solução de Problemas](#solução-de-problemas)

---

## 🎯 O Que Mudou na v4.5.1

### ✨ Nova Funcionalidade: Suporte a URLs Externas

O Bagus Browser agora pode:

- ✅ **Receber URLs de outros aplicativos**
- ✅ **Abrir links clicados em PDFs, emails, etc**
- ✅ **Ser configurado como navegador padrão do sistema**
- ✅ **Abrir múltiplas URLs de uma vez**
- ✅ **Respeitar autenticação (senha mestre)**

### 🔒 Segurança Mantida

- Se você tem senha mestre configurada, ela será solicitada **antes** de abrir qualquer URL externa
- Nenhuma URL é aberta sem autenticação
- Proteção contra abertura não autorizada

---

## 🔧 Como Funciona

### Linha de Comando

Agora você pode abrir URLs diretamente:

```bash
# Abrir uma URL
bagus-browser https://duckduckgo.com

# Abrir múltiplas URLs (cada uma em uma aba)
bagus-browser https://github.com https://duckduckgo.com

# Abrir domínio sem protocolo (adiciona https:// automaticamente)
bagus-browser duckduckgo.com

# Abrir arquivo local
bagus-browser file:///home/user/documento.html
```

### Integração com Sistema

Quando configurado como navegador padrão:

```
1. Você clica em um link no Thunderbird
2. Sistema chama: bagus-browser https://exemplo.com
3. Se tem senha mestre → Pede autenticação
4. Abre URL em nova aba
```

---

## ⚙️ Configurar como Padrão

### Método 1: Script Automático (Recomendado)

```bash
cd ~/bagus-browser-go
./scripts/set-default-browser.sh
```

**O que faz:**
- Configura Bagus como navegador padrão
- Registra para HTTP, HTTPS, HTML
- Verifica se funcionou

### Método 2: Manual (Ubuntu/GNOME)

1. **Abrir Configurações do Sistema**
   ```
   Configurações → Aplicativos Padrão
   ```

2. **Selecionar Navegador Web**
   ```
   Navegador Web → Bagus Browser
   ```

3. **Confirmar**
   ```
   Fechar configurações
   ```

### Método 3: Manual (KDE Plasma)

1. **Abrir Configurações do Sistema**
   ```
   Configurações do Sistema → Aplicativos → Aplicativos Padrão
   ```

2. **Selecionar Navegador Web**
   ```
   Navegador Web → Bagus Browser
   ```

### Método 4: Linha de Comando

```bash
# Configurar como padrão
xdg-settings set default-web-browser bagus-browser.desktop

# Verificar
xdg-settings get default-web-browser
# Deve retornar: bagus-browser.desktop
```

---

## 🧪 Testar

### Teste 1: Linha de Comando

```bash
# Teste simples
bagus-browser https://duckduckgo.com

# Deve:
# 1. Abrir Bagus Browser (ou focar se já aberto)
# 2. Criar nova aba
# 3. Carregar DuckDuckGo
```

### Teste 2: xdg-open

```bash
# Simular clique em link
xdg-open https://github.com

# Deve:
# 1. Abrir no Bagus Browser
# 2. Se for navegador padrão
```

### Teste 3: Múltiplas URLs

```bash
bagus-browser https://github.com https://duckduckgo.com https://gitlab.com

# Deve:
# 1. Abrir browser
# 2. Criar 3 novas abas
# 3. Carregar cada URL
```

### Teste 4: Com Senha Mestre

```bash
# 1. Configurar senha mestre (Ctrl+,)
# 2. Fechar browser
# 3. Executar:
bagus-browser https://duckduckgo.com

# Deve:
# 1. Pedir senha
# 2. Após autenticar, abrir URL
```

### Teste 5: Arquivo Local

```bash
# Criar arquivo HTML de teste
echo '<h1>Teste</h1>' > /tmp/teste.html

# Abrir
bagus-browser file:///tmp/teste.html

# Deve abrir o arquivo
```

---

## 📚 Exemplos de Uso

### Exemplo 1: Email com Link

```
Cenário: Você recebe email com link no Thunderbird

1. Clicar no link
2. Sistema abre Bagus Browser
3. Link carrega em nova aba
```

### Exemplo 2: PDF com Hyperlink

```
Cenário: PDF aberto no Evince com link

1. Clicar no link no PDF
2. Sistema abre Bagus Browser
3. Link carrega em nova aba
```

### Exemplo 3: Terminal

```bash
# Desenvolvedor abrindo documentação
bagus-browser https://go.dev/doc

# Múltiplas referências
bagus-browser \
  https://pkg.go.dev/github.com/gotk3/gotk3 \
  https://webkitgtk.org/reference/webkit2gtk/stable/ \
  https://developer.gnome.org/gtk3/
```

### Exemplo 4: Script de Automação

```bash
#!/bin/bash
# Script que abre várias páginas de monitoramento

bagus-browser \
  https://status.github.com \
  https://status.gitlab.com \
  https://downdetector.com
```

### Exemplo 5: Atalho de Desktop

```bash
# Criar atalho para site favorito
cat > ~/Desktop/DuckDuckGo.desktop <<EOF
[Desktop Entry]
Type=Application
Name=DuckDuckGo
Exec=bagus-browser https://duckduckgo.com
Icon=bagus-browser
Terminal=false
EOF

chmod +x ~/Desktop/DuckDuckGo.desktop
```

---

## 🔍 Solução de Problemas

### Problema 1: Links Não Abrem no Bagus

**Verificar navegador padrão:**
```bash
xdg-settings get default-web-browser
```

**Se não for `bagus-browser.desktop`:**
```bash
./scripts/set-default-browser.sh
```

### Problema 2: Abre Navegador Errado

**Resetar configuração:**
```bash
# Limpar cache
rm -rf ~/.local/share/applications/mimeapps.list

# Reconfigurar
./scripts/set-default-browser.sh
```

### Problema 3: URL Não Abre

**Verificar formato:**
```bash
# ✅ Correto
bagus-browser https://exemplo.com
bagus-browser http://exemplo.com
bagus-browser exemplo.com

# ❌ Incorreto
bagus-browser exemplo  # Falta .com ou protocolo
```

### Problema 4: Senha Não É Pedida

**Verificar configuração:**
```bash
# Abrir configurações
bagus-browser
# Pressionar Ctrl+,
# Verificar aba "Segurança"
# Confirmar que "Exigir senha" está marcado
```

### Problema 5: Múltiplas Instâncias

**Comportamento esperado:**
- Bagus Browser abre **uma nova aba** na instância existente
- Não abre múltiplas janelas

**Se abrir múltiplas janelas:**
```bash
# Verificar se há múltiplos processos
ps aux | grep bagus-browser

# Matar todos
killall bagus-browser

# Reabrir
bagus-browser
```

---

## 📊 Logs e Debug

### Ver Logs de URLs Externas

```bash
# Executar com logs
bagus-browser 2>&1 | grep "URL externa"

# Em outro terminal, testar
xdg-open https://duckduckgo.com

# Deve aparecer:
# 🔗 URL externa detectada: https://duckduckgo.com
# 📂 Abrindo URL em nova aba: https://duckduckgo.com
```

### Debug Completo

```bash
# Logs completos
bagus-browser https://duckduckgo.com 2>&1 | tee bagus-debug.log

# Verificar arquivo
cat bagus-debug.log | grep -E "URL|aba|auth"
```

---

## 🎯 Checklist de Configuração

- [ ] Bagus Browser instalado (`bagus-browser --version`)
- [ ] Script executado (`./scripts/set-default-browser.sh`)
- [ ] Navegador padrão configurado (`xdg-settings get default-web-browser`)
- [ ] Teste com `xdg-open https://duckduckgo.com` funciona
- [ ] Links em emails abrem no Bagus
- [ ] Links em PDFs abrem no Bagus
- [ ] Senha mestre funciona (se configurada)

---

## 📝 Notas Técnicas

### Argumentos Suportados

```bash
bagus-browser [URL1] [URL2] ... [URLn]
```

**Formatos aceitos:**
- `https://exemplo.com` - URL completa
- `http://exemplo.com` - HTTP não seguro
- `exemplo.com` - Domínio (adiciona https://)
- `file:///caminho/arquivo.html` - Arquivo local

**Ignorados:**
- Flags começando com `-` ou `--`
- Argumentos sem `.` e sem protocolo

### Arquivo .desktop

Localização: `/usr/share/applications/bagus-browser.desktop`

**Linha importante:**
```
Exec=bagus-browser %u
```

O `%u` permite receber URLs do sistema.

**MimeTypes registrados:**
```
text/html
text/xml
application/xhtml+xml
x-scheme-handler/http
x-scheme-handler/https
```

---

## 🚀 Próximos Passos

Após configurar:

1. **Testar com aplicativos reais:**
   - Thunderbird (email)
   - Evince (PDF)
   - LibreOffice (documentos)
   - Telegram (links)

2. **Configurar atalhos:**
   - Criar .desktop para sites favoritos
   - Adicionar ao dock/painel

3. **Automatizar:**
   - Scripts que abrem múltiplas páginas
   - Monitoramento de serviços

---

**Versão:** 4.5.1  
**Data:** 23/10/2025  
**Tipo:** Patch (correção/melhoria)
