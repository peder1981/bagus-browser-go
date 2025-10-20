# Guia Rápido - Publicação no GitHub

Este guia mostra como publicar rapidamente o Bagus Browser Go no GitHub.

## 🚀 Setup Inicial (Primeira Vez)

### Opção 1: Setup Automático (Recomendado)

```bash
# Linux/macOS
./scripts/setup-github.sh
```

Este script irá:
1. ✅ Inicializar repositório Git
2. ✅ Configurar remote do GitHub
3. ✅ Criar commit inicial
4. ✅ Fazer primeiro push

### Opção 2: Setup Manual

#### 1. Criar Repositório no GitHub

Acesse: https://github.com/new

**Configurações:**
- Owner: `peder1981`
- Repository name: `bagus-browser-go`
- Description: `Browser seguro e multiplataforma focado em privacidade - Versão Go`
- Visibility: **Public**
- ❌ NÃO inicialize com README, .gitignore ou LICENSE (já temos)

#### 2. Configurar Git Local

```bash
# Inicializar repositório (se ainda não fez)
git init

# Configurar branch principal
git branch -M main

# Adicionar remote (escolha SSH ou HTTPS)

# SSH (recomendado)
git remote add origin git@github.com:peder1981/bagus-browser-go.git

# OU HTTPS
git remote add origin https://github.com/peder1981/bagus-browser-go.git
```

#### 3. Primeiro Commit e Push

```bash
# Adicionar todos os arquivos
git add .

# Commit inicial
git commit -m "Initial commit: Bagus Browser Go v2.0.0-alpha

- Estrutura completa do projeto
- Motor do navegador básico
- Sistema de armazenamento
- Testes unitários (29 testes)
- Documentação completa
- Scripts de build/teste multiplataforma
- Manuais de usuário e desenvolvedor"

# Push inicial
git push -u origin main
```

## 📤 Publicações Futuras

Após o setup inicial, use os scripts de publicação rápida:

### Linux/macOS

```bash
# Com mensagem personalizada
./scripts/publish.sh "feat: adiciona nova funcionalidade"

# Sem mensagem (usa timestamp automático)
./scripts/publish.sh
```

### Windows (CMD)

```cmd
REM Com mensagem
scripts\publish.bat feat: adiciona nova funcionalidade

REM Sem mensagem
scripts\publish.bat
```

### Windows (PowerShell)

```powershell
# Com mensagem
.\scripts\publish.ps1 -Message "feat: adiciona nova funcionalidade"

# Sem mensagem
.\scripts\publish.ps1
```

## 🔍 O que o Script de Publicação Faz

1. **Validação de Qualidade**
   - ✅ Formata código (gofmt)
   - ✅ Executa lint (golangci-lint)
   - ✅ Roda todos os testes
   - ✅ Compila o projeto

2. **Git Operations**
   - ✅ Adiciona arquivos modificados
   - ✅ Cria commit com mensagem
   - ✅ Faz push para GitHub

3. **Feedback Visual**
   - ✅ Mostra status de cada etapa
   - ✅ Cores para fácil identificação
   - ✅ Mensagens de erro claras

## 🔑 Configurar SSH (Recomendado)

Se você escolheu SSH, configure sua chave:

### 1. Gerar Chave SSH

```bash
# Gerar nova chave
ssh-keygen -t ed25519 -C "seu-email@example.com"

# Pressione Enter para aceitar o local padrão
# Digite uma senha (opcional mas recomendado)
```

### 2. Adicionar Chave ao SSH Agent

```bash
# Iniciar ssh-agent
eval "$(ssh-agent -s)"

# Adicionar chave
ssh-add ~/.ssh/id_ed25519
```

### 3. Adicionar Chave ao GitHub

```bash
# Copiar chave pública
cat ~/.ssh/id_ed25519.pub
# Copie o output
```

1. Acesse: https://github.com/settings/keys
2. Clique em **New SSH key**
3. Cole a chave copiada
4. Clique em **Add SSH key**

### 4. Testar Conexão

```bash
ssh -T git@github.com
# Deve mostrar: "Hi peder1981! You've successfully authenticated..."
```

## 📋 Convenções de Commit

Use mensagens de commit semânticas:

```bash
# Nova funcionalidade
./scripts/publish.sh "feat: adiciona modo leitura"

# Correção de bug
./scripts/publish.sh "fix: corrige navegação de abas"

# Documentação
./scripts/publish.sh "docs: atualiza manual do usuário"

# Refatoração
./scripts/publish.sh "refactor: melhora estrutura do engine"

# Testes
./scripts/publish.sh "test: adiciona testes de integração"

# Performance
./scripts/publish.sh "perf: otimiza carregamento de páginas"

# Build/CI
./scripts/publish.sh "build: atualiza scripts de build"

# Outros
./scripts/publish.sh "chore: atualiza dependências"
```

## 🔧 Configurações Recomendadas no GitHub

Após publicar, configure no GitHub:

### 1. Proteção de Branch

Settings → Branches → Add rule

- Branch name pattern: `main`
- ✅ Require pull request reviews before merging
- ✅ Require status checks to pass before merging
- ✅ Require branches to be up to date before merging

### 2. Topics/Tags

Adicione topics para melhor descoberta:

- `browser`
- `privacy`
- `security`
- `golang`
- `cross-platform`
- `webview`
- `open-source`

### 3. About

Preencha a descrição:
```
Browser seguro e multiplataforma focado em privacidade - Versão Go
```

Website: `https://bagus-browser.dev` (quando disponível)

### 4. GitHub Actions (Futuro)

Crie `.github/workflows/ci.yml` para CI/CD automático.

## ❓ Troubleshooting

### Erro: "Permission denied (publickey)"

**Solução:** Configure SSH corretamente (veja seção acima)

### Erro: "remote: Repository not found"

**Solução:** 
1. Verifique se o repositório existe no GitHub
2. Verifique se você tem permissão de acesso
3. Verifique se o remote está configurado corretamente:
   ```bash
   git remote -v
   ```

### Erro: "Updates were rejected"

**Solução:**
```bash
# Puxar mudanças primeiro
git pull origin main --rebase

# Depois fazer push
git push origin main
```

### Testes Falhando

O script perguntará se você quer continuar. Recomendamos:
1. Corrigir os testes primeiro
2. Ou fazer commit em uma branch separada

## 📚 Recursos Adicionais

- [GitHub Docs](https://docs.github.com)
- [Git Cheat Sheet](https://education.github.com/git-cheat-sheet-education.pdf)
- [Conventional Commits](https://www.conventionalcommits.org/)

---

**Pronto para publicar! 🚀**

Execute: `./scripts/publish.sh "feat: publicação inicial"`
