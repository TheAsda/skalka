package fs

import "github.com/TheAsda/skalka/internal"

type PathReaderMock struct {
	storage map[string]string
}

func NewFileReaderMock(storage map[string]string) *PathReaderMock {
	return &PathReaderMock{storage: storage}
}

func (f PathReaderMock) Read(path string) ([]byte, error) {
	value, ok := f.storage[path]
	if !ok {
		return nil, internal.NewError("No such path")
	}
	return []byte(value), nil
}

func (f PathReaderMock) ReadString(path string) (string, error) {
	value, ok := f.storage[path]
	if !ok {
		return "", internal.NewError("No such path")
	}
	return value, nil
}
