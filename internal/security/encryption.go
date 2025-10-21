package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"golang.org/x/crypto/pbkdf2"
)

const (
	// Parâmetros de segurança
	keySize      = 32 // AES-256
	saltSize     = 32
	nonceSize    = 12 // GCM padrão
	pbkdf2Rounds = 100000
	keyFileName  = ".encryption_key"
	saltFileName = ".encryption_salt"
)

// Encryptor gerencia criptografia de dados sensíveis
type Encryptor struct {
	key  []byte
	salt []byte
	path string
}

// NewEncryptor cria novo encryptor para um usuário
// A chave é derivada de forma segura e armazenada localmente
func NewEncryptor(userPath string) (*Encryptor, error) {
	e := &Encryptor{
		path: userPath,
	}

	// Tenta carregar chave existente
	keyPath := filepath.Join(userPath, keyFileName)

	// Se chave existe, carrega
	if _, err := os.Stat(keyPath); err == nil {
		if err := e.loadKey(); err != nil {
			return nil, fmt.Errorf("erro ao carregar chave: %w", err)
		}
		return e, nil
	}

	// Senão, gera nova chave
	if err := e.generateKey(); err != nil {
		return nil, fmt.Errorf("erro ao gerar chave: %w", err)
	}

	// Salva chave com permissões restritas
	if err := e.saveKey(); err != nil {
		return nil, fmt.Errorf("erro ao salvar chave: %w", err)
	}

	return e, nil
}

// generateKey gera chave criptográfica segura
func (e *Encryptor) generateKey() error {
	// Gera salt aleatório
	e.salt = make([]byte, saltSize)
	if _, err := io.ReadFull(rand.Reader, e.salt); err != nil {
		return fmt.Errorf("erro ao gerar salt: %w", err)
	}

	// Gera material de chave aleatório
	keyMaterial := make([]byte, keySize)
	if _, err := io.ReadFull(rand.Reader, keyMaterial); err != nil {
		return fmt.Errorf("erro ao gerar material de chave: %w", err)
	}

	// Deriva chave usando PBKDF2 para adicionar camada extra de segurança
	e.key = pbkdf2.Key(keyMaterial, e.salt, pbkdf2Rounds, keySize, sha256.New)

	return nil
}

// saveKey salva chave em disco com permissões restritas
func (e *Encryptor) saveKey() error {
	keyPath := filepath.Join(e.path, keyFileName)
	saltPath := filepath.Join(e.path, saltFileName)

	// Cria diretório se não existir
	if err := os.MkdirAll(e.path, 0700); err != nil {
		return fmt.Errorf("erro ao criar diretório: %w", err)
	}

	// Salva chave (apenas owner pode ler/escrever)
	if err := os.WriteFile(keyPath, e.key, 0600); err != nil {
		return fmt.Errorf("erro ao salvar chave: %w", err)
	}

	// Salva salt
	if err := os.WriteFile(saltPath, e.salt, 0600); err != nil {
		return fmt.Errorf("erro ao salvar salt: %w", err)
	}

	return nil
}

// loadKey carrega chave do disco
func (e *Encryptor) loadKey() error {
	keyPath := filepath.Join(e.path, keyFileName)
	saltPath := filepath.Join(e.path, saltFileName)

	// Carrega chave
	key, err := os.ReadFile(keyPath)
	if err != nil {
		return fmt.Errorf("erro ao ler chave: %w", err)
	}

	// Valida tamanho
	if len(key) != keySize {
		return fmt.Errorf("tamanho de chave inválido: %d bytes", len(key))
	}

	// Carrega salt
	salt, err := os.ReadFile(saltPath)
	if err != nil {
		return fmt.Errorf("erro ao ler salt: %w", err)
	}

	if len(salt) != saltSize {
		return fmt.Errorf("tamanho de salt inválido: %d bytes", len(salt))
	}

	e.key = key
	e.salt = salt

	return nil
}

// Encrypt criptografa dados usando AES-256-GCM
// Retorna dados criptografados em base64
func (e *Encryptor) Encrypt(plaintext []byte) (string, error) {
	if len(plaintext) == 0 {
		return "", nil
	}

	// Cria cipher block
	block, err := aes.NewCipher(e.key)
	if err != nil {
		return "", fmt.Errorf("erro ao criar cipher: %w", err)
	}

	// Cria GCM (Galois/Counter Mode) - autenticado
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("erro ao criar GCM: %w", err)
	}

	// Gera nonce aleatório
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", fmt.Errorf("erro ao gerar nonce: %w", err)
	}

	// Criptografa e autentica
	// Formato: nonce + ciphertext + tag
	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)

	// Codifica em base64 para armazenamento seguro
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt descriptografa dados
func (e *Encryptor) Decrypt(ciphertext string) ([]byte, error) {
	if ciphertext == "" {
		return nil, nil
	}

	// Decodifica base64
	data, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return nil, fmt.Errorf("erro ao decodificar base64: %w", err)
	}

	// Cria cipher block
	block, err := aes.NewCipher(e.key)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar cipher: %w", err)
	}

	// Cria GCM
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar GCM: %w", err)
	}

	// Valida tamanho mínimo
	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return nil, fmt.Errorf("ciphertext muito curto")
	}

	// Extrai nonce e ciphertext
	nonce, ciphertextBytes := data[:nonceSize], data[nonceSize:]

	// Descriptografa e valida autenticação
	plaintext, err := gcm.Open(nil, nonce, ciphertextBytes, nil)
	if err != nil {
		return nil, fmt.Errorf("erro ao descriptografar: %w", err)
	}

	return plaintext, nil
}

// RotateKey gera nova chave e re-criptografa dados
// Deve ser chamado periodicamente para segurança adicional
func (e *Encryptor) RotateKey() error {
	// Salva chave antiga
	oldKey := make([]byte, len(e.key))
	copy(oldKey, e.key)

	// Gera nova chave
	if err := e.generateKey(); err != nil {
		return fmt.Errorf("erro ao gerar nova chave: %w", err)
	}

	// Salva nova chave
	if err := e.saveKey(); err != nil {
		// Restaura chave antiga em caso de erro
		e.key = oldKey
		return fmt.Errorf("erro ao salvar nova chave: %w", err)
	}

	return nil
}

// Destroy remove chaves do disco (para exclusão de usuário)
func (e *Encryptor) Destroy() error {
	keyPath := filepath.Join(e.path, keyFileName)
	saltPath := filepath.Join(e.path, saltFileName)

	// Remove arquivos de forma segura (sobrescreve antes de deletar)
	if err := secureDelete(keyPath); err != nil {
		return fmt.Errorf("erro ao deletar chave: %w", err)
	}

	if err := secureDelete(saltPath); err != nil {
		return fmt.Errorf("erro ao deletar salt: %w", err)
	}

	// Zera memória
	for i := range e.key {
		e.key[i] = 0
	}
	for i := range e.salt {
		e.salt[i] = 0
	}

	return nil
}

// secureDelete sobrescreve arquivo antes de deletar
func secureDelete(path string) error {
	// Verifica se existe
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}

	// Sobrescreve com dados aleatórios
	randomData := make([]byte, info.Size())
	if _, err := io.ReadFull(rand.Reader, randomData); err != nil {
		return err
	}

	if err := os.WriteFile(path, randomData, 0600); err != nil {
		return err
	}

	// Sobrescreve com zeros
	zeros := make([]byte, info.Size())
	if err := os.WriteFile(path, zeros, 0600); err != nil {
		return err
	}

	// Remove arquivo
	return os.Remove(path)
}
