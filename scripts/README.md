# Scripts de Automação

Scripts auxiliares para desenvolvimento e manutenção do Bagus Browser.

## 📋 Scripts Disponíveis

### 🔧 Build e Instalação

#### `build_cef.sh`
Compila a versão CEF (Chromium Embedded Framework) do browser.

```bash
./scripts/build_cef.sh
```

**Características:**
- 100% compatibilidade com todos os sites
- Requer CEF instalado (`./scripts/install_cef.sh`)
- Tamanho: ~165MB
- DevTools integrado (F12)

#### `install_cef.sh`
Instala o CEF (Chromium Embedded Framework).

```bash
./scripts/install_cef.sh
```

**Nota:** Download de ~500MB, instalação única.

#### `install-icon.sh`
Instala o ícone do Bagus Browser no sistema.

```bash
./scripts/install-icon.sh
```

Integra o browser com o menu de aplicativos do desktop.

### 🧪 Testes e Qualidade

#### `test.sh`
Executa testes automatizados, lint e formatação.

```bash
./scripts/test.sh [opção]
```

**Opções:**
- `unit` - Testes unitários
- `coverage` - Testes com coverage
- `lint` - Linter (golangci-lint)
- `fmt` - Formatação de código
- `all` - Tudo (padrão)

**Exemplos:**
```bash
./scripts/test.sh           # Executa tudo
./scripts/test.sh unit      # Apenas testes
./scripts/test.sh coverage  # Com coverage
```

#### `verify_privacy.sh`
Verifica que o código não contém telemetria.

```bash
./scripts/verify_privacy.sh
```

Audita o código para garantir zero telemetria e rastreamento.

### 🚀 Git e GitHub

#### `setup-github.sh`
Configura repositório Git e GitHub.

```bash
./scripts/setup-github.sh
```

**Funcionalidades:**
- Inicializa repositório Git
- Configura remote (SSH/HTTPS)
- Cria commit inicial
- Faz primeiro push

## 🎯 Scripts Principais (Raiz do Projeto)

Para uso diário, utilize os scripts na raiz:

- **`build-all.sh`** - Build completo multiplataforma
- **`build-deb.sh`** - Cria pacote .deb
- **`install.sh`** - Instalador rápido
- **`publish-release.sh`** - Publica release no GitHub

## 📚 Documentação Adicional

- [README.md](../README.md) - Documentação principal
- [RELEASE_INSTRUCTIONS.md](../RELEASE_INSTRUCTIONS.md) - Como publicar releases
- [SCRIPTS_ANALYSIS.md](../SCRIPTS_ANALYSIS.md) - Análise de scripts

## 🤝 Contribuindo

Ao modificar scripts:

1. Teste em ambiente limpo
2. Mantenha compatibilidade
3. Atualize esta documentação
4. Use shellcheck para validar

---

**Última Atualização:** 2025-10-21
