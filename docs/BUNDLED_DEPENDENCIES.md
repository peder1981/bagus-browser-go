# 📦 Estratégia de Dependências Embarcadas - Bagus Browser

## 🎯 Objetivo

Embarcar todas as dependências (GTK3, WebKit2GTK, GStreamer) nos pacotes de distribuição para garantir:

- ✅ **Robustez máxima** - Funciona em qualquer distribuição Linux
- ✅ **Compatibilidade** - Sem conflitos de versão
- ✅ **Simplicidade** - Usuário não precisa instalar nada
- ✅ **Segurança** - Versões testadas e validadas

## 📋 Dependências Embarcadas

### Principais
- **libgtk-3-0** - Interface gráfica
- **libwebkit2gtk-4.0-37** - Engine de navegação
- **libgstreamer1.0-0** - Multimídia
- **libgstreamer-plugins-base1.0-0** - Plugins base
- **libgstreamer-plugins-good1.0-0** - Plugins bons
- **libgstreamer-plugins-bad1.0-0** - Plugins avançados
- **libgstreamer-gl1.0-0** - Aceleração gráfica

### Suporte (50+ bibliotecas)
- Libs de sistema (libc, libssl, etc)
- Libs de renderização (cairo, pango, etc)
- Libs de mídia (libwebp, libjpeg, etc)
- Libs de desenvolvimento (libffi, libpcre, etc)

## 🚀 Workflow de Build

### 1. Compilar Binário
```bash
./scripts/bagus build
```

Gera:
- `build/bagus-browser` - Binário compilado
- `dist/bagus-browser_v5.0.0_amd64.deb` - .deb padrão
- `dist/bagus-browser_v5.0.0_linux_amd64.tar.gz` - Tarball

### 2. Embarcar Dependências
```bash
./scripts/bagus bundle
```

Gera:
- `build/bundle/lib/` - ~200 arquivos .so
- `build/bundle/include/` - Headers
- `build/bundle/lib/pkgconfig/` - Arquivos .pc
- `build/bundle/setup-env.sh` - Script de configuração
- `build/bundle/README.md` - Documentação

### 3. Criar .deb Bundled
```bash
./scripts/build-deb-bundled.sh v5.0.0
```

Gera:
- `dist/bagus-browser_v5.0.0_amd64_bundled.deb` - .deb com dependências

### 4. Testar Pacote
```bash
./scripts/test-bundled.sh
```

Valida:
- Estrutura do bundle
- Bibliotecas críticas
- .pc files
- Permissões
- Tamanho

## 📦 Estrutura do Pacote .deb Bundled

```
bagus-browser_v5.0.0_amd64_bundled.deb
│
├── usr/bin/
│   └── bagus-browser              # Binário principal
│
├── usr/lib/bagus-browser/
│   └── bundle/                    # Dependências embarcadas
│       ├── lib/                   # Bibliotecas compartilhadas
│       ├── include/               # Headers
│       ├── lib/pkgconfig/         # Arquivos .pc
│       ├── setup-env.sh           # Script de configuração
│       └── README.md              # Documentação
│
├── usr/share/applications/
│   └── bagus-browser.desktop      # Atalho de aplicação
│
├── usr/share/icons/
│   └── hicolor/                   # Ícones
│
└── DEBIAN/
    ├── control                    # Metadados
    ├── preinst                    # Verificações pré-instalação
    ├── postinst                   # Configuração pós-instalação
    └── postrm                     # Limpeza pós-remoção
```

## 🔧 Scripts Utilizados

### `scripts/bundle-dependencies.sh`
Coleta todas as dependências do sistema e as copia para `build/bundle/`.

**Uso:**
```bash
./scripts/bundle-dependencies.sh build/bundle
```

**O que faz:**
1. Cria estrutura de diretórios
2. Copia bibliotecas .so
3. Copia headers
4. Copia .pc files
5. Cria setup-env.sh
6. Gera README.md

### `scripts/build-deb-bundled.sh`
Cria pacote .deb com dependências embarcadas.

**Uso:**
```bash
./scripts/build-deb-bundled.sh v5.0.0
```

**O que faz:**
1. Verifica binário e bundle
2. Cria estrutura .deb
3. Copia binário e dependências
4. Cria scripts de instalação
5. Valida glibc mínima
6. Gera checksums

### `scripts/test-bundled.sh`
Testa integridade do pacote bundled.

**Uso:**
```bash
./scripts/test-bundled.sh
```

**Testes:**
1. Estrutura do bundle
2. Bibliotecas críticas
3. .pc files
4. setup-env.sh
5. Carregamento de libs
6. Tamanho
7. Permissões
8. Documentação

