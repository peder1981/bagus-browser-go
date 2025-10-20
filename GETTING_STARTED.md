# Getting Started - Bagus Browser Go

## 🎯 Bem-vindo ao Projeto

Este é o **Bagus Browser Go** - uma reescrita multiplataforma do Bagus Browser Python.

## 📁 Estrutura de Projetos

Você agora tem **dois projetos separados**:

### 1. Versão Python (Original)
```
~/bagus_browser_python/
```
- **Repositório**: https://github.com/peder1981/bagus_browser
- **Status**: Produção, funcional
- **Plataforma**: Linux apenas
- **Tecnologia**: Python + PySide6

### 2. Versão Go (Nova - Este Projeto)
```
~/bagus-browser-go/
```
- **Repositório**: Ainda não criado
- **Status**: Desenvolvimento inicial (Alpha)
- **Plataforma**: Linux, Windows, macOS
- **Tecnologia**: Go + WebView

## 🚀 Primeiros Passos

### 1. Instalar Go

Se ainda não tem Go instalado:

```bash
# Ubuntu/Debian
sudo apt update
sudo apt install golang-go

# Ou baixe a versão mais recente
wget https://go.dev/dl/go1.22.0.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.22.0.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
```

### 2. Verificar Instalação

```bash
go version  # Deve mostrar Go 1.21 ou superior
```

### 3. Instalar Dependências do Projeto

```bash
cd ~/bagus-browser-go
make deps
```

### 4. Compilar o Projeto

```bash
# Build para sua plataforma atual
make build

# Ou build para plataformas específicas
make build-linux    # Linux (amd64, arm64)
make build-windows  # Windows (amd64)
make build-macos    # macOS (amd64, arm64)

# Build para todas as plataformas
make build-all
```

### 5. Executar

```bash
# Executar diretamente
./build/bagus

# Ou usar o Makefile
make run

# Modo debug
make run-debug
```

## 📝 Comandos Úteis

```bash
# Ver todos os comandos disponíveis
make help

# Executar testes
make test

# Executar testes com coverage
make test-coverage

# Formatar código
make fmt

# Lint (verificar qualidade do código)
make lint

# Limpar arquivos de build
make clean

# Criar pacotes de distribuição
make dist
```

## 🔧 Desenvolvimento

### Estrutura do Código

```
bagus-browser-go/
├── cmd/bagus/          # Entry point (main.go)
├── internal/           # Código privado
│   ├── browser/        # Engine do navegador
│   ├── crypto/         # Criptografia
│   ├── security/       # Segurança
│   ├── ui/             # Interface
│   └── storage/        # Armazenamento
├── pkg/                # Código público/reutilizável
└── docs/               # Documentação
```

### Adicionar Nova Funcionalidade

1. Crie os arquivos necessários em `internal/`
2. Implemente a lógica
3. Adicione testes em `tests/`
4. Execute `make test`
5. Commit e push

### Exemplo: Adicionar Novo Módulo

```go
// internal/mymodule/mymodule.go
package mymodule

func DoSomething() error {
    // Sua implementação
    return nil
}
```

## 🐛 Debugging

### Modo Debug

```bash
./build/bagus --debug
```

### Logs

Os logs são escritos em:
- **Linux**: `~/.bagus/logs/`
- **Windows**: `%APPDATA%\bagus\logs\`
- **macOS**: `~/Library/Application Support/bagus/logs/`

## 📦 Criar Release

### Local

```bash
make dist
# Pacotes criados em dist/
```

### GitHub (Automático)

```bash
# Criar tag
git tag -a v2.0.0-alpha -m "Release v2.0.0-alpha"
git push origin v2.0.0-alpha

# GitHub Actions vai automaticamente:
# 1. Compilar para todas as plataformas
# 2. Executar testes
# 3. Criar release
# 4. Fazer upload dos binários
```

## 🔗 Configurar Repositório GitHub

### 1. Criar Repositório

Acesse: https://github.com/new

Configurações:
- **Owner**: peder1981
- **Repository name**: bagus-browser-go
- **Description**: Browser seguro e multiplataforma focado em privacidade - Versão Go
- **Visibility**: Public
- **NÃO** inicialize com README, .gitignore ou LICENSE (já temos)

### 2. Adicionar Remote

```bash
cd ~/bagus-browser-go

# SSH (recomendado)
git remote add origin git@github.com:peder1981/bagus-browser-go.git

# Ou HTTPS
git remote add origin https://github.com/peder1981/bagus-browser-go.git
```

### 3. Push Inicial

```bash
git push -u origin main
```

## 📚 Documentação

- [Arquitetura](docs/ARCHITECTURE.md) - Visão geral da arquitetura
- [README](README.md) - Informações gerais do projeto

## 🎯 Roadmap de Desenvolvimento

### Fase Atual: Alpha (Estrutura Base)
- [x] Estrutura de projeto
- [x] Storage manager
- [x] Browser engine básico
- [ ] WebView integration
- [ ] UI básica

### Próximas Fases
- [ ] Sistema de criptografia completo
- [ ] Bloqueador de conteúdo
- [ ] Interface gráfica completa
- [ ] Sistema de extensões
- [ ] Testes de integração

## 💡 Dicas

1. **Use o Makefile**: Todos os comandos comuns estão lá
2. **Rode testes frequentemente**: `make test`
3. **Mantenha código formatado**: `make fmt`
4. **Verifique lint antes de commit**: `make lint`
5. **Leia a documentação**: Especialmente [ARCHITECTURE.md](docs/ARCHITECTURE.md)

## 🆘 Problemas Comuns

### "go: command not found"
```bash
# Instale Go ou adicione ao PATH
export PATH=$PATH:/usr/local/go/bin
```

### "permission denied" ao executar
```bash
chmod +x build/bagus
```

### Erros de dependências
```bash
make clean
make deps
```

## 📞 Suporte

- **Issues**: https://github.com/peder1981/bagus-browser-go/issues
- **Versão Python**: https://github.com/peder1981/bagus_browser

---

**Pronto para começar! 🚀**

Execute `make help` para ver todos os comandos disponíveis.
