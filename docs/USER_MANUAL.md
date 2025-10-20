# Manual do Usuário - Bagus Browser Go

Bem-vindo ao Bagus Browser Go! Este manual irá guiá-lo através de todas as funcionalidades do navegador.

## 📋 Índice

- [Introdução](#introdução)
- [Instalação](#instalação)
- [Primeiros Passos](#primeiros-passos)
- [Interface](#interface)
- [Navegação](#navegação)
- [Abas](#abas)
- [Favoritos](#favoritos)
- [Histórico](#histórico)
- [Configurações](#configurações)
- [Privacidade e Segurança](#privacidade-e-segurança)
- [Atalhos de Teclado](#atalhos-de-teclado)
- [Solução de Problemas](#solução-de-problemas)
- [FAQ](#faq)

## Introdução

### O que é o Bagus Browser?

O Bagus Browser Go é um navegador web focado em **privacidade** e **segurança**, desenvolvido para usuários que valorizam o controle sobre seus dados pessoais.

### Características Principais

- 🔒 **Privacidade Total**: Seus dados nunca saem do seu computador
- 🛡️ **Segurança Avançada**: Criptografia AES-256 e bloqueio de conteúdo malicioso
- 🚀 **Performance**: Rápido e leve, consome menos memória
- 🌍 **Multiplataforma**: Funciona em Linux, Windows e macOS
- 🎨 **Interface Limpa**: Design minimalista e intuitivo

### Requisitos do Sistema

#### Linux
- Ubuntu 20.04+ / Fedora 35+ / Arch Linux
- 2GB RAM (recomendado 4GB)
- 200MB espaço em disco

#### Windows
- Windows 10/11
- 2GB RAM (recomendado 4GB)
- 200MB espaço em disco

#### macOS
- macOS 11 Big Sur ou superior
- 2GB RAM (recomendado 4GB)
- 200MB espaço em disco

## Instalação

### Linux

#### Método 1: Download Direto

```bash
# Baixar binário
wget https://github.com/peder1981/bagus-browser-go/releases/latest/download/bagus-linux-amd64

# Tornar executável
chmod +x bagus-linux-amd64

# Mover para /usr/local/bin (opcional)
sudo mv bagus-linux-amd64 /usr/local/bin/bagus

# Executar
bagus
```

#### Método 2: Compilar do Código

```bash
git clone https://github.com/peder1981/bagus-browser-go.git
cd bagus-browser-go
make build
./build/bagus
```

### Windows

#### Método 1: Download Direto

1. Acesse: https://github.com/peder1981/bagus-browser-go/releases/latest
2. Baixe `bagus-windows-amd64.exe`
3. Execute o arquivo

#### Método 2: PowerShell

```powershell
# Baixar
Invoke-WebRequest -Uri "https://github.com/peder1981/bagus-browser-go/releases/latest/download/bagus-windows-amd64.exe" -OutFile "bagus.exe"

# Executar
.\bagus.exe
```

### macOS

```bash
# Baixar
wget https://github.com/peder1981/bagus-browser-go/releases/latest/download/bagus-darwin-amd64

# Tornar executável
chmod +x bagus-darwin-amd64

# Executar
./bagus-darwin-amd64
```

## Primeiros Passos

### Primeira Execução

1. **Inicie o Bagus Browser**
   - Linux/macOS: `./bagus`
   - Windows: Clique duplo em `bagus.exe`

2. **Configuração Inicial**
   - Escolha seu idioma (Português, Inglês, Espanhol)
   - Configure preferências de privacidade
   - Opcionalmente, importe favoritos de outro navegador

3. **Página Inicial**
   - Por padrão, abre uma página de boas-vindas
   - Configure sua página inicial em Configurações

### Criando um Perfil

Perfis permitem separar diferentes contextos de navegação:

1. Clique em **Menu** → **Perfis**
2. Clique em **Novo Perfil**
3. Digite um nome (ex: "Trabalho", "Pessoal")
4. Opcionalmente, defina uma senha para criptografia

### Importando Dados

Importe favoritos e histórico de outros navegadores:

1. **Menu** → **Importar Dados**
2. Selecione o navegador de origem
3. Escolha o que importar:
   - ☑ Favoritos
   - ☑ Histórico
   - ☑ Senhas (criptografadas)
4. Clique em **Importar**

## Interface

### Elementos da Interface

```
┌─────────────────────────────────────────────────────────┐
│  [←] [→] [↻] [🏠]  [Barra de Endereço]  [⭐] [☰]      │ ← Barra de Navegação
├─────────────────────────────────────────────────────────┤
│  [Aba 1] [Aba 2] [+]                                   │ ← Abas
├─────────────────────────────────────────────────────────┤
│                                                         │
│                                                         │
│              Conteúdo da Página                        │
│                                                         │
│                                                         │
└─────────────────────────────────────────────────────────┘
```

### Barra de Navegação

- **← Voltar**: Volta para a página anterior
- **→ Avançar**: Avança para a próxima página
- **↻ Recarregar**: Recarrega a página atual
- **🏠 Início**: Vai para a página inicial
- **⭐ Favoritos**: Adiciona aos favoritos
- **☰ Menu**: Abre o menu principal

### Barra de Endereço

- Digite URLs ou termos de busca
- Sugestões automáticas baseadas em histórico
- Indicador de segurança (🔒 para HTTPS)
- Bloqueador de rastreadores ativo (🛡️)

## Navegação

### Navegação Básica

#### Acessar um Site

1. Clique na barra de endereço
2. Digite a URL (ex: `example.com`)
3. Pressione **Enter**

#### Buscar na Web

1. Digite sua busca na barra de endereço
2. Pressione **Enter**
3. O mecanismo de busca padrão será usado

#### Configurar Mecanismo de Busca

1. **Menu** → **Configurações** → **Busca**
2. Escolha entre:
   - DuckDuckGo (padrão - privado)
   - Google
   - Bing
   - Startpage

### Navegação por Abas

#### Abrir Nova Aba
- Clique no botão **[+]**
- Ou pressione **Ctrl+T** (Cmd+T no macOS)

#### Fechar Aba
- Clique no **[×]** na aba
- Ou pressione **Ctrl+W** (Cmd+W no macOS)

#### Alternar entre Abas
- **Ctrl+Tab**: Próxima aba
- **Ctrl+Shift+Tab**: Aba anterior
- **Ctrl+1-8**: Ir para aba específica
- **Ctrl+9**: Última aba

### Histórico de Navegação

#### Voltar/Avançar
- Clique nos botões **←** ou **→**
- Ou use **Alt+←** e **Alt+→**

#### Ver Histórico Completo
1. **Menu** → **Histórico**
2. Navegue por data ou busque
3. Clique em um item para abrir

#### Limpar Histórico
1. **Menu** → **Histórico** → **Limpar Histórico**
2. Escolha o período:
   - Última hora
   - Últimas 24 horas
   - Últimos 7 dias
   - Todo o período
3. Selecione o que limpar:
   - ☑ Histórico de navegação
   - ☑ Cookies
   - ☑ Cache
   - ☑ Senhas salvas
4. Clique em **Limpar**

## Abas

### Gerenciamento de Abas

#### Abrir em Nova Aba
- **Ctrl+Clique** em um link
- Ou **Clique com botão direito** → **Abrir em Nova Aba**

#### Reabrir Aba Fechada
- **Ctrl+Shift+T** (Cmd+Shift+T no macOS)
- Ou **Menu** → **Histórico** → **Reabrir Aba**

#### Fixar Aba
1. **Clique com botão direito** na aba
2. Selecione **Fixar Aba**
3. A aba ficará menor e sempre visível

#### Duplicar Aba
1. **Clique com botão direito** na aba
2. Selecione **Duplicar Aba**

#### Mover Aba
- Arraste a aba para a posição desejada
- Arraste para fora para criar nova janela

### Grupos de Abas

Organize abas relacionadas em grupos:

1. **Clique com botão direito** em uma aba
2. Selecione **Adicionar ao Grupo**
3. Escolha um grupo existente ou crie novo
4. Dê um nome e cor ao grupo

## Favoritos

### Adicionar Favorito

#### Método 1: Botão Estrela
1. Navegue até a página
2. Clique no ícone **⭐** na barra de endereço
3. Edite o nome (opcional)
4. Escolha uma pasta
5. Clique em **Salvar**

#### Método 2: Menu
1. **Menu** → **Favoritos** → **Adicionar aos Favoritos**
2. Configure e salve

#### Método 3: Atalho
- Pressione **Ctrl+D** (Cmd+D no macOS)

### Organizar Favoritos

#### Criar Pasta
1. **Menu** → **Favoritos** → **Gerenciar Favoritos**
2. Clique em **Nova Pasta**
3. Digite o nome
4. Arraste favoritos para a pasta

#### Editar Favorito
1. **Clique com botão direito** no favorito
2. Selecione **Editar**
3. Modifique nome ou URL
4. Clique em **Salvar**

#### Excluir Favorito
1. **Clique com botão direito** no favorito
2. Selecione **Excluir**

### Barra de Favoritos

Ative a barra de favoritos para acesso rápido:

1. **Menu** → **Visualizar** → **Barra de Favoritos**
2. Ou pressione **Ctrl+Shift+B**

## Histórico

### Visualizar Histórico

1. **Menu** → **Histórico**
2. Ou pressione **Ctrl+H** (Cmd+H no macOS)

### Buscar no Histórico

1. Abra o histórico
2. Digite na caixa de busca
3. Filtre por data, site ou título

### Exportar Histórico

1. **Menu** → **Histórico** → **Exportar**
2. Escolha o formato (JSON, CSV, HTML)
3. Selecione o local para salvar

## Configurações

### Acessar Configurações

- **Menu** → **Configurações**
- Ou digite na barra: `bagus://settings`

### Categorias de Configurações

#### Geral
- **Página Inicial**: Configure qual página abre ao iniciar
- **Ao Iniciar**: Abrir página inicial, continuar de onde parou, ou páginas específicas
- **Downloads**: Local padrão para downloads
- **Idioma**: Escolha o idioma da interface

#### Aparência
- **Tema**: Claro, Escuro, ou Automático (segue o sistema)
- **Fonte**: Tamanho e família da fonte
- **Zoom**: Nível de zoom padrão
- **Barra de Favoritos**: Sempre visível, oculta, ou apenas em nova aba

#### Privacidade
- **Bloqueio de Rastreadores**: Ativado por padrão
- **Cookies**: Aceitar todos, bloquear terceiros, ou bloquear todos
- **Não Rastrear (DNT)**: Enviar cabeçalho DNT
- **Limpar ao Fechar**: Limpar dados ao fechar o navegador

#### Segurança
- **HTTPS-Only**: Sempre usar HTTPS quando disponível
- **Bloqueio de Sites Maliciosos**: Ativado por padrão
- **Avisos de Segurança**: Mostrar avisos para sites inseguros
- **Senhas**: Gerenciar senhas salvas (criptografadas)

#### Busca
- **Mecanismo Padrão**: DuckDuckGo, Google, Bing, etc.
- **Sugestões**: Ativar/desativar sugestões de busca
- **Mecanismos Personalizados**: Adicionar mecanismos customizados

#### Avançado
- **Hardware Acceleration**: Usar aceleração por hardware
- **Prefetch**: Pré-carregar páginas para navegação mais rápida
- **Proxy**: Configurar proxy manual
- **Certificados**: Gerenciar certificados SSL

## Privacidade e Segurança

### Modo Privado

Navegue sem salvar histórico, cookies ou dados:

1. **Menu** → **Nova Janela Privada**
2. Ou pressione **Ctrl+Shift+N** (Cmd+Shift+N no macOS)

**O que é protegido:**
- ✅ Histórico não é salvo
- ✅ Cookies são descartados ao fechar
- ✅ Cache não é armazenado
- ✅ Formulários não são lembrados

**O que NÃO é protegido:**
- ❌ Seu provedor de internet ainda vê seu tráfego
- ❌ Sites visitados podem rastrear você
- ❌ Downloads são salvos normalmente

### Bloqueador de Conteúdo

O Bagus Browser bloqueia automaticamente:

- 🚫 **Rastreadores**: Scripts de tracking
- 🚫 **Anúncios**: Publicidade intrusiva
- 🚫 **Malware**: Sites e downloads maliciosos
- 🚫 **Phishing**: Tentativas de roubo de dados

#### Ver Sites Bloqueados

1. Clique no ícone **🛡️** na barra de endereço
2. Veja estatísticas de bloqueio
3. Opcionalmente, desative para o site atual

#### Adicionar Exceção

1. Clique no ícone **🛡️**
2. Clique em **Desativar Proteção para Este Site**
3. Recarregue a página

### Criptografia de Dados

Todos os dados sensíveis são criptografados:

- 🔐 **Senhas**: AES-256-GCM
- 🔐 **Histórico**: Opcional, ative nas configurações
- 🔐 **Favoritos**: Opcional
- 🔐 **Cookies**: Opcional

#### Ativar Criptografia

1. **Menu** → **Configurações** → **Privacidade**
2. Ative **Criptografar Dados Locais**
3. Defina uma senha mestra
4. **Importante**: Não perca esta senha!

### Gerenciador de Senhas

#### Salvar Senha

1. Faça login em um site
2. Clique em **Salvar** quando solicitado
3. A senha é criptografada e armazenada

#### Ver Senhas Salvas

1. **Menu** → **Configurações** → **Senhas**
2. Digite sua senha mestra (se configurada)
3. Veja, edite ou exclua senhas

#### Exportar Senhas

1. **Configurações** → **Senhas** → **Exportar**
2. Escolha o formato (CSV criptografado)
3. Salve em local seguro

## Atalhos de Teclado

### Navegação

| Ação | Windows/Linux | macOS |
|------|---------------|-------|
| Nova aba | Ctrl+T | Cmd+T |
| Fechar aba | Ctrl+W | Cmd+W |
| Reabrir aba | Ctrl+Shift+T | Cmd+Shift+T |
| Próxima aba | Ctrl+Tab | Cmd+Option+→ |
| Aba anterior | Ctrl+Shift+Tab | Cmd+Option+← |
| Nova janela | Ctrl+N | Cmd+N |
| Janela privada | Ctrl+Shift+N | Cmd+Shift+N |
| Fechar janela | Ctrl+Shift+W | Cmd+Shift+W |

### Página

| Ação | Windows/Linux | macOS |
|------|---------------|-------|
| Voltar | Alt+← | Cmd+[ |
| Avançar | Alt+→ | Cmd+] |
| Recarregar | Ctrl+R | Cmd+R |
| Recarregar (cache) | Ctrl+Shift+R | Cmd+Shift+R |
| Parar | Esc | Esc |
| Início | Alt+Home | Cmd+Home |
| Fim da página | End | End |
| Topo da página | Home | Home |

### Zoom

| Ação | Windows/Linux | macOS |
|------|---------------|-------|
| Aumentar | Ctrl++ | Cmd++ |
| Diminuir | Ctrl+- | Cmd+- |
| Resetar | Ctrl+0 | Cmd+0 |

### Outros

| Ação | Windows/Linux | macOS |
|------|---------------|-------|
| Buscar na página | Ctrl+F | Cmd+F |
| Imprimir | Ctrl+P | Cmd+P |
| Salvar página | Ctrl+S | Cmd+S |
| Favoritos | Ctrl+D | Cmd+D |
| Histórico | Ctrl+H | Cmd+H |
| Downloads | Ctrl+J | Cmd+J |
| Configurações | Ctrl+, | Cmd+, |
| Ferramentas Dev | F12 | F12 |

## Solução de Problemas

### O navegador não inicia

**Solução:**
```bash
# Linux/macOS
./bagus --debug

# Windows
bagus.exe --debug
```

Verifique os logs em:
- Linux: `~/.bagus/logs/`
- Windows: `%APPDATA%\bagus\logs\`
- macOS: `~/Library/Application Support/bagus/logs/`

### Página não carrega

1. Verifique sua conexão com a internet
2. Tente recarregar: **Ctrl+R**
3. Limpe o cache: **Menu** → **Limpar Dados**
4. Desative extensões temporariamente

### Navegador está lento

1. Feche abas não utilizadas
2. Limpe cache e cookies
3. Desative aceleração por hardware (Configurações → Avançado)
4. Verifique uso de memória no gerenciador de tarefas

### Senhas não são salvas

1. Verifique se o gerenciador está ativado
2. **Configurações** → **Senhas** → Ativar
3. Certifique-se de que o site permite salvar senhas

### Resetar Configurações

```bash
# Backup primeiro!
cp -r ~/.bagus ~/.bagus.backup

# Resetar
rm -rf ~/.bagus
```

## FAQ

### O Bagus Browser é gratuito?

Sim, é 100% gratuito e open source (licença MIT).

### Meus dados são enviados para algum servidor?

Não. Todos os dados ficam armazenados localmente no seu computador. Zero telemetria.

### Posso usar extensões?

Atualmente não, mas o suporte a extensões está planejado para versão 2.1.0.

### Como atualizar o navegador?

1. **Menu** → **Sobre**
2. Clique em **Verificar Atualizações**
3. Ou baixe manualmente a versão mais recente

### Posso sincronizar entre dispositivos?

Não na versão atual. Sincronização opcional (criptografada) está planejada para o futuro.

### O navegador suporta VPN?

O navegador respeita configurações de VPN do sistema. Não há VPN integrada.

### Como reportar bugs?

1. Acesse: https://github.com/peder1981/bagus-browser-go/issues
2. Clique em **New Issue**
3. Descreva o problema detalhadamente

### Onde obter suporte?

- **GitHub Issues**: Para bugs
- **Discussions**: Para perguntas
- **Email**: support@bagus-browser.dev

---

**Versão do Manual**: 2.0.0-alpha  
**Última Atualização**: 2024-10-20

**Aproveite sua navegação privada e segura! 🔒**
