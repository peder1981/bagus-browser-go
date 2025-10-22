# 📁 Estrutura do Projeto - Bagus Browser

## 📂 Organização de Diretórios

```
bagus-browser-go/
├── 📄 README.md              # Documentação principal
├── 📄 CHANGELOG.md           # Histórico de versões
├── 📄 LICENSE                # Licença MIT
├── 📄 ESTRUTURA.md           # Este arquivo
│
├── 📁 assets/                # Recursos visuais
│   └── icons/                # Ícones do aplicativo
│
├── 📁 build/                 # Arquivos de build (gitignored)
│   └── bagus-browser         # Binário compilado
│
├── 📁 dist/                  # Pacotes de distribuição (gitignored)
│   ├── *.deb                 # Pacote Debian
│   ├── *.tar.gz              # Tarball
│   └── SHA256SUMS            # Checksums
│
├── 📁 docs/                  # Documentação completa
│   ├── 📄 README.md          # Índice da documentação
│   ├── 📄 FEATURES.md        # Lista de funcionalidades
│   ├── 📄 SECURITY.md        # Política de segurança
│   ├── 📄 PRIVACY.md         # Política de privacidade
│   ├── 📄 TESTING_GUIDE.md   # Guia de testes
│   │
│   ├── 📁 releases/          # Documentação de releases
│   │   ├── RELEASE_NOTES_v4.2.0.md
│   │   ├── PUBLICAR_RELEASE.md
│   │   └── PUBLICAÇÃO_FINAL.md
│   │
│   └── 📁 development/       # Documentação de desenvolvimento
│       ├── MELHORIAS_UX.md
│       ├── CORREÇÕES_APLICADAS.md
│       ├── SCRIPTS_CORRIGIDOS.md
│       ├── VALIDATION_REPORT.md
│       └── ...
│
├── 📁 scripts/               # Scripts de automação
│   ├── 📄 README.md          # Documentação dos scripts
│   ├── build.sh              # Compilação e empacotamento
│   ├── release.sh            # Preparação de release
│   └── publish.sh            # Workflow completo
│
└── 📁 src/ (código Go)       # Código-fonte
    ├── main.go               # Arquivo principal
    ├── security.go           # Validação e sanitização
    ├── privacy.go            # Configurações de privacidade
    ├── crypto.go             # Criptografia
    ├── bookmarks.go          # Gerenciador de favoritos
    └── downloads.go          # Gerenciador de downloads
```

---

## 📚 Documentação por Categoria

### 📄 Raiz (Arquivos Principais)
- **README.md** - Documentação principal do projeto
- **CHANGELOG.md** - Histórico de todas as versões
- **LICENSE** - Licença MIT
- **ESTRUTURA.md** - Este arquivo (organização do projeto)

### 📁 docs/ (Documentação Geral)
- **FEATURES.md** - Lista completa de funcionalidades
- **SECURITY.md** - Política de segurança
- **PRIVACY.md** - Política de privacidade
- **TESTING_GUIDE.md** - Guia de testes

### 📁 docs/releases/ (Releases)
Documentação específica de cada release:
- **RELEASE_NOTES_vX.X.X.md** - Notas da release
- **PUBLICAR_RELEASE.md** - Guia de publicação
- **PUBLICAÇÃO_FINAL.md** - Checklist de publicação

### 📁 docs/development/ (Desenvolvimento)
Documentação técnica e relatórios:
- **MELHORIAS_UX.md** - Melhorias de UX implementadas
- **CORREÇÕES_APLICADAS.md** - Correções e fixes
- **SCRIPTS_CORRIGIDOS.md** - Mudanças nos scripts
- **VALIDATION_REPORT.md** - Relatórios de validação
- **TEST_VALIDATION.md** - Checklists de testes

### 📁 scripts/ (Scripts)
Scripts de automação com documentação própria:
- **README.md** - Documentação dos scripts
- **build.sh** - Compilação e empacotamento
- **release.sh** - Preparação de release
- **publish.sh** - Workflow completo

---

## 🎯 Convenções

### Nomenclatura de Arquivos
- **MAIÚSCULAS.md** - Documentação importante
- **minúsculas.go** - Código-fonte
- **minúsculas.sh** - Scripts

### Gitignore
Arquivos/pastas ignorados:
- `build/` - Arquivos de compilação
- `dist/` - Pacotes de distribuição
- `*.deb`, `*.tar.gz` - Artefatos
- Arquivos temporários e caches

### Versionamento
- Tags: `vX.Y.Z` (ex: v4.2.0)
- Branches: `main` (produção)
- Commits: Conventional Commits

---

## 📖 Onde Encontrar

### Quero saber sobre...

**Funcionalidades do browser:**
→ `docs/FEATURES.md`

**Como instalar:**
→ `README.md` seção Instalação

**Histórico de mudanças:**
→ `CHANGELOG.md`

**Segurança e privacidade:**
→ `docs/SECURITY.md` e `docs/PRIVACY.md`

**Como compilar:**
→ `scripts/README.md`

**Notas da última release:**
→ `docs/releases/RELEASE_NOTES_v4.2.0.md`

**Como contribuir:**
→ `README.md` seção Contribuindo

---

## 🔧 Manutenção

### Adicionar Nova Release
1. Atualizar `CHANGELOG.md`
2. Criar `docs/releases/RELEASE_NOTES_vX.X.X.md`
3. Executar `bash scripts/build.sh`
4. Executar `bash scripts/release.sh`

### Adicionar Nova Funcionalidade
1. Implementar código
2. Atualizar `docs/FEATURES.md`
3. Atualizar `CHANGELOG.md`
4. Criar documentação em `docs/development/` se necessário

### Limpar Arquivos Temporários
```bash
rm -rf build/ dist/
git clean -fdx
```

---

**Última atualização:** 21/10/2025  
**Versão:** 4.2.0
