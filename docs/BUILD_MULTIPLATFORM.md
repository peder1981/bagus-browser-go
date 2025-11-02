# 🌍 Guia de Build Multiplataforma - Bagus Browser v5.0.0

## 🚀 Quick Start

### Linux (Recomendado para começar)

```bash
# Compilar para Linux x86_64
./scripts/build-multiplatform.sh --os linux --arch amd64

# Resultado: dist/bagus-browser
```

### Windows (Compilação Cruzada)

```bash
# Compilar para Windows x86_64 (do Linux)
./scripts/build-multiplatform.sh --os windows --arch amd64

# Resultado: dist/bagus-browser-windows-amd64/bagus-browser.exe
```

### macOS (Compilação Cruzada)

```bash
# Compilar para macOS x86_64 (do Linux)
./scripts/build-multiplatform.sh --os macos --arch amd64

# Compilar para macOS ARM64 (Apple Silicon)
./scripts/build-multiplatform.sh --os macos --arch arm64

# Resultado: dist/bagus-browser-macos-amd64/bagus-browser
```

### Compilar para Todos os SOs

```bash
./scripts/build-multiplatform.sh --os all

# Resultado: Binários para Linux, Windows, macOS e FreeBSD
```

---

## 📋 Instruções Detalhadas por Plataforma

### 1. Linux (✅ Pronto)

#### Requisitos
```bash
sudo apt-get install golang-go libgtk-3-dev libwebkit2gtk-4.0-dev
```

#### Build
```bash
./scripts/build-multiplatform.sh --os linux --arch amd64
```

#### Resultado
- Binário: `dist/bagus-browser` (5,5M)
- Tarball: `dist/bagus-browser_v5.0.0_linux_amd64.tar.gz`

#### Instalar
```bash
sudo cp dist/bagus-browser /usr/bin/
```

---

### 2. Windows (⚠️ Requer Configuração)

#### Opção A: Compilação Cruzada (Recomendado)

**Do Linux para Windows:**

```bash
# Instalar dependências de compilação cruzada
sudo apt-get install mingw-w64 pkg-config-mingw-w64-x86-64

# Compilar
./scripts/build-multiplatform.sh --os windows --arch amd64
```

**Resultado:** `dist/bagus-browser-windows-amd64/bagus-browser.exe`

#### Opção B: Build Nativo no Windows

**Requisitos:**
1. Instalar Go 1.21+
2. Instalar MSYS2
3. Instalar GTK3 e WebKit2GTK via MSYS2:

```bash
pacman -S mingw-w64-x86_64-gtk3 mingw-w64-x86_64-webkit2-gtk3
```

**Build:**
```bash
go build -o bagus-browser.exe ./cmd/bagus-browser
```

#### Criar Instalador

**Usando NSIS:**

```nsis
; bagus-browser.nsi
!include "MUI2.nsh"

Name "Bagus Browser v5.0.0"
OutFile "bagus-browser-installer.exe"
InstallDir "$PROGRAMFILES\Bagus Browser"

!insertmacro MUI_PAGE_DIRECTORY
!insertmacro MUI_PAGE_INSTFILES
!insertmacro MUI_LANGUAGE "Portuguese"

Section "Install"
  SetOutPath "$INSTDIR"
  File "bagus-browser.exe"
  File "*.dll"  ; GTK3 e WebKit2GTK DLLs
  
  CreateDirectory "$SMPROGRAMS\Bagus Browser"
  CreateShortcut "$SMPROGRAMS\Bagus Browser\Bagus Browser.lnk" "$INSTDIR\bagus-browser.exe"
  CreateShortcut "$DESKTOP\Bagus Browser.lnk" "$INSTDIR\bagus-browser.exe"
SectionEnd
```

---

### 3. macOS (⚠️ Requer Configuração)

#### Opção A: Compilação Cruzada (Recomendado)

**Do Linux para macOS:**

```bash
# Instalar ferramentas de compilação cruzada
brew install osxcross

# Compilar para Intel
./scripts/build-multiplatform.sh --os macos --arch amd64

# Compilar para Apple Silicon
./scripts/build-multiplatform.sh --os macos --arch arm64
```

#### Opção B: Build Nativo no macOS

**Requisitos:**
```bash
brew install go gtk3 webkit2gtk
```

**Build:**
```bash
go build -o bagus-browser ./cmd/bagus-browser
```

#### Criar App Bundle

