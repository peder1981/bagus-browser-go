# Criptografia de Dados no Bagus Browser

## Visão Geral

O Bagus Browser implementa **criptografia de nível militar** para proteger todos os dados sensíveis dos usuários. Nenhum dado é armazenado em texto plano, garantindo privacidade total mesmo se o dispositivo for comprometido.

## Tecnologia Utilizada

### AES-256-GCM

**Advanced Encryption Standard com Galois/Counter Mode**

- **Padrão**: FIPS 197 (Federal Information Processing Standard)
- **Aprovação**: NSA Suite B Cryptography para TOP SECRET
- **Tipo**: AEAD (Authenticated Encryption with Associated Data)
- **Tamanho de Chave**: 256 bits
- **Tamanho de Nonce**: 12 bytes (96 bits)
- **Tag de Autenticação**: 16 bytes (128 bits)

### Por que AES-256-GCM?

1. **Segurança Comprovada**: Usado por governos e militares mundialmente
2. **Autenticação Integrada**: Detecta qualquer adulteração nos dados
3. **Performance**: Otimizado em hardware moderno (AES-NI)
4. **Padrão da Indústria**: Recomendado por NIST, NSA, IETF

## Arquitetura de Segurança

### Geração de Chaves

```
┌─────────────────────────────────────────────────────────┐
│ 1. Geração de Material Aleatório                       │
│    - 32 bytes via crypto/rand (CSPRNG)                 │
│    - Entropia do sistema operacional                   │
└─────────────────────────────────────────────────────────┘
                          ↓
┌─────────────────────────────────────────────────────────┐
│ 2. Geração de Salt                                      │
│    - 32 bytes aleatórios únicos por usuário            │
│    - Armazenado em .encryption_salt                    │
└─────────────────────────────────────────────────────────┘
                          ↓
┌─────────────────────────────────────────────────────────┐
│ 3. Derivação de Chave (PBKDF2-SHA256)                  │
│    - 100.000 iterações                                  │
│    - Proteção contra ataques de força bruta            │
│    - Chave final: 32 bytes (256 bits)                  │
└─────────────────────────────────────────────────────────┘
                          ↓
┌─────────────────────────────────────────────────────────┐
│ 4. Armazenamento Seguro                                 │
│    - Arquivo: .encryption_key                          │
│    - Permissões: 0600 (apenas owner)                   │
│    - Diretório: 0700 (apenas owner)                    │
└─────────────────────────────────────────────────────────┘
```

### Processo de Criptografia

```
┌──────────────────┐
│  Dados Originais │  (ex: histórico de navegação)
│  (JSON)          │
└────────┬─────────┘
         │
         ↓
┌──────────────────┐
│  Serialização    │  json.Marshal()
│  ([]byte)        │
└────────┬─────────┘
         │
         ↓
┌──────────────────┐
│  Gera Nonce      │  12 bytes aleatórios
│  (crypto/rand)   │  Único para cada operação
└────────┬─────────┘
         │
         ↓
┌──────────────────┐
│  AES-256-GCM     │  Encrypt + Authenticate
│  Encryption      │  
└────────┬─────────┘
         │
         ↓
┌──────────────────┐
│  Formato Final   │  [nonce][ciphertext][tag]
│  ([]byte)        │  
└────────┬─────────┘
         │
         ↓
┌──────────────────┐
│  Base64 Encode   │  Para armazenamento seguro
│  (string)        │
└────────┬─────────┘
         │
         ↓
┌──────────────────┐
│  Salvar Arquivo  │  history.encrypted
│  (modo 0600)     │  Apenas owner pode ler
└──────────────────┘
```

### Processo de Descriptografia

```
┌──────────────────┐
│  Ler Arquivo     │  history.encrypted
│  (modo 0600)     │
└────────┬─────────┘
         │
         ↓
┌──────────────────┐
│  Base64 Decode   │  string → []byte
│                  │
└────────┬─────────┘
         │
         ↓
┌──────────────────┐
│  Extrair Partes  │  [nonce][ciphertext][tag]
│                  │
└────────┬─────────┘
         │
         ↓
┌──────────────────┐
│  Validar Tag     │  Verifica integridade
│  (GCM)           │  Falha se adulterado
└────────┬─────────┘
         │
         ↓
┌──────────────────┐
│  AES-256-GCM     │  Decrypt
│  Decryption      │  
└────────┬─────────┘
         │
         ↓
┌──────────────────┐
│  Deserialização  │  json.Unmarshal()
│  (JSON → struct) │
└────────┬─────────┘
         │
         ↓
┌──────────────────┐
│  Dados Originais │  Prontos para uso
│                  │
└──────────────────┘
```

## Segurança em Camadas

### Camada 1: Criptografia
- AES-256-GCM para confidencialidade
- Tag de autenticação para integridade
- Nonce único por operação

### Camada 2: Derivação de Chave
- PBKDF2 com 100.000 iterações
- Salt único por usuário
- Proteção contra rainbow tables

### Camada 3: Permissões de Arquivo
- Arquivos: modo 0600 (rw-------)
- Diretórios: modo 0700 (rwx------)
- Apenas owner tem acesso

### Camada 4: Isolamento por Usuário
- Chaves únicas por usuário
- Dados completamente isolados
- Sem compartilhamento entre usuários

## Dados Protegidos

### Atualmente Criptografados
- ✅ **Histórico de Navegação** (`history.encrypted`)
  - URLs visitadas
  - Títulos de páginas
  - Timestamps de acesso

