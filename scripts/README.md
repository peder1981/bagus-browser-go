# 🛠️ Scripts do Bagus Browser

Scripts para build, release e gerenciamento de versões.

---

## 🚀 Workflow Simplificado (RECOMENDADO)

### Criar Nova Release

```bash
# Um único comando faz tudo\!
./scripts/version.sh release 4.5.0
```

**O que faz:**
1. ✅ Atualiza CHANGELOG.md
2. ✅ Cria release notes em `docs/releases/`
3. ✅ Commita mudanças
4. ✅ Cria tag git
5. ✅ Compila e empacota (.deb + .tar.gz)
6. ✅ Limpa arquivos temporários
7. ✅ Push para GitHub

---

## 📋 Scripts Disponíveis

### version.sh ⭐ NOVO\!
Gerenciador centralizado de versões.

```bash
# Ver versão atual
./scripts/version.sh current

# Criar release completa
./scripts/version.sh release 4.5.0
```

### build.sh
Compila e empacota o browser.

```bash
./scripts/build.sh
```

Cria:
- `dist/bagus-browser_vX.X.X_amd64.deb`
- `dist/bagus-browser_vX.X.X_linux_amd64.tar.gz`
- `dist/SHA256SUMS`

### release.sh
Prepara release notes e instruções.

```bash
./scripts/release.sh
```

### publish.sh
Build + Instalação local.

```bash
./scripts/publish.sh
```

---

## 📁 Organização de Arquivos

```
scripts/
├── README.md           # Este arquivo
├── version.sh          # Gerenciador de versões ⭐
├── build.sh            # Build e empacotamento
├── publish.sh          # Publicação local
└── release.sh          # Preparação de release

dist/                   # Pacotes finais (mantido)
├── bagus-browser_vX.X.X_amd64.deb
├── bagus-browser_vX.X.X_linux_amd64.tar.gz
└── SHA256SUMS

build/                  # Temporário (removido após build)

docs/
├── releases/           # Release notes
└── development/        # Documentação técnica
```

---

## 🎯 Pilares do Projeto

1. ✅ **Organização:** Arquivos nas pastas corretas
2. ✅ **Limpeza:** Sem lixo na raiz
3. ✅ **Automação:** Um comando para tudo
4. ✅ **Documentação:** Tudo documentado

---

**Atualizado em:** 22/10/2025  
**Versão:** 4.4.0
