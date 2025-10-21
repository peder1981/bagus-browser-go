# 🚀 Quick Start - Bagus Browser v3.0.0

## ⚡ Início Rápido

### 1. Compilar
```bash
go build -o build/bagus .
```

### 2. Executar
```bash
./build/bagus
```

### 3. Usar
- Digite uma URL ou termo de busca
- Pressione **Enter** ou clique em **Ir**
- Navegue!

---

## 🎯 Primeiros Passos

### Navegar para um Site
```
1. Digite: github.com
2. Pressione Enter
3. ✅ Abre GitHub
```

### Fazer uma Busca
```
1. Digite: golang tutorial
2. Pressione Enter
3. ✅ Busca no DuckDuckGo
```

### Usar Atalhos
```
Alt+←  → Voltar
Alt+→  → Avançar
F5     → Recarregar
Ctrl+L → Focar URL
```

---

## 🎨 Interface

### Janela de Controle
```
┌─────────────────────────────────────────┐
│ [←] [→] [⟳] [✕] [URL____________] [Ir] │
└─────────────────────────────────────────┘
```

**Botões:**
- **←** Voltar
- **→** Avançar
- **⟳** Recarregar
- **✕** Parar
- **Ir** Navegar

### Janela de Conteúdo
```
┌─────────────────────────────────────────┐
│                                          │
│         [Site renderizado aqui]         │
│                                          │
└─────────────────────────────────────────┘
```

---

## 📝 Exemplos de Uso

### Exemplo 1: Navegar para Site
```bash
./build/bagus
# Digite: wikipedia.org
# Pressione Enter
# ✅ Abre Wikipedia
```

### Exemplo 2: Buscar Informação
```bash
./build/bagus
# Digite: como instalar go
# Pressione Enter
# ✅ Busca no DuckDuckGo
```

### Exemplo 3: Usar Atalhos
```bash
./build/bagus
# Navegue para qualquer site
# Pressione Alt+← para voltar
# Pressione Alt+→ para avançar
# Pressione F5 para recarregar
```

---

## 🔧 Requisitos

### Sistema
- Linux (testado em Ubuntu 20.04+)
- GTK3 instalado
- WebKit2GTK instalado

### Instalar Dependências (Ubuntu/Debian)
```bash
sudo apt-get update
sudo apt-get install -y \
    libgtk-3-dev \
    libwebkit2gtk-4.0-dev \
    pkg-config
```

### Go
- Go 1.21 ou superior

---

## 🐛 Solução de Problemas

### Erro: "cannot find package gtk"
```bash
# Instalar GTK3
sudo apt-get install libgtk-3-dev
```

### Erro: "cannot find package webkit2gtk"
```bash
# Instalar WebKit2GTK
sudo apt-get install libwebkit2gtk-4.0-dev
```

### Erro: "Package gtk+-3.0 was not found"
```bash
# Instalar pkg-config
sudo apt-get install pkg-config
```

### Browser não abre
```bash
# Verificar se já há instância rodando
ps aux | grep bagus

# Matar processo se necessário
pkill bagus

# Tentar novamente
./build/bagus
```

---

## 📊 Verificar Instalação

### Testar Compilação
```bash
go build -o build/bagus .
echo $?
# Deve retornar: 0
```

### Verificar Binário
```bash
ls -lh build/bagus
# Deve mostrar: ~6-7MB
```

### Testar Execução
```bash
./build/bagus &
sleep 2
ps aux | grep bagus
# Deve mostrar processo rodando
```

---

## 🎯 Próximos Passos

### Após Testar
1. ✅ Verificar navegação funciona
2. ✅ Testar atalhos de teclado
3. ✅ Validar segurança (URLs maliciosas)
4. ✅ Verificar performance

### Reportar Problemas
Se encontrar bugs:
1. Abra issue no GitHub
2. Inclua logs de erro
3. Descreva passos para reproduzir

---

## 📚 Documentação Completa

- **REBUILD_PLAN.md** - Plano de reconstrução
- **STATUS_V3.md** - Status completo da implementação
- **README.md** - Documentação principal

---

## 🎊 Pronto!

Seu Bagus Browser v3.0.0 está pronto para uso!

```bash
./build/bagus
```

**Divirta-se navegando com privacidade! 🌐🔒**
