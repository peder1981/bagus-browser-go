# 📁 Scripts - Bagus Browser

## ⚠️ ATENÇÃO: Sistema Atualizado!

**Os scripts antigos foram movidos para `.old/` e não devem mais ser usados.**

---

## 🎯 Novo Sistema Unificado

Tudo agora é controlado por **UM ÚNICO SCRIPT** na raiz do projeto:

```bash
../bagus <comando>
```

### Por que mudou?

**Antes:**
- ❌ 5+ scripts diferentes (build.sh, version.sh, release.sh, publish.sh, etc)
- ❌ Confusão de qual script usar
- ❌ Workflow complicado
- ❌ Duplicação de código

**Agora:**
- ✅ 1 script master (`../bagus`)
- ✅ Comandos claros e intuitivos
- ✅ Workflow simplificado
- ✅ Código centralizado

---

## 🚀 Como Usar

### Da raiz do projeto:

```bash
# Ver ajuda
./bagus help

# Build
./bagus build

# Instalar
./bagus install

# Release completa
./bagus release 4.5.1

# Publicar
./bagus publish-auto

# Status
./bagus status
```

### Ou use Make (wrapper):

```bash
make build
make install
make release VERSION=4.5.1
make publish
```

---

## 📚 Documentação Completa

Veja: `../BUILD.md`

---

## 🗂️ Estrutura Atual

```
scripts/
├── .old/              # Scripts antigos (backup)
│   ├── build.sh
│   ├── version.sh
│   ├── release.sh
│   └── publish.sh
└── README.md          # Este arquivo

../
├── bagus              # ⭐ SCRIPT MASTER - USE ESTE!
├── Makefile           # Wrapper para ./bagus
└── BUILD.md           # Documentação completa
```

---

## 🔄 Migração

Se você estava usando os scripts antigos:

**Antes:**
```bash
./scripts/build.sh
./scripts/version.sh release 4.5.1
./scripts/publish.sh
```

**Agora:**
```bash
./bagus build
./bagus release 4.5.1
./bagus publish-auto
```

Ou simplesmente:
```bash
./bagus release 4.5.1  # Faz tudo de uma vez!
```

---

## 📝 Scripts Antigos

Os scripts antigos estão em `.old/` apenas como backup.

**NÃO USE MAIS:**
- ❌ build.sh
- ❌ version.sh  
- ❌ release.sh
- ❌ publish.sh
- ❌ github-auth.sh
- ❌ install-desktop-icon.sh

**USE:**
- ✅ `../bagus` (na raiz)

---

## 🆘 Ajuda

```bash
# Ver todos os comandos
cd ..
./bagus help

# Ou ver documentação
cat BUILD.md
```

---

**Sistema:** Unificado  
**Versão:** 4.5.1  
**Data:** 23/10/2025
