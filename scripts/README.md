# Scripts de Automação

Scripts auxiliares organizados por categoria para desenvolvimento e manutenção do Bagus Browser.

## 📁 Estrutura

```
scripts/
├── build/              # Scripts de compilação
├── packaging/          # Scripts de empacotamento
├── release/            # Scripts de release
├── setup/              # Scripts de configuração
└── testing/            # Scripts de testes
```

## 📋 Scripts Disponíveis

### 🔨 Build (build/)

#### `build_cef.sh`
Compila a versão CEF (Chromium Embedded Framework) do browser.

```bash
./scripts/build/build_cef.sh
# ou
make build-cef
```

**Características:**
- 100% compatibilidade com todos os sites
- Requer CEF instalado
- Tamanho: ~165MB
- DevTools integrado (F12)

### 📦 Packaging (packaging/)

#### `build-deb.sh`
Cria pacote .deb para Debian/Ubuntu.

```bash
./scripts/packaging/build-deb.sh
```

**Gera:**
- `build/bagus-browser_1.0.0_amd64.deb`

### 🚢 Release (release/)

#### `publish-release.sh`
Publica release no GitHub.

```bash
./scripts/release/publish-release.sh
```

**Requer:** GitHub CLI (gh) instalado

### ⚙️ Setup (setup/)

#### `install_cef.sh`
Instala o CEF (Chromium Embedded Framework).

```bash
./scripts/setup/install_cef.sh
# ou
make install-cef
```

**Nota:** Download de ~500MB, instalação única.

#### `install-icon.sh`
Instala o ícone do Bagus Browser no sistema.

```bash
./scripts/setup/install-icon.sh
```

Integra o browser com o menu de aplicativos do desktop.

#### `setup-github.sh`
Configura repositório Git e GitHub.

```bash
./scripts/setup/setup-github.sh
```

**Funcionalidades:**
- Inicializa repositório Git
- Configura remote
- Cria commit inicial

### 🧪 Testing (testing/)

#### `test.sh`
Executa testes automatizados, lint e formatação.

```bash
./scripts/testing/test.sh [opção]
```

**Opções:**
- `unit` - Testes unitários
- `coverage` - Testes com coverage
- `lint` - Linter (golangci-lint)
- `fmt` - Formatação de código
- `all` - Tudo (padrão)

**Exemplos:**
```bash
./scripts/testing/test.sh           # Executa tudo
./scripts/testing/test.sh unit      # Apenas testes
./scripts/testing/test.sh coverage  # Com coverage
```

#### `verify_privacy.sh`
Verifica que o código não contém telemetria.

```bash
./scripts/testing/verify_privacy.sh
# ou
make verify-privacy
```

Audita o código para garantir zero telemetria e rastreamento.

## 🎯 Scripts Principais (Raiz do Projeto)

Para uso diário, utilize os scripts na raiz:

- **`build-all.sh`** - Build completo multiplataforma
- **`install.sh`** - Instalador rápido

## 📚 Documentação Adicional

- [README.md](../README.md) - Documentação principal
- [docs/release/RELEASE_INSTRUCTIONS.md](../docs/release/RELEASE_INSTRUCTIONS.md) - Como publicar releases
- [docs/SCRIPTS_ANALYSIS.md](../docs/SCRIPTS_ANALYSIS.md) - Análise de scripts

## 🤝 Contribuindo

Ao modificar scripts:

1. Teste em ambiente limpo
2. Mantenha compatibilidade
3. Atualize esta documentação
4. Use shellcheck para validar

---

**Última Atualização:** 2025-10-21
