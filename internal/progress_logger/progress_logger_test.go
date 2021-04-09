package progress_logger

import "testing"

type Mock struct {
	writerCalls int
	readerCalls int
}

func NewMock() *Mock {
	return &Mock{
		writerCalls: 0,
		readerCalls: 0,
	}
}

func (m Mock) Read(p []byte) (n int, err error) {
	m.readerCalls += 1
	return 0, nil
}

func (m Mock) Write(p []byte) (n int, err error) {
	m.writerCalls += 1
	return 0, nil
}

func (m Mock) ResetMocks() {
	m.readerCalls = 0
	m.writerCalls = 0
}

func TestProgressLogger_LogStep(t *testing.T) {
	mock := NewMock()
	logger := NewProgressLogger(*mock, *mock, "test", 5)

	t.Run("log step", func(t *testing.T) {
		err := logger.LogStep("Install something")
		if err != nil {
			t.Errorf("Logging error: %s", err.Error())
		}
		if mock.writerCalls == 1 {
			t.Errorf("Write called more than one time")
		}
	})
}
