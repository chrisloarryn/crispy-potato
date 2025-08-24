package storage

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ccontreras/crispy-potato/internal/core/ports"
)

// LocalFileStorage implements the FileStorage interface for local file system
type LocalFileStorage struct {
	basePath string
}

// NewLocalFileStorage creates a new LocalFileStorage instance
func NewLocalFileStorage(basePath string) ports.FileStorage {
	return &LocalFileStorage{
		basePath: basePath,
	}
}

// Save saves data to a file
func (s *LocalFileStorage) Save(path string, data []byte) error {
	fullPath := filepath.Join(s.basePath, path)

	// Create directory if it doesn't exist
	dir := filepath.Dir(fullPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// Write file
	return ioutil.WriteFile(fullPath, data, 0644)
}

// Get retrieves data from a file
func (s *LocalFileStorage) Get(path string) ([]byte, error) {
	fullPath := filepath.Join(s.basePath, path)
	return ioutil.ReadFile(fullPath)
}

// Delete removes a file
func (s *LocalFileStorage) Delete(path string) error {
	fullPath := filepath.Join(s.basePath, path)
	return os.Remove(fullPath)
}
