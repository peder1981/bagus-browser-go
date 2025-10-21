package security

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	// Cria diretório temporário
	tmpDir := t.TempDir()

	// Cria encryptor
	enc, err := NewEncryptor(tmpDir)
	if err != nil {
		t.Fatalf("Erro ao criar encryptor: %v", err)
	}

	// Testa dados vazios
	t.Run("Empty data", func(t *testing.T) {
		ciphertext, err := enc.Encrypt(nil)
		if err != nil {
			t.Fatalf("Erro ao criptografar dados vazios: %v", err)
		}
		if ciphertext != "" {
			t.Errorf("Esperado string vazia, obteve: %s", ciphertext)
		}

		plaintext, err := enc.Decrypt("")
		if err != nil {
			t.Fatalf("Erro ao descriptografar string vazia: %v", err)
		}
		if plaintext != nil {
			t.Errorf("Esperado nil, obteve: %v", plaintext)
		}
	})

	// Testa dados pequenos
	t.Run("Small data", func(t *testing.T) {
		original := []byte("Hello, World!")

		ciphertext, err := enc.Encrypt(original)
		if err != nil {
			t.Fatalf("Erro ao criptografar: %v", err)
		}

		if ciphertext == "" {
			t.Fatal("Ciphertext vazio")
		}

		plaintext, err := enc.Decrypt(ciphertext)
		if err != nil {
			t.Fatalf("Erro ao descriptografar: %v", err)
		}

		if !bytes.Equal(original, plaintext) {
			t.Errorf("Dados não correspondem.\nOriginal: %s\nDescriptografado: %s", original, plaintext)
		}
	})

	// Testa dados grandes
	t.Run("Large data", func(t *testing.T) {
		// 1MB de dados
		original := make([]byte, 1024*1024)
		for i := range original {
			original[i] = byte(i % 256)
		}

		ciphertext, err := enc.Encrypt(original)
		if err != nil {
			t.Fatalf("Erro ao criptografar dados grandes: %v", err)
		}

		plaintext, err := enc.Decrypt(ciphertext)
		if err != nil {
			t.Fatalf("Erro ao descriptografar dados grandes: %v", err)
		}

		if !bytes.Equal(original, plaintext) {
			t.Error("Dados grandes não correspondem")
		}
	})

	// Testa dados com caracteres especiais
	t.Run("Special characters", func(t *testing.T) {
		original := []byte("Olá! 你好 🔒 Тест")

		ciphertext, err := enc.Encrypt(original)
		if err != nil {
			t.Fatalf("Erro ao criptografar caracteres especiais: %v", err)
		}

		plaintext, err := enc.Decrypt(ciphertext)
		if err != nil {
			t.Fatalf("Erro ao descriptografar caracteres especiais: %v", err)
		}

		if !bytes.Equal(original, plaintext) {
			t.Errorf("Caracteres especiais não correspondem.\nOriginal: %s\nDescriptografado: %s", original, plaintext)
		}
	})
}

func TestEncryptorPersistence(t *testing.T) {
	tmpDir := t.TempDir()

	// Cria primeiro encryptor
	enc1, err := NewEncryptor(tmpDir)
	if err != nil {
		t.Fatalf("Erro ao criar primeiro encryptor: %v", err)
	}

	// Criptografa dados
	original := []byte("Test data for persistence")
	ciphertext, err := enc1.Encrypt(original)
	if err != nil {
		t.Fatalf("Erro ao criptografar: %v", err)
	}

	// Cria segundo encryptor (deve carregar mesma chave)
	enc2, err := NewEncryptor(tmpDir)
	if err != nil {
		t.Fatalf("Erro ao criar segundo encryptor: %v", err)
	}

	// Descriptografa com segundo encryptor
	plaintext, err := enc2.Decrypt(ciphertext)
	if err != nil {
		t.Fatalf("Erro ao descriptografar com segundo encryptor: %v", err)
	}

	if !bytes.Equal(original, plaintext) {
		t.Error("Dados não correspondem após recarregar encryptor")
	}
}

func TestInvalidCiphertext(t *testing.T) {
	tmpDir := t.TempDir()

	enc, err := NewEncryptor(tmpDir)
	if err != nil {
		t.Fatalf("Erro ao criar encryptor: %v", err)
	}

	tests := []struct {
		name       string
		ciphertext string
	}{
		{"Invalid base64", "not-valid-base64!!!"},
		{"Too short", "YWJj"},               // "abc" em base64 - muito curto
		{"Random data", "SGVsbG8gV29ybGQh"}, // "Hello World!" - não criptografado
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := enc.Decrypt(tt.ciphertext)
			if err == nil {
				t.Error("Esperado erro ao descriptografar dados inválidos")
			}
		})
	}
}

