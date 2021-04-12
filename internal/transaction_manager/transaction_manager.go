package transaction_manager

import (
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"strconv"
	"time"
)

type TransactionManager struct {
	path    string
	isTmp   bool
	tmpPath string
	tmpDir  string
}

func NewTransactionManager(path string, isTmp bool) *TransactionManager {
	return &TransactionManager{path: path, isTmp: isTmp, tmpPath: os.TempDir()}
}

func (m *TransactionManager) getTempDir() (string, error) {
	rand.Seed(time.Now().UnixNano())
	folderName := rand.Int()
	m.tmpDir = path.Join(m.tmpPath, strconv.Itoa(folderName))
	err := os.Mkdir(m.tmpDir, 0777)
	if err != nil {
		return "", err
	}
	return m.tmpDir, nil
}

func (m *TransactionManager) GetPath() (string, error) {
	if m.isTmp == false {
		return m.path, nil
	}
	return m.getTempDir()
}

func (m *TransactionManager) Commit() error {
	if !m.isTmp {
		return nil
	}
	// Get temp directory items
	dir, err := ioutil.ReadDir(m.tmpDir)
	if err != nil {
		return err
	}
	// Move items to target folder
	for _, item := range dir {
		err := m.move(item.Name())
		if err != nil {
			return err
		}
	}
	// Remove temp directory
	err = os.Remove(m.tmpDir)
	return err
}

func (m *TransactionManager) move(name string) error {
	source := path.Join(m.tmpDir, name)
	target := path.Join(m.path, name)
	err := os.Rename(source, target)
	return err
}

func (m *TransactionManager) Rollback() error {
	if !m.isTmp {
		return nil
	}
	err := os.RemoveAll(m.tmpDir)
	return err
}
