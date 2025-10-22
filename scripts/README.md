# Scripts - Bagus Browser

Scripts de build, release e automação.

## 📜 Scripts Disponíveis

### build.sh
Compila o projeto e cria pacotes de instalação.

**Uso:**
```bash
./scripts/build.sh
```

**Cria:**
- `build/bagus-browser` - Binário
- `dist/bagus-browser_v4.1.0_amd64.deb` - Pacote Debian
- `dist/bagus-browser_v4.1.0_linux_amd64.tar.gz` - Tarball
- `dist/SHA256SUMS` - Checksums

---

### release.sh
Prepara arquivos e instruções para publicação manual no GitHub.

**Uso:**
```bash
./scripts/release.sh
```

**Funcionalidades:**
- Verifica tag e arquivos dist/
- Gera arquivo RELEASE_NOTES_vX.X.X.md
- Lista arquivos para upload
- Mostra instruções passo a passo para publicação manual

**Nota:** Este script NÃO publica automaticamente. Ele prepara tudo e fornece instruções para você publicar manualmente via interface web do GitHub.

---

### publish.sh
Build + Commit + Preparação de Release em um comando.

**Uso:**
```bash
./scripts/publish.sh
```

**Faz:**
1. Build completo (compila e empacota)
2. Commit e push para GitHub
3. Prepara instruções para release manual

**Após executar:** Siga as instruções exibidas para publicar a release manualmente no GitHub.

---

### install-desktop-icon.sh
Instala ícone e atalho do desktop.

**Uso:**
```bash
./scripts/install-desktop-icon.sh
```

**Funcionalidades:**
- Copia ícones para /usr/share/icons
- Cria arquivo .desktop
- Atualiza caches do sistema

---

## 🚀 Uso Rápido

### Workflow Completo
```bash
# 1. Build e empacotamento
bash scripts/build.sh

# 2. Commit e push
git add -A
git commit -m "Sua mensagem"
git push origin main

# 3. Criar tag
git tag -a v4.2.0 -m "Release v4.2.0"
git push origin v4.2.0

# 4. Preparar release
bash scripts/release.sh

# 5. Publicar manualmente
# Siga as instruções exibidas pelo script
# Acesse: https://github.com/peder1981/bagus-browser-go/releases/new
```

### Ou use o script completo
```bash
bash scripts/publish.sh
# Depois siga as instruções para publicação manual
```

---

**Veja a documentação em `/docs` para mais detalhes.**