func TestTamperedData(t *testing.T) {
	tmpDir := t.TempDir()

	enc, err := NewEncryptor(tmpDir)
	if err != nil {
		t.Fatalf("Erro ao criar encryptor: %v", err)
	}

	// Criptografa dados
	original := []byte("Important data")
	ciphertext, err := enc.Encrypt(original)
	if err != nil {
		t.Fatalf("Erro ao criptografar: %v", err)
	}

	// Adultera dados (muda um caractere)
	tamperedBytes := []byte(ciphertext)
	if len(tamperedBytes) > 10 {
		tamperedBytes[10] = tamperedBytes[10] ^ 0xFF // Inverte bits
	}
	tampered := string(tamperedBytes)

	// Tenta descriptografar dados adulterados
	_, err = enc.Decrypt(tampered)
	if err == nil {
		t.Error("Esperado erro ao descriptografar dados adulterados")
	}
}

func TestKeyFiles(t *testing.T) {
	tmpDir := t.TempDir()

	// Cria encryptor
	_, err := NewEncryptor(tmpDir)
	if err != nil {
		t.Fatalf("Erro ao criar encryptor: %v", err)
	}

	// Verifica se arquivos foram criados
	keyPath := filepath.Join(tmpDir, ".encryption_key")
	saltPath := filepath.Join(tmpDir, ".encryption_salt")

	if _, err := os.Stat(keyPath); os.IsNotExist(err) {
		t.Error("Arquivo de chave não foi criado")
	}

	if _, err := os.Stat(saltPath); os.IsNotExist(err) {
		t.Error("Arquivo de salt não foi criado")
	}

	// Verifica permissões
	keyInfo, err := os.Stat(keyPath)
	if err != nil {
		t.Fatalf("Erro ao obter info do arquivo de chave: %v", err)
	}

	expectedPerm := os.FileMode(0600)
	if keyInfo.Mode().Perm() != expectedPerm {
		t.Errorf("Permissões incorretas no arquivo de chave. Esperado: %v, Obteve: %v",
			expectedPerm, keyInfo.Mode().Perm())
	}
}

func TestRotateKey(t *testing.T) {
	tmpDir := t.TempDir()

	enc, err := NewEncryptor(tmpDir)
	if err != nil {
		t.Fatalf("Erro ao criar encryptor: %v", err)
	}

	// Salva chave antiga
	oldKey := make([]byte, len(enc.key))
	copy(oldKey, enc.key)

	// Rotaciona chave
	err = enc.RotateKey()
	if err != nil {
		t.Fatalf("Erro ao rotacionar chave: %v", err)
	}

	// Verifica se chave mudou
	if bytes.Equal(oldKey, enc.key) {
		t.Error("Chave não foi alterada após rotação")
	}

	// Verifica se nova chave foi salva
	enc2, err := NewEncryptor(tmpDir)
	if err != nil {
		t.Fatalf("Erro ao criar novo encryptor: %v", err)
	}

	if !bytes.Equal(enc.key, enc2.key) {
		t.Error("Nova chave não foi persistida corretamente")
	}
}

func TestDestroy(t *testing.T) {
	tmpDir := t.TempDir()

	enc, err := NewEncryptor(tmpDir)
	if err != nil {
		t.Fatalf("Erro ao criar encryptor: %v", err)
	}

	// Destrói chaves
	err = enc.Destroy()
	if err != nil {
		t.Fatalf("Erro ao destruir chaves: %v", err)
	}

	// Verifica se arquivos foram removidos
	keyPath := filepath.Join(tmpDir, ".encryption_key")
	saltPath := filepath.Join(tmpDir, ".encryption_salt")

	if _, err := os.Stat(keyPath); !os.IsNotExist(err) {
		t.Error("Arquivo de chave não foi removido")
	}

	if _, err := os.Stat(saltPath); !os.IsNotExist(err) {
		t.Error("Arquivo de salt não foi removido")
	}

	// Verifica se memória foi zerada
	for i, b := range enc.key {
		if b != 0 {
			t.Errorf("Chave não foi zerada na posição %d: %d", i, b)
			break
		}
	}
}

func BenchmarkEncrypt(b *testing.B) {
	tmpDir := b.TempDir()
	enc, _ := NewEncryptor(tmpDir)

	data := make([]byte, 1024) // 1KB
	for i := range data {
		data[i] = byte(i % 256)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = enc.Encrypt(data)
	}
}

func BenchmarkDecrypt(b *testing.B) {
	tmpDir := b.TempDir()
	enc, _ := NewEncryptor(tmpDir)

	data := make([]byte, 1024) // 1KB
	ciphertext, _ := enc.Encrypt(data)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = enc.Decrypt(ciphertext)
	}
}
