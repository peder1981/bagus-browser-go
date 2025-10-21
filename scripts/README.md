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
Publica release no GitHub com OAuth2.

**Uso:**
```bash
./scripts/release.sh
```

**Funcionalidades:**
- OAuth2 com cache
- Upload automático de pacotes
- Notas de release automáticas

---

### publish.sh
Build + Release em um comando.

**Uso:**
```bash
./scripts/publish.sh
```

**Faz:**
1. Build completo
2. Commit e push
3. Release no GitHub

---

### github-auth.sh
Gerenciador de autenticação OAuth2 para GitHub.

**Uso:**
```bash
# Obter token
./scripts/github-auth.sh get

# Ver status
./scripts/github-auth.sh status

# Limpar cache
./scripts/github-auth.sh clear
```

---

## 🚀 Uso Rápido

### Via Makefile (Recomendado)
```bash
make build      # Compilar
make release    # Publicar
make publish    # Tudo de uma vez
```

### Direto
```bash
./scripts/build.sh
./scripts/release.sh
```

---

**Veja a documentação em `/docs` para mais detalhes.**
