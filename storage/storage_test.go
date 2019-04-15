package storage_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/mpetavy/alexandria/storage"
	"github.com/mpetavy/common"

	_ "github.com/mpetavy/alexandria/storage"
)

func TestStorage(t *testing.T) {
	tempDir, err := common.TempDir()
	if err != nil {
		t.Error(err.Error())
	}

	path := filepath.Join(tempDir, "test")

	err = os.RemoveAll(path)
	if err != nil {
		t.Error(err.Error())
	}

	s, err := storage.NewStorage(path)
	if err != nil {
		t.Error(err.Error())
	}

	for i := 1; i < 10; i++ {
		docID, err := s.NewDocument()
		if err != nil {
			t.Error(err.Error())
		}

		t.Log(fmt.Sprintf("created document %v", docID))
	}
}
