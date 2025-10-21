# 🚀 Build e Release - Bagus Browser

## 📋 Pré-requisitos

### Dependências de Build
```bash
# Ubuntu/Debian
sudo apt-get install -y \
    libgtk-3-dev \
    libwebkit2gtk-4.0-dev \
    pkg-config \
    build-essential \
    git

# GitHub CLI (para releases)
sudo apt install gh
gh auth login
```

---

## 🔨 Build Local

### Opção 1: Script Automático
```bash
./build.sh
```

### Opção 2: Makefile
```bash
make build
```

### Opção 3: Manual
```bash
go build -ldflags="-s -w" -o bagus-browser .
```

**Resultado:**
- `build/bagus-browser` - Binário
- `dist/bagus-browser_v4.1.0_amd64.deb` - Pacote Debian
- `dist/bagus-browser_v4.1.0_linux_amd64.tar.gz` - Tarball
- `dist/SHA256SUMS` - Checksums

---

## 📦 Empacotamento

### Pacote .deb (Debian/Ubuntu)
```bash
./build.sh  # Cria automaticamente
```

**Estrutura:**
```
bagus-browser_v4.1.0_amd64.deb
├── DEBIAN/
│   └── control
├── usr/
│   ├── bin/
│   │   └── bagus-browser
│   ├── share/
│   │   ├── applications/
│   │   │   └── bagus-browser.desktop
│   │   └── doc/
│   │       └── bagus-browser/
│   │           ├── README.md
│   │           ├── LICENSE
│   │           └── CHANGELOG.md
```

### Tarball (Genérico)
```bash
./build.sh  # Cria automaticamente
```

**Conteúdo:**
- bagus-browser (binário)
- bagus-browser.desktop
- README.md
- LICENSE
- CHANGELOG.md

---

## 🚀 Release no GitHub

### Opção 1: Script Completo (Recomendado)
```bash
./publish.sh
```

**Faz:**
1. Build completo
2. Commit e push
3. Cria release no GitHub
4. Upload de todos os pacotes

### Opção 2: Release Manual
```bash
# 1. Build
./build.sh

# 2. Commit
git add -A
git commit -m "Release v4.1.0"
git push

# 3. Release
./release.sh
```

### Opção 3: Makefile
```bash
make publish
```

---

## 📝 Processo de Release

### 1. Atualizar Versão

**CHANGELOG.md:**
```markdown
## [4.1.0] - 2025-10-21
### Added
- Feature X
### Fixed
- Bug Y
```

**Criar Tag:**
```bash
git tag -a v4.1.0 -m "Release v4.1.0"
git push origin v4.1.0
```

### 2. Build e Empacotamento
```bash
./build.sh
```

**Verifica:**
- ✅ Compilação OK
- ✅ Pacote .deb criado
- ✅ Tarball criado
- ✅ Checksums gerados

### 3. Publicar no GitHub
```bash
./release.sh
```

**Cria:**
- ✅ Release no GitHub
- ✅ Upload de .deb
- ✅ Upload de .tar.gz
- ✅ Upload de SHA256SUMS
- ✅ Notas de release automáticas

---

## 🧪 Testar Instalação

### Debian/Ubuntu (.deb)
```bash
# Instalar
sudo dpkg -i dist/bagus-browser_v4.1.0_amd64.deb
sudo apt-get install -f  # Dependências

# Executar
bagus-browser

# Desinstalar
sudo apt-get remove bagus-browser
```

### Tarball
```bash
# Extrair
tar -xzf dist/bagus-browser_v4.1.0_linux_amd64.tar.gz

# Executar
./bagus-browser
```

### Verificar Checksums
```bash
cd dist/
sha256sum -c SHA256SUMS
```

---

## 🔐 Assinatura de Releases (Opcional)

### Configurar GPG
```bash
# Gerar chave
gpg --full-generate-key

# Listar chaves
gpg --list-secret-keys --keyid-format LONG

# Configurar Git
git config --global user.signingkey YOUR_KEY_ID
git config --global commit.gpgsign true
git config --global tag.gpgsign true
```

### Assinar Release
```bash
# Tag assinada
git tag -s v4.1.0 -m "Release v4.1.0"

# Assinar arquivos
cd dist/
gpg --armor --detach-sign bagus-browser_v4.1.0_amd64.deb
gpg --armor --detach-sign bagus-browser_v4.1.0_linux_amd64.tar.gz
```

---

## 📊 Checklist de Release

- [ ] Atualizar CHANGELOG.md
- [ ] Atualizar versão no README.md
- [ ] Commit todas as mudanças
- [ ] Criar e push tag
- [ ] Executar `./build.sh`
- [ ] Testar pacotes localmente
- [ ] Executar `./release.sh`
- [ ] Verificar release no GitHub
- [ ] Testar download e instalação
- [ ] Anunciar release

---

## 🐛 Troubleshooting

### Erro: "gh not found"
```bash
sudo apt install gh
gh auth login
```

### Erro: "Tag already exists"
```bash
# Deletar tag local e remota
git tag -d v4.1.0
git push origin :refs/tags/v4.1.0

# Recriar
git tag -a v4.1.0 -m "Release v4.1.0"
git push origin v4.1.0
```

### Erro: "dpkg-deb: error"
```bash
# Verificar permissões
chmod +x build/bagus-browser

# Verificar estrutura DEBIAN/
ls -la build/bagus-browser_v4.1.0_amd64/DEBIAN/
```

### Erro de Compilação
```bash
# Verificar dependências
pkg-config --list-all | grep -E "gtk|webkit"

# Reinstalar
sudo apt-get install --reinstall libgtk-3-dev libwebkit2gtk-4.0-dev
```

---

## 📚 Recursos

- [GitHub CLI](https://cli.github.com/)
- [Debian Packaging](https://www.debian.org/doc/manuals/maint-guide/)
- [Go Build](https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies)

---

**Pronto para publicar releases profissionais!** 🚀