### `scripts/bagus-browser-wrapper.sh`
Wrapper que detecta e usa dependências embarcadas.

**Uso:**
```bash
/usr/bin/bagus-browser-wrapper
```

**Prioridade:**
1. Dependências embarcadas em `/opt/bagus-browser/bundle`
2. WebKit compilado em `/opt/webkitgtk-webrtc`
3. Dependências do sistema

## 📊 Tamanho do Pacote

| Componente | Tamanho |
|-----------|---------|
| Binário | ~5 MB |
| Dependências | ~200-300 MB |
| Comprimido (.deb) | ~50-80 MB |
| **Total** | **~50-80 MB** |

## 🔒 Verificações de Segurança

### Pré-instalação (preinst)
- Verifica glibc >= 2.35
- Valida compatibilidade do sistema

### Pós-instalação (postinst)
- Atualiza cache de ícones
- Atualiza banco de dados de aplicações
- Cria symlinks necessários

### Pós-remoção (postrm)
- Remove symlinks
- Limpa cache de ícones
- Remove banco de dados

## 🚀 Instalação do Usuário

### Opção 1: Via .deb
```bash
wget https://github.com/peder1981/bagus-browser-go/releases/download/v5.0.0/bagus-browser_v5.0.0_amd64_bundled.deb
sudo dpkg -i bagus-browser_v5.0.0_amd64_bundled.deb
bagus-browser
```

### Opção 2: Via Tarball
```bash
wget https://github.com/peder1981/bagus-browser-go/releases/download/v5.0.0/bagus-browser_v5.0.0_linux_amd64_bundled.tar.gz
tar -xzf bagus-browser_v5.0.0_linux_amd64_bundled.tar.gz
cd bagus-browser-bundle
source setup-env.sh
./bagus-browser
```

## 🔍 Troubleshooting

### Erro: "libsoup3 symbols detected"
**Causa:** Conflito entre libsoup2 e libsoup3

**Solução:** Usar dependências embarcadas
```bash
source /usr/lib/bagus-browser/bundle/setup-env.sh
bagus-browser
```

### Erro: "glibc version not found"
**Causa:** Sistema com glibc < 2.35

**Solução:** Atualizar sistema ou compilar WebKit localmente

### Erro: "Cannot find library"
**Causa:** Dependência embarcada corrompida

**Solução:** Reinstalar pacote
```bash
sudo dpkg --remove bagus-browser
sudo dpkg -i bagus-browser_v5.0.0_amd64_bundled.deb
```

## 📝 Checklist de Release

- [ ] Compilar: `./scripts/bagus build`
- [ ] Embarcar: `./scripts/bagus bundle`
- [ ] Testar bundle: `./scripts/test-bundled.sh`
- [ ] Criar .deb: `./scripts/build-deb-bundled.sh v5.0.0`
- [ ] Verificar tamanho: `ls -lh dist/`
- [ ] Testar instalação: `sudo dpkg -i dist/*.deb`
- [ ] Testar execução: `bagus-browser`
- [ ] Gerar checksums: `sha256sum dist/*`
- [ ] Publicar: `./scripts/bagus publish`

## 🎯 Benefícios

### Para Usuários
- ✅ Instalação simples (um clique)
- ✅ Sem dependências externas
- ✅ Funciona em qualquer distribuição
- ✅ Sem conflitos de versão

### Para Desenvolvedores
- ✅ Menos suporte necessário
- ✅ Menos bugs relacionados a dependências
- ✅ Versões testadas e validadas
- ✅ Distribuição simplificada

### Para Distribuições
- ✅ Pacote auto-contido
- ✅ Sem impacto em outras aplicações
- ✅ Fácil remoção
- ✅ Compatibilidade garantida

## 🔄 Atualização de Dependências

Para atualizar as dependências embarcadas:

1. Atualizar sistema: `sudo apt update && sudo apt upgrade`
2. Limpar bundle antigo: `rm -rf build/bundle`
3. Embarcar novas dependências: `./scripts/bagus bundle`
4. Testar: `./scripts/test-bundled.sh`
5. Criar novo .deb: `./scripts/build-deb-bundled.sh v5.0.1`

## 📚 Referências

- [GTK3 Documentation](https://developer.gnome.org/gtk3/)
- [WebKit2GTK Documentation](https://webkitgtk.org/)
- [GStreamer Documentation](https://gstreamer.freedesktop.org/)
- [Debian Packaging Guide](https://www.debian.org/doc/manuals/debian-faq/ch-pkg_basics.en.html)

---

**Última atualização:** 02/11/2025
**Versão:** 5.0.0
**Status:** ✅ Pronto para Produção
