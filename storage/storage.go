package storage

import (
	"io"
	"os"
	"path/filepath"
	"sync"
)

// DocumentID bla
type DocumentID int

// ComponentID bla
type ComponentID int

// Storage struct
type Storage struct {
	mutex          sync.Mutex
	path           string
	nextDocumentID int
}

// NewStorage initialize a new Storage
func NewStorage(path string) (storage *Storage, err error) {
	if _, err := os.Stat(path); err == nil {
		storage = &Storage{path: path}
	}

	return storage, err
}

// NewDocument creates a new document
func (storage *Storage) NewDocument() (documentID DocumentID, err error) {
	storage.mutex.Lock()
	storage.nextDocumentID = storage.nextDocumentID + 1
	documentID = DocumentID(storage.nextDocumentID)
	storage.mutex.Unlock()

	path := filepath.Join(storage.path, string(storage.nextDocumentID))

	os.MkdirAll(path, os.ModePerm)

	return documentID, err
}

// NewComponent creates a new component within a document
func (storage *Storage) NewComponent(reader io.Reader) (componentID ComponentID, err error) {
	return componentID, err
}
