package fs

import "github.com/TheAsda/skalka/internal"

type FileReaderMock struct {
	storage map[string]string
}

func NewFileReaderMock(storage map[string]string) *FileReaderMock {
	return &FileReaderMock{storage: storage}
}

func (f FileReaderMock) Read(path string) ([]byte, error) {
	value, ok := f.storage[path]
	if !ok {
		return nil, internal.NewError("No such path")
	}
	return []byte(value), nil
}

func (f FileReaderMock) ReadString(path string) (string, error) {
	value, ok := f.storage[path]
	if !ok {
		return "", internal.NewError("No such path")
	}
	return value, nil
}
