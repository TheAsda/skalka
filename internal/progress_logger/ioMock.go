package progress_logger

type IOMock struct {
	writerCalls int
	readerCalls int
}

func NewIOMock() *IOMock {
	return &IOMock{
		writerCalls: 0,
		readerCalls: 0,
	}
}

func (m *IOMock) Read(p []byte) (n int, err error) {
	m.readerCalls += 1
	return len(p), nil
}

func (m *IOMock) Write(p []byte) (n int, err error) {
	m.writerCalls += 1
	return len(p), nil
}

func (m *IOMock) ResetMocks() {
	m.readerCalls = 0
	m.writerCalls = 0
}
