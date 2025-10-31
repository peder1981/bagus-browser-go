# 🔨 Guia de Build - Bagus Browser

## 🎯 Sistema Unificado

Tudo é controlado por **UM ÚNICO SCRIPT**: `./bagus`

Esqueça a confusão de múltiplos scripts. Agora é simples:

```bash
./bagus <comando>
```

---

## 🚀 Comandos Principais

### 📦 Build & Instalação

```bash
# Compilar e empacotar (.deb + .tar.gz)
./bagus build

# Compilar e instalar localmente
./bagus install

# Limpar builds e temporários
./bagus clean
```

### 🏷️ Versionamento & Release

```bash
# Ver versão atual
./bagus version

# Criar release completa (RECOMENDADO)
./bagus release 4.5.1

# O que faz:
# 1. Atualiza versão
# 2. Compila e empacota
# 3. Cria tag git
# 4. Commit
# 5. Push para GitHub
```

### 📤 Publicação

```bash
# Publicar no GitHub (requer gh CLI)
./bagus publish-auto

# Ou publicar manualmente
./bagus publish
```

### 🧪 Desenvolvimento

```bash
# Testar compilação
./bagus test

# Compilar e executar
./bagus run

# Ver status do projeto
./bagus status

# Ver changelog
./bagus changelog
```

---

## 📋 Workflow Completo

### Desenvolvimento Normal

```bash
# 1. Fazer mudanças no código
vim main.go

# 2. Testar
./bagus test

# 3. Instalar localmente
./bagus install

# 4. Testar manualmente
bagus-browser
```

### Criar Nova Release

```bash
# Um único comando faz TUDO:
./bagus release 4.5.1

# Depois publicar:
./bagus publish-auto
```

Pronto! Não precisa de mais nada.

---

## 🔧 Compatibilidade com Make

Se preferir usar `make`, ainda funciona:

```bash
make build          # = ./bagus build
make install        # = ./bagus install
make clean          # = ./bagus clean
make test           # = ./bagus test
make version        # = ./bagus version
make status         # = ./bagus status
make release VERSION=4.5.1  # = ./bagus release 4.5.1
make publish        # = ./bagus publish-auto
```

Mas internamente, tudo chama `./bagus`.

---

## 📁 Estrutura Limpa

```
bagus-browser-go/
├── bagus                 # ⭐ SCRIPT MASTER (use este!)
├── Makefile              # Wrapper para ./bagus
├── main.go               # Código fonte
├── *.go                  # Outros arquivos
├── assets/               # Ícones
├── docs/                 # Documentação
├── dist/                 # Pacotes finais (criado pelo build)
└── build/                # Temporário (criado e removido)
```

**Não há mais:**
- ❌ scripts/build.sh
- ❌ scripts/version.sh
- ❌ scripts/release.sh
- ❌ scripts/publish.sh
- ❌ Confusão de scripts

**Há apenas:**
- ✅ `./bagus` - Um script para governar todos

---

## 🎯 Exemplos Práticos

### Exemplo 1: Desenvolvimento Rápido

```bash
# Editar código
vim settings.go

# Testar
./bagus test

# Instalar e testar
./bagus install
bagus-browser
```

### Exemplo 2: Release Completa

```bash
# Criar release 4.5.1
./bagus release 4.5.1

# Publicar no GitHub
./bagus publish-auto

# Pronto! 🎉
```

### Exemplo 3: Apenas Compilar

```bash
# Build sem instalar
./bagus build

# Pacotes em dist/
ls -lh dist/
```

### Exemplo 4: Limpar Tudo

```bash
# Remover builds, dist, temporários
./bagus clean
```

---

## 🔍 Ajuda

```bash
# Ver todos os comandos
./bagus help

# Ou apenas
./bagus
```

---

## 📊 Dependências

### Obrigatórias

```bash
sudo apt-get install \
    golang-go \
    libgtk-3-dev \
    libwebkit2gtk-4.0-dev
```

### Opcionais (para publicação)

```bash
# GitHub CLI (para ./bagus publish)
sudo apt install gh
gh auth login
```

---

## ⚡ Atalhos Rápidos

```bash
# Build rápido
./bagus build

# Instalar rápido
./bagus install

# Release rápido
./bagus release 4.5.1

# Publicar rápido
./bagus publish-auto

# Status rápido
./bagus status
```

---

## 🎨 Cores e Interface

O script `./bagus` usa cores para melhor visualização:

- 🔵 **Azul**: Ações em progresso
- 🟢 **Verde**: Sucesso
- 🟡 **Amarelo**: Avisos/Info
- 🔴 **Vermelho**: Erros
- 🟣 **Roxo**: Release
- 🔷 **Ciano**: Títulos/Seções

---

## 📝 Notas

1. **Sempre use `./bagus`** para operações de build/release
2. **Makefile é apenas wrapper** para compatibilidade
3. **Scripts antigos foram removidos** (ou serão em breve)
4. **Um comando, uma responsabilidade** - simples e direto
5. **Tudo é versionado** automaticamente

---

## 🆘 Solução de Problemas

### Erro: "comando não encontrado"

```bash
chmod +x bagus
./bagus help
```

### Erro: "dependências não encontradas"

```bash
sudo apt-get install golang-go libgtk-3-dev libwebkit2gtk-4.0-dev
```

### Erro: "gh não encontrado" (ao publicar)

```bash
sudo apt install gh
gh auth login
```

### Limpar tudo e recomeçar

```bash
./bagus clean
./bagus build
./bagus install
```

---

## 🎉 Conclusão

**Antes:**
- 5+ scripts diferentes
- Confusão de comandos
- Workflow complicado

**Agora:**
- 1 script master
- Comandos claros
- Workflow simples

```bash
./bagus release 4.5.1
./bagus publish-auto
```

Pronto! 🚀

---

**Versão:** 4.5.1  
**Data:** 23/10/2025  
**Sistema:** Unificado e Simplificado
