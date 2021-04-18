package fs

type PathReader interface {
	Read(path string) ([]byte, error)
}