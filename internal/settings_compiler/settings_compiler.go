package settings_compiler

import (
	"github.com/TheAsda/skalka/internal/fs"
	"github.com/TheAsda/skalka/pkg/settings"
	"gopkg.in/yaml.v2"
)

type SettingsCompiler struct {
	Settings     settings.Settings
	settingsPath string
	reader       fs.PathReader
}

func NewSettingsCompiler(settingsPath string, reader fs.PathReader) *SettingsCompiler {
	return &SettingsCompiler{settingsPath: settingsPath, reader: reader}
}

func (c SettingsCompiler) ReadSettings() (settings.Settings, error) {
	data, err := c.reader.Read(c.settingsPath)
	if err != nil {
		return settings.Settings{}, err
	}
	err = yaml.Unmarshal(data, &c.Settings)
	if err != nil {
		return settings.Settings{}, err
	}
	return c.Settings, nil
}