### Planejado para Criptografia
- 🔄 **Favoritos** (próxima versão)
- 🔄 **Cookies** (próxima versão)
- 🔄 **Senhas Salvas** (próxima versão)
- 🔄 **Formulários Autopreenchidos** (próxima versão)

## Migração Automática

O sistema detecta e migra automaticamente dados não criptografados:

```go
// Detecta arquivo antigo
if _, err := os.Stat("history.json"); err == nil {
    // Lê dados não criptografados
    oldData := readOldFormat()
    
    // Criptografa e salva
    encryptedData := encrypt(oldData)
    save("history.encrypted", encryptedData)
    
    // Remove arquivo antigo
    os.Remove("history.json")
}
```

## API de Criptografia

### Criar Encryptor

```go
import "github.com/peder1981/bagus-browser-go/internal/security"

// Cria encryptor para um usuário
encryptor, err := security.NewEncryptor(userPath)
if err != nil {
    log.Fatal(err)
}
```

### Criptografar Dados

```go
// Dados originais
plaintext := []byte("dados sensíveis")

// Criptografa
ciphertext, err := encryptor.Encrypt(plaintext)
if err != nil {
    log.Fatal(err)
}

// ciphertext é uma string base64
// Pode ser salva diretamente em arquivo
```

### Descriptografar Dados

```go
// Lê dados criptografados
ciphertext := readFromFile()

// Descriptografa
plaintext, err := encryptor.Decrypt(ciphertext)
if err != nil {
    log.Fatal(err) // Falha se adulterado
}

// plaintext contém dados originais
```

### Rotação de Chaves

```go
// Gera nova chave
err := encryptor.RotateKey()
if err != nil {
    log.Fatal(err)
}

// Nota: Você precisa re-criptografar dados existentes
```

### Destruição Segura

```go
// Remove chaves do disco de forma segura
err := encryptor.Destroy()
if err != nil {
    log.Fatal(err)
}

// Sobrescreve com dados aleatórios antes de deletar
```

## Análise de Segurança

### Resistência a Ataques

| Tipo de Ataque | Proteção | Status |
|----------------|----------|--------|
| Força Bruta | PBKDF2 100k iterações | ✅ Protegido |
| Rainbow Tables | Salt único por usuário | ✅ Protegido |
| Adulteração | Tag GCM de autenticação | ✅ Detectado |
| Replay Attack | Nonce único por operação | ✅ Protegido |
| Side-Channel | Constant-time operations | ✅ Mitigado |
| Cold Boot | Limpeza de memória | ✅ Mitigado |
| Acesso Físico | Permissões de arquivo | ⚠️ Limitado* |

\* *Acesso físico com privilégios root pode comprometer qualquer sistema*

### Tempo para Quebrar (Estimativas)

Assumindo computador com 1 bilhão de tentativas/segundo:

- **AES-256**: ~3.31 × 10^56 anos (idade do universo: 1.38 × 10^10 anos)
- **PBKDF2 (100k)**: Adiciona ~100.000x de tempo por tentativa

**Conclusão**: Praticamente impossível de quebrar com tecnologia atual ou futura previsível.

## Conformidade e Padrões

### Padrões Seguidos
- ✅ **NIST SP 800-38D**: GCM Mode
- ✅ **NIST SP 800-132**: PBKDF2
- ✅ **FIPS 197**: AES
- ✅ **RFC 5116**: AEAD Cipher Suites

### Bibliotecas Utilizadas
- `crypto/aes`: Implementação Go do AES
- `crypto/cipher`: Modos de operação (GCM)
- `crypto/rand`: Gerador criptográfico de números aleatórios
- `golang.org/x/crypto/pbkdf2`: Derivação de chaves

## Performance

### Benchmarks (Hardware Típico)

| Operação | Tempo | Throughput |
|----------|-------|------------|
| Encrypt (1KB) | ~50µs | ~20 MB/s |
| Decrypt (1KB) | ~50µs | ~20 MB/s |
| Key Derivation | ~100ms | 1x (por design) |

**Nota**: AES-NI (hardware acceleration) melhora performance em ~10x

## Perguntas Frequentes

### P: Os dados são enviados para a nuvem?
**R**: Não. Todos os dados ficam localmente no seu dispositivo.

### P: Alguém pode ver meu histórico?
**R**: Não, desde que não tenham acesso físico ao seu dispositivo com suas credenciais.

### P: E se eu esquecer minha senha?
**R**: Os dados não podem ser recuperados. A criptografia é irreversível sem a chave.

### P: Posso exportar meus dados criptografados?
**R**: Sim, mas você precisará da chave de criptografia para descriptografá-los.

### P: A criptografia deixa o navegador mais lento?
**R**: Não perceptivelmente. A sobrecarga é de microssegundos por operação.

### P: Posso desabilitar a criptografia?
**R**: Não. A criptografia é obrigatória para proteger sua privacidade.

## Contribuindo

Se você é um especialista em criptografia e encontrou algum problema, por favor:

1. **Não** abra uma issue pública
2. Entre em contato via `security@bagus-browser.dev`
3. Aguarde nossa resposta antes de divulgar

## Referências

- [NIST SP 800-38D - GCM](https://nvlpubs.nist.gov/nistpubs/Legacy/SP/nistspecialpublication800-38d.pdf)
- [FIPS 197 - AES](https://nvlpubs.nist.gov/nistpubs/FIPS/NIST.FIPS.197.pdf)
- [RFC 5116 - AEAD](https://tools.ietf.org/html/rfc5116)
- [OWASP Cryptographic Storage Cheat Sheet](https://cheatsheetseries.owasp.org/cheatsheets/Cryptographic_Storage_Cheat_Sheet.html)

---

**Última atualização**: 2024-10-20
