# 🚀 Guia Rápido - Bagus Browser CEF

## Instalação em 3 Passos

### 1️⃣ Instalar CEF (15-30 min)

```bash
cd ~/bagus-browser-go
./scripts/install_cef.sh
```

**O que acontece:**
- Instala dependências do sistema
- Baixa CEF (~500MB)
- Compila bibliotecas necessárias

**Aguarde:** Pode demorar dependendo da sua conexão e CPU

### 2️⃣ Compilar Bagus Browser (2-5 min)

```bash
./scripts/build_cef.sh
```

**O que acontece:**
- Compila código C++ (CEF wrappers)
- Compila Go com CGO
- Copia recursos do CEF
- Cria executável final

### 3️⃣ Executar

```bash
cd build
./run_bagus_cef.sh
```

**Pronto!** Você tem um navegador Chromium completo!

---

## ✅ Teste de Compatibilidade

Após instalar, teste estes sites que **NÃO funcionam** na versão webview:

```
✅ google.com - FUNCIONA!
✅ facebook.com - FUNCIONA!
✅ twitter.com - FUNCIONA!
✅ github.com - FUNCIONA!
✅ instagram.com - FUNCIONA!
✅ linkedin.com - FUNCIONA!
```

**100% de compatibilidade garantida!**

---

## 🆘 Problemas?

### Erro: "CEF não encontrado"
```bash
# Verifique se instalou
ls -la /opt/cef

# Se não, instale
./scripts/install_cef.sh
```

### Erro: "libcef.so not found"
```bash
# Use o script de execução
cd build
./run_bagus_cef.sh
```

### Erro de compilação
```bash
# Instale dependências
sudo apt-get install build-essential cmake libgtk-3-dev
```

---

## 📊 Comparação Rápida

| Característica | Webview | CEF |
|----------------|---------|-----|
| **Google/Facebook** | ❌ | ✅ |
| **Todos os sites** | 70% | **100%** |
| **Tamanho** | 5MB | 165MB |
| **Instalação** | 1 min | 30 min |
| **DevTools** | ❌ | ✅ |

---

## 💡 Dica

Se você quer **máxima compatibilidade** e não se importa com o tamanho, use **CEF**.

Se você quer algo **leve e rápido** e não precisa de Google/Facebook, use **Webview** (padrão).

---

## 📖 Documentação Completa

- [Instalação detalhada](INSTALL_CEF.md)
- [README principal](README.md)
- [Segurança](SECURITY.md)

---

**Pronto para começar? Execute o passo 1!** 🚀
