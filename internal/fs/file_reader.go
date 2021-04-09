package fs

type FileReader interface {
	Read(path string) ([]byte, error)
	ReadString(path string) (string, error)
}
