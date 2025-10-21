# 🚀 Instalação do Bagus Browser

## Método 1: Script Único (Recomendado)

### Um comando faz tudo:

```bash
./install.sh
```

### O que o script faz automaticamente:

1. ✅ **Detecta seu sistema** (Ubuntu, Debian, Fedora, Arch)
2. ✅ **Verifica dependências** (Go, Git, GCC, etc.)
3. ✅ **Instala o que falta** (com sua permissão)
4. ✅ **Baixa dependências Go** (go mod download)
5. ✅ **Compila versão Webview** (~1 minuto)
6. ✅ **Pergunta sobre CEF** (opcional, 100% compatibilidade)
7. ✅ **Instala CEF se quiser** (~30 minutos)
8. ✅ **Compila versão CEF** (~5 minutos)
9. ✅ **Cria links simbólicos** ($HOME/.local/bin)
10. ✅ **Cria atalhos no menu** (aplicativos)
11. ✅ **Verifica privacidade** (zero telemetria)
12. ✅ **Mostra resumo completo**

### Tempo total:
- **Só Webview**: 2-5 minutos
- **Webview + CEF**: 30-40 minutos

---

## Método 2: Passo a Passo Manual

### 1. Instalar dependências

#### Ubuntu/Debian
```bash
sudo apt-get update
sudo apt-get install -y \
    golang \
    git \
    build-essential \
    libwebkit2gtk-4.0-dev \
    libgtk-3-dev
```

#### Fedora
```bash
sudo dnf install -y \
    golang \
    git \
    gcc-c++ \
    webkit2gtk3-devel \
    gtk3-devel
```

#### Arch Linux
```bash
sudo pacman -S --noconfirm \
    go \
    git \
    base-devel \
    webkit2gtk \
    gtk3
```

### 2. Clonar repositório
```bash
git clone https://github.com/peder1981/bagus-browser-go.git
cd bagus-browser-go
```

### 3. Instalar dependências Go
```bash
go mod download
go mod tidy
```

### 4. Compilar
```bash
# Versão Webview (padrão)
make build

# OU versão CEF (100% compatibilidade)
make install-cef  # Uma vez, ~30 min
make build-cef    # ~5 min
```

### 5. Executar
```bash
# Webview
./build/bagus

# CEF
cd build && ./run_bagus_cef.sh
```

---

## Método 3: Menu Interativo

```bash
make menu
```

Interface amigável que guia você passo a passo.

---

## 📋 Requisitos

### Mínimos
- **SO**: Linux (Ubuntu 20.04+, Debian 11+, Fedora 35+, Arch)
- **RAM**: 2GB
- **Disco**: 500MB (Webview) ou 3GB (CEF)
- **Go**: 1.21+

### Recomendados
- **RAM**: 4GB+
- **Disco**: 5GB+
- **CPU**: 2+ cores

---

## 🎯 Qual versão instalar?

### Instale Webview se:
- ✅ Quer algo leve e rápido
- ✅ Não precisa de Google/Facebook
- ✅ DuckDuckGo, Wikipedia são suficientes
- ✅ Quer instalar em 2 minutos

### Instale CEF se:
- ✅ Precisa de 100% compatibilidade
- ✅ Usa Google, Facebook, Twitter
- ✅ Quer DevTools (F12)
- ✅ Não se importa com 30 min de instalação

**Dica:** Instale Webview primeiro, teste, e depois instale CEF se precisar.

---

## 🔧 Após a Instalação

### Executar
```bash
# Se criou links simbólicos
bagus           # Webview
bagus-cef       # CEF

# Ou diretamente
./build/bagus
cd build && ./run_bagus_cef.sh
```

### Adicionar ao PATH
```bash
echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.bashrc
source ~/.bashrc
```

### Criar atalho desktop manualmente
```bash
cat > ~/.local/share/applications/bagus-browser.desktop << EOF
[Desktop Entry]
Version=1.0
Type=Application
Name=Bagus Browser
Comment=Navegador web seguro e privado
Exec=$HOME/.local/bin/bagus
Icon=web-browser
Terminal=false
Categories=Network;WebBrowser;
EOF
```

---

## 🆘 Problemas Comuns

### "go: command not found"
```bash
# Ubuntu/Debian
sudo apt-get install golang

# Fedora
sudo dnf install golang

# Arch
sudo pacman -S go
```

### "package libwebkit2gtk-4.0-dev is not available"
```bash
# Ubuntu/Debian
sudo apt-get install libwebkit2gtk-4.0-dev

# Se não encontrar, tente:
sudo apt-get install libwebkit2gtk-4.1-dev
```

### "CEF não encontrado"
```bash
# Instale CEF primeiro
make install-cef
# ou
./scripts/install_cef.sh
```

### "Permission denied"
```bash
chmod +x install.sh
./install.sh
```

### Erro de compilação
```bash
# Limpe e tente novamente
make clean
make build
```

---

## 🔄 Atualizar

```bash
git pull
./install.sh
```

O script detecta o que já está instalado e atualiza apenas o necessário.

---

## 🗑️ Desinstalar

```bash
# Remover binários
rm -rf build/
rm -f ~/.local/bin/bagus
rm -f ~/.local/bin/bagus-cef

# Remover atalhos
rm -f ~/.local/share/applications/bagus-browser*.desktop

# Remover configurações (opcional)
rm -rf ~/.bagus/

# Remover CEF (opcional)
sudo rm -rf /opt/cef
```

---

## 📊 Comparação de Métodos

| Método | Tempo | Dificuldade | Recomendado para |
|--------|-------|-------------|------------------|
| **Script Único** | 2-40 min | Fácil | Todos |
| **Menu Interativo** | 2-40 min | Fácil | Iniciantes |
| **Passo a Passo** | 5-45 min | Médio | Quem quer controle |
| **Makefile** | 2-40 min | Fácil | Desenvolvedores |

---

## 📖 Documentação Adicional

- [README.md](README.md) - Visão geral
- [BUILD.md](BUILD.md) - Guia de compilação detalhado
- [QUICKSTART_CEF.md](QUICKSTART_CEF.md) - CEF em 3 passos
- [PRIVACY.md](PRIVACY.md) - Política de privacidade
- [SECURITY.md](SECURITY.md) - Segurança

---

## ✅ Checklist de Instalação

- [ ] Dependências instaladas (Go, Git, GCC)
- [ ] Repositório clonado
- [ ] Dependências Go baixadas
- [ ] Versão Webview compilada
- [ ] (Opcional) CEF instalado
- [ ] (Opcional) Versão CEF compilada
- [ ] Links simbólicos criados
- [ ] PATH configurado
- [ ] Atalhos desktop criados
- [ ] Privacidade verificada
- [ ] Navegador testado

---

**Pronto para começar? Execute `./install.sh`!** 🚀
