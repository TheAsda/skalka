package settings_compiler

import (
	"github.com/TheAsda/skalka/internal/fs"
	"github.com/TheAsda/skalka/pkg/settings"
	"testing"
)

func TestSettingsCompiler_ReadSettings(t *testing.T) {
	paths := map[string]string{
		"/path/config.yml": `
log_level: info
`,
	}
	exampleSettings := settings.Settings{LogLevel: settings.Info}

	t.Run("read Settings", func(t *testing.T) {
		compiler := NewSettingsCompiler("/path/config.yml", fs.NewFileReaderMock(paths))
		s, err := compiler.ReadSettings()
		if err != nil {
			t.Fatalf("Error on reading Settings: %s", err.Error())
		}
		if s.LogLevel != exampleSettings.LogLevel {
			t.Fatalf("Got wrong log level '%s'", s.LogLevel)
		}
	})

	t.Run("path does not exist", func(t *testing.T) {
		compiler := NewSettingsCompiler("/path/random.yml", fs.NewFileReaderMock(paths))
		_, err := compiler.ReadSettings()
		if err == nil {
			t.Fatalf("Error did not return")
		}
	})
}
