package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"io"
	"os"
	
	"golang.org/x/crypto/pbkdf2"
)

// CryptoManager gerencia criptografia de dados locais
type CryptoManager struct {
	masterKey []byte
}

// NewCryptoManager cria um novo gerenciador de criptografia
// Se masterPassword for vazio, usa uma chave derivada do hostname
func NewCryptoManager(masterPassword string) (*CryptoManager, error) {
	var key []byte
	
	if masterPassword == "" {
		// Usar chave derivada do hostname + username para persistência
		// Isso garante que a mesma chave seja usada em todas as sessões
		hostname, _ := os.Hostname()
		username := os.Getenv("USER")
		if username == "" {
			username = os.Getenv("USERNAME") // Windows
		}
		
		// Criar senha mestre baseada no sistema
		systemKey := hostname + "-" + username + "-bagus-browser-v1"
		
		// Derivar chave usando PBKDF2
		salt := []byte("bagus-browser-salt-v1")
		key = pbkdf2.Key([]byte(systemKey), salt, 100000, 32, sha256.New)
	} else {
		// Derivar chave da senha mestre usando PBKDF2
		salt := []byte("bagus-browser-salt-v1")
		key = pbkdf2.Key([]byte(masterPassword), salt, 100000, 32, sha256.New)
	}
	
	return &CryptoManager{
		masterKey: key,
	}, nil
}

// Encrypt criptografa dados usando AES-256-GCM
func (cm *CryptoManager) Encrypt(plaintext []byte) (string, error) {
	block, err := aes.NewCipher(cm.masterKey)
	if err != nil {
		return "", err
	}
	
	// GCM mode (Galois/Counter Mode) - autenticado
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	
	// Nonce único para cada criptografia
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	
	// Criptografar e autenticar
	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
	
	// Retornar como base64 para fácil armazenamento
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt descriptografa dados
func (cm *CryptoManager) Decrypt(ciphertextB64 string) ([]byte, error) {
	// Decodificar base64
	ciphertext, err := base64.StdEncoding.DecodeString(ciphertextB64)
	if err != nil {
		return nil, err
	}
	
	block, err := aes.NewCipher(cm.masterKey)
	if err != nil {
		return nil, err
	}
	
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext muito curto")
	}
	
	// Extrair nonce e ciphertext
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	
	// Descriptografar e verificar autenticidade
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, errors.New("falha na descriptografia ou dados corrompidos")
	}
	
	return plaintext, nil
}

// GenerateSalt gera um salt aleatório
func GenerateSalt() (string, error) {
	salt := make([]byte, 32)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(salt), nil
}

// HashPassword cria hash de senha usando SHA-256
func HashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return base64.StdEncoding.EncodeToString(hash[:])
}
