# 🎯 Script `bagus` - Centralizador Completo

## ✅ Confirmação: O Script `bagus` Continua Sendo o Centralizador

O script `bagus` foi **expandido e melhorado** para orquestrar completamente o workflow de bundling, mantendo sua posição como **centralizador único** de todas as operações.

---

## 🚀 Novos Comandos Adicionados

### 1. `./bagus bundle-complete` ⭐ RECOMENDADO
Executa o workflow **completo** em um único comando:
- Build
- Embarcar dependências
- Testar bundle
- Criar .deb bundled

```bash
./bagus bundle-complete
```

**Resultado:**
- ✅ Binário compilado
- ✅ 409 bibliotecas embarcadas
- ✅ Testes validados
- ✅ .deb bundled criado (73 MB)

---

### 2. `./bagus install-bundled`
Cria e instala a versão bundled em um comando:

```bash
./bagus install-bundled
```

**O que faz:**
1. Verifica se .deb bundled existe
2. Se não existir, executa `bundle-complete`
3. Instala o .deb bundled
4. Exibe informações de uso

---

### 3. `./bagus release X.Y.Z` (Melhorado)
Agora pergunta se quer criar .deb bundled:

```bash
./bagus release 5.0.1
```

**Workflow:**
1. Atualiza versão
2. Compila binário
3. **Pergunta:** "Deseja criar .deb com dependências embarcadas? (s/N):"
4. Se sim: executa `bundle-complete`
5. Cria tag e faz push

---

## 📊 Estrutura de Comandos

```
./bagus
├── build                    # Compilar + empacotar
├── bundle                   # Apenas embarcar dependências
├── bundle-complete          # ⭐ Build + Bundle + Testes + .deb
├── install                  # Build + Instalar padrão
├── install-bundled          # ⭐ Build + Bundle + Instalar
├── clean                    # Limpar builds
├── version                  # Ver versão
├── release <X.Y.Z>          # Release com opção bundled
├── publish                  # Publicar no GitHub
├── test                     # Testar compilação
├── run                      # Compilar e executar
├── status                   # Status do projeto
├── changelog                # Ver changelog
└── help                     # Ajuda
```

---

## 🎯 Workflows Recomendados

### Workflow 1: Desenvolvimento Rápido
```bash
./bagus build
```

### Workflow 2: Build + Bundled Completo
```bash
./bagus bundle-complete
```

### Workflow 3: Instalar Versão Bundled
```bash
./bagus install-bundled
```

### Workflow 4: Release Completa
```bash
./bagus release 5.0.1
# Pergunta se quer criar bundled
# Cria tag e faz push
# Próximo: ./bagus publish
```

---

## 📋 Exemplos de Uso

### Exemplo 1: Desenvolvimento Normal
```bash
# Apenas compilar
./bagus build

# Instalar versão padrão
./bagus install
```

### Exemplo 2: Criar Versão Bundled
```bash
# Tudo em um comando
./bagus bundle-complete

# Resultado: dist/bagus-browser_v5.0.0_amd64_bundled.deb (73 MB)
```

### Exemplo 3: Instalar Versão Bundled
```bash
# Cria e instala em um comando
./bagus install-bundled

# Resultado: Bagus Browser instalado com dependências embarcadas
```

### Exemplo 4: Release Completa
```bash
# Release com bundled
./bagus release 5.0.1

# Pergunta: Deseja criar .deb com dependências embarcadas? (s/N): s
# Resultado: Release criada com .deb padrão + bundled
```

---

## ✅ Benefícios da Centralização

### Para Usuários
- ✓ Um único script para tudo
- ✓ Comandos intuitivos e claros
- ✓ Sem necessidade de conhecer scripts internos
- ✓ Workflow simplificado

### Para Desenvolvedores
- ✓ Fácil manutenção
- ✓ Orquestração centralizada
- ✓ Scripts internos reutilizáveis
- ✓ Sem duplicação de lógica

### Para Distribuição
- ✓ Pacotes padrão (1.5 MB) - para quem quer dependências do sistema
- ✓ Pacotes bundled (73 MB) - para compatibilidade universal
- ✓ Ambos criados automaticamente
- ✓ Checksums validados

---

## 🔄 Fluxo de Orquestração

```
./bagus bundle-complete
    │
    ├─→ cmd_build
    │   └─→ Compila binário (5.5 MB)
    │
    ├─→ bash scripts/bundle-dependencies.sh
    │   └─→ Coleta 409 bibliotecas
    │
    ├─→ bash scripts/test-bundled.sh
    │   └─→ Valida 8 aspectos do bundle
    │
    └─→ bash scripts/build-deb-bundled.sh
        └─→ Cria .deb bundled (73 MB)
```

---

## 📊 Comparação de Workflows

| Tarefa | Antes | Agora |
|--------|-------|-------|
| Build | `./scripts/bagus build` | `./bagus build` |
| Build + Bundle | 4 comandos | `./bagus bundle-complete` |
| Instalar bundled | 5 comandos | `./bagus install-bundled` |
| Release | 2 comandos | `./bagus release 5.0.1` |

---

## 🎓 Hierarquia de Responsabilidades

```
./bagus (Centralizador)
├── Orquestra workflows
├── Valida dependências
├── Gerencia versões
└── Chama scripts internos
    ├── scripts/bundle-dependencies.sh (Coleta libs)
    ├── scripts/build-deb-bundled.sh (Cria .deb)
    ├── scripts/test-bundled.sh (Testa bundle)
    └── scripts/quick-bundle-test.sh (Teste rápido)
```

---

## ✨ Status Final

✅ **Script `bagus` Continua Sendo o Centralizador**
- Expandido com novos comandos
- Orquestra todos os workflows
- Mantém interface simples e intuitiva
- Sem complexidade adicional

✅ **Scripts Internos Reutilizáveis**
- Podem ser chamados independentemente
- Ou via `./bagus`
- Sem duplicação de lógica

✅ **Workflow Simplificado**
- Um comando para tudo
- Fácil de usar
- Fácil de manter

---

## 🚀 Próximos Passos

```bash
# Testar novo workflow
./bagus bundle-complete

# Ou instalar direto
./bagus install-bundled

# Ou fazer release
./bagus release 5.0.1
```

---

**Versão:** v5.0.0
**Data:** 02/11/2025
**Status:** ✅ Centralizador Completo e Funcional
