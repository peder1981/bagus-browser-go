# 📚 Versões Antigas do Bagus Browser

## ℹ️ Informação

A partir da **v5.0.0**, o Bagus Browser foi completamente reescrito em **CGo puro**, eliminando todas as dependências Go problemáticas (gotk3).

As versões antigas foram preservadas para referência histórica e compatibilidade.

---

## 🗂️ Branches de Versões Antigas

### v4.x - Branch: `old-versions/v4`

**Última versão:** v4.6.3  
**Tecnologia:** Go + gotk3 (binding Go para GTK3)  
**Status:** Mantida para referência, não recomendada para uso

**Acesso:**
```bash
git checkout old-versions/v4
```

**Características:**
- ✅ Interface GTK3 via gotk3
- ✅ WebKitGTK 4.0
- ✅ Sistema de favoritos
- ✅ Configurações avançadas
- ⚠️ Dependência problemática: gotk3
- ⚠️ Crashes ocasionais
- ⚠️ Binário maior (~4 MB)

**Por que foi substituída:**
- Dependência gotk3 causava crashes frequentes
- Problemas de estabilidade
- Dificuldade de manutenção
- Binário maior

---

## 🚀 Versão Atual Recomendada

### v5.0.0+ - Branch: `main`

**Tecnologia:** Go + CGo Puro (C nativo)  
**Status:** Versão oficial e recomendada

**Vantagens sobre v4:**
- ✅ **100% CGo Puro** - Zero dependências Go problemáticas
- ✅ **Estabilidade Máxima** - Sem crashes
- ✅ **Binário Menor** - 2.3 MB (vs 4 MB da v4)
- ✅ **Performance Melhor** - Código C nativo
- ✅ **Manutenção Simples** - Menos dependências
- ✅ **Migração Automática** - Importa dados da v4

**Instalação:**
```bash
# Via .deb
wget https://github.com/peder1981/bagus-browser-go/releases/download/v5.0.0/bagus-browser_v5.0.0_amd64.deb
sudo dpkg -i bagus-browser_v5.0.0_amd64.deb

# Via tarball
wget https://github.com/peder1981/bagus-browser-go/releases/download/v5.0.0/bagus-browser_v5.0.0_linux_amd64.tar.gz
tar -xzf bagus-browser_v5.0.0_linux_amd64.tar.gz
sudo mv bagus-browser /usr/local/bin/
```

---

## 🔄 Migração da v4 para v5

A migração é **100% automática**!

Todos os seus dados da v4 serão importados:
- ✅ Favoritos (`~/.config/bagus-browser/bookmarks.enc`)
- ✅ Configurações (`~/.config/bagus-browser/config.enc`)
- ✅ Sessões (`~/.config/bagus-browser/session.enc`)
- ✅ Cookies (`~/.config/bagus-browser/cookies.sqlite`)

**Passos:**
1. Instale a v5.0.0
2. Execute `bagus-browser`
3. Pronto! Seus dados foram importados automaticamente

---

## 📊 Comparação de Versões

| Característica | v4.6.3 | v5.0.0 |
|----------------|--------|--------|
| Tecnologia | gotk3 | CGo Puro |
| Tamanho Binário | ~4 MB | 2.3 MB |
| Estabilidade | ⚠️ Média | ✅ Máxima |
| Crashes | Ocasionais | Zero |
| Dependências Go | gotk3 | Nenhuma |
| Performance | Boa | Excelente |
| Manutenção | Difícil | Simples |
| Migração de Dados | - | Automática |

---

## 🔗 Links Úteis

- **Versão Atual (v5):** [README.md](README.md)
- **Release Notes v5.0.0:** [docs/releases/v5.0.0.md](docs/releases/v5.0.0.md)
- **CHANGELOG:** [CHANGELOG.md](CHANGELOG.md)
- **Releases:** https://github.com/peder1981/bagus-browser-go/releases

---

## ❓ FAQ

### Posso ainda usar a v4?

Sim, mas **não é recomendado**. A v4 tem problemas conhecidos de estabilidade e não receberá mais atualizações.

### Meus dados da v4 serão perdidos?

Não! A v5 importa automaticamente todos os dados da v4.

### Posso voltar para v4 depois de usar v5?

Sim, mas não é recomendado. Os dados são compatíveis entre versões.

### A v4 receberá correções de bugs?

Não. Todo o desenvolvimento está focado na v5.

---

**Recomendação:** Use sempre a versão mais recente (v5.0.0+) para melhor experiência e estabilidade.
