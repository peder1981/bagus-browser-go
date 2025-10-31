# 📋 Requisitos do Sistema

## Versões Mínimas Requeridas

O Bagus Browser aceita **versões maiores ou iguais** aos requisitos listados abaixo:

### Linguagem
- **Go**: >= 1.21
  - Testado com: 1.21, 1.22, 1.23, 1.24+
  - Qualquer versão >= 1.21 deve funcionar

### Bibliotecas do Sistema

- **GTK3**: >= 3.0
  - Testado com: 3.24.x
  - Qualquer versão >= 3.0 deve funcionar
  
- **WebKit2GTK**: >= 2.0
  - Testado com: 2.40.x, 2.48.x
  - Qualquer versão >= 2.0 deve funcionar

### Instalação no Ubuntu/Debian

```bash
sudo apt-get install golang-go libgtk-3-dev libwebkit2gtk-4.0-dev
```

### Instalação no Fedora/RHEL

```bash
sudo dnf install golang gtk3-devel webkit2gtk4.0-devel
```

### Instalação no Arch Linux

```bash
sudo pacman -S go gtk3 webkit2gtk
```

## Verificação de Dependências

O script de build verifica automaticamente as versões:

```bash
./scripts/bagus build
```

Saída esperada:
```
✓ Go 1.24
✓ GTK3 3.24.33
✓ WebKit2GTK 2.48.7
```

## Notas

- O campo `toolchain` no `go.mod` é **opcional** e apenas sugere a versão usada no desenvolvimento
- O campo `go 1.21` define a versão **mínima** aceita
- Versões superiores são automaticamente aceitas pelo Go
