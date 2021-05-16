package progress_logger

import "testing"

func TestStepLogger_LogStep(t *testing.T) {
	mock := NewIOMock()
	logger := NewStepLogger(mock, mock, "test", 5)

	t.Run("step", func(t *testing.T) {
		err := logger.LogStep("Do something")
		if err != nil {
			t.Errorf("Logging error: %s", err.Error())
		}
		if mock.writerCalls != 1 {
			t.Errorf("Write is not called once")
		}
	})
}
