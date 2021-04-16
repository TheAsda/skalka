package progress_logger

//func TestProgressLogger_Error(t *testing.T) {
//	mock := NewIOMock()
//	logger := NewProgressLogger(mock, mock)
//
//	t.Run("error", func(t *testing.T) {
//		err := logger.Error("Install something")
//		if err != nil {
//			t.Errorf("Logging error: %s", err.Error())
//		}
//		if mock.writerCalls != 1 {
//			t.Errorf("Write is not called once")
//		}
//	})
//}
//
//func TestProgressLogger_Warn(t *testing.T) {
//	mock := NewIOMock()
//	logger := NewProgressLogger(mock, mock)
//
//	t.Run("warn", func(t *testing.T) {
//		err := logger.Warn("Install something")
//		if err != nil {
//			t.Errorf("Logging error: %s", err.Error())
//		}
//		if mock.writerCalls != 1 {
//			t.Errorf("Write is not called once")
//		}
//	})
//}
//
//func TestProgressLogger_Info(t *testing.T) {
//	mock := NewIOMock()
//	logger := NewProgressLogger(mock, mock)
//
//	t.Run("info", func(t *testing.T) {
//		err := logger.Info("Install something")
//		if err != nil {
//			t.Errorf("Logging error: %s", err.Error())
//		}
//		if mock.writerCalls != 1 {
//			t.Errorf("Write is not called once")
//		}
//	})
//}
