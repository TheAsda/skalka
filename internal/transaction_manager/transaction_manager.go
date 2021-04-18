package transaction_manager

import (
	"context"
	"fmt"
	"github.com/TheAsda/skalka/internal"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"strconv"
	"sync"
	"time"
)

type TransactionManager struct {
	targetPath    string
	isTmp         bool
	tmpPath       string
	tmpTargetPath string
}

func NewTransactionManager(path string, isTmp bool) *TransactionManager {
	return &TransactionManager{targetPath: path, isTmp: isTmp, tmpPath: os.TempDir(), tmpTargetPath: ""}
}

func (m *TransactionManager) getTempDir() (string, error) {
	if m.tmpTargetPath != "" {
		return m.tmpTargetPath, nil
	}
	rand.Seed(time.Now().UnixNano())
	folderName := rand.Int()
	m.tmpTargetPath = path.Join(m.tmpPath, strconv.Itoa(folderName))
	err := os.Mkdir(m.tmpTargetPath, 0777)
	if err != nil {
		return "", err
	}
	return m.tmpTargetPath, nil
}

func (m *TransactionManager) GetPath() (string, error) {
	if m.isTmp == false {
		return m.targetPath, nil
	}
	return m.getTempDir()
}

func (m *TransactionManager) Commit() error {
	if !m.isTmp {
		return nil
	}
	// Get temp directory items
	dir, err := ioutil.ReadDir(m.tmpTargetPath)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancelChan := make(chan bool)
	errorChan := make(chan error)
	progressChan := make(chan string, 10)
	var movedItems []string
	var endWg sync.WaitGroup

	endWg.Add(1)
	// On cancel signal it will call cancel
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			select {
			case <-cancelChan:
				cancel()
				close(progressChan)
				close(cancelChan)
				return
			case item := <-progressChan:
				movedItems = append(movedItems, item)
				if len(movedItems) == len(dir) {
					cancel()
					close(progressChan)
					close(cancelChan)
					return
				}
			}
		}
	}(&endWg)

	// Move items to target folder
	for _, item := range dir {
		go func(ctx context.Context, name string) {
			select {
			case <-ctx.Done():
				return
			default:
				err := m.move(name)
				if err != nil {
					cancelChan <- true
					errorChan <- err
				}
				progressChan <- name
			}
		}(ctx, item.Name())
	}
	endWg.Wait()
	select {
	case err = <-errorChan:
		close(errorChan)
		// If there was an error return error
		if err != nil {
			removeErr := os.Remove(m.tmpTargetPath)
			return internal.NewError(fmt.Sprintf("%s, remove error: %s", err.Error(), removeErr.Error()))
		}
	default:
		close(errorChan)
	}
	err = os.Remove(m.tmpTargetPath)
	return err
}

func (m *TransactionManager) move(name string) error {
	source := path.Join(m.tmpTargetPath, name)
	target := path.Join(m.targetPath, name)
	err := os.Rename(source, target)
	return err
}

func (m *TransactionManager) Rollback() error {
	if !m.isTmp {
		return nil
	}
	err := os.RemoveAll(m.tmpTargetPath)
	return err
}
