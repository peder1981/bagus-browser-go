# 📦 Scripts de Bundling - Bagus Browser

## 🎯 Objetivo

Estes scripts automatizam o processo de embarcar todas as dependências (GTK3, WebKit2GTK, GStreamer) nos pacotes de distribuição.

## 📋 Scripts Disponíveis

### 1. `bundle-dependencies.sh`
Coleta todas as dependências do sistema e as copia para um diretório bundle.

**Uso:**
```bash
./scripts/bundle-dependencies.sh [diretório]
```

**Exemplo:**
```bash
./scripts/bundle-dependencies.sh build/bundle
```

**O que faz:**
- Cria estrutura de diretórios (lib, include, lib/pkgconfig)
- Copia bibliotecas .so de todos os pacotes necessários
- Copia headers de desenvolvimento
- Copia arquivos .pc (pkg-config)
- Cria script de configuração (setup-env.sh)
- Gera documentação (README.md)

**Saída:**
```
build/bundle/
├── lib/                    # ~200 arquivos .so
├── include/                # Headers
├── lib/pkgconfig/          # Arquivos .pc
├── setup-env.sh            # Script de configuração
└── README.md               # Documentação
```

**Tempo:** ~2-5 minutos

---

### 2. `build-deb-bundled.sh`
Cria pacote .deb com todas as dependências embarcadas.

**Uso:**
```bash
./scripts/build-deb-bundled.sh [versão]
```

**Exemplo:**
```bash
./scripts/build-deb-bundled.sh v5.0.0
```

**Pré-requisitos:**
- Binário compilado em `build/bagus-browser`
- Bundle em `build/bundle`
- .desktop em `build/bagus-browser.desktop`

**O que faz:**
1. Cria estrutura .deb
2. Copia binário para `/usr/bin/`
3. Copia bundle para `/usr/lib/bagus-browser/bundle/`
4. Copia .desktop para `/usr/share/applications/`
5. Copia ícones
6. Cria scripts de instalação (preinst, postinst, postrm)
7. Valida glibc mínima
8. Gera checksums SHA256

**Saída:**
```
dist/bagus-browser_v5.0.0_amd64_bundled.deb (~50-80 MB)
dist/SHA256SUMS (atualizado)
```

**Tempo:** ~1-2 minutos

---

### 3. `test-bundled.sh`
Testa integridade do bundle criado.

**Uso:**
```bash
./scripts/test-bundled.sh
```

**Testes realizados:**
1. ✓ Estrutura de diretórios
2. ✓ Bibliotecas críticas (GTK3, WebKit, GStreamer)
3. ✓ Arquivos .pc
4. ✓ setup-env.sh
5. ✓ Carregamento de bibliotecas via pkg-config
6. ✓ Tamanho total
7. ✓ Permissões de leitura
8. ✓ Documentação

**Saída:**
```
✓ Estrutura OK
✓ Bibliotecas OK
✓ .pc files OK
✓ setup-env.sh OK
✓ Carregamento OK
✓ Tamanho: 250 MB
✓ Permissões OK
✓ Documentação OK
```

**Tempo:** ~30 segundos

---

### 4. `quick-bundle-test.sh`
Executa o workflow completo em sequência.

**Uso:**
```bash
./scripts/quick-bundle-test.sh
```

**Workflow:**
1. Build: `./scripts/bagus build`
2. Bundle: `./scripts/bagus bundle`
3. Teste: `./scripts/test-bundled.sh`
4. .deb: `./scripts/build-deb-bundled.sh v5.0.0`

**Saída:**
```
✅ Build OK
✅ Bundle OK
✅ Testes OK
✅ .deb OK

📊 Estatísticas:
  Binário: 5 MB
  Bundle: 250 MB (200 arquivos)
  .deb: 60 MB

📦 Pacotes criados:
  dist/bagus-browser_v5.0.0_amd64_bundled.deb (60 MB)
```

**Tempo:** ~10-15 minutos

---

## 🚀 Workflow Completo

### Opção 1: Passo a Passo
```bash
# 1. Compilar
./scripts/bagus build

# 2. Embarcar dependências
./scripts/bagus bundle

# 3. Testar bundle
./scripts/test-bundled.sh

# 4. Criar .deb
./scripts/build-deb-bundled.sh v5.0.0

# 5. Instalar (opcional)
sudo dpkg -i dist/bagus-browser_v5.0.0_amd64_bundled.deb

# 6. Executar
bagus-browser
```

### Opção 2: Automático (Recomendado)
```bash
./scripts/quick-bundle-test.sh
```

---

## 📊 Tamanho dos Pacotes

| Tipo | Tamanho | Descrição |
|------|---------|-----------|
| Binário | ~5 MB | Executável compilado |
| Bundle | ~250 MB | Dependências descompactadas |
| .deb | ~50-80 MB | Pacote comprimido |
| .tar.gz | ~40-60 MB | Tarball comprimido |

---

## 🔍 Troubleshooting

### Erro: "Script bundle-dependencies.sh não encontrado"
```bash
chmod +x ./scripts/bundle-dependencies.sh
```

### Erro: "Binário não encontrado"
```bash
./scripts/bagus build
```

### Erro: "Bundle não encontrado"
```bash
./scripts/bagus bundle
```

### Erro: "dpkg-deb not found"
```bash
sudo apt install dpkg-dev
```

### Erro: "Permission denied"
```bash
chmod +x ./scripts/*.sh
```

---

## 🔧 Customização

### Adicionar Dependência
Editar `scripts/bundle-dependencies.sh` e adicionar pacote em `MAIN_PACKAGES` ou `SUPPORT_PACKAGES`:

```bash
MAIN_PACKAGES=(
    "libgtk-3-0"
    "libwebkit2gtk-4.0-37"
    "libgstreamer1.0-0"
    "nova-dependencia"  # ← Adicionar aqui
)
```

### Mudar Versão
```bash
./scripts/build-deb-bundled.sh v5.1.0
```

### Mudar Diretório de Saída
Editar `BUILD_DIR` e `DIST_DIR` nos scripts:

```bash
BUILD_DIR="meu-build"
DIST_DIR="meu-dist"
```

---

## 📝 Checklist de Release

- [ ] Compilar: `./scripts/bagus build`
- [ ] Embarcar: `./scripts/bagus bundle`
- [ ] Testar: `./scripts/test-bundled.sh`
- [ ] .deb: `./scripts/build-deb-bundled.sh v5.0.0`
- [ ] Verificar tamanho: `ls -lh dist/`
- [ ] Testar instalação: `sudo dpkg -i dist/*.deb`
- [ ] Testar execução: `bagus-browser`
- [ ] Publicar: `./scripts/bagus publish`

---

## 🎯 Benefícios

✅ **Robustez** - Funciona em qualquer distribuição Linux
✅ **Compatibilidade** - Sem conflitos de versão
✅ **Simplicidade** - Usuário não precisa instalar nada
✅ **Segurança** - Versões testadas e validadas
✅ **Manutenção** - Menos suporte necessário

---

## 📚 Documentação Relacionada

- [BUNDLED_DEPENDENCIES.md](../docs/BUNDLED_DEPENDENCIES.md) - Documentação completa
- [BUILD_MULTIPLATFORM.md](../docs/BUILD_MULTIPLATFORM.md) - Build multiplataforma
- [README.md](../README.md) - README principal

---

**Última atualização:** 02/11/2025
**Versão:** 1.0
**Status:** ✅ Pronto para Uso
