package lockfile

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"syscall"
)

// LockFile representa um arquivo de lock para garantir instância única
type LockFile struct {
	path string
	file *os.File
}

// New cria um novo lock file
func New(name string) *LockFile {
	lockPath := filepath.Join(os.TempDir(), fmt.Sprintf("%s.lock", name))
	return &LockFile{
		path: lockPath,
	}
}

// TryLock tenta adquirir o lock
// Retorna erro se já houver outra instância rodando
func (l *LockFile) TryLock() error {
	// Abre ou cria o arquivo de lock
	file, err := os.OpenFile(l.path, os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		return fmt.Errorf("erro ao criar lock file: %w", err)
	}

	// Tenta adquirir lock exclusivo (não-bloqueante)
	err = syscall.Flock(int(file.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
	if err != nil {
		file.Close()

		// Lê PID da instância em execução
		existingPID := l.readPID()
		if existingPID > 0 {
			return fmt.Errorf("outra instância do Bagus Browser já está em execução (PID: %d)", existingPID)
		}
		return fmt.Errorf("outra instância do Bagus Browser já está em execução")
	}

	// Escreve PID atual no arquivo
	pid := os.Getpid()
	file.Truncate(0)
	file.Seek(0, 0)
	fmt.Fprintf(file, "%d\n", pid)
	file.Sync()

	l.file = file
	return nil
}

// Unlock libera o lock
func (l *LockFile) Unlock() error {
	if l.file == nil {
		return nil
	}

	// Libera o lock
	syscall.Flock(int(l.file.Fd()), syscall.LOCK_UN)

	// Fecha o arquivo
	err := l.file.Close()
	l.file = nil

	// Remove o arquivo de lock
	os.Remove(l.path)

	return err
}

// readPID lê o PID do arquivo de lock
func (l *LockFile) readPID() int {
	data, err := os.ReadFile(l.path)
	if err != nil {
		return 0
	}

	pid, err := strconv.Atoi(string(data[:len(data)-1])) // Remove \n
	if err != nil {
		return 0
	}

	return pid
}
