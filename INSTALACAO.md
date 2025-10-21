# 📦 Instalação do Bagus Browser

## Um Único Comando

```bash
./install.sh
```

Pronto! O script faz tudo automaticamente.

---

## O que vai acontecer?

### 1. Menu de Escolha

Você verá 3 opções:

```
1) Versão Rápida (Recomendado)
   ├─ 2 minutos
   ├─ ~4MB
   ├─ DuckDuckGo, Wikipedia, YouTube
   └─ Google/Facebook não funcionam

2) Versão Completa (100% Sites)
   ├─ 40 minutos
   ├─ ~165MB
   ├─ TODOS os sites
   └─ Google, Facebook, Twitter funcionam

3) Ambas (Melhor Opção)
   ├─ bagus (rápido, uso diário)
   └─ bagus-full (completo, quando precisar)
```

### 2. Instalação Automática

O script vai:
- ✅ Detectar seu sistema (Ubuntu/Debian/Fedora/Arch)
- ✅ Instalar dependências
- ✅ Compilar o navegador
- ✅ Criar atalhos
- ✅ Configurar comandos

### 3. Pronto!

Após a instalação:

```bash
# Se escolheu opção 1 ou 3
bagus

# Se escolheu opção 2 ou 3
bagus-full
```

---

## Qual escolher?

### Escolha Opção 1 se:
- ✅ Quer algo rápido (2 min)
- ✅ Usa DuckDuckGo, Wikipedia, YouTube
- ✅ Não precisa de Google/Facebook

### Escolha Opção 2 se:
- ✅ Precisa de Google, Facebook, Twitter
- ✅ Quer 100% compatibilidade
- ✅ Não se importa com 40 min de instalação

### Escolha Opção 3 se: ⭐
- ✅ Quer o melhor dos dois mundos
- ✅ Usa `bagus` no dia a dia (rápido)
- ✅ Usa `bagus-full` quando precisar

---

## Requisitos

- **SO**: Linux (Ubuntu 20.04+, Debian 11+, Fedora 35+, Arch)
- **RAM**: 2GB
- **Disco**: 500MB (opção 1) ou 3GB (opção 2/3)
- **Internet**: Para baixar dependências

---

## Problemas?

### "Comando não encontrado"
```bash
# Adicione ao PATH
echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.bashrc
source ~/.bashrc
```

### "Erro ao compilar"
```bash
# Instale dependências manualmente
sudo apt-get install golang git build-essential libwebkit2gtk-4.1-dev
```

### "CEF não instala"
```bash
# Escolha opção 1 (Versão Rápida)
# Funciona perfeitamente sem CEF
```

---

## Desinstalar

```bash
rm -rf ~/.bagus
rm -f ~/.local/bin/bagus
rm -f ~/.local/bin/bagus-full
rm -f ~/.local/share/applications/bagus-browser*.desktop
```

---

**Dúvidas? Execute `./install.sh` e siga o menu!** 🚀
