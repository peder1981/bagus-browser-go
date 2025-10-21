# 📚 Exemplos de Uso - Bagus Browser

## 🖱️ Modo Gráfico (Cliques)

### Uso Básico
```bash
bagus
```
1. Menu aparece
2. Clique em `1` ou `2`
3. Navegador abre

---

## ⌨️ Modo Linha de Comando

### Uso Direto (Sem Menu)

```bash
# Versão Rápida
bagus --fast

# Versão Completa
bagus --full
```

### Abrir URL Específica

```bash
# Rápida
bagus --fast duckduckgo.com
bagus --fast wikipedia.org

# Completa
bagus --full google.com
bagus --full facebook.com
```

### Scripts Automatizados

```bash
#!/bin/bash
# Script para abrir múltiplos sites

bagus --fast wikipedia.org &
bagus --fast stackoverflow.com &
bagus --full google.com &
```

### Atalhos de Teclado

**No seu ~/.bashrc ou ~/.zshrc:**

```bash
# Atalhos rápidos
alias bb='bagus --fast'
alias bbf='bagus --full'
alias bbg='bagus --full google.com'
alias bbw='bagus --fast wikipedia.org'
```

**Uso:**
```bash
bb              # Abre versão rápida
bbf             # Abre versão completa
bbg             # Abre Google direto
bbw             # Abre Wikipedia direto
```

---

## 🔧 Integração com Outras Ferramentas

### Rofi / Dmenu

```bash
# ~/.config/rofi/scripts/bagus.sh
#!/bin/bash
choice=$(echo -e "Rápida\nCompleta" | rofi -dmenu -p "Bagus Browser")

case $choice in
    "Rápida")
        bagus --fast
        ;;
    "Completa")
        bagus --full
        ;;
esac
```

### i3wm

```bash
# ~/.config/i3/config
bindsym $mod+b exec bagus --fast
bindsym $mod+Shift+b exec bagus --full
```

### Polybar

```ini
[module/bagus]
type = custom/text
content = 🌐
click-left = bagus --fast
click-right = bagus --full
```

---

## 📋 Casos de Uso

### 1. Desenvolvedor Web

```bash
# Testa site em ambas versões
bagus --fast localhost:3000 &
bagus --full localhost:3000 &
```

### 2. Pesquisador

```bash
# Busca rápida
bagus --fast "linux kernel documentation"

# Acessa Google Scholar
bagus --full scholar.google.com
```

### 3. Usuário Casual

```bash
# Dia a dia
bagus --fast

# Quando precisa de redes sociais
bagus --full
```

### 4. Administrador de Sistemas

```bash
#!/bin/bash
# Script de monitoramento
while true; do
    bagus --fast "http://monitor.local/dashboard"
    sleep 300
done
```

---

## 🎮 Atalhos Personalizados

### Criar Comandos Customizados

```bash
# ~/.local/bin/bb-wiki
#!/bin/bash
bagus --fast wikipedia.org "$@"

# ~/.local/bin/bb-google
#!/bin/bash
bagus --full google.com/search?q="$*"
```

**Uso:**
```bash
bb-wiki linux
bb-google "how to use bagus browser"
```

---

## 🚀 Automação Avançada

### Abrir Múltiplas Abas

```bash
#!/bin/bash
# morning-routine.sh

# Notícias
bagus --fast news.ycombinator.com &
sleep 2

# Email (precisa de versão completa)
bagus --full gmail.com &
sleep 2

# Trabalho
bagus --fast github.com &
```

### Perfis de Uso

```bash
# ~/.bashrc
work() {
    bagus --fast github.com &
    bagus --fast stackoverflow.com &
    bagus --full gmail.com &
}

study() {
    bagus --fast wikipedia.org &
    bagus --fast youtube.com &
}

social() {
    bagus --full facebook.com &
    bagus --full twitter.com &
}
```

**Uso:**
```bash
work    # Abre sites de trabalho
study   # Abre sites de estudo
social  # Abre redes sociais
```

---

## 💡 Dicas

### 1. Modo Silencioso

```bash
bagus --fast > /dev/null 2>&1 &
```

### 2. Com Timeout

```bash
timeout 3600 bagus --fast  # Fecha após 1 hora
```

### 3. Verificar Versão Instalada

```bash
bagus --help | head -1
```

### 4. Abrir em Background

```bash
bagus --fast &
disown
```

---

## 📊 Resumo

| Modo | Comando | Uso |
|------|---------|-----|
| **Interativo** | `bagus` | Menu gráfico |
| **Rápida** | `bagus --fast` | Direto, sem menu |
| **Completa** | `bagus --full` | Direto, sem menu |
| **URL** | `bagus --fast URL` | Abre URL específica |
| **Ajuda** | `bagus --help` | Mostra opções |

---

**Funciona tanto com cliques quanto por linha de comando!** 🎯✨
