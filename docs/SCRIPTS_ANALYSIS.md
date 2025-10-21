# Análise de Scripts - Redundâncias e Limpeza

## 📊 Scripts Encontrados

### Raiz do Projeto
1. `bagus-launcher.sh` - Launcher para escolher entre versão rápida/completa
2. `build-all.sh` - **NOVO** Build unificado multiplataforma
3. `build-deb.sh` - Cria pacote .deb
4. `build-windows.sh` - Build placeholder para Windows
5. `install.sh` - Instalador simplificado
6. `publish-release.sh` - **NOVO** Publicação de releases

### Diretório scripts/
7. `build.sh` - Build com opções de plataforma
8. `build_cef.sh` - Build versão CEF
9. `build_menu.sh` - Menu interativo de build
10. `install-icon.sh` - Instalação de ícone
11. `install_cef.sh` - Instalação do CEF
12. `publish.sh` - Publicação no GitHub
13. `setup-github.sh` - Setup inicial do GitHub
14. `test.sh` - Testes automatizados
15. `verify_privacy.sh` - Verificação de privacidade

### Scripts Windows
16. `scripts/build.bat`
17. `scripts/build.ps1`
18. `scripts/publish.bat`
19. `scripts/publish.ps1`
20. `scripts/test.bat`
21. `scripts/test.ps1`

## 🔍 Análise de Redundâncias

### ❌ REDUNDANTES - Para Remover

1. **`bagus-launcher.sh`**
   - **Motivo**: Obsoleto, referencia versões antigas (rápida/completa)
   - **Substituído por**: Executável direto `bagus-linux-amd64`

2. **`build-windows.sh`**
   - **Motivo**: Cria apenas placeholder, não funcional
   - **Substituído por**: `build-all.sh` (tenta build Windows com CGO_ENABLED=0)

3. **`scripts/build.sh`**
   - **Motivo**: Duplica funcionalidade do `build-all.sh`
   - **Substituído por**: `build-all.sh` (mais completo e atualizado)

4. **`scripts/build_menu.sh`**
   - **Motivo**: Menu para versões antigas (Webview vs CEF)
   - **Substituído por**: Versão única com BrowserSingleWindow

5. **`scripts/publish.sh`**
   - **Motivo**: Duplica funcionalidade do `publish-release.sh`
   - **Substituído por**: `publish-release.sh` (usa gh CLI)

6. **`scripts/build.bat`, `scripts/build.ps1`**
   - **Motivo**: Build Windows não funciona sem CGO nativo
   - **Nota**: Manter apenas como referência futura

7. **`scripts/publish.bat`, `scripts/publish.ps1`**
   - **Motivo**: Duplicam `publish-release.sh`
   - **Substituído por**: `publish-release.sh` ou git direto

8. **`scripts/test.bat`, `scripts/test.ps1`**
   - **Motivo**: Testes podem ser executados com `go test` direto
   - **Substituído por**: `scripts/test.sh` (mantido para Linux)

9. **`scripts/README.md`**
   - **Motivo**: Documenta scripts obsoletos
   - **Substituído por**: README.md principal + RELEASE_INSTRUCTIONS.md

## ✅ MANTER - Scripts Úteis

1. **`build-all.sh`** ⭐
   - Build unificado e atualizado
   - Suporta Linux AMD64
   - Cria pacotes de distribuição

2. **`build-deb.sh`** ⭐
   - Cria pacote .deb funcional
   - Necessário para distribuição Debian/Ubuntu

3. **`install.sh`** ⭐
   - Instalador simples e funcional
   - Compila e configura o browser

4. **`publish-release.sh`** ⭐
   - Publicação de releases no GitHub
   - Usa gh CLI

5. **`scripts/install_cef.sh`** ⭐
   - Instalação do CEF (versão completa)
   - Útil para versão 100% compatível

6. **`scripts/build_cef.sh`** ⭐
   - Build da versão CEF
   - Complementa versão webview

7. **`scripts/test.sh`** ⭐
   - Testes automatizados
   - Lint e formatação

8. **`scripts/verify_privacy.sh`** ⭐
   - Auditoria de privacidade
   - Importante para garantir zero telemetria

9. **`scripts/setup-github.sh`** ⭐
   - Setup inicial do repositório
   - Útil para novos contribuidores

10. **`scripts/install-icon.sh`** ⭐
    - Instalação de ícone do sistema
    - Integração com desktop

## 📋 Plano de Limpeza

### Fase 1: Remover Scripts Obsoletos
```bash
rm bagus-launcher.sh
rm build-windows.sh
rm scripts/build.sh
rm scripts/build_menu.sh
rm scripts/publish.sh
rm scripts/README.md
```

### Fase 2: Remover Scripts Windows (Opcional)
```bash
rm scripts/build.bat
rm scripts/build.ps1
rm scripts/publish.bat
rm scripts/publish.ps1
rm scripts/test.bat
rm scripts/test.ps1
```

### Fase 3: Organizar Documentação
- Atualizar README.md principal
- Manter apenas RELEASE_INSTRUCTIONS.md

## 📊 Resultado Final

### Scripts Mantidos (10 arquivos)
```
/
├── build-all.sh          # Build principal
├── build-deb.sh          # Pacote .deb
├── install.sh            # Instalador
├── publish-release.sh    # Publicação
└── scripts/
    ├── build_cef.sh      # Build CEF
    ├── install_cef.sh    # Instalar CEF
    ├── install-icon.sh   # Ícone sistema
    ├── setup-github.sh   # Setup Git
    ├── test.sh           # Testes
    └── verify_privacy.sh # Auditoria
```

### Redução
- **Antes**: 21 scripts
- **Depois**: 10 scripts
- **Redução**: 52% menos arquivos

## 🎯 Benefícios

1. **Menos Confusão**: Apenas scripts necessários e atualizados
2. **Manutenção Simples**: Menos arquivos para manter
3. **Documentação Clara**: Cada script tem propósito único
4. **Sem Duplicação**: Funcionalidades não se repetem
5. **Foco em Linux**: Plataforma principal e funcional

## ⚠️ Notas

- Scripts Windows removidos porque webview requer CGO nativo
- Para Windows, usuários devem compilar nativamente no Windows
- Versão CEF mantida para compatibilidade 100%
- Scripts de teste e privacidade mantidos para qualidade