```bash
#!/bin/bash

# Criar estrutura do app
mkdir -p Bagus.app/Contents/{MacOS,Resources}

# Copiar binário
cp bagus-browser Bagus.app/Contents/MacOS/

# Criar Info.plist
cat > Bagus.app/Contents/Info.plist << 'EOF'
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>CFBundleDevelopmentRegion</key>
    <string>pt_BR</string>
    <key>CFBundleExecutable</key>
    <string>bagus-browser</string>
    <key>CFBundleIdentifier</key>
    <string>com.bagus.browser</string>
    <key>CFBundleInfoDictionaryVersion</key>
    <string>6.0</string>
    <key>CFBundleName</key>
    <string>Bagus Browser</string>
    <key>CFBundlePackageType</key>
    <string>APPL</string>
    <key>CFBundleShortVersionString</key>
    <string>5.0.0</string>
    <key>CFBundleVersion</key>
    <string>1</string>
</dict>
</plist>
EOF

# Assinatura (opcional)
codesign -s - Bagus.app

# Criar DMG
hdiutil create -volname "Bagus Browser" -srcfolder . -ov -format UDZO bagus-browser.dmg
```

---

### 4. FreeBSD (✅ Possível)

#### Requisitos
```bash
pkg install go gtk3 webkit2-gtk3
```

#### Build
```bash
./scripts/build-multiplatform.sh --os freebsd --arch amd64
```

#### Resultado
- Binário: `dist/bagus-browser` (5,5M)
- Tarball: `dist/bagus-browser_v5.0.0_freebsd_amd64.tar.gz`

---

## 🔧 Configuração por Plataforma

### Caminhos de Configuração

O Bagus Browser detecta automaticamente o SO e usa os caminhos corretos:

**Linux:**
```
~/.config/bagus-browser/
~/.cache/bagus-browser/
~/Downloads/
```

**Windows:**
```
%APPDATA%\bagus-browser\
%LOCALAPPDATA%\bagus-browser\
%USERPROFILE%\Downloads\
```

**macOS:**
```
~/Library/Application Support/bagus-browser/
~/Library/Caches/bagus-browser/
~/Downloads/
```

**FreeBSD:**
```
~/.config/bagus-browser/
~/.cache/bagus-browser/
~/Downloads/
```

---

## 📊 Matriz de Compatibilidade

| Plataforma | Arquitetura | Build Nativo | Compilação Cruzada | Status |
|-----------|------------|--------------|-------------------|--------|
| Linux     | x86_64     | ✅           | ✅                | ✅ Pronto |
| Linux     | ARM64      | ✅           | ✅                | ✅ Pronto |
| Windows   | x86_64     | ✅           | ⚠️ Requer MSYS2   | ⚠️ Possível |
| Windows   | x86        | ✅           | ⚠️ Requer MSYS2   | ⚠️ Possível |
| macOS     | x86_64     | ✅           | ⚠️ Requer osxcross | ⚠️ Possível |
| macOS     | ARM64      | ✅           | ⚠️ Requer osxcross | ⚠️ Possível |
| FreeBSD   | x86_64     | ✅           | ✅                | ✅ Possível |

---

## 🚀 Distribuição

### Linux

**Debian/Ubuntu:**
```bash
# Usar o .deb existente
./scripts/bagus install
```

**Genérico:**
```bash
tar -xzf bagus-browser_v5.0.0_linux_amd64.tar.gz
sudo cp bagus-browser /usr/local/bin/
```

### Windows

**Instalador:**
- Usar NSIS para criar `.exe` instalador
- Incluir GTK3 e WebKit2GTK

**Portável:**
- Distribuir `.zip` com binário + DLLs

### macOS

**App Bundle:**
- Distribuir `.dmg` com app bundle
- Assinatura de código (opcional)

**Homebrew:**
```bash
brew install bagus-browser
```

### FreeBSD

**Package:**
```bash
pkg install bagus-browser
```

**Port:**
- Adicionar ao FreeBSD ports

---

## 🔍 Troubleshooting

### Erro: "pkg-config not found"

**Solução:**
```bash
# Linux
sudo apt-get install pkg-config

# macOS
brew install pkg-config

# FreeBSD
pkg install pkgconf
```

### Erro: "GTK3 not found"

**Solução:**
```bash
# Linux
sudo apt-get install libgtk-3-dev

# macOS
brew install gtk3

# FreeBSD
pkg install gtk3
```

### Erro: "WebKit2GTK not found"

**Solução:**
```bash
# Linux
sudo apt-get install libwebkit2gtk-4.0-dev

# macOS
brew install webkit2gtk

# FreeBSD
pkg install webkit2-gtk3
```

---

## 📝 Próximos Passos

1. ✅ Testar build multiplataforma
2. ⚠️ Criar instaladores para cada plataforma
3. ⚠️ Configurar CI/CD (GitHub Actions)
4. ⚠️ Publicar em repositórios de pacotes
5. ⚠️ Documentar processo de instalação

---

## 📞 Suporte

Para problemas de build em plataformas específicas:

1. Verificar requisitos de dependências
2. Consultar documentação do Go para compilação cruzada
3. Verificar logs de compilação
4. Abrir issue no GitHub

---

**Última atualização:** 02/11/2025
**Versão:** 5.0.0
