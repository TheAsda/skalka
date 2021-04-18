package fs

import (
	"io/ioutil"
)

type FileReader struct {
}

func NewFileReader() *FileReader {
	return &FileReader{}
}

func (f FileReader) Read(path string) ([]byte, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}
