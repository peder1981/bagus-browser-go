# ✅ Scripts Atualizados para Nova Estrutura

## 🎯 Mudanças Realizadas

Todos os scripts foram atualizados para usar os paths corretos da nova estrutura organizada.

---

## 📝 Arquivos Modificados

### 1. **scripts/bagus** (Script Master)

#### Mudanças:
```bash
# ANTES:
go build -ldflags="-s -w" -o ${BUILD_DIR}/${APP_NAME} .

# AGORA:
go build -ldflags="-s -w" -o ${BUILD_DIR}/${APP_NAME} ./cmd/bagus-browser
```

```bash
# ANTES:
sed -i "s/^CURRENT_VERSION=.*/CURRENT_VERSION=\"${NEW_VERSION}\"/" bagus
git add bagus

# AGORA:
sed -i "s/^CURRENT_VERSION=.*/CURRENT_VERSION=\"${NEW_VERSION}\"/" scripts/bagus
git add scripts/bagus
```

**Impacto:**
- ✅ Compilação agora aponta para `cmd/bagus-browser/`
- ✅ Versionamento atualiza o script correto
- ✅ Git commits rastreiam o arquivo correto

---

### 2. **Makefile**

#### Mudanças:
```makefile
# ANTES:
help:
	@./bagus help

build:
	@./bagus build

# AGORA:
help:
	@./scripts/bagus help

build:
	@./scripts/bagus build
```

**Todos os comandos atualizados:**
- `make help` → `./scripts/bagus help`
- `make build` → `./scripts/bagus build`
- `make install` → `./scripts/bagus install`
- `make clean` → `./scripts/bagus clean`
- `make test` → `./scripts/bagus test`
- `make version` → `./scripts/bagus version`
- `make status` → `./scripts/bagus status`
- `make release` → `./scripts/bagus release`
- `make publish` → `./scripts/bagus publish-auto`
- `make run` → `./scripts/bagus run`

**Impacto:**
- ✅ Makefile funciona com nova estrutura
- ✅ Comandos curtos continuam funcionando
- ✅ Compatibilidade mantida

---

### 3. **scripts/benchmark.sh**

#### Mudanças:
```bash
# ANTES:
echo "Instale primeiro: ./bagus install"

# AGORA:
echo "Instale primeiro: ./scripts/bagus install"
```

**Impacto:**
- ✅ Mensagens de erro corretas
- ✅ Usuário sabe o comando certo

---

## 🚀 Como Usar os Scripts

### Desenvolvimento Rápido
```bash
# Compilar e testar
./scripts/bagus build
./scripts/bagus run

# Ou usar Makefile
make build
make run
```

### Build Completo
```bash
# Build com pacotes .deb e .tar.gz
./scripts/bagus build

# Ou
make build
```

### Instalação Local
```bash
# Compilar e instalar
./scripts/bagus install

# Ou
make install
```

### Release Completa
```bash
# Criar release (version + build + tag + push)
./scripts/bagus release 4.6.4

# Ou
make release VERSION=4.6.4
```

### Publicação no GitHub
```bash
# Publicar release
./scripts/bagus publish

# Ou
make publish
```

---

## 🔍 Verificação

### Testar Compilação
```bash
# Limpar e compilar do zero
./scripts/bagus clean
./scripts/bagus build

# Verificar se binário foi criado
ls -lh build/bagus-browser
```

### Testar Instalação
```bash
# Instalar localmente
./scripts/bagus install

# Verificar se foi instalado
which bagus-browser
bagus-browser --version
```

### Testar Pacotes
```bash
# Build completo
./scripts/bagus build

# Verificar pacotes criados
ls -lh dist/
# Deve mostrar:
# - bagus-browser_v4.6.3_amd64.deb
# - bagus-browser_v4.6.3_linux_amd64.tar.gz
# - SHA256SUMS
```

---

## ✅ Checklist de Validação

Antes de fazer release, verifique:

- [ ] `./scripts/bagus build` compila sem erros
- [ ] `./scripts/bagus test` passa todos os testes
- [ ] `./scripts/bagus install` instala corretamente
- [ ] `make build` funciona
- [ ] `make install` funciona
- [ ] Pacotes .deb e .tar.gz são criados em `dist/`
- [ ] Binário está em `build/bagus-browser`
- [ ] Ícones estão em `assets/icons/`
- [ ] Documentação está atualizada

---

## 🎯 Estrutura de Paths

```
bagus-browser-go/
├── cmd/bagus-browser/          # ← Source principal
│   └── main.go
├── build/                      # ← Binários compilados
│   └── bagus-browser
├── dist/                       # ← Pacotes de distribuição
│   ├── *.deb
│   └── *.tar.gz
├── scripts/                    # ← Scripts de build
│   ├── bagus                  # ← Script master
│   └── benchmark.sh
└── Makefile                    # ← Atalhos make
```

---

## 🔧 Comandos Importantes

### Desenvolvimento
```bash
./scripts/bagus run              # Compilar e executar
./scripts/bagus test             # Rodar testes
./scripts/bagus clean            # Limpar builds
```

### Build
```bash
./scripts/bagus build            # Build completo
./scripts/bagus install          # Build + instalar
```

### Release
```bash
./scripts/bagus version          # Ver versão atual
./scripts/bagus release 4.6.4    # Criar release
./scripts/bagus publish          # Publicar no GitHub
```

### Informações
```bash
./scripts/bagus status           # Status do projeto
./scripts/bagus help             # Ajuda completa
```

---

## 📚 Documentação Relacionada

- [ESTRUTURA_PROJETO.md](ESTRUTURA_PROJETO.md) - Estrutura completa
- [BUILD.md](BUILD.md) - Guia de compilação
- [CHANGELOG.md](CHANGELOG.md) - Histórico de versões

---

## ✨ Resultado

**Todos os scripts agora:**
- ✅ Usam paths corretos
- ✅ Compilam da estrutura organizada
- ✅ Criam pacotes corretamente
- ✅ Funcionam com Makefile
- ✅ Não quebram o build

**Próxima release será:**
- ✅ Sem erros de path
- ✅ Build confiável
- ✅ Instalação correta
- ✅ Pacotes válidos

🎉 **Scripts prontos para produção!**
