package transaction_manager

import (
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func TestTransactionManager_GetPath(t *testing.T) {
	t.Run("get path tmp", func(t *testing.T) {
		tempPath := t.TempDir()
		manager := NewTransactionManager(tempPath, true)
		p, err := manager.GetPath()
		if err != nil {
			t.Errorf("Cannot get path: %s", err.Error())
			return
		}
		if p == tempPath {
			t.Errorf("Get tempPath must return temp tempPath if isTmp is true")
		}
	})

	t.Run("get path", func(t *testing.T) {
		tempPath := t.TempDir()
		manager := NewTransactionManager(tempPath, false)
		p, err := manager.GetPath()
		if err != nil {
			t.Errorf("Cannot get path: %s", err.Error())
			return
		}
		if p != tempPath {
			t.Errorf("Get path must return path if isTmp is false")
		}
	})
}

func TestTransactionManager_Commit(t *testing.T) {
	t.Run("get commit", func(t *testing.T) {
		tempPath := t.TempDir()
		manager := NewTransactionManager(tempPath, true)
		p, err := manager.GetPath()
		if err != nil {
			t.Errorf("Cannot get path: %s", err.Error())
			return
		}
		file, err := os.Create(path.Join(p, "file.txt"))
		if err != nil {
			t.Errorf("Cannot create file: %s", err.Error())
			return
		}
		err = file.Close()
		if err != nil {
			t.Errorf("Cannot close file: %s", err.Error())
			return
		}
		err = manager.Commit()
		if err != nil {
			t.Errorf("Error while commiting: %s", err.Error())
			return
		}
		if _, err := os.Stat(p); !os.IsNotExist(err) {
			t.Errorf("Temp folder did not delete")
		}
		dir, err := ioutil.ReadDir(tempPath)
		if err != nil {
			t.Errorf("Cannot read dir: %s", err.Error())
		}
		if len(dir) != 1 || dir[0].Name() != "file.txt" {
			t.Errorf("Target directory does not contain created file")
		}
	})
}

func TestTransactionManager_Rollback(t *testing.T) {
	t.Run("get rollback", func(t *testing.T) {
		tempPath := t.TempDir()
		manager := NewTransactionManager(tempPath, true)
		p, err := manager.GetPath()
		if err != nil {
			t.Errorf("Cannot get path: %s", err.Error())
			return
		}
		file, err := os.Create(path.Join(p, "file.txt"))
		if err != nil {
			t.Errorf("Cannot create file: %s", err.Error())
			return
		}
		err = file.Close()
		if err != nil {
			t.Errorf("Cannot close file: %s", err.Error())
			return
		}
		err = manager.Rollback()
		if err != nil {
			t.Errorf("Error while rolling back: %s", err.Error())
			return
		}
		if _, err := os.Stat(p); !os.IsNotExist(err) {
			t.Errorf("Temp folder did not delete")
		}
		dir, err := ioutil.ReadDir(tempPath)
		if err != nil {
			t.Errorf("Cannot read dir: %s", err.Error())
		}
		if len(dir) != 0 {
			t.Errorf("Target directory contains created file")
		}
	})
}
