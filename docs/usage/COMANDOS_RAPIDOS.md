# ⚡ Comandos Rápidos - Bagus Browser

## 🚀 Para Usuários

### Instalar
```bash
chmod +x install.sh
./install.sh
```

### Usar
```bash
bagus
```

---

## 📦 Para Criar Pacotes

### Pacote .deb (Linux)
```bash
./build-deb.sh
```

**Resultado:** `build/bagus-browser_1.0.0_amd64.deb`

### Executável Windows
```bash
./build-windows.sh
```

**Resultado:** `build/windows/bagus-browser.exe`

---

## 🔧 Para Desenvolvedores

### Compilar
```bash
go build -o build/bagus .
```

### Executar
```bash
./build/bagus
```

### Testar
```bash
go test ./...
```

### Limpar
```bash
rm -rf build/
```

---

## 📤 Para Publicar Release

### 1. Criar Pacotes
```bash
./build-deb.sh
./build-windows.sh
```

### 2. Commit
```bash
git add .
git commit -m "Release v1.0.0"
git push origin main
```

### 3. Tag
```bash
git tag -a v1.0.0 -m "Release v1.0.0"
git push origin v1.0.0
```

### 4. GitHub Release
```
1. Acesse: github.com/peder1981/bagus-browser-go/releases/new
2. Tag: v1.0.0
3. Anexe: build/bagus-browser_1.0.0_amd64.deb
4. Anexe: build/windows/bagus-browser.exe
5. Publique
```

---

## 🧪 Testes Rápidos

### Testar instalação .deb
```bash
sudo dpkg -i build/bagus-browser_1.0.0_amd64.deb
bagus-browser
sudo dpkg -r bagus-browser
```

### Testar código fonte
```bash
./install.sh
bagus
```

### Verificar sintaxe
```bash
bash -n install.sh
bash -n build-deb.sh
bash -n build-windows.sh
```

---

## 📊 Informações Rápidas

### Ver tamanho
```bash
du -h build/bagus
```

### Ver dependências
```bash
ldd build/bagus
```

### Ver versão Go
```bash
go version
```

### Listar arquivos .deb
```bash
dpkg -c build/bagus-browser_1.0.0_amd64.deb
```

---

## 🔍 Debug

### Logs detalhados
```bash
./build/bagus --verbose
```

### Verificar instalação
```bash
which bagus
ls -la ~/.local/bin/bagus
```

### Limpar cache
```bash
rm -rf ~/.bagus/
```

---

## 📝 Documentação

### Ler documentos
```bash
cat README.md
cat INSTALACAO_SIMPLES.md
cat CEF_STATUS.md
cat RELEASE_GUIDE.md
cat ENTREGA_COMPLETA.md
```

### Editar
```bash
nano README.md
```

---

## ✅ Checklist Rápido

Antes de publicar:

```bash
# 1. Compilar
go build -o build/bagus .

# 2. Testar
./build/bagus

# 3. Criar .deb
./build-deb.sh

# 4. Criar .exe
./build-windows.sh

# 5. Verificar
ls -lh build/

# 6. Commit
git add .
git commit -m "Release v1.0.0"
git push

# 7. Tag
git tag v1.0.0
git push --tags

# 8. Release no GitHub (manual)
```

---

**Pronto! 🚀**
